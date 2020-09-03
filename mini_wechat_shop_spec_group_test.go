package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatShopSpecGroupSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/spec_group").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(model.SpecGroupCount)
	obj.Value("data").Array().First().Object().Value("id").Equal(SpecGroup.ID)
	obj.Value("data").Array().First().Object().Value("name").Equal(SpecGroup.Name)
	obj.Value("data").Array().First().Object().Value("spec_params").Array().Length().Equal(model.SpecParamCount)
}
