package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestShopCateSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/inner/cate").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(model.ParentCateCount)

	data := obj.Value("data").Array().Last().Object()
	data.Value("id").Equal(Cate1.ID)

	data.Value("children").Array().Length().Equal(1)
	data.Value("children").Array().First().Object().Value("id").Equal(Cate2.ID)

	data.Value("children").Array().First().Object().Value("children").Array().Length().Equal(1)
	data.Value("children").Array().First().Object().Value("children").Array().First().Object().Value("id").Equal(Cate3.ID)
}
