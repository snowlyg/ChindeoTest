package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/snowlyg/ChindeoTest/common"
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
)

var miniCareOrderCarerId float64
var miniCarerPrice decimal.Decimal
var miniCarerTimeTypeText string

func TestMiniWechatCarrListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/carer").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithQuery("application_id", AppId).
		WithQuery("carer_tag_id", CarerTag.ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(1)

	fitst := obj.Value("data").Object().Value("data").Array().First().Object()
	fitst.Value("id").Equal(Carer.ID)
	fitst.Value("name").Equal(Carer.Name)
	fitst.Value("desc").Equal(Carer.Desc)
	fitst.Value("age").Equal(Carer.Age)
	fitst.Value("work_exp").Equal(Carer.WorkExp)
	fitst.Value("phone").Equal(Carer.Phone)
	fitst.Value("sex").Equal(Carer.Sex)
	fitst.Value("time_type").Equal(Carer.TimeType)
	fitst.Value("amount").Equal(CarerAmount)
	fitst.Value("avatar").Equal(Carer.Avatar)
}

func TestMiniWechatCarrNoTagIdListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/carer").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithQuery("application_id", AppId).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(2)

	fitst := obj.Value("data").Object().Value("data").Array().Last().Object()
	fitst.Value("id").Equal(Carer.ID)
	fitst.Value("name").Equal(Carer.Name)
	fitst.Value("desc").Equal(Carer.Desc)
	fitst.Value("age").Equal(Carer.Age)
	fitst.Value("work_exp").Equal(Carer.WorkExp)
	fitst.Value("phone").Equal(Carer.Phone)
	fitst.Value("sex").Equal(Carer.Sex)
	fitst.Value("time_type").Equal(Carer.TimeType)
	fitst.Value("amount").Equal(CarerAmount)
	fitst.Value("avatar").Equal(Carer.Avatar)
}

func TestMiniWechatCarrListNoPageSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/carer").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithQuery("application_id", AppId).
		WithQuery("page_size", "-1").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Last().Object().Value("id").Equal(Carer.ID)
}

func TestMiniWechatCarrShowSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/carer/{id}", Carer.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
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
	obj.Value("data").Object().Value("amount").Equal(CarerAmount)
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
		"application_id": AppId,
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.POST("/care/v1/order/add").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
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
	obj.Value("data").Object().Value("app_type").Equal(common.ORDER_APP_TYPE_MINI)
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

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.POST("/care/v1/order/add").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
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
		"application_id": AppId,
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.POST("/care/v1/order/add").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithJSON(careOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Contains("时间段已经被预约")
}

func TestMiniWechatCarrShowAfterOrderAddSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/carer/{id}", Carer.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
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
	obj.Value("data").Object().Value("amount").Equal(CarerAmount + 1)
	CarerAmount++
	obj.Value("data").Object().Value("avatar").Equal(Carer.Avatar)

	carerPriceF, _ := strconv.ParseFloat(common.GetS(obj.Value("data").Object().Value("price").Raw()), 10)
	miniCarerPrice = decimal.NewFromFloat(carerPriceF)
	miniCarerTimeTypeText = common.GetS(obj.Value("data").Object().Value("time_type").Raw())
}

func TestMiniWechatCarrOrderShowCarerSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/order/{id}", miniCareOrderCarerId).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(miniCareOrderCarerId)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(common.I_ORDER_STATUS_FOR_PAY)
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
	obj.Value("data").Object().Value("comments").Array().Length().Equal(0)

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
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/order/pay/{id}", miniCareOrderCarerId).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Contains("PARAM_ERROR:appid和openid不匹配")
}

