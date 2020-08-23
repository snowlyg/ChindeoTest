package main

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var delAddrId string
var addrId float64

func TestAddrListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.GET("/api/v1/outline/addr").
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	id := obj.Value("data").Array().First().Object().Value("id").Raw()
	data, ok := id.(float64)
	if ok {
		addrId = data
	}
}

func TestAddrAddSuccess(t *testing.T) {
	addr := map[string]interface{}{
		"name":          "name",
		"phone":         "13800138000",
		"addr":          "addr",
		"sex":           1,
		"age":           10,
		"is_default":    1,
		"patient_name":  "操蛋",
		"hospital_name": "蛋",
		"disease":       "你妹",
		"loc_name":      "泥马",
		"bed_num":       "05",
		"hospital_no":   "9556854545",
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.POST("/api/v1/outline/addr/add").
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(addr).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("操作成功")
	obj.Value("data").Object().Value("id").NotEqual(0)
	obj.Value("data").Object().Value("name").Equal("name")
	obj.Value("data").Object().Value("phone").Equal("13800138000")
	obj.Value("data").Object().Value("addr").Equal("addr")
	obj.Value("data").Object().Value("sex").Equal("男")
	obj.Value("data").Object().Value("age").Equal("10")
	obj.Value("data").Object().Value("is_default").Equal("1")
	obj.Value("data").Object().Value("hospital_name").Equal("蛋")
	obj.Value("data").Object().Value("disease").Equal("你妹")
	obj.Value("data").Object().Value("loc_name").Equal("泥马")
	obj.Value("data").Object().Value("bed_num").Equal("05")
	obj.Value("data").Object().Value("hospital_no").Equal("9556854545")

	id := obj.Value("data").Object().Value("id").Raw()
	data, ok := id.(string)
	if ok {
		delAddrId = data
	}
}

func TestAddrAddForOrderSuccess(t *testing.T) {
	addr := map[string]interface{}{
		"name":       "name",
		"phone":      "13800138000",
		"addr":       "addr",
		"sex":        1,
		"is_default": 1,
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.POST("/api/v1/outline/addr/add").
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(addr).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("操作成功")
	obj.Value("data").Object().Value("id").NotEqual(0)
	obj.Value("data").Object().Value("name").Equal("name")
	obj.Value("data").Object().Value("phone").Equal("13800138000")
	obj.Value("data").Object().Value("addr").Equal("addr")
	obj.Value("data").Object().Value("sex").Equal("男")
	obj.Value("data").Object().Value("is_default").Equal("1")
	obj.Value("data").Object().Value("hospital_name").Equal("")
	obj.Value("data").Object().Value("age").Equal("")
	obj.Value("data").Object().Value("disease").Equal("")
	obj.Value("data").Object().Value("loc_name").Equal("")
	obj.Value("data").Object().Value("bed_num").Equal("")
	obj.Value("data").Object().Value("hospital_no").Equal("")

	id := obj.Value("data").Object().Value("id").Raw()
	data, ok := id.(string)
	if ok {
		delAddrId = data
	}
}

func TestAddrAddError(t *testing.T) {
	addr := map[string]interface{}{
		"name":          "",
		"phone":         "13800138000",
		"addr":          "addr",
		"sex":           1,
		"age":           10,
		"is_default":    1,
		"patient_name":  "操蛋",
		"hospital_name": "蛋",
		"disease":       "你妹",
		"loc_name":      "泥马",
		"bed_num":       "05",
		"hospital_no":   "9556854545",
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.POST("/api/v1/outline/addr/add").
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(addr).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("联系人名称不能为空！")
}

func TestAddrUpdateSuccess(t *testing.T) {
	addr := map[string]interface{}{
		"name":          "update",
		"phone":         "13800138001",
		"addr":          "update",
		"sex":           0,
		"age":           12,
		"is_default":    0,
		"patient_name":  "update",
		"hospital_name": "update",
		"disease":       "update",
		"loc_name":      "update",
		"bed_num":       "update",
		"hospital_no":   "update",
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.POST("/api/v1/outline/addr/{id}", delAddrId).
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(addr).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("操作成功")
	obj.Value("data").Object().Value("id").NotEqual(0)
	obj.Value("data").Object().Value("name").Equal("update")
	obj.Value("data").Object().Value("phone").Equal("13800138001")
	obj.Value("data").Object().Value("addr").Equal("update")
	obj.Value("data").Object().Value("sex").Equal("女")
	obj.Value("data").Object().Value("age").Equal("12")
	obj.Value("data").Object().Value("is_default").Equal("0")
	obj.Value("data").Object().Value("hospital_name").Equal("update")
	obj.Value("data").Object().Value("disease").Equal("update")
	obj.Value("data").Object().Value("loc_name").Equal("update")
	obj.Value("data").Object().Value("bed_num").Equal("update")
	obj.Value("data").Object().Value("hospital_no").Equal("update")
}

func TestAddrUpdateNoName(t *testing.T) {
	addr := map[string]interface{}{
		"name":          "",
		"phone":         "13800138001",
		"addr":          "update",
		"sex":           0,
		"age":           12,
		"is_default":    0,
		"patient_name":  "update",
		"hospital_name": "update",
		"disease":       "update",
		"loc_name":      "update",
		"bed_num":       "update",
		"hospital_no":   "update",
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.POST("/api/v1/outline/addr/{id}", delAddrId).
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(addr).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("联系人名称不能为空！")
}

func TestAddrUpdateErrorSex(t *testing.T) {
	addr := map[string]interface{}{
		"name":          "13800138001",
		"phone":         "13800138001",
		"addr":          "update",
		"sex":           3,
		"age":           12,
		"is_default":    0,
		"patient_name":  "update",
		"hospital_name": "update",
		"disease":       "update",
		"loc_name":      "update",
		"bed_num":       "update",
		"hospital_no":   "update",
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.POST("/api/v1/outline/addr/{id}", delAddrId).
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(addr).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("性别必须为0或1！")
}

func TestAddrDeleteSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.DELETE("/api/v1/outline/addr/{id}", delAddrId).
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("操作成功")
}

func TestAddrDeleteError(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.DELETE("/api/v1/outline/addr/{id}", 0).
		WithHeaders(map[string]string{"X-Token": Token, "IsDev": "1", "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("错误数据")
}
