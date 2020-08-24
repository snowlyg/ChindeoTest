package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CarerCarerTagMgr struct {
	*_BaseMgr
}

// CarerCarerTagMgr open func
func CarerCarerTagMgr(db *gorm.DB) *_CarerCarerTagMgr {
	if db == nil {
		panic(fmt.Errorf("CarerCarerTagMgr need init by db"))
	}
	return &_CarerCarerTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("carer_carer_tag"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CarerCarerTagMgr) GetTableName() string {
	return "carer_carer_tag"
}

// Get 获取
func (obj *_CarerCarerTagMgr) Get() (result CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CarerCarerTagMgr) Gets() (results []*CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CarerCarerTagMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCarerID carer_id获取 carer_id
func (obj *_CarerCarerTagMgr) WithCarerID(carerID int) Option {
	return optionFunc(func(o *options) { o.query["carer_id"] = carerID })
}

// WithCarerTagID carer_tag_id获取 carer_tag_id
func (obj *_CarerCarerTagMgr) WithCarerTagID(carerTagID int) Option {
	return optionFunc(func(o *options) { o.query["carer_tag_id"] = carerTagID })
}

// WithCreateAt create_at获取
func (obj *_CarerCarerTagMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CarerCarerTagMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_CarerCarerTagMgr) GetByOption(opts ...Option) (result CarerCarerTag, err error) {
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
func (obj *_CarerCarerTagMgr) GetByOptions(opts ...Option) (results []*CarerCarerTag, err error) {
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
func (obj *_CarerCarerTagMgr) GetFromID(id int) (result CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CarerCarerTagMgr) GetBatchFromID(ids []int) (results []*CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCarerID 通过carer_id获取内容 carer_id
func (obj *_CarerCarerTagMgr) GetFromCarerID(carerID int) (results []*CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_id = ?", carerID).Find(&results).Error

	return
}

// GetBatchFromCarerID 批量唯一主键查找 carer_id
func (obj *_CarerCarerTagMgr) GetBatchFromCarerID(carerIDs []int) (results []*CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_id IN (?)", carerIDs).Find(&results).Error

	return
}

// GetFromCarerTagID 通过carer_tag_id获取内容 carer_tag_id
func (obj *_CarerCarerTagMgr) GetFromCarerTagID(carerTagID int) (results []*CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_tag_id = ?", carerTagID).Find(&results).Error

	return
}

// GetBatchFromCarerTagID 批量唯一主键查找 carer_tag_id
func (obj *_CarerCarerTagMgr) GetBatchFromCarerTagID(carerTagIDs []int) (results []*CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_tag_id IN (?)", carerTagIDs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CarerCarerTagMgr) GetFromCreateAt(createAt time.Time) (results []*CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CarerCarerTagMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CarerCarerTagMgr) GetFromUpdateAt(updateAt time.Time) (results []*CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CarerCarerTagMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CarerCarerTagMgr) FetchByPrimaryKey(id int) (result CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByCarerID  获取多个内容
func (obj *_CarerCarerTagMgr) FetchIndexByCarerID(carerID int, carerTagID int) (results []*CarerCarerTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_id = ? AND carer_tag_id = ?", carerID, carerTagID).Find(&results).Error

	return
}
