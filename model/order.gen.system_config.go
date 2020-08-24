package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _SystemConfigMgr struct {
	*_BaseMgr
}

// SystemConfigMgr open func
func SystemConfigMgr(db *gorm.DB) *_SystemConfigMgr {
	if db == nil {
		panic(fmt.Errorf("SystemConfigMgr need init by db"))
	}
	return &_SystemConfigMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_config"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemConfigMgr) GetTableName() string {
	return "system_config"
}

// Get 获取
func (obj *_SystemConfigMgr) Get() (result SystemConfig, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemConfigMgr) Gets() (results []*SystemConfig, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithType type获取 分类
func (obj *_SystemConfigMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithName name获取 配置名
func (obj *_SystemConfigMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithValue value获取 配置值
func (obj *_SystemConfigMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// GetByOption 功能选项模式获取
func (obj *_SystemConfigMgr) GetByOption(opts ...Option) (result SystemConfig, err error) {
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
func (obj *_SystemConfigMgr) GetByOptions(opts ...Option) (results []*SystemConfig, err error) {
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

// GetFromType 通过type获取内容 分类
func (obj *_SystemConfigMgr) GetFromType(_type string) (results []*SystemConfig, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 分类
func (obj *_SystemConfigMgr) GetBatchFromType(_types []string) (results []*SystemConfig, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 配置名
func (obj *_SystemConfigMgr) GetFromName(name string) (results []*SystemConfig, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 配置名
func (obj *_SystemConfigMgr) GetBatchFromName(names []string) (results []*SystemConfig, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 配置值
func (obj *_SystemConfigMgr) GetFromValue(value string) (results []*SystemConfig, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量唯一主键查找 配置值
func (obj *_SystemConfigMgr) GetBatchFromValue(values []string) (results []*SystemConfig, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByIDxSystemConfigType  获取多个内容
func (obj *_SystemConfigMgr) FetchIndexByIDxSystemConfigType(_type string) (results []*SystemConfig, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// FetchIndexByIDxSystemConfigName  获取多个内容
func (obj *_SystemConfigMgr) FetchIndexByIDxSystemConfigName(name string) (results []*SystemConfig, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
