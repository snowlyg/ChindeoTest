package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _MenuMenuTagMgr struct {
	*_BaseMgr
}

// MenuMenuTagMgr open func
func MenuMenuTagMgr(db *gorm.DB) *_MenuMenuTagMgr {
	if db == nil {
		panic(fmt.Errorf("MenuMenuTagMgr need init by db"))
	}
	return &_MenuMenuTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("menu_menu_tag"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MenuMenuTagMgr) GetTableName() string {
	return "menu_menu_tag"
}

// Get 获取
func (obj *_MenuMenuTagMgr) Get() (result MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MenuMenuTagMgr) Gets() (results []*MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MenuMenuTagMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMenuID menu_id获取 menu_id
func (obj *_MenuMenuTagMgr) WithMenuID(menuID int) Option {
	return optionFunc(func(o *options) { o.query["menu_id"] = menuID })
}

// WithMenuTagID menu_tag_id获取 menu_tag_id
func (obj *_MenuMenuTagMgr) WithMenuTagID(menuTagID int) Option {
	return optionFunc(func(o *options) { o.query["menu_tag_id"] = menuTagID })
}

// WithCreateAt create_at获取
func (obj *_MenuMenuTagMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_MenuMenuTagMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_MenuMenuTagMgr) GetByOption(opts ...Option) (result MenuMenuTag, err error) {
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
func (obj *_MenuMenuTagMgr) GetByOptions(opts ...Option) (results []*MenuMenuTag, err error) {
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
func (obj *_MenuMenuTagMgr) GetFromID(id int) (result MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_MenuMenuTagMgr) GetBatchFromID(ids []int) (results []*MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromMenuID 通过menu_id获取内容 menu_id
func (obj *_MenuMenuTagMgr) GetFromMenuID(menuID int) (results []*MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_id = ?", menuID).Find(&results).Error

	return
}

// GetBatchFromMenuID 批量唯一主键查找 menu_id
func (obj *_MenuMenuTagMgr) GetBatchFromMenuID(menuIDs []int) (results []*MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_id IN (?)", menuIDs).Find(&results).Error

	return
}

// GetFromMenuTagID 通过menu_tag_id获取内容 menu_tag_id
func (obj *_MenuMenuTagMgr) GetFromMenuTagID(menuTagID int) (results []*MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_tag_id = ?", menuTagID).Find(&results).Error

	return
}

// GetBatchFromMenuTagID 批量唯一主键查找 menu_tag_id
func (obj *_MenuMenuTagMgr) GetBatchFromMenuTagID(menuTagIDs []int) (results []*MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_tag_id IN (?)", menuTagIDs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_MenuMenuTagMgr) GetFromCreateAt(createAt time.Time) (results []*MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_MenuMenuTagMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_MenuMenuTagMgr) GetFromUpdateAt(updateAt time.Time) (results []*MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_MenuMenuTagMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_MenuMenuTagMgr) FetchByPrimaryKey(id int) (result MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByMenuID  获取多个内容
func (obj *_MenuMenuTagMgr) FetchIndexByMenuID(menuID int, menuTagID int) (results []*MenuMenuTag, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_id = ? AND menu_tag_id = ?", menuID, menuTagID).Find(&results).Error

	return
}
