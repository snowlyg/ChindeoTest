package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"strconv"
	"testing"
)

var miniShopOrderId float64

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
	spu := model.CreateSpu(brand.ID, Cate2.ID, 3, name, title, Spec)
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

	total := spu.Skus[0].Price*3 + spu.Skus[1].Price*3

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
	obj.Value("data").Object().Value("total").Equal(total)
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("app_type").Object().Value("value").Equal(model.OrderAppTypeMini)
	obj.Value("data").Object().Value("app_type").Object().Value("text").Equal("小程序")
	obj.Value("data").Object().Value("application_id").Equal(13)
	miniShopOrderId, _ = strconv.ParseFloat(model.GetS(obj.Value("data").Object().Value("id").Raw()), 10)
}

func TestMiniWechatShopOrderCommentSuccess(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "content",
		"id_card_no": "456952158962254456",
		"pics":       model.GetPics(),
		"order_id":   miniShopOrderId,
	}

	obj := model.GetE(t).POST("/shop/v1/comment").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("评论成功")
}

func TestMiniWechatShopOrderCommentNoContentError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "",
		"id_card_no": "456952158962254456",
		"pics":       model.GetPics(),
		"order_id":   miniShopOrderId,
	}

	obj := model.GetE(t).POST("/shop/v1/comment").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("评论内容不能为空！")
}

func TestMiniWechatShopOrderCommentOrderNotExistsError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "sdfsdf",
		"id_card_no": "456952158962254456",
		"pics":       model.GetPics(),
		"order_id":   9999,
	}

	obj := model.GetE(t).POST("/shop/v1/comment").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("订单 9999 不存在")
}

func TestMiniWechatShopOrderCommentNoIdCardNoError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "456952158962254456",
		"id_card_no": "",
		"pics":       model.GetPics(),
		"order_id":   miniShopOrderId,
	}

	obj := model.GetE(t).POST("/shop/v1/comment").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("评论成功")
}

func TestMiniWechatShopOrderCommentNoOrderIdError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "456952158962254456",
		"id_card_no": "456952158962254456",
		"pics":       model.GetPics(),
	}

	obj := model.GetE(t).POST("/shop/v1/comment").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("评论订单不能为空！")
}

