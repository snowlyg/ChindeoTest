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
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

var Db *gorm.DB

var Token string
var MiniWechatToken string
var PHPSESSID string
var MINIWECHATPHPSESSID string

var Addr *model.APIAddrs
var MenuType *model.APIMenuTypes
var MenuTag *model.APIMenuTags
var Menu *model.APIMenus

var OrderCount int
var MiniOrderCommentCount int
var Order *model.APIOOrders
var OrderAddr *model.APIOOrderAddrs
var OrderMenus []*model.APIOOrderMenus
var OrderComment *model.Comments
var MiniOrder *model.APIOOrders
var MiniOrderAddr *model.APIOOrderAddrs
var MiniOrderMenus []*model.APIOOrderMenus
var MiniOrderComment *model.Comments
var MiniOrderCommentForDel *model.Comments

var CareAmount int
var CareType *model.CareTypes
var CareTag *model.CareTags
var Care *model.Cares

var CareOrderCount int
var CareOrderCommentCount int
var CareOrder *model.CareOrders
var CareOrderAddr *model.CareOrderAddrs
var CareOrderInfo *model.CareOrderInfos
var CareOrderComment *model.Comments
var CareOrderCommentForDel *model.Comments

var MiniCareOrder *model.CareOrders
var MiniCareOrderAddr *model.CareOrderAddrs
var MiniCareOrderInfo *model.CareOrderInfos
var MiniCareOrderComment *model.Comments

var CarerAmount int
var CarerTag *model.CarerTags
var Carer *model.Carers
var CarerOrder *model.CareOrders
var CarerOrderAddr *model.CareOrderAddrs
var CarerOrderCarerInfo *model.CareOrderCarerInfos
var CarerOrderComment *model.Comments

var MiniCarerOrder *model.CareOrders
var MiniCarerOrderComment *model.Comments
var MiniCarerOrderAddr *model.CareOrderAddrs
var MiniCarerOrderCarerInfo *model.CareOrderCarerInfos

var Article *model.Articles
var LocType *model.LocTypes
var Loc *model.Locs
var UserType *model.UserTypes
var Patient *model.Patients

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

	common.GetUserByIdCardNo(db)
	common.GetFake()
	common.GetPics()

	Addr = common.CreateAddr(db)
	MenuType = common.CreateMenuType(db)
	MenuTag = common.CreateMenuTag(db)
	common.CreateMenu(db)
	common.CreateOrder(db)
	common.CreateMiniOrder(db)

	common.CreateCareTag(db)
	common.CreateCareType(db)
	common.CreateCare(db)
	common.CreateMiniCareOrder(db)
	common.CreateCareOrder(db)

	common.CreateCarerTag(db)
	common.CreateCarer(db)
	common.CreateMiniCarerOrder(db)
	common.CreateCarerOrder(db)

	common.CreateArticle(db)
	common.CreateLocType(db)
	common.CreateLoc(db)
	common.CreateUserType(db)
	common.CreatePatient(db)

	data := fmt.Sprintf("app_id=%s&app_secret=%s", "b44fc017043763eb5ac15f0069d77c", "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b")
	Token, PHPSESSID = getAuthToken(strconv.FormatInt(int64(common.AUTH_TYPE_SERVER), 10), data, false)

	data = fmt.Sprintf("uuid=%s&code=%s&iv=%s&encrypted_data=%s", "5205857593c2eacc6f6c1da376b32ca3", "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b", "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b", "106d1b47f6fa30c0ff6ae48da5f1c9e4b557a6363ed854e2e250de4e00127c2b")
	MiniWechatToken, MINIWECHATPHPSESSID = getAuthToken(strconv.FormatInt(int64(common.AUTH_TYPE_MINIWECHAT), 10), data, true)

	flag.Parse()
	exitCode := m.Run()

	Token = ""

	model.ClearTables(Db)
	model.Close(Db)

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

