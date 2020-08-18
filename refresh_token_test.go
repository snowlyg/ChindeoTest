package main

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestRefreshTokenSuccess(t *testing.T) {
	e := httpexpect.New(t, BaseUrl)
	obj := e.POST("/api/v1/refresh_access_token").
		WithHeader("X-Token", Token).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("刷新成功")
	obj.Value("data").Object().Value("AccessToken").NotNull()
}

func TestRefreshTokenError(t *testing.T) {
	e := httpexpect.New(t, BaseUrl)
	obj := e.POST("/api/v1/refresh_access_token").
		WithHeader("X-Token", "").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(401)
	obj.Value("message").String().Equal("请登录！")
	obj.Value("data").Equal("")
}
