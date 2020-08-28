package model

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/snowlyg/ChindeoTest/config"
	"net/http"
	"testing"
)

func GetE(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Jar: httpexpect.NewJar(), // used by default if Client is nil
		},
		BaseURL: config.Config.Url,
	})
}
