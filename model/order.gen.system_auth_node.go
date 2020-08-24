package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _SystemAuthNodeMgr struct {
	*_BaseMgr
}

// SystemAuthNodeMgr open func
func SystemAuthNodeMgr(db *gorm.DB) *_SystemAuthNodeMgr {
	if db == nil {
		panic(fmt.Errorf("SystemAuthNodeMgr need init by db"))
	}
	return &_SystemAuthNodeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_auth_node"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemAuthNodeMgr) GetTableName() string {
	return "system_auth_node"
}

// Get 获取
func (obj *_SystemAuthNodeMgr) Get() (result SystemAuthNode, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemAuthNodeMgr) Gets() (results []*SystemAuthNode, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SystemAuthNodeMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAuth auth获取 角色
func (obj *_SystemAuthNodeMgr) WithAuth(auth uint64) Option {
	return optionFunc(func(o *options) { o.query["auth"] = auth })
}

// WithNode node获取 节点
func (obj *_SystemAuthNodeMgr) WithNode(node string) Option {
	return optionFunc(func(o *options) { o.query["node"] = node })
}

// GetByOption 功能选项模式获取
func (obj *_SystemAuthNodeMgr) GetByOption(opts ...Option) (result SystemAuthNode, err error) {
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
func (obj *_SystemAuthNodeMgr) GetByOptions(opts ...Option) (results []*SystemAuthNode, err error) {
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
func (obj *_SystemAuthNodeMgr) GetFromID(id uint64) (result SystemAuthNode, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_SystemAuthNodeMgr) GetBatchFromID(ids []uint64) (results []*SystemAuthNode, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromAuth 通过auth获取内容 角色
func (obj *_SystemAuthNodeMgr) GetFromAuth(auth uint64) (results []*SystemAuthNode, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("auth = ?", auth).Find(&results).Error

	return
}

// GetBatchFromAuth 批量唯一主键查找 角色
func (obj *_SystemAuthNodeMgr) GetBatchFromAuth(auths []uint64) (results []*SystemAuthNode, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("auth IN (?)", auths).Find(&results).Error

	return
}

// GetFromNode 通过node获取内容 节点
func (obj *_SystemAuthNodeMgr) GetFromNode(node string) (results []*SystemAuthNode, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("node = ?", node).Find(&results).Error

	return
}

// GetBatchFromNode 批量唯一主键查找 节点
func (obj *_SystemAuthNodeMgr) GetBatchFromNode(nodes []string) (results []*SystemAuthNode, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("node IN (?)", nodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SystemAuthNodeMgr) FetchByPrimaryKey(id uint64) (result SystemAuthNode, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxSystemAuthAuth  获取多个内容
func (obj *_SystemAuthNodeMgr) FetchIndexByIDxSystemAuthAuth(auth uint64) (results []*SystemAuthNode, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("auth = ?", auth).Find(&results).Error

	return
}

// FetchIndexByIDxSystemAuthNode  获取多个内容
func (obj *_SystemAuthNodeMgr) FetchIndexByIDxSystemAuthNode(node string) (results []*SystemAuthNode, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("node = ?", node).Find(&results).Error

	return
}
