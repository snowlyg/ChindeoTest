package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

var careOrderCareId int
var carePrice decimal.Decimal
var careTimeTypeText string

func TestCareNoTagIdListSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/care/v1/inner/care").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.CareNoTagCount)

	fitst := obj.Value("data").Object().Value("data").Array().Last().Object()
	fitst.Value("id").Equal(Care.ID)
	fitst.Value("name").Equal(Care.Name)
	fitst.Value("desc").Equal(Care.Desc)
	fitst.Value("time_type").Equal(Care.TimeType)
	fitst.Value("amount").Equal(Care.Amount)
	fitst.Value("cover").Equal(Care.Cover)
	fitst.Value("care_type_id").Equal(Care.CareTypeID)

	careType := fitst.Value("type").Object()
	careType.Value("id").Equal(CareType.ID)
	careType.Value("name").Equal(CareType.Name)
}

func TestCareListSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/care/v1/inner/care").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("care_type_id", CareType.ID).
		WithQuery("care_tag_id", CareTag.ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.CarerCount)

	fitst := obj.Value("data").Object().Value("data").Array().First().Object()
	fitst.Value("id").Equal(Care.ID)
	fitst.Value("name").Equal(Care.Name)
	fitst.Value("desc").Equal(Care.Desc)
	fitst.Value("time_type").Equal(Care.TimeType)
	fitst.Value("amount").Equal(Care.Amount)
	fitst.Value("cover").Equal(Care.Cover)
	fitst.Value("care_type_id").Equal(Care.CareTypeID)

	careType := fitst.Value("type").Object()
	careType.Value("id").Equal(CareType.ID)
	careType.Value("name").Equal(CareType.Name)
}

func TestCareListNoPageSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/inner/care").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("care_type_id", CareType.ID).
		WithQuery("care_tag_id", CareTag.ID).
		WithQuery("page_size", "-1").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(model.CarerCount)

	fitst := obj.Value("data").Array().Last().Object()
	fitst.Value("id").Equal(Care.ID)
	fitst.Value("name").Equal(Care.Name)
	fitst.Value("desc").Equal(Care.Desc)
	fitst.Value("time_type").Equal(Care.TimeType)
	fitst.Value("amount").Equal(Care.Amount)
	fitst.Value("cover").Equal(Care.Cover)
	fitst.Value("care_type_id").Equal(Care.CareTypeID)

	careType := fitst.Value("type").Object()
	careType.Value("id").Equal(CareType.ID)
	careType.Value("name").Equal(CareType.Name)
}

func TestCareShowSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/inner/care/{id}", Care.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(Care.ID)
	obj.Value("data").Object().Value("name").Equal(Care.Name)
	obj.Value("data").Object().Value("desc").Equal(Care.Desc)
	obj.Value("data").Object().Value("time_type").Equal(Care.TimeType)
	obj.Value("data").Object().Value("amount").Equal(model.CareAmount)
	obj.Value("data").Object().Value("cover").Equal(Care.Cover)
	obj.Value("data").Object().Value("care_type_id").Equal(Care.CareTypeID)

	careType := obj.Value("data").Object().Value("type").Object()
	careType.Value("id").Equal(CareType.ID)
	careType.Value("name").Equal(CareType.Name)
}

func TestCareOrderListStatusForDeliveryingSuccess(t *testing.T) {
	careOrder := model.CreateCareOrder(Care.TimeType, "IC202008241612348468756914", User.ID, 0, model.OrderAppTypeMini, model.IOrderPayTypeAli, model.IOrderStatusForDeliverying, model.CareMaxPrice)

	re := map[string]interface{}{
		"status":      model.IOrderStatusForDeliverying,
		"page_size":   10,
		"hospital_no": "9556854545",
		"id_card_no":  model.IdCardNo,
	}
	obj := model.GetE(t).GET("/care/v1/inner/order").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQueryObject(re).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(1)
	model.DelCareOrder(careOrder)
}

