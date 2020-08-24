package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CarersMgr struct {
	*_BaseMgr
}

// CarersMgr open func
func CarersMgr(db *gorm.DB) *_CarersMgr {
	if db == nil {
		panic(fmt.Errorf("CarersMgr need init by db"))
	}
	return &_CarersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("carers"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CarersMgr) GetTableName() string {
	return "carers"
}

// Get 获取
func (obj *_CarersMgr) Get() (result Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CarersMgr) Gets() (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CarersMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 姓名
func (obj *_CarersMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAge age获取 年龄
func (obj *_CarersMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// WithWorkExp work_exp获取 工作年限
func (obj *_CarersMgr) WithWorkExp(workExp int) Option {
	return optionFunc(func(o *options) { o.query["work_exp"] = workExp })
}

// WithPrice price获取 价格
func (obj *_CarersMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithSex sex获取 性别:1:男,0:女
func (obj *_CarersMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithPhone phone获取 手机
func (obj *_CarersMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithDesc desc获取 简介
func (obj *_CarersMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithAvatar avatar获取 头像
func (obj *_CarersMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithAmount amount获取 服务数量
func (obj *_CarersMgr) WithAmount(amount int) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithCreateAt create_at获取
func (obj *_CarersMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CarersMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_CarersMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithApplicationID application_id获取 application_id
func (obj *_CarersMgr) WithApplicationID(applicationID int) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithStatus status获取 状态：上下架
func (obj *_CarersMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTimeType time_type获取 时间类型：时，天，月
func (obj *_CarersMgr) WithTimeType(timeType string) Option {
	return optionFunc(func(o *options) { o.query["time_type"] = timeType })
}

// WithCarerDetail carer_detail获取 服务内容:采用富文本
func (obj *_CarersMgr) WithCarerDetail(carerDetail string) Option {
	return optionFunc(func(o *options) { o.query["carer_detail"] = carerDetail })
}

// WithCardNo card_no获取 工号
func (obj *_CarersMgr) WithCardNo(cardNo string) Option {
	return optionFunc(func(o *options) { o.query["card_no"] = cardNo })
}

// GetByOption 功能选项模式获取
func (obj *_CarersMgr) GetByOption(opts ...Option) (result Carers, err error) {
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
func (obj *_CarersMgr) GetByOptions(opts ...Option) (results []*Carers, err error) {
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
func (obj *_CarersMgr) GetFromID(id int) (result Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CarersMgr) GetBatchFromID(ids []int) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_CarersMgr) GetFromName(name string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 姓名
func (obj *_CarersMgr) GetBatchFromName(names []string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容 年龄
func (obj *_CarersMgr) GetFromAge(age int) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量唯一主键查找 年龄
func (obj *_CarersMgr) GetBatchFromAge(ages []int) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age IN (?)", ages).Find(&results).Error

	return
}

// GetFromWorkExp 通过work_exp获取内容 工作年限
func (obj *_CarersMgr) GetFromWorkExp(workExp int) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("work_exp = ?", workExp).Find(&results).Error

	return
}

// GetBatchFromWorkExp 批量唯一主键查找 工作年限
func (obj *_CarersMgr) GetBatchFromWorkExp(workExps []int) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("work_exp IN (?)", workExps).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_CarersMgr) GetFromPrice(price float64) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量唯一主键查找 价格
func (obj *_CarersMgr) GetBatchFromPrice(prices []float64) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 性别:1:男,0:女
func (obj *_CarersMgr) GetFromSex(sex int) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找 性别:1:男,0:女
func (obj *_CarersMgr) GetBatchFromSex(sexs []int) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", sexs).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机
func (obj *_CarersMgr) GetFromPhone(phone string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 手机
func (obj *_CarersMgr) GetBatchFromPhone(phones []string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 简介
func (obj *_CarersMgr) GetFromDesc(desc string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量唯一主键查找 简介
func (obj *_CarersMgr) GetBatchFromDesc(descs []string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("desc IN (?)", descs).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像
func (obj *_CarersMgr) GetFromAvatar(avatar string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("avatar = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量唯一主键查找 头像
func (obj *_CarersMgr) GetBatchFromAvatar(avatars []string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("avatar IN (?)", avatars).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 服务数量
func (obj *_CarersMgr) GetFromAmount(amount int) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量唯一主键查找 服务数量
func (obj *_CarersMgr) GetBatchFromAmount(amounts []int) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("amount IN (?)", amounts).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CarersMgr) GetFromCreateAt(createAt time.Time) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CarersMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CarersMgr) GetFromUpdateAt(updateAt time.Time) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CarersMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CarersMgr) GetFromIsDeleted(isDeleted time.Time) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CarersMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容 application_id
func (obj *_CarersMgr) GetFromApplicationID(applicationID int) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量唯一主键查找 application_id
func (obj *_CarersMgr) GetBatchFromApplicationID(applicationIDs []int) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id IN (?)", applicationIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态：上下架
func (obj *_CarersMgr) GetFromStatus(status bool) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 状态：上下架
func (obj *_CarersMgr) GetBatchFromStatus(statuss []bool) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTimeType 通过time_type获取内容 时间类型：时，天，月
func (obj *_CarersMgr) GetFromTimeType(timeType string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type = ?", timeType).Find(&results).Error

	return
}

// GetBatchFromTimeType 批量唯一主键查找 时间类型：时，天，月
func (obj *_CarersMgr) GetBatchFromTimeType(timeTypes []string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("time_type IN (?)", timeTypes).Find(&results).Error

	return
}

// GetFromCarerDetail 通过carer_detail获取内容 服务内容:采用富文本
func (obj *_CarersMgr) GetFromCarerDetail(carerDetail string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_detail = ?", carerDetail).Find(&results).Error

	return
}

// GetBatchFromCarerDetail 批量唯一主键查找 服务内容:采用富文本
func (obj *_CarersMgr) GetBatchFromCarerDetail(carerDetails []string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("carer_detail IN (?)", carerDetails).Find(&results).Error

	return
}

// GetFromCardNo 通过card_no获取内容 工号
func (obj *_CarersMgr) GetFromCardNo(cardNo string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("card_no = ?", cardNo).Find(&results).Error

	return
}

// GetBatchFromCardNo 批量唯一主键查找 工号
func (obj *_CarersMgr) GetBatchFromCardNo(cardNos []string) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("card_no IN (?)", cardNos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CarersMgr) FetchByPrimaryKey(id int) (result Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByApplicationID  获取多个内容
func (obj *_CarersMgr) FetchIndexByApplicationID(applicationID int) (results []*Carers, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ?", applicationID).Find(&results).Error

	return
}
