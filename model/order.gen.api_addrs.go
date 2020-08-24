package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIAddrsMgr struct {
	*_BaseMgr
}

// APIAddrsMgr open func
func APIAddrsMgr(db *gorm.DB) *_APIAddrsMgr {
	if db == nil {
		panic(fmt.Errorf("APIAddrsMgr need init by db"))
	}
	return &_APIAddrsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_addrs"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIAddrsMgr) GetTableName() string {
	return "api_addrs"
}

// Get 获取
func (obj *_APIAddrsMgr) Get() (result APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIAddrsMgr) Gets() (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIAddrsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 联系人姓名
func (obj *_APIAddrsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPhone phone获取 手机
func (obj *_APIAddrsMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithAddr addr获取 地址
func (obj *_APIAddrsMgr) WithAddr(addr string) Option {
	return optionFunc(func(o *options) { o.query["addr"] = addr })
}

// WithSex sex获取
func (obj *_APIAddrsMgr) WithSex(sex int8) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithIsDefault is_default获取 默认地址
func (obj *_APIAddrsMgr) WithIsDefault(isDefault bool) Option {
	return optionFunc(func(o *options) { o.query["is_default"] = isDefault })
}

// WithCreateAt create_at获取
func (obj *_APIAddrsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIAddrsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIAddrsMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithUserID user_id获取 user_id
func (obj *_APIAddrsMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithHospitalName hospital_name获取 医院名称
func (obj *_APIAddrsMgr) WithHospitalName(hospitalName string) Option {
	return optionFunc(func(o *options) { o.query["hospital_name"] = hospitalName })
}

// WithLocName loc_name获取 病区名称
func (obj *_APIAddrsMgr) WithLocName(locName string) Option {
	return optionFunc(func(o *options) { o.query["loc_name"] = locName })
}

// WithBedNum bed_num获取 床号
func (obj *_APIAddrsMgr) WithBedNum(bedNum string) Option {
	return optionFunc(func(o *options) { o.query["bed_num"] = bedNum })
}

// WithHospitalNo hospital_no获取 住院号
func (obj *_APIAddrsMgr) WithHospitalNo(hospitalNo string) Option {
	return optionFunc(func(o *options) { o.query["hospital_no"] = hospitalNo })
}

// WithDisease disease获取 病种
func (obj *_APIAddrsMgr) WithDisease(disease string) Option {
	return optionFunc(func(o *options) { o.query["disease"] = disease })
}

// WithAge age获取 年龄
func (obj *_APIAddrsMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// GetByOption 功能选项模式获取
func (obj *_APIAddrsMgr) GetByOption(opts ...Option) (result APIAddrs, err error) {
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
func (obj *_APIAddrsMgr) GetByOptions(opts ...Option) (results []*APIAddrs, err error) {
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
func (obj *_APIAddrsMgr) GetFromID(id int) (result APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIAddrsMgr) GetBatchFromID(ids []int) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 联系人姓名
func (obj *_APIAddrsMgr) GetFromName(name string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 联系人姓名
func (obj *_APIAddrsMgr) GetBatchFromName(names []string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机
func (obj *_APIAddrsMgr) GetFromPhone(phone string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 手机
func (obj *_APIAddrsMgr) GetBatchFromPhone(phones []string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromAddr 通过addr获取内容 地址
func (obj *_APIAddrsMgr) GetFromAddr(addr string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addr = ?", addr).Find(&results).Error

	return
}

// GetBatchFromAddr 批量唯一主键查找 地址
func (obj *_APIAddrsMgr) GetBatchFromAddr(addrs []string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addr IN (?)", addrs).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容
func (obj *_APIAddrsMgr) GetFromSex(sex int8) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找
func (obj *_APIAddrsMgr) GetBatchFromSex(sexs []int8) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", sexs).Find(&results).Error

	return
}

// GetFromIsDefault 通过is_default获取内容 默认地址
func (obj *_APIAddrsMgr) GetFromIsDefault(isDefault bool) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_default = ?", isDefault).Find(&results).Error

	return
}

// GetBatchFromIsDefault 批量唯一主键查找 默认地址
func (obj *_APIAddrsMgr) GetBatchFromIsDefault(isDefaults []bool) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_default IN (?)", isDefaults).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIAddrsMgr) GetFromCreateAt(createAt time.Time) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIAddrsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIAddrsMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIAddrsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIAddrsMgr) GetFromIsDeleted(isDeleted time.Time) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIAddrsMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 user_id
func (obj *_APIAddrsMgr) GetFromUserID(userID int) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找 user_id
func (obj *_APIAddrsMgr) GetBatchFromUserID(userIDs []int) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromHospitalName 通过hospital_name获取内容 医院名称
func (obj *_APIAddrsMgr) GetFromHospitalName(hospitalName string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_name = ?", hospitalName).Find(&results).Error

	return
}

// GetBatchFromHospitalName 批量唯一主键查找 医院名称
func (obj *_APIAddrsMgr) GetBatchFromHospitalName(hospitalNames []string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_name IN (?)", hospitalNames).Find(&results).Error

	return
}

// GetFromLocName 通过loc_name获取内容 病区名称
func (obj *_APIAddrsMgr) GetFromLocName(locName string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loc_name = ?", locName).Find(&results).Error

	return
}

// GetBatchFromLocName 批量唯一主键查找 病区名称
func (obj *_APIAddrsMgr) GetBatchFromLocName(locNames []string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loc_name IN (?)", locNames).Find(&results).Error

	return
}

// GetFromBedNum 通过bed_num获取内容 床号
func (obj *_APIAddrsMgr) GetFromBedNum(bedNum string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bed_num = ?", bedNum).Find(&results).Error

	return
}

// GetBatchFromBedNum 批量唯一主键查找 床号
func (obj *_APIAddrsMgr) GetBatchFromBedNum(bedNums []string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bed_num IN (?)", bedNums).Find(&results).Error

	return
}

// GetFromHospitalNo 通过hospital_no获取内容 住院号
func (obj *_APIAddrsMgr) GetFromHospitalNo(hospitalNo string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_no = ?", hospitalNo).Find(&results).Error

	return
}

// GetBatchFromHospitalNo 批量唯一主键查找 住院号
func (obj *_APIAddrsMgr) GetBatchFromHospitalNo(hospitalNos []string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_no IN (?)", hospitalNos).Find(&results).Error

	return
}

// GetFromDisease 通过disease获取内容 病种
func (obj *_APIAddrsMgr) GetFromDisease(disease string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("disease = ?", disease).Find(&results).Error

	return
}

// GetBatchFromDisease 批量唯一主键查找 病种
func (obj *_APIAddrsMgr) GetBatchFromDisease(diseases []string) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("disease IN (?)", diseases).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容 年龄
func (obj *_APIAddrsMgr) GetFromAge(age int) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量唯一主键查找 年龄
func (obj *_APIAddrsMgr) GetBatchFromAge(ages []int) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age IN (?)", ages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIAddrsMgr) FetchByPrimaryKey(id int) (result APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_APIAddrsMgr) FetchIndexByUserID(userID int) (results []*APIAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}
