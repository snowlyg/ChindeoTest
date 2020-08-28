package main

import (
	"github.com/snowlyg/ChindeoTest/common"
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestMiniWechatOnlineMyPatientListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/online/v1/patient/my_patient").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(1)
	object := obj.Value("data").Object().Value("data").Array().First().Object()
	object.Value("id").Equal(Patient.ID)
	object.Value("username").Equal(Patient.Username)
	object.Value("nickname").Equal(Patient.Nickname)
	object.Value("phone").Equal(Patient.Phone)
	object.Value("sex").Object().Value("text").Equal("女")
	object.Value("sex").Object().Value("value").Equal(Patient.Sex)
	object.Value("status").Equal(1)
	object.Value("avatar_url").Equal(Patient.AvatarURL)
	object.Value("open_id").Equal(Patient.OpenID)
	object.Value("union_id").Equal(Patient.UnionID)
	object.Value("country").Equal(Patient.Country)
	object.Value("province").Equal(Patient.Province)
	object.Value("mac").Equal(Patient.Mac)
}

func TestMiniWechatOnlineMyPatientNoPageSizeListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/online/v1/patient/my_patient").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithQuery("page_size", -1).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(1)
	object := obj.Value("data").Array().First().Object()
	object.Value("id").Equal(Patient.ID)
	object.Value("username").Equal(Patient.Username)
	object.Value("nickname").Equal(Patient.Nickname)
	object.Value("phone").Equal(Patient.Phone)
	object.Value("sex").Object().Value("text").Equal("女")
	object.Value("sex").Object().Value("value").Equal(Patient.Sex)
	object.Value("status").Equal(1)
	object.Value("avatar_url").Equal(Patient.AvatarURL)
	object.Value("open_id").Equal(Patient.OpenID)
	object.Value("union_id").Equal(Patient.UnionID)
	object.Value("country").Equal(Patient.Country)
	object.Value("province").Equal(Patient.Province)
	object.Value("mac").Equal(Patient.Mac)
}

func TestMiniWechatOnlineMyPatientShowSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/online/v1/patient/{id}", Patient.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithQuery("page_size", -1).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	object := obj.Value("data").Object()
	object.Value("id").Equal(Patient.ID)
	object.Value("username").Equal(Patient.Username)
	object.Value("nickname").Equal(Patient.Nickname)
	object.Value("phone").Equal(Patient.Phone)
	object.Value("sex").Object().Value("text").Equal("女")
	object.Value("sex").Object().Value("value").Equal(Patient.Sex)
	object.Value("status").Equal(1)
	object.Value("avatar_url").Equal(Patient.AvatarURL)
	object.Value("open_id").Equal(Patient.OpenID)
	object.Value("union_id").Equal(Patient.UnionID)
	object.Value("country").Equal(Patient.Country)
	object.Value("province").Equal(Patient.Province)
	object.Value("mac").Equal(Patient.Mac)
}
