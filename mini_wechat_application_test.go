package main

import (
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var applicationId float64

func TestMiniWechatApplicationSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/api/v1/outline/application/0").
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().NotEmpty()

	id := obj.Value("data").Array().First().Object().Value("id").Raw()
	data, ok := id.(float64)
	if ok {
		applicationId = data
	}

}
