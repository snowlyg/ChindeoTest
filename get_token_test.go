package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"strconv"
	"testing"
)

func TestGetTokenSuccess(t *testing.T) {
	auth := map[string]interface{}{
		"app_id":     "b44fc017043763eb5ac15f0069d77c",
		"app_secret": "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b",
	}

	obj := model.GetE(t).POST("/api/v1/get_access_token").
		WithHeaders(map[string]string{"AuthType": strconv.FormatInt(int64(model.AUTH_TYPE_SERVER), 10)}).
		WithJSON(auth).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("授权成功")
	obj.Value("data").Object().Value("AccessToken").NotNull()

	token := obj.Value("data").Object().Value("AccessToken").Raw()
	data, ok := token.(string)
	if ok {
		model.Token = data
	}
}

func TestGetTokenErrorAuthType(t *testing.T) {
	auth := map[string]interface{}{
		"app_id":     "b44fc017043763eb5ac15f0069d77c",
		"app_secret": "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b",
	}

	obj := model.GetE(t).POST("/api/v1/get_access_token").
		WithHeaders(map[string]string{"AuthType": "10"}).
		WithJSON(auth).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("auth_type 错误!")
	obj.Value("data").Null()
}

func TestGetTokenEmptyAppId(t *testing.T) {
	auth := map[string]interface{}{
		"app_id":     "",
		"app_secret": "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b",
	}

	obj := model.GetE(t).POST("/api/v1/get_access_token").
		WithHeaders(map[string]string{"AuthType": strconv.FormatInt(int64(model.AUTH_TYPE_SERVER), 10)}).
		WithJSON(auth).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("app_id 不能为空！")
	obj.Value("data").Null()
}

func TestGetTokenEmptyAppSecret(t *testing.T) {
	auth := map[string]interface{}{
		"app_id":     "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b",
		"app_secret": "",
	}

	obj := model.GetE(t).POST("/api/v1/get_access_token").
		WithHeaders(map[string]string{"AuthType": strconv.FormatInt(int64(model.AUTH_TYPE_SERVER), 10)}).
		WithJSON(auth).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("app_secret 不能为空！")
	obj.Value("data").Null()
}

func TestGetTokenErrorAppSecretOrAppId(t *testing.T) {
	auth := map[string]interface{}{
		"app_id":     "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b",
		"app_secret": "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b",
	}

	obj := model.GetE(t).POST("/api/v1/get_access_token").
		WithHeaders(map[string]string{"AuthType": strconv.FormatInt(int64(model.AUTH_TYPE_SERVER), 10)}).
		WithJSON(auth).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(401)
	obj.Value("message").String().Equal("应用不存在!")
	obj.Value("data").Null()
}
