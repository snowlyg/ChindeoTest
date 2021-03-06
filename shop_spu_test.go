package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestShopSpuSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/inner/spu").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("page_size", -1).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(model.SpuCount)
}
func TestShopSpuPaginateSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/inner/spu").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	spuCount := model.SpuCount
	if spuCount > 10 {
		spuCount = 10
	}
	obj.Value("data").Object().Value("data").Array().Length().Equal(spuCount)
}

func TestShopSpuSortByPriceAscSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	shopSpu := model.CreateSpu(brand.ID, Cate1.ID, 1, "", "", 1.00, 10.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/inner/spu").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("order_by", "min_price").
		WithQuery("sort", "asc").
		WithQuery("page_size", -1).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(model.SpuCount)
	obj.Value("data").Array().First().Object().Value("id").Equal(shopSpu.ID)
	model.DelShopSpu(shopSpu)

}

func TestShopSpuSortByPriceDescSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	shopSpu := model.CreateSpu(brand.ID, Cate1.ID, 1, "", "", 500.00, 600.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/inner/spu").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("order_by", "min_price").
		WithQuery("sort", "desc").
		WithQuery("page_size", -1).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(model.SpuCount)
	obj.Value("data").Array().First().Object().Value("id").Equal(shopSpu.ID)
	model.DelShopSpu(shopSpu)
}

func TestShopSpuWithBrandIdSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	model.CreateSpu(brand.ID, Cate1.ID, 1, "", "", 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/inner/spu").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("brand_id", Brand.ID).
		WithQuery("page_size", -1).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(1)
}

func TestShopSpuWithCateIdSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	model.CreateSpu(brand.ID, Cate2.ID, 1, "", "", 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/inner/spu").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("cate_id", Cate2.ID).
		WithQuery("page_size", -1).
		Expect().
		Status(http.StatusOK).JSON().Object()

	spus := model.GetSpuByCateId(Cate2.ID)

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(len(spus))
}

func TestShopSpuWithCateIdAndBrandSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	model.CreateSpu(brand.ID, Cate2.ID, 1, "", "", 10.00, 100.00, Spec)
	model.CreateSpu(brand.ID, Cate1.ID, 1, "", "", 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/inner/spu").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("cate_id", Cate2.ID).
		WithQuery("brand_id", brand.ID).
		WithQuery("page_size", -1).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(1)
}

func TestShopSpuWithKeyWordSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	model.CreateSpu(brand.ID, Cate1.ID, 1, "这是一个很牛逼的商品", "这是一个很牛逼的商品的超厉害的副标题", 10.00, 100.00, Spec)
	model.CreateSpu(brand.ID, Cate1.ID, 1, "很牛逼的商品", "很牛逼的商品的超厉害的副标题", 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/inner/spu").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("key_word", "牛逼").
		WithQuery("page_size", -1).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(model.SpuTitleNameCount)
}

func TestShopSpuWithCateIdAndBrandAndKeyWordSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	model.CreateSpu(brand.ID, Cate2.ID, 1, "这是一个很神奇的商品", "这是一个很神奇的商品的超厉害的副标题", 10.00, 100.00, Spec)
	model.CreateSpu(brand.ID, Cate1.ID, 1, "这是一个很神奇的商品", "这是一个很神奇的商品的超厉害的副标题", 10.00, 100.00, Spec)
	model.CreateSpu(brand.ID, Cate1.ID, 1, "很神奇的商品", "很神奇的商品的超厉害的副标题", 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/inner/spu").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("key_word", "神奇").
		WithQuery("cate_id", Cate2.ID).
		WithQuery("brand_id", brand.ID).
		WithQuery("page_size", -1).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(1)
}

func TestShopSpuShowSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	name := "这是一个很神奇的商品"
	title := "这是一个很神奇的商品的超厉害的副标题"
	spu := model.CreateSpu(brand.ID, Cate1.ID, 3, name, title, 10.00, 100.00, Spec)
	obj := model.GetE(t).GET("/shop/v1/inner/spu/{id}", spu.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Object().Value("id").Equal(spu.ID)
	obj.Value("data").Object().Value("name").Equal(name)
	obj.Value("data").Object().Value("sub_title").Equal(title)
	obj.Value("data").Object().Value("specs").Array().Length().Equal(0)
	obj.Value("data").Object().Value("detail").Object().Value("id").Equal(spu.Detail.ID)
	obj.Value("data").Object().Value("detail").Object().Value("description").Equal(spu.Detail.Description)
	obj.Value("data").Object().Value("detail").Object().Value("packing_list").Equal(spu.Detail.PackingList)
	obj.Value("data").Object().Value("detail").Object().Value("after_service").Equal(spu.Detail.AfterService)
	obj.Value("data").Object().Value("cates").Array().Length().Equal(1)
	obj.Value("data").Object().Value("cates").Array().First().Object().Value("id").Equal(Cate1.ID)
	obj.Value("data").Object().Value("brand").Object().Value("id").Equal(brand.ID)
	obj.Value("data").Object().Value("skus").Array().Length().Equal(3)
}
