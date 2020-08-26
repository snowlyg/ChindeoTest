package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"github.com/snowlyg/ChindeoTest/common"
	"github.com/snowlyg/ChindeoTest/config"
	"github.com/snowlyg/ChindeoTest/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

const AppId = 13
const IdCardNo = "430923198901156623"
const CareMiniPrice = 1.00
const CareMaxPrice = 50.00
const CarerPrice = 50.00

var Token string
var MiniWechatToken string

var PHPSESSID string
var MINIWECHATPHPSESSID string

var MenuAmount int
var Addr *model.APIAddrs
var MenuType *model.APIMenuTypes
var MenuTag *model.APIMenuTags
var Menu *model.APIMenus
var User *model.APIUsers
var Order *model.APIOOrders
var OrderAddr *model.APIOOrderAddrs
var OrderMenus []*model.APIOOrderMenus

var CareAmount int
var CareType *model.CareTypes
var CareTag *model.CareTags
var Care *model.Cares
var CareOrder *model.CareOrders
var CareOrderAddr *model.CareOrderAddrs
var CareOrderInfo *model.CareOrderInfos

var CarerAmount int
var CarerTag *model.CarerTags
var Carer *model.Carers
var CarerOrder *model.CareOrders
var CarerOrderAddr *model.CareOrderAddrs
var CarerOrderCarerInfo *model.CareOrderCarerInfos

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

	CreateAddr(db)

	CreateMenuType(db)
	CreateMenuTag(db)
	CreateMenu(db)
	CreateOrder(db)

	CreateCareTag(db)
	CreateCareType(db)
	CreateCare(db)
	CreateCareOrder(db)

	CreateCarerTag(db)
	CreateCarer(db)
	CreateCarerOrder(db)

	data := fmt.Sprintf("app_id=%s&app_secret=%s", "b44fc017043763eb5ac15f0069d77c", "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b")
	Token, PHPSESSID = getAuthToken(strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10), data, false)

	data = fmt.Sprintf("uuid=%s&code=%s&iv=%s&encrypted_data=%s", "5205857593c2eacc6f6c1da376b32ca3", "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b", "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b", "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b")
	MiniWechatToken, MINIWECHATPHPSESSID = getAuthToken(strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10), data, true)

	flag.Parse()
	exitCode := m.Run()

	Token = ""

	model.ClearTables(db)
	model.Close(db)

	println(fmt.Sprintf("-------------------main end-------------------"))

	os.Exit(exitCode)
}

func getAuthToken(authType, data string, isDev bool) (string, string) {
	var re getToken
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

func GetUserByIdCardNo(db *gorm.DB) {
	user := model.APIUsers{}
	if err := db.Where("id_card_no = ?", IdCardNo).First(&user).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuType create error :%v", err))
	}
	User = &user
}

func CreateAddr(db *gorm.DB) {
	addr := model.APIAddrs{
		Name:         "name",
		Phone:        "13412569874",
		Addr:         "",
		Sex:          1,
		UserID:       User.ID,
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
	if err := db.Create(&addr).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuType create error :%v", err))
	}
	Addr = &addr
}

func CreateCareType(db *gorm.DB) {
	careType := model.CareTypes{Name: "护理项目类型", EnName: "en_name", Status: true, CreateAt: time.Now(), UpdateAt: time.Now(), IsDeleted: sql.NullTime{}, Icon: "icon"}
	if err := db.Create(&careType).Error; err != nil {
		fmt.Println(fmt.Sprintf("careType create error :%v", err))
	}
	CareType = &careType
}

func CreateCareTag(db *gorm.DB) {
	careTag := model.CareTags{Name: "护理项目标签", CreateAt: time.Now(), UpdateAt: time.Now(), IsDeleted: sql.NullTime{}, Icon: "icon"}
	if err := db.Create(&careTag).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuTag create error :%v", err))
	}
	CareTag = &careTag
}

func CreateCare(db *gorm.DB) {
	care := model.Cares{
		Name:          "护理项目名称",
		Desc:          "护理项目名称",
		TimeType:      "时",
		Status:        true,
		Amount:        CareAmount,
		MinPrice:      CareMiniPrice,
		MaxPrice:      CareMaxPrice,
		Cover:         "sdfsdfds",
		CareDetail:    "dsfsdfdsf",
		CareTypeID:    CareType.ID,
		ApplicationID: AppId,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}
	if err := db.Create(&care).Error; err != nil {
		fmt.Println(fmt.Sprintf("care create error :%v", err))
	}
	Care = &care

	tagCare := model.CareCareTag{CareID: Care.ID, CareTagID: CareTag.ID, UpdateAt: time.Now(), CreateAt: time.Now()}
	if err := db.Table("care_care_tag").Create(&tagCare).Error; err != nil {
		fmt.Println(fmt.Sprintf("tagCare create error :%v", err))
	}
}

