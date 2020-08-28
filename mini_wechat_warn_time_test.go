package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"strconv"
	"testing"
)

var warmTimeId int

func TestMiniWechatWarmTimeAddSuccess(t *testing.T) {
	warmTime := map[string]interface{}{
		"weeks":  "1,2,3,4,5,6,7",
		"time":   "10:26",
		"status": 1,
	}

	obj := model.GetE(t).POST("/api/v1/outline/warn_time/add").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
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

	obj := model.GetE(t).POST("/api/v1/outline/warn_time/{id}", warmTimeId).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
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

	obj := model.GetE(t).GET("/api/v1/outline/warn_time/status/{id}", warmTimeId).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
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
	obj := model.GetE(t).GET("/api/v1/outline/warn_time").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(1)
}

func TestMiniWechatWarmTimeDeleteSuccess(t *testing.T) {
	warmTime := map[string]interface{}{
		"weeks":  "1,2,3,4,5",
		"time":   "09:26",
		"status": 0,
	}

	obj := model.GetE(t).DELETE("/api/v1/outline/warn_time/{id}", warmTimeId).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(warmTime).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("删除成功")
}

func TestMiniWechatWarmTimeAfterDelSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/warn_time").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("查询成功")
	obj.Value("data").Array().Length().Equal(0)
}
