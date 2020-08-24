package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIMenuTagsMgr struct {
	*_BaseMgr
}

// APIMenuTagsMgr open func
func APIMenuTagsMgr(db *gorm.DB) *_APIMenuTagsMgr {
	if db == nil {
		panic(fmt.Errorf("APIMenuTagsMgr need init by db"))
	}
	return &_APIMenuTagsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_menu_tags"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIMenuTagsMgr) GetTableName() string {
	return "api_menu_tags"
}

// Get 获取
func (obj *_APIMenuTagsMgr) Get() (result APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIMenuTagsMgr) Gets() (results []*APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIMenuTagsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 标签名称
func (obj *_APIMenuTagsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIsDeleted is_deleted获取
func (obj *_APIMenuTagsMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithCreateAt create_at获取
func (obj *_APIMenuTagsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIMenuTagsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_APIMenuTagsMgr) GetByOption(opts ...Option) (result APIMenuTags, err error) {
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
func (obj *_APIMenuTagsMgr) GetByOptions(opts ...Option) (results []*APIMenuTags, err error) {
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
func (obj *_APIMenuTagsMgr) GetFromID(id int) (result APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIMenuTagsMgr) GetBatchFromID(ids []int) (results []*APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 标签名称
func (obj *_APIMenuTagsMgr) GetFromName(name string) (results []*APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 标签名称
func (obj *_APIMenuTagsMgr) GetBatchFromName(names []string) (results []*APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIMenuTagsMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIMenuTagsMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIMenuTagsMgr) GetFromCreateAt(createAt time.Time) (results []*APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIMenuTagsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIMenuTagsMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIMenuTagsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIMenuTagsMgr) FetchByPrimaryKey(id int) (result APIMenuTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
