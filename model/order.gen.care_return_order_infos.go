package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CareReturnOrderInfosMgr struct {
	*_BaseMgr
}

// CareReturnOrderInfosMgr open func
func CareReturnOrderInfosMgr(db *gorm.DB) *_CareReturnOrderInfosMgr {
	if db == nil {
		panic(fmt.Errorf("CareReturnOrderInfosMgr need init by db"))
	}
	return &_CareReturnOrderInfosMgr{_BaseMgr: &_BaseMgr{DB: db.Table("care_return_order_infos"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CareReturnOrderInfosMgr) GetTableName() string {
	return "care_return_order_infos"
}

// Get 获取
func (obj *_CareReturnOrderInfosMgr) Get() (result CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CareReturnOrderInfosMgr) Gets() (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CareReturnOrderInfosMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_CareReturnOrderInfosMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDesc desc获取 简介
func (obj *_CareReturnOrderInfosMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithTimeType time_type获取 时间类型：时，天，月
func (obj *_CareReturnOrderInfosMgr) WithTimeType(timeType string) Option {
	return optionFunc(func(o *options) { o.query["time_type"] = timeType })
}

// WithCareType care_type获取 类型
func (obj *_CareReturnOrderInfosMgr) WithCareType(careType string) Option {
	return optionFunc(func(o *options) { o.query["care_type"] = careType })
}

// WithCareTags care_tags获取 标签：用 | 分隔
func (obj *_CareReturnOrderInfosMgr) WithCareTags(careTags string) Option {
	return optionFunc(func(o *options) { o.query["care_tags"] = careTags })
}

// WithMinPrice min_price获取 最小价格
func (obj *_CareReturnOrderInfosMgr) WithMinPrice(minPrice float64) Option {
	return optionFunc(func(o *options) { o.query["min_price"] = minPrice })
}

// WithMaxPrice max_price获取 最大价格
func (obj *_CareReturnOrderInfosMgr) WithMaxPrice(maxPrice float64) Option {
	return optionFunc(func(o *options) { o.query["max_price"] = maxPrice })
}

// WithCover cover获取 封面
func (obj *_CareReturnOrderInfosMgr) WithCover(cover string) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

// WithCareDetail care_detail获取 服务内容:采用富文本
func (obj *_CareReturnOrderInfosMgr) WithCareDetail(careDetail string) Option {
	return optionFunc(func(o *options) { o.query["care_detail"] = careDetail })
}

// WithApplicationName application_name获取 医院名称
func (obj *_CareReturnOrderInfosMgr) WithApplicationName(applicationName string) Option {
	return optionFunc(func(o *options) { o.query["application_name"] = applicationName })
}

// WithCreateAt create_at获取
func (obj *_CareReturnOrderInfosMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CareReturnOrderInfosMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_CareReturnOrderInfosMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithCareReturnOrderID care_return_order_id获取 care_return_order_id
func (obj *_CareReturnOrderInfosMgr) WithCareReturnOrderID(careReturnOrderID int) Option {
	return optionFunc(func(o *options) { o.query["care_return_order_id"] = careReturnOrderID })
}

// GetByOption 功能选项模式获取
func (obj *_CareReturnOrderInfosMgr) GetByOption(opts ...Option) (result CareReturnOrderInfos, err error) {
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
func (obj *_CareReturnOrderInfosMgr) GetByOptions(opts ...Option) (results []*CareReturnOrderInfos, err error) {
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
func (obj *_CareReturnOrderInfosMgr) GetFromID(id int) (result CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CareReturnOrderInfosMgr) GetBatchFromID(ids []int) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_CareReturnOrderInfosMgr) GetFromName(name string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_CareReturnOrderInfosMgr) GetBatchFromName(names []string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 简介
func (obj *_CareReturnOrderInfosMgr) GetFromDesc(desc string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 简介
func (obj *_CareReturnOrderInfosMgr) GetBatchFromDesc(descs []string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

// GetFromTimeType 通过time_type获取内容 时间类型：时，天，月
func (obj *_CareReturnOrderInfosMgr) GetFromTimeType(timeType string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type = ?", timeType).Find(&results).Error

	return
}

// GetBatchFromTimeType 批量唯一主键查找 时间类型：时，天，月
func (obj *_CareReturnOrderInfosMgr) GetBatchFromTimeType(timeTypes []string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type IN (?)", timeTypes).Find(&results).Error

	return
}

// GetFromCareType 通过care_type获取内容 类型
func (obj *_CareReturnOrderInfosMgr) GetFromCareType(careType string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_type = ?", careType).Find(&results).Error

	return
}

// GetBatchFromCareType 批量唯一主键查找 类型
func (obj *_CareReturnOrderInfosMgr) GetBatchFromCareType(careTypes []string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_type IN (?)", careTypes).Find(&results).Error

	return
}

// GetFromCareTags 通过care_tags获取内容 标签：用 | 分隔
func (obj *_CareReturnOrderInfosMgr) GetFromCareTags(careTags string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_tags = ?", careTags).Find(&results).Error

	return
}

// GetBatchFromCareTags 批量唯一主键查找 标签：用 | 分隔
func (obj *_CareReturnOrderInfosMgr) GetBatchFromCareTags(careTagss []string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_tags IN (?)", careTagss).Find(&results).Error

	return
}

// GetFromMinPrice 通过min_price获取内容 最小价格
func (obj *_CareReturnOrderInfosMgr) GetFromMinPrice(minPrice float64) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("min_price = ?", minPrice).Find(&results).Error

	return
}

// GetBatchFromMinPrice 批量唯一主键查找 最小价格
func (obj *_CareReturnOrderInfosMgr) GetBatchFromMinPrice(minPrices []float64) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("min_price IN (?)", minPrices).Find(&results).Error

	return
}

// GetFromMaxPrice 通过max_price获取内容 最大价格
func (obj *_CareReturnOrderInfosMgr) GetFromMaxPrice(maxPrice float64) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("max_price = ?", maxPrice).Find(&results).Error

	return
}

// GetBatchFromMaxPrice 批量唯一主键查找 最大价格
func (obj *_CareReturnOrderInfosMgr) GetBatchFromMaxPrice(maxPrices []float64) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("max_price IN (?)", maxPrices).Find(&results).Error

	return
}

// GetFromCover 通过cover获取内容 封面
func (obj *_CareReturnOrderInfosMgr) GetFromCover(cover string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover = ?", cover).Find(&results).Error

	return
}

// GetBatchFromCover 批量唯一主键查找 封面
func (obj *_CareReturnOrderInfosMgr) GetBatchFromCover(covers []string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover IN (?)", covers).Find(&results).Error

	return
}

// GetFromCareDetail 通过care_detail获取内容 服务内容:采用富文本
func (obj *_CareReturnOrderInfosMgr) GetFromCareDetail(careDetail string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_detail = ?", careDetail).Find(&results).Error

	return
}

// GetBatchFromCareDetail 批量唯一主键查找 服务内容:采用富文本
func (obj *_CareReturnOrderInfosMgr) GetBatchFromCareDetail(careDetails []string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_detail IN (?)", careDetails).Find(&results).Error

	return
}

// GetFromApplicationName 通过application_name获取内容 医院名称
func (obj *_CareReturnOrderInfosMgr) GetFromApplicationName(applicationName string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_name = ?", applicationName).Find(&results).Error

	return
}

// GetBatchFromApplicationName 批量唯一主键查找 医院名称
func (obj *_CareReturnOrderInfosMgr) GetBatchFromApplicationName(applicationNames []string) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_name IN (?)", applicationNames).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CareReturnOrderInfosMgr) GetFromCreateAt(createAt time.Time) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CareReturnOrderInfosMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CareReturnOrderInfosMgr) GetFromUpdateAt(updateAt time.Time) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CareReturnOrderInfosMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CareReturnOrderInfosMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CareReturnOrderInfosMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCareReturnOrderID 通过care_return_order_id获取内容 care_return_order_id
func (obj *_CareReturnOrderInfosMgr) GetFromCareReturnOrderID(careReturnOrderID int) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_return_order_id = ?", careReturnOrderID).Find(&results).Error

	return
}

// GetBatchFromCareReturnOrderID 批量唯一主键查找 care_return_order_id
func (obj *_CareReturnOrderInfosMgr) GetBatchFromCareReturnOrderID(careReturnOrderIDs []int) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_return_order_id IN (?)", careReturnOrderIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CareReturnOrderInfosMgr) FetchByPrimaryKey(id int) (result CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByCareReturnOrderID  获取多个内容
func (obj *_CareReturnOrderInfosMgr) FetchIndexByCareReturnOrderID(careReturnOrderID int) (results []*CareReturnOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_return_order_id = ?", careReturnOrderID).Find(&results).Error

	return
}
