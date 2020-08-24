package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CareCareTagMgr struct {
	*_BaseMgr
}

// CareCareTagMgr open func
func CareCareTagMgr(db *gorm.DB) *_CareCareTagMgr {
	if db == nil {
		panic(fmt.Errorf("CareCareTagMgr need init by db"))
	}
	return &_CareCareTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("care_care_tag"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CareCareTagMgr) GetTableName() string {
	return "care_care_tag"
}

// Get 获取
func (obj *_CareCareTagMgr) Get() (result CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CareCareTagMgr) Gets() (results []*CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CareCareTagMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCareID care_id获取 care_id
func (obj *_CareCareTagMgr) WithCareID(careID int) Option {
	return optionFunc(func(o *options) { o.query["care_id"] = careID })
}

// WithCareTagID care_tag_id获取 care_tag_id
func (obj *_CareCareTagMgr) WithCareTagID(careTagID int) Option {
	return optionFunc(func(o *options) { o.query["care_tag_id"] = careTagID })
}

// WithCreateAt create_at获取
func (obj *_CareCareTagMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CareCareTagMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_CareCareTagMgr) GetByOption(opts ...Option) (result CareCareTag, err error) {
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
func (obj *_CareCareTagMgr) GetByOptions(opts ...Option) (results []*CareCareTag, err error) {
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
func (obj *_CareCareTagMgr) GetFromID(id int) (result CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CareCareTagMgr) GetBatchFromID(ids []int) (results []*CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCareID 通过care_id获取内容 care_id
func (obj *_CareCareTagMgr) GetFromCareID(careID int) (results []*CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_id = ?", careID).Find(&results).Error

	return
}

// GetBatchFromCareID 批量唯一主键查找 care_id
func (obj *_CareCareTagMgr) GetBatchFromCareID(careIDs []int) (results []*CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_id IN (?)", careIDs).Find(&results).Error

	return
}

// GetFromCareTagID 通过care_tag_id获取内容 care_tag_id
func (obj *_CareCareTagMgr) GetFromCareTagID(careTagID int) (results []*CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_tag_id = ?", careTagID).Find(&results).Error

	return
}

// GetBatchFromCareTagID 批量唯一主键查找 care_tag_id
func (obj *_CareCareTagMgr) GetBatchFromCareTagID(careTagIDs []int) (results []*CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_tag_id IN (?)", careTagIDs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CareCareTagMgr) GetFromCreateAt(createAt time.Time) (results []*CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CareCareTagMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CareCareTagMgr) GetFromUpdateAt(updateAt time.Time) (results []*CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CareCareTagMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CareCareTagMgr) FetchByPrimaryKey(id int) (result CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByCareID  获取多个内容
func (obj *_CareCareTagMgr) FetchIndexByCareID(careID int, careTagID int) (results []*CareCareTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_id = ? AND care_tag_id = ?", careID, careTagID).Find(&results).Error

	return
}
