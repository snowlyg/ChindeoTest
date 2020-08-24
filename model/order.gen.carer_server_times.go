package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CarerServerTimesMgr struct {
	*_BaseMgr
}

// CarerServerTimesMgr open func
func CarerServerTimesMgr(db *gorm.DB) *_CarerServerTimesMgr {
	if db == nil {
		panic(fmt.Errorf("CarerServerTimesMgr need init by db"))
	}
	return &_CarerServerTimesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("carer_server_times"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CarerServerTimesMgr) GetTableName() string {
	return "carer_server_times"
}

// Get 获取
func (obj *_CarerServerTimesMgr) Get() (result CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CarerServerTimesMgr) Gets() (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CarerServerTimesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStartAt start_at获取 服务开始时间
func (obj *_CarerServerTimesMgr) WithStartAt(startAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_at"] = startAt })
}

// WithEndAt end_at获取 服务结束时间
func (obj *_CarerServerTimesMgr) WithEndAt(endAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_at"] = endAt })
}

// WithCarerID carer_id获取 carer_id
func (obj *_CarerServerTimesMgr) WithCarerID(carerID int) Option {
	return optionFunc(func(o *options) { o.query["carer_id"] = carerID })
}

// WithTimeType time_type获取 时间类型：时，天
func (obj *_CarerServerTimesMgr) WithTimeType(timeType string) Option {
	return optionFunc(func(o *options) { o.query["time_type"] = timeType })
}

// WithCareOrderID care_order_id获取 care_order_id
func (obj *_CarerServerTimesMgr) WithCareOrderID(careOrderID int) Option {
	return optionFunc(func(o *options) { o.query["care_order_id"] = careOrderID })
}

// WithCreateAt create_at获取
func (obj *_CarerServerTimesMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CarerServerTimesMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_CarerServerTimesMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// GetByOption 功能选项模式获取
func (obj *_CarerServerTimesMgr) GetByOption(opts ...Option) (result CarerServerTimes, err error) {
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
func (obj *_CarerServerTimesMgr) GetByOptions(opts ...Option) (results []*CarerServerTimes, err error) {
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
func (obj *_CarerServerTimesMgr) GetFromID(id int) (result CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CarerServerTimesMgr) GetBatchFromID(ids []int) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromStartAt 通过start_at获取内容 服务开始时间
func (obj *_CarerServerTimesMgr) GetFromStartAt(startAt time.Time) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("start_at = ?", startAt).Find(&results).Error

	return
}

// GetBatchFromStartAt 批量唯一主键查找 服务开始时间
func (obj *_CarerServerTimesMgr) GetBatchFromStartAt(startAts []time.Time) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("start_at IN (?)", startAts).Find(&results).Error

	return
}

// GetFromEndAt 通过end_at获取内容 服务结束时间
func (obj *_CarerServerTimesMgr) GetFromEndAt(endAt time.Time) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("end_at = ?", endAt).Find(&results).Error

	return
}

// GetBatchFromEndAt 批量唯一主键查找 服务结束时间
func (obj *_CarerServerTimesMgr) GetBatchFromEndAt(endAts []time.Time) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("end_at IN (?)", endAts).Find(&results).Error

	return
}

// GetFromCarerID 通过carer_id获取内容 carer_id
func (obj *_CarerServerTimesMgr) GetFromCarerID(carerID int) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_id = ?", carerID).Find(&results).Error

	return
}

// GetBatchFromCarerID 批量唯一主键查找 carer_id
func (obj *_CarerServerTimesMgr) GetBatchFromCarerID(carerIDs []int) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_id IN (?)", carerIDs).Find(&results).Error

	return
}

// GetFromTimeType 通过time_type获取内容 时间类型：时，天
func (obj *_CarerServerTimesMgr) GetFromTimeType(timeType string) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type = ?", timeType).Find(&results).Error

	return
}

// GetBatchFromTimeType 批量唯一主键查找 时间类型：时，天
func (obj *_CarerServerTimesMgr) GetBatchFromTimeType(timeTypes []string) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type IN (?)", timeTypes).Find(&results).Error

	return
}

// GetFromCareOrderID 通过care_order_id获取内容 care_order_id
func (obj *_CarerServerTimesMgr) GetFromCareOrderID(careOrderID int) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id = ?", careOrderID).Find(&results).Error

	return
}

// GetBatchFromCareOrderID 批量唯一主键查找 care_order_id
func (obj *_CarerServerTimesMgr) GetBatchFromCareOrderID(careOrderIDs []int) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id IN (?)", careOrderIDs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CarerServerTimesMgr) GetFromCreateAt(createAt time.Time) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CarerServerTimesMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CarerServerTimesMgr) GetFromUpdateAt(updateAt time.Time) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CarerServerTimesMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CarerServerTimesMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CarerServerTimesMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CarerServerTimesMgr) FetchByPrimaryKey(id int) (result CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByCareOrderID  获取多个内容
func (obj *_CarerServerTimesMgr) FetchIndexByCareOrderID(carerID int, careOrderID int) (results []*CarerServerTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_id = ? AND care_order_id = ?", carerID, careOrderID).Find(&results).Error

	return
}
