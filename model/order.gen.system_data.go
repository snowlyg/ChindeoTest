package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _SystemDataMgr struct {
	*_BaseMgr
}

// SystemDataMgr open func
func SystemDataMgr(db *gorm.DB) *_SystemDataMgr {
	if db == nil {
		panic(fmt.Errorf("SystemDataMgr need init by db"))
	}
	return &_SystemDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_data"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemDataMgr) GetTableName() string {
	return "system_data"
}

// Get 获取
func (obj *_SystemDataMgr) Get() (result SystemData, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemDataMgr) Gets() (results []*SystemData, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SystemDataMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 配置名
func (obj *_SystemDataMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithValue value获取 配置值
func (obj *_SystemDataMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// GetByOption 功能选项模式获取
func (obj *_SystemDataMgr) GetByOption(opts ...Option) (result SystemData, err error) {
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
func (obj *_SystemDataMgr) GetByOptions(opts ...Option) (results []*SystemData, err error) {
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
func (obj *_SystemDataMgr) GetFromID(id uint64) (result SystemData, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_SystemDataMgr) GetBatchFromID(ids []uint64) (results []*SystemData, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 配置名
func (obj *_SystemDataMgr) GetFromName(name string) (results []*SystemData, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 配置名
func (obj *_SystemDataMgr) GetBatchFromName(names []string) (results []*SystemData, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 配置值
func (obj *_SystemDataMgr) GetFromValue(value string) (results []*SystemData, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量唯一主键查找 配置值
func (obj *_SystemDataMgr) GetBatchFromValue(values []string) (results []*SystemData, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SystemDataMgr) FetchByPrimaryKey(id uint64) (result SystemData, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxSystemDataName  获取多个内容
func (obj *_SystemDataMgr) FetchIndexByIDxSystemDataName(name string) (results []*SystemData, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
