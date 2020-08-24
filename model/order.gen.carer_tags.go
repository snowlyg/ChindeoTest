package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CarerTagsMgr struct {
	*_BaseMgr
}

// CarerTagsMgr open func
func CarerTagsMgr(db *gorm.DB) *_CarerTagsMgr {
	if db == nil {
		panic(fmt.Errorf("CarerTagsMgr need init by db"))
	}
	return &_CarerTagsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("carer_tags"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CarerTagsMgr) GetTableName() string {
	return "carer_tags"
}

// Get 获取
func (obj *_CarerTagsMgr) Get() (result CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CarerTagsMgr) Gets() (results []*CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CarerTagsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_CarerTagsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCreateAt create_at获取
func (obj *_CarerTagsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CarerTagsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_CarerTagsMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithIcon icon获取 图标
func (obj *_CarerTagsMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// GetByOption 功能选项模式获取
func (obj *_CarerTagsMgr) GetByOption(opts ...Option) (result CarerTags, err error) {
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
func (obj *_CarerTagsMgr) GetByOptions(opts ...Option) (results []*CarerTags, err error) {
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
func (obj *_CarerTagsMgr) GetFromID(id int) (result CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CarerTagsMgr) GetBatchFromID(ids []int) (results []*CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_CarerTagsMgr) GetFromName(name string) (results []*CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_CarerTagsMgr) GetBatchFromName(names []string) (results []*CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CarerTagsMgr) GetFromCreateAt(createAt time.Time) (results []*CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CarerTagsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CarerTagsMgr) GetFromUpdateAt(updateAt time.Time) (results []*CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CarerTagsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CarerTagsMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CarerTagsMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 图标
func (obj *_CarerTagsMgr) GetFromIcon(icon string) (results []*CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找 图标
func (obj *_CarerTagsMgr) GetBatchFromIcon(icons []string) (results []*CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CarerTagsMgr) FetchByPrimaryKey(id int) (result CarerTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
