package main

import (
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var miniWechatOrderId int

func TestMiniWechatOrderListSuccess(t *testing.T) {
	re := map[string]interface{}{
		"status":      1,
		"page_size":   10,
		"hospital_no": "9556854545",
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.POST("/api/v1/outline/o_order").
		WithQuery("page", 2).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4", "IsDev": "1"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(re).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	for _, private := range obj.Value("data").Object().Value("data").Array().Iter() {
		private.Object().Value("status").Equal("待付款")
	}
}

func TestMiniWechatOrderAddSuccess(t *testing.T) {
	order := map[string]interface{}{
		"amount":         12,
		"total":          32.00,
		"addr_id":        addrId,
		"application_id": applicationId,
		"rmk":            "455445455544",
		"menus": []map[string]interface{}{
			{
				"menu_name":      "红烧肉",
				"menu_type":      1,
				"menu_time_type": "中餐",
				"menu_desc":      "好吃就吃红烧肉",
				"price":          10,
				"menu_id":        2,
				"cover":          "http://",
				"amount":         4,
			},
			{
				"menu_name":      "红烧肉",
				"menu_type":      1,
				"menu_time_type": "中餐",
				"menu_desc":      "好吃就吃红烧肉",
				"price":          10,
				"menu_id":        2,
				"cover":          "http://",
				"amount":         4,
			},
		},
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.POST("/api/v1/outline/o_order/add").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4", "IsDev": "1"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(order).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").NotEqual(0)
	obj.Value("data").Object().Value("order_no").String().Contains("O")
	obj.Value("data").Object().Value("amount").String().Equal("12")
	obj.Value("data").Object().Value("total").String().Equal("32")
	obj.Value("data").Object().Value("rmk").String().Equal("455445455544")
	obj.Value("data").Object().Value("app_type").Number().Equal(1)
	obj.Value("data").Object().Value("pay_type").Number().Equal(0)
	obj.Value("data").Object().Value("application_id").Equal(strconv.FormatFloat(applicationId, 'g', -1, 64))

	id := obj.Value("data").Object().Value("id").Raw()
	data, ok := id.(string)
	if ok {
		miniWechatOrderId, _ = strconv.Atoi(data)
	}
}

func TestMiniWechatOrderAddNoApplicationId(t *testing.T) {
	order := map[string]interface{}{
		"amount":  12,
		"total":   32.00,
		"addr_id": addrId,
		"rmk":     "455445455544",
		"menus": []map[string]interface{}{
			{
				"menu_name":      "红烧肉",
				"menu_type":      1,
				"menu_time_type": "中餐",
				"menu_desc":      "好吃就吃红烧肉",
				"price":          10,
				"menu_id":        2,
				"cover":          "http://",
				"amount":         4,
			},
			{
				"menu_name":      "红烧肉",
				"menu_type":      1,
				"menu_time_type": "中餐",
				"menu_desc":      "好吃就吃红烧肉",
				"price":          10,
				"menu_id":        2,
				"cover":          "http://",
				"amount":         4,
			},
		},
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.POST("/api/v1/outline/o_order/add").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4", "IsDev": "1"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(order).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("医院不能为空！")
}

func TestMiniWechatOrderAddNoAddrId(t *testing.T) {
	order := map[string]interface{}{
		"amount":         12,
		"total":          32.00,
		"application_id": applicationId,
		"rmk":            "455445455544",
		"menus": []map[string]interface{}{
			{
				"menu_name":      "红烧肉",
				"menu_type":      1,
				"menu_time_type": "中餐",
				"menu_desc":      "好吃就吃红烧肉",
				"price":          10,
				"menu_id":        2,
				"cover":          "http://",
				"amount":         4,
			},
			{
				"menu_name":      "红烧肉",
				"menu_type":      1,
				"menu_time_type": "中餐",
				"menu_desc":      "好吃就吃红烧肉",
				"price":          10,
				"menu_id":        2,
				"cover":          "http://",
				"amount":         4,
			},
		},
	}
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.POST("/api/v1/outline/o_order/add").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4", "IsDev": "1"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(order).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("订单地址错误")
}

func TestMiniWechatOrderShowSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/api/v1/outline/o_order/{id}", miniWechatOrderId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4", "IsDev": "1"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(miniWechatOrderId)
	obj.Value("data").Object().Value("order_no").String().Contains("O")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(1)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("待付款")
	obj.Value("data").Object().Value("amount").Number().Equal(12)
	obj.Value("data").Object().Value("total").String().Equal("32.00")
	obj.Value("data").Object().Value("is_return").String().Equal("有退款")
	obj.Value("data").Object().Value("menus").Array().Length().Equal(2)
	obj.Value("data").Object().Value("return_order").Null()
	obj.Value("data").Object().Value("addr").Object().Value("id").NotNull()
	obj.Value("data").Object().Value("addr").Object().Value("o_order_id").Equal(miniWechatOrderId)

}

func TestMiniWechatOrderPaySuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/api/v1/outline/o_order/pay/{id}", miniWechatOrderId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4", "IsDev": "1"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(401)
	obj.Value("message").String().Equal("小程序配置错误")
}

func TestMiniWechatOrderCancelSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/api/v1/outline/o_order/cancel/{id}", miniWechatOrderId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4", "IsDev": "1"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}
