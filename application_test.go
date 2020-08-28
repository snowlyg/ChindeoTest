package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"

	"testing"
)

func TestApplicationSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/application").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").NotEqual(0)
	obj.Value("data").Object().Value("name").Equal("我的医院")
	obj.Value("data").Object().Value("tel").Equal("0755-65236253")
	obj.Value("data").Object().Value("addr").Equal("我的医院")
	obj.Value("data").Object().Value("describle").Equal("我的医院")
	obj.Value("data").Object().Value("business_hours").Equal("09:00-10:00")
}
