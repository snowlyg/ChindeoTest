package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CareReturnOrdersMgr struct {
	*_BaseMgr
}

// CareReturnOrdersMgr open func
func CareReturnOrdersMgr(db *gorm.DB) *_CareReturnOrdersMgr {
	if db == nil {
		panic(fmt.Errorf("CareReturnOrdersMgr need init by db"))
	}
	return &_CareReturnOrdersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("care_return_orders"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CareReturnOrdersMgr) GetTableName() string {
	return "care_return_orders"
}

// Get 获取
func (obj *_CareReturnOrdersMgr) Get() (result CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CareReturnOrdersMgr) Gets() (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CareReturnOrdersMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNo order_no获取 订单号
func (obj *_CareReturnOrdersMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithStatus status获取 订单状态：1，待审核，2：审核通过，3：审核驳回，4：退款成功，5：退款失败，
func (obj *_CareReturnOrdersMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithPayType pay_type获取 订单支付类型：1：微信，2：支付宝
func (obj *_CareReturnOrdersMgr) WithPayType(payType int) Option {
	return optionFunc(func(o *options) { o.query["pay_type"] = payType })
}

// WithTotal total获取 总金额
func (obj *_CareReturnOrdersMgr) WithTotal(total float64) Option {
	return optionFunc(func(o *options) { o.query["total"] = total })
}

// WithOpenID open_id获取 open_id
func (obj *_CareReturnOrdersMgr) WithOpenID(openID string) Option {
	return optionFunc(func(o *options) { o.query["open_id"] = openID })
}

// WithTradeType trade_type获取 JSAPI、NATIVE、APP
func (obj *_CareReturnOrdersMgr) WithTradeType(tradeType string) Option {
	return optionFunc(func(o *options) { o.query["trade_type"] = tradeType })
}

// WithCreateAt create_at获取
func (obj *_CareReturnOrdersMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CareReturnOrdersMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_CareReturnOrdersMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithCareOrderID care_order_id获取 care_order_id
func (obj *_CareReturnOrdersMgr) WithCareOrderID(careOrderID int) Option {
	return optionFunc(func(o *options) { o.query["care_order_id"] = careOrderID })
}

// WithApplicationID application_id获取 application_id
func (obj *_CareReturnOrdersMgr) WithApplicationID(applicationID int) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithUserID user_id获取 user_id
func (obj *_CareReturnOrdersMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithAppType app_type获取 订单应用类型，1：小程序，2：床旁
func (obj *_CareReturnOrdersMgr) WithAppType(appType int) Option {
	return optionFunc(func(o *options) { o.query["app_type"] = appType })
}

// WithCarerID carer_id获取 护工id
func (obj *_CareReturnOrdersMgr) WithCarerID(carerID int) Option {
	return optionFunc(func(o *options) { o.query["carer_id"] = carerID })
}

// GetByOption 功能选项模式获取
func (obj *_CareReturnOrdersMgr) GetByOption(opts ...Option) (result CareReturnOrders, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CareReturnOrdersMgr) GetByOptions(opts ...Option) (results []*CareReturnOrders, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_CareReturnOrdersMgr) GetFromID(id int) (result CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CareReturnOrdersMgr) GetBatchFromID(ids []int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 订单号
func (obj *_CareReturnOrdersMgr) GetFromOrderNo(orderNo string) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_no = ?", orderNo).Find(&results).Error

	return
}

// GetBatchFromOrderNo 批量唯一主键查找 订单号
func (obj *_CareReturnOrdersMgr) GetBatchFromOrderNo(orderNos []string) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_no IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 订单状态：1，待审核，2：审核通过，3：审核驳回，4：退款成功，5：退款失败，
func (obj *_CareReturnOrdersMgr) GetFromStatus(status int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 订单状态：1，待审核，2：审核通过，3：审核驳回，4：退款成功，5：退款失败，
func (obj *_CareReturnOrdersMgr) GetBatchFromStatus(statuss []int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromPayType 通过pay_type获取内容 订单支付类型：1：微信，2：支付宝
func (obj *_CareReturnOrdersMgr) GetFromPayType(payType int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pay_type = ?", payType).Find(&results).Error

	return
}

// GetBatchFromPayType 批量唯一主键查找 订单支付类型：1：微信，2：支付宝
func (obj *_CareReturnOrdersMgr) GetBatchFromPayType(payTypes []int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pay_type IN (?)", payTypes).Find(&results).Error

	return
}

// GetFromTotal 通过total获取内容 总金额
func (obj *_CareReturnOrdersMgr) GetFromTotal(total float64) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("total = ?", total).Find(&results).Error

	return
}

// GetBatchFromTotal 批量唯一主键查找 总金额
func (obj *_CareReturnOrdersMgr) GetBatchFromTotal(totals []float64) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("total IN (?)", totals).Find(&results).Error

	return
}

// GetFromOpenID 通过open_id获取内容 open_id
func (obj *_CareReturnOrdersMgr) GetFromOpenID(openID string) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("open_id = ?", openID).Find(&results).Error

	return
}

// GetBatchFromOpenID 批量唯一主键查找 open_id
func (obj *_CareReturnOrdersMgr) GetBatchFromOpenID(openIDs []string) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("open_id IN (?)", openIDs).Find(&results).Error

	return
}

// GetFromTradeType 通过trade_type获取内容 JSAPI、NATIVE、APP
func (obj *_CareReturnOrdersMgr) GetFromTradeType(tradeType string) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("trade_type = ?", tradeType).Find(&results).Error

	return
}

// GetBatchFromTradeType 批量唯一主键查找 JSAPI、NATIVE、APP
func (obj *_CareReturnOrdersMgr) GetBatchFromTradeType(tradeTypes []string) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("trade_type IN (?)", tradeTypes).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CareReturnOrdersMgr) GetFromCreateAt(createAt time.Time) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CareReturnOrdersMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CareReturnOrdersMgr) GetFromUpdateAt(updateAt time.Time) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CareReturnOrdersMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CareReturnOrdersMgr) GetFromIsDeleted(isDeleted time.Time) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CareReturnOrdersMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCareOrderID 通过care_order_id获取内容 care_order_id
func (obj *_CareReturnOrdersMgr) GetFromCareOrderID(careOrderID int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id = ?", careOrderID).Find(&results).Error

	return
}

