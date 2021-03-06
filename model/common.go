package model

import (
	"encoding/json"
	"fmt"
	"github.com/manveru/faker"
	"github.com/shopspring/decimal"
	"strconv"
)

const AppId = 13
const IdCardNo = "456952158962254456"
const CareMiniPrice = 1.00
const CareMaxPrice = 50.00
const CarerPrice = 50.00

const AuthTypeMiniWechat = 1 // 小程序
const AUTH_TYPE_APP = 2      //APP
const AUTH_TYPE_WECHAT = 3   //公众号
const AuthTypeServer = 4     //服务

const IOrderStatusForPay = 1         // 待付款
const IOrderStatusForDelivery = 2    // 已付款
const IOrderStatusForDeliverying = 3 // 配送中/进行中/待收货
const IOrderStatusForCancel = 4      // 已取消
const IOrderStatusForFinish = 5      // 已完成
const IOrderStatusForComment = 6     // 待评价/已收货

const I_ORDER_PAY_TYPE_NOPAY = 0 // 未支付
const IOrderPayTypeWechat = 1    // 微信
const IOrderPayTypeAli = 2       // 支付宝

const OrderAppTypeMini = 1 // 小程序
const OrderAppTypeBed = 2  // 床旁

const IReturnOrderStatusForAudit = 1       // 待审核
const I_RETURN_ORDER_STATUS_FOR_PASS = 2   // 审核通过
const I_RETURN_ORDER_STATUS_FOR_RE = 3     // 审核驳回
const I_RETURN_ORDER_STATUS_FOR_FINISH = 4 // 退款成功
const I_RETURN_ORDER_STATUS_FOR_FAULT = 5  // 退款失败

type MenuTimeType int

const (
	MENU_TIME_TYPE_O MenuTimeType = iota // 早餐
	MenuTimeTypeB                        // 早餐
	MenuTimeTypeM                        // 中餐
	MenuTimeTypeN                        // 晚餐
)

func (m MenuTimeType) String() string {
	switch m {
	case MenuTimeTypeB:
		return "早餐"
	case MenuTimeTypeM:
		return "中餐"
	case MenuTimeTypeN:
		return "晚餐"
	}

	return ""
}

func GetSToI(i interface{}) int {
	atoi, _ := strconv.Atoi(GetS(i))
	return atoi

}

func GetSToF(i interface{}) float64 {
	atoi, _ := strconv.ParseFloat(GetS(i), 10)
	return atoi
}

func GetS(i interface{}) string {
	tt, ok := i.(string)
	if ok {
		return tt
	}
	return ""
}

func GetI(i interface{}) int64 {
	tt, ok := i.(int64)
	if ok {
		return tt
	}
	return 0
}
func GetF(i interface{}) float64 {
	tt, ok := i.(float64)
	if ok {
		return tt
	}
	return 0
}

func Ftos(f float64) string {
	df, _ := decimal.NewFromFloat(f).Float64()
	return fmt.Sprintf("%.2f", df)
}

var Fake *faker.Faker

func GetPics() string {
	pics := []string{"https", "https"}
	jpics, _ := json.Marshal(pics)
	return string(jpics)
}

func GetFake() {
	var err error
	Fake, err = faker.New("en")
	if err != nil {
		fmt.Println(fmt.Sprintf("faker create error :%v", err))
	}
}
