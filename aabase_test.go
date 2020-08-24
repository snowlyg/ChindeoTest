package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/snowlyg/ChindeoTest/common"
	"github.com/snowlyg/ChindeoTest/config"
	"github.com/snowlyg/ChindeoTest/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

const AppId = 13
const IdCardNo = "430923198901156623"

var Token string
var PHPSESSID string
var MenuType *model.APIMenuTypes
var MenuTag *model.APIMenuTags
var Menu *model.APIMenus
var User *model.APIUsers
var Order *model.APIOOrders
var OrderAddr *model.APIOOrderAddrs
var OrderMenus []*model.APIOOrderMenus

type getToken struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"AccessToken"`
	} `json:"data"`
}

//单元测试基境
func TestMain(m *testing.M) {

	println(fmt.Sprintf("-------------------main start-------------------"))

	db := model.Init()

	GetUserByIdCardNo(db)
	CreateMenuType(db)
	CreateMenuTag(db)
	CreateMenu(db)
	CreateOrder(db)

	var re getToken
	appid := "b44fc017043763eb5ac15f0069d77c"
	appsecret := "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b"
	result := DoPOST(fmt.Sprintf("%s/%s", config.Config.Url, "api/v1/get_access_token"), fmt.Sprintf("app_id=%s&app_secret=%s&auth_type=%d", appid, appsecret, 4))
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

	model.ClearTables(db)
	model.Close(db)

	println(fmt.Sprintf("-------------------main end-------------------"))

	os.Exit(exitCode)
}

func GetUserByIdCardNo(db *gorm.DB) {
	user := model.APIUsers{}
	if err := db.Where("id_card_no = ?", IdCardNo).First(&user).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuType create error :%v", err))
	}
	User = &user
}

func CreateMenuType(db *gorm.DB) {
	menuType := model.APIMenuTypes{Name: "菜单类型", ApplicationID: AppId, CreateAt: time.Now(), UpdateAt: time.Now(), IsDeleted: sql.NullTime{}}
	if err := db.Create(&menuType).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuType create error :%v", err))
	}
	MenuType = &menuType
}

func CreateMenuTag(db *gorm.DB) {
	menuTag := model.APIMenuTags{Name: "菜单标签", CreateAt: time.Now(), UpdateAt: time.Now(), IsDeleted: sql.NullTime{}}
	if err := db.Create(&menuTag).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuTag create error :%v", err))
	}
	MenuTag = &menuTag
}

func CreateMenu(db *gorm.DB) {
	menu := model.APIMenus{
		Name:          "菜单名称",
		TimeType:      common.MENU_TIME_TYPE_B,
		Desc:          "菜品介绍",
		Status:        true,
		Amount:        0,
		Price:         10.00,
		ApplicationID: AppId,
		MenuTypeID:    MenuType.ID,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}

	if err := db.Create(&menu).Error; err != nil {
		fmt.Println(fmt.Sprintf("menu create error :%v", err))
	}

	Menu = &menu

	tagMenu := model.MenuMenuTag{MenuID: menu.ID, MenuTagID: MenuTag.ID, UpdateAt: time.Now(), CreateAt: time.Now()}
	if err := db.Table("menu_menu_tag").Create(&tagMenu).Error; err != nil {
		fmt.Println(fmt.Sprintf("tagMenu create error :%v", err))
	}
}

func CreateOrder(db *gorm.DB) {
	order := model.APIOOrders{
		OrderNo:       "O202008241612348468756914",
		Status:        common.I_ORDER_STATUS_FOR_DELIVERY,
		Amount:        0,
		Total:         10.00,
		ApplicationID: AppId,
		PayType:       common.I_ORDER_PAY_TYPE_WECHAT,
		Rmk:           "备注",
		AppType:       common.ORDER_APP_TYPE_MINI,
		OpenID:        "",
		TradeType:     "",
		IsReturn:      false,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
		UserID:        User.ID,
	}
	if err := db.Create(&order).Error; err != nil {
		fmt.Println(fmt.Sprintf("order create error :%v", err))
	}

	orderAddr := model.APIOOrderAddrs{
		Name:         "name",
		Phone:        "13412569874",
		Addr:         "",
		Sex:          1,
		OOrderID:     order.ID,
		HospitalName: "HospitalName",
		LocName:      "LocName",
		BedNum:       "BedNum",
		HospitalNo:   "9556854545",
		Disease:      "Disease",
		Age:          10,
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
		IsDeleted:    sql.NullTime{},
	}
	if err := db.Create(&orderAddr).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderAddr create error :%v", err))
	}

	orderMenu := model.APIOOrderMenus{
		MenuName:     Menu.Name,
		MenuType:     MenuType.Name,
		MenuTimeType: Menu.TimeType.String(),
		MenuDesc:     Menu.Desc,
		MenuID:       Menu.ID,
		Price:        Menu.Price,
		Amount:       Menu.Amount,
		Cover:        Menu.Cover,
		OOrderID:     order.ID,
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
		IsDeleted:    sql.NullTime{},
	}
	if err := db.Create(&orderMenu).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderAddr create error :%v", err))
	}
	var orderMenus []*model.APIOOrderMenus
	orderMenus = append(orderMenus, &orderMenu)

	OrderMenus = orderMenus
	OrderAddr = &orderAddr
	Order = &order

}

func DoPOST(url string, data string) []byte {
	client, req := getClient("POST", url, data)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Set("AuthType", "4")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("请求出错：%v", err)
	}
	if resp != nil {

		for _, c := range resp.Cookies() {
			if c.Name == "PHPSESSID" {
				PHPSESSID = c.Value
			}
		}

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
