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
var MenuTag *model.APIMenuTags
var Menu *model.APIMenus

var Order *model.APIOOrders
var OrderAddr *model.APIOOrderAddrs
var OrderMenu *model.APIOOrderMenus
var OrderMenus []*model.APIOOrderMenus
var OrderComment *model.Comments

var MiniOrder *model.APIOOrders
var MiniOrderAddr *model.APIOOrderAddrs
var MiniOrderMenu *model.APIOOrderMenus
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

//单元测试基境
func TestMain(m *testing.M) {

	println(fmt.Sprintf("-------------------main start-------------------"))

	model.GetAuthToken()
	App = model.GetApp()
	User = model.GetUserByIdCardNo()
	//model.GetFake()
	//model.GetPics()

	Addr = model.CreateAddr(User.ID)
	MenuType = model.CreateMenuType()
	MenuTag = model.CreateMenuTag()
	Menu = model.CreateMenu()

	Order = model.CreateOrder("I202008241612348468756914", User.ID, model.ORDER_APP_TYPE_BED, model.I_ORDER_PAY_TYPE_WECHAT)
	OrderAddr = model.CreateOrderAddr(Order.ID)
	OrderMenu = model.CreateOrderMenu(Menu, MenuType.Name, Order.ID)
	OrderMenus = model.CreateOrderMenus(OrderMenu)
	OrderComment = model.CreateOrderComment(User.ID, Order.ID, COMMENTABLETYPE)

	MiniOrder = model.CreateOrder("O202008241612348468756914", User.ID, model.ORDER_APP_TYPE_MINI, model.I_ORDER_PAY_TYPE_ALI)
	MiniOrderAddr = model.CreateOrderAddr(Order.ID)
	MiniOrderMenu = model.CreateOrderMenu(Menu, MenuType.Name, Order.ID)
	MiniOrderMenus = model.CreateOrderMenus(MiniOrderMenu)
	MiniOrderComment = model.CreateOrderComment(User.ID, Order.ID, COMMENTABLETYPE)

	order := model.CreateOrder("O202008241612348468756914", User.ID, model.ORDER_APP_TYPE_MINI, model.I_ORDER_PAY_TYPE_ALI)
	MiniOrderCommentForDel = model.CreateOrderComment(User.ID, order.ID, COMMENTABLETYPE)

	CareTag = model.CreateCareTag()
	CareType = model.CreateCareType()
	Care = model.CreateCare(CareTag.ID, CareType.ID, true)
	model.CreateCare(CareTag.ID, CareType.ID, false)
	MiniCareOrder = model.CreateCareOrder(Care.TimeType, "OC202008241612348468756914", User.ID, 0, model.ORDER_APP_TYPE_MINI, model.I_ORDER_PAY_TYPE_ALI, model.CareMaxPrice)
	MiniCareOrderComment = model.CreateCareOrderComment(User.ID, CareOrder.ID, 1, CARECOMMENTABLETYPE)
	MiniCareOrderInfo = model.CreateCareOrderInfo(CareOrder.ID, Care, App.Name, CareType.Name, CareTag.Name)
	MiniCareOrderAddr = model.CreateCareOrderAddr(CareOrder.ID, Addr)

	CareOrder = model.CreateCareOrder(Care.TimeType, "IC202008241612348468756914", User.ID, 0, model.ORDER_APP_TYPE_MINI, model.I_ORDER_PAY_TYPE_ALI, model.CareMaxPrice)
	CareOrderComment = model.CreateCareOrderComment(User.ID, CareOrder.ID, 1, CARECOMMENTABLETYPE)
	CareOrderInfo = model.CreateCareOrderInfo(CareOrder.ID, Care, App.Name, CareType.Name, CareTag.Name)
	CareOrderAddr = model.CreateCareOrderAddr(CareOrder.ID, Addr)

	CarerTag = model.CreateCarerTag()
	Carer = model.CreateCarer(CarerTag.ID, true)
	model.CreateCarer(CarerTag.ID, false)
	MiniCarerOrder = model.CreateCareOrder(Carer.TimeType, "OC202008241612348468756914", User.ID, Carer.ID, model.ORDER_APP_TYPE_MINI, model.I_ORDER_PAY_TYPE_ALI, model.CarerPrice)
	MiniCarerOrderComment = model.CreateCareOrderComment(User.ID, CarerOrder.ID, 1, CARECOMMENTABLETYPE)
	MiniCarerOrderCarerInfo = model.CreateCareOrderCarerInfo(CarerOrder.ID, Carer, App.Name, CarerTag.Name)
	MiniCarerOrderAddr = model.CreateCareOrderAddr(CarerOrder.ID, Addr)

	CarerOrder = model.CreateCareOrder(Carer.TimeType, "IC202008241612348468756914", User.ID, Carer.ID, model.ORDER_APP_TYPE_MINI, model.I_ORDER_PAY_TYPE_ALI, model.CarerPrice)
	CarerOrderComment = model.CreateCareOrderComment(User.ID, CarerOrder.ID, 1, CARECOMMENTABLETYPE)
	CarerOrderCarerInfo = model.CreateCareOrderCarerInfo(CarerOrder.ID, Carer, App.Name, CarerTag.Name)
	CarerOrderAddr = model.CreateCareOrderAddr(CarerOrder.ID, Addr)

	Article = model.CreateArticle()
	LocType = model.CreateLocType()
	Loc = model.CreateLoc()
	UserType = model.CreateUserType()
	Patient = model.CreatePatient(User.ID)

	flag.Parse()
	exitCode := m.Run()

	model.ClearToken()
	model.ClearTables()
	model.Close()

	println(fmt.Sprintf("-------------------main end-------------------"))

	os.Exit(exitCode)
}
