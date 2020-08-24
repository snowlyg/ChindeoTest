package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _WechatFansTagsMgr struct {
	*_BaseMgr
}

// WechatFansTagsMgr open func
func WechatFansTagsMgr(db *gorm.DB) *_WechatFansTagsMgr {
	if db == nil {
		panic(fmt.Errorf("WechatFansTagsMgr need init by db"))
	}
	return &_WechatFansTagsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wechat_fans_tags"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WechatFansTagsMgr) GetTableName() string {
	return "wechat_fans_tags"
}

// Get 获取
func (obj *_WechatFansTagsMgr) Get() (result WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WechatFansTagsMgr) Gets() (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 标签ID
func (obj *_WechatFansTagsMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAppid appid获取 公众号APPID
func (obj *_WechatFansTagsMgr) WithAppid(appid string) Option {
	return optionFunc(func(o *options) { o.query["appid"] = appid })
}

// WithName name获取 标签名称
func (obj *_WechatFansTagsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCount count获取 总数
func (obj *_WechatFansTagsMgr) WithCount(count uint64) Option {
	return optionFunc(func(o *options) { o.query["count"] = count })
}

// WithCreateAt create_at获取 创建日期
func (obj *_WechatFansTagsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_WechatFansTagsMgr) GetByOption(opts ...Option) (result WechatFansTags, err error) {
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
func (obj *_WechatFansTagsMgr) GetByOptions(opts ...Option) (results []*WechatFansTags, err error) {
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

// GetFromID 通过id获取内容 标签ID
func (obj *_WechatFansTagsMgr) GetFromID(id uint64) (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量唯一主键查找 标签ID
func (obj *_WechatFansTagsMgr) GetBatchFromID(ids []uint64) (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromAppid 通过appid获取内容 公众号APPID
func (obj *_WechatFansTagsMgr) GetFromAppid(appid string) (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("appid = ?", appid).Find(&results).Error

	return
}

// GetBatchFromAppid 批量唯一主键查找 公众号APPID
func (obj *_WechatFansTagsMgr) GetBatchFromAppid(appids []string) (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("appid IN (?)", appids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 标签名称
func (obj *_WechatFansTagsMgr) GetFromName(name string) (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 标签名称
func (obj *_WechatFansTagsMgr) GetBatchFromName(names []string) (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromCount 通过count获取内容 总数
func (obj *_WechatFansTagsMgr) GetFromCount(count uint64) (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("count = ?", count).Find(&results).Error

	return
}

// GetBatchFromCount 批量唯一主键查找 总数
func (obj *_WechatFansTagsMgr) GetBatchFromCount(counts []uint64) (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("count IN (?)", counts).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建日期
func (obj *_WechatFansTagsMgr) GetFromCreateAt(createAt time.Time) (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建日期
func (obj *_WechatFansTagsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByIndexWechatFansTagsID  获取多个内容
func (obj *_WechatFansTagsMgr) FetchIndexByIndexWechatFansTagsID(id uint64) (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&results).Error

	return
}

// FetchIndexByIndexWechatFansTagsAppid  获取多个内容
func (obj *_WechatFansTagsMgr) FetchIndexByIndexWechatFansTagsAppid(appid string) (results []*WechatFansTags, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("appid = ?", appid).Find(&results).Error

	return
}
