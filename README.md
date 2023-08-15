# pro-seeds backend

Golangバージョン
```
go >= 1.19
```

1. PCの環境に合わせてdocker-composeファイルの編集を行う。  
    - M1チップのMacの場合docker-compose.deps.ymlに`platform: linux/x86_64`を追加する。
2. `make up-db`でDBを立ち上げる。  
    - Docker起動時、`network 〇〇 not found`のエラーが出た場合、`docker network ls`コマンドの実行で出てくる既存のdocker_networkの名前をnetwork_modeに設定する。
3. 環境変数の設定
    1. `cp env.example env.dev`
    2. env.devの環境変数を埋める。

4. `make run-dev`で全体を起動する。
windowsの場合は`docker-compose -f docker-compose.dev-windows.yml up -d`
5. コンテナを立ち上げた状態で`make init-db`を実行し、初期データを作成する。

## 株式データ投入

```
go run cmd/importonlocal/importonlocal.go

go run cmd/searchbycondition/searchbycondition.go
```
## テスト

テストDBと接続しない場合のテストは

```
make run-dev
```

でサーバーコンテナが起動中に

```
make test-without-db
```

を実行する。

テストDBと接続してテストする場合は、まずテスト用DBコンテナを

```
make up-test-db
```

で起動後、env.testを用意しDB_PORT=13306としたうえで、サーバーコンテナを

```
make run-for-test
```

で立ち上げ、

```
make test
```

を実行する。

## その他コマンド
```sh
# lintの実行
make lint

# フォーマットの実行
make format

# テストの実行
make test
```

### mockの生成
```
go install github.com/golang/mock/mockgen@v1.6.0

for go_file in $(find . -name "*.go" | grep -v "_test.go"); do
if grep -q "^type.*interface" $go_file; then
# $(dirname $go_file)から.を取り除く
package_name=$(dirname $go_file | sed -e "s/^\.\///")
go_file_name=$(basename $go_file)
# mockgen -source=パッケージ名/ファイル名 -destination=パッケージ名_mock/ファイル名_mock.go
mockgen -source=$package_name/$go_file_name -destination=${package_name}_mock/$go_file_name
fi
done
```
