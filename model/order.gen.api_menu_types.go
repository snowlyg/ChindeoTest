package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIMenuTypesMgr struct {
	*_BaseMgr
}

// APIMenuTypesMgr open func
func APIMenuTypesMgr(db *gorm.DB) *_APIMenuTypesMgr {
	if db == nil {
		panic(fmt.Errorf("APIMenuTypesMgr need init by db"))
	}
	return &_APIMenuTypesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_menu_types"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIMenuTypesMgr) GetTableName() string {
	return "api_menu_types"
}

// Get 获取
func (obj *_APIMenuTypesMgr) Get() (result APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIMenuTypesMgr) Gets() (results []*APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIMenuTypesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 菜品类型名称
func (obj *_APIMenuTypesMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCreateAt create_at获取
func (obj *_APIMenuTypesMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIMenuTypesMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIMenuTypesMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithApplicationID application_id获取
func (obj *_APIMenuTypesMgr) WithApplicationID(applicationID int) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// GetByOption 功能选项模式获取
func (obj *_APIMenuTypesMgr) GetByOption(opts ...Option) (result APIMenuTypes, err error) {
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
func (obj *_APIMenuTypesMgr) GetByOptions(opts ...Option) (results []*APIMenuTypes, err error) {
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
func (obj *_APIMenuTypesMgr) GetFromID(id int) (result APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIMenuTypesMgr) GetBatchFromID(ids []int) (results []*APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 菜品类型名称
func (obj *_APIMenuTypesMgr) GetFromName(name string) (results []*APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 菜品类型名称
func (obj *_APIMenuTypesMgr) GetBatchFromName(names []string) (results []*APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIMenuTypesMgr) GetFromCreateAt(createAt time.Time) (results []*APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIMenuTypesMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIMenuTypesMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIMenuTypesMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIMenuTypesMgr) GetFromIsDeleted(isDeleted time.Time) (results []*APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIMenuTypesMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容
func (obj *_APIMenuTypesMgr) GetFromApplicationID(applicationID int) (results []*APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量唯一主键查找
func (obj *_APIMenuTypesMgr) GetBatchFromApplicationID(applicationIDs []int) (results []*APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id IN (?)", applicationIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIMenuTypesMgr) FetchByPrimaryKey(id int) (result APIMenuTypes, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
