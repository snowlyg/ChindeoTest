package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIApplicationsMgr struct {
	*_BaseMgr
}

// APIApplicationsMgr open func
func APIApplicationsMgr(db *gorm.DB) *_APIApplicationsMgr {
	if db == nil {
		panic(fmt.Errorf("APIApplicationsMgr need init by db"))
	}
	return &_APIApplicationsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_applications"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIApplicationsMgr) GetTableName() string {
	return "api_applications"
}

// Get 获取
func (obj *_APIApplicationsMgr) Get() (result APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIApplicationsMgr) Gets() (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIApplicationsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_APIApplicationsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAppID app_id获取 app_id
func (obj *_APIApplicationsMgr) WithAppID(appID string) Option {
	return optionFunc(func(o *options) { o.query["app_id"] = appID })
}

// WithAppSecret app_secret获取 app_secret
func (obj *_APIApplicationsMgr) WithAppSecret(appSecret string) Option {
	return optionFunc(func(o *options) { o.query["app_secret"] = appSecret })
}

// WithTel tel获取
func (obj *_APIApplicationsMgr) WithTel(tel string) Option {
	return optionFunc(func(o *options) { o.query["tel"] = tel })
}

// WithAddr addr获取
func (obj *_APIApplicationsMgr) WithAddr(addr string) Option {
	return optionFunc(func(o *options) { o.query["addr"] = addr })
}

// WithDescrible describle获取 应用描述
func (obj *_APIApplicationsMgr) WithDescrible(describle string) Option {
	return optionFunc(func(o *options) { o.query["describle"] = describle })
}

// WithAppType app_type获取 app_type
func (obj *_APIApplicationsMgr) WithAppType(appType string) Option {
	return optionFunc(func(o *options) { o.query["app_type"] = appType })
}

// WithCreateAt create_at获取
func (obj *_APIApplicationsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIApplicationsMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithUserID user_id获取
func (obj *_APIApplicationsMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCnareaID cnarea_id获取
func (obj *_APIApplicationsMgr) WithCnareaID(cnareaID int) Option {
	return optionFunc(func(o *options) { o.query["cnarea_id"] = cnareaID })
}

// WithUpdateAt update_at获取
func (obj *_APIApplicationsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithBusinessHours business_hours获取 营业时间
func (obj *_APIApplicationsMgr) WithBusinessHours(businessHours string) Option {
	return optionFunc(func(o *options) { o.query["business_hours"] = businessHours })
}

// GetByOption 功能选项模式获取
func (obj *_APIApplicationsMgr) GetByOption(opts ...Option) (result APIApplications, err error) {
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
func (obj *_APIApplicationsMgr) GetByOptions(opts ...Option) (results []*APIApplications, err error) {
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
func (obj *_APIApplicationsMgr) GetFromID(id int) (result APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIApplicationsMgr) GetBatchFromID(ids []int) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_APIApplicationsMgr) GetFromName(name string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_APIApplicationsMgr) GetBatchFromName(names []string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromAppID 通过app_id获取内容 app_id
func (obj *_APIApplicationsMgr) GetFromAppID(appID string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_id = ?", appID).Find(&results).Error

	return
}

// GetBatchFromAppID 批量唯一主键查找 app_id
func (obj *_APIApplicationsMgr) GetBatchFromAppID(appIDs []string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_id IN (?)", appIDs).Find(&results).Error

	return
}

// GetFromAppSecret 通过app_secret获取内容 app_secret
func (obj *_APIApplicationsMgr) GetFromAppSecret(appSecret string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_secret = ?", appSecret).Find(&results).Error

	return
}

// GetBatchFromAppSecret 批量唯一主键查找 app_secret
func (obj *_APIApplicationsMgr) GetBatchFromAppSecret(appSecrets []string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_secret IN (?)", appSecrets).Find(&results).Error

	return
}

// GetFromTel 通过tel获取内容
func (obj *_APIApplicationsMgr) GetFromTel(tel string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("tel = ?", tel).Find(&results).Error

	return
}

// GetBatchFromTel 批量唯一主键查找
func (obj *_APIApplicationsMgr) GetBatchFromTel(tels []string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("tel IN (?)", tels).Find(&results).Error

	return
}

// GetFromAddr 通过addr获取内容
func (obj *_APIApplicationsMgr) GetFromAddr(addr string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addr = ?", addr).Find(&results).Error

	return
}

// GetBatchFromAddr 批量唯一主键查找
func (obj *_APIApplicationsMgr) GetBatchFromAddr(addrs []string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addr IN (?)", addrs).Find(&results).Error

	return
}

// GetFromDescrible 通过describle获取内容 应用描述
func (obj *_APIApplicationsMgr) GetFromDescrible(describle string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("describle = ?", describle).Find(&results).Error

	return
}

// GetBatchFromDescrible 批量唯一主键查找 应用描述
func (obj *_APIApplicationsMgr) GetBatchFromDescrible(describles []string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("describle IN (?)", describles).Find(&results).Error

	return
}

// GetFromAppType 通过app_type获取内容 app_type
func (obj *_APIApplicationsMgr) GetFromAppType(appType string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_type = ?", appType).Find(&results).Error

	return
}

// GetBatchFromAppType 批量唯一主键查找 app_type
func (obj *_APIApplicationsMgr) GetBatchFromAppType(appTypes []string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("app_type IN (?)", appTypes).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIApplicationsMgr) GetFromCreateAt(createAt time.Time) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIApplicationsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIApplicationsMgr) GetFromIsDeleted(isDeleted time.Time) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIApplicationsMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_APIApplicationsMgr) GetFromUserID(userID int) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_APIApplicationsMgr) GetBatchFromUserID(userIDs []int) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCnareaID 通过cnarea_id获取内容
func (obj *_APIApplicationsMgr) GetFromCnareaID(cnareaID int) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cnarea_id = ?", cnareaID).Find(&results).Error

	return
}

// GetBatchFromCnareaID 批量唯一主键查找
func (obj *_APIApplicationsMgr) GetBatchFromCnareaID(cnareaIDs []int) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cnarea_id IN (?)", cnareaIDs).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIApplicationsMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIApplicationsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromBusinessHours 通过business_hours获取内容 营业时间
func (obj *_APIApplicationsMgr) GetFromBusinessHours(businessHours string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("business_hours = ?", businessHours).Find(&results).Error

	return
}

// GetBatchFromBusinessHours 批量唯一主键查找 营业时间
func (obj *_APIApplicationsMgr) GetBatchFromBusinessHours(businessHourss []string) (results []*APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("business_hours IN (?)", businessHourss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIApplicationsMgr) FetchByPrimaryKey(id int) (result APIApplications, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
