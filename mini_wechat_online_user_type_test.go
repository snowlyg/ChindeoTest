package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatOnlineUserTypeListSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/online/v1/user_type").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("application_id", model.AppId).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(1)
	obj.Value("data").Array().First().Object().Value("id").Equal(UserType.ID)
	obj.Value("data").Array().First().Object().Value("utp_desc").Equal(UserType.UtpDesc)
}
