package main

import (
	"github.com/snowlyg/ChindeoTest/model"
	"net/http"
	"testing"
)

func TestMiniWechatOnlineArticleListSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/online/v1/article").
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(model.ArticleCount)
	obj.Value("data").Object().Value("data").Array().First().Object().Value("id").Equal(Article.ID)
	obj.Value("data").Object().Value("data").Array().First().Object().Value("title").Equal(Article.Title)
	obj.Value("data").Object().Value("data").Array().First().Object().Value("digest").Equal(Article.Digest)
	obj.Value("data").Object().Value("data").Array().First().Object().Value("content").Equal(Article.Content)
	obj.Value("data").Object().Value("data").Array().First().Object().Value("author").Equal(Article.Author)
	obj.Value("data").Object().Value("data").Array().First().Object().Value("local_url").Equal(Article.LocalURL)
	obj.Value("data").Object().Value("data").Array().First().Object().Value("content_source_url").Equal(Article.ContentSourceURL)
	obj.Value("data").Object().Value("data").Array().First().Object().Value("create_at").NotNull()
	obj.Value("data").Object().Value("data").Array().First().Object().Value("update_at").NotNull()
}

func TestMiniWechatOnlineArticleShowSuccess(t *testing.T) {

	obj := model.GetE(t).GET("/online/v1/article/{id}", Article.ID).
		WithHeaders(model.GetMiniHeader()).
		WithCookie("PHPSESSID", model.GetMiniSessionId()).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Value("id").Equal(Article.ID)
	obj.Value("data").Object().Value("title").Equal(Article.Title)
	obj.Value("data").Object().Value("digest").Equal(Article.Digest)
	obj.Value("data").Object().Value("content").Equal(Article.Content)
	obj.Value("data").Object().Value("author").Equal(Article.Author)
	obj.Value("data").Object().Value("local_url").Equal(Article.LocalURL)
	obj.Value("data").Object().Value("content_source_url").Equal(Article.ContentSourceURL)
	obj.Value("data").Object().Value("create_at").NotNull()
	obj.Value("data").Object().Value("update_at").NotNull()
}
