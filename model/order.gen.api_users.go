package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIUsersMgr struct {
	*_BaseMgr
}

// APIUsersMgr open func
func APIUsersMgr(db *gorm.DB) *_APIUsersMgr {
	if db == nil {
		panic(fmt.Errorf("APIUsersMgr need init by db"))
	}
	return &_APIUsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_users"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIUsersMgr) GetTableName() string {
	return "api_users"
}

// Get 获取
func (obj *_APIUsersMgr) Get() (result APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIUsersMgr) Gets() (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIUsersMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取 用户名
func (obj *_APIUsersMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithNickname nickname获取 名称
func (obj *_APIUsersMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithPhone phone获取 手机
func (obj *_APIUsersMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithEmail email获取 邮箱
func (obj *_APIUsersMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithSex sex获取 性别
func (obj *_APIUsersMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithPassword password获取 用户密码
func (obj *_APIUsersMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithStatus status获取 账号状态
func (obj *_APIUsersMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithAvatarURL avatar_url获取 用户头像
func (obj *_APIUsersMgr) WithAvatarURL(avatarURL string) Option {
	return optionFunc(func(o *options) { o.query["avatar_url"] = avatarURL })
}

// WithOpenID open_id获取 open_id
func (obj *_APIUsersMgr) WithOpenID(openID string) Option {
	return optionFunc(func(o *options) { o.query["open_id"] = openID })
}

// WithUnionID union_id获取 union_id
func (obj *_APIUsersMgr) WithUnionID(unionID string) Option {
	return optionFunc(func(o *options) { o.query["union_id"] = unionID })
}

// WithCountry country获取 国家
func (obj *_APIUsersMgr) WithCountry(country string) Option {
	return optionFunc(func(o *options) { o.query["country"] = country })
}

// WithProvince province获取 省份
func (obj *_APIUsersMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithCity city获取 城镇
func (obj *_APIUsersMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithMac mac获取 mac 地址
func (obj *_APIUsersMgr) WithMac(mac string) Option {
	return optionFunc(func(o *options) { o.query["mac"] = mac })
}

// WithCreateAt create_at获取
func (obj *_APIUsersMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIUsersMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIDCardNo id_card_no获取 身份证
func (obj *_APIUsersMgr) WithIDCardNo(idCardNo string) Option {
	return optionFunc(func(o *options) { o.query["id_card_no"] = idCardNo })
}

// WithIsAuth is_auth获取 是否实名认证
func (obj *_APIUsersMgr) WithIsAuth(isAuth bool) Option {
	return optionFunc(func(o *options) { o.query["is_auth"] = isAuth })
}

// WithRealname realname获取 真实姓名
func (obj *_APIUsersMgr) WithRealname(realname string) Option {
	return optionFunc(func(o *options) { o.query["realname"] = realname })
}

// GetByOption 功能选项模式获取
func (obj *_APIUsersMgr) GetByOption(opts ...Option) (result APIUsers, err error) {
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
func (obj *_APIUsersMgr) GetByOptions(opts ...Option) (results []*APIUsers, err error) {
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
func (obj *_APIUsersMgr) GetFromID(id int) (result APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIUsersMgr) GetBatchFromID(ids []int) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户名
func (obj *_APIUsersMgr) GetFromUsername(username string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量唯一主键查找 用户名
func (obj *_APIUsersMgr) GetBatchFromUsername(usernames []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username IN (?)", usernames).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 名称
func (obj *_APIUsersMgr) GetFromNickname(nickname string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nickname = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量唯一主键查找 名称
func (obj *_APIUsersMgr) GetBatchFromNickname(nicknames []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nickname IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机
func (obj *_APIUsersMgr) GetFromPhone(phone string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 手机
func (obj *_APIUsersMgr) GetBatchFromPhone(phones []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱
func (obj *_APIUsersMgr) GetFromEmail(email string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找 邮箱
func (obj *_APIUsersMgr) GetBatchFromEmail(emails []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 性别
func (obj *_APIUsersMgr) GetFromSex(sex int) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找 性别
func (obj *_APIUsersMgr) GetBatchFromSex(sexs []int) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", sexs).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 用户密码
func (obj *_APIUsersMgr) GetFromPassword(password string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("password = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量唯一主键查找 用户密码
func (obj *_APIUsersMgr) GetBatchFromPassword(passwords []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("password IN (?)", passwords).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 账号状态
func (obj *_APIUsersMgr) GetFromStatus(status bool) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 账号状态
func (obj *_APIUsersMgr) GetBatchFromStatus(statuss []bool) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromAvatarURL 通过avatar_url获取内容 用户头像
func (obj *_APIUsersMgr) GetFromAvatarURL(avatarURL string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("avatar_url = ?", avatarURL).Find(&results).Error

	return
}

// GetBatchFromAvatarURL 批量唯一主键查找 用户头像
func (obj *_APIUsersMgr) GetBatchFromAvatarURL(avatarURLs []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("avatar_url IN (?)", avatarURLs).Find(&results).Error

	return
}

// GetFromOpenID 通过open_id获取内容 open_id
func (obj *_APIUsersMgr) GetFromOpenID(openID string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("open_id = ?", openID).Find(&results).Error

	return
}

// GetBatchFromOpenID 批量唯一主键查找 open_id
func (obj *_APIUsersMgr) GetBatchFromOpenID(openIDs []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("open_id IN (?)", openIDs).Find(&results).Error

	return
}

// GetFromUnionID 通过union_id获取内容 union_id
func (obj *_APIUsersMgr) GetFromUnionID(unionID string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("union_id = ?", unionID).Find(&results).Error

	return
}

// GetBatchFromUnionID 批量唯一主键查找 union_id
func (obj *_APIUsersMgr) GetBatchFromUnionID(unionIDs []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("union_id IN (?)", unionIDs).Find(&results).Error

	return
}

// GetFromCountry 通过country获取内容 国家
func (obj *_APIUsersMgr) GetFromCountry(country string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("country = ?", country).Find(&results).Error

	return
}

// GetBatchFromCountry 批量唯一主键查找 国家
func (obj *_APIUsersMgr) GetBatchFromCountry(countrys []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("country IN (?)", countrys).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 省份
func (obj *_APIUsersMgr) GetFromProvince(province string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("province = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量唯一主键查找 省份
func (obj *_APIUsersMgr) GetBatchFromProvince(provinces []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("province IN (?)", provinces).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 城镇
func (obj *_APIUsersMgr) GetFromCity(city string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("city = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量唯一主键查找 城镇
func (obj *_APIUsersMgr) GetBatchFromCity(citys []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("city IN (?)", citys).Find(&results).Error

	return
}

// GetFromMac 通过mac获取内容 mac 地址
func (obj *_APIUsersMgr) GetFromMac(mac string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mac = ?", mac).Find(&results).Error

	return
}

// GetBatchFromMac 批量唯一主键查找 mac 地址
func (obj *_APIUsersMgr) GetBatchFromMac(macs []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("mac IN (?)", macs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIUsersMgr) GetFromCreateAt(createAt time.Time) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIUsersMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIUsersMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIUsersMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIDCardNo 通过id_card_no获取内容 身份证
func (obj *_APIUsersMgr) GetFromIDCardNo(idCardNo string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id_card_no = ?", idCardNo).Find(&results).Error

	return
}

// GetBatchFromIDCardNo 批量唯一主键查找 身份证
func (obj *_APIUsersMgr) GetBatchFromIDCardNo(idCardNos []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id_card_no IN (?)", idCardNos).Find(&results).Error

	return
}

// GetFromIsAuth 通过is_auth获取内容 是否实名认证
func (obj *_APIUsersMgr) GetFromIsAuth(isAuth bool) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_auth = ?", isAuth).Find(&results).Error

	return
}

// GetBatchFromIsAuth 批量唯一主键查找 是否实名认证
func (obj *_APIUsersMgr) GetBatchFromIsAuth(isAuths []bool) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_auth IN (?)", isAuths).Find(&results).Error

	return
}

// GetFromRealname 通过realname获取内容 真实姓名
func (obj *_APIUsersMgr) GetFromRealname(realname string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("realname = ?", realname).Find(&results).Error

	return
}

// GetBatchFromRealname 批量唯一主键查找 真实姓名
func (obj *_APIUsersMgr) GetBatchFromRealname(realnames []string) (results []*APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("realname IN (?)", realnames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIUsersMgr) FetchByPrimaryKey(id int) (result APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByOpenID primay or index 获取唯一内容
func (obj *_APIUsersMgr) FetchUniqueIndexByOpenID(username string, phone string, email string, openID string, unionID string) (result APIUsers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username = ? AND phone = ? AND email = ? AND open_id = ? AND union_id = ?", username, phone, email, openID, unionID).Find(&result).Error

	return
}
