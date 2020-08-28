package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestRefreshTokenSuccess(t *testing.T) {
	obj := model.GetE(t).POST("/api/v1/refresh_access_token").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("刷新成功")
	obj.Value("data").Object().Value("AccessToken").NotNull()

	token := obj.Value("data").Object().Value("AccessToken").Raw()
	data, ok := token.(string)
	if ok {
		model.Token = data
	}
}

func TestRefreshTokenError(t *testing.T) {
	obj := model.GetE(t).POST("/api/v1/refresh_access_token").
		WithHeaders(model.GetHeaderNoToken()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(401)
	obj.Value("message").String().Equal("错误 token 结构")
	obj.Value("data").Null()
}