func TestCareOrderListStatusForDeliveryingNoIdCardNoSuccess(t *testing.T) {
	careOrder := model.CreateCareOrder(Care.TimeType, "IC202008241612348468756914", User.ID, 0, model.OrderAppTypeMini, model.IOrderPayTypeAli, model.IOrderStatusForDeliverying, model.CareMaxPrice)

	re := map[string]interface{}{
		"status":    model.IOrderStatusForDeliverying,
		"page_size": 10,
	}
	obj := model.GetE(t).GET("/care/v1/inner/order").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQueryObject(re).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(0)
	model.DelCareOrder(careOrder)
}

func TestCareOrderListNoStatusSuccess(t *testing.T) {
	careOrder := model.CreateCareOrder(Care.TimeType, "IC202008241612348468756914", User.ID, 0, model.OrderAppTypeMini, model.IOrderPayTypeAli, model.IOrderStatusForDeliverying, model.CareMaxPrice)

	re := map[string]interface{}{
		"status":      0,
		"page_size":   10,
		"hospital_no": "9556854545",
		"id_card_no":  model.IdCardNo,
	}
	obj := model.GetE(t).GET("/care/v1/inner/order").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQueryObject(re).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.CareOrderCount)
	model.DelCareOrder(careOrder)
}

func TestCareOrderListStatusForFinishSuccess(t *testing.T) {
	careOrder := model.CreateCareOrder(Care.TimeType, "IC202008241612348468756914", User.ID, 0, model.OrderAppTypeMini, model.IOrderPayTypeAli, model.IOrderStatusForDeliverying, model.CareMaxPrice)

	re := map[string]interface{}{
		"status":      model.IOrderStatusForFinish,
		"page_size":   10,
		"hospital_no": "9556854545",
		"id_card_no":  model.IdCardNo,
	}
	obj := model.GetE(t).GET("/care/v1/inner/order").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQueryObject(re).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(0)
	model.DelCareOrder(careOrder)
}

var careStartAt time.Time
var careEndAt time.Time

func TestCareOrderAddCareSuccess(t *testing.T) {

	careStartAt = time.Now().AddDate(0, 0, 1)
	careEndAt = time.Now().AddDate(0, 0, 2)
	careOrder := map[string]interface{}{
		"patient_name": "操蛋",
		"loc_name":     "泥马",
		"bed_num":      "05",
		"hospital_no":  "9556854545",
		"disease":      "玩玩",
		"care_id":      Care.ID,
		"start_at":     careStartAt.Format("2006-01-02 15:04:05"),
		"end_at":       careEndAt.Format("2006-01-02 15:04:05"),
		"rmk":          "年轻貌美",
		"id_card_no":   model.IdCardNo,
	}

	obj := model.GetE(t).POST("/care/v1/inner/order/add").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(careOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").NotEqual(0)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("start_at").Equal(careStartAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("end_at").Equal(careEndAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("app_type").Equal(model.OrderAppTypeBed)
	obj.Value("data").Object().Value("application_id").Equal(13)
	careOrderCareId = model.GetSToI(obj.Value("data").Object().Value("id").Raw())
}

func TestCareShowAfterOrderAddSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/inner/care/{id}", Care.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(Care.ID)
	obj.Value("data").Object().Value("name").Equal(Care.Name)
	obj.Value("data").Object().Value("desc").Equal(Care.Desc)
	obj.Value("data").Object().Value("time_type").Equal(Care.TimeType)
	obj.Value("data").Object().Value("amount").Equal(model.CareAmount + 1)
	model.CareAmount++
	obj.Value("data").Object().Value("cover").Equal(Care.Cover)
	obj.Value("data").Object().Value("care_type_id").Equal(Care.CareTypeID)

	careType := obj.Value("data").Object().Value("type").Object()
	careType.Value("id").Equal(CareType.ID)
	careType.Value("name").Equal(CareType.Name)

	carePriceF := model.GetSToF(obj.Value("data").Object().Value("max_price").Raw())
	carePrice = decimal.NewFromFloat(carePriceF)
	careTimeTypeText = model.GetS(obj.Value("data").Object().Value("time_type").Raw())

}

