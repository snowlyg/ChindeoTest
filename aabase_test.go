package main

import (
	"flag"
	"fmt"
	"github.com/snowlyg/ChindeoTest/model"
	"os"
	"testing"
)

const COMMENTABLETYPE = "app\\api\\model\\Order"
const CARECOMMENTABLETYPE = "app\\care\\model\\CareOrder"

var App *model.APIApplications
var User *model.APIUsers
var Addr *model.APIAddrs
var MenuType *model.APIMenuTypes
var Menu *model.APIMenus
var CareType *model.CareTypes
var CareTag *model.CareTags
var Care *model.Cares
var CarerTag *model.CarerTags
var Carer *model.Carers

var Order *model.APIOOrders
var MiniOrder *model.APIOOrders
var CareOrder *model.CareOrders
var MiniCareOrder *model.CareOrders
var CarerOrder *model.CareOrders
var MiniCarerOrder *model.CareOrders

var Article *model.OnlineArticles
var LocType *model.OnlineLocTypes
var Loc *model.OnlineLocs
var UserType *model.OnlineUserTypes
var Patient *model.OnlinePatients
var Spec *model.SpecGroup
var Brand *model.ShopBrands
var Cate1 *model.ShopCates
var Cate2 *model.ShopCates
var Cate3 *model.ShopCates
var SpecGroup *model.ShopSpecGroups

