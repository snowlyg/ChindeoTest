package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

var miniShopOrderId interface{}
var miniShopOrdeTotal float64
var miniShopOrderSkuId interface{}

func TestMiniWechatShopOrderSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/order").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("application_id", model.AppId).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.ShopOrderCount)
}
func TestMiniWechatShopOrderAddSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	name := "这是一个很神奇的商品"
	title := "这是一个很神奇的商品的超厉害的副标题"
	spu := model.CreateSpu(brand.ID, Cate1.ID, 3, name, title, Spec)
	shopOrder := map[string]interface{}{
		"sku_ids": []map[string]interface{}{
			{
				"id":  spu.Skus[0].ID,
				"num": 3,
			},
			{
				"id":  spu.Skus[1].ID,
				"num": 3,
			},
		},
		"rmk":            "年轻貌美",
		"application_id": model.AppId,
		"addr_id":        Addr.ID,
	}

	miniShopOrdeTotal = spu.Skus[0].Price*3 + spu.Skus[1].Price*3

	obj := model.GetE(t).POST("/shop/v1/order/add").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(shopOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").NotEqual(0)
	obj.Value("data").Object().Value("order_no").String().Contains("S")
	obj.Value("data").Object().Value("total").Equal(miniShopOrdeTotal)
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("app_type").Object().Value("value").Equal(model.OrderAppTypeMini)
	obj.Value("data").Object().Value("app_type").Object().Value("text").Equal("小程序")
	obj.Value("data").Object().Value("application_id").Equal(13)
	miniShopOrderId = obj.Value("data").Object().Value("id").Raw()
}

func TestMiniWechatShopOrderNoAddrError(t *testing.T) {
	brand := model.CreateBrand(false)
	name := "这是一个很神奇的商品"
	title := "这是一个很神奇的商品的超厉害的副标题"
	spu := model.CreateSpu(brand.ID, Cate1.ID, 3, name, title, Spec)
	shopOrder := map[string]interface{}{
		"sku_ids": []map[string]interface{}{
			{
				"id":  spu.Skus[0].ID,
				"num": 3,
			},
			{
				"id":  spu.Skus[1].ID,
				"num": 3,
			},
		},
		"rmk":            "年轻貌美",
		"application_id": model.AppId,
	}

	obj := model.GetE(t).POST("/shop/v1/order/add").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(shopOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("收货地址不存在")
}

func TestMiniWechatShopOrderNoAppError(t *testing.T) {
	brand := model.CreateBrand(false)
	name := "这是一个很神奇的商品"
	title := "这是一个很神奇的商品的超厉害的副标题"
	spu := model.CreateSpu(brand.ID, Cate1.ID, 3, name, title, Spec)
	shopOrder := map[string]interface{}{
		"sku_ids": []map[string]interface{}{
			{
				"id":  spu.Skus[0].ID,
				"num": 3,
			},
			{
				"id":  spu.Skus[1].ID,
				"num": 3,
			},
		},
		"rmk":     "年轻貌美",
		"addr_id": Addr.ID,
	}
	obj := model.GetE(t).POST("/shop/v1/order/add").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(shopOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("医院不能为空！")
}

func TestMiniWechatShopOrderShowSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/order/{id}", miniShopOrderId).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(model.GetSToI(miniShopOrderId))
	obj.Value("data").Object().Value("order_no").String().Contains("S")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForPay)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("待付款")

	obj.Value("data").Object().Value("total").Equal(model.Ftos(miniShopOrdeTotal))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("pay_type").Object().Value("value").Equal(model.IOrderPayTypeWechat)
	obj.Value("data").Object().Value("pay_type").Object().Value("text").Equal("微信")
	obj.Value("data").Object().Value("is_return").Object().Value("value").Equal(0)
	obj.Value("data").Object().Value("is_return").Object().Value("text").Equal("未退款")

	obj.Value("data").Object().Value("skus").Array().Length().Equal(2)
	sku := obj.Value("data").Object().Value("skus").Array().First().Object()

	miniShopOrderSkuId = sku.Value("id").Raw()

	sku.Value("id").NotNull()
	sku.Value("title").NotNull()
	sku.Value("sku_no").NotNull()
	sku.Value("images").Array().NotEmpty()
	sku.Value("price").NotNull()
	sku.Value("indexes").NotNull()
	sku.Value("own_spec").NotNull()
	sku.Value("num").NotNull()
	sku.Value("shop_order_id").NotNull()
	sku.Value("sku_id").NotNull()
	sku.Value("application_id").NotNull()
	sku.Value("comments").NotNull()

	obj.Value("data").Object().Value("return_order").Null()

	obj.Value("data").Object().Value("comments").Array().Length().Equal(0)

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal(Addr.Name)
	addr.Value("sex").Equal(Addr.Sex)
	addr.Value("shop_order_id").Equal(model.GetSToI(miniShopOrderId))
	addr.Value("loc_name").Equal(Addr.LocName)
	addr.Value("hospital_no").Equal(Addr.HospitalNo)
	addr.Value("hospital_name").Equal(Addr.HospitalName)
	addr.Value("age").Equal(Addr.Age)
	addr.Value("disease").Equal(Addr.Disease)
	addr.Value("phone").Equal(Addr.Phone)
	addr.Value("addr").Equal(Addr.Addr)
}

