package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatCnAreaSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/cnarea").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(34)
}
