package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatTempsCareSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/temps").
		WithHeaders(model.GetMiniHeader("406429e1284a86b92283cbb52dacda6a")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").NotNull()
}

func TestMiniWechatTempsOrderSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/temps").
		WithHeaders(model.GetMiniHeader("5205857593c2eacc6f6c1da376b32ca3")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").NotNull()
}

func TestMiniWechatTempsShopSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/temps").
		WithHeaders(model.GetMiniHeader("ebba9bb355864815eacf09598515e562")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").NotNull()

}

func TestMiniWechatTempsIndexSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/temps").
		WithHeaders(model.GetMiniHeader("ebba9bb355864815eacf09598515e562")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").NotNull()
}
