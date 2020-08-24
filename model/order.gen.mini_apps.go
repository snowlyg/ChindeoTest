package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _MiniAppsMgr struct {
	*_BaseMgr
}

// MiniAppsMgr open func
func MiniAppsMgr(db *gorm.DB) *_MiniAppsMgr {
	if db == nil {
		panic(fmt.Errorf("MiniAppsMgr need init by db"))
	}
	return &_MiniAppsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mini_apps"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MiniAppsMgr) GetTableName() string {
	return "mini_apps"
}

// Get 获取
func (obj *_MiniAppsMgr) Get() (result MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MiniAppsMgr) Gets() (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MiniAppsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAppID app_id获取 app_id
func (obj *_MiniAppsMgr) WithAppID(appID string) Option {
	return optionFunc(func(o *options) { o.query["app_id"] = appID })
}

// WithAppSecret app_secret获取 app_secret
func (obj *_MiniAppsMgr) WithAppSecret(appSecret string) Option {
	return optionFunc(func(o *options) { o.query["app_secret"] = appSecret })
}

// WithUUID uuid获取 uuid
func (obj *_MiniAppsMgr) WithUUID(uuid string) Option {
	return optionFunc(func(o *options) { o.query["uuid"] = uuid })
}

// WithName name获取 应用名称
func (obj *_MiniAppsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDescrible describle获取 应用描述
func (obj *_MiniAppsMgr) WithDescrible(describle string) Option {
	return optionFunc(func(o *options) { o.query["describle"] = describle })
}

// WithCreateAt create_at获取
func (obj *_MiniAppsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_MiniAppsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_MiniAppsMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// GetByOption 功能选项模式获取
func (obj *_MiniAppsMgr) GetByOption(opts ...Option) (result MiniApps, err error) {
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
func (obj *_MiniAppsMgr) GetByOptions(opts ...Option) (results []*MiniApps, err error) {
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
func (obj *_MiniAppsMgr) GetFromID(id int) (result MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_MiniAppsMgr) GetBatchFromID(ids []int) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromAppID 通过app_id获取内容 app_id
func (obj *_MiniAppsMgr) GetFromAppID(appID string) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_id = ?", appID).Find(&results).Error

	return
}

// GetBatchFromAppID 批量唯一主键查找 app_id
func (obj *_MiniAppsMgr) GetBatchFromAppID(appIDs []string) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_id IN (?)", appIDs).Find(&results).Error

	return
}

// GetFromAppSecret 通过app_secret获取内容 app_secret
func (obj *_MiniAppsMgr) GetFromAppSecret(appSecret string) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_secret = ?", appSecret).Find(&results).Error

	return
}

// GetBatchFromAppSecret 批量唯一主键查找 app_secret
func (obj *_MiniAppsMgr) GetBatchFromAppSecret(appSecrets []string) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_secret IN (?)", appSecrets).Find(&results).Error

	return
}

// GetFromUUID 通过uuid获取内容 uuid
func (obj *_MiniAppsMgr) GetFromUUID(uuid string) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid = ?", uuid).Find(&results).Error

	return
}

// GetBatchFromUUID 批量唯一主键查找 uuid
func (obj *_MiniAppsMgr) GetBatchFromUUID(uuids []string) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("uuid IN (?)", uuids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 应用名称
func (obj *_MiniAppsMgr) GetFromName(name string) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 应用名称
func (obj *_MiniAppsMgr) GetBatchFromName(names []string) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromDescrible 通过describle获取内容 应用描述
func (obj *_MiniAppsMgr) GetFromDescrible(describle string) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("describle = ?", describle).Find(&results).Error

	return
}

// GetBatchFromDescrible 批量唯一主键查找 应用描述
func (obj *_MiniAppsMgr) GetBatchFromDescrible(describles []string) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("describle IN (?)", describles).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_MiniAppsMgr) GetFromCreateAt(createAt time.Time) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_MiniAppsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_MiniAppsMgr) GetFromUpdateAt(updateAt time.Time) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_MiniAppsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_MiniAppsMgr) GetFromIsDeleted(isDeleted time.Time) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_MiniAppsMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_MiniAppsMgr) FetchByPrimaryKey(id int) (result MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUUID primay or index 获取唯一内容
func (obj *_MiniAppsMgr) FetchUniqueIndexByUUID(appID string, uuid string) (result MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_id = ? AND uuid = ?", appID, uuid).Find(&result).Error

	return
}

// FetchIndexByName  获取多个内容
func (obj *_MiniAppsMgr) FetchIndexByName(name string) (results []*MiniApps, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
