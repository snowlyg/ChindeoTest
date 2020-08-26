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

var miniCareOrderCareId float64
var miniCarePrice decimal.Decimal
var miniCareTimeTypeText string

func TestMiniWechatCareListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/care").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithQuery("application_id", AppId).
		WithQuery("care_type_id", CareType.ID).
		WithQuery("care_tag_id", CareTag.ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(1)

	fitst := obj.Value("data").Object().Value("data").Array().Last().Object()
	fitst.Value("id").Equal(Care.ID)
	fitst.Value("name").Equal(Care.Name)
	fitst.Value("desc").Equal(Care.Desc)
	fitst.Value("time_type").Equal(Care.TimeType)
	fitst.Value("amount").Equal(CareAmount)
	fitst.Value("cover").Equal(Care.Cover)
	fitst.Value("care_type_id").Equal(Care.CareTypeID)

	careType := fitst.Value("type").Object()
	careType.Value("id").Equal(CareType.ID)
	careType.Value("name").Equal(CareType.Name)
}

func TestMiniWechatCareNoTagIdListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/care").
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
	fitst.Value("id").Equal(Care.ID)
	fitst.Value("name").Equal(Care.Name)
	fitst.Value("desc").Equal(Care.Desc)
	fitst.Value("time_type").Equal(Care.TimeType)
	fitst.Value("amount").Equal(CareAmount)
	fitst.Value("cover").Equal(Care.Cover)
	fitst.Value("care_type_id").Equal(Care.CareTypeID)

	careType := fitst.Value("type").Object()
	careType.Value("id").Equal(CareType.ID)
	careType.Value("name").Equal(CareType.Name)
}

func TestMiniWechatCareListNoAppError(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/care").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithQuery("care_type_id", CareType.ID).
		WithQuery("care_tag_id", CareTag.ID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("请选择商家/医院")

}

func TestMiniWechatCareListNoPageSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/care").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithQuery("application_id", AppId).
		WithQuery("care_type_id", CareType.ID).
		WithQuery("care_tag_id", CareTag.ID).
		WithQuery("page_size", "-1").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Array().Length().Equal(1)

	fitst := obj.Value("data").Array().First().Object()
	fitst.Value("id").Equal(Care.ID)
	fitst.Value("name").Equal(Care.Name)
	fitst.Value("desc").Equal(Care.Desc)
	fitst.Value("time_type").Equal(Care.TimeType)
	fitst.Value("amount").Equal(CareAmount)
	fitst.Value("cover").Equal(Care.Cover)
	fitst.Value("care_type_id").Equal(Care.CareTypeID)

	careType := fitst.Value("type").Object()
	careType.Value("id").Equal(CareType.ID)
	careType.Value("name").Equal(CareType.Name)

}

func TestMiniWechatCareShowSuccess(t *testing.T) {

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/care/{id}", Care.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
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

func TestMiniWechatCareOrderListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/order").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithQuery("application_id", AppId).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(CareOrderCount)
}

var miniCareStartAt time.Time
var miniCareEndAt time.Time