func TestMiniWechatShopOrderNoAddrError(t *testing.T) {
	brand := model.CreateBrand(false)
	name := "这是一个很神奇的商品"
	title := "这是一个很神奇的商品的超厉害的副标题"
	spu := model.CreateSpu(brand.ID, Cate2.ID, 3, name, title, Spec)
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
	spu := model.CreateSpu(brand.ID, Cate2.ID, 3, name, title, Spec)
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
	obj.Value("data").Object().Value("id").Equal(miniShopOrderId)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForPay)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("待付款")
	obj.Value("data").Object().Value("start_at").Equal(miniCareStartAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("end_at").Equal(miniCareEndAt.Format("2006-01-02 15:04:05"))

	var total decimal.Decimal
	sub := int(miniCareEndAt.Sub(miniCareStartAt).Hours())
	miniCarePrice := decimal.NewFromFloat(model.CareMaxPrice)
	timeType := model.GetS(obj.Value("data").Object().Value("order_carer_info").Object().Value("time_type").Raw())
	if timeType == "天" {
		total = miniCarePrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if timeType == "时" {
		total = miniCarePrice.Mul(decimal.NewFromFloat(float64(sub)))
	}

	f, _ := total.Float64()

	obj.Value("data").Object().Value("total").Equal(model.Ftos(f))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("pay_type").Equal(1)
	obj.Value("data").Object().Value("is_return").Equal(0)
	obj.Value("data").Object().Value("order_carer_info").Null()
	obj.Value("data").Object().Value("return_order").Null()
	obj.Value("data").Object().Value("comments").Array().Length().Equal(2)
	orderInfo := obj.Value("data").Object().Value("order_info").Object()
	orderInfo.Value("id").NotNull()
	orderInfo.Value("name").Equal(MiniCareOrder.CareOrderInfo.Name)
	orderInfo.Value("desc").Equal(MiniCareOrder.CareOrderInfo.Desc)
	orderInfo.Value("application_name").Equal("我的医院")
	orderInfo.Value("time_type").Equal(MiniCareOrder.CareOrderInfo.TimeType)
	orderInfo.Value("care_detail").Equal(MiniCareOrder.CareOrderInfo.CareDetail)
	orderInfo.Value("care_tags").Equal(MiniCareOrder.CareOrderInfo.CareTags)
	orderInfo.Value("min_price").Equal(model.Ftos(MiniCareOrder.CareOrderInfo.MinPrice))
	orderInfo.Value("max_price").Equal(model.Ftos(MiniCareOrder.CareOrderInfo.MaxPrice))
	orderInfo.Value("cover").Equal(MiniCareOrder.CareOrderInfo.Cover)
	orderInfo.Value("care_type").Equal(MiniCareOrder.CareOrderInfo.CareType)
	orderInfo.Value("care_order_id").Equal(miniShopOrderId)

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal(Addr.Name)
	addr.Value("sex").Equal(Addr.Sex)
	addr.Value("care_order_id").Equal(miniShopOrderId)
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
	obj := model.GetE(t).GET("/shop/v1/order/cancel/{id}", MiniCareOrder.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatShopOrderShowReturnSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/shop/v1/order/{id}", MiniCareOrder.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(MiniCareOrder.ID)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForCancel)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")
	obj.Value("data").Object().Value("total").Equal(model.Ftos(MiniCareOrder.Total))
	obj.Value("data").Object().Value("start_at").String().Contains(MiniCareOrder.StartAt.Format("2006-01-02 15:04"))
	obj.Value("data").Object().Value("end_at").String().Contains(MiniCareOrder.EndAt.Format("2006-01-02 15:04"))

	var total decimal.Decimal
	sub := int(MiniCareOrder.EndAt.Sub(MiniCareOrder.StartAt).Hours())
	maxCarePrice := decimal.NewFromFloat(Care.MaxPrice)
	if Care.TimeType == "天" {
		total = maxCarePrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if Care.TimeType == "时" {
		total = maxCarePrice.Mul(decimal.NewFromFloat(float64(sub)))
	}
	f, _ := total.Float64()

	obj.Value("data").Object().Value("total").Equal(fmt.Sprintf("%.2f", f))
	obj.Value("data").Object().Value("rmk").Equal(MiniCareOrder.Rmk)
	obj.Value("data").Object().Value("pay_type").Equal(MiniCareOrder.PayType)
	obj.Value("data").Object().Value("is_return").Equal(1)
	obj.Value("data").Object().Value("order_carer_info").Null()
	obj.Value("data").Object().Value("comments").Array().Length().Equal(len(MiniCareOrder.CareOrderComments))
	comment := obj.Value("data").Object().Value("comments").Array().First().Object()
	comment.Value("id").NotNull()
	comment.Value("user_id").Equal(User.ID)
	comment.Value("application_id").Equal(model.AppId)
	comment.Value("content").Equal(MiniCareOrder.CareOrderComments[0].Content)
	comment.Value("star").Equal(MiniCareOrder.CareOrderComments[0].Star)
	comment.Value("pics").Array().Length().Equal(2)
	comment.Value("pics").Array().First().Equal("https")
	comment.Value("pics").Array().Last().Equal("https")

	orderInfo := obj.Value("data").Object().Value("order_info").Object()
	orderInfo.Value("id").NotNull()
	orderInfo.Value("name").Equal(MiniCareOrder.CareOrderInfo.Name)
	orderInfo.Value("desc").Equal(MiniCareOrder.CareOrderInfo.Desc)
	orderInfo.Value("application_name").Equal(MiniCareOrder.CareOrderInfo.ApplicationName)
	orderInfo.Value("time_type").Equal(MiniCareOrder.CareOrderInfo.TimeType)
	orderInfo.Value("care_detail").Equal(MiniCareOrder.CareOrderInfo.CareDetail)
	orderInfo.Value("care_tags").Equal(MiniCareOrder.CareOrderInfo.CareTags)
	orderInfo.Value("min_price").Equal(model.Ftos(MiniCareOrder.CareOrderInfo.MinPrice))
	orderInfo.Value("max_price").Equal(model.Ftos(MiniCareOrder.CareOrderInfo.MaxPrice))
	orderInfo.Value("cover").Equal(MiniCareOrder.CareOrderInfo.Cover)
	orderInfo.Value("care_type").Equal(MiniCareOrder.CareOrderInfo.CareType)
	orderInfo.Value("care_order_id").Equal(MiniCareOrder.ID)

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal(MiniCareOrder.CareOrderAddr.Name)
	addr.Value("loc_name").Equal(MiniCareOrder.CareOrderAddr.LocName)
	addr.Value("bed_num").Equal(MiniCareOrder.CareOrderAddr.BedNum)
	addr.Value("hospital_no").Equal(MiniCareOrder.CareOrderAddr.HospitalNo)
	addr.Value("disease").Equal(MiniCareOrder.CareOrderAddr.Disease)
	addr.Value("care_order_id").Equal(MiniCareOrder.CareOrderAddr.CareOrderID)
	addr.Value("sex").Equal(MiniCareOrder.CareOrderAddr.Sex)
	addr.Value("hospital_name").Equal(MiniCareOrder.CareOrderAddr.HospitalName)
	addr.Value("phone").Equal(MiniCareOrder.CareOrderAddr.Phone)
	addr.Value("age").Equal(MiniCareOrder.CareOrderAddr.Age)
	addr.Value("addr").Equal(MiniCareOrder.CareOrderAddr.Addr)

	orderReturn := obj.Value("data").Object().Value("return_order").Object()
	orderReturn.Value("id").NotNull()
	orderReturn.Value("order_no").String().Contains("RC")
	orderReturn.Value("status").Object().Value("value").Equal(model.IReturnOrderStatusForAudit)
	orderReturn.Value("status").Object().Value("text").Equal("待审核")
	orderReturn.Value("total").Equal(model.Ftos(MiniCareOrder.Total))
	orderReturn.Value("open_id").Equal("")
	orderReturn.Value("app_type").Equal(MiniCareOrder.AppType)
	orderReturn.Value("trade_type").Equal("")
	orderReturn.Value("care_order_id").Equal(MiniCareOrder.ID)
	orderReturn.Value("application_id").Equal(MiniCareOrder.ApplicationID)
	orderReturn.Value("pay_type").Equal(MiniCareOrder.PayType)
	orderReturn.Value("user_id").Equal(MiniCareOrder.UserID)
	orderReturn.Value("carer_id").Equal(MiniCareOrder.CarerID)

	returnAddr := orderReturn.Value("addr").Object()
	returnAddr.Value("id").NotNull()
	returnAddr.Value("name").Equal(MiniCareOrder.CareOrderAddr.Name)
	returnAddr.Value("sex").Equal(MiniCareOrder.CareOrderAddr.Sex)
	returnAddr.Value("care_return_order_id").NotNull()
	returnAddr.Value("loc_name").Equal(MiniCareOrder.CareOrderAddr.LocName)
	returnAddr.Value("hospital_no").Equal(MiniCareOrder.CareOrderAddr.HospitalNo)
	returnAddr.Value("hospital_name").Equal(MiniCareOrder.CareOrderAddr.HospitalName)
	returnAddr.Value("age").Equal(MiniCareOrder.CareOrderAddr.Age)
	returnAddr.Value("disease").Equal(MiniCareOrder.CareOrderAddr.Disease)
	returnAddr.Value("phone").Equal(MiniCareOrder.CareOrderAddr.Phone)
	returnAddr.Value("addr").Equal(MiniCareOrder.CareOrderAddr.Addr)

	returnOrderInfo := orderReturn.Value("order_info").Object()
	returnOrderInfo.Value("id").NotNull()
	returnOrderInfo.Value("name").Equal(MiniCareOrder.CareOrderInfo.Name)
	returnOrderInfo.Value("desc").Equal(MiniCareOrder.CareOrderInfo.Desc)
	returnOrderInfo.Value("application_name").Equal(MiniCareOrder.CareOrderInfo.ApplicationName)
	returnOrderInfo.Value("time_type").Equal(MiniCareOrder.CareOrderInfo.TimeType)
	returnOrderInfo.Value("care_detail").Equal(MiniCareOrder.CareOrderInfo.CareDetail)
	returnOrderInfo.Value("care_tags").Equal(MiniCareOrder.CareOrderInfo.CareTags)
	returnOrderInfo.Value("min_price").Equal(model.Ftos(MiniCareOrder.CareOrderInfo.MinPrice))
	returnOrderInfo.Value("max_price").Equal(model.Ftos(MiniCareOrder.CareOrderInfo.MaxPrice))
	returnOrderInfo.Value("cover").Equal(MiniCareOrder.CareOrderInfo.Cover)
	returnOrderInfo.Value("care_type").Equal(MiniCareOrder.CareOrderInfo.CareType)
	returnOrderInfo.Value("care_return_order_id").NotNull()

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
	obj := model.GetE(t).DELETE("/shop/v1/order/{id}", MiniCareOrder.ID).
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	model.CareOrderCount--
}
