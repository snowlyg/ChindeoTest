package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/snowlyg/ChindeoTest/model"

	"net/http"
	"strconv"
	"testing"
	"time"
)

var careOrderCarerId float64
var carerPrice decimal.Decimal
var carerTimeTypeText string

func TestCarrListSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/care/v1/inner/carer").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("carer_tag_id", CarerTag.ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.CarerCount)

	fitst := obj.Value("data").Object().Value("data").Array().First().Object()
	fitst.Value("id").Equal(Carer.ID)
	fitst.Value("name").Equal(Carer.Name)
	fitst.Value("desc").Equal(Carer.Desc)
	fitst.Value("age").Equal(Carer.Age)
	fitst.Value("work_exp").Equal(Carer.WorkExp)
	fitst.Value("phone").Equal(Carer.Phone)
	fitst.Value("sex").Equal(Carer.Sex)
	fitst.Value("time_type").Equal(Carer.TimeType)
	fitst.Value("amount").Equal(model.CarerAmount)
	fitst.Value("avatar").Equal(Carer.Avatar)
}

func TestCarrListNoTagIdSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/inner/carer").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.CarerNoTagCount)

	fitst := obj.Value("data").Object().Value("data").Array().Last().Object()
	fitst.Value("id").Equal(Carer.ID)
	fitst.Value("name").Equal(Carer.Name)
	fitst.Value("desc").Equal(Carer.Desc)
	fitst.Value("age").Equal(Carer.Age)
	fitst.Value("work_exp").Equal(Carer.WorkExp)
	fitst.Value("phone").Equal(Carer.Phone)
	fitst.Value("sex").Equal(Carer.Sex)
	fitst.Value("time_type").Equal(Carer.TimeType)
	fitst.Value("amount").Equal(model.CarerAmount)
	fitst.Value("avatar").Equal(Carer.Avatar)
}

func TestCarrListNoPageSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/inner/carer").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("page_size", "-1").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(model.CarerCount)
	obj.Value("data").Array().Last().Object().Value("id").Equal(Carer.ID)
}

func TestCarrShowSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/inner/carer/{id}", Carer.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(Carer.ID)
	obj.Value("data").Object().Value("name").Equal(Carer.Name)
	obj.Value("data").Object().Value("desc").Equal(Carer.Desc)
	obj.Value("data").Object().Value("age").Equal(Carer.Age)
	obj.Value("data").Object().Value("work_exp").Equal(Carer.WorkExp)
	obj.Value("data").Object().Value("phone").Equal(Carer.Phone)
	obj.Value("data").Object().Value("sex").Equal(Carer.Sex)
	obj.Value("data").Object().Value("time_type").Equal(Carer.TimeType)
	obj.Value("data").Object().Value("amount").Equal(model.CarerAmount)
	obj.Value("data").Object().Value("avatar").Equal(Carer.Avatar)
}

var carerStartAt time.Time
var carerEndAt time.Time

