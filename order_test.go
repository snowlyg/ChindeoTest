package main

import (
	"github.com/snowlyg/ChindeoTest/common"
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var orderId int

func TestOrderListSuccess(t *testing.T) {
	re := map[string]interface{}{
		"status":      0,
		"page_size":   10,
		"hospital_no": "9556854545",
		"id_card_no":  IdCardNo,
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.POST("/api/v1/i_order").
		WithQuery("page", 1).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(re).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("data").Array().Length().Equal(OrderCount)

	last := obj.Value("data").Object().Value("data").Array().Last().Object()
	last.Value("status").Object().Value("text").Equal("已付款")
	last.Value("status").Object().Value("value").Equal(common.I_ORDER_STATUS_FOR_DELIVERY)
	last.Value("is_return").Equal("未退款")

}

func TestOrderAddSuccess(t *testing.T) {
	order := map[string]interface{}{
		"amount":       12,
		"total":        32.00,
		"patient_name": "操蛋",
		"loc_name":     "泥马",
		"bed_num":      05,
		"hospital_no":  "9556854545",
		"disease":      "disease",
		"addr":         "addr",
		"phone":        "13800138000",
		"id_card_no":   IdCardNo,
		"rmk":          "455445455544",
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
	obj := e.POST("/api/v1/i_order/add").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(order).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").NotEqual(0)
	obj.Value("data").Object().Value("order_no").String().Contains("I")
	obj.Value("data").Object().Value("amount").Equal(12)
	obj.Value("data").Object().Value("total").Equal(32)
	obj.Value("data").Object().Value("rmk").String().Equal("455445455544")
	obj.Value("data").Object().Value("app_type").Number().Equal(common.ORDER_APP_TYPE_BED)
	obj.Value("data").Object().Value("pay_type").Number().Equal(0)
	obj.Value("data").Object().Value("application_id").Number().Equal(13)

	id := obj.Value("data").Object().Value("id").Raw()
	data, ok := id.(string)
	if ok {
		orderId, _ = strconv.Atoi(data)
	}
	OrderCount++
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
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.POST("/api/v1/i_order/add").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(order).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("id_card_no不符合指定规则")
}

func TestOrderAfterOrderAddMenuShowSuccess(t *testing.T) {
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

func TestOrderShowSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/api/v1/i_order/{id}", orderId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(orderId)
	obj.Value("data").Object().Value("order_no").String().Contains("I")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(common.I_ORDER_STATUS_FOR_PAY)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("待付款")
	obj.Value("data").Object().Value("amount").Number().Equal(12)
	obj.Value("data").Object().Value("total").String().Equal("32.00")
	obj.Value("data").Object().Value("is_return").String().Equal("未退款")
	obj.Value("data").Object().Value("menus").Array().Length().Equal(1)

	obj.Value("data").Object().Value("return_order").Null()
	obj.Value("data").Object().Value("comments").Array().Length().Equal(0)

	obj.Value("data").Object().Value("addr").Object().Value("id").NotNull()
	obj.Value("data").Object().Value("addr").Object().Value("name").Equal("操蛋")
	obj.Value("data").Object().Value("addr").Object().Value("sex").Equal(1)
	obj.Value("data").Object().Value("addr").Object().Value("o_order_id").Equal(orderId)
	obj.Value("data").Object().Value("addr").Object().Value("loc_name").Equal("泥马")
	obj.Value("data").Object().Value("addr").Object().Value("hospital_no").Equal("9556854545")
	obj.Value("data").Object().Value("addr").Object().Value("hospital_name").Equal("我的医院")
	obj.Value("data").Object().Value("addr").Object().Value("age").Equal(1)
	obj.Value("data").Object().Value("addr").Object().Value("disease").Equal("disease")
	obj.Value("data").Object().Value("addr").Object().Value("phone").Equal("13800138000")
	obj.Value("data").Object().Value("addr").Object().Value("addr").Equal("addr")
}

func TestOrderPaySuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/api/v1/i_order/pay/{id}", orderId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("qrcode").NotNull()
}

func TestOrderCancelNoPaySuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/api/v1/i_order/cancel/{id}", orderId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestOrderCancelPaySuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/api/v1/i_order/cancel/{id}", Order.ID).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestOrderShowReturnSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/api/v1/i_order/{id}", Order.ID).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	orderMenuPrice := strconv.FormatFloat(OrderMenus[0].Price, 'g', -1, 64)

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(Order.ID)
	obj.Value("data").Object().Value("order_no").String().Contains("O")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(common.I_ORDER_STATUS_FOR_CANCEL)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")
	obj.Value("data").Object().Value("amount").Number().Equal(Order.Amount)
	obj.Value("data").Object().Value("total").Equal("10.00")
	obj.Value("data").Object().Value("is_return").String().Equal("有退款")

	obj.Value("data").Object().Value("comments").Array().Length().Equal(1)
	comment := obj.Value("data").Object().Value("comments").Array().First().Object()
	comment.Value("id").NotNull()
	comment.Value("user_id").Equal(User.ID)
	comment.Value("application_id").Equal(AppId)
	comment.Value("content").Equal(OrderComment.Content)
	comment.Value("star").Equal(5)
	comment.Value("pics").Array().Length().Equal(2)
	comment.Value("pics").Array().First().Equal("https")
	comment.Value("pics").Array().Last().Equal("https")
	obj.Value("data").Object().Value("menus").Array().Length().Equal(len(OrderMenus))

	orderMenu := obj.Value("data").Object().Value("menus").Array().First().Object()
	orderMenu.Value("id").NotNull()
	orderMenu.Value("menu_name").Equal(OrderMenus[0].MenuName)
	orderMenu.Value("menu_type").Equal(OrderMenus[0].MenuType)
	orderMenu.Value("menu_time_type").Equal(OrderMenus[0].MenuTimeType)
	orderMenu.Value("menu_desc").Equal(OrderMenus[0].MenuDesc)
	orderMenu.Value("menu_id").Equal(OrderMenus[0].MenuID)
	orderMenu.Value("price").Equal(orderMenuPrice)
	orderMenu.Value("amount").Equal(OrderMenus[0].Amount)
	orderMenu.Value("cover").Equal(OrderMenus[0].Cover)

	orderAddr := obj.Value("data").Object().Value("addr").Object()
	orderAddr.Value("id").Equal(OrderAddr.ID)
	orderAddr.Value("name").Equal(OrderAddr.Name)
	orderAddr.Value("sex").Equal(OrderAddr.Sex)
	orderAddr.Value("o_order_id").Equal(OrderAddr.OOrderID)
	orderAddr.Value("loc_name").Equal(OrderAddr.LocName)
	orderAddr.Value("hospital_no").Equal(OrderAddr.HospitalNo)
	orderAddr.Value("hospital_name").Equal(OrderAddr.HospitalName)
	orderAddr.Value("age").Equal(OrderAddr.Age)
	orderAddr.Value("disease").Equal(OrderAddr.Disease)
	orderAddr.Value("phone").Equal(OrderAddr.Phone)
	orderAddr.Value("addr").Equal(OrderAddr.Addr)

	orderReturn := obj.Value("data").Object().Value("return_order").Object()
	orderReturn.Value("id").NotNull()
	orderReturn.Value("order_no").String().Contains("RI")
	orderReturn.Value("status").Object().Value("value").Equal(common.I_RETURN_ORDER_STATUS_FOR_AUDIT)
	orderReturn.Value("status").Object().Value("text").Equal("待审核")
	orderReturn.Value("amount").Equal(Order.Amount)
	orderReturn.Value("total").Equal("10")
	orderReturn.Value("open_id").Equal("")
	orderReturn.Value("app_type").Equal(Order.AppType)
	orderReturn.Value("trade_type").Equal("")
	orderReturn.Value("o_order_id").Equal(Order.ID)
	orderReturn.Value("application_id").Equal(Order.ApplicationID)
	orderReturn.Value("pay_type").Equal(Order.PayType)
	orderReturn.Value("user_id").Equal(Order.UserID)

	returnAddr := orderReturn.Value("addr").Object()
	returnAddr.Value("id").NotNull()
	returnAddr.Value("name").Equal(OrderAddr.Name)
	returnAddr.Value("sex").Equal(OrderAddr.Sex)
	returnAddr.Value("o_return_order_id").NotNull()
	returnAddr.Value("loc_name").Equal(OrderAddr.LocName)
	returnAddr.Value("hospital_no").Equal(OrderAddr.HospitalNo)
	returnAddr.Value("hospital_name").Equal(OrderAddr.HospitalName)
	returnAddr.Value("age").Equal(OrderAddr.Age)
	returnAddr.Value("disease").Equal(OrderAddr.Disease)
	returnAddr.Value("phone").Equal(OrderAddr.Phone)
	returnAddr.Value("addr").Equal(OrderAddr.Addr)

	orderReturn.Value("menus").Array().Length().Equal(len(OrderMenus))
	returnOrderMenu := orderReturn.Value("menus").Array().First().Object()
	returnOrderMenu.Value("id").NotNull()
	returnOrderMenu.Value("menu_name").Equal(OrderMenus[0].MenuName)
	returnOrderMenu.Value("menu_type").Equal(OrderMenus[0].MenuType)
	returnOrderMenu.Value("menu_time_type").Equal(OrderMenus[0].MenuTimeType)
	returnOrderMenu.Value("menu_desc").Equal(OrderMenus[0].MenuDesc)
	returnOrderMenu.Value("menu_id").Equal(OrderMenus[0].MenuID)
	returnOrderMenu.Value("price").Equal(orderMenuPrice)
	returnOrderMenu.Value("amount").Equal(OrderMenus[0].Amount)
	returnOrderMenu.Value("cover").Equal(OrderMenus[0].Cover)
}

func TestOrderDeleteSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.DELETE("/api/v1/i_order/{id}", orderId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	OrderCount--
}

func TestOrderDeleteReturnSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.DELETE("/api/v1/i_order/{id}", Order.ID).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	OrderCount--
}

func TestOrderDeleteError(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.DELETE("/api/v1/i_order/{id}", 99999).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("数据不存在")
}

func TestOrderDeleteErrorForZore(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.DELETE("/api/v1/i_order/{id}", 0).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("错误数据")
}
