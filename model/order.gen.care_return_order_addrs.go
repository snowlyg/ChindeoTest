package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CareReturnOrderAddrsMgr struct {
	*_BaseMgr
}

// CareReturnOrderAddrsMgr open func
func CareReturnOrderAddrsMgr(db *gorm.DB) *_CareReturnOrderAddrsMgr {
	if db == nil {
		panic(fmt.Errorf("CareReturnOrderAddrsMgr need init by db"))
	}
	return &_CareReturnOrderAddrsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("care_return_order_addrs"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CareReturnOrderAddrsMgr) GetTableName() string {
	return "care_return_order_addrs"
}

// Get 获取
func (obj *_CareReturnOrderAddrsMgr) Get() (result CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CareReturnOrderAddrsMgr) Gets() (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CareReturnOrderAddrsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 联系人姓名
func (obj *_CareReturnOrderAddrsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPhone phone获取 手机
func (obj *_CareReturnOrderAddrsMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithAddr addr获取 地址
func (obj *_CareReturnOrderAddrsMgr) WithAddr(addr string) Option {
	return optionFunc(func(o *options) { o.query["addr"] = addr })
}

// WithSex sex获取 性别:1:男,0:女
func (obj *_CareReturnOrderAddrsMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithHospitalName hospital_name获取 医院名称
func (obj *_CareReturnOrderAddrsMgr) WithHospitalName(hospitalName string) Option {
	return optionFunc(func(o *options) { o.query["hospital_name"] = hospitalName })
}

// WithLocName loc_name获取 病区名称
func (obj *_CareReturnOrderAddrsMgr) WithLocName(locName string) Option {
	return optionFunc(func(o *options) { o.query["loc_name"] = locName })
}

// WithBedNum bed_num获取 床号
func (obj *_CareReturnOrderAddrsMgr) WithBedNum(bedNum string) Option {
	return optionFunc(func(o *options) { o.query["bed_num"] = bedNum })
}

// WithHospitalNo hospital_no获取 住院号
func (obj *_CareReturnOrderAddrsMgr) WithHospitalNo(hospitalNo string) Option {
	return optionFunc(func(o *options) { o.query["hospital_no"] = hospitalNo })
}

// WithDisease disease获取 病种
func (obj *_CareReturnOrderAddrsMgr) WithDisease(disease string) Option {
	return optionFunc(func(o *options) { o.query["disease"] = disease })
}

// WithCreateAt create_at获取
func (obj *_CareReturnOrderAddrsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CareReturnOrderAddrsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_CareReturnOrderAddrsMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithApplicationID application_id获取 application_id
func (obj *_CareReturnOrderAddrsMgr) WithApplicationID(applicationID int) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithCareReturnOrderID care_return_order_id获取 care_return_order_id
func (obj *_CareReturnOrderAddrsMgr) WithCareReturnOrderID(careReturnOrderID int) Option {
	return optionFunc(func(o *options) { o.query["care_return_order_id"] = careReturnOrderID })
}

// WithAge age获取 年龄
func (obj *_CareReturnOrderAddrsMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// GetByOption 功能选项模式获取
func (obj *_CareReturnOrderAddrsMgr) GetByOption(opts ...Option) (result CareReturnOrderAddrs, err error) {
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
func (obj *_CareReturnOrderAddrsMgr) GetByOptions(opts ...Option) (results []*CareReturnOrderAddrs, err error) {
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
func (obj *_CareReturnOrderAddrsMgr) GetFromID(id int) (result CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromID(ids []int) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 联系人姓名
func (obj *_CareReturnOrderAddrsMgr) GetFromName(name string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 联系人姓名
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromName(names []string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机
func (obj *_CareReturnOrderAddrsMgr) GetFromPhone(phone string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 手机
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromPhone(phones []string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromAddr 通过addr获取内容 地址
func (obj *_CareReturnOrderAddrsMgr) GetFromAddr(addr string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addr = ?", addr).Find(&results).Error

	return
}

// GetBatchFromAddr 批量唯一主键查找 地址
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromAddr(addrs []string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addr IN (?)", addrs).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 性别:1:男,0:女
func (obj *_CareReturnOrderAddrsMgr) GetFromSex(sex int) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找 性别:1:男,0:女
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromSex(sexs []int) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", sexs).Find(&results).Error

	return
}

// GetFromHospitalName 通过hospital_name获取内容 医院名称
func (obj *_CareReturnOrderAddrsMgr) GetFromHospitalName(hospitalName string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_name = ?", hospitalName).Find(&results).Error

	return
}

// GetBatchFromHospitalName 批量唯一主键查找 医院名称
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromHospitalName(hospitalNames []string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_name IN (?)", hospitalNames).Find(&results).Error

	return
}

// GetFromLocName 通过loc_name获取内容 病区名称
func (obj *_CareReturnOrderAddrsMgr) GetFromLocName(locName string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loc_name = ?", locName).Find(&results).Error

	return
}

// GetBatchFromLocName 批量唯一主键查找 病区名称
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromLocName(locNames []string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loc_name IN (?)", locNames).Find(&results).Error

	return
}

// GetFromBedNum 通过bed_num获取内容 床号
func (obj *_CareReturnOrderAddrsMgr) GetFromBedNum(bedNum string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bed_num = ?", bedNum).Find(&results).Error

	return
}

// GetBatchFromBedNum 批量唯一主键查找 床号
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromBedNum(bedNums []string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bed_num IN (?)", bedNums).Find(&results).Error

	return
}

// GetFromHospitalNo 通过hospital_no获取内容 住院号
func (obj *_CareReturnOrderAddrsMgr) GetFromHospitalNo(hospitalNo string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_no = ?", hospitalNo).Find(&results).Error

	return
}

// GetBatchFromHospitalNo 批量唯一主键查找 住院号
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromHospitalNo(hospitalNos []string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_no IN (?)", hospitalNos).Find(&results).Error

	return
}

// GetFromDisease 通过disease获取内容 病种
func (obj *_CareReturnOrderAddrsMgr) GetFromDisease(disease string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("disease = ?", disease).Find(&results).Error

	return
}

// GetBatchFromDisease 批量唯一主键查找 病种
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromDisease(diseases []string) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("disease IN (?)", diseases).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CareReturnOrderAddrsMgr) GetFromCreateAt(createAt time.Time) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CareReturnOrderAddrsMgr) GetFromUpdateAt(updateAt time.Time) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CareReturnOrderAddrsMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容 application_id
func (obj *_CareReturnOrderAddrsMgr) GetFromApplicationID(applicationID int) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量唯一主键查找 application_id
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromApplicationID(applicationIDs []int) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id IN (?)", applicationIDs).Find(&results).Error

	return
}

// GetFromCareReturnOrderID 通过care_return_order_id获取内容 care_return_order_id
func (obj *_CareReturnOrderAddrsMgr) GetFromCareReturnOrderID(careReturnOrderID int) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_return_order_id = ?", careReturnOrderID).Find(&results).Error

	return
}

// GetBatchFromCareReturnOrderID 批量唯一主键查找 care_return_order_id
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromCareReturnOrderID(careReturnOrderIDs []int) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_return_order_id IN (?)", careReturnOrderIDs).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容 年龄
func (obj *_CareReturnOrderAddrsMgr) GetFromAge(age int) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量唯一主键查找 年龄
func (obj *_CareReturnOrderAddrsMgr) GetBatchFromAge(ages []int) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age IN (?)", ages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CareReturnOrderAddrsMgr) FetchByPrimaryKey(id int) (result CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByCareReturnOrderID  获取多个内容
func (obj *_CareReturnOrderAddrsMgr) FetchIndexByCareReturnOrderID(careReturnOrderID int) (results []*CareReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_return_order_id = ?", careReturnOrderID).Find(&results).Error

	return
}
