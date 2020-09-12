package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"strconv"
	"testing"
)

var orderId int

func TestOrderListSuccess(t *testing.T) {
	re := map[string]interface{}{
		"status":      0,
		"page_size":   10,
		"hospital_no": "9556854545",
		"id_card_no":  model.IdCardNo,
	}

	obj := model.GetE(t).POST("/api/v1/i_order").
		WithQuery("page", 1).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(re).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.OrderCount)

	last := obj.Value("data").Object().Value("data").Array().Last().Object()
	last.Value("status").Object().Value("text").Equal("已付款")
	last.Value("status").Object().Value("value").Equal(model.IOrderStatusForDelivery)
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
		"id_card_no":   model.IdCardNo,
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

	obj := model.GetE(t).POST("/api/v1/i_order/add").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
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
	obj.Value("data").Object().Value("app_type").Number().Equal(model.OrderAppTypeBed)
	obj.Value("data").Object().Value("pay_type").Number().Equal(0)
	obj.Value("data").Object().Value("application_id").Number().Equal(13)

	id := obj.Value("data").Object().Value("id").Raw()
	data, ok := id.(string)
	if ok {
		orderId, _ = strconv.Atoi(data)
	}
	model.OrderCount++
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

	obj := model.GetE(t).POST("/api/v1/i_order/add").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(order).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("id_card_no不符合指定规则")
}

func TestOrderCommentSuccess(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "content",
		"id_card_no": model.IdCardNo,
		"pics":       model.GetPics(),
		"order_id":   orderId,
	}

	obj := model.GetE(t).POST("/common/v1/inner/comment/order").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("评论成功")
}

func TestOrderCommentNoContentError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "",
		"id_card_no": model.IdCardNo,
		"pics":       model.GetPics(),
		"order_id":   orderId,
	}

	obj := model.GetE(t).POST("/common/v1/inner/comment/order").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("评论内容不能为空！")
}

func TestOrderCommentOrderNotExistsError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "dsfsdfsd",
		"id_card_no": model.IdCardNo,
		"pics":       model.GetPics(),
		"order_id":   9999,
	}

	obj := model.GetE(t).POST("/common/v1/inner/comment/order").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("订单 9999 不存在")
}

func TestOrderCommentNoIdCardNoError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "456952158962254456",
		"id_card_no": "",
		"pics":       model.GetPics(),
		"order_id":   orderId,
	}

	obj := model.GetE(t).POST("/common/v1/inner/comment/order").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("评论成功")
}

func TestOrderCommentNoOrderIdError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "456952158962254456",
		"id_card_no": model.IdCardNo,
		"pics":       model.GetPics(),
	}

	obj := model.GetE(t).POST("/common/v1/inner/comment/order").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("评论订单不能为空！")
}

func TestOrderAfterOrderAddMenuShowSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/outline/menu/{id}", Menu.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("application_id", model.AppId).
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
	obj.Value("data").Object().Value("amount").Equal(model.MenuAmount + 1)
	model.MenuAmount++
	obj.Value("data").Object().Value("price").Equal("10.00")
	obj.Value("data").Object().Value("cover").Equal(Menu.Cover)
	obj.Value("data").Object().Value("pics").Array().Length().Equal(len(Menu.Pics))
	obj.Value("data").Object().Value("create_at").String().Contains(Menu.CreateAt.Format("2006-01-02 15:04"))
	obj.Value("data").Object().Value("create_at").String().Contains(Menu.UpdateAt.Format("2006-01-02 15:04"))

	menuType := obj.Value("data").Object().Value("type").Object()
	menuType.Value("id").Equal(Menu.MenuTypeID)
	menuType.Value("name").Equal(Menu.MenuType.Name)

	obj.Value("data").Object().Value("tags").Array().Length().Equal(len(Menu.MenuTags))
	if len(Menu.MenuTags) > 0 {
		menuTag := obj.Value("data").Object().Value("tags").Array().First().Object()
		menuTag.Value("id").Equal(Menu.MenuTags[0].ID)
		menuTag.Value("name").Equal(Menu.MenuTags[0].Name)
	}

}

func TestOrderShowSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/i_order/{id}", orderId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(orderId)
	obj.Value("data").Object().Value("order_no").String().Contains("I")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForPay)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("待付款")
	obj.Value("data").Object().Value("amount").Number().Equal(12)
	obj.Value("data").Object().Value("total").String().Equal("32.00")
	obj.Value("data").Object().Value("is_return").String().Equal("未退款")
	obj.Value("data").Object().Value("menus").Array().Length().Equal(1)

	obj.Value("data").Object().Value("return_order").Null()
	obj.Value("data").Object().Value("comments").Array().Length().Equal(2)

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
	obj := model.GetE(t).GET("/api/v1/i_order/pay/{id}", orderId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("qrcode").NotNull()
}