//单元测试基境
func TestMain(m *testing.M) {

	println(fmt.Sprintf("-------------------main start-------------------"))

	cleanTables := flag.Bool("clean", false, "Set Clean tables (default false)")
	flag.Parse()

	model.GetAuthToken()
	App = model.GetApp()
	User = model.GetUserByIdCardNo()
	model.GetFake()

	Addr = model.CreateAddr(User.ID)
	MenuType = model.CreateMenuType()
	model.CreateMenuTag()
	Menu = model.CreateMenu(true)
	model.CreateMenu(false)

	Order = model.CreateOrder("I202008241612348468756914", User.ID, model.OrderAppTypeBed, model.IOrderPayTypeWechat)
	Order.OrderAddr = model.CreateOrderAddr(Order.ID)
	Order.OrderMenus = model.CreateOrderMenus(Menu, MenuType.Name, Order.ID, 5)
	Order.OrderComments = model.CreateOrderComments(5, User.ID, Order.ID, COMMENTABLETYPE)

	MiniOrder = model.CreateOrder("O202008241612348468756914", User.ID, model.OrderAppTypeMini, model.IOrderPayTypeAli)
	MiniOrder.OrderAddr = model.CreateOrderAddr(MiniOrder.ID)
	MiniOrder.OrderMenus = model.CreateOrderMenus(Menu, MenuType.Name, MiniOrder.ID, 3)
	MiniOrder.OrderComments = model.CreateOrderComments(4, User.ID, MiniOrder.ID, COMMENTABLETYPE)

	CareTag = model.CreateCareTag()
	CareType = model.CreateCareType()
	Care = model.CreateCare(CareTag.ID, CareType.ID, true)
	model.CreateCare(CareTag.ID, CareType.ID, false)

	MiniCareOrder = model.CreateCareOrder(Care.TimeType, "OC202008241612348468756914", User.ID, 0, model.OrderAppTypeMini, model.IOrderPayTypeAli, model.CareMaxPrice)
	MiniCareOrder.CareOrderComments = model.CreateCareOrderComments(4, User.ID, MiniCareOrder.ID, 1, CARECOMMENTABLETYPE)
	MiniCareOrder.CareOrderInfo = model.CreateCareOrderInfo(MiniCareOrder.ID, Care, App.Name, CareType.Name, CareTag.Name)
	MiniCareOrder.CareOrderAddr = model.CreateCareOrderAddr(MiniCareOrder.ID, Addr)

	CareOrder = model.CreateCareOrder(Care.TimeType, "IC202008241612348468756914", User.ID, 0, model.OrderAppTypeMini, model.IOrderPayTypeAli, model.CareMaxPrice)
	CareOrder.CareOrderComments = model.CreateCareOrderComments(3, User.ID, CareOrder.ID, 1, CARECOMMENTABLETYPE)
	CareOrder.CareOrderInfo = model.CreateCareOrderInfo(CareOrder.ID, Care, App.Name, CareType.Name, CareTag.Name)
	CareOrder.CareOrderAddr = model.CreateCareOrderAddr(CareOrder.ID, Addr)

	CarerTag = model.CreateCarerTag()
	Carer = model.CreateCarer(CarerTag.ID, true)
	model.CreateCarer(CarerTag.ID, false)

	MiniCarerOrder = model.CreateCareOrder(Carer.TimeType, "OC202008241612348468756914", User.ID, Carer.ID, model.OrderAppTypeMini, model.IOrderPayTypeAli, model.CarerPrice)
	MiniCarerOrder.CareOrderComments = model.CreateCareOrderComments(4, User.ID, MiniCarerOrder.ID, 1, CARECOMMENTABLETYPE)
	MiniCarerOrder.CareOrderCarerInfo = model.CreateCareOrderCarerInfo(MiniCarerOrder.ID, Carer, App.Name, CarerTag.Name)
	MiniCarerOrder.CareOrderAddr = model.CreateCareOrderAddr(MiniCarerOrder.ID, Addr)

	CarerOrder = model.CreateCareOrder(Carer.TimeType, "IC202008241612348468756914", User.ID, Carer.ID, model.OrderAppTypeMini, model.IOrderPayTypeAli, model.CarerPrice)
	CarerOrder.CareOrderComments = model.CreateCareOrderComments(7, User.ID, CarerOrder.ID, 1, CARECOMMENTABLETYPE)
	CarerOrder.CareOrderCarerInfo = model.CreateCareOrderCarerInfo(CarerOrder.ID, Carer, App.Name, CarerTag.Name)
	CarerOrder.CareOrderAddr = model.CreateCareOrderAddr(CarerOrder.ID, Addr)

	Article = model.CreateArticle()
	LocType = model.CreateLocType()
	Loc = model.CreateLoc()
	UserType = model.CreateUserType()
	Patient = model.CreatePatient(User.ID)

	model.CreateBrand(false)
	Brand = model.CreateBrand(true)
	Cate1 = model.CreateCate(0, 1)
	Cate2 = model.CreateCate(Cate1.ID, 2)
	Cate3 = model.CreateCate(Cate2.ID, 3)
	Spec := &model.SpecGroup{
		Name: "基本信息",
		Params: []*model.SpecParam{
			{
				Name:      "机身颜色",
				Unit:      "",
				Numeric:   false,
				Generic:   false,
				Searching: true,
				Value:     []string{"白色", "金色", "黑色"},
			}, {
				Name:      "机身长度",
				Unit:      "mm",
				Numeric:   true,
				Generic:   true,
				Searching: true,
				Value:     []string{"112"},
			}, {
				Name:      "输入方法",
				Unit:      "",
				Numeric:   false,
				Generic:   true,
				Searching: true,
				Value:     []string{"触控"},
			}, {
				Name:      "内存",
				Unit:      "G",
				Numeric:   true,
				Generic:   false,
				Searching: true,
				Value:     []string{"6", "8"},
			}, {
				Name:      "机身存储",
				Unit:      "GB",
				Numeric:   true,
				Generic:   false,
				Searching: true,
				Value:     []string{"16", "32", "64", "128"},
			},
		},
	}
	SpecGroup = model.CreateSpecGroup(Cate1.ID, Spec)
	spu := model.CreateSpu(Brand.ID, Cate1.ID, 1, "", "", Spec)

	model.CreateShopOrder("S202008241612348468756914", User.ID, model.IOrderPayTypeAli, model.OrderAppTypeMini, spu.Skus)

	flag.Parse()
	exitCode := m.Run()

	model.ClearToken()
	model.ClearTables(*cleanTables)
	model.Close()

	println(fmt.Sprintf("-------------------main end-------------------"))

	os.Exit(exitCode)
}