func TestMiniWechatShopOrderPaySuccess(t *testing.T) {

	obj := model.GetE(t).GET("/shop/v1/order/pay/{id}", miniShopOrderId).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Contains("PARAM_ERROR:appid和openid不匹配")
}

func TestMiniWechatShopOrderCancelNoPaySuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/order/cancel/{id}", miniShopOrderId).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatShopOrderCancelPaySuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/order/cancel/{id}", MiniShopOrder.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatShopOrderShowReturnSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/shop/v1/order/{id}", MiniShopOrder.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(MiniShopOrder.ID)
	obj.Value("data").Object().Value("order_no").String().Contains("S")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForCancel)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")
	obj.Value("data").Object().Value("total").Equal(model.Ftos(MiniShopOrder.Total))

	obj.Value("data").Object().Value("total").Equal("0.00")
	obj.Value("data").Object().Value("rmk").Equal(MiniShopOrder.Rmk)
	obj.Value("data").Object().Value("pay_type").Object().Value("value").Equal(MiniShopOrder.PayType)
	obj.Value("data").Object().Value("pay_type").Object().Value("text").Equal("支付宝")
	obj.Value("data").Object().Value("is_return").Object().Value("value").Equal(1)
	obj.Value("data").Object().Value("is_return").Object().Value("text").Equal("有退款")

	obj.Value("data").Object().Value("skus").Array().Length().Equal(1)
	sku := obj.Value("data").Object().Value("skus").Array().First().Object()
	sku.Value("id").NotNull()
	sku.Value("title").NotNull()
	sku.Value("sku_no").NotNull()
	sku.Value("images").Array().NotEmpty()
	sku.Value("price").NotNull()
	sku.Value("indexes").NotNull()
	sku.Value("own_spec").NotNull()
	sku.Value("num").NotNull()
	sku.Value("shop_order_id").NotNull()
	sku.Value("sku_id").NotNull()
	sku.Value("application_id").NotNull()
	sku.Value("comments").NotNull()

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal(MiniShopOrder.Addr.Name)
	addr.Value("loc_name").Equal(MiniShopOrder.Addr.LocName)
	addr.Value("bed_num").Equal(MiniShopOrder.Addr.BedNum)
	addr.Value("hospital_no").Equal(MiniShopOrder.Addr.HospitalNo)
	addr.Value("disease").Equal(MiniShopOrder.Addr.Disease)
	addr.Value("shop_order_id").Equal(MiniShopOrder.Addr.ShopOrderID)
	addr.Value("sex").Equal(MiniShopOrder.Addr.Sex)
	addr.Value("hospital_name").Equal(MiniShopOrder.Addr.HospitalName)
	addr.Value("phone").Equal(MiniShopOrder.Addr.Phone)
	addr.Value("age").Equal(MiniShopOrder.Addr.Age)
	addr.Value("addr").Equal(MiniShopOrder.Addr.Addr)

	orderReturn := obj.Value("data").Object().Value("return_order").Object()
	orderReturn.Value("id").NotNull()
	orderReturn.Value("order_no").String().Contains("SC")
	orderReturn.Value("status").Object().Value("value").Equal(model.IReturnOrderStatusForAudit)
	orderReturn.Value("status").Object().Value("text").Equal("待审核")
	orderReturn.Value("total").Equal(model.Ftos(MiniShopOrder.Total))
	orderReturn.Value("open_id").Equal("")
	orderReturn.Value("app_type").Equal(MiniShopOrder.AppType)
	orderReturn.Value("trade_type").Equal("")
	orderReturn.Value("shop_order_id").Equal(MiniShopOrder.ID)
	orderReturn.Value("application_id").Equal(MiniShopOrder.ApplicationID)
	orderReturn.Value("pay_type").Equal(MiniShopOrder.PayType)
	orderReturn.Value("user_id").Equal(MiniShopOrder.UserID)

	returnAddr := orderReturn.Value("addr").Object()
	returnAddr.Value("id").NotNull()
	returnAddr.Value("name").Equal(MiniShopOrder.Addr.Name)
	returnAddr.Value("sex").Equal(MiniShopOrder.Addr.Sex)
	returnAddr.Value("care_return_order_id").NotNull()
	returnAddr.Value("loc_name").Equal(MiniShopOrder.Addr.LocName)
	returnAddr.Value("hospital_no").Equal(MiniShopOrder.Addr.HospitalNo)
	returnAddr.Value("hospital_name").Equal(MiniShopOrder.Addr.HospitalName)
	returnAddr.Value("age").Equal(MiniShopOrder.Addr.Age)
	returnAddr.Value("disease").Equal(MiniShopOrder.Addr.Disease)
	returnAddr.Value("phone").Equal(MiniShopOrder.Addr.Phone)
	returnAddr.Value("addr").Equal(MiniShopOrder.Addr.Addr)

	returnOrderInfo := orderReturn.Value("skus").Object()
	returnOrderInfo.Value("data").Object().Value("skus").Array().Length().Equal(1)
	returnSku := returnOrderInfo.Value("data").Object().Value("skus").Array().First().Object()
	returnSku.Value("id").NotNull()
	returnSku.Value("title").NotNull()
	returnSku.Value("sku_no").NotNull()
	returnSku.Value("images").Array().NotEmpty()
	returnSku.Value("price").NotNull()
	returnSku.Value("indexes").NotNull()
	returnSku.Value("own_spec").NotNull()
	returnSku.Value("num").NotNull()
	returnSku.Value("shop_order_id").NotNull()
	returnSku.Value("sku_id").NotNull()
	returnSku.Value("application_id").NotNull()
	returnSku.Value("comments").NotNull()
}

