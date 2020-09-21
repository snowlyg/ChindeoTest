package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var shopOrderId interface{}
var shopOrderTotal float64
var shopOrderSkuId interface{}

func TestShopOrderSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/inner/order").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("application_id", model.AppId).
		WithQuery("id_card_no", model.IdCardNo).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.ShopOrderCount)
}

func TestShopOrderStatusForDeliveryingSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	cate := model.CreateCate(Cate1.ID, 1)
	spu := model.CreateSpu(brand.ID, cate.ID, 1, "这是一个很神奇的中德澳商品", "", 10.00, 100.00, Spec)
	shopOrder := model.CreateShopOrder("S202008241612348468756915", User.ID, model.IOrderPayTypeAli, model.OrderAppTypeBed, model.IOrderStatusForDeliverying, spu.Skus)
	re := map[string]interface{}{
		"status":      model.IOrderStatusForDeliverying,
		"page_size":   10,
		"hospital_no": "9556854545",
		"id_card_no":  model.IdCardNo,
	}

	obj := model.GetE(t).GET("/shop/v1/inner/order").
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
	model.DelShopOrder(shopOrder)
}

func TestShopOrderStatusForDeliveryingNoIdCardNoSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	cate := model.CreateCate(Cate1.ID, 1)
	spu := model.CreateSpu(brand.ID, cate.ID, 1, "这是一个很神奇的中德澳商品", "", 10.00, 100.00, Spec)
	shopOrder := model.CreateShopOrder("S202008241612348468756915", User.ID, model.IOrderPayTypeAli, model.OrderAppTypeBed, model.IOrderStatusForDeliverying, spu.Skus)
	re := map[string]interface{}{
		"status":    model.IOrderStatusForDeliverying,
		"page_size": 10,
	}

	obj := model.GetE(t).GET("/shop/v1/inner/order").
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
	model.DelShopOrder(shopOrder)
}

func TestShopOrderNoStatusSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	cate := model.CreateCate(Cate1.ID, 1)
	spu := model.CreateSpu(brand.ID, cate.ID, 1, "这是一个很神奇的中德澳商品", "", 10.00, 100.00, Spec)
	shopOrder := model.CreateShopOrder("S202008241612348468756915", User.ID, model.IOrderPayTypeAli, model.OrderAppTypeBed, model.IOrderStatusForDeliverying, spu.Skus)
	re := map[string]interface{}{
		"status":      0,
		"page_size":   10,
		"hospital_no": "9556854545",
		"id_card_no":  model.IdCardNo,
	}

	obj := model.GetE(t).GET("/shop/v1/inner/order").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQueryObject(re).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.ShopOrderCount)
	model.DelShopOrder(shopOrder)
}

func TestShopOrderStatusForFinishSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	cate := model.CreateCate(Cate1.ID, 1)
	spu := model.CreateSpu(brand.ID, cate.ID, 1, "这是一个很神奇的中德澳商品", "", 10.00, 100.00, Spec)
	shopOrder := model.CreateShopOrder("S202008241612348468756915", User.ID, model.IOrderPayTypeAli, model.OrderAppTypeBed, model.IOrderStatusForDeliverying, spu.Skus)
	re := map[string]interface{}{
		"status":      model.IOrderStatusForFinish,
		"page_size":   10,
		"hospital_no": "9556854545",
		"id_card_no":  model.IdCardNo,
	}

	obj := model.GetE(t).GET("/shop/v1/inner/order").
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
	model.DelShopOrder(shopOrder)
}

func TestShopOrderWithKeyWordSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	cate := model.CreateCate(Cate1.ID, 1)
	spu := model.CreateSpu(brand.ID, cate.ID, 1, "这是一个很神奇的中德澳商品", "", 10.00, 100.00, Spec)
	ShopOrder = model.CreateShopOrder("S202008241612348468756915", User.ID, model.IOrderPayTypeAli, model.OrderAppTypeBed, model.IOrderStatusForDelivery, spu.Skus)
	obj := model.GetE(t).GET("/shop/v1/inner/order").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithQuery("application_id", model.AppId).
		WithQuery("id_card_no", model.IdCardNo).
		WithQuery("key_word", "中德澳").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.ShopOrderCount)
}

