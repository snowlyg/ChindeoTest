package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

const BaseUrl = "https://diancan.t.chindeo.com"

var Token string

type getToken struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"AccessToken"`
	} `json:"data"`
}

//单元测试基境
func TestMain(m *testing.M) {
	var re getToken

	appid := "b44fc017043763eb5ac15f0069d77c"
	appsecret := "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b"
	result := DoPOST(fmt.Sprintf("%s/%s", BaseUrl, "api/v1/get_access_token"), fmt.Sprintf("app_id=%s&app_secret=%s&auth_type=%d", appid, appsecret, 4))
	err := json.Unmarshal(result, &re)
	if err != nil {
		log.Printf("GetToken error：%v -result:%v", err, result)
	}
	if re.Code == 200 {
		Token = re.Data.Token
	} else {
		println(fmt.Sprintf("token 获取失败 : %s", re.Message))
	}

	flag.Parse()
	exitCode := m.Run()

	Token = ""

	os.Exit(exitCode)
}

func DoPOST(url string, data string) []byte {
	client, req := getClient("POST", url, data)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("请求出错：%v", err)
	}
	if resp != nil {
		defer resp.Body.Close()
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("DoPOST error：%v", err)
		}
		return result
	}

	return nil
}

func getClient(reType string, url string, data string) (*http.Client, *http.Request) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(reType, url, strings.NewReader(data))
	if err != nil {
		log.Printf("请求出错：%v", err)
	}
	return client, req
}
