package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestVersionSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/version").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("version").Equal("V0.0.1")
}
