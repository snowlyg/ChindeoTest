package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatProfileSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/profile").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Object().Keys().ContainsOnly(
		"id",
		"username",
		"nickname",
		"phone",
		"email",
		"sex",
		"password",
		"status",
		"avatar_url",
		"open_id",
		"union_id",
		"country",
		"province",
		"city",
		"mac",
		"create_at",
		"update_at",
		"id_card_no",
		"is_auth",
		"realname",
		"area",
		"birthday",
		"addrs",
		"warn_times",
		"collects",
		"cares",
		"profile",
		"qual",
		"patients",
		"families",
		"vitals",
		"companies",
		"mini_apps",
		"visit_spus",
	)
	obj.Value("data").Object().Value("id").Equal(15)
}

func TestMiniWechatProfileWithCompaniesRelationSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/profile").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("relation", "companies").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Object().Keys().ContainsOnly(
		"id",
		"username",
		"nickname",
		"phone",
		"email",
		"sex", "password",
		"status",
		"avatar_url",
		"open_id",
		"union_id",
		"country",
		"province",
		"city",
		"mac",
		"create_at",
		"update_at",
		"id_card_no",
		"is_auth",
		"realname",
		"companies",
		"area",
		"birthday",
	)
	obj.Value("data").Object().Value("id").Equal(15)
}

func TestMiniWechatProfileUpdateSuccess(t *testing.T) {
	info := map[string]interface{}{
		"realname":   "小样",
		"id_card_no": model.IdCardNo,
	}

	obj := model.GetE(t).POST("/api/v1/profile/update").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(info).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("操作成功")
	obj.Value("data").Object().Value("id").Equal(15)
	obj.Value("data").Object().Value("id_card_no").Equal(model.IdCardNo)
	obj.Value("data").Object().Value("realname").Equal("小样")
}

func TestMiniWechatProfileSetServerSuccess(t *testing.T) {
	info := map[string]interface{}{
		"server_date":     "1,2,3,4,5,6,7",
		"server_start_at": "9:00",
		"server_end_at":   "17:00",
	}

	obj := model.GetE(t).POST("/api/v1/profile/set_server").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(info).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("新增成功")
}

func TestMiniWechatProfileNoDevHeader(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/profile").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("内部服务类型,无法访问外部接口类型")
}

func TestMiniWechatProfileSetServerOpenSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/profile/set_server_open").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("设置成功")
}

func TestMiniWechatProfileAuthSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/profile/auth").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("设置成功")
}

func TestMiniWechatProfileSubSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/profile/sub").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("订阅成功")
}

func TestMiniWechatProfileUnSubSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/profile/unsub").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("取消订阅成功")
}
