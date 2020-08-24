package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CareOrdersMgr struct {
	*_BaseMgr
}

// CareOrdersMgr open func
func CareOrdersMgr(db *gorm.DB) *_CareOrdersMgr {
	if db == nil {
		panic(fmt.Errorf("CareOrdersMgr need init by db"))
	}
	return &_CareOrdersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("care_orders"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CareOrdersMgr) GetTableName() string {
	return "care_orders"
}

// Get 获取
func (obj *_CareOrdersMgr) Get() (result CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CareOrdersMgr) Gets() (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CareOrdersMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNo order_no获取 订单号
func (obj *_CareOrdersMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithStatus status获取 订单状态：1：待付款，2：已付款，3，进行中，4：已取消，5：已完成
func (obj *_CareOrdersMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithPayType pay_type获取 订单支付类型：1：微信，2：支付宝
func (obj *_CareOrdersMgr) WithPayType(payType int) Option {
	return optionFunc(func(o *options) { o.query["pay_type"] = payType })
}

// WithTotal total获取 总金额
func (obj *_CareOrdersMgr) WithTotal(total float64) Option {
	return optionFunc(func(o *options) { o.query["total"] = total })
}

// WithRmk rmk获取 备注
func (obj *_CareOrdersMgr) WithRmk(rmk string) Option {
	return optionFunc(func(o *options) { o.query["rmk"] = rmk })
}

// WithStartAt start_at获取 服务开始时间
func (obj *_CareOrdersMgr) WithStartAt(startAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_at"] = startAt })
}

// WithEndAt end_at获取 服务结束时间
func (obj *_CareOrdersMgr) WithEndAt(endAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_at"] = endAt })
}

// WithOpenID open_id获取 open_id
func (obj *_CareOrdersMgr) WithOpenID(openID string) Option {
	return optionFunc(func(o *options) { o.query["open_id"] = openID })
}

// WithTradeType trade_type获取 JSAPI、NATIVE、APP
func (obj *_CareOrdersMgr) WithTradeType(tradeType string) Option {
	return optionFunc(func(o *options) { o.query["trade_type"] = tradeType })
}

// WithIsReturn is_return获取 是否退款
func (obj *_CareOrdersMgr) WithIsReturn(isReturn bool) Option {
	return optionFunc(func(o *options) { o.query["is_return"] = isReturn })
}

// WithApplicationID application_id获取 application_id
func (obj *_CareOrdersMgr) WithApplicationID(applicationID int) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithUserID user_id获取 user_id
func (obj *_CareOrdersMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithIsDeleted is_deleted获取
func (obj *_CareOrdersMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithCreateAt create_at获取
func (obj *_CareOrdersMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CareOrdersMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithAppType app_type获取 订单应用类型，1：小程序，2：床旁
func (obj *_CareOrdersMgr) WithAppType(appType int) Option {
	return optionFunc(func(o *options) { o.query["app_type"] = appType })
}

// WithCarerID carer_id获取 护工id
func (obj *_CareOrdersMgr) WithCarerID(carerID int) Option {
	return optionFunc(func(o *options) { o.query["carer_id"] = carerID })
}

// GetByOption 功能选项模式获取
func (obj *_CareOrdersMgr) GetByOption(opts ...Option) (result CareOrders, err error) {
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
func (obj *_CareOrdersMgr) GetByOptions(opts ...Option) (results []*CareOrders, err error) {
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
func (obj *_CareOrdersMgr) GetFromID(id int) (result CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CareOrdersMgr) GetBatchFromID(ids []int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 订单号
func (obj *_CareOrdersMgr) GetFromOrderNo(orderNo string) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_no = ?", orderNo).Find(&results).Error

	return
}

// GetBatchFromOrderNo 批量唯一主键查找 订单号
func (obj *_CareOrdersMgr) GetBatchFromOrderNo(orderNos []string) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_no IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 订单状态：1：待付款，2：已付款，3，进行中，4：已取消，5：已完成
func (obj *_CareOrdersMgr) GetFromStatus(status int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 订单状态：1：待付款，2：已付款，3，进行中，4：已取消，5：已完成
func (obj *_CareOrdersMgr) GetBatchFromStatus(statuss []int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromPayType 通过pay_type获取内容 订单支付类型：1：微信，2：支付宝
func (obj *_CareOrdersMgr) GetFromPayType(payType int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pay_type = ?", payType).Find(&results).Error

	return
}

// GetBatchFromPayType 批量唯一主键查找 订单支付类型：1：微信，2：支付宝
func (obj *_CareOrdersMgr) GetBatchFromPayType(payTypes []int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pay_type IN (?)", payTypes).Find(&results).Error

	return
}

// GetFromTotal 通过total获取内容 总金额
func (obj *_CareOrdersMgr) GetFromTotal(total float64) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("total = ?", total).Find(&results).Error

	return
}

// GetBatchFromTotal 批量唯一主键查找 总金额
func (obj *_CareOrdersMgr) GetBatchFromTotal(totals []float64) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("total IN (?)", totals).Find(&results).Error

	return
}

// GetFromRmk 通过rmk获取内容 备注
func (obj *_CareOrdersMgr) GetFromRmk(rmk string) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("rmk = ?", rmk).Find(&results).Error

	return
}

// GetBatchFromRmk 批量唯一主键查找 备注
func (obj *_CareOrdersMgr) GetBatchFromRmk(rmks []string) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("rmk IN (?)", rmks).Find(&results).Error

	return
}

// GetFromStartAt 通过start_at获取内容 服务开始时间
func (obj *_CareOrdersMgr) GetFromStartAt(startAt time.Time) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("start_at = ?", startAt).Find(&results).Error

	return
}

// GetBatchFromStartAt 批量唯一主键查找 服务开始时间
func (obj *_CareOrdersMgr) GetBatchFromStartAt(startAts []time.Time) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("start_at IN (?)", startAts).Find(&results).Error

	return
}

// GetFromEndAt 通过end_at获取内容 服务结束时间
func (obj *_CareOrdersMgr) GetFromEndAt(endAt time.Time) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("end_at = ?", endAt).Find(&results).Error

	return
}

// GetBatchFromEndAt 批量唯一主键查找 服务结束时间
func (obj *_CareOrdersMgr) GetBatchFromEndAt(endAts []time.Time) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("end_at IN (?)", endAts).Find(&results).Error

	return
}

// GetFromOpenID 通过open_id获取内容 open_id
func (obj *_CareOrdersMgr) GetFromOpenID(openID string) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("open_id = ?", openID).Find(&results).Error

	return
}

// GetBatchFromOpenID 批量唯一主键查找 open_id
func (obj *_CareOrdersMgr) GetBatchFromOpenID(openIDs []string) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("open_id IN (?)", openIDs).Find(&results).Error

	return
}

// GetFromTradeType 通过trade_type获取内容 JSAPI、NATIVE、APP
func (obj *_CareOrdersMgr) GetFromTradeType(tradeType string) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("trade_type = ?", tradeType).Find(&results).Error

	return
}

