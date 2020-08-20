package main

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestRefreshTokenSuccess(t *testing.T) {

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.POST("/api/v1/refresh_access_token").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("刷新成功")
	obj.Value("data").Object().Value("AccessToken").NotNull()

	token := obj.Value("data").Object().Value("AccessToken").Raw()
	data, ok := token.(string)
	if ok {
		Token = data
	}
}

func TestRefreshTokenError(t *testing.T) {

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.POST("/api/v1/refresh_access_token").
		WithHeaders(map[string]string{"X-Token": "", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(401)
	obj.Value("message").String().Equal("请登录！")
	obj.Value("data").Equal("")
}
