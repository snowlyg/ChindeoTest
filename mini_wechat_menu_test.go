package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatMenuSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/menu").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("application_id", model.AppId).
		WithQuery("menu_type_id", Menu.MenuType.ID).
		WithQuery("time_type", Menu.TimeType).
		WithQuery("menu_tag_id", Menu.MenuTags[0].ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.MenuCount)
	obj.Value("data").Object().Value("data").Array().First().Object().Value("id").Equal(Menu.ID)
}

func TestMiniWechatMenuNoPageSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/menu").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("application_id", model.AppId).
		WithQuery("menu_type_id", Menu.MenuType.ID).
		WithQuery("time_type", Menu.TimeType).
		WithQuery("menu_tag_id", Menu.MenuTags[0].ID).
		WithQuery("page_size", "-1").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(model.MenuCount)
	obj.Value("data").Array().First().Object().Keys().ContainsOnly(
		"cover",
		"pics",
		"is_deleted",
		"name",
		"amount",
		"application_id",
		"menu_type_id",
		"tags",
		"status",
		"sort",
		"type",
		"id",
		"desc",
		"price",
		"create_at",
		"update_at",
		"pivot",
		"time_type",
		"users",
	)
	obj.Value("data").Array().First().Object().Value("id").Equal(Menu.ID)
}

func TestMiniWechatMenuNoPageNoTagSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/menu").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("application_id", model.AppId).
		WithQuery("page_size", "-1").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(model.MenuNoTagCount)
	obj.Value("data").Array().Last().Object().Value("id").Equal(Menu.ID)
}

func TestMiniWechatMenuShowSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/menu/{id}", Menu.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("application_id", model.AppId).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(Menu.ID)
	obj.Value("data").Object().Value("name").Equal(Menu.Name)
	obj.Value("data").Object().Value("time_type").Equal(Menu.TimeType.String())
	obj.Value("data").Object().Value("menu_type_id").Equal(Menu.MenuTypeID)
	obj.Value("data").Object().Value("desc").Equal(Menu.Desc)
	obj.Value("data").Object().Value("amount").Equal(Menu.Amount)
	obj.Value("data").Object().Value("price").Equal("10.00")
	obj.Value("data").Object().Value("cover").Equal(Menu.Cover)
	obj.Value("data").Object().Value("pics").Array().Length().Equal(0)
	obj.Value("data").Object().Value("create_at").String().Contains(Menu.CreateAt.Format("2006-01-02 15:04"))
	obj.Value("data").Object().Value("create_at").String().Contains(Menu.UpdateAt.Format("2006-01-02 15:04"))

	menuType := obj.Value("data").Object().Value("type").Object()
	menuType.Value("id").Equal(Menu.MenuType.ID)
	menuType.Value("name").Equal(Menu.MenuType.Name)

	obj.Value("data").Object().Value("tags").Array().Length().Equal(len(Menu.MenuTags))
	menuTag := obj.Value("data").Object().Value("tags").Array().First().Object()
	menuTag.Value("id").Equal(Menu.MenuTags[0].ID)
	menuTag.Value("name").Equal(Menu.MenuTags[0].Name)
}

func TestMiniWechatMenuCollectAddSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/collect/add/{id}", Menu.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("收藏成功")
}

func TestMiniWechatMenuCollectSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/collect").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().NotEmpty()
}

func TestMiniWechatMenuCollectCancelSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/collect/cancel/{id}", Menu.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("取消成功")
}
