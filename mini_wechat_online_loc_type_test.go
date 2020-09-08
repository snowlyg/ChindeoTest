package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatOnlineLocTypeListSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/online/v1/loc_type").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("application_id", model.AppId).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(model.LocTypeCount)
	obj.Value("data").Array().First().Object().Value("id").Equal(LocType.ID)
	obj.Value("data").Array().First().Object().Value("name").Equal(LocType.Name)
	obj.Value("data").Array().First().Object().Value("application_id").Equal(model.AppId)
}

func TestMiniWechatOnlineLocTypeListError(t *testing.T) {
	obj := model.GetE(t).GET("/online/v1/loc_type").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("医院商家不能为空")
}
