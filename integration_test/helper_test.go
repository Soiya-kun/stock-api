/*
func TestSample(t *testing.T) {
	t.Parallel()

	// テスト用サーバー
	testServer, tx, closeAndRollback := NewTestServer(t)

	// テスト終了時に呼ばれる
	t.Cleanup(func() {
		// テスト用サーバーのClose & NewTestServer()以降に加えられたDBの変更をすべてロールバックする
		closeAndRollback()
	})

	// HTTPクライアント
	restyClient := resty.New()

	// 初期状態の設定
	// DBの中身は空なのでここでテストに必要な環境を個別に作成する
	// レコードの登録等にはNewTestServer()から返されるtxが使用できる。
	// あるいは r := repository.NewXXXRepository(tx) としてリポジトリ構造体を作成し、メソッドを呼ぶことも可能。
	// 統合テストなのでエンドポイントを順番に叩く(例: ユーザーの作成 -> ユーザーの詳細を編集 → ユーザーの取得)ようにしないと不整合なデータが作成できるので、内部の実装の一部のみを呼ばない方が良いかもしれないが、それだと大変なのでtxを返している。
	// あらかじめテスト用データを流し込んでおくということも考えられるが、他のテストと干渉しないようにする必要がある。
	user := &model.User{
		UserID:         "01GPZSDXPPH0FFSVNR2R9NY5TC",
		Email:          "user0of0@test.com",
		HashedPassword: "JDJhJDEwJExyeTMuSmI5OWxKbmM0WXVBVUFBTWVNazl3SC8wMGJKWm0vTjJMVlN6UExqRkQ1WEdxMjhl",
	}
	// ユーザーの作成
	if err := tx.Create(user).Error; err != nil {
		t.Fatal(err)
	}
	loginRes := &schema.LoginRes{}
	// ログイン用エンドポイントを先に叩いてログインしておく
	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"email": "user0of0@test.com", "password": "pass"}`).
		SetResult(loginRes).
		Post(testServer.URL + "/api/auth/access-token")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode() != 200 {
		t.Errorf("Status Code %d, want = %d", resp.StatusCode(), 200)
	}
	token := loginRes.AccessToken

	// 実際にテストしたいエンドポイント
	findMeRes := &schema.UserRes{}
	resp, err = restyClient.R().
		SetHeader("Authorization", "Bearer"+" "+token).
		SetResult(findMeRes).
		Get(testServer.URL + "/api/user/me")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode() != 200 {
		t.Errorf("Status Code %d, want = %d", resp.StatusCode(), 200)
	}

	// 理想のレスポンス
	wantRes := &schema.UserRes{
		UserID:        "01GPZSDXPPH0FFSVNR2R9NY5TC",
		Email:         "user0of0@test.com",
		Groups:        []schema.GroupRes{},
		IsSystemAdmin: false,
	}

	// アサーション
	if !cmp.Equal(findMeRes, wantRes) {
		t.Errorf("FindMe Response diff =%v", cmp.Diff(findMeRes, wantRes))
	}
}
*/

package integrationtests_test

import (
	"fmt"
	l "log"
	"math"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"gitlab.com/soy-app/stock-api/usecase/port"

	"github.com/go-resty/resty/v2"

	"gitlab.com/soy-app/stock-api/adapter/authentication"
	"gitlab.com/soy-app/stock-api/adapter/database"
	"gitlab.com/soy-app/stock-api/api/router"
	"gitlab.com/soy-app/stock-api/config"
	"gitlab.com/soy-app/stock-api/interface/repository"
	"gitlab.com/soy-app/stock-api/log"
	"gitlab.com/soy-app/stock-api/usecase/interactor"

	"gorm.io/gorm"
)

var (
	testDB *gorm.DB
)

// 初回読み込み時にテスト用DBと接続し、Migrationを実行
func init() {
	if config.IsTest() {
		if os.Getenv("DB_USER") != "template-test" {
			panic("DB_USER is not template-test")
		}
		if os.Getenv("DB_HOST") != "template-mysql-test" {
			panic("DB_HOST is not template-mysql-test")
		}
		if os.Getenv("DB_PORT") != "13306" {
			panic("DB_PORT is not 13306")
		}
		if os.Getenv("DB_NAME") != "template-test" {
			panic("DB_NAME is not template-test")
		}
	} else if config.IsGitLabCI() {
		if os.Getenv("DB_USER") != "root" {
			panic("DB_USER is not root")
		}
		if os.Getenv("DB_HOST") != "mysql" {
			panic("DB_HOST is not mysql")
		}
		if os.Getenv("DB_PORT") != "3306" {
			panic("DB_PORT is not 3306")
		}
		if os.Getenv("DB_NAME") != "template-test" {
			panic("DB_NAME is not template-test")
		}
	} else {
		panic("Invalid ENV")
	}

	var err error
	logger, _ := log.NewLogger()

	// MySQLの起動を待つ
	const retryMax = 10
	for i := 0; i < retryMax; i++ {
		testDB, err = database.NewDB(logger)
		if err == nil {
			break
		}
		if i == retryMax-1 {
			panic(fmt.Sprintf("failed to connect to database: %v", err))
		}
		duration := time.Millisecond * time.Duration(math.Pow(1.5, float64(i))*1000)
		l.Printf("failed to connect to database retrying: %v: %v\n", err, duration)
		time.Sleep(duration)
	}

	if err := database.Migrate(testDB); err != nil {
		panic(err)
	}
}

func NewTestServer(
	t *testing.T,
	tx *gorm.DB,
	f func() (port.ULID, port.Email),
) (s *httptest.Server, restyClient *resty.Client) {
	t.Helper()

	ulidDriver, email := f()
	userRepo := repository.NewUserRepository(tx, ulidDriver)
	userAuth := authentication.NewUserAuth()
	userUC := interactor.NewUserUseCase(email, ulidDriver, userAuth, userRepo)

	stockRepo := repository.NewStockRepository(tx)
	searchStockPatternRepo := repository.NewSearchStockPatternRepository(tx)
	searchedStockRepo := repository.NewSearchedStockPatternRepository(tx)
	stockUC := interactor.NewStockUseCase(ulidDriver, stockRepo, searchStockPatternRepo, searchedStockRepo)

	e := router.NewServer(
		userUC,
		stockUC,
	)

	s = httptest.NewServer(e.Server.Handler)

	t.Cleanup(func() {
		s.Close()
		tx.Rollback()
	})

	// HTTPクライアント
	restyClient = resty.New()
	return
}
