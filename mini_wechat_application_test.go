package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatApplicationSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/application/482024").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(2)

	firstId := obj.Value("data").Array().First().Object().Value("id")
	firstId.Equal(model.AppId)

}
