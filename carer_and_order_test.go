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

var careOrderCarerId float64
var carerId float64
var carerTimeTypeText string
var carerPrice decimal.Decimal

func TestCarrListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/inner/carer").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")

}

func TestCarrListNoPageSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/inner/carer").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithQuery("page_size", "-1").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	id := obj.Value("data").Array().First().Object().Value("id").Raw()
	data, ok := id.(float64)
	if ok {
		carerId = data
	}

	maxPrice := obj.Value("data").Array().First().Object().Value("price").Raw()
	total, ok := maxPrice.(string)
	if ok {
		carerPriceF, _ := strconv.ParseFloat(total, 10)
		carerPrice = decimal.NewFromFloat(carerPriceF)
	}

	timeType := obj.Value("data").Array().First().Object().Value("time_type").Raw()
	tt, ok := timeType.(string)
	if ok {
		carerTimeTypeText = tt
	}
}

func TestCarrShowSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/inner/carer/{id}", carerId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(carerId)

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
		"carer_id":     carerId,
		"start_at":     carerStartAt.Format("2006-01-02 15:04:05"),
		"end_at":       carerEndAt.Format("2006-01-02 15:04:05"),
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
	obj.Value("data").Object().Value("start_at").Equal(carerStartAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("end_at").Equal(carerEndAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("app_type").Equal(2)
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
		"carer_id":     carerId,
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
		"carer_id":     carerId,
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
	obj.Value("message").String().Contains("时间段已经被预约")
}

func TestCarrOrderShowCarerSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/inner/order/{id}", careOrderCarerId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(careOrderCarerId)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(1)
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
	orderInfo := obj.Value("data").Object().Value("order_carer_info").Object()
	orderInfo.Value("id").NotNull()

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal("操蛋")
	addr.Value("loc_name").Equal("泥马")
	addr.Value("bed_num").Equal("05")
	addr.Value("hospital_no").Equal("9556854545")
	addr.Value("disease").Equal("玩玩")
}

func TestCarrOrderPayCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/inner/order/pay/{id}", careOrderCarerId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("qrcode").NotNull()
}

func TestCarrOrderCancelCarerSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(),
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/inner/order/cancel/{id}", careOrderCarerId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestCarrOrderDeleteCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.DELETE("/care/v1/inner/order/{id}", careOrderCarerId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestCarrOrderDeleteCareRetrunSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.DELETE("/care/v1/inner/order/{id}", CarerOrder.ID).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10)}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}
