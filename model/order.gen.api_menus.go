package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIMenusMgr struct {
	*_BaseMgr
}

// APIMenusMgr open func
func APIMenusMgr(db *gorm.DB) *_APIMenusMgr {
	if db == nil {
		panic(fmt.Errorf("APIMenusMgr need init by db"))
	}
	return &_APIMenusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_menus"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIMenusMgr) GetTableName() string {
	return "api_menus"
}

// Get 获取
func (obj *_APIMenusMgr) Get() (result APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIMenusMgr) Gets() (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIMenusMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 菜品名称
func (obj *_APIMenusMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTimeType time_type获取 菜品时段类型
func (obj *_APIMenusMgr) WithTimeType(timeType int) Option {
	return optionFunc(func(o *options) { o.query["time_type"] = timeType })
}

// WithDesc desc获取 菜品介绍
func (obj *_APIMenusMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithStatus status获取 菜品状态：上下架
func (obj *_APIMenusMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithAmount amount获取 销量
func (obj *_APIMenusMgr) WithAmount(amount int) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithPrice price获取 价格
func (obj *_APIMenusMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithCover cover获取 封面
func (obj *_APIMenusMgr) WithCover(cover string) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

// WithSort sort获取 排序
func (obj *_APIMenusMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithPics pics获取 餐品图片
func (obj *_APIMenusMgr) WithPics(pics string) Option {
	return optionFunc(func(o *options) { o.query["pics"] = pics })
}

// WithCreateAt create_at获取
func (obj *_APIMenusMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIMenusMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIMenusMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithApplicationID application_id获取
func (obj *_APIMenusMgr) WithApplicationID(applicationID int) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithMenuTypeID menu_type_id获取 menu_type_id
func (obj *_APIMenusMgr) WithMenuTypeID(menuTypeID int) Option {
	return optionFunc(func(o *options) { o.query["menu_type_id"] = menuTypeID })
}

// GetByOption 功能选项模式获取
func (obj *_APIMenusMgr) GetByOption(opts ...Option) (result APIMenus, err error) {
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
func (obj *_APIMenusMgr) GetByOptions(opts ...Option) (results []*APIMenus, err error) {
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
func (obj *_APIMenusMgr) GetFromID(id int) (result APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIMenusMgr) GetBatchFromID(ids []int) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 菜品名称
func (obj *_APIMenusMgr) GetFromName(name string) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 菜品名称
func (obj *_APIMenusMgr) GetBatchFromName(names []string) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromTimeType 通过time_type获取内容 菜品时段类型
func (obj *_APIMenusMgr) GetFromTimeType(timeType int) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type = ?", timeType).Find(&results).Error

	return
}

// GetBatchFromTimeType 批量唯一主键查找 菜品时段类型
func (obj *_APIMenusMgr) GetBatchFromTimeType(timeTypes []int) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type IN (?)", timeTypes).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 菜品介绍
func (obj *_APIMenusMgr) GetFromDesc(desc string) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 菜品介绍
func (obj *_APIMenusMgr) GetBatchFromDesc(descs []string) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 菜品状态：上下架
func (obj *_APIMenusMgr) GetFromStatus(status bool) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 菜品状态：上下架
func (obj *_APIMenusMgr) GetBatchFromStatus(statuss []bool) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 销量
func (obj *_APIMenusMgr) GetFromAmount(amount int) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量唯一主键查找 销量
func (obj *_APIMenusMgr) GetBatchFromAmount(amounts []int) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount IN (?)", amounts).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_APIMenusMgr) GetFromPrice(price float64) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量唯一主键查找 价格
func (obj *_APIMenusMgr) GetBatchFromPrice(prices []float64) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

// GetFromCover 通过cover获取内容 封面
func (obj *_APIMenusMgr) GetFromCover(cover string) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover = ?", cover).Find(&results).Error

	return
}

// GetBatchFromCover 批量唯一主键查找 封面
func (obj *_APIMenusMgr) GetBatchFromCover(covers []string) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover IN (?)", covers).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_APIMenusMgr) GetFromSort(sort int) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序
func (obj *_APIMenusMgr) GetBatchFromSort(sorts []int) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromPics 通过pics获取内容 餐品图片
func (obj *_APIMenusMgr) GetFromPics(pics string) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pics = ?", pics).Find(&results).Error

	return
}

// GetBatchFromPics 批量唯一主键查找 餐品图片
func (obj *_APIMenusMgr) GetBatchFromPics(picss []string) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pics IN (?)", picss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIMenusMgr) GetFromCreateAt(createAt time.Time) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIMenusMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIMenusMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIMenusMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIMenusMgr) GetFromIsDeleted(isDeleted time.Time) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIMenusMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容
func (obj *_APIMenusMgr) GetFromApplicationID(applicationID int) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量唯一主键查找
func (obj *_APIMenusMgr) GetBatchFromApplicationID(applicationIDs []int) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id IN (?)", applicationIDs).Find(&results).Error

	return
}

// GetFromMenuTypeID 通过menu_type_id获取内容 menu_type_id
func (obj *_APIMenusMgr) GetFromMenuTypeID(menuTypeID int) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_type_id = ?", menuTypeID).Find(&results).Error

	return
}

// GetBatchFromMenuTypeID 批量唯一主键查找 menu_type_id
func (obj *_APIMenusMgr) GetBatchFromMenuTypeID(menuTypeIDs []int) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_type_id IN (?)", menuTypeIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIMenusMgr) FetchByPrimaryKey(id int) (result APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByMenuTypeID  获取多个内容
func (obj *_APIMenusMgr) FetchIndexByMenuTypeID(menuTypeID int) (results []*APIMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_type_id = ?", menuTypeID).Find(&results).Error

	return
}
