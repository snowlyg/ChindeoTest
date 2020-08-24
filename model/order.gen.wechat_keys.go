package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _WechatKeysMgr struct {
	*_BaseMgr
}

// WechatKeysMgr open func
func WechatKeysMgr(db *gorm.DB) *_WechatKeysMgr {
	if db == nil {
		panic(fmt.Errorf("WechatKeysMgr need init by db"))
	}
	return &_WechatKeysMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wechat_keys"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WechatKeysMgr) GetTableName() string {
	return "wechat_keys"
}

// Get 获取
func (obj *_WechatKeysMgr) Get() (result WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WechatKeysMgr) Gets() (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WechatKeysMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAppid appid获取 公众号APPID
func (obj *_WechatKeysMgr) WithAppid(appid string) Option {
	return optionFunc(func(o *options) { o.query["appid"] = appid })
}

// WithType type获取 类型(text,image,news)
func (obj *_WechatKeysMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithKeys keys获取 关键字
func (obj *_WechatKeysMgr) WithKeys(keys string) Option {
	return optionFunc(func(o *options) { o.query["keys"] = keys })
}

// WithContent content获取 文本内容
func (obj *_WechatKeysMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithImageURL image_url获取 图片链接
func (obj *_WechatKeysMgr) WithImageURL(imageURL string) Option {
	return optionFunc(func(o *options) { o.query["image_url"] = imageURL })
}

// WithVoiceURL voice_url获取 语音链接
func (obj *_WechatKeysMgr) WithVoiceURL(voiceURL string) Option {
	return optionFunc(func(o *options) { o.query["voice_url"] = voiceURL })
}

// WithMusicTitle music_title获取 音乐标题
func (obj *_WechatKeysMgr) WithMusicTitle(musicTitle string) Option {
	return optionFunc(func(o *options) { o.query["music_title"] = musicTitle })
}

// WithMusicURL music_url获取 音乐链接
func (obj *_WechatKeysMgr) WithMusicURL(musicURL string) Option {
	return optionFunc(func(o *options) { o.query["music_url"] = musicURL })
}

// WithMusicImage music_image获取 缩略图片
func (obj *_WechatKeysMgr) WithMusicImage(musicImage string) Option {
	return optionFunc(func(o *options) { o.query["music_image"] = musicImage })
}

// WithMusicDesc music_desc获取 音乐描述
func (obj *_WechatKeysMgr) WithMusicDesc(musicDesc string) Option {
	return optionFunc(func(o *options) { o.query["music_desc"] = musicDesc })
}

// WithVideoTitle video_title获取 视频标题
func (obj *_WechatKeysMgr) WithVideoTitle(videoTitle string) Option {
	return optionFunc(func(o *options) { o.query["video_title"] = videoTitle })
}

// WithVideoURL video_url获取 视频URL
func (obj *_WechatKeysMgr) WithVideoURL(videoURL string) Option {
	return optionFunc(func(o *options) { o.query["video_url"] = videoURL })
}

// WithVideoDesc video_desc获取 视频描述
func (obj *_WechatKeysMgr) WithVideoDesc(videoDesc string) Option {
	return optionFunc(func(o *options) { o.query["video_desc"] = videoDesc })
}

// WithNewsID news_id获取 图文ID
func (obj *_WechatKeysMgr) WithNewsID(newsID uint64) Option {
	return optionFunc(func(o *options) { o.query["news_id"] = newsID })
}

// WithSort sort获取 排序字段
func (obj *_WechatKeysMgr) WithSort(sort uint64) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithStatus status获取 状态(0禁用,1启用)
func (obj *_WechatKeysMgr) WithStatus(status uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateBy create_by获取 创建人
func (obj *_WechatKeysMgr) WithCreateBy(createBy uint64) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithCreateAt create_at获取 创建时间
func (obj *_WechatKeysMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_WechatKeysMgr) GetByOption(opts ...Option) (result WechatKeys, err error) {
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
func (obj *_WechatKeysMgr) GetByOptions(opts ...Option) (results []*WechatKeys, err error) {
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
func (obj *_WechatKeysMgr) GetFromID(id int64) (result WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WechatKeysMgr) GetBatchFromID(ids []int64) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromAppid 通过appid获取内容 公众号APPID
func (obj *_WechatKeysMgr) GetFromAppid(appid string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("appid = ?", appid).Find(&results).Error

	return
}

// GetBatchFromAppid 批量唯一主键查找 公众号APPID
func (obj *_WechatKeysMgr) GetBatchFromAppid(appids []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("appid IN (?)", appids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型(text,image,news)
func (obj *_WechatKeysMgr) GetFromType(_type string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 类型(text,image,news)
func (obj *_WechatKeysMgr) GetBatchFromType(_types []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromKeys 通过keys获取内容 关键字
func (obj *_WechatKeysMgr) GetFromKeys(keys string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("keys = ?", keys).Find(&results).Error

	return
}

// GetBatchFromKeys 批量唯一主键查找 关键字
func (obj *_WechatKeysMgr) GetBatchFromKeys(keyss []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("keys IN (?)", keyss).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 文本内容
func (obj *_WechatKeysMgr) GetFromContent(content string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找 文本内容
func (obj *_WechatKeysMgr) GetBatchFromContent(contents []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

// GetFromImageURL 通过image_url获取内容 图片链接
func (obj *_WechatKeysMgr) GetFromImageURL(imageURL string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("image_url = ?", imageURL).Find(&results).Error

	return
}

// GetBatchFromImageURL 批量唯一主键查找 图片链接
func (obj *_WechatKeysMgr) GetBatchFromImageURL(imageURLs []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("image_url IN (?)", imageURLs).Find(&results).Error

	return
}

// GetFromVoiceURL 通过voice_url获取内容 语音链接
func (obj *_WechatKeysMgr) GetFromVoiceURL(voiceURL string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("voice_url = ?", voiceURL).Find(&results).Error

	return
}

// GetBatchFromVoiceURL 批量唯一主键查找 语音链接
func (obj *_WechatKeysMgr) GetBatchFromVoiceURL(voiceURLs []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("voice_url IN (?)", voiceURLs).Find(&results).Error

	return
}

// GetFromMusicTitle 通过music_title获取内容 音乐标题
func (obj *_WechatKeysMgr) GetFromMusicTitle(musicTitle string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("music_title = ?", musicTitle).Find(&results).Error

	return
}

// GetBatchFromMusicTitle 批量唯一主键查找 音乐标题
func (obj *_WechatKeysMgr) GetBatchFromMusicTitle(musicTitles []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("music_title IN (?)", musicTitles).Find(&results).Error

	return
}

// GetFromMusicURL 通过music_url获取内容 音乐链接
func (obj *_WechatKeysMgr) GetFromMusicURL(musicURL string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("music_url = ?", musicURL).Find(&results).Error

	return
}

// GetBatchFromMusicURL 批量唯一主键查找 音乐链接
func (obj *_WechatKeysMgr) GetBatchFromMusicURL(musicURLs []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("music_url IN (?)", musicURLs).Find(&results).Error

	return
}

// GetFromMusicImage 通过music_image获取内容 缩略图片
func (obj *_WechatKeysMgr) GetFromMusicImage(musicImage string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("music_image = ?", musicImage).Find(&results).Error

	return
}

// GetBatchFromMusicImage 批量唯一主键查找 缩略图片
func (obj *_WechatKeysMgr) GetBatchFromMusicImage(musicImages []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("music_image IN (?)", musicImages).Find(&results).Error

	return
}

// GetFromMusicDesc 通过music_desc获取内容 音乐描述
func (obj *_WechatKeysMgr) GetFromMusicDesc(musicDesc string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("music_desc = ?", musicDesc).Find(&results).Error

	return
}

// GetBatchFromMusicDesc 批量唯一主键查找 音乐描述
func (obj *_WechatKeysMgr) GetBatchFromMusicDesc(musicDescs []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("music_desc IN (?)", musicDescs).Find(&results).Error

	return
}

// GetFromVideoTitle 通过video_title获取内容 视频标题
func (obj *_WechatKeysMgr) GetFromVideoTitle(videoTitle string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("video_title = ?", videoTitle).Find(&results).Error

	return
}

// GetBatchFromVideoTitle 批量唯一主键查找 视频标题
func (obj *_WechatKeysMgr) GetBatchFromVideoTitle(videoTitles []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("video_title IN (?)", videoTitles).Find(&results).Error

	return
}

// GetFromVideoURL 通过video_url获取内容 视频URL
func (obj *_WechatKeysMgr) GetFromVideoURL(videoURL string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("video_url = ?", videoURL).Find(&results).Error

	return
}

// GetBatchFromVideoURL 批量唯一主键查找 视频URL
func (obj *_WechatKeysMgr) GetBatchFromVideoURL(videoURLs []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("video_url IN (?)", videoURLs).Find(&results).Error

	return
}

// GetFromVideoDesc 通过video_desc获取内容 视频描述
func (obj *_WechatKeysMgr) GetFromVideoDesc(videoDesc string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("video_desc = ?", videoDesc).Find(&results).Error

	return
}

// GetBatchFromVideoDesc 批量唯一主键查找 视频描述
func (obj *_WechatKeysMgr) GetBatchFromVideoDesc(videoDescs []string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("video_desc IN (?)", videoDescs).Find(&results).Error

	return
}

// GetFromNewsID 通过news_id获取内容 图文ID
func (obj *_WechatKeysMgr) GetFromNewsID(newsID uint64) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("news_id = ?", newsID).Find(&results).Error

	return
}

// GetBatchFromNewsID 批量唯一主键查找 图文ID
func (obj *_WechatKeysMgr) GetBatchFromNewsID(newsIDs []uint64) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("news_id IN (?)", newsIDs).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序字段
func (obj *_WechatKeysMgr) GetFromSort(sort uint64) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序字段
func (obj *_WechatKeysMgr) GetBatchFromSort(sorts []uint64) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态(0禁用,1启用)
func (obj *_WechatKeysMgr) GetFromStatus(status uint8) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态(0禁用,1启用)
func (obj *_WechatKeysMgr) GetBatchFromStatus(statuss []uint8) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建人
func (obj *_WechatKeysMgr) GetFromCreateBy(createBy uint64) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_by = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量唯一主键查找 创建人
func (obj *_WechatKeysMgr) GetBatchFromCreateBy(createBys []uint64) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_by IN (?)", createBys).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_WechatKeysMgr) GetFromCreateAt(createAt time.Time) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_WechatKeysMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WechatKeysMgr) FetchByPrimaryKey(id int64) (result WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIndexWechatKeysAppid  获取多个内容
func (obj *_WechatKeysMgr) FetchIndexByIndexWechatKeysAppid(appid string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("appid = ?", appid).Find(&results).Error

	return
}

// FetchIndexByIndexWechatKeysType  获取多个内容
func (obj *_WechatKeysMgr) FetchIndexByIndexWechatKeysType(_type string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// FetchIndexByIndexWechatKeysKeys  获取多个内容
func (obj *_WechatKeysMgr) FetchIndexByIndexWechatKeysKeys(keys string) (results []*WechatKeys, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("keys = ?", keys).Find(&results).Error

	return
}
