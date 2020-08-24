package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CareTagsMgr struct {
	*_BaseMgr
}

// CareTagsMgr open func
func CareTagsMgr(db *gorm.DB) *_CareTagsMgr {
	if db == nil {
		panic(fmt.Errorf("CareTagsMgr need init by db"))
	}
	return &_CareTagsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("care_tags"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CareTagsMgr) GetTableName() string {
	return "care_tags"
}

// Get 获取
func (obj *_CareTagsMgr) Get() (result CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CareTagsMgr) Gets() (results []*CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CareTagsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_CareTagsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIsDeleted is_deleted获取
func (obj *_CareTagsMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithCreateAt create_at获取
func (obj *_CareTagsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CareTagsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIcon icon获取 图标
func (obj *_CareTagsMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// GetByOption 功能选项模式获取
func (obj *_CareTagsMgr) GetByOption(opts ...Option) (result CareTags, err error) {
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
func (obj *_CareTagsMgr) GetByOptions(opts ...Option) (results []*CareTags, err error) {
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
func (obj *_CareTagsMgr) GetFromID(id int) (result CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CareTagsMgr) GetBatchFromID(ids []int) (results []*CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_CareTagsMgr) GetFromName(name string) (results []*CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_CareTagsMgr) GetBatchFromName(names []string) (results []*CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CareTagsMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CareTagsMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CareTagsMgr) GetFromCreateAt(createAt time.Time) (results []*CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CareTagsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CareTagsMgr) GetFromUpdateAt(updateAt time.Time) (results []*CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CareTagsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 图标
func (obj *_CareTagsMgr) GetFromIcon(icon string) (results []*CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找 图标
func (obj *_CareTagsMgr) GetBatchFromIcon(icons []string) (results []*CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CareTagsMgr) FetchByPrimaryKey(id int) (result CareTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
