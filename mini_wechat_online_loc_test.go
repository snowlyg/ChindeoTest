package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatOnlineLocListSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/online/v1/loc").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("application_id", model.AppId).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(model.LocCount)
	obj.Value("data").Array().First().Object().Value("id").Equal(Loc.ID)
	obj.Value("data").Array().First().Object().Value("loc_desc").Equal(Loc.LocDesc)
}

func TestMiniWechatOnlineLocListError(t *testing.T) {
	obj := model.GetE(t).GET("/online/v1/loc").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("医院商家不能为空")

}
