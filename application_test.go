package main

import (
	"github.com/snowlyg/ChindeoTest/common"
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestApplicationSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/api/v1/application").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
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
