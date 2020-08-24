package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CareTypesMgr struct {
	*_BaseMgr
}

// CareTypesMgr open func
func CareTypesMgr(db *gorm.DB) *_CareTypesMgr {
	if db == nil {
		panic(fmt.Errorf("CareTypesMgr need init by db"))
	}
	return &_CareTypesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("care_types"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CareTypesMgr) GetTableName() string {
	return "care_types"
}

// Get 获取
func (obj *_CareTypesMgr) Get() (result CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CareTypesMgr) Gets() (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CareTypesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_CareTypesMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithEnName en_name获取 英文名称
func (obj *_CareTypesMgr) WithEnName(enName string) Option {
	return optionFunc(func(o *options) { o.query["en_name"] = enName })
}

// WithStatus status获取 状态：启用，禁用
func (obj *_CareTypesMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取
func (obj *_CareTypesMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CareTypesMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_CareTypesMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithIcon icon获取 图标
func (obj *_CareTypesMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// GetByOption 功能选项模式获取
func (obj *_CareTypesMgr) GetByOption(opts ...Option) (result CareTypes, err error) {
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
func (obj *_CareTypesMgr) GetByOptions(opts ...Option) (results []*CareTypes, err error) {
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
func (obj *_CareTypesMgr) GetFromID(id int) (result CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CareTypesMgr) GetBatchFromID(ids []int) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_CareTypesMgr) GetFromName(name string) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_CareTypesMgr) GetBatchFromName(names []string) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromEnName 通过en_name获取内容 英文名称
func (obj *_CareTypesMgr) GetFromEnName(enName string) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("en_name = ?", enName).Find(&results).Error

	return
}

// GetBatchFromEnName 批量唯一主键查找 英文名称
func (obj *_CareTypesMgr) GetBatchFromEnName(enNames []string) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("en_name IN (?)", enNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态：启用，禁用
func (obj *_CareTypesMgr) GetFromStatus(status bool) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态：启用，禁用
func (obj *_CareTypesMgr) GetBatchFromStatus(statuss []bool) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CareTypesMgr) GetFromCreateAt(createAt time.Time) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CareTypesMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CareTypesMgr) GetFromUpdateAt(updateAt time.Time) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CareTypesMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CareTypesMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CareTypesMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 图标
func (obj *_CareTypesMgr) GetFromIcon(icon string) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找 图标
func (obj *_CareTypesMgr) GetBatchFromIcon(icons []string) (results []*CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CareTypesMgr) FetchByPrimaryKey(id int) (result CareTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
