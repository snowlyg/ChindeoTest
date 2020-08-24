package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _SystemAuthMgr struct {
	*_BaseMgr
}

// SystemAuthMgr open func
func SystemAuthMgr(db *gorm.DB) *_SystemAuthMgr {
	if db == nil {
		panic(fmt.Errorf("SystemAuthMgr need init by db"))
	}
	return &_SystemAuthMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_auth"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemAuthMgr) GetTableName() string {
	return "system_auth"
}

// Get 获取
func (obj *_SystemAuthMgr) Get() (result SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemAuthMgr) Gets() (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SystemAuthMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取 权限名称
func (obj *_SystemAuthMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithDesc desc获取 备注说明
func (obj *_SystemAuthMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithSort sort获取 排序权重
func (obj *_SystemAuthMgr) WithSort(sort uint64) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithStatus status获取 权限状态(1使用,0禁用)
func (obj *_SystemAuthMgr) WithStatus(status uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取 创建时间
func (obj *_SystemAuthMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_SystemAuthMgr) GetByOption(opts ...Option) (result SystemAuth, err error) {
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
func (obj *_SystemAuthMgr) GetByOptions(opts ...Option) (results []*SystemAuth, err error) {
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
func (obj *_SystemAuthMgr) GetFromID(id uint64) (result SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_SystemAuthMgr) GetBatchFromID(ids []uint64) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 权限名称
func (obj *_SystemAuthMgr) GetFromTitle(title string) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找 权限名称
func (obj *_SystemAuthMgr) GetBatchFromTitle(titles []string) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 备注说明
func (obj *_SystemAuthMgr) GetFromDesc(desc string) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 备注说明
func (obj *_SystemAuthMgr) GetBatchFromDesc(descs []string) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序权重
func (obj *_SystemAuthMgr) GetFromSort(sort uint64) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序权重
func (obj *_SystemAuthMgr) GetBatchFromSort(sorts []uint64) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 权限状态(1使用,0禁用)
func (obj *_SystemAuthMgr) GetFromStatus(status uint8) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 权限状态(1使用,0禁用)
func (obj *_SystemAuthMgr) GetBatchFromStatus(statuss []uint8) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_SystemAuthMgr) GetFromCreateAt(createAt time.Time) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_SystemAuthMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SystemAuthMgr) FetchByPrimaryKey(id uint64) (result SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxSystemAuthTitle  获取多个内容
func (obj *_SystemAuthMgr) FetchIndexByIDxSystemAuthTitle(title string) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// FetchIndexByIDxSystemAuthStatus  获取多个内容
func (obj *_SystemAuthMgr) FetchIndexByIDxSystemAuthStatus(status uint8) (results []*SystemAuth, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}