func TestOrderCancelNoPaySuccess(t *testing.T) {

	obj := model.GetE(t).GET("/api/v1/i_order/cancel/{id}", orderId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestOrderCancelPaySuccess(t *testing.T) {

	obj := model.GetE(t).GET("/api/v1/i_order/cancel/{id}", Order.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestOrderShowReturnSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/api/v1/i_order/{id}", Order.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	orderMenuPrice := strconv.FormatFloat(Order.OrderMenus[0].Price, 'g', -1, 64)

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(Order.ID)
	obj.Value("data").Object().Value("order_no").String().Equal("I202008241612348468756914")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForCancel)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")
	obj.Value("data").Object().Value("amount").Number().Equal(Order.Amount)
	obj.Value("data").Object().Value("total").Equal("10.00")
	obj.Value("data").Object().Value("is_return").String().Equal("有退款")

	obj.Value("data").Object().Value("comments").Array().Length().Equal(len(Order.OrderComments))
	comment := obj.Value("data").Object().Value("comments").Array().First().Object()
	comment.Value("id").NotNull()
	comment.Value("user_id").Equal(User.ID)
	comment.Value("application_id").Equal(model.AppId)
	comment.Value("content").Equal(Order.OrderComments[0].Content)
	comment.Value("star").Equal(5)
	comment.Value("pics").Array().Length().Equal(2)
	comment.Value("pics").Array().First().Equal("https")
	comment.Value("pics").Array().Last().Equal("https")
	obj.Value("data").Object().Value("menus").Array().Length().Equal(len(Order.OrderMenus))

	orderMenu := obj.Value("data").Object().Value("menus").Array().First().Object()
	orderMenu.Value("id").NotNull()
	orderMenu.Value("menu_name").Equal(Order.OrderMenus[0].MenuName)
	orderMenu.Value("menu_type").Equal(Order.OrderMenus[0].MenuType)
	orderMenu.Value("menu_time_type").Equal(Order.OrderMenus[0].MenuTimeType)
	orderMenu.Value("menu_desc").Equal(Order.OrderMenus[0].MenuDesc)
	orderMenu.Value("menu_id").Equal(Order.OrderMenus[0].MenuID)
	orderMenu.Value("price").Equal(orderMenuPrice)
	orderMenu.Value("amount").Equal(Order.OrderMenus[0].Amount)
	orderMenu.Value("cover").Equal(Order.OrderMenus[0].Cover)

	orderAddr := obj.Value("data").Object().Value("addr").Object()
	orderAddr.Value("id").Equal(Order.OrderAddr.ID)
	orderAddr.Value("name").Equal(Order.OrderAddr.Name)
	orderAddr.Value("sex").Equal(Order.OrderAddr.Sex)
	orderAddr.Value("o_order_id").Equal(Order.OrderAddr.OOrderID)
	orderAddr.Value("loc_name").Equal(Order.OrderAddr.LocName)
	orderAddr.Value("hospital_no").Equal(Order.OrderAddr.HospitalNo)
	orderAddr.Value("hospital_name").Equal(Order.OrderAddr.HospitalName)
	orderAddr.Value("age").Equal(Order.OrderAddr.Age)
	orderAddr.Value("disease").Equal(Order.OrderAddr.Disease)
	orderAddr.Value("phone").Equal(Order.OrderAddr.Phone)
	orderAddr.Value("addr").Equal(Order.OrderAddr.Addr)

	orderReturn := obj.Value("data").Object().Value("return_order").Object()
	orderReturn.Value("id").NotNull()
	orderReturn.Value("order_no").String().Contains("RI")
	orderReturn.Value("status").Object().Value("value").Equal(model.IReturnOrderStatusForAudit)
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
	returnAddr.Value("name").Equal(Order.OrderAddr.Name)
	returnAddr.Value("sex").Equal(Order.OrderAddr.Sex)
	returnAddr.Value("o_return_order_id").NotNull()
	returnAddr.Value("loc_name").Equal(Order.OrderAddr.LocName)
	returnAddr.Value("hospital_no").Equal(Order.OrderAddr.HospitalNo)
	returnAddr.Value("hospital_name").Equal(Order.OrderAddr.HospitalName)
	returnAddr.Value("age").Equal(Order.OrderAddr.Age)
	returnAddr.Value("disease").Equal(Order.OrderAddr.Disease)
	returnAddr.Value("phone").Equal(Order.OrderAddr.Phone)
	returnAddr.Value("addr").Equal(Order.OrderAddr.Addr)

	orderReturn.Value("menus").Array().Length().Equal(len(Order.OrderMenus))
	returnOrderMenu := orderReturn.Value("menus").Array().First().Object()
	returnOrderMenu.Value("id").NotNull()
	returnOrderMenu.Value("menu_name").Equal(Order.OrderMenus[0].MenuName)
	returnOrderMenu.Value("menu_type").Equal(Order.OrderMenus[0].MenuType)
	returnOrderMenu.Value("menu_time_type").Equal(Order.OrderMenus[0].MenuTimeType)
	returnOrderMenu.Value("menu_desc").Equal(Order.OrderMenus[0].MenuDesc)
	returnOrderMenu.Value("menu_id").Equal(Order.OrderMenus[0].MenuID)
	returnOrderMenu.Value("price").Equal(orderMenuPrice)
	returnOrderMenu.Value("amount").Equal(Order.OrderMenus[0].Amount)
	returnOrderMenu.Value("cover").Equal(Order.OrderMenus[0].Cover)
}

func TestOrderDeleteSuccess(t *testing.T) {
	obj := model.GetE(t).DELETE("/api/v1/i_order/{id}", orderId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	model.OrderCount--
}

func TestOrderDeleteReturnSuccess(t *testing.T) {
	obj := model.GetE(t).DELETE("/api/v1/i_order/{id}", Order.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	model.OrderCount--
}

func TestOrderDeleteError(t *testing.T) {
	obj := model.GetE(t).DELETE("/api/v1/i_order/{id}", 99999).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("数据不存在")
}

func TestOrderDeleteErrorForZore(t *testing.T) {
	obj := model.GetE(t).DELETE("/api/v1/i_order/{id}", 0).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("错误数据")
}
