package email

import (
	"gitlab.com/soy-app/stock-api/config"
)

func ContentToResetPassword(token string) (subject, body string) {
	subject = "【システム】パスワード再設定のお知らせ"
	body = "下記のURLにアクセスしてパスワードを再設定してください。\n" +
		config.FrontendURL() + "/password-change?token=" + token
	return
}
