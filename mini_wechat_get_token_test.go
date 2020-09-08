package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatGetTokenSuccess(t *testing.T) {
	auth := map[string]interface{}{
		"code":           "5205857593c2eacc6f6c1da376b32ca3",
		"iv":             "5205857593c2eacc6f6c1da376b32ca3",
		"encrypted_data": "5205857593c2eacc6f6c1da376b32ca3",
	}

	obj := model.GetE(t).POST("/api/v1/get_access_token").
		WithHeaders(model.GetMiniHeader("")).
		WithJSON(auth).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("授权成功")
	obj.Value("data").Object().Value("AccessToken").NotNull()
	token := obj.Value("data").Object().Value("AccessToken").Raw()
	data, ok := token.(string)
	if ok {
		model.MiniWechatToken = data
	}
}

func TestMiniWechatGetTokenErrorAuthType(t *testing.T) {
	auth := map[string]interface{}{
		"code":           "5205857593c2eacc6f6c1da376b32ca3",
		"iv":             "5205857593c2eacc6f6c1da376b32ca3",
		"encrypted_data": "5205857593c2eacc6f6c1da376b32ca3",
	}

	obj := model.GetE(t).POST("/api/v1/get_access_token").
		WithHeaders(model.GetMiniHeaderAuthServer("")).
		WithJSON(auth).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("app_id 不能为空！")
	obj.Value("data").Null()
}
