package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"strconv"
	"testing"
)

var cartShop *model.ShopSpus

func TestMiniWechatShopCartAddSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	cartShop = model.CreateSpu(brand.ID, Cate1.ID, 2, "这是一个很神奇的商品", "这是一个很神奇的商品的超厉害的副标题", 10.00, 100.00, Spec)
	ids := map[string]interface{}{
		"sku_id": cartShop.Skus[0].ID,
	}

	obj := model.GetE(t).POST("/shop/v1/cart/add").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(ids).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestMiniWechatShopCartAddAgainiSuccess(t *testing.T) {
	ids := map[string]interface{}{
		"sku_id": cartShop.Skus[0].ID,
	}

	obj := model.GetE(t).POST("/shop/v1/cart/add").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(ids).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestMiniWechatShopCartSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/cart").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	strconv.FormatInt(int64(cartShop.Skus[0].ID), 10)
	datail := obj.Value("data").Object().Value("detail").Object().Value(strconv.FormatInt(int64(cartShop.Skus[0].ID), 10))
	datail.Array().Length().Equal(1)
	datail.Array().First().Object().Value("id").Equal(cartShop.Skus[0].ID)
	datail.Array().First().Object().Value("quantity").Equal(2)
}

func TestMiniWechatShopCartSubSuccess(t *testing.T) {
	ids := map[string]interface{}{
		"sku_id": cartShop.Skus[0].ID,
	}
	obj := model.GetE(t).POST("/shop/v1/cart/sub").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(ids).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestMiniWechatShopCartAfterSubSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/cart").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	strconv.FormatInt(int64(cartShop.Skus[0].ID), 10)
	datail := obj.Value("data").Object().Value("detail").Object().Value(strconv.FormatInt(int64(cartShop.Skus[0].ID), 10))
	datail.Array().Length().Equal(1)
	datail.Array().First().Object().Value("id").Equal(cartShop.Skus[0].ID)
	datail.Array().First().Object().Value("quantity").Equal(1)
}

func TestMiniWechatShopCartAddThereSuccess(t *testing.T) {
	ids := map[string]interface{}{
		"sku_id": cartShop.Skus[1].ID,
	}

	obj := model.GetE(t).POST("/shop/v1/cart/add").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(ids).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestMiniWechatShopCartClearSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/shop/v1/cart/clear").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestMiniWechatShopCartAfterSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/cart").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("detail").Array().Empty()
}