func TestMiniWechatCareOrderAddCareSuccess(t *testing.T) {

	miniCareStartAt = time.Now().AddDate(0, 0, 1)
	miniCareEndAt = time.Now().AddDate(0, 0, 2)
	careOrder := map[string]interface{}{
		"care_id":        Care.ID,
		"start_at":       miniCareStartAt.Format("2006-01-02 15:04:05"),
		"end_at":         miniCareEndAt.Format("2006-01-02 15:04:05"),
		"rmk":            "年轻貌美",
		"application_id": AppId,
		"addr_id":        Addr.ID,
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
	obj.Value("data").Object().Value("start_at").Equal(miniCareStartAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("end_at").Equal(miniCareEndAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("app_type").Equal(common.ORDER_APP_TYPE_MINI)
	obj.Value("data").Object().Value("application_id").Equal(13)
	miniCareOrderCareId, _ = strconv.ParseFloat(common.GetS(obj.Value("data").Object().Value("id").Raw()), 10)
}

func TestMiniWechatCareShowAfterOrderAddSuccess(t *testing.T) {

	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/care/{id}", Care.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
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
	miniCarePrice = decimal.NewFromFloat(carePriceF)
	miniCareTimeTypeText = common.GetS(obj.Value("data").Object().Value("time_type").Raw())
}

func TestMiniWechatCareOrderNoAddrError(t *testing.T) {
	startAt := time.Now().AddDate(0, 0, 5)
	endAt := time.Now().AddDate(0, 0, 6)
	careOrder := map[string]interface{}{
		"start_at":       startAt.Format("2006-01-02 15:04:05"),
		"end_at":         endAt.Format("2006-01-02 15:04:05"),
		"rmk":            "年轻貌美",
		"care_id":        Care.ID,
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
	obj.Value("message").String().Equal("收货地址不存在")
}

func TestMiniWechatCareOrderNoAppError(t *testing.T) {
	careOrder := map[string]interface{}{
		"care_id":  Care.ID,
		"start_at": miniCareStartAt.Format("2006-01-02 15:04:05"),
		"end_at":   miniCareEndAt.Format("2006-01-02 15:04:05"),
		"rmk":      "年轻貌美",
		"addr_id":  Addr.ID,
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

func TestMiniWechatCareOrderShowCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/order/{id}", miniCareOrderCareId).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(miniCareOrderCareId)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(1)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("待付款")
	obj.Value("data").Object().Value("start_at").Equal(miniCareStartAt.Format("2006-01-02 15:04:05"))
	obj.Value("data").Object().Value("end_at").Equal(miniCareEndAt.Format("2006-01-02 15:04:05"))

	var total decimal.Decimal
	sub := int(miniCareEndAt.Sub(miniCareStartAt).Hours())
	if miniCareTimeTypeText == "天" {
		total = miniCarePrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if miniCareTimeTypeText == "时" {
		total = miniCarePrice.Mul(decimal.NewFromFloat(float64(sub)))
	}

	f, _ := total.Float64()

	obj.Value("data").Object().Value("total").Equal(common.Ftos(f))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("pay_type").Equal(1)
	obj.Value("data").Object().Value("is_return").Equal(0)
	obj.Value("data").Object().Value("order_carer_info").Null()
	obj.Value("data").Object().Value("return_order").Null()
	orderInfo := obj.Value("data").Object().Value("order_info").Object()
	orderInfo.Value("id").NotNull()
	orderInfo.Value("name").Equal(MiniCareOrderInfo.Name)
	orderInfo.Value("desc").Equal(MiniCareOrderInfo.Desc)
	orderInfo.Value("application_name").Equal("我的医院")
	orderInfo.Value("time_type").Equal(MiniCareOrderInfo.TimeType)
	orderInfo.Value("care_detail").Equal(MiniCareOrderInfo.CareDetail)
	orderInfo.Value("care_tags").Equal(MiniCareOrderInfo.CareTags)
	orderInfo.Value("min_price").Equal(common.Ftos(MiniCareOrderInfo.MinPrice))
	orderInfo.Value("max_price").Equal(common.Ftos(MiniCareOrderInfo.MaxPrice))
	orderInfo.Value("cover").Equal(MiniCareOrderInfo.Cover)
	orderInfo.Value("care_type").Equal(MiniCareOrderInfo.CareType)
	orderInfo.Value("care_order_id").Equal(miniCareOrderCareId)

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal(Addr.Name)
	addr.Value("sex").Equal(Addr.Sex)
	addr.Value("care_order_id").Equal(miniCareOrderCareId)
	addr.Value("loc_name").Equal(Addr.LocName)
	addr.Value("hospital_no").Equal(Addr.HospitalNo)
	addr.Value("hospital_name").Equal(Addr.HospitalName)
	addr.Value("age").Equal(Addr.Age)
	addr.Value("disease").Equal(Addr.Disease)
	addr.Value("phone").Equal(Addr.Phone)
	addr.Value("addr").Equal(Addr.Addr)
}

func TestMiniWechatCareOrderPayCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/order/pay/{id}", miniCareOrderCareId).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Contains("PARAM_ERROR:appid和openid不匹配")
}

func TestMiniWechatCareOrderCancelNoPayCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/order/cancel/{id}", miniCareOrderCareId).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatCareOrderCancelPayCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/care/v1/order/cancel/{id}", MiniCareOrder.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatCareOrderShowReturnSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
	obj := e.GET("/care/v1/order/{id}", MiniCareOrder.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(MiniCareOrder.ID)
	obj.Value("data").Object().Value("order_no").String().Contains("C")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(4)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")
	obj.Value("data").Object().Value("total").Equal(common.Ftos(MiniCareOrder.Total))
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

	orderInfo := obj.Value("data").Object().Value("order_info").Object()
	orderInfo.Value("id").NotNull()
	orderInfo.Value("name").Equal(MiniCareOrderInfo.Name)
	orderInfo.Value("desc").Equal(MiniCareOrderInfo.Desc)
	orderInfo.Value("application_name").Equal(MiniCareOrderInfo.ApplicationName)
	orderInfo.Value("time_type").Equal(MiniCareOrderInfo.TimeType)
	orderInfo.Value("care_detail").Equal(MiniCareOrderInfo.CareDetail)
	orderInfo.Value("care_tags").Equal(MiniCareOrderInfo.CareTags)
	orderInfo.Value("min_price").Equal(common.Ftos(MiniCareOrderInfo.MinPrice))
	orderInfo.Value("max_price").Equal(common.Ftos(MiniCareOrderInfo.MaxPrice))
	orderInfo.Value("cover").Equal(MiniCareOrderInfo.Cover)
	orderInfo.Value("care_type").Equal(MiniCareOrderInfo.CareType)
	orderInfo.Value("care_order_id").Equal(MiniCareOrder.ID)

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal(MiniCareOrderAddr.Name)
	addr.Value("loc_name").Equal(MiniCareOrderAddr.LocName)
	addr.Value("bed_num").Equal(MiniCareOrderAddr.BedNum)
	addr.Value("hospital_no").Equal(MiniCareOrderAddr.HospitalNo)
	addr.Value("disease").Equal(MiniCareOrderAddr.Disease)
	addr.Value("care_order_id").Equal(MiniCareOrderAddr.CareOrderID)
	addr.Value("sex").Equal(MiniCareOrderAddr.Sex)
	addr.Value("hospital_name").Equal(MiniCareOrderAddr.HospitalName)
	addr.Value("phone").Equal(MiniCareOrderAddr.Phone)
	addr.Value("age").Equal(MiniCareOrderAddr.Age)
	addr.Value("addr").Equal(MiniCareOrderAddr.Addr)

	orderReturn := obj.Value("data").Object().Value("return_order").Object()
	orderReturn.Value("id").NotNull()
	orderReturn.Value("order_no").String().Contains("RC")
	orderReturn.Value("status").Object().Value("value").Equal(common.I_RETURN_ORDER_STATUS_FOR_AUDIT)
	orderReturn.Value("status").Object().Value("text").Equal("待审核")
	orderReturn.Value("total").Equal(common.Ftos(MiniCareOrder.Total))
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
	returnAddr.Value("name").Equal(MiniCareOrderAddr.Name)
	returnAddr.Value("sex").Equal(MiniCareOrderAddr.Sex)
	returnAddr.Value("care_return_order_id").NotNull()
	returnAddr.Value("loc_name").Equal(MiniCareOrderAddr.LocName)
	returnAddr.Value("hospital_no").Equal(MiniCareOrderAddr.HospitalNo)
	returnAddr.Value("hospital_name").Equal(MiniCareOrderAddr.HospitalName)
	returnAddr.Value("age").Equal(MiniCareOrderAddr.Age)
	returnAddr.Value("disease").Equal(MiniCareOrderAddr.Disease)
	returnAddr.Value("phone").Equal(MiniCareOrderAddr.Phone)
	returnAddr.Value("addr").Equal(MiniCareOrderAddr.Addr)

	returnOrderInfo := orderReturn.Value("order_info").Object()
	returnOrderInfo.Value("id").NotNull()
	returnOrderInfo.Value("name").Equal(MiniCareOrderInfo.Name)
	returnOrderInfo.Value("desc").Equal(MiniCareOrderInfo.Desc)
	returnOrderInfo.Value("application_name").Equal(MiniCareOrderInfo.ApplicationName)
	returnOrderInfo.Value("time_type").Equal(MiniCareOrderInfo.TimeType)
	returnOrderInfo.Value("care_detail").Equal(MiniCareOrderInfo.CareDetail)
	returnOrderInfo.Value("care_tags").Equal(MiniCareOrderInfo.CareTags)
	returnOrderInfo.Value("min_price").Equal(common.Ftos(MiniCareOrderInfo.MinPrice))
	returnOrderInfo.Value("max_price").Equal(common.Ftos(MiniCareOrderInfo.MaxPrice))
	returnOrderInfo.Value("cover").Equal(MiniCareOrderInfo.Cover)
	returnOrderInfo.Value("care_type").Equal(MiniCareOrderInfo.CareType)
	returnOrderInfo.Value("care_return_order_id").NotNull()

}

func TestMiniWechatCareOrderDeleteCareSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.DELETE("/care/v1/order/{id}", miniCareOrderCareId).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestMiniWechatCareOrderDeleteCareRetrunSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.DELETE("/care/v1/order/{id}", MiniCareOrder.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	CareOrderCount--
}
