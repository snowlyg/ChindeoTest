package main

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var orderId int

func TestOrderListSuccess(t *testing.T) {
	re := map[string]interface{}{
		"status":      1,
		"page_size":   10,
		"hospital_no": "9556854545",
	}

	e := httpexpect.New(t, BaseUrl)
	obj := e.POST("/api/v1/i_order").
		WithQuery("page", 2).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
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

func TestOrderAddSuccess(t *testing.T) {
	order := map[string]interface{}{
		"amount":       12,
		"total":        32.00,
		"patient_name": "操蛋",
		"loc_name":     "泥马",
		"bed_num":      05,
		"hospital_no":  "hospital_no",
		"disease":      "disease",
		"addr":         "addr",
		"phone":        "13800138000",
		"id_card_no":   "430923155811586698",
		"rmk":          "455445455544",
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
	e := httpexpect.New(t, BaseUrl)
	obj := e.POST("/api/v1/i_order/add").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(order).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").NotEqual(0)
	obj.Value("data").Object().Value("order_no").String().Contains("I")
	obj.Value("data").Object().Value("amount").String().Equal("12")
	obj.Value("data").Object().Value("total").String().Equal("32")
	obj.Value("data").Object().Value("rmk").String().Equal("455445455544")
	obj.Value("data").Object().Value("app_type").Number().Equal(2)
	obj.Value("data").Object().Value("pay_type").Number().Equal(0)
	obj.Value("data").Object().Value("application_id").Number().Equal(13)

	id := obj.Value("data").Object().Value("id").Raw()
	data, ok := id.(string)
	if ok {
		orderId, _ = strconv.Atoi(data)
	}
}

func TestOrderAddErrorIdCardNo(t *testing.T) {
	order := map[string]interface{}{
		"amount":       12,
		"total":        32.00,
		"patient_name": "操蛋",
		"loc_name":     "泥马",
		"bed_num":      05,
		"hospital_no":  "hospital_no",
		"disease":      "disease",
		"addr":         "addr",
		"phone":        "13800138000",
		"id_card_no":   "455445455544",
		"rmk":          "455445455544",
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
	e := httpexpect.New(t, BaseUrl)
	obj := e.POST("/api/v1/i_order/add").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(order).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("id_card_no不符合指定规则")
}

func TestOrderShowSuccess(t *testing.T) {
	e := httpexpect.New(t, BaseUrl)
	obj := e.GET("/api/v1/i_order/{id}", orderId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(orderId)
	obj.Value("data").Object().Value("order_no").String().Contains("I")
	obj.Value("data").Object().Value("status").String().Equal("待付款")
	obj.Value("data").Object().Value("amount").Number().Equal(12)
	obj.Value("data").Object().Value("total").String().Equal("32.00")
	obj.Value("data").Object().Value("is_return").String().Equal("有退款")
	obj.Value("data").Object().Value("menus").Array().Length().Equal(2)
	obj.Value("data").Object().Value("return_order").Null()
	obj.Value("data").Object().Value("addr").Object().Value("id").NotNull()
	obj.Value("data").Object().Value("addr").Object().Value("name").Equal("操蛋")
	obj.Value("data").Object().Value("addr").Object().Value("sex").Equal(1)
	obj.Value("data").Object().Value("addr").Object().Value("o_order_id").Equal(orderId)
	obj.Value("data").Object().Value("addr").Object().Value("loc_name").Equal("泥马")
	obj.Value("data").Object().Value("addr").Object().Value("hospital_no").Equal("hospital_no")
	obj.Value("data").Object().Value("addr").Object().Value("hospital_name").Equal("我的医院")
	obj.Value("data").Object().Value("addr").Object().Value("age").Equal(1)
	obj.Value("data").Object().Value("addr").Object().Value("disease").Equal("disease")
	obj.Value("data").Object().Value("addr").Object().Value("phone").Equal("13800138000")
	obj.Value("data").Object().Value("addr").Object().Value("addr").Equal("addr")
}

func TestOrderPaySuccess(t *testing.T) {
	e := httpexpect.New(t, BaseUrl)
	obj := e.GET("/api/v1/i_order/pay/{id}", orderId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("qrcode").NotNull()
}

func TestOrderCancelSuccess(t *testing.T) {
	e := httpexpect.New(t, BaseUrl)
	obj := e.GET("/api/v1/i_order/cancel/{id}", orderId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}
