package main

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestVersionSuccess(t *testing.T) {
	e := httpexpect.New(t, BaseUrl)
	obj := e.GET("/api/v1/version").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("version").Equal("V0.0.1")
}