func TestCareOrderAddCareError(t *testing.T) {
	startAt := time.Now().AddDate(0, 0, 5)
	endAt := time.Now().AddDate(0, 0, 6)
	careOrder := map[string]interface{}{
		"patient_name": "",
		"loc_name":     "泥马",
		"bed_num":      "05",
		"hospital_no":  "9556854545",
		"disease":      "玩玩",
		"start_at":     startAt.Format("2006-01-02 15:04:05"),
		"end_at":       endAt.Format("2006-01-02 15:04:05"),
		"rmk":          "年轻貌美",
		"care_id":      Care.ID,
		"id_card_no":   model.IdCardNo,
	}

	obj := model.GetE(t).POST("/care/v1/inner/order/add").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(careOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("患者名称不能为空！")
}

func TestCareOrderCommentSuccess(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "content",
		"id_card_no": model.IdCardNo,
		"pics":       model.GetPics(),
		"order_id":   careOrderCareId,
	}

	obj := model.GetE(t).POST("/common/v1/inner/comment/care").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("评论成功")
}

func TestCareOrderCommentNoContentError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "",
		"id_card_no": model.IdCardNo,
		"pics":       model.GetPics(),
		"order_id":   careOrderCareId,
	}

	obj := model.GetE(t).POST("/common/v1/inner/comment/care").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("评论内容不能为空！")
}

func TestCareOrderCommentOrderExistsError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "sdfsdfs",
		"id_card_no": model.IdCardNo,
		"pics":       model.GetPics(),
		"order_id":   9999,
	}

	obj := model.GetE(t).POST("/common/v1/inner/comment/care").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("订单 9999 不存在")
}

func TestCareOrderCommentNoIdCardNoError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "456952158962254456",
		"id_card_no": "",
		"pics":       model.GetPics(),
		"order_id":   careOrderCareId,
	}

	obj := model.GetE(t).POST("/common/v1/inner/comment/care").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("评论成功")
}

func TestCareOrderCommentNoOrderIdError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "456952158962254456",
		"id_card_no": model.IdCardNo,
		"pics":       model.GetPics(),
	}

	obj := model.GetE(t).POST("/common/v1/inner/comment/care").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("评论订单不能为空！")
}

func TestCareOrderShowCareSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/inner/order/{id}", careOrderCareId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(careOrderCareId)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForPay)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("待付款")
	obj.Value("data").Object().Value("start_at").Equal(careStartAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("end_at").Equal(careEndAt.Format("2006-01-02 15:04:05"))

	var total decimal.Decimal
	careStartAt = time.Now().AddDate(0, 0, 1)
	careEndAt = time.Now().AddDate(0, 0, 2)
	sub := int(careEndAt.Sub(careStartAt).Hours())
	if careTimeTypeText == "天" {
		total = carePrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if careTimeTypeText == "时" {
		total = carePrice.Mul(decimal.NewFromFloat(float64(sub)))
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
	orderInfo.Value("name").Equal(CareOrder.CareOrderInfo.Name)
	orderInfo.Value("desc").Equal(CareOrder.CareOrderInfo.Desc)
	orderInfo.Value("application_name").Equal("我的医院")
	orderInfo.Value("time_type").Equal(CareOrder.CareOrderInfo.TimeType)
	orderInfo.Value("care_detail").Equal(CareOrder.CareOrderInfo.CareDetail)
	orderInfo.Value("care_tags").Equal(CareOrder.CareOrderInfo.CareTags)
	orderInfo.Value("min_price").Equal(model.Ftos(CareOrder.CareOrderInfo.MinPrice))
	orderInfo.Value("max_price").Equal(model.Ftos(CareOrder.CareOrderInfo.MaxPrice))
	orderInfo.Value("cover").Equal(CareOrder.CareOrderInfo.Cover)
	orderInfo.Value("care_type").Equal(CareOrder.CareOrderInfo.CareType)
	orderInfo.Value("care_order_id").Equal(careOrderCareId)

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("name").Equal("操蛋")
	addr.Value("loc_name").Equal("泥马")
	addr.Value("bed_num").Equal("05")
	addr.Value("hospital_no").Equal("9556854545")
	addr.Value("disease").Equal("玩玩")
	addr.Value("care_order_id").Equal(careOrderCareId)
	addr.Value("sex").Equal(1)
	addr.Value("hospital_name").Equal("我的医院")
	addr.Value("phone").Equal("")
	addr.Value("age").Equal(0)
	addr.Value("addr").Equal("")
}

func TestCareOrderPayCareSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/inner/order/pay/{id}", careOrderCareId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("qrcode").NotNull()
}

func TestCareOrderCancelNoPayCareSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/care/v1/inner/order/cancel/{id}", careOrderCareId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestCareOrderCancelPayCareSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/care/v1/inner/order/cancel/{id}", CareOrder.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestCareOrderShowReturnSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/inner/order/{id}", CareOrder.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(CareOrder.ID)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForCancel)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")
	obj.Value("data").Object().Value("total").Equal(model.Ftos(CareOrder.Total))
	obj.Value("data").Object().Value("start_at").String().Contains(CareOrder.StartAt.Format("2006-01-02 15:04"))
	obj.Value("data").Object().Value("end_at").String().Contains(CareOrder.EndAt.Format("2006-01-02 15:04"))

	var total decimal.Decimal
	sub := int(CareOrder.EndAt.Sub(CareOrder.StartAt).Hours())
	maxCarePrice := decimal.NewFromFloat(Care.MaxPrice)
	if Care.TimeType == "天" {
		total = maxCarePrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if Care.TimeType == "时" {
		total = maxCarePrice.Mul(decimal.NewFromFloat(float64(sub)))
	}
	f, _ := total.Float64()

	obj.Value("data").Object().Value("total").Equal(model.Ftos(f))
	obj.Value("data").Object().Value("rmk").Equal(CareOrder.Rmk)
	obj.Value("data").Object().Value("pay_type").Equal(CareOrder.PayType)
	obj.Value("data").Object().Value("is_return").Equal(1)

	obj.Value("data").Object().Value("order_carer_info").Null()
	obj.Value("data").Object().Value("comments").Array().Length().Equal(len(CareOrder.CareOrderComments))
	comment := obj.Value("data").Object().Value("comments").Array().First().Object()
	comment.Value("id").NotNull()
	comment.Value("user_id").Equal(User.ID)
	comment.Value("application_id").Equal(model.AppId)
	comment.Value("content").Equal(CareOrder.CareOrderComments[0].Content)
	comment.Value("pics").Array().Length().Equal(2)
	comment.Value("pics").Array().First().Equal("https")
	comment.Value("pics").Array().Last().Equal("https")

	orderInfo := obj.Value("data").Object().Value("order_info").Object()
	orderInfo.Value("id").NotNull()
	orderInfo.Value("name").Equal(CareOrder.CareOrderInfo.Name)
	orderInfo.Value("desc").Equal(CareOrder.CareOrderInfo.Desc)
	orderInfo.Value("application_name").Equal(CareOrder.CareOrderInfo.ApplicationName)
	orderInfo.Value("time_type").Equal(CareOrder.CareOrderInfo.TimeType)
	orderInfo.Value("care_detail").Equal(CareOrder.CareOrderInfo.CareDetail)
	orderInfo.Value("care_tags").Equal(CareOrder.CareOrderInfo.CareTags)
	orderInfo.Value("min_price").Equal(model.Ftos(CareOrder.CareOrderInfo.MinPrice))
	orderInfo.Value("max_price").Equal(model.Ftos(CareOrder.CareOrderInfo.MaxPrice))
	orderInfo.Value("cover").Equal(CareOrder.CareOrderInfo.Cover)
	orderInfo.Value("care_type").Equal(CareOrder.CareOrderInfo.CareType)
	orderInfo.Value("care_order_id").Equal(CareOrder.ID)

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal(CareOrder.CareOrderAddr.Name)
	addr.Value("loc_name").Equal(CareOrder.CareOrderAddr.LocName)
	addr.Value("bed_num").Equal(CareOrder.CareOrderAddr.BedNum)
	addr.Value("hospital_no").Equal(CareOrder.CareOrderAddr.HospitalNo)
	addr.Value("disease").Equal(CareOrder.CareOrderAddr.Disease)
	addr.Value("care_order_id").Equal(CareOrder.CareOrderAddr.CareOrderID)
	addr.Value("sex").Equal(CareOrder.CareOrderAddr.Sex)
	addr.Value("hospital_name").Equal(CareOrder.CareOrderAddr.HospitalName)
	addr.Value("phone").Equal(CareOrder.CareOrderAddr.Phone)
	addr.Value("age").Equal(CareOrder.CareOrderAddr.Age)
	addr.Value("addr").Equal(CareOrder.CareOrderAddr.Addr)

	orderReturn := obj.Value("data").Object().Value("return_order").Object()
	orderReturn.Value("id").NotNull()
	orderReturn.Value("order_no").String().Contains("RC")
	orderReturn.Value("status").Object().Value("value").Equal(model.IReturnOrderStatusForAudit)
	orderReturn.Value("status").Object().Value("text").Equal("待审核")
	orderReturn.Value("total").Equal(model.Ftos(CareOrder.Total))
	orderReturn.Value("open_id").Equal("")
	orderReturn.Value("app_type").Equal(CareOrder.AppType)
	orderReturn.Value("trade_type").Equal("")
	orderReturn.Value("care_order_id").Equal(CareOrder.ID)
	orderReturn.Value("application_id").Equal(CareOrder.ApplicationID)
	orderReturn.Value("pay_type").Equal(CareOrder.PayType)
	orderReturn.Value("user_id").Equal(CareOrder.UserID)
	orderReturn.Value("carer_id").Equal(CareOrder.CarerID)

	returnAddr := orderReturn.Value("addr").Object()
	returnAddr.Value("id").NotNull()
	returnAddr.Value("name").Equal(CareOrder.CareOrderAddr.Name)
	returnAddr.Value("sex").Equal(CareOrder.CareOrderAddr.Sex)
	returnAddr.Value("care_return_order_id").NotNull()
	returnAddr.Value("loc_name").Equal(CareOrder.CareOrderAddr.LocName)
	returnAddr.Value("hospital_no").Equal(CareOrder.CareOrderAddr.HospitalNo)
	returnAddr.Value("hospital_name").Equal(CareOrder.CareOrderAddr.HospitalName)
	returnAddr.Value("age").Equal(CareOrder.CareOrderAddr.Age)
	returnAddr.Value("disease").Equal(CareOrder.CareOrderAddr.Disease)
	returnAddr.Value("phone").Equal(CareOrder.CareOrderAddr.Phone)
	returnAddr.Value("addr").Equal(CareOrder.CareOrderAddr.Addr)

	returnOrderInfo := orderReturn.Value("order_info").Object()
	returnOrderInfo.Value("id").NotNull()
	returnOrderInfo.Value("name").Equal(CareOrder.CareOrderInfo.Name)
	returnOrderInfo.Value("desc").Equal(CareOrder.CareOrderInfo.Desc)
	returnOrderInfo.Value("application_name").Equal(CareOrder.CareOrderInfo.ApplicationName)
	returnOrderInfo.Value("time_type").Equal(CareOrder.CareOrderInfo.TimeType)
	returnOrderInfo.Value("care_detail").Equal(CareOrder.CareOrderInfo.CareDetail)
	returnOrderInfo.Value("care_tags").Equal(CareOrder.CareOrderInfo.CareTags)
	returnOrderInfo.Value("min_price").Equal(model.Ftos(CareOrder.CareOrderInfo.MinPrice))
	returnOrderInfo.Value("max_price").Equal(model.Ftos(CareOrder.CareOrderInfo.MaxPrice))
	returnOrderInfo.Value("cover").Equal(CareOrder.CareOrderInfo.Cover)
	returnOrderInfo.Value("care_type").Equal(CareOrder.CareOrderInfo.CareType)
	returnOrderInfo.Value("care_return_order_id").NotNull()
}

func TestCareOrderDeleteCareSuccess(t *testing.T) {

	obj := model.GetE(t).DELETE("/care/v1/inner/order/{id}", careOrderCareId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestCareOrderDeleteCareRetrunSuccess(t *testing.T) {
	obj := model.GetE(t).DELETE("/care/v1/inner/order/{id}", CareOrder.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	model.CareOrderCount--
}
