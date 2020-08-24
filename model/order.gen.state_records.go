package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _StateRecordsMgr struct {
	*_BaseMgr
}

// StateRecordsMgr open func
func StateRecordsMgr(db *gorm.DB) *_StateRecordsMgr {
	if db == nil {
		panic(fmt.Errorf("StateRecordsMgr need init by db"))
	}
	return &_StateRecordsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("state_records"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StateRecordsMgr) GetTableName() string {
	return "state_records"
}

// Get 获取
func (obj *_StateRecordsMgr) Get() (result StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StateRecordsMgr) Gets() (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_StateRecordsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStateType state_type获取 对账类型：1：微信，2：支付宝
func (obj *_StateRecordsMgr) WithStateType(stateType int) Option {
	return optionFunc(func(o *options) { o.query["state_type"] = stateType })
}

// WithOrderType order_type获取 订单类型：1：付款，2：退款
func (obj *_StateRecordsMgr) WithOrderType(orderType int) Option {
	return optionFunc(func(o *options) { o.query["order_type"] = orderType })
}

// WithOrderNum order_num获取 订单数量
func (obj *_StateRecordsMgr) WithOrderNum(orderNum int) Option {
	return optionFunc(func(o *options) { o.query["order_num"] = orderNum })
}

// WithTotal total获取 总金额
func (obj *_StateRecordsMgr) WithTotal(total float64) Option {
	return optionFunc(func(o *options) { o.query["total"] = total })
}

// WithRealOrderNum real_order_num获取 支付平台订单数量
func (obj *_StateRecordsMgr) WithRealOrderNum(realOrderNum int) Option {
	return optionFunc(func(o *options) { o.query["real_order_num"] = realOrderNum })
}

// WithRealTotal real_total获取 支付平台总金额
func (obj *_StateRecordsMgr) WithRealTotal(realTotal float64) Option {
	return optionFunc(func(o *options) { o.query["real_total"] = realTotal })
}

// WithIsException is_exception获取 是否异常
func (obj *_StateRecordsMgr) WithIsException(isException bool) Option {
	return optionFunc(func(o *options) { o.query["is_exception"] = isException })
}

// WithOrderNos order_nos获取 异常订单号用 ， 隔开
func (obj *_StateRecordsMgr) WithOrderNos(orderNos string) Option {
	return optionFunc(func(o *options) { o.query["order_nos"] = orderNos })
}

// WithExtraType extra_type获取 平账方式：1：盈余，2：亏损
func (obj *_StateRecordsMgr) WithExtraType(extraType int) Option {
	return optionFunc(func(o *options) { o.query["extra_type"] = extraType })
}

// WithExtra extra获取 平账金额
func (obj *_StateRecordsMgr) WithExtra(extra float64) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithStartAt start_at获取 对账订单开始时间范围
func (obj *_StateRecordsMgr) WithStartAt(startAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_at"] = startAt })
}

// WithEndAt end_at获取 对账订单结束时间范围
func (obj *_StateRecordsMgr) WithEndAt(endAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_at"] = endAt })
}

// WithUsername username获取 平账人
func (obj *_StateRecordsMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithRmk rmk获取 平账备注
func (obj *_StateRecordsMgr) WithRmk(rmk string) Option {
	return optionFunc(func(o *options) { o.query["rmk"] = rmk })
}

// WithRecordAt record_at获取 对账时间
func (obj *_StateRecordsMgr) WithRecordAt(recordAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["record_at"] = recordAt })
}

// WithCreateAt create_at获取
func (obj *_StateRecordsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_StateRecordsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_StateRecordsMgr) WithIsDeleted(isDeleted bool) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// GetByOption 功能选项模式获取
func (obj *_StateRecordsMgr) GetByOption(opts ...Option) (result StateRecords, err error) {
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
func (obj *_StateRecordsMgr) GetByOptions(opts ...Option) (results []*StateRecords, err error) {
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
func (obj *_StateRecordsMgr) GetFromID(id int) (result StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_StateRecordsMgr) GetBatchFromID(ids []int) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromStateType 通过state_type获取内容 对账类型：1：微信，2：支付宝
func (obj *_StateRecordsMgr) GetFromStateType(stateType int) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_type = ?", stateType).Find(&results).Error

	return
}

// GetBatchFromStateType 批量唯一主键查找 对账类型：1：微信，2：支付宝
func (obj *_StateRecordsMgr) GetBatchFromStateType(stateTypes []int) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("state_type IN (?)", stateTypes).Find(&results).Error

	return
}

// GetFromOrderType 通过order_type获取内容 订单类型：1：付款，2：退款
func (obj *_StateRecordsMgr) GetFromOrderType(orderType int) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_type = ?", orderType).Find(&results).Error

	return
}

// GetBatchFromOrderType 批量唯一主键查找 订单类型：1：付款，2：退款
func (obj *_StateRecordsMgr) GetBatchFromOrderType(orderTypes []int) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_type IN (?)", orderTypes).Find(&results).Error

	return
}

// GetFromOrderNum 通过order_num获取内容 订单数量
func (obj *_StateRecordsMgr) GetFromOrderNum(orderNum int) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_num = ?", orderNum).Find(&results).Error

	return
}

// GetBatchFromOrderNum 批量唯一主键查找 订单数量
func (obj *_StateRecordsMgr) GetBatchFromOrderNum(orderNums []int) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_num IN (?)", orderNums).Find(&results).Error

	return
}

// GetFromTotal 通过total获取内容 总金额
func (obj *_StateRecordsMgr) GetFromTotal(total float64) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("total = ?", total).Find(&results).Error

	return
}

