package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMenuTypeSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/menu_type").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(1)
	obj.Value("data").Array().First().Object().Value("id").Equal(MenuType.ID)
	obj.Value("data").Array().First().Object().Value("name").Equal(MenuType.Name)
}
