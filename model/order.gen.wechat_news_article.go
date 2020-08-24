package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _WechatNewsArticleMgr struct {
	*_BaseMgr
}

// WechatNewsArticleMgr open func
func WechatNewsArticleMgr(db *gorm.DB) *_WechatNewsArticleMgr {
	if db == nil {
		panic(fmt.Errorf("WechatNewsArticleMgr need init by db"))
	}
	return &_WechatNewsArticleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wechat_news_article"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WechatNewsArticleMgr) GetTableName() string {
	return "wechat_news_article"
}

// Get 获取
func (obj *_WechatNewsArticleMgr) Get() (result WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WechatNewsArticleMgr) Gets() (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WechatNewsArticleMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取 素材标题
func (obj *_WechatNewsArticleMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithLocalURL local_url获取 永久素材显示URL
func (obj *_WechatNewsArticleMgr) WithLocalURL(localURL string) Option {
	return optionFunc(func(o *options) { o.query["local_url"] = localURL })
}

// WithShowCoverPic show_cover_pic获取 显示封面(0不显示,1显示)
func (obj *_WechatNewsArticleMgr) WithShowCoverPic(showCoverPic uint8) Option {
	return optionFunc(func(o *options) { o.query["show_cover_pic"] = showCoverPic })
}

// WithAuthor author获取 文章作者
func (obj *_WechatNewsArticleMgr) WithAuthor(author string) Option {
	return optionFunc(func(o *options) { o.query["author"] = author })
}

// WithDigest digest获取 摘要内容
func (obj *_WechatNewsArticleMgr) WithDigest(digest string) Option {
	return optionFunc(func(o *options) { o.query["digest"] = digest })
}

// WithContent content获取 图文内容
func (obj *_WechatNewsArticleMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithContentSourceURL content_source_url获取 原文地址
func (obj *_WechatNewsArticleMgr) WithContentSourceURL(contentSourceURL string) Option {
	return optionFunc(func(o *options) { o.query["content_source_url"] = contentSourceURL })
}

// WithReadNum read_num获取 阅读数量
func (obj *_WechatNewsArticleMgr) WithReadNum(readNum uint64) Option {
	return optionFunc(func(o *options) { o.query["read_num"] = readNum })
}

// WithCreateAt create_at获取 创建时间
func (obj *_WechatNewsArticleMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_WechatNewsArticleMgr) GetByOption(opts ...Option) (result WechatNewsArticle, err error) {
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
func (obj *_WechatNewsArticleMgr) GetByOptions(opts ...Option) (results []*WechatNewsArticle, err error) {
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
func (obj *_WechatNewsArticleMgr) GetFromID(id uint64) (result WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WechatNewsArticleMgr) GetBatchFromID(ids []uint64) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 素材标题
func (obj *_WechatNewsArticleMgr) GetFromTitle(title string) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找 素材标题
func (obj *_WechatNewsArticleMgr) GetBatchFromTitle(titles []string) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromLocalURL 通过local_url获取内容 永久素材显示URL
func (obj *_WechatNewsArticleMgr) GetFromLocalURL(localURL string) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("local_url = ?", localURL).Find(&results).Error

	return
}

// GetBatchFromLocalURL 批量唯一主键查找 永久素材显示URL
func (obj *_WechatNewsArticleMgr) GetBatchFromLocalURL(localURLs []string) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("local_url IN (?)", localURLs).Find(&results).Error

	return
}

// GetFromShowCoverPic 通过show_cover_pic获取内容 显示封面(0不显示,1显示)
func (obj *_WechatNewsArticleMgr) GetFromShowCoverPic(showCoverPic uint8) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("show_cover_pic = ?", showCoverPic).Find(&results).Error

	return
}

// GetBatchFromShowCoverPic 批量唯一主键查找 显示封面(0不显示,1显示)
func (obj *_WechatNewsArticleMgr) GetBatchFromShowCoverPic(showCoverPics []uint8) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("show_cover_pic IN (?)", showCoverPics).Find(&results).Error

	return
}

// GetFromAuthor 通过author获取内容 文章作者
func (obj *_WechatNewsArticleMgr) GetFromAuthor(author string) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("author = ?", author).Find(&results).Error

	return
}

// GetBatchFromAuthor 批量唯一主键查找 文章作者
func (obj *_WechatNewsArticleMgr) GetBatchFromAuthor(authors []string) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("author IN (?)", authors).Find(&results).Error

	return
}

// GetFromDigest 通过digest获取内容 摘要内容
func (obj *_WechatNewsArticleMgr) GetFromDigest(digest string) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("digest = ?", digest).Find(&results).Error

	return
}

// GetBatchFromDigest 批量唯一主键查找 摘要内容
func (obj *_WechatNewsArticleMgr) GetBatchFromDigest(digests []string) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("digest IN (?)", digests).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 图文内容
func (obj *_WechatNewsArticleMgr) GetFromContent(content string) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找 图文内容
func (obj *_WechatNewsArticleMgr) GetBatchFromContent(contents []string) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

// GetFromContentSourceURL 通过content_source_url获取内容 原文地址
func (obj *_WechatNewsArticleMgr) GetFromContentSourceURL(contentSourceURL string) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content_source_url = ?", contentSourceURL).Find(&results).Error

	return
}

// GetBatchFromContentSourceURL 批量唯一主键查找 原文地址
func (obj *_WechatNewsArticleMgr) GetBatchFromContentSourceURL(contentSourceURLs []string) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content_source_url IN (?)", contentSourceURLs).Find(&results).Error

	return
}

// GetFromReadNum 通过read_num获取内容 阅读数量
func (obj *_WechatNewsArticleMgr) GetFromReadNum(readNum uint64) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("read_num = ?", readNum).Find(&results).Error

	return
}

// GetBatchFromReadNum 批量唯一主键查找 阅读数量
func (obj *_WechatNewsArticleMgr) GetBatchFromReadNum(readNums []uint64) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("read_num IN (?)", readNums).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_WechatNewsArticleMgr) GetFromCreateAt(createAt time.Time) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_WechatNewsArticleMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WechatNewsArticleMgr) FetchByPrimaryKey(id uint64) (result WechatNewsArticle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
