package main

import (
	"github.com/snowlyg/ChindeoTest/common"
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestMiniWechatRefreshTokenSuccess(t *testing.T) {
	auth := map[string]interface{}{
		"uuid": "5205857593c2eacc6f6c1da376b32ca3",
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.POST("/api/v1/refresh_access_token").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithJSON(auth).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("刷新成功")
	obj.Value("data").Object().Value("AccessToken").NotNull()

	token := obj.Value("data").Object().Value("AccessToken").Raw()
	data, ok := token.(string)
	if ok {
		MiniWechatToken = data
	}
}

func TestMiniWechatRefreshTokenError(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.POST("/api/v1/refresh_access_token").
		WithHeaders(map[string]string{"X-Token": "", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(401)
	obj.Value("message").String().Equal("错误 token 结构")
	obj.Value("data").Null()
}
