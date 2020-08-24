package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIOReturnOrdersMgr struct {
	*_BaseMgr
}

// APIOReturnOrdersMgr open func
func APIOReturnOrdersMgr(db *gorm.DB) *_APIOReturnOrdersMgr {
	if db == nil {
		panic(fmt.Errorf("APIOReturnOrdersMgr need init by db"))
	}
	return &_APIOReturnOrdersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_o_return_orders"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIOReturnOrdersMgr) GetTableName() string {
	return "api_o_return_orders"
}

// Get 获取
func (obj *_APIOReturnOrdersMgr) Get() (result APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIOReturnOrdersMgr) Gets() (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIOReturnOrdersMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNo order_no获取 订单号
func (obj *_APIOReturnOrdersMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithStatus status获取 订单状态：1，待审核，2：审核通过，3：审核驳回，4：退款成功，5：退款失败，
func (obj *_APIOReturnOrdersMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithAmount amount获取 总数量
func (obj *_APIOReturnOrdersMgr) WithAmount(amount int) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithTotal total获取 总金额
func (obj *_APIOReturnOrdersMgr) WithTotal(total float64) Option {
	return optionFunc(func(o *options) { o.query["total"] = total })
}

// WithOpenID open_id获取 open_id
func (obj *_APIOReturnOrdersMgr) WithOpenID(openID string) Option {
	return optionFunc(func(o *options) { o.query["open_id"] = openID })
}

// WithAppType app_type获取 订单应用类型，1：小程序，2：床旁
func (obj *_APIOReturnOrdersMgr) WithAppType(appType int) Option {
	return optionFunc(func(o *options) { o.query["app_type"] = appType })
}

// WithTradeType trade_type获取 JSAPI、NATIVE、APP
func (obj *_APIOReturnOrdersMgr) WithTradeType(tradeType string) Option {
	return optionFunc(func(o *options) { o.query["trade_type"] = tradeType })
}

// WithCreateAt create_at获取
func (obj *_APIOReturnOrdersMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIOReturnOrdersMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIOReturnOrdersMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithOOrderID o_order_id获取 o_order_id
func (obj *_APIOReturnOrdersMgr) WithOOrderID(oOrderID int) Option {
	return optionFunc(func(o *options) { o.query["o_order_id"] = oOrderID })
}

// WithApplicationID application_id获取 application_id
func (obj *_APIOReturnOrdersMgr) WithApplicationID(applicationID int) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithPayType pay_type获取
func (obj *_APIOReturnOrdersMgr) WithPayType(payType bool) Option {
	return optionFunc(func(o *options) { o.query["pay_type"] = payType })
}

// WithUserID user_id获取 user_id
func (obj *_APIOReturnOrdersMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// GetByOption 功能选项模式获取
func (obj *_APIOReturnOrdersMgr) GetByOption(opts ...Option) (result APIOReturnOrders, err error) {
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
func (obj *_APIOReturnOrdersMgr) GetByOptions(opts ...Option) (results []*APIOReturnOrders, err error) {
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
func (obj *_APIOReturnOrdersMgr) GetFromID(id int) (result APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIOReturnOrdersMgr) GetBatchFromID(ids []int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 订单号
func (obj *_APIOReturnOrdersMgr) GetFromOrderNo(orderNo string) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_no = ?", orderNo).Find(&results).Error

	return
}

// GetBatchFromOrderNo 批量唯一主键查找 订单号
func (obj *_APIOReturnOrdersMgr) GetBatchFromOrderNo(orderNos []string) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_no IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 订单状态：1，待审核，2：审核通过，3：审核驳回，4：退款成功，5：退款失败，
func (obj *_APIOReturnOrdersMgr) GetFromStatus(status int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 订单状态：1，待审核，2：审核通过，3：审核驳回，4：退款成功，5：退款失败，
func (obj *_APIOReturnOrdersMgr) GetBatchFromStatus(statuss []int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 总数量
func (obj *_APIOReturnOrdersMgr) GetFromAmount(amount int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量唯一主键查找 总数量
func (obj *_APIOReturnOrdersMgr) GetBatchFromAmount(amounts []int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount IN (?)", amounts).Find(&results).Error

	return
}

// GetFromTotal 通过total获取内容 总金额
func (obj *_APIOReturnOrdersMgr) GetFromTotal(total float64) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("total = ?", total).Find(&results).Error

	return
}

// GetBatchFromTotal 批量唯一主键查找 总金额
func (obj *_APIOReturnOrdersMgr) GetBatchFromTotal(totals []float64) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("total IN (?)", totals).Find(&results).Error

	return
}

// GetFromOpenID 通过open_id获取内容 open_id
func (obj *_APIOReturnOrdersMgr) GetFromOpenID(openID string) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("open_id = ?", openID).Find(&results).Error

	return
}

// GetBatchFromOpenID 批量唯一主键查找 open_id
func (obj *_APIOReturnOrdersMgr) GetBatchFromOpenID(openIDs []string) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("open_id IN (?)", openIDs).Find(&results).Error

	return
}

// GetFromAppType 通过app_type获取内容 订单应用类型，1：小程序，2：床旁
func (obj *_APIOReturnOrdersMgr) GetFromAppType(appType int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_type = ?", appType).Find(&results).Error

	return
}

// GetBatchFromAppType 批量唯一主键查找 订单应用类型，1：小程序，2：床旁
func (obj *_APIOReturnOrdersMgr) GetBatchFromAppType(appTypes []int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_type IN (?)", appTypes).Find(&results).Error

	return
}

// GetFromTradeType 通过trade_type获取内容 JSAPI、NATIVE、APP
func (obj *_APIOReturnOrdersMgr) GetFromTradeType(tradeType string) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("trade_type = ?", tradeType).Find(&results).Error

	return
}

// GetBatchFromTradeType 批量唯一主键查找 JSAPI、NATIVE、APP
func (obj *_APIOReturnOrdersMgr) GetBatchFromTradeType(tradeTypes []string) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("trade_type IN (?)", tradeTypes).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIOReturnOrdersMgr) GetFromCreateAt(createAt time.Time) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIOReturnOrdersMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIOReturnOrdersMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIOReturnOrdersMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIOReturnOrdersMgr) GetFromIsDeleted(isDeleted time.Time) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIOReturnOrdersMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromOOrderID 通过o_order_id获取内容 o_order_id
func (obj *_APIOReturnOrdersMgr) GetFromOOrderID(oOrderID int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("o_order_id = ?", oOrderID).Find(&results).Error

	return
}

// GetBatchFromOOrderID 批量唯一主键查找 o_order_id
func (obj *_APIOReturnOrdersMgr) GetBatchFromOOrderID(oOrderIDs []int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("o_order_id IN (?)", oOrderIDs).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容 application_id
func (obj *_APIOReturnOrdersMgr) GetFromApplicationID(applicationID int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量唯一主键查找 application_id
func (obj *_APIOReturnOrdersMgr) GetBatchFromApplicationID(applicationIDs []int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id IN (?)", applicationIDs).Find(&results).Error

	return
}

// GetFromPayType 通过pay_type获取内容
func (obj *_APIOReturnOrdersMgr) GetFromPayType(payType bool) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pay_type = ?", payType).Find(&results).Error

	return
}

// GetBatchFromPayType 批量唯一主键查找
func (obj *_APIOReturnOrdersMgr) GetBatchFromPayType(payTypes []bool) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pay_type IN (?)", payTypes).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 user_id
func (obj *_APIOReturnOrdersMgr) GetFromUserID(userID int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 user_id
func (obj *_APIOReturnOrdersMgr) GetBatchFromUserID(userIDs []int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIOReturnOrdersMgr) FetchByPrimaryKey(id int) (result APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByOOrderID  获取多个内容
func (obj *_APIOReturnOrdersMgr) FetchIndexByOOrderID(orderNo string, oOrderID int, applicationID int) (results []*APIOReturnOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_no = ? AND o_order_id = ? AND application_id = ?", orderNo, oOrderID, applicationID).Find(&results).Error

	return
}
