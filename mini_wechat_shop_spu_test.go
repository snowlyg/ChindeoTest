package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

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
	model.CreateSpu(brand.ID, Cate1.ID, 1, "", "", nil)
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
	model.CreateSpu(brand.ID, Cate2.ID, 1, "", "", nil)
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
	model.CreateSpu(brand.ID, Cate2.ID, 1, "", "", nil)
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
	model.CreateSpu(brand.ID, Cate2.ID, 1, "这是一个很牛逼的商品", "这是一个很牛逼的商品的超厉害的副标题", nil)
	model.CreateSpu(brand.ID, Cate2.ID, 1, "很牛逼的商品", "很牛逼的商品的超厉害的副标题", nil)
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
	model.CreateSpu(brand.ID, Cate2.ID, 1, "这是一个很神奇的商品", "这是一个很神奇的商品的超厉害的副标题", nil)
	model.CreateSpu(brand.ID, Cate2.ID, 1, "很神奇的商品", "很神奇的商品的超厉害的副标题", nil)
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
	obj.Value("data").Array().Length().Equal(2)
}

func TestMiniWechatShopSpuShowSuccess(t *testing.T) {
	spec := &model.SpecGroup{
		Name: "基本信息",
		Params: []*model.SpecParam{
			{
				Name:      "机身颜色",
				Unit:      "",
				Numeric:   false,
				Generic:   false,
				Searching: true,
				Value:     []string{"白色", "金色", "黑色"},
			}, {
				Name:      "机身长度",
				Unit:      "mm",
				Numeric:   true,
				Generic:   true,
				Searching: true,
				Value:     []string{"112"},
			}, {
				Name:      "输入方法",
				Unit:      "",
				Numeric:   false,
				Generic:   true,
				Searching: true,
				Value:     []string{"触控"},
			}, {
				Name:      "内存",
				Unit:      "G",
				Numeric:   true,
				Generic:   false,
				Searching: true,
				Value:     []string{"6", "8"},
			}, {
				Name:      "机身存储",
				Unit:      "GB",
				Numeric:   true,
				Generic:   false,
				Searching: true,
				Value:     []string{"16", "32", "64", "128"},
			},
		},
	}
	brand := model.CreateBrand(false)
	name := "这是一个很神奇的商品"
	title := "这是一个很神奇的商品的超厉害的副标题"
	spu := model.CreateSpu(brand.ID, Cate2.ID, 3, name, title, spec)
	obj := model.GetE(t).GET("/shop/v1/spu/{id}", spu.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Object().Value("id").Equal(spu.ID)
	obj.Value("data").Object().Value("name").Equal(name)
	obj.Value("data").Object().Value("sub_title").Equal(title)
	obj.Value("data").Object().Value("specs").Array().NotEmpty()
	obj.Value("data").Object().Value("detail").Object().Value("id").Equal(spu.Detail.ID)
	obj.Value("data").Object().Value("detail").Object().Value("description").Equal(spu.Detail.Description)
	obj.Value("data").Object().Value("detail").Object().Value("packing_list").Equal(spu.Detail.PackingList)
	obj.Value("data").Object().Value("detail").Object().Value("after_service").Equal(spu.Detail.AfterService)
	obj.Value("data").Object().Value("cates").Array().Length().Equal(1)
	obj.Value("data").Object().Value("cates").Array().First().Object().Value("id").Equal(Cate2.ID)
	obj.Value("data").Object().Value("brand").Object().Value("id").Equal(brand.ID)
	obj.Value("data").Object().Value("skus").Array().Length().Equal(3)
}