func TestShopOrderAddSuccess(t *testing.T) {
	brand := model.CreateBrand(false)
	name := "这是一个很神奇的商品"
	title := "这是一个很神奇的商品的超厉害的副标题"
	spu := model.CreateSpu(brand.ID, Cate1.ID, 3, name, title, 10.00, 100.00, Spec)
	stock1 := spu.Skus[0].Stock
	stock2 := spu.Skus[1].Stock
	mun := 3
	shopOrder := map[string]interface{}{
		"sku_ids": []map[string]interface{}{
			{
				"id":  spu.Skus[0].ID,
				"num": mun,
			},
			{
				"id":  spu.Skus[1].ID,
				"num": mun,
			},
		},
		"rmk":            "年轻貌美",
		"application_id": model.AppId,
		"id_card_no":     model.IdCardNo,
		"patient_name":   "操蛋",
		"loc_name":       "泥马",
		"bed_num":        05,
		"hospital_no":    "9556854545",
		"disease":        "disease",
		"addr":           "addr",
		"phone":          "13800138000",
		"age":            1,
	}

	shopOrderTotal = spu.Skus[0].Price*3 + spu.Skus[1].Price*3

	obj := model.GetE(t).POST("/shop/v1/inner/order/add").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(shopOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").NotEqual(0)
	obj.Value("data").Object().Value("order_no").String().Contains("S")
	obj.Value("data").Object().Value("total").Equal(shopOrderTotal)
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("app_type").Object().Value("value").Equal(model.OrderAppTypeMini)
	obj.Value("data").Object().Value("app_type").Object().Value("text").Equal("小程序")
	obj.Value("data").Object().Value("application_id").Equal(13)
	shopOrderId = obj.Value("data").Object().Value("id").Raw()

	skus := model.GetSkuByIds([]int{spu.Skus[0].ID, spu.Skus[1].ID})
	assert.Equal(t, stock1-mun, skus[0].Stock)
	assert.Equal(t, stock2-mun, skus[1].Stock)

	model.ShopOrderCount++

}

func TestShopOrderAddNoBedNumError(t *testing.T) {
	brand := model.CreateBrand(false)
	name := "这是一个很神奇的商品"
	title := "这是一个很神奇的商品的超厉害的副标题"
	spu := model.CreateSpu(brand.ID, Cate1.ID, 3, name, title, 10.00, 100.00, Spec)
	stock1 := spu.Skus[0].Stock
	stock2 := spu.Skus[1].Stock
	mun := 3
	shopOrder := map[string]interface{}{
		"sku_ids": []map[string]interface{}{
			{
				"id":  spu.Skus[0].ID,
				"num": mun,
			},
			{
				"id":  spu.Skus[1].ID,
				"num": mun,
			},
		},
		"rmk":            "年轻貌美",
		"application_id": model.AppId,
		"id_card_no":     model.IdCardNo,
		"patient_name":   "操蛋",
		"loc_name":       "泥马",
		"hospital_no":    "9556854545",
		"disease":        "disease",
		"addr":           "addr",
		"phone":          "13800138000",
		"age":            1,
	}

	shopOrderTotal = spu.Skus[0].Price*3 + spu.Skus[1].Price*3

	obj := model.GetE(t).POST("/shop/v1/inner/order/add").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(shopOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("床号不能为空！")

	skus := model.GetSkuByIds([]int{spu.Skus[0].ID, spu.Skus[1].ID})
	assert.Equal(t, skus[0].Stock, stock1)
	assert.Equal(t, skus[1].Stock, stock2)
}

func TestShopOrderNoPatientNameError(t *testing.T) {
	brand := model.CreateBrand(false)
	name := "这是一个很神奇的商品"
	title := "这是一个很神奇的商品的超厉害的副标题"
	spu := model.CreateSpu(brand.ID, Cate1.ID, 3, name, title, 10.00, 100.00, Spec)
	stock1 := spu.Skus[0].Stock
	stock2 := spu.Skus[1].Stock
	mun := 3
	shopOrder := map[string]interface{}{
		"sku_ids": []map[string]interface{}{
			{
				"id":  spu.Skus[0].ID,
				"num": mun,
			},
			{
				"id":  spu.Skus[1].ID,
				"num": mun,
			},
		},
		"rmk":            "年轻貌美",
		"application_id": model.AppId,
		"id_card_no":     model.IdCardNo,
		"loc_name":       "泥马",
		"bed_num":        05,
		"hospital_no":    "9556854545",
		"disease":        "disease",
		"addr":           "addr",
		"phone":          "13800138000",
		"age":            1,
	}

	obj := model.GetE(t).POST("/shop/v1/inner/order/add").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(shopOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("患者名称不能为空！")
	skus := model.GetSkuByIds([]int{spu.Skus[0].ID, spu.Skus[1].ID})
	assert.Equal(t, skus[0].Stock, stock1)
	assert.Equal(t, skus[1].Stock, stock2)
}

func TestShopOrderNoLocNameError(t *testing.T) {
	brand := model.CreateBrand(false)
	name := "这是一个很神奇的商品"
	title := "这是一个很神奇的商品的超厉害的副标题"
	spu := model.CreateSpu(brand.ID, Cate1.ID, 3, name, title, 10.00, 100.00, Spec)
	stock1 := spu.Skus[0].Stock
	stock2 := spu.Skus[1].Stock
	mun := 3
	shopOrder := map[string]interface{}{
		"sku_ids": []map[string]interface{}{
			{
				"id":  spu.Skus[0].ID,
				"num": mun,
			},
			{
				"id":  spu.Skus[1].ID,
				"num": mun,
			},
		},
		"rmk":          "年轻貌美",
		"id_card_no":   model.IdCardNo,
		"patient_name": "操蛋",
		"bed_num":      05,
		"hospital_no":  "9556854545",
		"disease":      "disease",
		"addr":         "addr",
		"phone":        "13800138000",
		"age":          1,
	}
	obj := model.GetE(t).POST("/shop/v1/inner/order/add").
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(shopOrder).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("病区/科室不能为空！")
	skus := model.GetSkuByIds([]int{spu.Skus[0].ID, spu.Skus[1].ID})
	assert.Equal(t, skus[0].Stock, stock1)
	assert.Equal(t, skus[1].Stock, stock2)
}

func TestShopOrderShowSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/inner/order/{id}", shopOrderId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(model.GetSToI(shopOrderId))
	obj.Value("data").Object().Value("order_no").String().Contains("S")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForPay)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("待付款")

	obj.Value("data").Object().Value("total").Equal(model.Ftos(shopOrderTotal))
	obj.Value("data").Object().Value("rmk").Equal("年轻貌美")
	obj.Value("data").Object().Value("pay_type").Object().Value("value").Equal(model.IOrderPayTypeWechat)
	obj.Value("data").Object().Value("pay_type").Object().Value("text").Equal("微信")
	obj.Value("data").Object().Value("is_return").Object().Value("value").Equal(0)
	obj.Value("data").Object().Value("is_return").Object().Value("text").Equal("未退款")

	obj.Value("data").Object().Value("skus").Array().Length().Equal(2)
	sku := obj.Value("data").Object().Value("skus").Array().First().Object()

	shopOrderSkuId = sku.Value("id").Raw()

	sku.Value("id").NotNull()
	sku.Value("title").NotNull()
	sku.Value("sku_no").NotNull()
	sku.Value("images").Array().NotEmpty()
	sku.Value("price").NotNull()
	sku.Value("indexes").NotNull()
	sku.Value("own_spec").NotNull()
	sku.Value("num").NotNull()
	sku.Value("shop_order_id").Equal(model.GetSToI(shopOrderId))
	sku.Value("sku_id").NotNull()
	sku.Value("application_id").NotNull()
	sku.Value("comments").NotNull()

	obj.Value("data").Object().Value("return_order").Null()

	obj.Value("data").Object().Value("comments").Array().Length().Equal(0)

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal("操蛋")
	addr.Value("sex").Equal(1)
	addr.Value("shop_order_id").Equal(model.GetSToI(shopOrderId))
	addr.Value("loc_name").Equal("泥马")
	addr.Value("hospital_no").Equal("9556854545")
	addr.Value("hospital_name").Equal("我的医院")
	addr.Value("age").Equal(1)
	addr.Value("disease").Equal("disease")
	addr.Value("phone").Equal("13800138000")
	addr.Value("addr").Equal("")

}

func TestShopOrderPaySuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/inner/order/pay/{id}", shopOrderId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	//data := fmt.Sprintf("支付返回数据 %v",obj.Value("data").Raw())
	//fmt.Print(data)
}

func TestShopOrderCancelNoPaySuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/inner/order/cancel/{id}", shopOrderId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestShopOrderCancelPaySuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/inner/order/cancel/{id}", ShopOrder.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
}

func TestShopOrderShowReturnSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/shop/v1/inner/order/{id}", ShopOrder.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(ShopOrder.ID)
	obj.Value("data").Object().Value("order_no").String().Contains("S")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForCancel)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")
	obj.Value("data").Object().Value("total").Equal(model.Ftos(ShopOrder.Total))

	obj.Value("data").Object().Value("total").Equal("0.00")
	obj.Value("data").Object().Value("rmk").Equal(ShopOrder.Rmk)
	obj.Value("data").Object().Value("pay_type").Object().Value("value").Equal(ShopOrder.PayType)
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
	sku.Value("shop_order_id").Equal(ShopOrder.ID)
	sku.Value("sku_id").NotNull()
	sku.Value("application_id").NotNull()
	sku.Value("comments").NotNull()

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal(ShopOrder.Addr.Name)
	addr.Value("sex").Equal(ShopOrder.Addr.Sex)
	addr.Value("shop_order_id").Equal(ShopOrder.ID)
	addr.Value("loc_name").Equal(ShopOrder.Addr.LocName)
	addr.Value("hospital_no").Equal(ShopOrder.Addr.HospitalNo)
	addr.Value("hospital_name").Equal(ShopOrder.Addr.HospitalName)
	addr.Value("age").Equal(ShopOrder.Addr.Age)
	addr.Value("disease").Equal(ShopOrder.Addr.Disease)
	addr.Value("phone").Equal(ShopOrder.Addr.Phone)
	addr.Value("addr").Equal(ShopOrder.Addr.Addr)

	orderReturn := obj.Value("data").Object().Value("return_order").Object()
	orderReturn.Value("id").NotNull()
	orderReturnId := model.GetF(orderReturn.Value("id").Raw())
	orderReturn.Value("order_no").String().Contains("RS")
	orderReturn.Value("status").Object().Value("value").Equal(model.IReturnOrderStatusForAudit)
	orderReturn.Value("status").Object().Value("text").Equal("待审核")
	orderReturn.Value("total").Equal(model.Ftos(ShopOrder.Total))
	orderReturn.Value("open_id").Equal("")
	orderReturn.Value("app_type").Equal(ShopOrder.AppType)
	orderReturn.Value("trade_type").Equal("")
	orderReturn.Value("shop_order_id").Equal(ShopOrder.ID)
	orderReturn.Value("application_id").Equal(ShopOrder.ApplicationID)
	orderReturn.Value("pay_type").Equal(ShopOrder.PayType)

	returnAddr := orderReturn.Value("addr").Object()
	returnAddr.Value("id").NotNull()
	returnAddr.Value("name").Equal(ShopOrder.Addr.Name)
	returnAddr.Value("sex").Equal(ShopOrder.Addr.Sex)
	returnAddr.Value("shop_return_order_id").Equal(orderReturnId)
	returnAddr.Value("loc_name").Equal(ShopOrder.Addr.LocName)
	returnAddr.Value("hospital_no").Equal(ShopOrder.Addr.HospitalNo)
	returnAddr.Value("hospital_name").Equal(ShopOrder.Addr.HospitalName)
	returnAddr.Value("age").Equal(ShopOrder.Addr.Age)
	returnAddr.Value("disease").Equal(ShopOrder.Addr.Disease)
	returnAddr.Value("phone").Equal(ShopOrder.Addr.Phone)
	returnAddr.Value("addr").Equal(ShopOrder.Addr.Addr)

	orderReturn.Value("skus").Array().Length().Equal(1)
	returnSku := orderReturn.Value("skus").Array().First().Object()
	returnSku.Value("id").NotNull()
	returnSku.Value("title").NotNull()
	returnSku.Value("sku_no").NotNull()
	returnSku.Value("images").Array().NotEmpty()
	returnSku.Value("price").NotNull()
	returnSku.Value("indexes").NotNull()
	returnSku.Value("own_spec").NotNull()
	returnSku.Value("num").NotNull()
	returnSku.Value("shop_return_order_id").Equal(orderReturnId)
	returnSku.Value("application_id").NotNull()
}

