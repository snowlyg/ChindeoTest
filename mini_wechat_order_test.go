package main

import (
	"github.com/snowlyg/ChindeoTest/common"
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var miniWechatOrderId int

func TestMiniWechatOrderListSuccess(t *testing.T) {
	re := map[string]interface{}{
		"status":         2,
		"page_size":      10,
		"application_id": AppId,
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.POST("/api/v1/outline/o_order").
		WithQuery("page", 1).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithJSON(re).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("data").Array().Length().Equal(1)

	order := obj.Value("data").Object().Value("data").Array().First().Object()
	order.Value("id").Equal(Order.ID)
	order.Value("order_no").String().Contains("O")
	order.Value("status").Object().Value("value").Equal(2)
	order.Value("status").Object().Value("text").Equal("已付款")
	order.Value("amount").Number().Equal(Order.Amount)
	order.Value("total").String().Equal("10.00")
	order.Value("is_return").String().Equal("未退款")
	order.Value("menus").Array().Length().Equal(1)
	order.Value("return_order").Null()

	addr := order.Value("addr").Object()
	addr.Value("id").Equal(OrderAddr.ID)
	addr.Value("name").Equal(OrderAddr.Name)
	addr.Value("sex").Equal(OrderAddr.Sex)
	addr.Value("o_order_id").Equal(OrderAddr.OOrderID)
	addr.Value("loc_name").Equal(OrderAddr.LocName)
	addr.Value("hospital_no").Equal(OrderAddr.HospitalNo)
	addr.Value("hospital_name").Equal(OrderAddr.HospitalName)
	addr.Value("age").Equal(OrderAddr.Age)
	addr.Value("disease").Equal(OrderAddr.Disease)
	addr.Value("phone").Equal(OrderAddr.Phone)
	addr.Value("addr").Equal(OrderAddr.Addr)
}

func TestMiniWechatOrderAddSuccess(t *testing.T) {
	order := map[string]interface{}{
		"amount":         12,
		"total":          32.00,
		"addr_id":        Addr.ID,
		"application_id": AppId,
		"rmk":            "455445455544",
		"menus": []map[string]interface{}{
			{
				"menu_name":      Menu.Name,
				"menu_type":      Menu.MenuTypeID,
				"menu_time_type": Menu.TimeType,
				"menu_desc":      Menu.Desc,
				"price":          Menu.Price,
				"menu_id":        Menu.ID,
				"cover":          Menu.Cover,
				"amount":         Menu.Amount,
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
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithJSON(order).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").NotEqual(0)
	obj.Value("data").Object().Value("order_no").String().Contains("O")
	obj.Value("data").Object().Value("amount").Equal(12)
	obj.Value("data").Object().Value("total").Equal(32)
	obj.Value("data").Object().Value("rmk").String().Equal("455445455544")
	obj.Value("data").Object().Value("app_type").Number().Equal(common.ORDER_APP_TYPE_MINI)
	obj.Value("data").Object().Value("pay_type").Number().Equal(0)
	obj.Value("data").Object().Value("application_id").Equal(AppId)

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
		"addr_id": OrderAddr.ID,
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
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
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
		"application_id": AppId,
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
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithJSON(order).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("订单地址错误")

}

func TestMiniWechatOrderAfterOrderAddMenuShowSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/api/v1/outline/menu/{id}", Menu.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithQuery("application_id", AppId).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(Menu.ID)
	obj.Value("data").Object().Value("name").Equal(Menu.Name)
	obj.Value("data").Object().Value("time_type").Equal(Menu.TimeType.String())
	obj.Value("data").Object().Value("menu_type_id").Equal(Menu.MenuTypeID)
	obj.Value("data").Object().Value("desc").Equal(Menu.Desc)
	obj.Value("data").Object().Value("amount").Equal(MenuAmount + 1)
	MenuAmount++
	obj.Value("data").Object().Value("price").Equal("10.00")
	obj.Value("data").Object().Value("cover").Equal(Menu.Cover)
	obj.Value("data").Object().Value("pics").Array().Length().Equal(0)
	obj.Value("data").Object().Value("create_at").String().Contains(Menu.CreateAt.Format("2006-01-02 15:04"))
	obj.Value("data").Object().Value("create_at").String().Contains(Menu.UpdateAt.Format("2006-01-02 15:04"))

	menuType := obj.Value("data").Object().Value("type").Object()
	menuType.Value("id").Equal(MenuType.ID)
	menuType.Value("name").Equal(MenuType.Name)

	obj.Value("data").Object().Value("tags").Array().Length().Equal(1)
	menuTag := obj.Value("data").Object().Value("tags").Array().First().Object()
	menuTag.Value("id").Equal(MenuTag.ID)
	menuTag.Value("name").Equal(MenuTag.Name)
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
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
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
	obj.Value("data").Object().Value("is_return").String().Equal("未退款")
	obj.Value("data").Object().Value("menus").Array().Length().Equal(1)
	obj.Value("data").Object().Value("return_order").Null()
	obj.Value("data").Object().Value("addr").Object().Value("id").NotNull()
	obj.Value("data").Object().Value("addr").Object().Value("name").Equal(Addr.Name)
	obj.Value("data").Object().Value("addr").Object().Value("sex").Equal(Addr.Sex)
	obj.Value("data").Object().Value("addr").Object().Value("o_order_id").Equal(miniWechatOrderId)
	obj.Value("data").Object().Value("addr").Object().Value("loc_name").Equal(Addr.LocName)
	obj.Value("data").Object().Value("addr").Object().Value("hospital_no").Equal(Addr.HospitalNo)
	obj.Value("data").Object().Value("addr").Object().Value("hospital_name").Equal(Addr.HospitalName)
	obj.Value("data").Object().Value("addr").Object().Value("age").Equal(Addr.Age)
	obj.Value("data").Object().Value("addr").Object().Value("disease").Equal(Addr.Disease)
	obj.Value("data").Object().Value("addr").Object().Value("phone").Equal(Addr.Phone)
	obj.Value("data").Object().Value("addr").Object().Value("addr").Equal(Addr.Addr)

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
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("PARAM_ERROR:appid和openid不匹配")
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
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestMiniWechatOrderDeleteSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.DELETE("/api/v1/outline/o_order/{id}", miniWechatOrderId).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestMiniWechatOrderDeleteErrorForZore(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.DELETE("/api/v1/outline/o_order/{id}", 0).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("错误数据")
}

func TestMiniWechatOrderDeleteError(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.DELETE("/api/v1/outline/o_order/{id}", 99999).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("数据不存在")
}
