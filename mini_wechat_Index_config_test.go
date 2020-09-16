package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatIndexConfigRelationshipsSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/index/config/relationship").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("获取家庭关系成功")
}

func TestMiniWechatIndexConfigVitalTimesSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/index/config/vitalTimes").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("监测时间信息获取成功")
}

func TestMiniWechatIndexConfigVitalConfigSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/index/config/vitalConfigs").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("获取检测配置成功")
}
