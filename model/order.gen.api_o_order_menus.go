package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIOOrderMenusMgr struct {
	*_BaseMgr
}

// APIOOrderMenusMgr open func
func APIOOrderMenusMgr(db *gorm.DB) *_APIOOrderMenusMgr {
	if db == nil {
		panic(fmt.Errorf("APIOOrderMenusMgr need init by db"))
	}
	return &_APIOOrderMenusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_o_order_menus"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIOOrderMenusMgr) GetTableName() string {
	return "api_o_order_menus"
}

// Get 获取
func (obj *_APIOOrderMenusMgr) Get() (result APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIOOrderMenusMgr) Gets() (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIOOrderMenusMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMenuName menu_name获取 菜品名称
func (obj *_APIOOrderMenusMgr) WithMenuName(menuName string) Option {
	return optionFunc(func(o *options) { o.query["menu_name"] = menuName })
}

// WithMenuType menu_type获取 菜品类型
func (obj *_APIOOrderMenusMgr) WithMenuType(menuType string) Option {
	return optionFunc(func(o *options) { o.query["menu_type"] = menuType })
}

// WithMenuTimeType menu_time_type获取 菜品时段类型
func (obj *_APIOOrderMenusMgr) WithMenuTimeType(menuTimeType string) Option {
	return optionFunc(func(o *options) { o.query["menu_time_type"] = menuTimeType })
}

// WithMenuDesc menu_desc获取 菜品介绍
func (obj *_APIOOrderMenusMgr) WithMenuDesc(menuDesc string) Option {
	return optionFunc(func(o *options) { o.query["menu_desc"] = menuDesc })
}

// WithMenuID menu_id获取 menu_id
func (obj *_APIOOrderMenusMgr) WithMenuID(menuID int) Option {
	return optionFunc(func(o *options) { o.query["menu_id"] = menuID })
}

// WithPrice price获取 价格
func (obj *_APIOOrderMenusMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithAmount amount获取 数量
func (obj *_APIOOrderMenusMgr) WithAmount(amount int) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithCover cover获取 菜品图片
func (obj *_APIOOrderMenusMgr) WithCover(cover string) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

// WithCreateAt create_at获取
func (obj *_APIOOrderMenusMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIOOrderMenusMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIOOrderMenusMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithOOrderID o_order_id获取 o_order_id
func (obj *_APIOOrderMenusMgr) WithOOrderID(oOrderID int) Option {
	return optionFunc(func(o *options) { o.query["o_order_id"] = oOrderID })
}

// GetByOption 功能选项模式获取
func (obj *_APIOOrderMenusMgr) GetByOption(opts ...Option) (result APIOOrderMenus, err error) {
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
func (obj *_APIOOrderMenusMgr) GetByOptions(opts ...Option) (results []*APIOOrderMenus, err error) {
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
func (obj *_APIOOrderMenusMgr) GetFromID(id int) (result APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIOOrderMenusMgr) GetBatchFromID(ids []int) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromMenuName 通过menu_name获取内容 菜品名称
func (obj *_APIOOrderMenusMgr) GetFromMenuName(menuName string) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_name = ?", menuName).Find(&results).Error

	return
}

// GetBatchFromMenuName 批量唯一主键查找 菜品名称
func (obj *_APIOOrderMenusMgr) GetBatchFromMenuName(menuNames []string) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_name IN (?)", menuNames).Find(&results).Error

	return
}

// GetFromMenuType 通过menu_type获取内容 菜品类型
func (obj *_APIOOrderMenusMgr) GetFromMenuType(menuType string) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_type = ?", menuType).Find(&results).Error

	return
}

// GetBatchFromMenuType 批量唯一主键查找 菜品类型
func (obj *_APIOOrderMenusMgr) GetBatchFromMenuType(menuTypes []string) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_type IN (?)", menuTypes).Find(&results).Error

	return
}

// GetFromMenuTimeType 通过menu_time_type获取内容 菜品时段类型
func (obj *_APIOOrderMenusMgr) GetFromMenuTimeType(menuTimeType string) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_time_type = ?", menuTimeType).Find(&results).Error

	return
}

// GetBatchFromMenuTimeType 批量唯一主键查找 菜品时段类型
func (obj *_APIOOrderMenusMgr) GetBatchFromMenuTimeType(menuTimeTypes []string) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_time_type IN (?)", menuTimeTypes).Find(&results).Error

	return
}

// GetFromMenuDesc 通过menu_desc获取内容 菜品介绍
func (obj *_APIOOrderMenusMgr) GetFromMenuDesc(menuDesc string) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_desc = ?", menuDesc).Find(&results).Error

	return
}

// GetBatchFromMenuDesc 批量唯一主键查找 菜品介绍
func (obj *_APIOOrderMenusMgr) GetBatchFromMenuDesc(menuDescs []string) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_desc IN (?)", menuDescs).Find(&results).Error

	return
}

// GetFromMenuID 通过menu_id获取内容 menu_id
func (obj *_APIOOrderMenusMgr) GetFromMenuID(menuID int) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_id = ?", menuID).Find(&results).Error

	return
}

// GetBatchFromMenuID 批量唯一主键查找 menu_id
func (obj *_APIOOrderMenusMgr) GetBatchFromMenuID(menuIDs []int) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_id IN (?)", menuIDs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_APIOOrderMenusMgr) GetFromPrice(price float64) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量唯一主键查找 价格
func (obj *_APIOOrderMenusMgr) GetBatchFromPrice(prices []float64) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 数量
func (obj *_APIOOrderMenusMgr) GetFromAmount(amount int) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量唯一主键查找 数量
func (obj *_APIOOrderMenusMgr) GetBatchFromAmount(amounts []int) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount IN (?)", amounts).Find(&results).Error

	return
}

// GetFromCover 通过cover获取内容 菜品图片
func (obj *_APIOOrderMenusMgr) GetFromCover(cover string) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover = ?", cover).Find(&results).Error

	return
}

// GetBatchFromCover 批量唯一主键查找 菜品图片
func (obj *_APIOOrderMenusMgr) GetBatchFromCover(covers []string) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover IN (?)", covers).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIOOrderMenusMgr) GetFromCreateAt(createAt time.Time) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIOOrderMenusMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIOOrderMenusMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIOOrderMenusMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIOOrderMenusMgr) GetFromIsDeleted(isDeleted time.Time) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIOOrderMenusMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromOOrderID 通过o_order_id获取内容 o_order_id
func (obj *_APIOOrderMenusMgr) GetFromOOrderID(oOrderID int) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("o_order_id = ?", oOrderID).Find(&results).Error

	return
}

// GetBatchFromOOrderID 批量唯一主键查找 o_order_id
func (obj *_APIOOrderMenusMgr) GetBatchFromOOrderID(oOrderIDs []int) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("o_order_id IN (?)", oOrderIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIOOrderMenusMgr) FetchByPrimaryKey(id int) (result APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByOOrderID  获取多个内容
func (obj *_APIOOrderMenusMgr) FetchIndexByOOrderID(menuID int, oOrderID int) (results []*APIOOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_id = ? AND o_order_id = ?", menuID, oOrderID).Find(&results).Error

	return
}
