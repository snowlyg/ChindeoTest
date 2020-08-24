package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _UserMenuMgr struct {
	*_BaseMgr
}

// UserMenuMgr open func
func UserMenuMgr(db *gorm.DB) *_UserMenuMgr {
	if db == nil {
		panic(fmt.Errorf("UserMenuMgr need init by db"))
	}
	return &_UserMenuMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_menu"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMenuMgr) GetTableName() string {
	return "user_menu"
}

// Get 获取
func (obj *_UserMenuMgr) Get() (result UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMenuMgr) Gets() (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserMenuMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMenuID menu_id获取 menu_id
func (obj *_UserMenuMgr) WithMenuID(menuID int) Option {
	return optionFunc(func(o *options) { o.query["menu_id"] = menuID })
}

// WithUserID user_id获取 user_id
func (obj *_UserMenuMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithType type获取 类型: 默认 1 收藏
func (obj *_UserMenuMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCreateAt create_at获取
func (obj *_UserMenuMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_UserMenuMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_UserMenuMgr) GetByOption(opts ...Option) (result UserMenu, err error) {
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
func (obj *_UserMenuMgr) GetByOptions(opts ...Option) (results []*UserMenu, err error) {
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
func (obj *_UserMenuMgr) GetFromID(id int) (result UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserMenuMgr) GetBatchFromID(ids []int) (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromMenuID 通过menu_id获取内容 menu_id
func (obj *_UserMenuMgr) GetFromMenuID(menuID int) (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_id = ?", menuID).Find(&results).Error

	return
}

// GetBatchFromMenuID 批量唯一主键查找 menu_id
func (obj *_UserMenuMgr) GetBatchFromMenuID(menuIDs []int) (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_id IN (?)", menuIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 user_id
func (obj *_UserMenuMgr) GetFromUserID(userID int) (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 user_id
func (obj *_UserMenuMgr) GetBatchFromUserID(userIDs []int) (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型: 默认 1 收藏
func (obj *_UserMenuMgr) GetFromType(_type int) (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 类型: 默认 1 收藏
func (obj *_UserMenuMgr) GetBatchFromType(_types []int) (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_UserMenuMgr) GetFromCreateAt(createAt time.Time) (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_UserMenuMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_UserMenuMgr) GetFromUpdateAt(updateAt time.Time) (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_UserMenuMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserMenuMgr) FetchByPrimaryKey(id int) (result UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByMenuID  获取多个内容
func (obj *_UserMenuMgr) FetchIndexByMenuID(menuID int, userID int) (results []*UserMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_id = ? AND user_id = ?", menuID, userID).Find(&results).Error

	return
}
