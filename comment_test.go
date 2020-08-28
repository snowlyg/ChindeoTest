package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestCommentDeleteSuccess(t *testing.T) {
	obj := model.GetE(t).DELETE("/common/v1/comment/{id}", MiniOrderCommentForDel.ID).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("删除成功")
}

func TestCommentDeleteInnerSuccess(t *testing.T) {
	obj := model.GetE(t).DELETE("/common/v1/inner/comment/{id}", CareOrderCommentForDel.ID).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("删除成功")
	CareOrderCommentCount--
}
func TestCommentDeleteInnerIdNotExistsError(t *testing.T) {
	obj := model.GetE(t).DELETE("/common/v1/inner/comment/{id}", 9999).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("数据不存在")
}

func TestCommentDeleteInnerForZoreIdError(t *testing.T) {
	obj := model.GetE(t).DELETE("/common/v1/inner/comment/{id}", 0).
		WithHeaders(model.GetHeader()).
		WithCookie("PHPSESSID", model.GetSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("错误数据")
}
