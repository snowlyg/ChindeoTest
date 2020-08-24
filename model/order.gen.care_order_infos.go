package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CareOrderInfosMgr struct {
	*_BaseMgr
}

// CareOrderInfosMgr open func
func CareOrderInfosMgr(db *gorm.DB) *_CareOrderInfosMgr {
	if db == nil {
		panic(fmt.Errorf("CareOrderInfosMgr need init by db"))
	}
	return &_CareOrderInfosMgr{_BaseMgr: &_BaseMgr{DB: db.Table("care_order_infos"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CareOrderInfosMgr) GetTableName() string {
	return "care_order_infos"
}

// Get 获取
func (obj *_CareOrderInfosMgr) Get() (result CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CareOrderInfosMgr) Gets() (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CareOrderInfosMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_CareOrderInfosMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDesc desc获取 简介
func (obj *_CareOrderInfosMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithTimeType time_type获取 时间类型：时，天，月
func (obj *_CareOrderInfosMgr) WithTimeType(timeType string) Option {
	return optionFunc(func(o *options) { o.query["time_type"] = timeType })
}

// WithCareType care_type获取 类型
func (obj *_CareOrderInfosMgr) WithCareType(careType string) Option {
	return optionFunc(func(o *options) { o.query["care_type"] = careType })
}

// WithCareTags care_tags获取 标签：用 | 分隔
func (obj *_CareOrderInfosMgr) WithCareTags(careTags string) Option {
	return optionFunc(func(o *options) { o.query["care_tags"] = careTags })
}

// WithMinPrice min_price获取 最小价格
func (obj *_CareOrderInfosMgr) WithMinPrice(minPrice float64) Option {
	return optionFunc(func(o *options) { o.query["min_price"] = minPrice })
}

// WithMaxPrice max_price获取 最大价格
func (obj *_CareOrderInfosMgr) WithMaxPrice(maxPrice float64) Option {
	return optionFunc(func(o *options) { o.query["max_price"] = maxPrice })
}

// WithCover cover获取 封面
func (obj *_CareOrderInfosMgr) WithCover(cover string) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

// WithCareDetail care_detail获取 服务内容:采用富文本
func (obj *_CareOrderInfosMgr) WithCareDetail(careDetail string) Option {
	return optionFunc(func(o *options) { o.query["care_detail"] = careDetail })
}

// WithApplicationName application_name获取 医院名称
func (obj *_CareOrderInfosMgr) WithApplicationName(applicationName string) Option {
	return optionFunc(func(o *options) { o.query["application_name"] = applicationName })
}

// WithCreateAt create_at获取
func (obj *_CareOrderInfosMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CareOrderInfosMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_CareOrderInfosMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithCareOrderID care_order_id获取 care_order_id
func (obj *_CareOrderInfosMgr) WithCareOrderID(careOrderID int) Option {
	return optionFunc(func(o *options) { o.query["care_order_id"] = careOrderID })
}

// GetByOption 功能选项模式获取
func (obj *_CareOrderInfosMgr) GetByOption(opts ...Option) (result CareOrderInfos, err error) {
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
func (obj *_CareOrderInfosMgr) GetByOptions(opts ...Option) (results []*CareOrderInfos, err error) {
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
func (obj *_CareOrderInfosMgr) GetFromID(id int) (result CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CareOrderInfosMgr) GetBatchFromID(ids []int) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_CareOrderInfosMgr) GetFromName(name string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_CareOrderInfosMgr) GetBatchFromName(names []string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 简介
func (obj *_CareOrderInfosMgr) GetFromDesc(desc string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 简介
func (obj *_CareOrderInfosMgr) GetBatchFromDesc(descs []string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

// GetFromTimeType 通过time_type获取内容 时间类型：时，天，月
func (obj *_CareOrderInfosMgr) GetFromTimeType(timeType string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type = ?", timeType).Find(&results).Error

	return
}

// GetBatchFromTimeType 批量唯一主键查找 时间类型：时，天，月
func (obj *_CareOrderInfosMgr) GetBatchFromTimeType(timeTypes []string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type IN (?)", timeTypes).Find(&results).Error

	return
}

// GetFromCareType 通过care_type获取内容 类型
func (obj *_CareOrderInfosMgr) GetFromCareType(careType string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_type = ?", careType).Find(&results).Error

	return
}

// GetBatchFromCareType 批量唯一主键查找 类型
func (obj *_CareOrderInfosMgr) GetBatchFromCareType(careTypes []string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_type IN (?)", careTypes).Find(&results).Error

	return
}

// GetFromCareTags 通过care_tags获取内容 标签：用 | 分隔
func (obj *_CareOrderInfosMgr) GetFromCareTags(careTags string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_tags = ?", careTags).Find(&results).Error

	return
}

// GetBatchFromCareTags 批量唯一主键查找 标签：用 | 分隔
func (obj *_CareOrderInfosMgr) GetBatchFromCareTags(careTagss []string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_tags IN (?)", careTagss).Find(&results).Error

	return
}

// GetFromMinPrice 通过min_price获取内容 最小价格
func (obj *_CareOrderInfosMgr) GetFromMinPrice(minPrice float64) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("min_price = ?", minPrice).Find(&results).Error

	return
}

// GetBatchFromMinPrice 批量唯一主键查找 最小价格
func (obj *_CareOrderInfosMgr) GetBatchFromMinPrice(minPrices []float64) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("min_price IN (?)", minPrices).Find(&results).Error

	return
}

// GetFromMaxPrice 通过max_price获取内容 最大价格
func (obj *_CareOrderInfosMgr) GetFromMaxPrice(maxPrice float64) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("max_price = ?", maxPrice).Find(&results).Error

	return
}

// GetBatchFromMaxPrice 批量唯一主键查找 最大价格
func (obj *_CareOrderInfosMgr) GetBatchFromMaxPrice(maxPrices []float64) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("max_price IN (?)", maxPrices).Find(&results).Error

	return
}

// GetFromCover 通过cover获取内容 封面
func (obj *_CareOrderInfosMgr) GetFromCover(cover string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover = ?", cover).Find(&results).Error

	return
}

// GetBatchFromCover 批量唯一主键查找 封面
func (obj *_CareOrderInfosMgr) GetBatchFromCover(covers []string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover IN (?)", covers).Find(&results).Error

	return
}

// GetFromCareDetail 通过care_detail获取内容 服务内容:采用富文本
func (obj *_CareOrderInfosMgr) GetFromCareDetail(careDetail string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_detail = ?", careDetail).Find(&results).Error

	return
}

// GetBatchFromCareDetail 批量唯一主键查找 服务内容:采用富文本
func (obj *_CareOrderInfosMgr) GetBatchFromCareDetail(careDetails []string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_detail IN (?)", careDetails).Find(&results).Error

	return
}

// GetFromApplicationName 通过application_name获取内容 医院名称
func (obj *_CareOrderInfosMgr) GetFromApplicationName(applicationName string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_name = ?", applicationName).Find(&results).Error

	return
}

// GetBatchFromApplicationName 批量唯一主键查找 医院名称
func (obj *_CareOrderInfosMgr) GetBatchFromApplicationName(applicationNames []string) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_name IN (?)", applicationNames).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CareOrderInfosMgr) GetFromCreateAt(createAt time.Time) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CareOrderInfosMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CareOrderInfosMgr) GetFromUpdateAt(updateAt time.Time) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CareOrderInfosMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CareOrderInfosMgr) GetFromIsDeleted(isDeleted time.Time) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CareOrderInfosMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCareOrderID 通过care_order_id获取内容 care_order_id
func (obj *_CareOrderInfosMgr) GetFromCareOrderID(careOrderID int) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id = ?", careOrderID).Find(&results).Error

	return
}

// GetBatchFromCareOrderID 批量唯一主键查找 care_order_id
func (obj *_CareOrderInfosMgr) GetBatchFromCareOrderID(careOrderIDs []int) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id IN (?)", careOrderIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CareOrderInfosMgr) FetchByPrimaryKey(id int) (result CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByCareOrderID  获取多个内容
func (obj *_CareOrderInfosMgr) FetchIndexByCareOrderID(careOrderID int) (results []*CareOrderInfos, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id = ?", careOrderID).Find(&results).Error

	return
}
