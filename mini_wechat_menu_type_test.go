package main

import (
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestMiniWechatMenuTypeSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/api/v1/menu_type").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4", "IsDev": "1"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithQuery("application_id", "3").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}
