package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIRolesMgr struct {
	*_BaseMgr
}

// APIRolesMgr open func
func APIRolesMgr(db *gorm.DB) *_APIRolesMgr {
	if db == nil {
		panic(fmt.Errorf("APIRolesMgr need init by db"))
	}
	return &_APIRolesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_roles"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIRolesMgr) GetTableName() string {
	return "api_roles"
}

// Get 获取
func (obj *_APIRolesMgr) Get() (result APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIRolesMgr) Gets() (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIRolesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 角色名称
func (obj *_APIRolesMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDesc desc获取 描述
func (obj *_APIRolesMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithRemark remark获取 备注
func (obj *_APIRolesMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithStatus status获取 账号状态
func (obj *_APIRolesMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取
func (obj *_APIRolesMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIRolesMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIRolesMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// GetByOption 功能选项模式获取
func (obj *_APIRolesMgr) GetByOption(opts ...Option) (result APIRoles, err error) {
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
func (obj *_APIRolesMgr) GetByOptions(opts ...Option) (results []*APIRoles, err error) {
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
func (obj *_APIRolesMgr) GetFromID(id int) (result APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIRolesMgr) GetBatchFromID(ids []int) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 角色名称
func (obj *_APIRolesMgr) GetFromName(name string) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 角色名称
func (obj *_APIRolesMgr) GetBatchFromName(names []string) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 描述
func (obj *_APIRolesMgr) GetFromDesc(desc string) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 描述
func (obj *_APIRolesMgr) GetBatchFromDesc(descs []string) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_APIRolesMgr) GetFromRemark(remark string) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_APIRolesMgr) GetBatchFromRemark(remarks []string) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 账号状态
func (obj *_APIRolesMgr) GetFromStatus(status bool) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 账号状态
func (obj *_APIRolesMgr) GetBatchFromStatus(statuss []bool) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIRolesMgr) GetFromCreateAt(createAt time.Time) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIRolesMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIRolesMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIRolesMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIRolesMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIRolesMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIRolesMgr) FetchByPrimaryKey(id int) (result APIRoles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
