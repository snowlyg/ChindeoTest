package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _SystemMenuMgr struct {
	*_BaseMgr
}

// SystemMenuMgr open func
func SystemMenuMgr(db *gorm.DB) *_SystemMenuMgr {
	if db == nil {
		panic(fmt.Errorf("SystemMenuMgr need init by db"))
	}
	return &_SystemMenuMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_menu"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemMenuMgr) GetTableName() string {
	return "system_menu"
}

// Get 获取
func (obj *_SystemMenuMgr) Get() (result SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemMenuMgr) Gets() (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SystemMenuMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPid pid获取 上级ID
func (obj *_SystemMenuMgr) WithPid(pid uint64) Option {
	return optionFunc(func(o *options) { o.query["pid"] = pid })
}

// WithTitle title获取 菜单名称
func (obj *_SystemMenuMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithIcon icon获取 菜单图标
func (obj *_SystemMenuMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithNode node获取 节点代码
func (obj *_SystemMenuMgr) WithNode(node string) Option {
	return optionFunc(func(o *options) { o.query["node"] = node })
}

// WithURL url获取 链接节点
func (obj *_SystemMenuMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithParams params获取 链接参数
func (obj *_SystemMenuMgr) WithParams(params string) Option {
	return optionFunc(func(o *options) { o.query["params"] = params })
}

// WithTarget target获取 打开方式
func (obj *_SystemMenuMgr) WithTarget(target string) Option {
	return optionFunc(func(o *options) { o.query["target"] = target })
}

// WithSort sort获取 排序权重
func (obj *_SystemMenuMgr) WithSort(sort uint) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithStatus status获取 状态(0:禁用,1:启用)
func (obj *_SystemMenuMgr) WithStatus(status uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取 创建时间
func (obj *_SystemMenuMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_SystemMenuMgr) GetByOption(opts ...Option) (result SystemMenu, err error) {
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
func (obj *_SystemMenuMgr) GetByOptions(opts ...Option) (results []*SystemMenu, err error) {
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
func (obj *_SystemMenuMgr) GetFromID(id uint64) (result SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_SystemMenuMgr) GetBatchFromID(ids []uint64) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromPid 通过pid获取内容 上级ID
func (obj *_SystemMenuMgr) GetFromPid(pid uint64) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pid = ?", pid).Find(&results).Error

	return
}

// GetBatchFromPid 批量唯一主键查找 上级ID
func (obj *_SystemMenuMgr) GetBatchFromPid(pids []uint64) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("pid IN (?)", pids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 菜单名称
func (obj *_SystemMenuMgr) GetFromTitle(title string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找 菜单名称
func (obj *_SystemMenuMgr) GetBatchFromTitle(titles []string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 菜单图标
func (obj *_SystemMenuMgr) GetFromIcon(icon string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找 菜单图标
func (obj *_SystemMenuMgr) GetBatchFromIcon(icons []string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

// GetFromNode 通过node获取内容 节点代码
func (obj *_SystemMenuMgr) GetFromNode(node string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("node = ?", node).Find(&results).Error

	return
}

// GetBatchFromNode 批量唯一主键查找 节点代码
func (obj *_SystemMenuMgr) GetBatchFromNode(nodes []string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("node IN (?)", nodes).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 链接节点
func (obj *_SystemMenuMgr) GetFromURL(url string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量唯一主键查找 链接节点
func (obj *_SystemMenuMgr) GetBatchFromURL(urls []string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error

	return
}

// GetFromParams 通过params获取内容 链接参数
func (obj *_SystemMenuMgr) GetFromParams(params string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("params = ?", params).Find(&results).Error

	return
}

// GetBatchFromParams 批量唯一主键查找 链接参数
func (obj *_SystemMenuMgr) GetBatchFromParams(paramss []string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("params IN (?)", paramss).Find(&results).Error

	return
}

// GetFromTarget 通过target获取内容 打开方式
func (obj *_SystemMenuMgr) GetFromTarget(target string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("target = ?", target).Find(&results).Error

	return
}

// GetBatchFromTarget 批量唯一主键查找 打开方式
func (obj *_SystemMenuMgr) GetBatchFromTarget(targets []string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("target IN (?)", targets).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序权重
func (obj *_SystemMenuMgr) GetFromSort(sort uint) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序权重
func (obj *_SystemMenuMgr) GetBatchFromSort(sorts []uint) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态(0:禁用,1:启用)
func (obj *_SystemMenuMgr) GetFromStatus(status uint8) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态(0:禁用,1:启用)
func (obj *_SystemMenuMgr) GetBatchFromStatus(statuss []uint8) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_SystemMenuMgr) GetFromCreateAt(createAt time.Time) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_SystemMenuMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SystemMenuMgr) FetchByPrimaryKey(id uint64) (result SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxSystemMenuNode  获取多个内容
func (obj *_SystemMenuMgr) FetchIndexByIDxSystemMenuNode(node string) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("node = ?", node).Find(&results).Error

	return
}

// FetchIndexByIDxSystemMenuStatus  获取多个内容
func (obj *_SystemMenuMgr) FetchIndexByIDxSystemMenuStatus(status uint8) (results []*SystemMenu, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}
