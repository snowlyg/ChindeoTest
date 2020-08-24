package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIWarnTimesMgr struct {
	*_BaseMgr
}

// APIWarnTimesMgr open func
func APIWarnTimesMgr(db *gorm.DB) *_APIWarnTimesMgr {
	if db == nil {
		panic(fmt.Errorf("APIWarnTimesMgr need init by db"))
	}
	return &_APIWarnTimesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_warn_times"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIWarnTimesMgr) GetTableName() string {
	return "api_warn_times"
}

// Get 获取
func (obj *_APIWarnTimesMgr) Get() (result APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIWarnTimesMgr) Gets() (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIWarnTimesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWeeks weeks获取 星期
func (obj *_APIWarnTimesMgr) WithWeeks(weeks string) Option {
	return optionFunc(func(o *options) { o.query["weeks"] = weeks })
}

// WithTime time获取 提醒时间
func (obj *_APIWarnTimesMgr) WithTime(time string) Option {
	return optionFunc(func(o *options) { o.query["time"] = time })
}

// WithStatus status获取 是否启用
func (obj *_APIWarnTimesMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取
func (obj *_APIWarnTimesMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIWarnTimesMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIWarnTimesMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithUserID user_id获取 user_id
func (obj *_APIWarnTimesMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// GetByOption 功能选项模式获取
func (obj *_APIWarnTimesMgr) GetByOption(opts ...Option) (result APIWarnTimes, err error) {
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
func (obj *_APIWarnTimesMgr) GetByOptions(opts ...Option) (results []*APIWarnTimes, err error) {
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
func (obj *_APIWarnTimesMgr) GetFromID(id int) (result APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIWarnTimesMgr) GetBatchFromID(ids []int) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromWeeks 通过weeks获取内容 星期
func (obj *_APIWarnTimesMgr) GetFromWeeks(weeks string) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("weeks = ?", weeks).Find(&results).Error

	return
}

// GetBatchFromWeeks 批量唯一主键查找 星期
func (obj *_APIWarnTimesMgr) GetBatchFromWeeks(weekss []string) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("weeks IN (?)", weekss).Find(&results).Error

	return
}

// GetFromTime 通过time获取内容 提醒时间
func (obj *_APIWarnTimesMgr) GetFromTime(time string) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time = ?", time).Find(&results).Error

	return
}

// GetBatchFromTime 批量唯一主键查找 提醒时间
func (obj *_APIWarnTimesMgr) GetBatchFromTime(times []string) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time IN (?)", times).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 是否启用
func (obj *_APIWarnTimesMgr) GetFromStatus(status bool) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 是否启用
func (obj *_APIWarnTimesMgr) GetBatchFromStatus(statuss []bool) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIWarnTimesMgr) GetFromCreateAt(createAt time.Time) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIWarnTimesMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIWarnTimesMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIWarnTimesMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIWarnTimesMgr) GetFromIsDeleted(isDeleted time.Time) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIWarnTimesMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 user_id
func (obj *_APIWarnTimesMgr) GetFromUserID(userID int) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 user_id
func (obj *_APIWarnTimesMgr) GetBatchFromUserID(userIDs []int) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIWarnTimesMgr) FetchByPrimaryKey(id int) (result APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_APIWarnTimesMgr) FetchIndexByUserID(userID int) (results []*APIWarnTimes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}
