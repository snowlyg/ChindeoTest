package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIOOrdersMgr struct {
	*_BaseMgr
}

// APIOOrdersMgr open func
func APIOOrdersMgr(db *gorm.DB) *_APIOOrdersMgr {
	if db == nil {
		panic(fmt.Errorf("APIOOrdersMgr need init by db"))
	}
	return &_APIOOrdersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_o_orders"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIOOrdersMgr) GetTableName() string {
	return "api_o_orders"
}

// Get 获取
func (obj *_APIOOrdersMgr) Get() (result APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIOOrdersMgr) Gets() (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIOOrdersMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNo order_no获取 订单号
func (obj *_APIOOrdersMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithStatus status获取 订单状态：1：待付款，2：已付款，3，配送中，4：已取消，5：已完成
func (obj *_APIOOrdersMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithAmount amount获取 总数量
func (obj *_APIOOrdersMgr) WithAmount(amount int) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithTotal total获取 总金额
func (obj *_APIOOrdersMgr) WithTotal(total float64) Option {
	return optionFunc(func(o *options) { o.query["total"] = total })
}

// WithPayType pay_type获取
func (obj *_APIOOrdersMgr) WithPayType(payType bool) Option {
	return optionFunc(func(o *options) { o.query["pay_type"] = payType })
}

// WithRmk rmk获取 备注
func (obj *_APIOOrdersMgr) WithRmk(rmk string) Option {
	return optionFunc(func(o *options) { o.query["rmk"] = rmk })
}

// WithAppType app_type获取 订单应用类型，1：小程序，2：床旁
func (obj *_APIOOrdersMgr) WithAppType(appType int) Option {
	return optionFunc(func(o *options) { o.query["app_type"] = appType })
}

// WithOpenID open_id获取 open_id
func (obj *_APIOOrdersMgr) WithOpenID(openID string) Option {
	return optionFunc(func(o *options) { o.query["open_id"] = openID })
}

// WithTradeType trade_type获取 JSAPI、NATIVE、APP
func (obj *_APIOOrdersMgr) WithTradeType(tradeType string) Option {
	return optionFunc(func(o *options) { o.query["trade_type"] = tradeType })
}

// WithCreateAt create_at获取
func (obj *_APIOOrdersMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIOOrdersMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIOOrdersMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithIsReturn is_return获取 是否退款
func (obj *_APIOOrdersMgr) WithIsReturn(isReturn bool) Option {
	return optionFunc(func(o *options) { o.query["is_return"] = isReturn })
}

// WithApplicationID application_id获取 application_id
func (obj *_APIOOrdersMgr) WithApplicationID(applicationID int) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithUserID user_id获取 user_id
func (obj *_APIOOrdersMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// GetByOption 功能选项模式获取
func (obj *_APIOOrdersMgr) GetByOption(opts ...Option) (result APIOOrders, err error) {
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
func (obj *_APIOOrdersMgr) GetByOptions(opts ...Option) (results []*APIOOrders, err error) {
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
func (obj *_APIOOrdersMgr) GetFromID(id int) (result APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIOOrdersMgr) GetBatchFromID(ids []int) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 订单号
func (obj *_APIOOrdersMgr) GetFromOrderNo(orderNo string) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_no = ?", orderNo).Find(&results).Error

	return
}

// GetBatchFromOrderNo 批量唯一主键查找 订单号
func (obj *_APIOOrdersMgr) GetBatchFromOrderNo(orderNos []string) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_no IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 订单状态：1：待付款，2：已付款，3，配送中，4：已取消，5：已完成
func (obj *_APIOOrdersMgr) GetFromStatus(status int) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 订单状态：1：待付款，2：已付款，3，配送中，4：已取消，5：已完成
func (obj *_APIOOrdersMgr) GetBatchFromStatus(statuss []int) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 总数量
func (obj *_APIOOrdersMgr) GetFromAmount(amount int) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量唯一主键查找 总数量
func (obj *_APIOOrdersMgr) GetBatchFromAmount(amounts []int) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount IN (?)", amounts).Find(&results).Error

	return
}

// GetFromTotal 通过total获取内容 总金额
func (obj *_APIOOrdersMgr) GetFromTotal(total float64) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("total = ?", total).Find(&results).Error

	return
}

// GetBatchFromTotal 批量唯一主键查找 总金额
func (obj *_APIOOrdersMgr) GetBatchFromTotal(totals []float64) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("total IN (?)", totals).Find(&results).Error

	return
}

// GetFromPayType 通过pay_type获取内容
func (obj *_APIOOrdersMgr) GetFromPayType(payType bool) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pay_type = ?", payType).Find(&results).Error

	return
}

// GetBatchFromPayType 批量唯一主键查找
func (obj *_APIOOrdersMgr) GetBatchFromPayType(payTypes []bool) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pay_type IN (?)", payTypes).Find(&results).Error

	return
}

// GetFromRmk 通过rmk获取内容 备注
func (obj *_APIOOrdersMgr) GetFromRmk(rmk string) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("rmk = ?", rmk).Find(&results).Error

	return
}

