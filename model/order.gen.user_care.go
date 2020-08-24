package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _UserCareMgr struct {
	*_BaseMgr
}

// UserCareMgr open func
func UserCareMgr(db *gorm.DB) *_UserCareMgr {
	if db == nil {
		panic(fmt.Errorf("UserCareMgr need init by db"))
	}
	return &_UserCareMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_care"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserCareMgr) GetTableName() string {
	return "user_care"
}

// Get 获取
func (obj *_UserCareMgr) Get() (result UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserCareMgr) Gets() (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserCareMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCareID care_id获取 care_id
func (obj *_UserCareMgr) WithCareID(careID int) Option {
	return optionFunc(func(o *options) { o.query["care_id"] = careID })
}

// WithUserID user_id获取 user_id
func (obj *_UserCareMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithType type获取 类型: 默认 ,1 收藏
func (obj *_UserCareMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCreateAt create_at获取
func (obj *_UserCareMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_UserCareMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_UserCareMgr) GetByOption(opts ...Option) (result UserCare, err error) {
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
func (obj *_UserCareMgr) GetByOptions(opts ...Option) (results []*UserCare, err error) {
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
func (obj *_UserCareMgr) GetFromID(id int) (result UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserCareMgr) GetBatchFromID(ids []int) (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCareID 通过care_id获取内容 care_id
func (obj *_UserCareMgr) GetFromCareID(careID int) (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_id = ?", careID).Find(&results).Error

	return
}

// GetBatchFromCareID 批量唯一主键查找 care_id
func (obj *_UserCareMgr) GetBatchFromCareID(careIDs []int) (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_id IN (?)", careIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 user_id
func (obj *_UserCareMgr) GetFromUserID(userID int) (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 user_id
func (obj *_UserCareMgr) GetBatchFromUserID(userIDs []int) (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型: 默认 ,1 收藏
func (obj *_UserCareMgr) GetFromType(_type int) (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 类型: 默认 ,1 收藏
func (obj *_UserCareMgr) GetBatchFromType(_types []int) (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_UserCareMgr) GetFromCreateAt(createAt time.Time) (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_UserCareMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_UserCareMgr) GetFromUpdateAt(updateAt time.Time) (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_UserCareMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserCareMgr) FetchByPrimaryKey(id int) (result UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByCareID  获取多个内容
func (obj *_UserCareMgr) FetchIndexByCareID(careID int, userID int) (results []*UserCare, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_id = ? AND user_id = ?", careID, userID).Find(&results).Error

	return
}
