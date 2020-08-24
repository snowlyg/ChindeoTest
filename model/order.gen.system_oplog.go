package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _SystemOplogMgr struct {
	*_BaseMgr
}

// SystemOplogMgr open func
func SystemOplogMgr(db *gorm.DB) *_SystemOplogMgr {
	if db == nil {
		panic(fmt.Errorf("SystemOplogMgr need init by db"))
	}
	return &_SystemOplogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_oplog"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemOplogMgr) GetTableName() string {
	return "system_oplog"
}

// Get 获取
func (obj *_SystemOplogMgr) Get() (result SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemOplogMgr) Gets() (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SystemOplogMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNode node获取 当前操作节点
func (obj *_SystemOplogMgr) WithNode(node string) Option {
	return optionFunc(func(o *options) { o.query["node"] = node })
}

// WithGeoip geoip获取 操作者IP地址
func (obj *_SystemOplogMgr) WithGeoip(geoip string) Option {
	return optionFunc(func(o *options) { o.query["geoip"] = geoip })
}

// WithAction action获取 操作行为名称
func (obj *_SystemOplogMgr) WithAction(action string) Option {
	return optionFunc(func(o *options) { o.query["action"] = action })
}

// WithContent content获取 操作内容描述
func (obj *_SystemOplogMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithUsername username获取 操作人用户名
func (obj *_SystemOplogMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithCreateAt create_at获取 创建时间
func (obj *_SystemOplogMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_SystemOplogMgr) GetByOption(opts ...Option) (result SystemOplog, err error) {
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
func (obj *_SystemOplogMgr) GetByOptions(opts ...Option) (results []*SystemOplog, err error) {
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
func (obj *_SystemOplogMgr) GetFromID(id uint64) (result SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_SystemOplogMgr) GetBatchFromID(ids []uint64) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromNode 通过node获取内容 当前操作节点
func (obj *_SystemOplogMgr) GetFromNode(node string) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("node = ?", node).Find(&results).Error

	return
}

// GetBatchFromNode 批量唯一主键查找 当前操作节点
func (obj *_SystemOplogMgr) GetBatchFromNode(nodes []string) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("node IN (?)", nodes).Find(&results).Error

	return
}

// GetFromGeoip 通过geoip获取内容 操作者IP地址
func (obj *_SystemOplogMgr) GetFromGeoip(geoip string) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("geoip = ?", geoip).Find(&results).Error

	return
}

// GetBatchFromGeoip 批量唯一主键查找 操作者IP地址
func (obj *_SystemOplogMgr) GetBatchFromGeoip(geoips []string) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("geoip IN (?)", geoips).Find(&results).Error

	return
}

// GetFromAction 通过action获取内容 操作行为名称
func (obj *_SystemOplogMgr) GetFromAction(action string) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("action = ?", action).Find(&results).Error

	return
}

// GetBatchFromAction 批量唯一主键查找 操作行为名称
func (obj *_SystemOplogMgr) GetBatchFromAction(actions []string) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("action IN (?)", actions).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 操作内容描述
func (obj *_SystemOplogMgr) GetFromContent(content string) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找 操作内容描述
func (obj *_SystemOplogMgr) GetBatchFromContent(contents []string) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 操作人用户名
func (obj *_SystemOplogMgr) GetFromUsername(username string) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量唯一主键查找 操作人用户名
func (obj *_SystemOplogMgr) GetBatchFromUsername(usernames []string) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username IN (?)", usernames).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_SystemOplogMgr) GetFromCreateAt(createAt time.Time) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_SystemOplogMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SystemOplogMgr) FetchByPrimaryKey(id uint64) (result SystemOplog, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
