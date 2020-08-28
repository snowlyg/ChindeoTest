package main

import (
	"github.com/snowlyg/ChindeoTest/common"
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestMiniWechatOnlineArticleListSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/online/v1/article").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("请求成功")
	obj.Value("data").Object().Keys().ContainsOnly("total", "per_page", "current_page", "last_page", "data")
	obj.Value("data").Object().Value("data").Array().Length().Equal(1)
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
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	obj := e.GET("/online/v1/article/{id}", Article.ID).
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
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
