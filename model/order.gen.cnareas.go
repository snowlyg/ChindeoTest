package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _CnareasMgr struct {
	*_BaseMgr
}

// CnareasMgr open func
func CnareasMgr(db *gorm.DB) *_CnareasMgr {
	if db == nil {
		panic(fmt.Errorf("CnareasMgr need init by db"))
	}
	return &_CnareasMgr{_BaseMgr: &_BaseMgr{DB: db.Table("cnareas"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CnareasMgr) GetTableName() string {
	return "cnareas"
}

// Get 获取
func (obj *_CnareasMgr) Get() (result Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CnareasMgr) Gets() (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CnareasMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithLevel level获取 层级
func (obj *_CnareasMgr) WithLevel(level uint8) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithParentCode parent_code获取 父级行政代码
func (obj *_CnareasMgr) WithParentCode(parentCode uint64) Option {
	return optionFunc(func(o *options) { o.query["parent_code"] = parentCode })
}

// WithAreaCode area_code获取 行政代码
func (obj *_CnareasMgr) WithAreaCode(areaCode uint64) Option {
	return optionFunc(func(o *options) { o.query["area_code"] = areaCode })
}

// WithZipCode zip_code获取 邮政编码
func (obj *_CnareasMgr) WithZipCode(zipCode string) Option {
	return optionFunc(func(o *options) { o.query["zip_code"] = zipCode })
}

// WithCityCode city_code获取 区号
func (obj *_CnareasMgr) WithCityCode(cityCode string) Option {
	return optionFunc(func(o *options) { o.query["city_code"] = cityCode })
}

// WithName name获取 名称
func (obj *_CnareasMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithShortName short_name获取 简称
func (obj *_CnareasMgr) WithShortName(shortName string) Option {
	return optionFunc(func(o *options) { o.query["short_name"] = shortName })
}

// WithMergerName merger_name获取 组合名
func (obj *_CnareasMgr) WithMergerName(mergerName string) Option {
	return optionFunc(func(o *options) { o.query["merger_name"] = mergerName })
}

// WithPinyin pinyin获取 拼音
func (obj *_CnareasMgr) WithPinyin(pinyin string) Option {
	return optionFunc(func(o *options) { o.query["pinyin"] = pinyin })
}

// WithLng lng获取 经度
func (obj *_CnareasMgr) WithLng(lng float64) Option {
	return optionFunc(func(o *options) { o.query["lng"] = lng })
}

// WithLat lat获取 纬度
func (obj *_CnareasMgr) WithLat(lat float64) Option {
	return optionFunc(func(o *options) { o.query["lat"] = lat })
}

// GetByOption 功能选项模式获取
func (obj *_CnareasMgr) GetByOption(opts ...Option) (result Cnareas, err error) {
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
func (obj *_CnareasMgr) GetByOptions(opts ...Option) (results []*Cnareas, err error) {
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
func (obj *_CnareasMgr) GetFromID(id string) (result Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CnareasMgr) GetBatchFromID(ids []string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容 层级
func (obj *_CnareasMgr) GetFromLevel(level uint8) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("level = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量唯一主键查找 层级
func (obj *_CnareasMgr) GetBatchFromLevel(levels []uint8) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("level IN (?)", levels).Find(&results).Error

	return
}

// GetFromParentCode 通过parent_code获取内容 父级行政代码
func (obj *_CnareasMgr) GetFromParentCode(parentCode uint64) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("parent_code = ?", parentCode).Find(&results).Error

	return
}

// GetBatchFromParentCode 批量唯一主键查找 父级行政代码
func (obj *_CnareasMgr) GetBatchFromParentCode(parentCodes []uint64) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("parent_code IN (?)", parentCodes).Find(&results).Error

	return
}

// GetFromAreaCode 通过area_code获取内容 行政代码
func (obj *_CnareasMgr) GetFromAreaCode(areaCode uint64) (result Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("area_code = ?", areaCode).Find(&result).Error

	return
}

// GetBatchFromAreaCode 批量唯一主键查找 行政代码
func (obj *_CnareasMgr) GetBatchFromAreaCode(areaCodes []uint64) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("area_code IN (?)", areaCodes).Find(&results).Error

	return
}

// GetFromZipCode 通过zip_code获取内容 邮政编码
func (obj *_CnareasMgr) GetFromZipCode(zipCode string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("zip_code = ?", zipCode).Find(&results).Error

	return
}

// GetBatchFromZipCode 批量唯一主键查找 邮政编码
func (obj *_CnareasMgr) GetBatchFromZipCode(zipCodes []string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("zip_code IN (?)", zipCodes).Find(&results).Error

	return
}

// GetFromCityCode 通过city_code获取内容 区号
func (obj *_CnareasMgr) GetFromCityCode(cityCode string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("city_code = ?", cityCode).Find(&results).Error

	return
}

// GetBatchFromCityCode 批量唯一主键查找 区号
func (obj *_CnareasMgr) GetBatchFromCityCode(cityCodes []string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("city_code IN (?)", cityCodes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_CnareasMgr) GetFromName(name string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_CnareasMgr) GetBatchFromName(names []string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromShortName 通过short_name获取内容 简称
func (obj *_CnareasMgr) GetFromShortName(shortName string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("short_name = ?", shortName).Find(&results).Error

	return
}

// GetBatchFromShortName 批量唯一主键查找 简称
func (obj *_CnareasMgr) GetBatchFromShortName(shortNames []string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("short_name IN (?)", shortNames).Find(&results).Error

	return
}

// GetFromMergerName 通过merger_name获取内容 组合名
func (obj *_CnareasMgr) GetFromMergerName(mergerName string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("merger_name = ?", mergerName).Find(&results).Error

	return
}

// GetBatchFromMergerName 批量唯一主键查找 组合名
func (obj *_CnareasMgr) GetBatchFromMergerName(mergerNames []string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("merger_name IN (?)", mergerNames).Find(&results).Error

	return
}

// GetFromPinyin 通过pinyin获取内容 拼音
func (obj *_CnareasMgr) GetFromPinyin(pinyin string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pinyin = ?", pinyin).Find(&results).Error

	return
}

// GetBatchFromPinyin 批量唯一主键查找 拼音
func (obj *_CnareasMgr) GetBatchFromPinyin(pinyins []string) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pinyin IN (?)", pinyins).Find(&results).Error

	return
}

// GetFromLng 通过lng获取内容 经度
func (obj *_CnareasMgr) GetFromLng(lng float64) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("lng = ?", lng).Find(&results).Error

	return
}

// GetBatchFromLng 批量唯一主键查找 经度
func (obj *_CnareasMgr) GetBatchFromLng(lngs []float64) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("lng IN (?)", lngs).Find(&results).Error

	return
}

// GetFromLat 通过lat获取内容 纬度
func (obj *_CnareasMgr) GetFromLat(lat float64) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("lat = ?", lat).Find(&results).Error

	return
}

// GetBatchFromLat 批量唯一主键查找 纬度
func (obj *_CnareasMgr) GetBatchFromLat(lats []float64) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("lat IN (?)", lats).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CnareasMgr) FetchByPrimaryKey(id string) (result Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByUkCode primay or index 获取唯一内容
func (obj *_CnareasMgr) FetchUniqueByUkCode(areaCode uint64) (result Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("area_code = ?", areaCode).Find(&result).Error

	return
}

// FetchIndexByIDxParentCode  获取多个内容
func (obj *_CnareasMgr) FetchIndexByIDxParentCode(parentCode uint64) (results []*Cnareas, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("parent_code = ?", parentCode).Find(&results).Error

	return
}
