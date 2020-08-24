package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIAuthsMgr struct {
	*_BaseMgr
}

// APIAuthsMgr open func
func APIAuthsMgr(db *gorm.DB) *_APIAuthsMgr {
	if db == nil {
		panic(fmt.Errorf("APIAuthsMgr need init by db"))
	}
	return &_APIAuthsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_auths"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIAuthsMgr) GetTableName() string {
	return "api_auths"
}

// Get 获取
func (obj *_APIAuthsMgr) Get() (result APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIAuthsMgr) Gets() (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIAuthsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 权限名称
func (obj *_APIAuthsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithController controller获取 控制器
func (obj *_APIAuthsMgr) WithController(controller string) Option {
	return optionFunc(func(o *options) { o.query["controller"] = controller })
}

// WithAction action获取 操作方法
func (obj *_APIAuthsMgr) WithAction(action string) Option {
	return optionFunc(func(o *options) { o.query["action"] = action })
}

// WithMethod method获取 请求方法
func (obj *_APIAuthsMgr) WithMethod(method string) Option {
	return optionFunc(func(o *options) { o.query["method"] = method })
}

// WithDesc desc获取 描述
func (obj *_APIAuthsMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithRemark remark获取 备注
func (obj *_APIAuthsMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithStatus status获取 账号状态
func (obj *_APIAuthsMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取
func (obj *_APIAuthsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIAuthsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIAuthsMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// GetByOption 功能选项模式获取
func (obj *_APIAuthsMgr) GetByOption(opts ...Option) (result APIAuths, err error) {
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
func (obj *_APIAuthsMgr) GetByOptions(opts ...Option) (results []*APIAuths, err error) {
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
func (obj *_APIAuthsMgr) GetFromID(id int) (result APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIAuthsMgr) GetBatchFromID(ids []int) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 权限名称
func (obj *_APIAuthsMgr) GetFromName(name string) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 权限名称
func (obj *_APIAuthsMgr) GetBatchFromName(names []string) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromController 通过controller获取内容 控制器
func (obj *_APIAuthsMgr) GetFromController(controller string) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("controller = ?", controller).Find(&results).Error

	return
}

// GetBatchFromController 批量唯一主键查找 控制器
func (obj *_APIAuthsMgr) GetBatchFromController(controllers []string) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("controller IN (?)", controllers).Find(&results).Error

	return
}

// GetFromAction 通过action获取内容 操作方法
func (obj *_APIAuthsMgr) GetFromAction(action string) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("action = ?", action).Find(&results).Error

	return
}

// GetBatchFromAction 批量唯一主键查找 操作方法
func (obj *_APIAuthsMgr) GetBatchFromAction(actions []string) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("action IN (?)", actions).Find(&results).Error

	return
}

// GetFromMethod 通过method获取内容 请求方法
func (obj *_APIAuthsMgr) GetFromMethod(method string) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("method = ?", method).Find(&results).Error

	return
}

// GetBatchFromMethod 批量唯一主键查找 请求方法
func (obj *_APIAuthsMgr) GetBatchFromMethod(methods []string) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("method IN (?)", methods).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 描述
func (obj *_APIAuthsMgr) GetFromDesc(desc string) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 描述
func (obj *_APIAuthsMgr) GetBatchFromDesc(descs []string) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_APIAuthsMgr) GetFromRemark(remark string) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_APIAuthsMgr) GetBatchFromRemark(remarks []string) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 账号状态
func (obj *_APIAuthsMgr) GetFromStatus(status bool) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 账号状态
func (obj *_APIAuthsMgr) GetBatchFromStatus(statuss []bool) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIAuthsMgr) GetFromCreateAt(createAt time.Time) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIAuthsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIAuthsMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIAuthsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIAuthsMgr) GetFromIsDeleted(isDeleted time.Time) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIAuthsMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIAuthsMgr) FetchByPrimaryKey(id int) (result APIAuths, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
