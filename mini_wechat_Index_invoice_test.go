package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

var invoiceId int

func TestMiniWechatIndexInvoiceCreateSuccess(t *testing.T) {
	re := map[string]interface{}{
		"type":        "type",
		"title":       "title",
		"taxNumber":   "taxNumber",
		"telephone":   "telephone",
		"bankName":    "bankName",
		"bankAccount": "bankAccount",
	}
	obj := model.GetE(t).POST("/index/invoice/create").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithJSON(re).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("创建发票信息成功")
	id := obj.Value("data").Object().Value("id").Raw()
	invoiceId = model.GetSToI(id)
	println(id)
	println(invoiceId)
	println()
}

func TestMiniWechatIndexInvoiceIndexSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/index/invoice/index").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("获取发票信息成功")
}

func TestMiniWechatIndexInvoiceDeleteSuccess(t *testing.T) {

	obj := model.GetE(t).POST("/index/invoice/delete").
		WithHeaders(model.GetMiniHeader("")).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		WithQuery("id", invoiceId).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("删除发票信息成功")
}
