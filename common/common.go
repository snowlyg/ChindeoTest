package common

const I_ORDER_STATUS_FOR_PAY = 1         // 待付款
const I_ORDER_STATUS_FOR_DELIVERY = 2    // 已付款
const I_ORDER_STATUS_FOR_DELIVERYING = 3 // 配送中/进行中
const I_ORDER_STATUS_FOR_CANCEL = 4      // 已取消
const I_ORDER_STATUS_FOR_FINISH = 5      // 已完成

const I_ORDER_PAY_TYPE_NOPAY = 0  // 未支付
const I_ORDER_PAY_TYPE_WECHAT = 1 // 微信
const I_ORDER_PAY_TYPE_ALI = 2    // 支付宝

const ORDER_APP_TYPE_MINI = 1 // 小程序
const ORDER_APP_TYPE_BED = 2  // 床旁

const I_RETURN_ORDER_STATUS_FOR_AUDIT = 1  // 待审核
const I_RETURN_ORDER_STATUS_FOR_PASS = 2   // 审核通过
const I_RETURN_ORDER_STATUS_FOR_RE = 3     // 审核驳回
const I_RETURN_ORDER_STATUS_FOR_FINISH = 4 // 退款成功
const I_RETURN_ORDER_STATUS_FOR_FAULT = 5  // 退款失败

type MenuTimeType int

const (
	MENU_TIME_TYPE_B MenuTimeType = iota // 早餐
	MENU_TIME_TYPE_M                     // 中餐
	MENU_TIME_TYPE_N                     // 晚餐
)

func (m MenuTimeType) String() string {
	switch m {
	case MENU_TIME_TYPE_B:
		return "早餐"
	case MENU_TIME_TYPE_M:
		return "中餐"
	case MENU_TIME_TYPE_N:
		return "晚餐"
	}

	return ""
}