func CreateCareOrder(db *gorm.DB) {
	tt := time.Now()
	var dd decimal.Decimal
	sub := int(tt.AddDate(0, 0, 1).Sub(tt).Hours())
	maxCarePrice := decimal.NewFromFloat(CareMaxPrice)
	if Care.TimeType == "天" {
		dd = maxCarePrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if Care.TimeType == "时" {
		dd = maxCarePrice.Mul(decimal.NewFromFloat(float64(sub)))
	}
	total, _ := dd.Float64()

	order := model.CareOrders{
		OrderNo:       "C202008241612348468756914",
		Status:        common.I_ORDER_STATUS_FOR_DELIVERY,
		Total:         total,
		ApplicationID: AppId,
		PayType:       common.I_ORDER_PAY_TYPE_WECHAT,
		Rmk:           "备注",
		AppType:       common.ORDER_APP_TYPE_MINI,
		StartAt:       tt,
		EndAt:         tt.AddDate(0, 0, 1),
		OpenID:        "",
		TradeType:     "",
		IsReturn:      false,
		CreateAt:      tt,
		UpdateAt:      tt,
		IsDeleted:     sql.NullTime{},
		UserID:        User.ID,
		CarerID:       0,
	}
	if err := db.Create(&order).Error; err != nil {
		fmt.Println(fmt.Sprintf("order create error :%v", err))
	}

	orderAddr := model.CareOrderAddrs{
		Name:         "name",
		Phone:        "13412569874",
		Addr:         "",
		Sex:          1,
		CareOrderID:  order.ID,
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

	orderInfo := model.CareOrderInfos{
		Name:            Care.Name,
		Desc:            Care.Desc,
		TimeType:        Care.TimeType,
		CareType:        CareType.Name,
		CareTags:        CareTag.Name,
		MinPrice:        Care.MinPrice,
		MaxPrice:        Care.MaxPrice,
		CareDetail:      Care.CareDetail,
		ApplicationName: "name",
		Cover:           Care.Cover,
		CareOrderID:     order.ID,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := db.Create(&orderInfo).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderInfo create error :%v", err))
	}

	CareOrderInfo = &orderInfo
	CareOrderAddr = &orderAddr
	CareOrder = &order
}

func CreateCarerTag(db *gorm.DB) {
	carerTag := model.CarerTags{Name: "护工标签", CreateAt: time.Now(), UpdateAt: time.Now(), IsDeleted: sql.NullTime{}, Icon: "icon"}
	if err := db.Create(&carerTag).Error; err != nil {
		fmt.Println(fmt.Sprintf("carerTag create error :%v", err))
	}
	CarerTag = &carerTag
}

func CreateCarer(db *gorm.DB) {
	carer := model.Carers{
		Name:          "护理项目名称",
		Desc:          "护理项目名称",
		TimeType:      "时",
		Status:        true,
		Amount:        CareAmount,
		Price:         CarerPrice,
		Avatar:        "dsfsdfs",
		CarerDetail:   "dfsdfsf",
		ApplicationID: AppId,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}
	if err := db.Create(&carer).Error; err != nil {
		fmt.Println(fmt.Sprintf("carer create error :%v", err))
	}
	Carer = &carer

	tagCarer := model.CarerCarerTag{CarerID: Care.ID, CarerTagID: CarerTag.ID, UpdateAt: time.Now(), CreateAt: time.Now()}
	if err := db.Table("carer_carer_tag").Create(&tagCarer).Error; err != nil {
		fmt.Println(fmt.Sprintf("tagCarer create error :%v", err))
	}
}

func CreateCarerOrder(db *gorm.DB) {
	tt := time.Now()
	var dd decimal.Decimal
	sub := int(tt.AddDate(0, 0, 1).Sub(tt).Hours())
	carerPrice := decimal.NewFromFloat(CarerPrice)
	if Care.TimeType == "天" {
		dd = carerPrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if Care.TimeType == "时" {
		dd = carerPrice.Mul(decimal.NewFromFloat(float64(sub)))
	}
	total, _ := dd.Float64()

	order := model.CareOrders{
		OrderNo:       "C202008241612348468756914",
		Status:        common.I_ORDER_STATUS_FOR_DELIVERY,
		Total:         total,
		ApplicationID: AppId,
		PayType:       common.I_ORDER_PAY_TYPE_WECHAT,
		Rmk:           "备注",
		AppType:       common.ORDER_APP_TYPE_MINI,
		StartAt:       tt,
		EndAt:         tt.AddDate(0, 0, 1),
		OpenID:        "",
		TradeType:     "",
		IsReturn:      false,
		CreateAt:      tt,
		UpdateAt:      tt,
		IsDeleted:     sql.NullTime{},
		UserID:        User.ID,
		CarerID:       0,
	}
	if err := db.Create(&order).Error; err != nil {
		fmt.Println(fmt.Sprintf("order create error :%v", err))
	}

	orderAddr := model.CareOrderAddrs{
		Name:         "name",
		Phone:        "13412569874",
		Addr:         "",
		Sex:          1,
		CareOrderID:  order.ID,
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

	orderCarerInfo := model.CareOrderCarerInfos{
		Name:            Carer.Name,
		Desc:            Carer.Desc,
		TimeType:        Carer.TimeType,
		CarerTags:       CarerTag.Name,
		Price:           Carer.Price,
		Age:             Carer.Age,
		WorkExp:         Carer.WorkExp,
		Sex:             Carer.Sex,
		Phone:           Carer.Phone,
		CarerDetail:     Carer.CarerDetail,
		ApplicationName: "name",
		Avatar:          Carer.Avatar,
		CareOrderID:     order.ID,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := db.Create(&orderCarerInfo).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderInfo create error :%v", err))
	}

	CarerOrderCarerInfo = &orderCarerInfo
	CarerOrderAddr = &orderAddr
	CarerOrder = &order
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
		Amount:        MenuAmount,
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
