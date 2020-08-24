package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _SystemUserMgr struct {
	*_BaseMgr
}

// SystemUserMgr open func
func SystemUserMgr(db *gorm.DB) *_SystemUserMgr {
	if db == nil {
		panic(fmt.Errorf("SystemUserMgr need init by db"))
	}
	return &_SystemUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_user"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemUserMgr) GetTableName() string {
	return "system_user"
}

// Get 获取
func (obj *_SystemUserMgr) Get() (result SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemUserMgr) Gets() (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SystemUserMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取 用户账号
func (obj *_SystemUserMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 用户密码
func (obj *_SystemUserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithNickname nickname获取 用户昵称
func (obj *_SystemUserMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithHeadimg headimg获取 头像地址
func (obj *_SystemUserMgr) WithHeadimg(headimg string) Option {
	return optionFunc(func(o *options) { o.query["headimg"] = headimg })
}

// WithAuthorize authorize获取 权限授权
func (obj *_SystemUserMgr) WithAuthorize(authorize string) Option {
	return optionFunc(func(o *options) { o.query["authorize"] = authorize })
}

// WithContactQq contact_qq获取 联系QQ
func (obj *_SystemUserMgr) WithContactQq(contactQq string) Option {
	return optionFunc(func(o *options) { o.query["contact_qq"] = contactQq })
}

// WithContactMail contact_mail获取 联系邮箱
func (obj *_SystemUserMgr) WithContactMail(contactMail string) Option {
	return optionFunc(func(o *options) { o.query["contact_mail"] = contactMail })
}

// WithContactPhone contact_phone获取 联系手机
func (obj *_SystemUserMgr) WithContactPhone(contactPhone string) Option {
	return optionFunc(func(o *options) { o.query["contact_phone"] = contactPhone })
}

// WithLoginIP login_ip获取 登录地址
func (obj *_SystemUserMgr) WithLoginIP(loginIP string) Option {
	return optionFunc(func(o *options) { o.query["login_ip"] = loginIP })
}

// WithLoginAt login_at获取 登录时间
func (obj *_SystemUserMgr) WithLoginAt(loginAt string) Option {
	return optionFunc(func(o *options) { o.query["login_at"] = loginAt })
}

// WithLoginNum login_num获取 登录次数
func (obj *_SystemUserMgr) WithLoginNum(loginNum int64) Option {
	return optionFunc(func(o *options) { o.query["login_num"] = loginNum })
}

// WithDescribe describe获取 备注说明
func (obj *_SystemUserMgr) WithDescribe(describe string) Option {
	return optionFunc(func(o *options) { o.query["describe"] = describe })
}

// WithStatus status获取 状态(0禁用,1启用)
func (obj *_SystemUserMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSort sort获取 排序权重
func (obj *_SystemUserMgr) WithSort(sort int64) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithIsDeleted is_deleted获取 删除(1删除,0未删)
func (obj *_SystemUserMgr) WithIsDeleted(isDeleted bool) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithCreateAt create_at获取 创建时间
func (obj *_SystemUserMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_SystemUserMgr) GetByOption(opts ...Option) (result SystemUser, err error) {
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
func (obj *_SystemUserMgr) GetByOptions(opts ...Option) (results []*SystemUser, err error) {
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
func (obj *_SystemUserMgr) GetFromID(id int64) (result SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_SystemUserMgr) GetBatchFromID(ids []int64) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户账号
func (obj *_SystemUserMgr) GetFromUsername(username string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量唯一主键查找 用户账号
func (obj *_SystemUserMgr) GetBatchFromUsername(usernames []string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 用户密码
func (obj *_SystemUserMgr) GetFromPassword(password string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("password = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量唯一主键查找 用户密码
func (obj *_SystemUserMgr) GetBatchFromPassword(passwords []string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("password IN (?)", passwords).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 用户昵称
func (obj *_SystemUserMgr) GetFromNickname(nickname string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nickname = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量唯一主键查找 用户昵称
func (obj *_SystemUserMgr) GetBatchFromNickname(nicknames []string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("nickname IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromHeadimg 通过headimg获取内容 头像地址
func (obj *_SystemUserMgr) GetFromHeadimg(headimg string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("headimg = ?", headimg).Find(&results).Error

	return
}

// GetBatchFromHeadimg 批量唯一主键查找 头像地址
func (obj *_SystemUserMgr) GetBatchFromHeadimg(headimgs []string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("headimg IN (?)", headimgs).Find(&results).Error

	return
}

// GetFromAuthorize 通过authorize获取内容 权限授权
func (obj *_SystemUserMgr) GetFromAuthorize(authorize string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("authorize = ?", authorize).Find(&results).Error

	return
}

// GetBatchFromAuthorize 批量唯一主键查找 权限授权
func (obj *_SystemUserMgr) GetBatchFromAuthorize(authorizes []string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("authorize IN (?)", authorizes).Find(&results).Error

	return
}

// GetFromContactQq 通过contact_qq获取内容 联系QQ
func (obj *_SystemUserMgr) GetFromContactQq(contactQq string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("contact_qq = ?", contactQq).Find(&results).Error

	return
}

// GetBatchFromContactQq 批量唯一主键查找 联系QQ
func (obj *_SystemUserMgr) GetBatchFromContactQq(contactQqs []string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("contact_qq IN (?)", contactQqs).Find(&results).Error

	return
}

// GetFromContactMail 通过contact_mail获取内容 联系邮箱
func (obj *_SystemUserMgr) GetFromContactMail(contactMail string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("contact_mail = ?", contactMail).Find(&results).Error

	return
}

// GetBatchFromContactMail 批量唯一主键查找 联系邮箱
func (obj *_SystemUserMgr) GetBatchFromContactMail(contactMails []string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("contact_mail IN (?)", contactMails).Find(&results).Error

	return
}

// GetFromContactPhone 通过contact_phone获取内容 联系手机
func (obj *_SystemUserMgr) GetFromContactPhone(contactPhone string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("contact_phone = ?", contactPhone).Find(&results).Error

	return
}

// GetBatchFromContactPhone 批量唯一主键查找 联系手机
func (obj *_SystemUserMgr) GetBatchFromContactPhone(contactPhones []string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("contact_phone IN (?)", contactPhones).Find(&results).Error

	return
}

// GetFromLoginIP 通过login_ip获取内容 登录地址
func (obj *_SystemUserMgr) GetFromLoginIP(loginIP string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("login_ip = ?", loginIP).Find(&results).Error

	return
}

// GetBatchFromLoginIP 批量唯一主键查找 登录地址
func (obj *_SystemUserMgr) GetBatchFromLoginIP(loginIPs []string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("login_ip IN (?)", loginIPs).Find(&results).Error

	return
}

// GetFromLoginAt 通过login_at获取内容 登录时间
func (obj *_SystemUserMgr) GetFromLoginAt(loginAt string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("login_at = ?", loginAt).Find(&results).Error

	return
}

// GetBatchFromLoginAt 批量唯一主键查找 登录时间
func (obj *_SystemUserMgr) GetBatchFromLoginAt(loginAts []string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("login_at IN (?)", loginAts).Find(&results).Error

	return
}

// GetFromLoginNum 通过login_num获取内容 登录次数
func (obj *_SystemUserMgr) GetFromLoginNum(loginNum int64) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("login_num = ?", loginNum).Find(&results).Error

	return
}

// GetBatchFromLoginNum 批量唯一主键查找 登录次数
func (obj *_SystemUserMgr) GetBatchFromLoginNum(loginNums []int64) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("login_num IN (?)", loginNums).Find(&results).Error

	return
}

// GetFromDescribe 通过describe获取内容 备注说明
func (obj *_SystemUserMgr) GetFromDescribe(describe string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("describe = ?", describe).Find(&results).Error

	return
}

// GetBatchFromDescribe 批量唯一主键查找 备注说明
func (obj *_SystemUserMgr) GetBatchFromDescribe(describes []string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("describe IN (?)", describes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态(0禁用,1启用)
func (obj *_SystemUserMgr) GetFromStatus(status bool) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态(0禁用,1启用)
func (obj *_SystemUserMgr) GetBatchFromStatus(statuss []bool) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序权重
func (obj *_SystemUserMgr) GetFromSort(sort int64) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量唯一主键查找 排序权重
func (obj *_SystemUserMgr) GetBatchFromSort(sorts []int64) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容 删除(1删除,0未删)
func (obj *_SystemUserMgr) GetFromIsDeleted(isDeleted bool) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找 删除(1删除,0未删)
func (obj *_SystemUserMgr) GetBatchFromIsDeleted(isDeleteds []bool) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_SystemUserMgr) GetFromCreateAt(createAt time.Time) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_SystemUserMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SystemUserMgr) FetchByPrimaryKey(id int64) (result SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxSystemUserUsername  获取多个内容
func (obj *_SystemUserMgr) FetchIndexByIDxSystemUserUsername(username string) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username = ?", username).Find(&results).Error

	return
}

// FetchIndexByIDxSystemUserStatus  获取多个内容
func (obj *_SystemUserMgr) FetchIndexByIDxSystemUserStatus(status bool) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// FetchIndexByIDxSystemUserDeleted  获取多个内容
func (obj *_SystemUserMgr) FetchIndexByIDxSystemUserDeleted(isDeleted bool) (results []*SystemUser, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}
