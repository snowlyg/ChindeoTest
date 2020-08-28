package main

import (
	"fmt"
	"github.com/snowlyg/ChindeoTest/common"
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/shopspring/decimal"
)

var careOrderCareId float64
var carePrice decimal.Decimal
var careTimeTypeText string

func TestCareNoTagIdListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/inner/care").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(2)

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
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/inner/care").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithQuery("care_type_id", CareType.ID).
		WithQuery("care_tag_id", CareTag.ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(1)

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
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/inner/care").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithQuery("care_type_id", CareType.ID).
		WithQuery("care_tag_id", CareTag.ID).
		WithQuery("page_size", "-1").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(1)

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

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/inner/care/{id}", Care.ID).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(Care.ID)
	obj.Value("data").Object().Value("name").Equal(Care.Name)
	obj.Value("data").Object().Value("desc").Equal(Care.Desc)
	obj.Value("data").Object().Value("time_type").Equal(Care.TimeType)
	obj.Value("data").Object().Value("amount").Equal(CareAmount)
	obj.Value("data").Object().Value("cover").Equal(Care.Cover)
	obj.Value("data").Object().Value("care_type_id").Equal(Care.CareTypeID)

	careType := obj.Value("data").Object().Value("type").Object()
	careType.Value("id").Equal(CareType.ID)
	careType.Value("name").Equal(CareType.Name)
}

func TestCareOrderListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/inner/order").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(CareOrderCount)
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
		"id_card_no":   "456952158962254456",
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.POST("/care/v1/inner/order/add").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
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
	obj.Value("data").Object().Value("app_type").Equal(common.ORDER_APP_TYPE_BED)
	obj.Value("data").Object().Value("application_id").Equal(13)
	careOrderCareId, _ = strconv.ParseFloat(common.GetS(obj.Value("data").Object().Value("id").Raw()), 10)
}

func TestCareShowAfterOrderAddSuccess(t *testing.T) {

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/inner/care/{id}", Care.ID).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(Care.ID)
	obj.Value("data").Object().Value("name").Equal(Care.Name)
	obj.Value("data").Object().Value("desc").Equal(Care.Desc)
	obj.Value("data").Object().Value("time_type").Equal(Care.TimeType)
	obj.Value("data").Object().Value("amount").Equal(CareAmount + 1)
	CareAmount++
	obj.Value("data").Object().Value("cover").Equal(Care.Cover)
	obj.Value("data").Object().Value("care_type_id").Equal(Care.CareTypeID)

	careType := obj.Value("data").Object().Value("type").Object()
	careType.Value("id").Equal(CareType.ID)
	careType.Value("name").Equal(CareType.Name)

	carePriceF, _ := strconv.ParseFloat(common.GetS(obj.Value("data").Object().Value("max_price").Raw()), 10)
	carePrice = decimal.NewFromFloat(carePriceF)
	careTimeTypeText = common.GetS(obj.Value("data").Object().Value("time_type").Raw())

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
		"id_card_no":   "456952158962254456",
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.POST("/care/v1/inner/order/add").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
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
		"id_card_no": "456952158962254456",
		"pics":       Pics,
		"order_id":   careOrderCareId,
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.POST("/common/v1/inner/comment/care").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
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
		"id_card_no": "456952158962254456",
		"pics":       Pics,
		"order_id":   careOrderCareId,
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.POST("/common/v1/inner/comment/care").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
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
		"id_card_no": "456952158962254456",
		"pics":       Pics,
		"order_id":   9999,
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.POST("/common/v1/inner/comment/care").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
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
		"pics":       Pics,
		"order_id":   careOrderCareId,
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.POST("/common/v1/inner/comment/care").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
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
		"id_card_no": "456952158962254456",
		"pics":       Pics,
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.POST("/common/v1/inner/comment/care").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("评论订单不能为空！")
}

func TestCareOrderShowCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/inner/order/{id}", careOrderCareId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(careOrderCareId)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(common.I_ORDER_STATUS_FOR_PAY)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("待付款")
	obj.Value("data").Object().Value("start_at").Equal(careStartAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("end_at").Equal(careEndAt.Format("2006-01-02 15:04:05"))

	var total decimal.Decimal
	sub := int(careEndAt.Sub(careStartAt).Hours())
	if careTimeTypeText == "天" {
		total = carePrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if careTimeTypeText == "时" {
		total = carePrice.Mul(decimal.NewFromFloat(float64(sub)))
	}

	f, _ := total.Float64()

	obj.Value("data").Object().Value("total").Equal(common.Ftos(f))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("pay_type").Equal(1)
	obj.Value("data").Object().Value("is_return").Equal(0)

	obj.Value("data").Object().Value("order_carer_info").Null()
	obj.Value("data").Object().Value("return_order").Null()
	obj.Value("data").Object().Value("comments").Array().Length().Equal(2)
	orderInfo := obj.Value("data").Object().Value("order_info").Object()
	orderInfo.Value("id").NotNull()
	orderInfo.Value("name").Equal(CareOrderInfo.Name)
	orderInfo.Value("desc").Equal(CareOrderInfo.Desc)
	orderInfo.Value("application_name").Equal("我的医院")
	orderInfo.Value("time_type").Equal(CareOrderInfo.TimeType)
	orderInfo.Value("care_detail").Equal(CareOrderInfo.CareDetail)
	orderInfo.Value("care_tags").Equal(CareOrderInfo.CareTags)
	orderInfo.Value("min_price").Equal(common.Ftos(CareOrderInfo.MinPrice))
	orderInfo.Value("max_price").Equal(common.Ftos(CareOrderInfo.MaxPrice))
	orderInfo.Value("cover").Equal(CareOrderInfo.Cover)
	orderInfo.Value("care_type").Equal(CareOrderInfo.CareType)
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
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/inner/order/pay/{id}", careOrderCareId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("qrcode").NotNull()
}

func TestCareOrderCancelNoPayCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/inner/order/cancel/{id}", careOrderCareId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestCareOrderCancelPayCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/inner/order/cancel/{id}", CareOrder.ID).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestCareOrderShowReturnSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/inner/order/{id}", CareOrder.ID).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(CareOrder.ID)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(common.I_ORDER_STATUS_FOR_CANCEL)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")
	obj.Value("data").Object().Value("total").Equal(common.Ftos(CareOrder.Total))
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

	obj.Value("data").Object().Value("total").Equal(fmt.Sprintf("%.2f", f))
	obj.Value("data").Object().Value("rmk").Equal(CareOrder.Rmk)
	obj.Value("data").Object().Value("pay_type").Equal(CareOrder.PayType)
	obj.Value("data").Object().Value("is_return").Equal(1)

	obj.Value("data").Object().Value("order_carer_info").Null()
	obj.Value("data").Object().Value("comments").Array().Length().Equal(1)
	comment := obj.Value("data").Object().Value("comments").Array().First().Object()
	comment.Value("id").NotNull()
	comment.Value("user_id").Equal(User.ID)
	comment.Value("application_id").Equal(AppId)
	comment.Value("content").Equal(CareOrderComment.Content)
	comment.Value("pics").Array().Length().Equal(2)
	comment.Value("pics").Array().First().Equal("https")
	comment.Value("pics").Array().Last().Equal("https")

	orderInfo := obj.Value("data").Object().Value("order_info").Object()
	orderInfo.Value("id").NotNull()
	orderInfo.Value("name").Equal(CareOrderInfo.Name)
	orderInfo.Value("desc").Equal(CareOrderInfo.Desc)
	orderInfo.Value("application_name").Equal(CareOrderInfo.ApplicationName)
	orderInfo.Value("time_type").Equal(CareOrderInfo.TimeType)
	orderInfo.Value("care_detail").Equal(CareOrderInfo.CareDetail)
	orderInfo.Value("care_tags").Equal(CareOrderInfo.CareTags)
	orderInfo.Value("min_price").Equal(common.Ftos(CareOrderInfo.MinPrice))
	orderInfo.Value("max_price").Equal(common.Ftos(CareOrderInfo.MaxPrice))
	orderInfo.Value("cover").Equal(CareOrderInfo.Cover)
	orderInfo.Value("care_type").Equal(CareOrderInfo.CareType)
	orderInfo.Value("care_order_id").Equal(CareOrder.ID)

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal(CareOrderAddr.Name)
	addr.Value("loc_name").Equal(CareOrderAddr.LocName)
	addr.Value("bed_num").Equal(CareOrderAddr.BedNum)
	addr.Value("hospital_no").Equal(CareOrderAddr.HospitalNo)
	addr.Value("disease").Equal(CareOrderAddr.Disease)
	addr.Value("care_order_id").Equal(CareOrderAddr.CareOrderID)
	addr.Value("sex").Equal(CareOrderAddr.Sex)
	addr.Value("hospital_name").Equal(CareOrderAddr.HospitalName)
	addr.Value("phone").Equal(CareOrderAddr.Phone)
	addr.Value("age").Equal(CareOrderAddr.Age)
	addr.Value("addr").Equal(CareOrderAddr.Addr)

	orderReturn := obj.Value("data").Object().Value("return_order").Object()
	orderReturn.Value("id").NotNull()
	orderReturn.Value("order_no").String().Contains("RC")
	orderReturn.Value("status").Object().Value("value").Equal(common.I_RETURN_ORDER_STATUS_FOR_AUDIT)
	orderReturn.Value("status").Object().Value("text").Equal("待审核")
	orderReturn.Value("total").Equal(common.Ftos(CareOrder.Total))
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
	returnAddr.Value("name").Equal(CareOrderAddr.Name)
	returnAddr.Value("sex").Equal(CareOrderAddr.Sex)
	returnAddr.Value("care_return_order_id").NotNull()
	returnAddr.Value("loc_name").Equal(CareOrderAddr.LocName)
	returnAddr.Value("hospital_no").Equal(CareOrderAddr.HospitalNo)
	returnAddr.Value("hospital_name").Equal(CareOrderAddr.HospitalName)
	returnAddr.Value("age").Equal(CareOrderAddr.Age)
	returnAddr.Value("disease").Equal(CareOrderAddr.Disease)
	returnAddr.Value("phone").Equal(CareOrderAddr.Phone)
	returnAddr.Value("addr").Equal(CareOrderAddr.Addr)

	returnOrderInfo := orderReturn.Value("order_info").Object()
	returnOrderInfo.Value("id").NotNull()
	returnOrderInfo.Value("name").Equal(CareOrderInfo.Name)
	returnOrderInfo.Value("desc").Equal(CareOrderInfo.Desc)
	returnOrderInfo.Value("application_name").Equal(CareOrderInfo.ApplicationName)
	returnOrderInfo.Value("time_type").Equal(CareOrderInfo.TimeType)
	returnOrderInfo.Value("care_detail").Equal(CareOrderInfo.CareDetail)
	returnOrderInfo.Value("care_tags").Equal(CareOrderInfo.CareTags)
	returnOrderInfo.Value("min_price").Equal(common.Ftos(CareOrderInfo.MinPrice))
	returnOrderInfo.Value("max_price").Equal(common.Ftos(CareOrderInfo.MaxPrice))
	returnOrderInfo.Value("cover").Equal(CareOrderInfo.Cover)
	returnOrderInfo.Value("care_type").Equal(CareOrderInfo.CareType)
	returnOrderInfo.Value("care_return_order_id").NotNull()

}

func TestCareOrderDeleteCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.DELETE("/care/v1/inner/order/{id}", careOrderCareId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestCareOrderDeleteCareRetrunSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.DELETE("/care/v1/inner/order/{id}", CareOrder.ID).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	CareOrderCount--
}
