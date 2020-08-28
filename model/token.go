package model

import (
	"encoding/json"
	"fmt"
	"github.com/snowlyg/ChindeoTest/config"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type token struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"AccessToken"`
	} `json:"data"`
}

var Token string
var MiniWechatToken string

var PHPSESSID string
var MINIWECHATPHPSESSID string

func GetAuthToken() {
	data := fmt.Sprintf("app_id=%s&app_secret=%s", "b44fc017043763eb5ac15f0069d77c", "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b")
	Token, PHPSESSID = getToken(strconv.FormatInt(int64(AuthTypeServer), 10), data, false)

	data = fmt.Sprintf("uuid=%s&code=%s&iv=%s&encrypted_data=%s", "5205857593c2eacc6f6c1da376b32ca3", "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b", "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b", "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b")
	MiniWechatToken, MINIWECHATPHPSESSID = getToken(strconv.FormatInt(int64(AuthTypeMiniWechat), 10), data, true)
}

func ClearToken() {
	Token, PHPSESSID = "", ""
	MiniWechatToken, MINIWECHATPHPSESSID = "", ""
}

func getToken(authType, data string, isDev bool) (string, string) {
	var re token
	url := fmt.Sprintf("%s/%s", config.Config.Url, "api/v1/get_access_token")

	result, session := DoPOST(url, data, authType, isDev)
	err := json.Unmarshal(result, &re)
	if err != nil {
		log.Printf("GetToken error：%v -result:%v", err, result)
	}
	if re.Code == 200 {
		return re.Data.Token, session
	} else {
		println(fmt.Sprintf("token 获取失败 : %s", re.Message))
	}
	return "", ""
}

func DoPOST(url, data, authType string, isDev bool) ([]byte, string) {
	var session string
	client, req := getClient("POST", url, data)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	if isDev {
		req.Header.Set("IsDev", "1")
	}
	req.Header.Set("AuthType", authType)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("请求出错：%v", err)
	}
	if resp != nil {

		for _, c := range resp.Cookies() {
			if c.Name == "PHPSESSID" {
				session = c.Value
			}
		}

		defer resp.Body.Close()
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("DoPOST error：%v", err)
		}
		return result, session
	}

	return nil, ""
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
