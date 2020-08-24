package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CaresMgr struct {
	*_BaseMgr
}

// CaresMgr open func
func CaresMgr(db *gorm.DB) *_CaresMgr {
	if db == nil {
		panic(fmt.Errorf("CaresMgr need init by db"))
	}
	return &_CaresMgr{_BaseMgr: &_BaseMgr{DB: db.Table("cares"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CaresMgr) GetTableName() string {
	return "cares"
}

// Get 获取
func (obj *_CaresMgr) Get() (result Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CaresMgr) Gets() (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CaresMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_CaresMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDesc desc获取 简介
func (obj *_CaresMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithTimeType time_type获取 时间类型：时，天，月
func (obj *_CaresMgr) WithTimeType(timeType string) Option {
	return optionFunc(func(o *options) { o.query["time_type"] = timeType })
}

// WithStatus status获取 状态：上下架
func (obj *_CaresMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithAmount amount获取 销量
func (obj *_CaresMgr) WithAmount(amount int) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithMinPrice min_price获取 最小价格
func (obj *_CaresMgr) WithMinPrice(minPrice float64) Option {
	return optionFunc(func(o *options) { o.query["min_price"] = minPrice })
}

// WithMaxPrice max_price获取 最大价格
func (obj *_CaresMgr) WithMaxPrice(maxPrice float64) Option {
	return optionFunc(func(o *options) { o.query["max_price"] = maxPrice })
}

// WithCover cover获取 封面
func (obj *_CaresMgr) WithCover(cover string) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

// WithSort sort获取 排序
func (obj *_CaresMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithCareDetail care_detail获取 服务内容:采用富文本
func (obj *_CaresMgr) WithCareDetail(careDetail string) Option {
	return optionFunc(func(o *options) { o.query["care_detail"] = careDetail })
}

// WithCreateAt create_at获取
func (obj *_CaresMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CaresMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_CaresMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithCareTypeID care_type_id获取 care_type_id
func (obj *_CaresMgr) WithCareTypeID(careTypeID int) Option {
	return optionFunc(func(o *options) { o.query["care_type_id"] = careTypeID })
}

// WithApplicationID application_id获取 application_id
func (obj *_CaresMgr) WithApplicationID(applicationID int) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// GetByOption 功能选项模式获取
func (obj *_CaresMgr) GetByOption(opts ...Option) (result Cares, err error) {
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
func (obj *_CaresMgr) GetByOptions(opts ...Option) (results []*Cares, err error) {
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
func (obj *_CaresMgr) GetFromID(id int) (result Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CaresMgr) GetBatchFromID(ids []int) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_CaresMgr) GetFromName(name string) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_CaresMgr) GetBatchFromName(names []string) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 简介
func (obj *_CaresMgr) GetFromDesc(desc string) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 简介
func (obj *_CaresMgr) GetBatchFromDesc(descs []string) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

// GetFromTimeType 通过time_type获取内容 时间类型：时，天，月
func (obj *_CaresMgr) GetFromTimeType(timeType string) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type = ?", timeType).Find(&results).Error

	return
}

// GetBatchFromTimeType 批量唯一主键查找 时间类型：时，天，月
func (obj *_CaresMgr) GetBatchFromTimeType(timeTypes []string) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type IN (?)", timeTypes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态：上下架
func (obj *_CaresMgr) GetFromStatus(status bool) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态：上下架
func (obj *_CaresMgr) GetBatchFromStatus(statuss []bool) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 销量
func (obj *_CaresMgr) GetFromAmount(amount int) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量唯一主键查找 销量
func (obj *_CaresMgr) GetBatchFromAmount(amounts []int) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount IN (?)", amounts).Find(&results).Error

	return
}

// GetFromMinPrice 通过min_price获取内容 最小价格
func (obj *_CaresMgr) GetFromMinPrice(minPrice float64) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("min_price = ?", minPrice).Find(&results).Error

	return
}

// GetBatchFromMinPrice 批量唯一主键查找 最小价格
func (obj *_CaresMgr) GetBatchFromMinPrice(minPrices []float64) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("min_price IN (?)", minPrices).Find(&results).Error

	return
}

// GetFromMaxPrice 通过max_price获取内容 最大价格
func (obj *_CaresMgr) GetFromMaxPrice(maxPrice float64) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("max_price = ?", maxPrice).Find(&results).Error

	return
}

// GetBatchFromMaxPrice 批量唯一主键查找 最大价格
func (obj *_CaresMgr) GetBatchFromMaxPrice(maxPrices []float64) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("max_price IN (?)", maxPrices).Find(&results).Error

	return
}

// GetFromCover 通过cover获取内容 封面
func (obj *_CaresMgr) GetFromCover(cover string) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover = ?", cover).Find(&results).Error

	return
}

// GetBatchFromCover 批量唯一主键查找 封面
func (obj *_CaresMgr) GetBatchFromCover(covers []string) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover IN (?)", covers).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_CaresMgr) GetFromSort(sort int) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序
func (obj *_CaresMgr) GetBatchFromSort(sorts []int) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromCareDetail 通过care_detail获取内容 服务内容:采用富文本
func (obj *_CaresMgr) GetFromCareDetail(careDetail string) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_detail = ?", careDetail).Find(&results).Error

	return
}

// GetBatchFromCareDetail 批量唯一主键查找 服务内容:采用富文本
func (obj *_CaresMgr) GetBatchFromCareDetail(careDetails []string) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_detail IN (?)", careDetails).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CaresMgr) GetFromCreateAt(createAt time.Time) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CaresMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CaresMgr) GetFromUpdateAt(updateAt time.Time) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CaresMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CaresMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CaresMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCareTypeID 通过care_type_id获取内容 care_type_id
func (obj *_CaresMgr) GetFromCareTypeID(careTypeID int) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_type_id = ?", careTypeID).Find(&results).Error

	return
}

// GetBatchFromCareTypeID 批量唯一主键查找 care_type_id
func (obj *_CaresMgr) GetBatchFromCareTypeID(careTypeIDs []int) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_type_id IN (?)", careTypeIDs).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容 application_id
func (obj *_CaresMgr) GetFromApplicationID(applicationID int) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量唯一主键查找 application_id
func (obj *_CaresMgr) GetBatchFromApplicationID(applicationIDs []int) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id IN (?)", applicationIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CaresMgr) FetchByPrimaryKey(id int) (result Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByCareTypeID  获取多个内容
func (obj *_CaresMgr) FetchIndexByCareTypeID(careTypeID int, applicationID int) (results []*Cares, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_type_id = ? AND application_id = ?", careTypeID, applicationID).Find(&results).Error

	return
}
