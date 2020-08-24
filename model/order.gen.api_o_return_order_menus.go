package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIOReturnOrderMenusMgr struct {
	*_BaseMgr
}

// APIOReturnOrderMenusMgr open func
func APIOReturnOrderMenusMgr(db *gorm.DB) *_APIOReturnOrderMenusMgr {
	if db == nil {
		panic(fmt.Errorf("APIOReturnOrderMenusMgr need init by db"))
	}
	return &_APIOReturnOrderMenusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_o_return_order_menus"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIOReturnOrderMenusMgr) GetTableName() string {
	return "api_o_return_order_menus"
}

// Get 获取
func (obj *_APIOReturnOrderMenusMgr) Get() (result APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIOReturnOrderMenusMgr) Gets() (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIOReturnOrderMenusMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMenuName menu_name获取 菜品名称
func (obj *_APIOReturnOrderMenusMgr) WithMenuName(menuName string) Option {
	return optionFunc(func(o *options) { o.query["menu_name"] = menuName })
}

// WithMenuType menu_type获取 菜品类型
func (obj *_APIOReturnOrderMenusMgr) WithMenuType(menuType string) Option {
	return optionFunc(func(o *options) { o.query["menu_type"] = menuType })
}

// WithMenuTimeType menu_time_type获取 菜品时段类型
func (obj *_APIOReturnOrderMenusMgr) WithMenuTimeType(menuTimeType string) Option {
	return optionFunc(func(o *options) { o.query["menu_time_type"] = menuTimeType })
}

// WithMenuDesc menu_desc获取 菜品介绍
func (obj *_APIOReturnOrderMenusMgr) WithMenuDesc(menuDesc string) Option {
	return optionFunc(func(o *options) { o.query["menu_desc"] = menuDesc })
}

// WithMenuID menu_id获取 menu_id
func (obj *_APIOReturnOrderMenusMgr) WithMenuID(menuID int) Option {
	return optionFunc(func(o *options) { o.query["menu_id"] = menuID })
}

// WithPrice price获取 价格
func (obj *_APIOReturnOrderMenusMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithAmount amount获取 数量
func (obj *_APIOReturnOrderMenusMgr) WithAmount(amount int) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithCover cover获取 菜品图片
func (obj *_APIOReturnOrderMenusMgr) WithCover(cover string) Option {
	return optionFunc(func(o *options) { o.query["cover"] = cover })
}

// WithCreateAt create_at获取
func (obj *_APIOReturnOrderMenusMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIOReturnOrderMenusMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIOReturnOrderMenusMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithApplicationID application_id获取 application_id
func (obj *_APIOReturnOrderMenusMgr) WithApplicationID(applicationID int) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithOReturnOrderID o_return_order_id获取 o_return_order_id
func (obj *_APIOReturnOrderMenusMgr) WithOReturnOrderID(oReturnOrderID int) Option {
	return optionFunc(func(o *options) { o.query["o_return_order_id"] = oReturnOrderID })
}

// GetByOption 功能选项模式获取
func (obj *_APIOReturnOrderMenusMgr) GetByOption(opts ...Option) (result APIOReturnOrderMenus, err error) {
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
func (obj *_APIOReturnOrderMenusMgr) GetByOptions(opts ...Option) (results []*APIOReturnOrderMenus, err error) {
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
func (obj *_APIOReturnOrderMenusMgr) GetFromID(id int) (result APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromID(ids []int) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromMenuName 通过menu_name获取内容 菜品名称
func (obj *_APIOReturnOrderMenusMgr) GetFromMenuName(menuName string) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_name = ?", menuName).Find(&results).Error

	return
}

// GetBatchFromMenuName 批量唯一主键查找 菜品名称
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromMenuName(menuNames []string) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_name IN (?)", menuNames).Find(&results).Error

	return
}

// GetFromMenuType 通过menu_type获取内容 菜品类型
func (obj *_APIOReturnOrderMenusMgr) GetFromMenuType(menuType string) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_type = ?", menuType).Find(&results).Error

	return
}

// GetBatchFromMenuType 批量唯一主键查找 菜品类型
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromMenuType(menuTypes []string) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_type IN (?)", menuTypes).Find(&results).Error

	return
}

// GetFromMenuTimeType 通过menu_time_type获取内容 菜品时段类型
func (obj *_APIOReturnOrderMenusMgr) GetFromMenuTimeType(menuTimeType string) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_time_type = ?", menuTimeType).Find(&results).Error

	return
}

// GetBatchFromMenuTimeType 批量唯一主键查找 菜品时段类型
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromMenuTimeType(menuTimeTypes []string) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_time_type IN (?)", menuTimeTypes).Find(&results).Error

	return
}

// GetFromMenuDesc 通过menu_desc获取内容 菜品介绍
func (obj *_APIOReturnOrderMenusMgr) GetFromMenuDesc(menuDesc string) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_desc = ?", menuDesc).Find(&results).Error

	return
}

// GetBatchFromMenuDesc 批量唯一主键查找 菜品介绍
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromMenuDesc(menuDescs []string) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_desc IN (?)", menuDescs).Find(&results).Error

	return
}

// GetFromMenuID 通过menu_id获取内容 menu_id
func (obj *_APIOReturnOrderMenusMgr) GetFromMenuID(menuID int) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_id = ?", menuID).Find(&results).Error

	return
}

// GetBatchFromMenuID 批量唯一主键查找 menu_id
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromMenuID(menuIDs []int) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_id IN (?)", menuIDs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_APIOReturnOrderMenusMgr) GetFromPrice(price float64) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量唯一主键查找 价格
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromPrice(prices []float64) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 数量
func (obj *_APIOReturnOrderMenusMgr) GetFromAmount(amount int) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量唯一主键查找 数量
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromAmount(amounts []int) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount IN (?)", amounts).Find(&results).Error

	return
}

// GetFromCover 通过cover获取内容 菜品图片
func (obj *_APIOReturnOrderMenusMgr) GetFromCover(cover string) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover = ?", cover).Find(&results).Error

	return
}

// GetBatchFromCover 批量唯一主键查找 菜品图片
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromCover(covers []string) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover IN (?)", covers).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIOReturnOrderMenusMgr) GetFromCreateAt(createAt time.Time) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIOReturnOrderMenusMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIOReturnOrderMenusMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容 application_id
func (obj *_APIOReturnOrderMenusMgr) GetFromApplicationID(applicationID int) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量唯一主键查找 application_id
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromApplicationID(applicationIDs []int) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id IN (?)", applicationIDs).Find(&results).Error

	return
}

// GetFromOReturnOrderID 通过o_return_order_id获取内容 o_return_order_id
func (obj *_APIOReturnOrderMenusMgr) GetFromOReturnOrderID(oReturnOrderID int) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("o_return_order_id = ?", oReturnOrderID).Find(&results).Error

	return
}

// GetBatchFromOReturnOrderID 批量唯一主键查找 o_return_order_id
func (obj *_APIOReturnOrderMenusMgr) GetBatchFromOReturnOrderID(oReturnOrderIDs []int) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("o_return_order_id IN (?)", oReturnOrderIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIOReturnOrderMenusMgr) FetchByPrimaryKey(id int) (result APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByOReturnOrderID  获取多个内容
func (obj *_APIOReturnOrderMenusMgr) FetchIndexByOReturnOrderID(menuID int, oReturnOrderID int) (results []*APIOReturnOrderMenus, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("menu_id = ? AND o_return_order_id = ?", menuID, oReturnOrderID).Find(&results).Error

	return
}