// GetBatchFromTotal 批量唯一主键查找 总金额
func (obj *_StateRecordsMgr) GetBatchFromTotal(totals []float64) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("total IN (?)", totals).Find(&results).Error

	return
}

// GetFromRealOrderNum 通过real_order_num获取内容 支付平台订单数量
func (obj *_StateRecordsMgr) GetFromRealOrderNum(realOrderNum int) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("real_order_num = ?", realOrderNum).Find(&results).Error

	return
}

// GetBatchFromRealOrderNum 批量唯一主键查找 支付平台订单数量
func (obj *_StateRecordsMgr) GetBatchFromRealOrderNum(realOrderNums []int) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("real_order_num IN (?)", realOrderNums).Find(&results).Error

	return
}

// GetFromRealTotal 通过real_total获取内容 支付平台总金额
func (obj *_StateRecordsMgr) GetFromRealTotal(realTotal float64) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("real_total = ?", realTotal).Find(&results).Error

	return
}

// GetBatchFromRealTotal 批量唯一主键查找 支付平台总金额
func (obj *_StateRecordsMgr) GetBatchFromRealTotal(realTotals []float64) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("real_total IN (?)", realTotals).Find(&results).Error

	return
}

// GetFromIsException 通过is_exception获取内容 是否异常
func (obj *_StateRecordsMgr) GetFromIsException(isException bool) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_exception = ?", isException).Find(&results).Error

	return
}

// GetBatchFromIsException 批量唯一主键查找 是否异常
func (obj *_StateRecordsMgr) GetBatchFromIsException(isExceptions []bool) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_exception IN (?)", isExceptions).Find(&results).Error

	return
}

// GetFromOrderNos 通过order_nos获取内容 异常订单号用 ， 隔开
func (obj *_StateRecordsMgr) GetFromOrderNos(orderNos string) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_nos = ?", orderNos).Find(&results).Error

	return
}

// GetBatchFromOrderNos 批量唯一主键查找 异常订单号用 ， 隔开
func (obj *_StateRecordsMgr) GetBatchFromOrderNos(orderNoss []string) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("order_nos IN (?)", orderNoss).Find(&results).Error

	return
}

// GetFromExtraType 通过extra_type获取内容 平账方式：1：盈余，2：亏损
func (obj *_StateRecordsMgr) GetFromExtraType(extraType int) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("extra_type = ?", extraType).Find(&results).Error

	return
}

// GetBatchFromExtraType 批量唯一主键查找 平账方式：1：盈余，2：亏损
func (obj *_StateRecordsMgr) GetBatchFromExtraType(extraTypes []int) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("extra_type IN (?)", extraTypes).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 平账金额
func (obj *_StateRecordsMgr) GetFromExtra(extra float64) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("extra = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量唯一主键查找 平账金额
func (obj *_StateRecordsMgr) GetBatchFromExtra(extras []float64) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("extra IN (?)", extras).Find(&results).Error

	return
}

// GetFromStartAt 通过start_at获取内容 对账订单开始时间范围
func (obj *_StateRecordsMgr) GetFromStartAt(startAt time.Time) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("start_at = ?", startAt).Find(&results).Error

	return
}

// GetBatchFromStartAt 批量唯一主键查找 对账订单开始时间范围
func (obj *_StateRecordsMgr) GetBatchFromStartAt(startAts []time.Time) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("start_at IN (?)", startAts).Find(&results).Error

	return
}

// GetFromEndAt 通过end_at获取内容 对账订单结束时间范围
func (obj *_StateRecordsMgr) GetFromEndAt(endAt time.Time) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("end_at = ?", endAt).Find(&results).Error

	return
}

// GetBatchFromEndAt 批量唯一主键查找 对账订单结束时间范围
func (obj *_StateRecordsMgr) GetBatchFromEndAt(endAts []time.Time) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("end_at IN (?)", endAts).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 平账人
func (obj *_StateRecordsMgr) GetFromUsername(username string) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量唯一主键查找 平账人
func (obj *_StateRecordsMgr) GetBatchFromUsername(usernames []string) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username IN (?)", usernames).Find(&results).Error

	return
}

// GetFromRmk 通过rmk获取内容 平账备注
func (obj *_StateRecordsMgr) GetFromRmk(rmk string) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("rmk = ?", rmk).Find(&results).Error

	return
}

// GetBatchFromRmk 批量唯一主键查找 平账备注
func (obj *_StateRecordsMgr) GetBatchFromRmk(rmks []string) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("rmk IN (?)", rmks).Find(&results).Error

	return
}

// GetFromRecordAt 通过record_at获取内容 对账时间
func (obj *_StateRecordsMgr) GetFromRecordAt(recordAt time.Time) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("record_at = ?", recordAt).Find(&results).Error

	return
}

// GetBatchFromRecordAt 批量唯一主键查找 对账时间
func (obj *_StateRecordsMgr) GetBatchFromRecordAt(recordAts []time.Time) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("record_at IN (?)", recordAts).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_StateRecordsMgr) GetFromCreateAt(createAt time.Time) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_StateRecordsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_StateRecordsMgr) GetFromUpdateAt(updateAt time.Time) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_StateRecordsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_StateRecordsMgr) GetFromIsDeleted(isDeleted bool) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_StateRecordsMgr) GetBatchFromIsDeleted(isDeleteds []bool) (results []*StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_StateRecordsMgr) FetchByPrimaryKey(id int) (result StateRecords, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
