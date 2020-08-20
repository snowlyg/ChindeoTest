package main

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/shopspring/decimal"
)

var careOrderCareId float64
var careId float64
var carePrice decimal.Decimal
var careTimeTypeText string

func TestCareListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.GET("/care/v1/inner/care").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")

}

func TestCareListNoPageSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.GET("/care/v1/inner/care").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
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
		careId = data
	}

	maxPrice := obj.Value("data").Array().First().Object().Value("max_price").Raw()
	total, ok := maxPrice.(string)
	if ok {
		carePriceF, _ := strconv.ParseFloat(total, 10)
		carePrice = decimal.NewFromFloat(carePriceF)
	}

	timeType := obj.Value("data").Array().First().Object().Value("time_type").Raw()
	tt, ok := timeType.(string)
	if ok {
		careTimeTypeText = tt
	}

}

func TestCareShowSuccess(t *testing.T) {

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.GET("/care/v1/inner/care/{id}", careId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(careId)

}

func TestCareOrderListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})
	obj := e.GET("/care/v1/inner/order").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
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
		"care_id":      careId,
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
		BaseURL: BaseUrl,
	})

	obj := e.POST("/care/v1/inner/order/add").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
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
	obj.Value("data").Object().Value("app_type").Equal(2)
	obj.Value("data").Object().Value("application_id").Equal(13)

	id := obj.Value("data").Object().Value("id").Raw()
	data, ok := id.(string)
	if ok {
		careOrderCareId, _ = strconv.ParseFloat(data, 10)
	}
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
		"care_id":      careId,
		"id_card_no":   "456952158962254456",
	}

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})

	obj := e.POST("/care/v1/inner/order/add").
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		WithJSON(careOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("患者名称不能为空！")
}

func TestCareOrderShowCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})

	obj := e.GET("/care/v1/inner/order/{id}", careOrderCareId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(careOrderCareId)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(1)
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
	obj.Value("data").Object().Value("total").Equal(fmt.Sprintf("%.2f", f))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("pay_type").Equal(1)
	obj.Value("data").Object().Value("is_return").Equal(0)
	obj.Value("data").Object().Value("order_carer_info").Null()
	obj.Value("data").Object().Value("return_order").Null()
	orderInfo := obj.Value("data").Object().Value("order_info").Object()
	orderInfo.Value("id").NotNull()
	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal("操蛋")
	addr.Value("loc_name").Equal("泥马")
	addr.Value("bed_num").Equal("05")
	addr.Value("hospital_no").Equal("9556854545")
	addr.Value("disease").Equal("玩玩")
}

func TestCareOrderPayCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})

	obj := e.GET("/care/v1/inner/order/pay/{id}", careOrderCareId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("qrcode").NotNull()
}

func TestCareOrderCancelCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: BaseUrl,
	})

	obj := e.GET("/care/v1/inner/order/cancel/{id}", careOrderCareId).
		WithHeaders(map[string]string{"X-Token": Token, "AuthType": "4"}).
		WithCookie("PHPSESSID", PHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}
