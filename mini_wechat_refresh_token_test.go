package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatRefreshTokenSuccess(t *testing.T) {
	obj := model.GetE(t).POST("/api/v1/refresh_access_token").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("刷新成功")
	obj.Value("data").Object().Value("AccessToken").NotNull()

	token := obj.Value("data").Object().Value("AccessToken").Raw()
	data, ok := token.(string)
	if ok {
		model.MiniWechatToken = data
	}
}

func TestMiniWechatRefreshTokenError(t *testing.T) {

	obj := model.GetE(t).POST("/api/v1/refresh_access_token").
		WithHeaders(model.GetMiniHeaderNoToken("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(401)
	obj.Value("message").String().Equal("错误 token 结构")
	obj.Value("data").Null()
}