func TestCarrOrderAddCarerSuccess(t *testing.T) {
	carerStartAt = time.Now().AddDate(0, 0, 1)
	carerEndAt = time.Now().AddDate(0, 0, 2)
	careOrder := map[string]interface{}{
		"patient_name": "操蛋",
		"loc_name":     "泥马",
		"bed_num":      "05",
		"hospital_no":  "9556854545",
		"disease":      "玩玩",
		"carer_id":     Carer.ID,
		"start_at":     carerStartAt.Format("2006-01-02 15:04:05"),
		"end_at":       carerEndAt.Format("2006-01-02 15:04:05"),
		"rmk":          "年轻貌美",
		"id_card_no":   "456952158962254456",
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
	obj.Value("data").Object().Value("start_at").Equal(carerStartAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("end_at").Equal(carerEndAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("app_type").Equal(model.OrderAppTypeBed)
	obj.Value("data").Object().Value("application_id").Equal(13)

	id := obj.Value("data").Object().Value("id").Raw()
	data, ok := id.(string)
	if ok {
		careOrderCarerId, _ = strconv.ParseFloat(data, 10)
	}
}

func TestCarrOrderAddCarerError(t *testing.T) {
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
		"carer_id":     Carer.ID,
		"id_card_no":   "456952158962254456",
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

func TestCarrOrderAddCarerErrorTime(t *testing.T) {
	startAt := time.Now().AddDate(0, 0, 1)
	endAt := time.Now().AddDate(0, 0, 2)
	careOrder := map[string]interface{}{
		"patient_name": "泥马",
		"loc_name":     "泥马",
		"bed_num":      "05",
		"hospital_no":  "9556854545",
		"disease":      "玩玩",
		"start_at":     startAt.Format("2006-01-02 15:04:05"),
		"end_at":       endAt.Format("2006-01-02 15:04:05"),
		"rmk":          "年轻貌美",
		"carer_id":     Carer.ID,
		"id_card_no":   "456952158962254456",
	}

	obj := model.GetE(t).POST("/care/v1/inner/order/add").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(careOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Contains("时间段已经被预约")
}

func TestCarrOrderCommentSuccess(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "content",
		"id_card_no": "456952158962254456",
		"pics":       model.Pics,
		"order_id":   careOrderCarerId,
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

func TestCarrOrderCommentNoContentError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "",
		"id_card_no": "456952158962254456",
		"pics":       model.Pics,
		"order_id":   careOrderCarerId,
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

func TestCarrOrderCommentNoIdCardNoError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "456952158962254456",
		"id_card_no": "",
		"pics":       model.Pics,
		"order_id":   careOrderCarerId,
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

func TestCarrOrderCommentNoOrderIdError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "456952158962254456",
		"id_card_no": "456952158962254456",
		"pics":       model.Pics,
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

func TestCarrOrderCommentOrderNotExistsError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "456952158962254456",
		"id_card_no": "456952158962254456",
		"pics":       model.Pics,
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

func TestCarrShowAfterOrderAddSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/inner/carer/{id}", Carer.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(Carer.ID)
	obj.Value("data").Object().Value("name").Equal(Carer.Name)
	obj.Value("data").Object().Value("desc").Equal(Carer.Desc)
	obj.Value("data").Object().Value("age").Equal(Carer.Age)
	obj.Value("data").Object().Value("work_exp").Equal(Carer.WorkExp)
	obj.Value("data").Object().Value("phone").Equal(Carer.Phone)
	obj.Value("data").Object().Value("sex").Equal(Carer.Sex)
	obj.Value("data").Object().Value("time_type").Equal(Carer.TimeType)
	obj.Value("data").Object().Value("amount").Equal(model.CarerAmount + 1)
	model.CarerAmount++
	obj.Value("data").Object().Value("avatar").Equal(Carer.Avatar)

	carerPriceF, _ := strconv.ParseFloat(model.GetS(obj.Value("data").Object().Value("price").Raw()), 10)
	carerPrice = decimal.NewFromFloat(carerPriceF)
	carerTimeTypeText = model.GetS(obj.Value("data").Object().Value("time_type").Raw())
}

func TestCarrOrderShowCarerSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/inner/order/{id}", careOrderCarerId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(careOrderCarerId)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForPay)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("待付款")
	obj.Value("data").Object().Value("start_at").Equal(carerStartAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("end_at").Equal(carerEndAt.Format("2006-01-02 15:04:05"))
	var total decimal.Decimal
	sub := int(carerEndAt.Sub(carerStartAt).Hours())
	if carerTimeTypeText == "天" {
		total = carerPrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if carerTimeTypeText == "时" {
		total = carerPrice.Mul(decimal.NewFromFloat(float64(sub)))
	}
	f, _ := total.Float64()
	obj.Value("data").Object().Value("total").Equal(fmt.Sprintf("%.2f", f))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("pay_type").Equal(1)
	obj.Value("data").Object().Value("is_return").Equal(0)

	obj.Value("data").Object().Value("order_info").Null()
	obj.Value("data").Object().Value("return_order").Null()
	obj.Value("data").Object().Value("comments").Array().Length().Equal(2)

	orderCarerInfo := obj.Value("data").Object().Value("order_carer_info").Object()
	orderCarerInfo.Value("id").NotNull()
	orderCarerInfo.Value("name").Equal(Carer.Name)
	orderCarerInfo.Value("desc").Equal(Carer.Desc)
	orderCarerInfo.Value("age").Equal(Carer.Age)
	orderCarerInfo.Value("work_exp").Equal(Carer.WorkExp)
	orderCarerInfo.Value("phone").Equal(Carer.Phone)
	orderCarerInfo.Value("sex").Equal(Carer.Sex)
	orderCarerInfo.Value("time_type").Equal(Carer.TimeType)
	orderCarerInfo.Value("avatar").Equal(Carer.Avatar)
	orderCarerInfo.Value("carer_detail").Equal(Carer.CarerDetail)

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal("操蛋")
	addr.Value("loc_name").Equal("泥马")
	addr.Value("bed_num").Equal("05")
	addr.Value("hospital_no").Equal("9556854545")
	addr.Value("disease").Equal("玩玩")
	addr.Value("care_order_id").Equal(careOrderCarerId)
	addr.Value("sex").Equal(1)
	addr.Value("hospital_name").Equal("我的医院")
	addr.Value("phone").Equal("")
	addr.Value("age").Equal(0)
	addr.Value("addr").Equal("")
}

func TestCarrOrderPayCareSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/care/v1/inner/order/pay/{id}", careOrderCarerId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("qrcode").NotNull()
}

func TestCarrOrderCancelNoPayCarerSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/care/v1/inner/order/cancel/{id}", careOrderCarerId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestCarrOrderCancelPayCarerSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/care/v1/inner/order/cancel/{id}", CarerOrder.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestCarrOrderShowReturnCarerSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/inner/order/{id}", CarerOrder.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(CarerOrder.ID)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForCancel)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")
	obj.Value("data").Object().Value("start_at").String().Contains(CarerOrder.StartAt.Format("2006-01-02 15:04"))
	obj.Value("data").Object().Value("end_at").String().Contains(CarerOrder.EndAt.Format("2006-01-02 15:04"))
	var total decimal.Decimal
	sub := int(carerEndAt.Sub(carerStartAt).Hours())
	if carerTimeTypeText == "天" {
		total = carerPrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if carerTimeTypeText == "时" {
		total = carerPrice.Mul(decimal.NewFromFloat(float64(sub)))
	}
	f, _ := total.Float64()
	obj.Value("data").Object().Value("total").Equal(fmt.Sprintf("%.2f", f))
	obj.Value("data").Object().Value("rmk").Equal(CarerOrder.Rmk)
	obj.Value("data").Object().Value("pay_type").Equal(CarerOrder.PayType)
	obj.Value("data").Object().Value("is_return").Equal(1)

	obj.Value("data").Object().Value("order_info").Null()
	obj.Value("data").Object().Value("comments").Array().Length().Equal(len(CarerOrder.CareOrderComments))
	comment := obj.Value("data").Object().Value("comments").Array().First().Object()
	comment.Value("id").NotNull()
	comment.Value("user_id").Equal(User.ID)
	comment.Value("application_id").Equal(model.AppId)
	comment.Value("content").Equal(CarerOrder.CareOrderComments[0].Content)
	comment.Value("pics").Array().Length().Equal(2)
	comment.Value("pics").Array().First().Equal("https")
	comment.Value("pics").Array().Last().Equal("https")

	orderCarerInfo := obj.Value("data").Object().Value("order_carer_info").Object()
	orderCarerInfo.Value("id").NotNull()
	orderCarerInfo.Value("name").Equal(Carer.Name)
	orderCarerInfo.Value("desc").Equal(Carer.Desc)
	orderCarerInfo.Value("age").Equal(Carer.Age)
	orderCarerInfo.Value("work_exp").Equal(Carer.WorkExp)
	orderCarerInfo.Value("phone").Equal(Carer.Phone)
	orderCarerInfo.Value("sex").Equal(Carer.Sex)
	orderCarerInfo.Value("time_type").Equal(Carer.TimeType)
	orderCarerInfo.Value("avatar").Equal(Carer.Avatar)
	orderCarerInfo.Value("carer_detail").Equal(Carer.CarerDetail)

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal(CarerOrder.CareOrderAddr.Name)
	addr.Value("loc_name").Equal(CarerOrder.CareOrderAddr.LocName)
	addr.Value("bed_num").Equal(CarerOrder.CareOrderAddr.BedNum)
	addr.Value("hospital_no").Equal(CarerOrder.CareOrderAddr.HospitalNo)
	addr.Value("disease").Equal(CarerOrder.CareOrderAddr.Disease)
	addr.Value("care_order_id").Equal(CarerOrder.CareOrderAddr.CareOrderID)
	addr.Value("sex").Equal(CarerOrder.CareOrderAddr.Sex)
	addr.Value("hospital_name").Equal(CarerOrder.CareOrderAddr.HospitalName)
	addr.Value("phone").Equal(CarerOrder.CareOrderAddr.Phone)
	addr.Value("age").Equal(CarerOrder.CareOrderAddr.Age)
	addr.Value("addr").Equal(CarerOrder.CareOrderAddr.Addr)

	orderReturn := obj.Value("data").Object().Value("return_order").Object()
	orderReturn.Value("id").NotNull()
	orderReturn.Value("order_no").String().Contains("RC")
	orderReturn.Value("status").Object().Value("value").Equal(model.IReturnOrderStatusForAudit)
	orderReturn.Value("status").Object().Value("text").Equal("待审核")
	orderReturn.Value("total").Equal(model.Ftos(CarerOrder.Total))
	orderReturn.Value("open_id").Equal("")
	orderReturn.Value("app_type").Equal(CarerOrder.AppType)
	orderReturn.Value("trade_type").Equal("")
	orderReturn.Value("care_order_id").Equal(CarerOrder.ID)
	orderReturn.Value("application_id").Equal(CarerOrder.ApplicationID)
	orderReturn.Value("pay_type").Equal(CarerOrder.PayType)
	orderReturn.Value("user_id").Equal(CarerOrder.UserID)
	orderReturn.Value("carer_id").Equal(CarerOrder.CarerID)

	returnAddr := orderReturn.Value("addr").Object()
	returnAddr.Value("id").NotNull()
	returnAddr.Value("name").Equal(CarerOrder.CareOrderAddr.Name)
	returnAddr.Value("sex").Equal(CarerOrder.CareOrderAddr.Sex)
	returnAddr.Value("care_return_order_id").NotNull()
	returnAddr.Value("loc_name").Equal(CarerOrder.CareOrderAddr.LocName)
	returnAddr.Value("hospital_no").Equal(CarerOrder.CareOrderAddr.HospitalNo)
	returnAddr.Value("hospital_name").Equal(CarerOrder.CareOrderAddr.HospitalName)
	returnAddr.Value("age").Equal(CarerOrder.CareOrderAddr.Age)
	returnAddr.Value("disease").Equal(CarerOrder.CareOrderAddr.Disease)
	returnAddr.Value("phone").Equal(CarerOrder.CareOrderAddr.Phone)
	returnAddr.Value("addr").Equal(CarerOrder.CareOrderAddr.Addr)

	returnOrderInfo := orderReturn.Value("order_carer_info").Object()
	returnOrderInfo.Value("id").NotNull()
	returnOrderInfo.Value("name").Equal(CarerOrder.CareOrderCarerInfo.Name)
	returnOrderInfo.Value("desc").Equal(CarerOrder.CareOrderCarerInfo.Desc)
	returnOrderInfo.Value("age").Equal(CarerOrder.CareOrderCarerInfo.Age)
	returnOrderInfo.Value("work_exp").Equal(CarerOrder.CareOrderCarerInfo.WorkExp)
	returnOrderInfo.Value("phone").Equal(CarerOrder.CareOrderCarerInfo.Phone)
	returnOrderInfo.Value("sex").Equal(CarerOrder.CareOrderCarerInfo.Sex)
	returnOrderInfo.Value("time_type").Equal(CarerOrder.CareOrderCarerInfo.TimeType)
	returnOrderInfo.Value("avatar").Equal(CarerOrder.CareOrderCarerInfo.Avatar)
	returnOrderInfo.Value("carer_detail").Equal(CarerOrder.CareOrderCarerInfo.CarerDetail)
	returnOrderInfo.Value("care_return_order_id").NotNull()
}

func TestCarrOrderDeleteCareSuccess(t *testing.T) {

	obj := model.GetE(t).DELETE("/care/v1/inner/order/{id}", careOrderCarerId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestCarrOrderDeleteCareRetrunSuccess(t *testing.T) {
	obj := model.GetE(t).DELETE("/care/v1/inner/order/{id}", CarerOrder.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	model.CareOrderCount--
}