func TestMiniWechatShopOrderCommentSuccess(t *testing.T) {
	comment := map[string]interface{}{
		"star":           1,
		"content":        "content",
		"pics":           model.GetPics(),
		"is_unnamed":     1,
		"application_id": model.AppId,
	}

	obj := model.GetE(t).POST("/shop/v1/comment/{shop_order_sku_id}", miniShopOrderSkuId).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("评论成功")
}

func TestMiniWechatShopOrderShowAfterCommentSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/order/{id}", miniShopOrderId).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(model.GetSToI(miniShopOrderId))
	obj.Value("data").Object().Value("order_no").String().Contains("S")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForCancel)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")

	obj.Value("data").Object().Value("total").Equal(model.Ftos(miniShopOrdeTotal))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("pay_type").Object().Value("value").Equal(model.IOrderPayTypeWechat)
	obj.Value("data").Object().Value("pay_type").Object().Value("text").Equal("微信")
	obj.Value("data").Object().Value("is_return").Object().Value("value").Equal(0)
	obj.Value("data").Object().Value("is_return").Object().Value("text").Equal("未退款")

	obj.Value("data").Object().Value("skus").Array().Length().Equal(2)
	sku := obj.Value("data").Object().Value("skus").Array().First().Object()

	sku.Value("id").NotNull()
	sku.Value("title").NotNull()
	sku.Value("sku_no").NotNull()
	sku.Value("images").Array().NotEmpty()
	sku.Value("price").NotNull()
	sku.Value("indexes").NotNull()
	sku.Value("own_spec").NotNull()
	sku.Value("num").NotNull()
	sku.Value("shop_order_id").NotNull()
	sku.Value("sku_id").NotNull()
	sku.Value("application_id").NotNull()
	sku.Value("comments").NotNull()

	obj.Value("data").Object().Value("return_order").Null()

	obj.Value("data").Object().Value("comments").Array().Length().Equal(1)
	comments := obj.Value("data").Object().Value("comments").Array().First().Object()
	comments.Value("id").NotNull()

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal(Addr.Name)
	addr.Value("sex").Equal(Addr.Sex)
	addr.Value("shop_order_id").Equal(model.GetSToI(miniShopOrderId))
	addr.Value("loc_name").Equal(Addr.LocName)
	addr.Value("hospital_no").Equal(Addr.HospitalNo)
	addr.Value("hospital_name").Equal(Addr.HospitalName)
	addr.Value("age").Equal(Addr.Age)
	addr.Value("disease").Equal(Addr.Disease)
	addr.Value("phone").Equal(Addr.Phone)
	addr.Value("addr").Equal(Addr.Addr)
}

func TestMiniWechatShopOrderCommentNoContentError(t *testing.T) {
	comment := map[string]interface{}{
		"star":           1,
		"content":        "",
		"pics":           model.GetPics(),
		"is_unnamed":     0,
		"application_id": model.AppId,
	}

	obj := model.GetE(t).POST("/shop/v1/comment/{shop_order_sku_id}", miniShopOrderSkuId).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("评论内容不能为空！")
}

func TestMiniWechatShopOrderCommentNoOrderIdError(t *testing.T) {
	comment := map[string]interface{}{
		"star":           1,
		"content":        "456952158962254456",
		"pics":           model.GetPics(),
		"is_unnamed":     0,
		"application_id": model.AppId,
	}

	obj := model.GetE(t).POST("/shop/v1/comment/{shop_order_sku_id}", 0).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("参数错误")
}

func TestMiniWechatShopOrderDeleteSuccess(t *testing.T) {

	obj := model.GetE(t).DELETE("/shop/v1/order/{id}", miniShopOrderId).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatShopOrderDeleteRetrunSuccess(t *testing.T) {
	obj := model.GetE(t).DELETE("/shop/v1/order/{id}", MiniShopOrder.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	model.CareOrderCount--
}
