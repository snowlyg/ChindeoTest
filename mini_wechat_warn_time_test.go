package main

import (
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var warmTimeId int

func TestMiniWechatWarmTimeAddSuccess(t *testing.T) {
	warmTime := map[string]interface{}{
		"weeks":  "1,2,3,4,5,6,7",
		"time":   "10:26",
		"status": 1,
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.POST("/api/v1/outline/warn_time/add").
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(warmTime).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("添加成功")
	obj.Value("data").Object().Value("id").NotNull()
	obj.Value("data").Object().Value("status").Equal("1")
	obj.Value("data").Object().Value("time").Equal("10:26")
	obj.Value("data").Object().Value("weeks").Equal("1,2,3,4,5,6,7")

	id := obj.Value("data").Object().Value("id").Raw()
	data, ok := id.(string)
	if ok {
		warmTimeId, _ = strconv.Atoi(data)
	}
}

func TestMiniWechatWarmTimeUpdateSuccess(t *testing.T) {
	warmTime := map[string]interface{}{
		"weeks":  "1,2,3,4,5",
		"time":   "09:26",
		"status": 0,
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.POST("/api/v1/outline/warn_time/{id}", warmTimeId).
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(warmTime).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("更新成功")
	obj.Value("data").Object().Value("id").NotNull()
	obj.Value("data").Object().Value("status").Equal("0")
	obj.Value("data").Object().Value("time").Equal("09:26")
	obj.Value("data").Object().Value("weeks").Equal("1,2,3,4,5")
}
func TestMiniWechatWarmTimeStatusSuccess(t *testing.T) {
	warmTime := map[string]interface{}{
		"weeks":  "1,2,3,4,5",
		"time":   "09:26",
		"status": 0,
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/api/v1/outline/warn_time/status/{id}", warmTimeId).
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(warmTime).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("修改成功")
	obj.Value("data").Object().Value("id").NotNull()
	obj.Value("data").Object().Value("status").Boolean().True()
	obj.Value("data").Object().Value("time").Equal("09:26")
	obj.Value("data").Object().Value("weeks").Equal("1,2,3,4,5")
}

func TestMiniWechatWarmTimeSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/api/v1/outline/warn_time").
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().NotEmpty()
}

func TestMiniWechatWarmTimeDeleteSuccess(t *testing.T) {
	warmTime := map[string]interface{}{
		"weeks":  "1,2,3,4,5",
		"time":   "09:26",
		"status": 0,
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.DELETE("/api/v1/outline/warn_time/{id}", warmTimeId).
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(warmTime).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("删除成功")
}