// GetBatchFromCareOrderID 批量唯一主键查找 care_order_id
func (obj *_CareReturnOrdersMgr) GetBatchFromCareOrderID(careOrderIDs []int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id IN (?)", careOrderIDs).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容 application_id
func (obj *_CareReturnOrdersMgr) GetFromApplicationID(applicationID int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量唯一主键查找 application_id
func (obj *_CareReturnOrdersMgr) GetBatchFromApplicationID(applicationIDs []int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id IN (?)", applicationIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 user_id
func (obj *_CareReturnOrdersMgr) GetFromUserID(userID int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 user_id
func (obj *_CareReturnOrdersMgr) GetBatchFromUserID(userIDs []int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromAppType 通过app_type获取内容 订单应用类型，1：小程序，2：床旁
func (obj *_CareReturnOrdersMgr) GetFromAppType(appType int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_type = ?", appType).Find(&results).Error

	return
}

// GetBatchFromAppType 批量唯一主键查找 订单应用类型，1：小程序，2：床旁
func (obj *_CareReturnOrdersMgr) GetBatchFromAppType(appTypes []int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_type IN (?)", appTypes).Find(&results).Error

	return
}

// GetFromCarerID 通过carer_id获取内容 护工id
func (obj *_CareReturnOrdersMgr) GetFromCarerID(carerID int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_id = ?", carerID).Find(&results).Error

	return
}

// GetBatchFromCarerID 批量唯一主键查找 护工id
func (obj *_CareReturnOrdersMgr) GetBatchFromCarerID(carerIDs []int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_id IN (?)", carerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CareReturnOrdersMgr) FetchByPrimaryKey(id int) (result CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByCareOrderID  获取多个内容
func (obj *_CareReturnOrdersMgr) FetchIndexByCareOrderID(careOrderID int, applicationID int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id = ? AND application_id = ?", careOrderID, applicationID).Find(&results).Error

	return
}

// FetchIndexByCarerID  获取多个内容
func (obj *_CareReturnOrdersMgr) FetchIndexByCarerID(carerID int) (results []*CareReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_id = ?", carerID).Find(&results).Error

	return
}
