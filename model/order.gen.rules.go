package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _RulesMgr struct {
	*_BaseMgr
}

// RulesMgr open func
func RulesMgr(db *gorm.DB) *_RulesMgr {
	if db == nil {
		panic(fmt.Errorf("RulesMgr need init by db"))
	}
	return &_RulesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("rules"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RulesMgr) GetTableName() string {
	return "rules"
}

// Get 获取
func (obj *_RulesMgr) Get() (result Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RulesMgr) Gets() (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_RulesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPtype ptype获取
func (obj *_RulesMgr) WithPtype(ptype string) Option {
	return optionFunc(func(o *options) { o.query["ptype"] = ptype })
}

// WithV0 v0获取
func (obj *_RulesMgr) WithV0(v0 string) Option {
	return optionFunc(func(o *options) { o.query["v0"] = v0 })
}

// WithV1 v1获取
func (obj *_RulesMgr) WithV1(v1 string) Option {
	return optionFunc(func(o *options) { o.query["v1"] = v1 })
}

// WithV2 v2获取
func (obj *_RulesMgr) WithV2(v2 string) Option {
	return optionFunc(func(o *options) { o.query["v2"] = v2 })
}

// WithV3 v3获取
func (obj *_RulesMgr) WithV3(v3 string) Option {
	return optionFunc(func(o *options) { o.query["v3"] = v3 })
}

// WithV4 v4获取
func (obj *_RulesMgr) WithV4(v4 string) Option {
	return optionFunc(func(o *options) { o.query["v4"] = v4 })
}

// WithV5 v5获取
func (obj *_RulesMgr) WithV5(v5 string) Option {
	return optionFunc(func(o *options) { o.query["v5"] = v5 })
}

// GetByOption 功能选项模式获取
func (obj *_RulesMgr) GetByOption(opts ...Option) (result Rules, err error) {
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
func (obj *_RulesMgr) GetByOptions(opts ...Option) (results []*Rules, err error) {
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
func (obj *_RulesMgr) GetFromID(id int) (result Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_RulesMgr) GetBatchFromID(ids []int) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromPtype 通过ptype获取内容
func (obj *_RulesMgr) GetFromPtype(ptype string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ptype = ?", ptype).Find(&results).Error

	return
}

// GetBatchFromPtype 批量唯一主键查找
func (obj *_RulesMgr) GetBatchFromPtype(ptypes []string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("ptype IN (?)", ptypes).Find(&results).Error

	return
}

// GetFromV0 通过v0获取内容
func (obj *_RulesMgr) GetFromV0(v0 string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("v0 = ?", v0).Find(&results).Error

	return
}

// GetBatchFromV0 批量唯一主键查找
func (obj *_RulesMgr) GetBatchFromV0(v0s []string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("v0 IN (?)", v0s).Find(&results).Error

	return
}

// GetFromV1 通过v1获取内容
func (obj *_RulesMgr) GetFromV1(v1 string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("v1 = ?", v1).Find(&results).Error

	return
}

// GetBatchFromV1 批量唯一主键查找
func (obj *_RulesMgr) GetBatchFromV1(v1s []string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("v1 IN (?)", v1s).Find(&results).Error

	return
}

// GetFromV2 通过v2获取内容
func (obj *_RulesMgr) GetFromV2(v2 string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("v2 = ?", v2).Find(&results).Error

	return
}

// GetBatchFromV2 批量唯一主键查找
func (obj *_RulesMgr) GetBatchFromV2(v2s []string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("v2 IN (?)", v2s).Find(&results).Error

	return
}

// GetFromV3 通过v3获取内容
func (obj *_RulesMgr) GetFromV3(v3 string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("v3 = ?", v3).Find(&results).Error

	return
}

// GetBatchFromV3 批量唯一主键查找
func (obj *_RulesMgr) GetBatchFromV3(v3s []string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("v3 IN (?)", v3s).Find(&results).Error

	return
}

// GetFromV4 通过v4获取内容
func (obj *_RulesMgr) GetFromV4(v4 string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("v4 = ?", v4).Find(&results).Error

	return
}

// GetBatchFromV4 批量唯一主键查找
func (obj *_RulesMgr) GetBatchFromV4(v4s []string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("v4 IN (?)", v4s).Find(&results).Error

	return
}

// GetFromV5 通过v5获取内容
func (obj *_RulesMgr) GetFromV5(v5 string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("v5 = ?", v5).Find(&results).Error

	return
}

// GetBatchFromV5 批量唯一主键查找
func (obj *_RulesMgr) GetBatchFromV5(v5s []string) (results []*Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("v5 IN (?)", v5s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_RulesMgr) FetchByPrimaryKey(id int) (result Rules, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
