package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

var showSpu *model.ShopSpus

func TestMiniWechatShopSpuSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/spu").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(model.SpuCount)
}

func TestMiniWechatShopSpuWithBrandIdSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	model.CreateSpu(brand.ID, Cate1.ID, 1, "", "", 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/spu").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("brand_id", Brand.ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(1)
}

func TestMiniWechatShopSpuWithCateIdSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	model.CreateSpu(brand.ID, Cate2.ID, 1, "", "", 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/spu").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("cate_id", Cate2.ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(1)
}

func TestMiniWechatShopSpuWithCateIdAndBrandSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	model.CreateSpu(brand.ID, Cate2.ID, 1, "", "", 10.00, 100.00, Spec)
	model.CreateSpu(brand.ID, Cate1.ID, 1, "", "", 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/spu").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("cate_id", Cate2.ID).
		WithQuery("brand_id", brand.ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(1)
}

func TestMiniWechatShopSpuWithKeyWordSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	model.CreateSpu(brand.ID, Cate1.ID, 1, "这是一个很牛逼的商品", "这是一个很牛逼的商品的超厉害的副标题", 10.00, 100.00, Spec)
	model.CreateSpu(brand.ID, Cate1.ID, 1, "很牛逼的商品", "很牛逼的商品的超厉害的副标题", 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/spu").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("key_word", "牛逼").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(2)
}

func TestMiniWechatShopSpuWithCateIdAndBrandAndKeyWordSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	model.CreateSpu(brand.ID, Cate2.ID, 1, "这是一个很神奇的商品", "这是一个很神奇的商品的超厉害的副标题", 10.00, 100.00, Spec)
	model.CreateSpu(brand.ID, Cate1.ID, 1, "这是一个很神奇的商品", "这是一个很神奇的商品的超厉害的副标题", 10.00, 100.00, Spec)
	model.CreateSpu(brand.ID, Cate1.ID, 1, "很神奇的商品", "很神奇的商品的超厉害的副标题", 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/spu").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("key_word", "神奇").
		WithQuery("cate_id", Cate2.ID).
		WithQuery("brand_id", brand.ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(1)
}

func TestMiniWechatShopSpuShowSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	name := "这是一个很神奇的商品"
	title := "这是一个很神奇的商品的超厉害的副标题"
	showSpu = model.CreateSpu(brand.ID, Cate1.ID, 3, name, title, 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/spu/{id}", showSpu.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Object().Value("id").Equal(showSpu.ID)
	obj.Value("data").Object().Value("name").Equal(name)
	obj.Value("data").Object().Value("sub_title").Equal(title)
	obj.Value("data").Object().Value("specs").Array().Length().Equal(0)
	obj.Value("data").Object().Value("detail").Object().Value("id").Equal(showSpu.Detail.ID)
	obj.Value("data").Object().Value("detail").Object().Value("description").Equal(showSpu.Detail.Description)
	obj.Value("data").Object().Value("detail").Object().Value("packing_list").Equal(showSpu.Detail.PackingList)
	obj.Value("data").Object().Value("detail").Object().Value("after_service").Equal(showSpu.Detail.AfterService)
	obj.Value("data").Object().Value("cates").Array().Length().Equal(1)
	obj.Value("data").Object().Value("cates").Array().First().Object().Value("id").Equal(Cate1.ID)
	obj.Value("data").Object().Value("brand").Object().Value("id").Equal(brand.ID)
	obj.Value("data").Object().Value("skus").Array().Length().Equal(3)
}

func TestMiniWechatShopSpuVisitSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/visit").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(1)
}

func TestMiniWechatShopSpuVisitCancelSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/visit/cancel/{id}", showSpu.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("取消成功")
}