func TestMiniWechatCarrOrderCancelNoPayCarerSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(),
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/order/cancel/{id}", miniCareOrderCarerId).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatCarrOrderCancelPayCarerSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(),
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/order/cancel/{id}", MiniCarerOrder.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatCarrOrderShowReturnCarerSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/order/{id}", MiniCarerOrder.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(MiniCarerOrder.ID)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(common.I_ORDER_STATUS_FOR_CANCEL)
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
	obj.Value("data").Object().Value("comments").Array().Length().Equal(1)
	comment := obj.Value("data").Object().Value("comments").Array().First().Object()
	comment.Value("id").NotNull()
	comment.Value("user_id").Equal(User.ID)
	comment.Value("application_id").Equal(AppId)
	comment.Value("star").Equal(5)
	comment.Value("content").Equal(MiniCarerOrderComment.Content)
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
	addr.Value("name").Equal(MiniCarerOrderAddr.Name)
	addr.Value("loc_name").Equal(MiniCarerOrderAddr.LocName)
	addr.Value("bed_num").Equal(MiniCarerOrderAddr.BedNum)
	addr.Value("hospital_no").Equal(MiniCarerOrderAddr.HospitalNo)
	addr.Value("disease").Equal(MiniCarerOrderAddr.Disease)
	addr.Value("care_order_id").Equal(MiniCarerOrderAddr.CareOrderID)
	addr.Value("sex").Equal(MiniCarerOrderAddr.Sex)
	addr.Value("hospital_name").Equal(MiniCarerOrderAddr.HospitalName)
	addr.Value("phone").Equal(MiniCarerOrderAddr.Phone)
	addr.Value("age").Equal(MiniCarerOrderAddr.Age)
	addr.Value("addr").Equal(MiniCarerOrderAddr.Addr)

	orderReturn := obj.Value("data").Object().Value("return_order").Object()
	orderReturn.Value("id").NotNull()
	orderReturn.Value("order_no").String().Contains("RC")
	orderReturn.Value("status").Object().Value("value").Equal(common.I_RETURN_ORDER_STATUS_FOR_AUDIT)
	orderReturn.Value("status").Object().Value("text").Equal("待审核")
	orderReturn.Value("total").Equal(common.Ftos(MiniCarerOrder.Total))
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
	returnAddr.Value("name").Equal(MiniCarerOrderAddr.Name)
	returnAddr.Value("sex").Equal(MiniCarerOrderAddr.Sex)
	returnAddr.Value("care_return_order_id").NotNull()
	returnAddr.Value("loc_name").Equal(MiniCarerOrderAddr.LocName)
	returnAddr.Value("hospital_no").Equal(MiniCarerOrderAddr.HospitalNo)
	returnAddr.Value("hospital_name").Equal(MiniCarerOrderAddr.HospitalName)
	returnAddr.Value("age").Equal(MiniCarerOrderAddr.Age)
	returnAddr.Value("disease").Equal(MiniCarerOrderAddr.Disease)
	returnAddr.Value("phone").Equal(MiniCarerOrderAddr.Phone)
	returnAddr.Value("addr").Equal(MiniCarerOrderAddr.Addr)

	returnOrderInfo := orderReturn.Value("order_carer_info").Object()
	returnOrderInfo.Value("id").NotNull()
	returnOrderInfo.Value("name").Equal(MiniCarerOrderCarerInfo.Name)
	returnOrderInfo.Value("desc").Equal(MiniCarerOrderCarerInfo.Desc)
	returnOrderInfo.Value("age").Equal(MiniCarerOrderCarerInfo.Age)
	returnOrderInfo.Value("work_exp").Equal(MiniCarerOrderCarerInfo.WorkExp)
	returnOrderInfo.Value("phone").Equal(MiniCarerOrderCarerInfo.Phone)
	returnOrderInfo.Value("sex").Equal(MiniCarerOrderCarerInfo.Sex)
	returnOrderInfo.Value("time_type").Equal(MiniCarerOrderCarerInfo.TimeType)
	returnOrderInfo.Value("avatar").Equal(MiniCarerOrderCarerInfo.Avatar)
	returnOrderInfo.Value("carer_detail").Equal(MiniCarerOrderCarerInfo.CarerDetail)
	returnOrderInfo.Value("care_return_order_id").NotNull()
}

func TestMiniWechatCarrOrderDeleteCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.DELETE("/care/v1/order/{id}", miniCareOrderCarerId).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatCarrOrderDeleteCareRetrunSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.DELETE("/care/v1/order/{id}", MiniCarerOrder.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	CareOrderCount--
}