func TestShopOrderCommentSuccess(t *testing.T) {
	comment := map[string]interface{}{
		"star":           1,
		"content":        "content",
		"pics":           model.GetPics(),
		"is_unnamed":     1,
		"application_id": model.AppId,
	}

	obj := model.GetE(t).POST("/shop/v1/inner/comment/{shop_order_sku_id}", shopOrderSkuId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("评论成功")
}

func TestShopOrderShowAfterCommentSuccess(t *testing.T) {
	obj := model.GetE(t).GET("/shop/v1/inner/order/{id}", shopOrderId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	obj.Value("data").Object().Value("id").Equal(model.GetSToI(shopOrderId))
	obj.Value("data").Object().Value("order_no").String().Contains("S")
	obj.Value("data").Object().Value("status").Object().Value("value").Equal(model.IOrderStatusForCancel)
	obj.Value("data").Object().Value("status").Object().Value("text").Equal("已取消")

	obj.Value("data").Object().Value("total").Equal(model.Ftos(shopOrderTotal))
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
	sku.Value("shop_order_id").Equal(model.GetSToI(shopOrderId))
	sku.Value("sku_id").NotNull()
	sku.Value("application_id").NotNull()
	sku.Value("comments").NotNull()

	obj.Value("data").Object().Value("return_order").Null()

	obj.Value("data").Object().Value("comments").Array().Length().Equal(1)
	comments := obj.Value("data").Object().Value("comments").Array().First().Object()
	comments.Value("id").NotNull()

	addr := obj.Value("data").Object().Value("addr").Object()
	addr.Value("id").NotNull()
	addr.Value("name").Equal("操蛋")
	addr.Value("sex").Equal(1)
	addr.Value("shop_order_id").Equal(model.GetSToI(shopOrderId))
	addr.Value("loc_name").Equal("泥马")
	addr.Value("hospital_no").Equal("9556854545")
	addr.Value("hospital_name").Equal("我的医院")
	addr.Value("age").Equal(1)
	addr.Value("disease").Equal("disease")
	addr.Value("phone").Equal("13800138000")
	addr.Value("addr").Equal("")

}

func TestShopOrderCommentNoContentError(t *testing.T) {
	comment := map[string]interface{}{
		"star":           1,
		"content":        "",
		"pics":           model.GetPics(),
		"is_unnamed":     0,
		"application_id": model.AppId,
	}

	obj := model.GetE(t).POST("/shop/v1/inner/comment/{shop_order_sku_id}", shopOrderSkuId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("评论内容不能为空！")
}

func TestShopOrderCommentNoOrderIdError(t *testing.T) {
	comment := map[string]interface{}{
		"star":           1,
		"content":        "456952158962254456",
		"pics":           model.GetPics(),
		"is_unnamed":     0,
		"application_id": model.AppId,
	}

	obj := model.GetE(t).POST("/shop/v1/inner/comment/{shop_order_sku_id}", 0).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		WithJSON(comment).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("参数错误")
}

func TestShopOrderDeleteSuccess(t *testing.T) {

	obj := model.GetE(t).DELETE("/shop/v1/inner/order/{id}", shopOrderId).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	model.ShopOrderCount--
}

func TestShopOrderDeleteRetrunSuccess(t *testing.T) {
	obj := model.GetE(t).DELETE("/shop/v1/inner/order/{id}", ShopOrder.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Contains("请求成功")
	model.ShopOrderCount--
}
