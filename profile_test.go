package main

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestProfileSuccess(t *testing.T) {
	e := httpexpect.New(t, BaseUrl)
	obj := e.GET("/api/v1/profile").
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Object().Value("id").Equal(15)
}

func TestProfileUpdateSuccess(t *testing.T) {
	info := map[string]interface{}{
		"realname":   "小样",
		"id_card_no": "430923198901156623",
	}

	e := httpexpect.New(t, BaseUrl)
	obj := e.POST("/api/v1/profile/update").
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(info).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("操作成功")
	obj.Value("data").Object().Value("id").Equal(15)
	obj.Value("data").Object().Value("id_card_no").Equal("430923198901156623")
	obj.Value("data").Object().Value("realname").Equal("小样")
}

func TestProfileNoDevHeader(t *testing.T) {
	e := httpexpect.New(t, BaseUrl)
	obj := e.GET("/api/v1/profile").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("内部服务类型,无法访问外部接口类型")
}
