package integrationtests_test

import (
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"

	"gitlab.com/soy-app/stock-api/domain/entity"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"

	"gitlab.com/soy-app/stock-api/domain/entconst"

	"gitlab.com/soy-app/stock-api/api/schema"
)

// GET api/user/me
func TestFindMe(t *testing.T) {
	t.Parallel()

	tx := testDB.Begin()
	// 初期状態の設定
	user := &entity.User{
		UserID:         "01GPZSDXPPH0FFSVNR2R9NY5TC",
		Email:          "user@digeon.co",
		HashedPassword: "JDJhJDEwJExyeTMuSmI5OWxKbmM0WXVBVUFBTWVNazl3SC8wMGJKWm0vTjJMVlN6UExqRkQ1WEdxMjhl",
		UserType:       entconst.User,
		IsDeleted:      false,
	}
	if err := tx.Create(user).Error; err != nil {
		t.Fatal(err)
	}
	// テストケース
	tests := []struct {
		name        string
		successCase bool
		wantCode    int
		request     func(findMeRes *schema.UserRes, errRes *echo.HTTPError) (resp *resty.Response, err error)
		wantRes     *schema.UserRes
		wantErr     *echo.HTTPError
	}{
		{
			name:        "成功",
			successCase: true,
			wantCode:    200,
			request: func(findMeRes *schema.UserRes, errRes *echo.HTTPError) (resp *resty.Response, err error) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				s, c := NewTestServer(t, tx, ctrl)
				token := Login(t, user.Email, "pass", c, s)
				resp, err = c.R().
					SetHeader("Authorization", "Bearer"+" "+token).
					SetResult(findMeRes).
					Get(s.URL + "/api/user/me")
				return
			},
			wantRes: &schema.UserRes{
				UserId:   "01GPZSDXPPH0FFSVNR2R9NY5TC",
				Email:    "user@digeon.co",
				UserType: "User",
			},
			wantErr: nil,
		},
		{
			name:        "失敗: ログインしていない",
			successCase: false,
			wantCode:    401,
			request: func(findMeRes *schema.UserRes, errRes *echo.HTTPError) (resp *resty.Response, err error) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				s, c := NewTestServer(t, tx, ctrl)
				resp, err = c.R().
					SetHeader("Authorization", ""). // Authorizationヘッダが空
					SetError(errRes).
					Get(s.URL + "/api/user/me")
				return
			},
			wantRes: nil,
			wantErr: echo.NewHTTPError(http.StatusUnauthorized),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := &schema.UserRes{}
			errRes := &echo.HTTPError{}
			resp, err := tt.request(res, errRes)
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode() != tt.wantCode {
				t.Errorf("Status Code %d, want = %d", resp.StatusCode(), tt.wantCode)
			}
			if tt.successCase && !cmp.Equal(res, tt.wantRes) {
				t.Errorf("FindMe Response diff =%v", cmp.Diff(res, tt.wantRes))
				return
			}
			if diff := cmp.Diff(errRes, tt.wantErr, cmpopts.IgnoreFields(*errRes, "Code")); !tt.successCase && diff != "" {
				t.Errorf("FindMe Response diff =%v", diff)
			}
		})
	}
}

// GET api/user/me
// TODO SearchUserのテストを書く
// searchQueryはbindの仕様から、必ず0値が入る。
// 0値での検索は許容するため、handlerでの単体テストでは検証できないので、integrationテストで検証する。

// TODO 各APIの権限回りを検証できるテストを書く
