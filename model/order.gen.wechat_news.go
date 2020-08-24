package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _WechatNewsMgr struct {
	*_BaseMgr
}

// WechatNewsMgr open func
func WechatNewsMgr(db *gorm.DB) *_WechatNewsMgr {
	if db == nil {
		panic(fmt.Errorf("WechatNewsMgr need init by db"))
	}
	return &_WechatNewsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wechat_news"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WechatNewsMgr) GetTableName() string {
	return "wechat_news"
}

// Get 获取
func (obj *_WechatNewsMgr) Get() (result WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WechatNewsMgr) Gets() (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WechatNewsMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMediaID media_id获取 永久素材MediaID
func (obj *_WechatNewsMgr) WithMediaID(mediaID string) Option {
	return optionFunc(func(o *options) { o.query["media_id"] = mediaID })
}

// WithLocalURL local_url获取 永久素材外网URL
func (obj *_WechatNewsMgr) WithLocalURL(localURL string) Option {
	return optionFunc(func(o *options) { o.query["local_url"] = localURL })
}

// WithArticleID article_id获取 关联图文ID(用英文逗号做分割)
func (obj *_WechatNewsMgr) WithArticleID(articleID string) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithIsDeleted is_deleted获取 删除状态(0未删除,1已删除)
func (obj *_WechatNewsMgr) WithIsDeleted(isDeleted uint8) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithCreateAt create_at获取 创建时间
func (obj *_WechatNewsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithCreateBy create_by获取 创建人
func (obj *_WechatNewsMgr) WithCreateBy(createBy int64) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// GetByOption 功能选项模式获取
func (obj *_WechatNewsMgr) GetByOption(opts ...Option) (result WechatNews, err error) {
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
func (obj *_WechatNewsMgr) GetByOptions(opts ...Option) (results []*WechatNews, err error) {
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
func (obj *_WechatNewsMgr) GetFromID(id uint64) (result WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WechatNewsMgr) GetBatchFromID(ids []uint64) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromMediaID 通过media_id获取内容 永久素材MediaID
func (obj *_WechatNewsMgr) GetFromMediaID(mediaID string) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("media_id = ?", mediaID).Find(&results).Error

	return
}

// GetBatchFromMediaID 批量唯一主键查找 永久素材MediaID
func (obj *_WechatNewsMgr) GetBatchFromMediaID(mediaIDs []string) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("media_id IN (?)", mediaIDs).Find(&results).Error

	return
}

// GetFromLocalURL 通过local_url获取内容 永久素材外网URL
func (obj *_WechatNewsMgr) GetFromLocalURL(localURL string) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("local_url = ?", localURL).Find(&results).Error

	return
}

// GetBatchFromLocalURL 批量唯一主键查找 永久素材外网URL
func (obj *_WechatNewsMgr) GetBatchFromLocalURL(localURLs []string) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("local_url IN (?)", localURLs).Find(&results).Error

	return
}

// GetFromArticleID 通过article_id获取内容 关联图文ID(用英文逗号做分割)
func (obj *_WechatNewsMgr) GetFromArticleID(articleID string) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}

// GetBatchFromArticleID 批量唯一主键查找 关联图文ID(用英文逗号做分割)
func (obj *_WechatNewsMgr) GetBatchFromArticleID(articleIDs []string) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("article_id IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容 删除状态(0未删除,1已删除)
func (obj *_WechatNewsMgr) GetFromIsDeleted(isDeleted uint8) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找 删除状态(0未删除,1已删除)
func (obj *_WechatNewsMgr) GetBatchFromIsDeleted(isDeleteds []uint8) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_WechatNewsMgr) GetFromCreateAt(createAt time.Time) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_WechatNewsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建人
func (obj *_WechatNewsMgr) GetFromCreateBy(createBy int64) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量唯一主键查找 创建人
func (obj *_WechatNewsMgr) GetBatchFromCreateBy(createBys []int64) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WechatNewsMgr) FetchByPrimaryKey(id uint64) (result WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIndexWechatNewsMediaID  获取多个内容
func (obj *_WechatNewsMgr) FetchIndexByIndexWechatNewsMediaID(mediaID string) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("media_id = ?", mediaID).Find(&results).Error

	return
}

// FetchIndexByIndexWechatNewsArtcleID  获取多个内容
func (obj *_WechatNewsMgr) FetchIndexByIndexWechatNewsArtcleID(articleID string) (results []*WechatNews, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("article_id = ?", articleID).Find(&results).Error

	return
}
