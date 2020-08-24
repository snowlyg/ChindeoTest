package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _WechatMediaMgr struct {
	*_BaseMgr
}

// WechatMediaMgr open func
func WechatMediaMgr(db *gorm.DB) *_WechatMediaMgr {
	if db == nil {
		panic(fmt.Errorf("WechatMediaMgr need init by db"))
	}
	return &_WechatMediaMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wechat_media"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WechatMediaMgr) GetTableName() string {
	return "wechat_media"
}

// Get 获取
func (obj *_WechatMediaMgr) Get() (result WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WechatMediaMgr) Gets() (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WechatMediaMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAppid appid获取 公众号ID
func (obj *_WechatMediaMgr) WithAppid(appid string) Option {
	return optionFunc(func(o *options) { o.query["appid"] = appid })
}

// WithMd5 md5获取 文件md5
func (obj *_WechatMediaMgr) WithMd5(md5 string) Option {
	return optionFunc(func(o *options) { o.query["md5"] = md5 })
}

// WithType type获取 媒体类型
func (obj *_WechatMediaMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithMediaID media_id获取 永久素材MediaID
func (obj *_WechatMediaMgr) WithMediaID(mediaID string) Option {
	return optionFunc(func(o *options) { o.query["media_id"] = mediaID })
}

// WithLocalURL local_url获取 本地文件链接
func (obj *_WechatMediaMgr) WithLocalURL(localURL string) Option {
	return optionFunc(func(o *options) { o.query["local_url"] = localURL })
}

// WithMediaURL media_url获取 远程图片链接
func (obj *_WechatMediaMgr) WithMediaURL(mediaURL string) Option {
	return optionFunc(func(o *options) { o.query["media_url"] = mediaURL })
}

// WithCreateAt create_at获取 创建时间
func (obj *_WechatMediaMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_WechatMediaMgr) GetByOption(opts ...Option) (result WechatMedia, err error) {
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
func (obj *_WechatMediaMgr) GetByOptions(opts ...Option) (results []*WechatMedia, err error) {
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
func (obj *_WechatMediaMgr) GetFromID(id uint64) (result WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WechatMediaMgr) GetBatchFromID(ids []uint64) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromAppid 通过appid获取内容 公众号ID
func (obj *_WechatMediaMgr) GetFromAppid(appid string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("appid = ?", appid).Find(&results).Error

	return
}

// GetBatchFromAppid 批量唯一主键查找 公众号ID
func (obj *_WechatMediaMgr) GetBatchFromAppid(appids []string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("appid IN (?)", appids).Find(&results).Error

	return
}

// GetFromMd5 通过md5获取内容 文件md5
func (obj *_WechatMediaMgr) GetFromMd5(md5 string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("md5 = ?", md5).Find(&results).Error

	return
}

// GetBatchFromMd5 批量唯一主键查找 文件md5
func (obj *_WechatMediaMgr) GetBatchFromMd5(md5s []string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("md5 IN (?)", md5s).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 媒体类型
func (obj *_WechatMediaMgr) GetFromType(_type string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 媒体类型
func (obj *_WechatMediaMgr) GetBatchFromType(_types []string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromMediaID 通过media_id获取内容 永久素材MediaID
func (obj *_WechatMediaMgr) GetFromMediaID(mediaID string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("media_id = ?", mediaID).Find(&results).Error

	return
}

// GetBatchFromMediaID 批量唯一主键查找 永久素材MediaID
func (obj *_WechatMediaMgr) GetBatchFromMediaID(mediaIDs []string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("media_id IN (?)", mediaIDs).Find(&results).Error

	return
}

// GetFromLocalURL 通过local_url获取内容 本地文件链接
func (obj *_WechatMediaMgr) GetFromLocalURL(localURL string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("local_url = ?", localURL).Find(&results).Error

	return
}

// GetBatchFromLocalURL 批量唯一主键查找 本地文件链接
func (obj *_WechatMediaMgr) GetBatchFromLocalURL(localURLs []string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("local_url IN (?)", localURLs).Find(&results).Error

	return
}

// GetFromMediaURL 通过media_url获取内容 远程图片链接
func (obj *_WechatMediaMgr) GetFromMediaURL(mediaURL string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("media_url = ?", mediaURL).Find(&results).Error

	return
}

// GetBatchFromMediaURL 批量唯一主键查找 远程图片链接
func (obj *_WechatMediaMgr) GetBatchFromMediaURL(mediaURLs []string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("media_url IN (?)", mediaURLs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_WechatMediaMgr) GetFromCreateAt(createAt time.Time) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_WechatMediaMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WechatMediaMgr) FetchByPrimaryKey(id uint64) (result WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIndexWechatMediaAppid  获取多个内容
func (obj *_WechatMediaMgr) FetchIndexByIndexWechatMediaAppid(appid string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("appid = ?", appid).Find(&results).Error

	return
}

// FetchIndexByIndexWechatMediaMd5  获取多个内容
func (obj *_WechatMediaMgr) FetchIndexByIndexWechatMediaMd5(md5 string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("md5 = ?", md5).Find(&results).Error

	return
}

// FetchIndexByIndexWechatMediaType  获取多个内容
func (obj *_WechatMediaMgr) FetchIndexByIndexWechatMediaType(_type string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// FetchIndexByIndexWechatMediaMediaID  获取多个内容
func (obj *_WechatMediaMgr) FetchIndexByIndexWechatMediaMediaID(mediaID string) (results []*WechatMedia, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("media_id = ?", mediaID).Find(&results).Error

	return
}
