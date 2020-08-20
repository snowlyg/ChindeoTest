package main

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var menuId float64

func TestMenuSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})

	obj := e.GET("/api/v1/menu").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithQuery("menu_type_id", "3").
		WithQuery("time_type", "3").
		WithQuery("menu_tag_id", "3").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
}

func TestMenuNoPageSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})

	obj := e.GET("/api/v1/menu?page_size=-1").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	id := obj.Value("data").Array().First().Object().Value("id").Raw()
	data, ok := id.(float64)
	if ok {
		menuId = data
	}
}

func TestMenuShowSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})

	obj := e.GET("/api/v1/menu/{id}", menuId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(menuId)
	obj.Value("data").Object().Keys().ContainsOnly("id", "name", "menu_type_id", "time_type", "desc", "amount", "price", "cover", "pics", "create_at", "update_at", "type", "tags")
}
