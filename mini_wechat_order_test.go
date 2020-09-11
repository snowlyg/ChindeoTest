package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"strconv"
	"testing"
)

var miniWechatOrderId int

func TestMiniWechatOrderListSuccess(t *testing.T) {
	re := map[string]interface{}{
		"status":         2,
		"page_size":      10,
		"application_id": model.AppId,
	}

	obj := model.GetE(t).POST("/api/v1/outline/o_order").
		WithQuery("page", 1).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(re).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.OrderCount)
}

func TestMiniWechatOrderAddSuccess(t *testing.T) {
	order := map[string]interface{}{
		"amount":         12,
		"total":          32.00,
		"addr_id":        Addr.ID,
		"application_id": model.AppId,
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

	obj := model.GetE(t).POST("/api/v1/outline/o_order/add").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
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
	obj.Value("data").Object().Value("app_type").Number().Equal(model.OrderAppTypeMini)
	obj.Value("data").Object().Value("pay_type").Number().Equal(0)
	obj.Value("data").Object().Value("application_id").Equal(model.AppId)

	id := obj.Value("data").Object().Value("id").Raw()
	data, ok := id.(string)
	if ok {
		miniWechatOrderId, _ = strconv.Atoi(data)
	}
	model.OrderCount++
}

func TestMiniWechatOrderAddNoApplicationId(t *testing.T) {
	order := map[string]interface{}{
		"amount":  12,
		"total":   32.00,
		"addr_id": MiniOrder.OrderAddr.ID,
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

	obj := model.GetE(t).POST("/api/v1/outline/o_order/add").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
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
		"application_id": model.AppId,
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

	obj := model.GetE(t).POST("/api/v1/outline/o_order/add").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(order).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("订单地址错误")

}

func TestMiniWechatOrderAfterOrderAddMenuShowSuccess(t *testing.T) {

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
	obj.Value("data").Object().Value("pics").Array().Length().Equal(0)
	obj.Value("data").Object().Value("create_at").String().Contains(Menu.CreateAt.Format("2006-01-02 15:04"))
	obj.Value("data").Object().Value("create_at").String().Contains(Menu.UpdateAt.Format("2006-01-02 15:04"))

	menuType := obj.Value("data").Object().Value("type").Object()
	menuType.Value("id").Equal(Menu.MenuType.ID)
	menuType.Value("name").Equal(Menu.MenuType.Name)

	obj.Value("data").Object().Value("tags").Array().Length().Equal(len(Menu.MenuTags))
	menuTag := obj.Value("data").Object().Value("tags").Array().First().Object()
	menuTag.Value("id").Equal(Menu.MenuTags[0].ID)
	menuTag.Value("name").Equal(Menu.MenuTags[0].Name)
}

func TestMiniWechatOrderShowSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/api/v1/outline/o_order/{id}", miniWechatOrderId).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(miniWechatOrderId)
	obj.Value("data").Object().Value("order_no").String().Contains("O")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForPay)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("待付款")
	obj.Value("data").Object().Value("amount").Number().Equal(12)
	obj.Value("data").Object().Value("total").String().Equal("32.00")
	obj.Value("data").Object().Value("is_return").String().Equal("未退款")
	obj.Value("data").Object().Value("comments").Array().Length().Equal(0)
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

	obj := model.GetE(t).GET("/api/v1/outline/o_order/pay/{id}", miniWechatOrderId).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("PARAM_ERROR:appid和openid不匹配")
}

func TestMiniWechatOrderCancelSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/api/v1/outline/o_order/cancel/{id}", miniWechatOrderId).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestMiniWechatOrderCancelPaySuccess(t *testing.T) {

	obj := model.GetE(t).GET("/api/v1/outline/o_order/cancel/{id}", MiniOrder.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
}

func TestMiniWechatOrderShowReturnSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/api/v1/outline/o_order/{id}", MiniOrder.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(MiniOrder.ID)
	obj.Value("data").Object().Value("order_no").String().Contains("O")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForCancel)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")
	obj.Value("data").Object().Value("amount").Number().Equal(MiniOrder.Amount)
	obj.Value("data").Object().Value("total").Equal("10.00")
	obj.Value("data").Object().Value("is_return").String().Equal("有退款")

	obj.Value("data").Object().Value("comments").Array().Length().Equal(len(MiniOrder.OrderComments))
	comment := obj.Value("data").Object().Value("comments").Array().First().Object()
	comment.Value("id").NotNull()
	comment.Value("user_id").Equal(User.ID)
	comment.Value("application_id").Equal(model.AppId)
	comment.Value("content").Equal(MiniOrder.OrderComments[0].Content)
	comment.Value("star").Equal(5)
	comment.Value("pics").Array().Length().Equal(2)
	comment.Value("pics").Array().First().Equal("https")
	comment.Value("pics").Array().Last().Equal("https")
	obj.Value("data").Object().Value("menus").Array().Length().Equal(len(MiniOrder.OrderMenus))

	orderMenu := obj.Value("data").Object().Value("menus").Array().First().Object()
	orderMenu.Value("id").NotNull()
	orderMenu.Value("menu_name").Equal(MiniOrder.OrderMenus[0].MenuName)
	orderMenu.Value("menu_type").Equal(MiniOrder.OrderMenus[0].MenuType)
	orderMenu.Value("menu_time_type").Equal(MiniOrder.OrderMenus[0].MenuTimeType)
	orderMenu.Value("menu_desc").Equal(MiniOrder.OrderMenus[0].MenuDesc)
	orderMenu.Value("menu_id").Equal(MiniOrder.OrderMenus[0].MenuID)
	orderMenuPrice := strconv.FormatFloat(MiniOrder.OrderMenus[0].Price, 'g', -1, 64)
	orderMenu.Value("price").Equal(orderMenuPrice)
	orderMenu.Value("amount").Equal(MiniOrder.OrderMenus[0].Amount)
	orderMenu.Value("cover").Equal(MiniOrder.OrderMenus[0].Cover)

	orderAddr := obj.Value("data").Object().Value("addr").Object()
	orderAddr.Value("id").Equal(MiniOrder.OrderAddr.ID)
	orderAddr.Value("name").Equal(MiniOrder.OrderAddr.Name)
	orderAddr.Value("sex").Equal(MiniOrder.OrderAddr.Sex)
	orderAddr.Value("o_order_id").Equal(MiniOrder.OrderAddr.OOrderID)
	orderAddr.Value("loc_name").Equal(MiniOrder.OrderAddr.LocName)
	orderAddr.Value("hospital_no").Equal(MiniOrder.OrderAddr.HospitalNo)
	orderAddr.Value("hospital_name").Equal(MiniOrder.OrderAddr.HospitalName)
	orderAddr.Value("age").Equal(MiniOrder.OrderAddr.Age)
	orderAddr.Value("disease").Equal(MiniOrder.OrderAddr.Disease)
	orderAddr.Value("phone").Equal(MiniOrder.OrderAddr.Phone)
	orderAddr.Value("addr").Equal(MiniOrder.OrderAddr.Addr)

	orderReturn := obj.Value("data").Object().Value("return_order").Object()
	orderReturn.Value("id").NotNull()
	orderReturn.Value("order_no").String().Contains("RO")
	orderReturn.Value("status").Object().Value("value").Equal(model.IReturnOrderStatusForAudit)
	orderReturn.Value("status").Object().Value("text").Equal("待审核")
	orderReturn.Value("amount").Equal(MiniOrder.Amount)
	orderReturn.Value("total").Equal("10")
	orderReturn.Value("open_id").Equal("")
	orderReturn.Value("app_type").Equal(MiniOrder.AppType)
	orderReturn.Value("trade_type").Equal("")
	orderReturn.Value("o_order_id").Equal(MiniOrder.ID)
	orderReturn.Value("application_id").Equal(MiniOrder.ApplicationID)
	orderReturn.Value("pay_type").Equal(MiniOrder.PayType)
	orderReturn.Value("user_id").Equal(MiniOrder.UserID)

	returnAddr := orderReturn.Value("addr").Object()
	returnAddr.Value("id").NotNull()
	returnAddr.Value("name").Equal(MiniOrder.OrderAddr.Name)
	returnAddr.Value("sex").Equal(MiniOrder.OrderAddr.Sex)
	returnAddr.Value("o_return_order_id").NotNull()
	returnAddr.Value("loc_name").Equal(MiniOrder.OrderAddr.LocName)
	returnAddr.Value("hospital_no").Equal(MiniOrder.OrderAddr.HospitalNo)
	returnAddr.Value("hospital_name").Equal(MiniOrder.OrderAddr.HospitalName)
	returnAddr.Value("age").Equal(MiniOrder.OrderAddr.Age)
	returnAddr.Value("disease").Equal(MiniOrder.OrderAddr.Disease)
	returnAddr.Value("phone").Equal(MiniOrder.OrderAddr.Phone)
	returnAddr.Value("addr").Equal(MiniOrder.OrderAddr.Addr)

	orderReturn.Value("menus").Array().Length().Equal(len(MiniOrder.OrderMenus))
	returnOrderMenu := orderReturn.Value("menus").Array().First().Object()
	returnOrderMenu.Value("id").NotNull()
	returnOrderMenu.Value("menu_name").Equal(MiniOrder.OrderMenus[0].MenuName)
	returnOrderMenu.Value("menu_type").Equal(MiniOrder.OrderMenus[0].MenuType)
	returnOrderMenu.Value("menu_time_type").Equal(MiniOrder.OrderMenus[0].MenuTimeType)
	returnOrderMenu.Value("menu_desc").Equal(MiniOrder.OrderMenus[0].MenuDesc)
	returnOrderMenu.Value("menu_id").Equal(MiniOrder.OrderMenus[0].MenuID)
	returnOrderMenu.Value("price").Equal(orderMenuPrice)
	returnOrderMenu.Value("amount").Equal(MiniOrder.OrderMenus[0].Amount)
	returnOrderMenu.Value("cover").Equal(MiniOrder.OrderMenus[0].Cover)
}

func TestMiniWechatOrderDeleteSuccess(t *testing.T) {
	obj := model.GetE(t).DELETE("/api/v1/outline/o_order/{id}", miniWechatOrderId).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	model.OrderCount--
}

func TestMiniWechatOrderDeleteReturnSuccess(t *testing.T) {

	obj := model.GetE(t).DELETE("/api/v1/outline/o_order/{id}", MiniOrder.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	model.OrderCount--
}

func TestMiniWechatOrderDeleteErrorForZore(t *testing.T) {
	obj := model.GetE(t).DELETE("/api/v1/outline/o_order/{id}", 0).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("错误数据")
}

func TestMiniWechatOrderDeleteError(t *testing.T) {
	obj := model.GetE(t).DELETE("/api/v1/outline/o_order/{id}", 99999).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("数据不存在")
}
