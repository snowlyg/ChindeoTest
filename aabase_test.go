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
	model.GetFake()

	Addr = model.CreateAddr(User.ID)
	MenuType = model.CreateMenuType()
	MenuTag = model.CreateMenuTag()
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

	flag.Parse()
	exitCode := m.Run()

	model.ClearToken()
	model.ClearTables()
	model.Close()

	println(fmt.Sprintf("-------------------main end-------------------"))

	os.Exit(exitCode)
}
