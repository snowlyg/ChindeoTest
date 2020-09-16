package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

var collectSpu *model.ShopSpus

func TestMiniWechatSpuCollectAddSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	collectSpu = model.CreateSpu(brand.ID, Cate2.ID, 1, "这是一个很神奇的商品", "这是一个很神奇的商品的超厉害的副标题", 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/collect/add/{id}", collectSpu.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("收藏成功")
}

func TestMiniWechatSpuCollectSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/collect").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(1)
}

func TestMiniWechatSpuCollectCancelSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/collect/cancel/{id}", collectSpu.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("取消成功")
}