func CreateCareType() {
	careType := model.CareTypes{Name: "护理项目类型", EnName: "en_name", Status: true, CreateAt: time.Now(), UpdateAt: time.Now(), IsDeleted: sql.NullTime{}, Icon: "icon"}
	if err := Db.Create(&careType).Error; err != nil {
		fmt.Println(fmt.Sprintf("careType create error :%v", err))
	}
	CareType = &careType
}

func CreateCareTag() {
	careTag := model.CareTags{Name: "护理项目标签", CreateAt: time.Now(), UpdateAt: time.Now(), IsDeleted: sql.NullTime{}, Icon: "icon"}
	if err := Db.Create(&careTag).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuTag create error :%v", err))
	}
	CareTag = &careTag
}

func CreateCare() {
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

	if err := Db.Create(&care).Error; err != nil {
		fmt.Println(fmt.Sprintf("care create error :%v", err))
	}
	Care = &care

	care_no_tag := model.Cares{
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

	if err := Db.Create(&care_no_tag).Error; err != nil {
		fmt.Println(fmt.Sprintf("care_no_tag create error :%v", err))
	}

	tagCare := model.CareCareTag{CareID: Care.ID, CareTagID: CareTag.ID, UpdateAt: time.Now(), CreateAt: time.Now()}
	if err := Db.Table("care_care_tag").Create(&tagCare).Error; err != nil {
		fmt.Println(fmt.Sprintf("tagCare create error :%v", err))
	}
}

func CreateCareOrder() {
	CareOrderCount = CareOrderCount + 2
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
	if err := Db.Create(&order).Error; err != nil {
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
	if err := Db.Create(&orderAddr).Error; err != nil {
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
	if err := Db.Create(&orderInfo).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderInfo create error :%v", err))
	}
	orderComment := model.Comments{
		Content:         Fake.Paragraph(1, true),
		UserID:          User.ID,
		CommentableID:   order.ID,
		CommentableType: "app\\care\\model\\CareOrder",
		ApplicationID:   AppId,
		Pics:            Pics,
		Star:            5,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := Db.Create(&orderComment).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderComment create error :%v", err))
	}

	orderForCommentDel := model.CareOrders{
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
	if err := Db.Create(&orderForCommentDel).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderForCommentDel create error :%v", err))
	}

	CareOrderCommentForDel = &model.Comments{
		Content:         Fake.Paragraph(1, true),
		UserID:          User.ID,
		CommentableID:   orderForCommentDel.ID,
		CommentableType: "app\\care\\model\\CareOrder",
		ApplicationID:   AppId,
		Pics:            Pics,
		Star:            5,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := Db.Create(CareOrderCommentForDel).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderCommentForDel create error :%v", err))
	}

	CareOrderInfo = &orderInfo
	CareOrderAddr = &orderAddr
	CareOrder = &order
	CareOrderComment = &orderComment
	CareOrderCommentCount++
}

func CreateMiniCareOrder() {
	CareOrderCount++
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
	if err := Db.Create(&order).Error; err != nil {
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
	if err := Db.Create(&orderAddr).Error; err != nil {
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
	if err := Db.Create(&orderInfo).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderInfo create error :%v", err))
	}
	orderComment := model.Comments{
		Content:         Fake.Paragraph(1, true),
		UserID:          User.ID,
		CommentableID:   order.ID,
		CommentableType: "app\\care\\model\\CareOrder",
		ApplicationID:   AppId,
		Pics:            Pics,
		Star:            5,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := Db.Create(&orderComment).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderComment create error :%v", err))
	}

	MiniCareOrderInfo = &orderInfo
	MiniCareOrderAddr = &orderAddr
	MiniCareOrder = &order
	MiniCareOrderComment = &orderComment
}

func CreateCarerTag() {
	carerTag := model.CarerTags{Name: "护工标签", CreateAt: time.Now(), UpdateAt: time.Now(), IsDeleted: sql.NullTime{}, Icon: "icon"}
	if err := Db.Create(&carerTag).Error; err != nil {
		fmt.Println(fmt.Sprintf("carerTag create error :%v", err))
	}
	CarerTag = &carerTag
}

func CreateCarer() {
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
	if err := Db.Create(&carer).Error; err != nil {
		fmt.Println(fmt.Sprintf("carer create error :%v", err))
	}
	Carer = &carer

	carer_no_tag := model.Carers{
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
	if err := Db.Create(&carer_no_tag).Error; err != nil {
		fmt.Println(fmt.Sprintf("carer_no_tag create error :%v", err))
	}

	tagCarer := model.CarerCarerTag{CarerID: Carer.ID, CarerTagID: CarerTag.ID, UpdateAt: time.Now(), CreateAt: time.Now()}
	if err := Db.Table("carer_carer_tag").Create(&tagCarer).Error; err != nil {
		fmt.Println(fmt.Sprintf("tagCarer create error :%v", err))
	}
}

func CreateCarerOrder() {
	CareOrderCount++
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
	if err := Db.Create(&order).Error; err != nil {
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
	if err := Db.Create(&orderAddr).Error; err != nil {
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
	if err := Db.Create(&orderCarerInfo).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderInfo create error :%v", err))
	}

	orderComment := model.Comments{
		Content:         Fake.Paragraph(1, true),
		UserID:          User.ID,
		CommentableID:   order.ID,
		CommentableType: "app\\care\\model\\CareOrder",
		ApplicationID:   AppId,
		Pics:            Pics,
		Star:            5,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := Db.Create(&orderComment).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderComment create error :%v", err))
	}

	CarerOrderCarerInfo = &orderCarerInfo
	CarerOrderAddr = &orderAddr
	CarerOrder = &order
	CarerOrderComment = &orderComment
}

func CreateMiniCarerOrder() {
	CareOrderCount++
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
	if err := Db.Create(&order).Error; err != nil {
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
	if err := Db.Create(&orderAddr).Error; err != nil {
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
	if err := Db.Create(&orderCarerInfo).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderInfo create error :%v", err))
	}
	orderComment := model.Comments{
		Content:         Fake.Paragraph(1, true),
		UserID:          User.ID,
		CommentableID:   order.ID,
		CommentableType: "app\\care\\model\\CareOrder",
		ApplicationID:   AppId,
		Pics:            Pics,
		Star:            5,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := Db.Create(&orderComment).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderComment create error :%v", err))
	}

	MiniCarerOrderCarerInfo = &orderCarerInfo
	MiniCarerOrderAddr = &orderAddr
	MiniCarerOrder = &order
	MiniCarerOrderComment = &orderComment
}

func CreateOrder() {
	OrderCount++
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
	if err := Db.Create(&order).Error; err != nil {
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
	if err := Db.Create(&orderAddr).Error; err != nil {
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
	if err := Db.Create(&orderMenu).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderAddr create error :%v", err))
	}
	var orderMenus []*model.APIOOrderMenus
	orderMenus = append(orderMenus, &orderMenu)

	orderComment := model.Comments{
		Content:         Fake.Paragraph(1, true),
		UserID:          User.ID,
		CommentableID:   order.ID,
		CommentableType: "app\\api\\model\\Order",
		ApplicationID:   AppId,
		Pics:            Pics,
		Star:            5,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := Db.Create(&orderComment).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderComment create error :%v", err))
	}

	OrderMenus = orderMenus
	OrderAddr = &orderAddr
	OrderComment = &orderComment
	Order = &order
}

func CreateMiniOrder() {
	OrderCount = OrderCount + 2
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
	if err := Db.Create(&order).Error; err != nil {
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
	if err := Db.Create(&orderAddr).Error; err != nil {
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
	if err := Db.Create(&orderMenu).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderAddr create error :%v", err))
	}
	var orderMenus []*model.APIOOrderMenus
	orderMenus = append(orderMenus, &orderMenu)

	orderComment := model.Comments{
		Content:         Fake.Paragraph(1, true),
		UserID:          User.ID,
		CommentableID:   order.ID,
		CommentableType: "app\\api\\model\\Order",
		ApplicationID:   AppId,
		Pics:            Pics,
		Star:            5,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := Db.Create(&orderComment).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderComment create error :%v", err))
	}

	orderForCommentDel := model.APIOOrders{
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
	if err := Db.Create(&orderForCommentDel).Error; err != nil {
		fmt.Println(fmt.Sprintf("order create error :%v", err))
	}

	MiniOrderCommentForDel = &model.Comments{
		Content:         Fake.Paragraph(1, true),
		UserID:          User.ID,
		CommentableID:   orderForCommentDel.ID,
		CommentableType: "app\\api\\model\\Order",
		ApplicationID:   AppId,
		Pics:            Pics,
		Star:            5,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := Db.Create(MiniOrderCommentForDel).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderCommentFordel create error :%v", err))
	}

	MiniOrderMenus = orderMenus
	MiniOrderAddr = &orderAddr
	MiniOrder = &order
	MiniOrderComment = &orderComment
	MiniOrderCommentCount = 1
}

func CreateArticle() {

	artice := model.Articles{
		Title:            Fake.JobTitle(),
		Digest:           Fake.Paragraph(1, true),
		Content:          Fake.Paragraph(1, true),
		Author:           Fake.Name(),
		LocalURL:         Fake.URL(),
		ContentSourceURL: Fake.URL(),
		CreateAt:         time.Now(),
		UpdateAt:         time.Now(),
		IsDeleted:        sql.NullTime{},
	}
	if err := Db.Create(&artice).Error; err != nil {
		fmt.Println(fmt.Sprintf("artice create error :%v", err))
	}

	Article = &artice
}

func CreateLocType() {
	locType := model.LocTypes{
		Name:          Fake.JobTitle(),
		ApplicationID: AppId,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}
	if err := Db.Create(&locType).Error; err != nil {
		fmt.Println(fmt.Sprintf("locType create error :%v", err))
	}

	LocType = &locType
}

func CreateLoc() {
	loc := model.Locs{
		LocID:         rand.Int(),
		CtHospitalID:  rand.Int(),
		LocWardFlag:   rand.Int(),
		LocActiveFlag: rand.Int(),
		LocDesc:       Fake.Paragraph(1, true),
		ApplicationID: AppId,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}
	if err := Db.Create(&loc).Error; err != nil {
		fmt.Println(fmt.Sprintf("loc create error :%v", err))
	}

	Loc = &loc
}

func CreateUserType() {
	letters := []byte("NDO")
	userType := model.UserTypes{
		UtpID:       rand.Int(),
		UtpCode:     Fake.CellPhoneNumber(),
		UtpDesc:     Fake.Paragraph(1, true),
		UtpType:     string(letters[rand.Intn(2)]),
		UtpActive:   true,
		UtpContrast: Fake.Paragraph(1, true),
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
		IsDeleted:   sql.NullTime{},
	}
	if err := Db.Create(&userType).Error; err != nil {
		fmt.Println(fmt.Sprintf("userType create error :%v", err))
	}

	UserType = &userType
}

func CreatePatient() {
	patient := model.Patients{
		Username:  Fake.Name(),
		Nickname:  Fake.Name(),
		Phone:     Fake.PhoneNumber(),
		Email:     Fake.Email(),
		Sex:       rand.Intn(1),
		Password:  Fake.PhoneNumber(),
		Status:    true,
		AvatarURL: Fake.URL(),
		OpenID:    Fake.PhoneNumber(),
		UnionID:   Fake.PhoneNumber(),
		Country:   Fake.Country(),
		Province:  Fake.Country(),
		City:      Fake.City(),
		Mac:       Fake.City(),
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		IsDeleted: sql.NullTime{},
	}
	if err := Db.Create(&patient).Error; err != nil {
		fmt.Println(fmt.Sprintf("patient create error :%v", err))
	}

	Patient = &patient

	userPatient := model.UserPatient{PatientID: Patient.ID, UserID: User.ID, UpdateAt: time.Now(), CreateAt: time.Now()}
	if err := Db.Table("user_patient").Create(&userPatient).Error; err != nil {
		fmt.Println(fmt.Sprintf("userPatient create error :%v", err))
	}
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