// GetBatchFromRmk 批量唯一主键查找 备注
func (obj *_APIOOrdersMgr) GetBatchFromRmk(rmks []string) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("rmk IN (?)", rmks).Find(&results).Error

	return
}

// GetFromAppType 通过app_type获取内容 订单应用类型，1：小程序，2：床旁
func (obj *_APIOOrdersMgr) GetFromAppType(appType int) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_type = ?", appType).Find(&results).Error

	return
}

// GetBatchFromAppType 批量唯一主键查找 订单应用类型，1：小程序，2：床旁
func (obj *_APIOOrdersMgr) GetBatchFromAppType(appTypes []int) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_type IN (?)", appTypes).Find(&results).Error

	return
}

// GetFromOpenID 通过open_id获取内容 open_id
func (obj *_APIOOrdersMgr) GetFromOpenID(openID string) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("open_id = ?", openID).Find(&results).Error

	return
}

// GetBatchFromOpenID 批量唯一主键查找 open_id
func (obj *_APIOOrdersMgr) GetBatchFromOpenID(openIDs []string) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("open_id IN (?)", openIDs).Find(&results).Error

	return
}

// GetFromTradeType 通过trade_type获取内容 JSAPI、NATIVE、APP
func (obj *_APIOOrdersMgr) GetFromTradeType(tradeType string) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("trade_type = ?", tradeType).Find(&results).Error

	return
}

// GetBatchFromTradeType 批量唯一主键查找 JSAPI、NATIVE、APP
func (obj *_APIOOrdersMgr) GetBatchFromTradeType(tradeTypes []string) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("trade_type IN (?)", tradeTypes).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIOOrdersMgr) GetFromCreateAt(createAt time.Time) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIOOrdersMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIOOrdersMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIOOrdersMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIOOrdersMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIOOrdersMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromIsReturn 通过is_return获取内容 是否退款
func (obj *_APIOOrdersMgr) GetFromIsReturn(isReturn bool) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_return = ?", isReturn).Find(&results).Error

	return
}

// GetBatchFromIsReturn 批量唯一主键查找 是否退款
func (obj *_APIOOrdersMgr) GetBatchFromIsReturn(isReturns []bool) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_return IN (?)", isReturns).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容 application_id
func (obj *_APIOOrdersMgr) GetFromApplicationID(applicationID int) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量唯一主键查找 application_id
func (obj *_APIOOrdersMgr) GetBatchFromApplicationID(applicationIDs []int) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id IN (?)", applicationIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 user_id
func (obj *_APIOOrdersMgr) GetFromUserID(userID int) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 user_id
func (obj *_APIOOrdersMgr) GetBatchFromUserID(userIDs []int) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIOOrdersMgr) FetchByPrimaryKey(id int) (result APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_APIOOrdersMgr) FetchIndexByUserID(orderNo string, applicationID int, userID int) (results []*APIOOrders, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_no = ? AND application_id = ? AND user_id = ?", orderNo, applicationID, userID).Find(&results).Error

	return
}
