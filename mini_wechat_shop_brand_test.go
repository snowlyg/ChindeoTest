package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatShopBrandSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/brand").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(model.BrandCount)

}

func TestMiniWechatShopBrandWithCateSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/brand").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("cate_id", Brand.Cates[0].ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(model.BrandWithCateCount)

}
