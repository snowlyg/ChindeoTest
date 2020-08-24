package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _WechatFansMgr struct {
	*_BaseMgr
}

// WechatFansMgr open func
func WechatFansMgr(db *gorm.DB) *_WechatFansMgr {
	if db == nil {
		panic(fmt.Errorf("WechatFansMgr need init by db"))
	}
	return &_WechatFansMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wechat_fans"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WechatFansMgr) GetTableName() string {
	return "wechat_fans"
}

// Get 获取
func (obj *_WechatFansMgr) Get() (result WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WechatFansMgr) Gets() (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WechatFansMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAppid appid获取 公众号APPID
func (obj *_WechatFansMgr) WithAppid(appid string) Option {
	return optionFunc(func(o *options) { o.query["appid"] = appid })
}

// WithUnionid unionid获取 粉丝unionid
func (obj *_WechatFansMgr) WithUnionid(unionid string) Option {
	return optionFunc(func(o *options) { o.query["unionid"] = unionid })
}

// WithOpenid openid获取 粉丝openid
func (obj *_WechatFansMgr) WithOpenid(openid string) Option {
	return optionFunc(func(o *options) { o.query["openid"] = openid })
}

// WithTagidList tagid_list获取 粉丝标签id
func (obj *_WechatFansMgr) WithTagidList(tagidList string) Option {
	return optionFunc(func(o *options) { o.query["tagid_list"] = tagidList })
}

// WithIsBlack is_black获取 是否为黑名单状态
func (obj *_WechatFansMgr) WithIsBlack(isBlack uint8) Option {
	return optionFunc(func(o *options) { o.query["is_black"] = isBlack })
}

// WithSubscribe subscribe获取 关注状态(0未关注,1已关注)
func (obj *_WechatFansMgr) WithSubscribe(subscribe uint8) Option {
	return optionFunc(func(o *options) { o.query["subscribe"] = subscribe })
}

// WithNickname nickname获取 用户昵称
func (obj *_WechatFansMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithSex sex获取 用户性别(1男性,2女性,0未知)
func (obj *_WechatFansMgr) WithSex(sex uint8) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithCountry country获取 用户所在国家
func (obj *_WechatFansMgr) WithCountry(country string) Option {
	return optionFunc(func(o *options) { o.query["country"] = country })
}

// WithProvince province获取 用户所在省份
func (obj *_WechatFansMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithCity city获取 用户所在城市
func (obj *_WechatFansMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithLanguage language获取 用户的语言(zh_CN)
func (obj *_WechatFansMgr) WithLanguage(language string) Option {
	return optionFunc(func(o *options) { o.query["language"] = language })
}

// WithHeadimgurl headimgurl获取 用户头像
func (obj *_WechatFansMgr) WithHeadimgurl(headimgurl string) Option {
	return optionFunc(func(o *options) { o.query["headimgurl"] = headimgurl })
}

// WithSubscribeTime subscribe_time获取 关注时间
func (obj *_WechatFansMgr) WithSubscribeTime(subscribeTime uint64) Option {
	return optionFunc(func(o *options) { o.query["subscribe_time"] = subscribeTime })
}

// WithSubscribeAt subscribe_at获取 关注时间
func (obj *_WechatFansMgr) WithSubscribeAt(subscribeAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["subscribe_at"] = subscribeAt })
}

// WithRemark remark获取 备注
func (obj *_WechatFansMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithSubscribeScene subscribe_scene获取 扫码关注场景
func (obj *_WechatFansMgr) WithSubscribeScene(subscribeScene string) Option {
	return optionFunc(func(o *options) { o.query["subscribe_scene"] = subscribeScene })
}

// WithQrScene qr_scene获取 二维码场景值
func (obj *_WechatFansMgr) WithQrScene(qrScene string) Option {
	return optionFunc(func(o *options) { o.query["qr_scene"] = qrScene })
}

// WithQrSceneStr qr_scene_str获取 二维码场景内容
func (obj *_WechatFansMgr) WithQrSceneStr(qrSceneStr string) Option {
	return optionFunc(func(o *options) { o.query["qr_scene_str"] = qrSceneStr })
}

// WithCreateAt create_at获取 创建时间
func (obj *_WechatFansMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_WechatFansMgr) GetByOption(opts ...Option) (result WechatFans, err error) {
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
func (obj *_WechatFansMgr) GetByOptions(opts ...Option) (results []*WechatFans, err error) {
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
func (obj *_WechatFansMgr) GetFromID(id uint64) (result WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_WechatFansMgr) GetBatchFromID(ids []uint64) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromAppid 通过appid获取内容 公众号APPID
func (obj *_WechatFansMgr) GetFromAppid(appid string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("appid = ?", appid).Find(&results).Error

	return
}

// GetBatchFromAppid 批量唯一主键查找 公众号APPID
func (obj *_WechatFansMgr) GetBatchFromAppid(appids []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("appid IN (?)", appids).Find(&results).Error

	return
}

// GetFromUnionid 通过unionid获取内容 粉丝unionid
func (obj *_WechatFansMgr) GetFromUnionid(unionid string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unionid = ?", unionid).Find(&results).Error

	return
}

// GetBatchFromUnionid 批量唯一主键查找 粉丝unionid
func (obj *_WechatFansMgr) GetBatchFromUnionid(unionids []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unionid IN (?)", unionids).Find(&results).Error

	return
}

// GetFromOpenid 通过openid获取内容 粉丝openid
func (obj *_WechatFansMgr) GetFromOpenid(openid string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("openid = ?", openid).Find(&results).Error

	return
}

// GetBatchFromOpenid 批量唯一主键查找 粉丝openid
func (obj *_WechatFansMgr) GetBatchFromOpenid(openids []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("openid IN (?)", openids).Find(&results).Error

	return
}

// GetFromTagidList 通过tagid_list获取内容 粉丝标签id
func (obj *_WechatFansMgr) GetFromTagidList(tagidList string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("tagid_list = ?", tagidList).Find(&results).Error

	return
}

// GetBatchFromTagidList 批量唯一主键查找 粉丝标签id
func (obj *_WechatFansMgr) GetBatchFromTagidList(tagidLists []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("tagid_list IN (?)", tagidLists).Find(&results).Error

	return
}

// GetFromIsBlack 通过is_black获取内容 是否为黑名单状态
func (obj *_WechatFansMgr) GetFromIsBlack(isBlack uint8) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_black = ?", isBlack).Find(&results).Error

	return
}

// GetBatchFromIsBlack 批量唯一主键查找 是否为黑名单状态
func (obj *_WechatFansMgr) GetBatchFromIsBlack(isBlacks []uint8) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_black IN (?)", isBlacks).Find(&results).Error

	return
}

// GetFromSubscribe 通过subscribe获取内容 关注状态(0未关注,1已关注)
func (obj *_WechatFansMgr) GetFromSubscribe(subscribe uint8) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("subscribe = ?", subscribe).Find(&results).Error

	return
}

// GetBatchFromSubscribe 批量唯一主键查找 关注状态(0未关注,1已关注)
func (obj *_WechatFansMgr) GetBatchFromSubscribe(subscribes []uint8) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("subscribe IN (?)", subscribes).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 用户昵称
func (obj *_WechatFansMgr) GetFromNickname(nickname string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nickname = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量唯一主键查找 用户昵称
func (obj *_WechatFansMgr) GetBatchFromNickname(nicknames []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nickname IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 用户性别(1男性,2女性,0未知)
func (obj *_WechatFansMgr) GetFromSex(sex uint8) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找 用户性别(1男性,2女性,0未知)
func (obj *_WechatFansMgr) GetBatchFromSex(sexs []uint8) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", sexs).Find(&results).Error

	return
}

// GetFromCountry 通过country获取内容 用户所在国家
func (obj *_WechatFansMgr) GetFromCountry(country string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("country = ?", country).Find(&results).Error

	return
}

// GetBatchFromCountry 批量唯一主键查找 用户所在国家
func (obj *_WechatFansMgr) GetBatchFromCountry(countrys []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("country IN (?)", countrys).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 用户所在省份
func (obj *_WechatFansMgr) GetFromProvince(province string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("province = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量唯一主键查找 用户所在省份
func (obj *_WechatFansMgr) GetBatchFromProvince(provinces []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("province IN (?)", provinces).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 用户所在城市
func (obj *_WechatFansMgr) GetFromCity(city string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("city = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量唯一主键查找 用户所在城市
func (obj *_WechatFansMgr) GetBatchFromCity(citys []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("city IN (?)", citys).Find(&results).Error

	return
}

// GetFromLanguage 通过language获取内容 用户的语言(zh_CN)
func (obj *_WechatFansMgr) GetFromLanguage(language string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("language = ?", language).Find(&results).Error

	return
}

// GetBatchFromLanguage 批量唯一主键查找 用户的语言(zh_CN)
func (obj *_WechatFansMgr) GetBatchFromLanguage(languages []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("language IN (?)", languages).Find(&results).Error

	return
}

// GetFromHeadimgurl 通过headimgurl获取内容 用户头像
func (obj *_WechatFansMgr) GetFromHeadimgurl(headimgurl string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("headimgurl = ?", headimgurl).Find(&results).Error

	return
}

// GetBatchFromHeadimgurl 批量唯一主键查找 用户头像
func (obj *_WechatFansMgr) GetBatchFromHeadimgurl(headimgurls []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("headimgurl IN (?)", headimgurls).Find(&results).Error

	return
}

// GetFromSubscribeTime 通过subscribe_time获取内容 关注时间
func (obj *_WechatFansMgr) GetFromSubscribeTime(subscribeTime uint64) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("subscribe_time = ?", subscribeTime).Find(&results).Error

	return
}

// GetBatchFromSubscribeTime 批量唯一主键查找 关注时间
func (obj *_WechatFansMgr) GetBatchFromSubscribeTime(subscribeTimes []uint64) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("subscribe_time IN (?)", subscribeTimes).Find(&results).Error

	return
}

// GetFromSubscribeAt 通过subscribe_at获取内容 关注时间
func (obj *_WechatFansMgr) GetFromSubscribeAt(subscribeAt time.Time) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("subscribe_at = ?", subscribeAt).Find(&results).Error

	return
}

// GetBatchFromSubscribeAt 批量唯一主键查找 关注时间
func (obj *_WechatFansMgr) GetBatchFromSubscribeAt(subscribeAts []time.Time) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("subscribe_at IN (?)", subscribeAts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WechatFansMgr) GetFromRemark(remark string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量唯一主键查找 备注
func (obj *_WechatFansMgr) GetBatchFromRemark(remarks []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromSubscribeScene 通过subscribe_scene获取内容 扫码关注场景
func (obj *_WechatFansMgr) GetFromSubscribeScene(subscribeScene string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("subscribe_scene = ?", subscribeScene).Find(&results).Error

	return
}

// GetBatchFromSubscribeScene 批量唯一主键查找 扫码关注场景
func (obj *_WechatFansMgr) GetBatchFromSubscribeScene(subscribeScenes []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("subscribe_scene IN (?)", subscribeScenes).Find(&results).Error

	return
}

// GetFromQrScene 通过qr_scene获取内容 二维码场景值
func (obj *_WechatFansMgr) GetFromQrScene(qrScene string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("qr_scene = ?", qrScene).Find(&results).Error

	return
}

// GetBatchFromQrScene 批量唯一主键查找 二维码场景值
func (obj *_WechatFansMgr) GetBatchFromQrScene(qrScenes []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("qr_scene IN (?)", qrScenes).Find(&results).Error

	return
}

// GetFromQrSceneStr 通过qr_scene_str获取内容 二维码场景内容
func (obj *_WechatFansMgr) GetFromQrSceneStr(qrSceneStr string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("qr_scene_str = ?", qrSceneStr).Find(&results).Error

	return
}

// GetBatchFromQrSceneStr 批量唯一主键查找 二维码场景内容
func (obj *_WechatFansMgr) GetBatchFromQrSceneStr(qrSceneStrs []string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("qr_scene_str IN (?)", qrSceneStrs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_WechatFansMgr) GetFromCreateAt(createAt time.Time) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_WechatFansMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WechatFansMgr) FetchByPrimaryKey(id uint64) (result WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIndexWechatFansUnionid  获取多个内容
func (obj *_WechatFansMgr) FetchIndexByIndexWechatFansUnionid(unionid string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("unionid = ?", unionid).Find(&results).Error

	return
}

// FetchIndexByIndexWechatFansOpenid  获取多个内容
func (obj *_WechatFansMgr) FetchIndexByIndexWechatFansOpenid(openid string) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("openid = ?", openid).Find(&results).Error

	return
}

// FetchIndexByIndexWechatFansIsBack  获取多个内容
func (obj *_WechatFansMgr) FetchIndexByIndexWechatFansIsBack(isBlack uint8) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_black = ?", isBlack).Find(&results).Error

	return
}

// FetchIndexByIndexWechatFansSubscribe  获取多个内容
func (obj *_WechatFansMgr) FetchIndexByIndexWechatFansSubscribe(subscribe uint8) (results []*WechatFans, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("subscribe = ?", subscribe).Find(&results).Error

	return
}
