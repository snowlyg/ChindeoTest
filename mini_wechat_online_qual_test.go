package main

import (
	"github.com/snowlyg/ChindeoTest/common"
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"strconv"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestMiniWechatOnlineQualAddSuccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	data := map[string]interface{}{
		"loc":           "科室名称",
		"hospital":      "医院名称",
		"postion":       "职位名称",
		"qual_img":      "00/cc1513a9d57701c86848159f6ed5c9.png",
		"expertise_img": "00/cc1513a9d57701c86848159f6ed5c9.png",
		"loc_img":       "00/cc1513a9d57701c86848159f6ed5c9.png",
	}

	obj := e.POST("/online/v1/qual/add").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(200)
	obj.Value("message").String().Equal("新增成功")
	obj.Value("data").Object().Value("id").NotNull()
	obj.Value("data").Object().Value("loc").Equal("科室名称")
	obj.Value("data").Object().Value("hospital").Equal("医院名称")
	obj.Value("data").Object().Value("postion").Equal("职位名称")
	obj.Value("data").Object().Value("qual_img").Equal("00/cc1513a9d57701c86848159f6ed5c9.png")
	obj.Value("data").Object().Value("expertise_img").Equal("00/cc1513a9d57701c86848159f6ed5c9.png")
	obj.Value("data").Object().Value("loc_img").Equal("00/cc1513a9d57701c86848159f6ed5c9.png")
}

func TestMiniWechatOnlineQualAddNoLocError(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	data := map[string]interface{}{
		"loc":           "",
		"hospital":      "医院名称",
		"postion":       "职位名称",
		"qual_img":      "00/cc1513a9d57701c86848159f6ed5c9.png",
		"expertise_img": "00/cc1513a9d57701c86848159f6ed5c9.png",
		"loc_img":       "00/cc1513a9d57701c86848159f6ed5c9.png",
	}

	obj := e.POST("/online/v1/qual/add").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("科室不能为空！")
}

func TestMiniWechatOnlineQualAddNoHospitalError(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	data := map[string]interface{}{
		"loc":           "医院名称",
		"hospital":      "",
		"postion":       "职位名称",
		"qual_img":      "00/cc1513a9d57701c86848159f6ed5c9.png",
		"expertise_img": "00/cc1513a9d57701c86848159f6ed5c9.png",
		"loc_img":       "00/cc1513a9d57701c86848159f6ed5c9.png",
	}

	obj := e.POST("/online/v1/qual/add").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("医院不能为空！")
}

func TestMiniWechatOnlineQualAddNoQualImgError(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	data := map[string]interface{}{
		"loc":           "医院名称",
		"hospital":      "职位名称",
		"postion":       "职位名称",
		"qual_img":      "",
		"expertise_img": "00/cc1513a9d57701c86848159f6ed5c9.png",
		"loc_img":       "00/cc1513a9d57701c86848159f6ed5c9.png",
	}

	obj := e.POST("/online/v1/qual/add").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("执业证书图片不能为空！")
}

func TestMiniWechatOnlineQualAddNoPostionError(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	data := map[string]interface{}{
		"loc":           "医院名称",
		"hospital":      "职位名称",
		"postion":       "",
		"qual_img":      "00/cc1513a9d57701c86848159f6ed5c9.png",
		"expertise_img": "00/cc1513a9d57701c86848159f6ed5c9.png",
		"loc_img":       "00/cc1513a9d57701c86848159f6ed5c9.png",
	}

	obj := e.POST("/online/v1/qual/add").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("职称不能为空！")
}

func TestMiniWechatOnlineQualAddNoExpertiseImgError(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	data := map[string]interface{}{
		"loc":           "医院名称",
		"hospital":      "职位名称",
		"postion":       "职位名称",
		"qual_img":      "00/cc1513a9d57701c86848159f6ed5c9.png",
		"expertise_img": "",
		"loc_img":       "00/cc1513a9d57701c86848159f6ed5c9.png",
	}

	obj := e.POST("/online/v1/qual/add").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("专业技术资格证图片不能为空！")
}

func TestMiniWechatOnlineQualAddNoLocImgError(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})

	data := map[string]interface{}{
		"loc":           "医院名称",
		"hospital":      "职位名称",
		"postion":       "职位名称",
		"qual_img":      "00/cc1513a9d57701c86848159f6ed5c9.png",
		"expertise_img": "00/cc1513a9d57701c86848159f6ed5c9.png",
		"loc_img":       "",
	}

	obj := e.POST("/online/v1/qual/add").
		WithHeaders(map[string]string{"X-Token": MiniWechatToken, "IsDev": "1", "AuthType": strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10)}).
		WithCookie("PHPSESSID", MINIWECHATPHPSESSID).
		WithJSON(data).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "message")
	obj.Value("code").Equal(400)
	obj.Value("message").String().Equal("职称证图片不能为空！")
}
