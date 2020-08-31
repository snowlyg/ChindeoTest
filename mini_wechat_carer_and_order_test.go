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

var miniCareOrderCarerId float64
var miniCarerPrice decimal.Decimal
var miniCarerTimeTypeText string

func TestMiniWechatCarrListSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/carer").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("application_id", model.AppId).
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

func TestMiniWechatCarrNoTagIdListSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/carer").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("application_id", model.AppId).
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

func TestMiniWechatCarrListNoPageSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/carer").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("application_id", model.AppId).
		WithQuery("page_size", "-1").
		WithQuery("carer_tag_id", CarerTag.ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(model.CarerCount)
	obj.Value("data").Array().Last().Object().Value("id").Equal(Carer.ID)
}

func TestMiniWechatCarrShowSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/care/v1/carer/{id}", Carer.ID).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
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

var miniCarerStartAt time.Time
var miniCarerEndAt time.Time

func TestMiniWechatCarrOrderAddCarerSuccess(t *testing.T) {
	miniCarerStartAt = time.Now().AddDate(0, 0, 1)
	miniCarerEndAt = time.Now().AddDate(0, 0, 2)
	careOrder := map[string]interface{}{
		"carer_id":       Carer.ID,
		"start_at":       miniCarerStartAt.Format("2006-01-02 15:04:05"),
		"end_at":         miniCarerEndAt.Format("2006-01-02 15:04:05"),
		"rmk":            "年轻貌美",
		"id_card_no":     "456952158962254456",
		"addr_id":        Addr.ID,
		"application_id": model.AppId,
	}

	obj := model.GetE(t).POST("/care/v1/order/add").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(careOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").NotEqual(0)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("start_at").Equal(miniCarerStartAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("end_at").Equal(miniCarerEndAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("app_type").Equal(model.OrderAppTypeMini)
	obj.Value("data").Object().Value("application_id").Equal(13)

	id := obj.Value("data").Object().Value("id").Raw()
	data, ok := id.(string)
	if ok {
		miniCareOrderCarerId, _ = strconv.ParseFloat(data, 10)
	}
}

func TestMiniWechatCarrOrderAddCarerError(t *testing.T) {
	startAt := time.Now().AddDate(0, 0, 5)
	endAt := time.Now().AddDate(0, 0, 6)
	careOrder := map[string]interface{}{
		"start_at": startAt.Format("2006-01-02 15:04:05"),
		"end_at":   endAt.Format("2006-01-02 15:04:05"),
		"rmk":      "年轻貌美",
		"carer_id": Carer.ID,
	}

	obj := model.GetE(t).POST("/care/v1/order/add").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(careOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("医院不能为空！")
}

func TestMiniWechatCarrOrderAddCarerErrorTime(t *testing.T) {
	startAt := time.Now().AddDate(0, 0, 1)
	endAt := time.Now().AddDate(0, 0, 2)
	careOrder := map[string]interface{}{
		"start_at":       startAt.Format("2006-01-02 15:04:05"),
		"end_at":         endAt.Format("2006-01-02 15:04:05"),
		"rmk":            "年轻貌美",
		"carer_id":       Carer.ID,
		"addr_id":        Addr.ID,
		"application_id": model.AppId,
	}

	obj := model.GetE(t).POST("/care/v1/order/add").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(careOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Contains("时间段已经被预约")
}

func TestMiniWechatCarrOrderCommentSuccess(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "content",
		"id_card_no": "456952158962254456",
		"pics":       model.GetPics(),
		"order_id":   miniCareOrderCarerId,
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

func TestMiniWechatCarrOrderCommentOrderNotExistsError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "fdssdfs",
		"id_card_no": "456952158962254456",
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

func TestMiniWechatCarrOrderCommentNoContentError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "",
		"id_card_no": "456952158962254456",
		"pics":       model.GetPics(),
		"order_id":   miniCareOrderCarerId,
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

func TestMiniWechatCarrOrderCommentNoIdCardNoError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "456952158962254456",
		"id_card_no": "",
		"pics":       model.GetPics(),
		"order_id":   miniCareOrderCarerId,
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

func TestMiniWechatCarrOrderCommentNoOrderIdError(t *testing.T) {
	comment := map[string]interface{}{
		"star":       1,
		"content":    "456952158962254456",
		"id_card_no": "456952158962254456",
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

func TestMiniWechatCarrShowAfterOrderAddSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/carer/{id}", Carer.ID).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
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
	miniCarerPrice = decimal.NewFromFloat(carerPriceF)
	miniCarerTimeTypeText = model.GetS(obj.Value("data").Object().Value("time_type").Raw())
}

func TestMiniWechatCarrOrderShowCarerSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/order/{id}", miniCareOrderCarerId).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(miniCareOrderCarerId)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForPay)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("待付款")
	obj.Value("data").Object().Value("start_at").Equal(miniCarerStartAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("end_at").Equal(miniCarerEndAt.Format("2006-01-02 15:04:05"))
	var total decimal.Decimal
	sub := int(miniCarerEndAt.Sub(miniCarerStartAt).Hours())
	if miniCarerTimeTypeText == "天" {
		total = miniCarerPrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if miniCarerTimeTypeText == "时" {
		total = miniCarerPrice.Mul(decimal.NewFromFloat(float64(sub)))
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
	addr.Value("name").Equal(Addr.Name)
	addr.Value("sex").Equal(Addr.Sex)
	addr.Value("care_order_id").Equal(miniCareOrderCarerId)
	addr.Value("loc_name").Equal(Addr.LocName)
	addr.Value("hospital_no").Equal(Addr.HospitalNo)
	addr.Value("hospital_name").Equal(Addr.HospitalName)
	addr.Value("age").Equal(Addr.Age)
	addr.Value("disease").Equal(Addr.Disease)
	addr.Value("phone").Equal(Addr.Phone)
	addr.Value("addr").Equal(Addr.Addr)
}

func TestMiniWechatCarrOrderPayCareSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/order/pay/{id}", miniCareOrderCarerId).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Contains("PARAM_ERROR:appid和openid不匹配")
}

func TestMiniWechatCarrOrderCancelNoPayCarerSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/care/v1/order/cancel/{id}", miniCareOrderCarerId).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatCarrOrderCancelPayCarerSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/care/v1/order/cancel/{id}", MiniCarerOrder.ID).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatCarrOrderShowReturnCarerSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/care/v1/order/{id}", MiniCarerOrder.ID).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(MiniCarerOrder.ID)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForCancel)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")
	obj.Value("data").Object().Value("start_at").String().Contains(MiniCarerOrder.StartAt.Format("2006-01-02 15:04"))
	obj.Value("data").Object().Value("end_at").String().Contains(MiniCarerOrder.EndAt.Format("2006-01-02 15:04"))
	var total decimal.Decimal
	sub := int(miniCarerEndAt.Sub(miniCarerStartAt).Hours())
	if miniCarerTimeTypeText == "天" {
		total = miniCarerPrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if miniCarerTimeTypeText == "时" {
		total = miniCarerPrice.Mul(decimal.NewFromFloat(float64(sub)))
	}
	f, _ := total.Float64()
	obj.Value("data").Object().Value("total").Equal(fmt.Sprintf("%.2f", f))
	obj.Value("data").Object().Value("rmk").Equal(MiniCarerOrder.Rmk)
	obj.Value("data").Object().Value("pay_type").Equal(MiniCarerOrder.PayType)
	obj.Value("data").Object().Value("is_return").Equal(1)
	obj.Value("data").Object().Value("order_info").Null()
	obj.Value("data").Object().Value("comments").Array().Length().Equal(len(MiniCarerOrder.CareOrderComments))
	comment := obj.Value("data").Object().Value("comments").Array().First().Object()
	comment.Value("id").NotNull()
	comment.Value("user_id").Equal(User.ID)
	comment.Value("application_id").Equal(model.AppId)
	comment.Value("star").Equal(MiniCarerOrder.CareOrderComments[0].Star)
	comment.Value("content").Equal(MiniCarerOrder.CareOrderComments[0].Content)
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
	addr.Value("name").Equal(MiniCarerOrder.CareOrderAddr.Name)
	addr.Value("loc_name").Equal(MiniCarerOrder.CareOrderAddr.LocName)
	addr.Value("bed_num").Equal(MiniCarerOrder.CareOrderAddr.BedNum)
	addr.Value("hospital_no").Equal(MiniCarerOrder.CareOrderAddr.HospitalNo)
	addr.Value("disease").Equal(MiniCarerOrder.CareOrderAddr.Disease)
	addr.Value("care_order_id").Equal(MiniCarerOrder.CareOrderAddr.CareOrderID)
	addr.Value("sex").Equal(MiniCarerOrder.CareOrderAddr.Sex)
	addr.Value("hospital_name").Equal(MiniCarerOrder.CareOrderAddr.HospitalName)
	addr.Value("phone").Equal(MiniCarerOrder.CareOrderAddr.Phone)
	addr.Value("age").Equal(MiniCarerOrder.CareOrderAddr.Age)
	addr.Value("addr").Equal(MiniCarerOrder.CareOrderAddr.Addr)

	orderReturn := obj.Value("data").Object().Value("return_order").Object()
	orderReturn.Value("id").NotNull()
	orderReturn.Value("order_no").String().Contains("RC")
	orderReturn.Value("status").Object().Value("value").Equal(model.IReturnOrderStatusForAudit)
	orderReturn.Value("status").Object().Value("text").Equal("待审核")
	orderReturn.Value("total").Equal(model.Ftos(MiniCarerOrder.Total))
	orderReturn.Value("open_id").Equal("")
	orderReturn.Value("app_type").Equal(MiniCarerOrder.AppType)
	orderReturn.Value("trade_type").Equal("")
	orderReturn.Value("care_order_id").Equal(MiniCarerOrder.ID)
	orderReturn.Value("application_id").Equal(MiniCarerOrder.ApplicationID)
	orderReturn.Value("pay_type").Equal(MiniCarerOrder.PayType)
	orderReturn.Value("user_id").Equal(MiniCarerOrder.UserID)
	orderReturn.Value("carer_id").Equal(MiniCarerOrder.CarerID)

	returnAddr := orderReturn.Value("addr").Object()
	returnAddr.Value("id").NotNull()
	returnAddr.Value("name").Equal(MiniCarerOrder.CareOrderAddr.Name)
	returnAddr.Value("sex").Equal(MiniCarerOrder.CareOrderAddr.Sex)
	returnAddr.Value("care_return_order_id").NotNull()
	returnAddr.Value("loc_name").Equal(MiniCarerOrder.CareOrderAddr.LocName)
	returnAddr.Value("hospital_no").Equal(MiniCarerOrder.CareOrderAddr.HospitalNo)
	returnAddr.Value("hospital_name").Equal(MiniCarerOrder.CareOrderAddr.HospitalName)
	returnAddr.Value("age").Equal(MiniCarerOrder.CareOrderAddr.Age)
	returnAddr.Value("disease").Equal(MiniCarerOrder.CareOrderAddr.Disease)
	returnAddr.Value("phone").Equal(MiniCarerOrder.CareOrderAddr.Phone)
	returnAddr.Value("addr").Equal(MiniCarerOrder.CareOrderAddr.Addr)

	returnOrderInfo := orderReturn.Value("order_carer_info").Object()
	returnOrderInfo.Value("id").NotNull()
	returnOrderInfo.Value("name").Equal(MiniCarerOrder.CareOrderCarerInfo.Name)
	returnOrderInfo.Value("desc").Equal(MiniCarerOrder.CareOrderCarerInfo.Desc)
	returnOrderInfo.Value("age").Equal(MiniCarerOrder.CareOrderCarerInfo.Age)
	returnOrderInfo.Value("work_exp").Equal(MiniCarerOrder.CareOrderCarerInfo.WorkExp)
	returnOrderInfo.Value("phone").Equal(MiniCarerOrder.CareOrderCarerInfo.Phone)
	returnOrderInfo.Value("sex").Equal(MiniCarerOrder.CareOrderCarerInfo.Sex)
	returnOrderInfo.Value("time_type").Equal(MiniCarerOrder.CareOrderCarerInfo.TimeType)
	returnOrderInfo.Value("avatar").Equal(MiniCarerOrder.CareOrderCarerInfo.Avatar)
	returnOrderInfo.Value("carer_detail").Equal(MiniCarerOrder.CareOrderCarerInfo.CarerDetail)
	returnOrderInfo.Value("care_return_order_id").NotNull()
}

func TestMiniWechatCarrOrderDeleteCareSuccess(t *testing.T) {

	obj := model.GetE(t).DELETE("/care/v1/order/{id}", miniCareOrderCarerId).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatCarrOrderDeleteCareRetrunSuccess(t *testing.T) {

	obj := model.GetE(t).DELETE("/care/v1/order/{id}", MiniCarerOrder.ID).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	model.CareOrderCount--
}
