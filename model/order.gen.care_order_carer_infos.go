package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CareOrderCarerInfosMgr struct {
	*_BaseMgr
}

// CareOrderCarerInfosMgr open func
func CareOrderCarerInfosMgr(db *gorm.DB) *_CareOrderCarerInfosMgr {
	if db == nil {
		panic(fmt.Errorf("CareOrderCarerInfosMgr need init by db"))
	}
	return &_CareOrderCarerInfosMgr{_BaseMgr: &_BaseMgr{DB: db.Table("care_order_carer_infos"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CareOrderCarerInfosMgr) GetTableName() string {
	return "care_order_carer_infos"
}

// Get 获取
func (obj *_CareOrderCarerInfosMgr) Get() (result CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CareOrderCarerInfosMgr) Gets() (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CareOrderCarerInfosMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 姓名
func (obj *_CareOrderCarerInfosMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAge age获取 年龄
func (obj *_CareOrderCarerInfosMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// WithWorkExp work_exp获取 工作年限
func (obj *_CareOrderCarerInfosMgr) WithWorkExp(workExp int) Option {
	return optionFunc(func(o *options) { o.query["work_exp"] = workExp })
}

// WithPrice price获取 价格
func (obj *_CareOrderCarerInfosMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithSex sex获取 性别:1:男,0:女
func (obj *_CareOrderCarerInfosMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithPhone phone获取 手机
func (obj *_CareOrderCarerInfosMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithDesc desc获取 简介
func (obj *_CareOrderCarerInfosMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithAvatar avatar获取 头像
func (obj *_CareOrderCarerInfosMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithCarerTags carer_tags获取 标签
func (obj *_CareOrderCarerInfosMgr) WithCarerTags(carerTags string) Option {
	return optionFunc(func(o *options) { o.query["carer_tags"] = carerTags })
}

// WithApplicationName application_name获取 医院名称
func (obj *_CareOrderCarerInfosMgr) WithApplicationName(applicationName string) Option {
	return optionFunc(func(o *options) { o.query["application_name"] = applicationName })
}

// WithCreateAt create_at获取
func (obj *_CareOrderCarerInfosMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CareOrderCarerInfosMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_CareOrderCarerInfosMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithCareOrderID care_order_id获取 care_order_id
func (obj *_CareOrderCarerInfosMgr) WithCareOrderID(careOrderID int) Option {
	return optionFunc(func(o *options) { o.query["care_order_id"] = careOrderID })
}

// WithTimeType time_type获取 时间类型：时，天，月
func (obj *_CareOrderCarerInfosMgr) WithTimeType(timeType string) Option {
	return optionFunc(func(o *options) { o.query["time_type"] = timeType })
}

// GetByOption 功能选项模式获取
func (obj *_CareOrderCarerInfosMgr) GetByOption(opts ...Option) (result CareOrderCarerInfos, err error) {
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
func (obj *_CareOrderCarerInfosMgr) GetByOptions(opts ...Option) (results []*CareOrderCarerInfos, err error) {
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
func (obj *_CareOrderCarerInfosMgr) GetFromID(id int) (result CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CareOrderCarerInfosMgr) GetBatchFromID(ids []int) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_CareOrderCarerInfosMgr) GetFromName(name string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 姓名
func (obj *_CareOrderCarerInfosMgr) GetBatchFromName(names []string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容 年龄
func (obj *_CareOrderCarerInfosMgr) GetFromAge(age int) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量唯一主键查找 年龄
func (obj *_CareOrderCarerInfosMgr) GetBatchFromAge(ages []int) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age IN (?)", ages).Find(&results).Error

	return
}

// GetFromWorkExp 通过work_exp获取内容 工作年限
func (obj *_CareOrderCarerInfosMgr) GetFromWorkExp(workExp int) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("work_exp = ?", workExp).Find(&results).Error

	return
}

// GetBatchFromWorkExp 批量唯一主键查找 工作年限
func (obj *_CareOrderCarerInfosMgr) GetBatchFromWorkExp(workExps []int) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("work_exp IN (?)", workExps).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_CareOrderCarerInfosMgr) GetFromPrice(price float64) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量唯一主键查找 价格
func (obj *_CareOrderCarerInfosMgr) GetBatchFromPrice(prices []float64) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 性别:1:男,0:女
func (obj *_CareOrderCarerInfosMgr) GetFromSex(sex int) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找 性别:1:男,0:女
func (obj *_CareOrderCarerInfosMgr) GetBatchFromSex(sexs []int) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", sexs).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机
func (obj *_CareOrderCarerInfosMgr) GetFromPhone(phone string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 手机
func (obj *_CareOrderCarerInfosMgr) GetBatchFromPhone(phones []string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 简介
func (obj *_CareOrderCarerInfosMgr) GetFromDesc(desc string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 简介
func (obj *_CareOrderCarerInfosMgr) GetBatchFromDesc(descs []string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像
func (obj *_CareOrderCarerInfosMgr) GetFromAvatar(avatar string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("avatar = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量唯一主键查找 头像
func (obj *_CareOrderCarerInfosMgr) GetBatchFromAvatar(avatars []string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("avatar IN (?)", avatars).Find(&results).Error

	return
}

// GetFromCarerTags 通过carer_tags获取内容 标签
func (obj *_CareOrderCarerInfosMgr) GetFromCarerTags(carerTags string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_tags = ?", carerTags).Find(&results).Error

	return
}

// GetBatchFromCarerTags 批量唯一主键查找 标签
func (obj *_CareOrderCarerInfosMgr) GetBatchFromCarerTags(carerTagss []string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_tags IN (?)", carerTagss).Find(&results).Error

	return
}

// GetFromApplicationName 通过application_name获取内容 医院名称
func (obj *_CareOrderCarerInfosMgr) GetFromApplicationName(applicationName string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_name = ?", applicationName).Find(&results).Error

	return
}

// GetBatchFromApplicationName 批量唯一主键查找 医院名称
func (obj *_CareOrderCarerInfosMgr) GetBatchFromApplicationName(applicationNames []string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_name IN (?)", applicationNames).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CareOrderCarerInfosMgr) GetFromCreateAt(createAt time.Time) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CareOrderCarerInfosMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CareOrderCarerInfosMgr) GetFromUpdateAt(updateAt time.Time) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CareOrderCarerInfosMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CareOrderCarerInfosMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CareOrderCarerInfosMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCareOrderID 通过care_order_id获取内容 care_order_id
func (obj *_CareOrderCarerInfosMgr) GetFromCareOrderID(careOrderID int) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id = ?", careOrderID).Find(&results).Error

	return
}

// GetBatchFromCareOrderID 批量唯一主键查找 care_order_id
func (obj *_CareOrderCarerInfosMgr) GetBatchFromCareOrderID(careOrderIDs []int) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id IN (?)", careOrderIDs).Find(&results).Error

	return
}

// GetFromTimeType 通过time_type获取内容 时间类型：时，天，月
func (obj *_CareOrderCarerInfosMgr) GetFromTimeType(timeType string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type = ?", timeType).Find(&results).Error

	return
}

// GetBatchFromTimeType 批量唯一主键查找 时间类型：时，天，月
func (obj *_CareOrderCarerInfosMgr) GetBatchFromTimeType(timeTypes []string) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type IN (?)", timeTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CareOrderCarerInfosMgr) FetchByPrimaryKey(id int) (result CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByCareOrderID  获取多个内容
func (obj *_CareOrderCarerInfosMgr) FetchIndexByCareOrderID(careOrderID int) (results []*CareOrderCarerInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id = ?", careOrderID).Find(&results).Error

	return
}