// GetBatchFromTradeType 批量唯一主键查找 JSAPI、NATIVE、APP
func (obj *_CareOrdersMgr) GetBatchFromTradeType(tradeTypes []string) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("trade_type IN (?)", tradeTypes).Find(&results).Error

	return
}

// GetFromIsReturn 通过is_return获取内容 是否退款
func (obj *_CareOrdersMgr) GetFromIsReturn(isReturn bool) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_return = ?", isReturn).Find(&results).Error

	return
}

// GetBatchFromIsReturn 批量唯一主键查找 是否退款
func (obj *_CareOrdersMgr) GetBatchFromIsReturn(isReturns []bool) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_return IN (?)", isReturns).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容 application_id
func (obj *_CareOrdersMgr) GetFromApplicationID(applicationID int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量唯一主键查找 application_id
func (obj *_CareOrdersMgr) GetBatchFromApplicationID(applicationIDs []int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id IN (?)", applicationIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 user_id
func (obj *_CareOrdersMgr) GetFromUserID(userID int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 user_id
func (obj *_CareOrdersMgr) GetBatchFromUserID(userIDs []int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CareOrdersMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CareOrdersMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CareOrdersMgr) GetFromCreateAt(createAt time.Time) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CareOrdersMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CareOrdersMgr) GetFromUpdateAt(updateAt time.Time) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CareOrdersMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromAppType 通过app_type获取内容 订单应用类型，1：小程序，2：床旁
func (obj *_CareOrdersMgr) GetFromAppType(appType int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_type = ?", appType).Find(&results).Error

	return
}

// GetBatchFromAppType 批量唯一主键查找 订单应用类型，1：小程序，2：床旁
func (obj *_CareOrdersMgr) GetBatchFromAppType(appTypes []int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_type IN (?)", appTypes).Find(&results).Error

	return
}

// GetFromCarerID 通过carer_id获取内容 护工id
func (obj *_CareOrdersMgr) GetFromCarerID(carerID int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_id = ?", carerID).Find(&results).Error

	return
}

// GetBatchFromCarerID 批量唯一主键查找 护工id
func (obj *_CareOrdersMgr) GetBatchFromCarerID(carerIDs []int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_id IN (?)", carerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CareOrdersMgr) FetchByPrimaryKey(id int) (result CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_CareOrdersMgr) FetchIndexByUserID(applicationID int, userID int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ? AND user_id = ?", applicationID, userID).Find(&results).Error

	return
}

// FetchIndexByCarerID  获取多个内容
func (obj *_CareOrdersMgr) FetchIndexByCarerID(carerID int) (results []*CareOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_id = ?", carerID).Find(&results).Error

	return
}
