package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _CareOrderAddrsMgr struct {
	*_BaseMgr
}

// CareOrderAddrsMgr open func
func CareOrderAddrsMgr(db *gorm.DB) *_CareOrderAddrsMgr {
	if db == nil {
		panic(fmt.Errorf("CareOrderAddrsMgr need init by db"))
	}
	return &_CareOrderAddrsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("care_order_addrs"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CareOrderAddrsMgr) GetTableName() string {
	return "care_order_addrs"
}

// Get 获取
func (obj *_CareOrderAddrsMgr) Get() (result CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CareOrderAddrsMgr) Gets() (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CareOrderAddrsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 联系人姓名
func (obj *_CareOrderAddrsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPhone phone获取 手机
func (obj *_CareOrderAddrsMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithAddr addr获取 地址
func (obj *_CareOrderAddrsMgr) WithAddr(addr string) Option {
	return optionFunc(func(o *options) { o.query["addr"] = addr })
}

// WithSex sex获取 性别:1:男,0:女
func (obj *_CareOrderAddrsMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithHospitalName hospital_name获取 医院名称
func (obj *_CareOrderAddrsMgr) WithHospitalName(hospitalName string) Option {
	return optionFunc(func(o *options) { o.query["hospital_name"] = hospitalName })
}

// WithLocName loc_name获取 病区名称
func (obj *_CareOrderAddrsMgr) WithLocName(locName string) Option {
	return optionFunc(func(o *options) { o.query["loc_name"] = locName })
}

// WithBedNum bed_num获取 床号
func (obj *_CareOrderAddrsMgr) WithBedNum(bedNum string) Option {
	return optionFunc(func(o *options) { o.query["bed_num"] = bedNum })
}

// WithHospitalNo hospital_no获取 住院号
func (obj *_CareOrderAddrsMgr) WithHospitalNo(hospitalNo string) Option {
	return optionFunc(func(o *options) { o.query["hospital_no"] = hospitalNo })
}

// WithDisease disease获取 病种
func (obj *_CareOrderAddrsMgr) WithDisease(disease string) Option {
	return optionFunc(func(o *options) { o.query["disease"] = disease })
}

// WithIsDeleted is_deleted获取
func (obj *_CareOrderAddrsMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithCareOrderID care_order_id获取 care_order_id
func (obj *_CareOrderAddrsMgr) WithCareOrderID(careOrderID int) Option {
	return optionFunc(func(o *options) { o.query["care_order_id"] = careOrderID })
}

// WithCreateAt create_at获取
func (obj *_CareOrderAddrsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_CareOrderAddrsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithAge age获取 年龄
func (obj *_CareOrderAddrsMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// GetByOption 功能选项模式获取
func (obj *_CareOrderAddrsMgr) GetByOption(opts ...Option) (result CareOrderAddrs, err error) {
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
func (obj *_CareOrderAddrsMgr) GetByOptions(opts ...Option) (results []*CareOrderAddrs, err error) {
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
func (obj *_CareOrderAddrsMgr) GetFromID(id int) (result CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_CareOrderAddrsMgr) GetBatchFromID(ids []int) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 联系人姓名
func (obj *_CareOrderAddrsMgr) GetFromName(name string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 联系人姓名
func (obj *_CareOrderAddrsMgr) GetBatchFromName(names []string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机
func (obj *_CareOrderAddrsMgr) GetFromPhone(phone string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 手机
func (obj *_CareOrderAddrsMgr) GetBatchFromPhone(phones []string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromAddr 通过addr获取内容 地址
func (obj *_CareOrderAddrsMgr) GetFromAddr(addr string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addr = ?", addr).Find(&results).Error

	return
}

// GetBatchFromAddr 批量唯一主键查找 地址
func (obj *_CareOrderAddrsMgr) GetBatchFromAddr(addrs []string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addr IN (?)", addrs).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 性别:1:男,0:女
func (obj *_CareOrderAddrsMgr) GetFromSex(sex int) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找 性别:1:男,0:女
func (obj *_CareOrderAddrsMgr) GetBatchFromSex(sexs []int) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", sexs).Find(&results).Error

	return
}

// GetFromHospitalName 通过hospital_name获取内容 医院名称
func (obj *_CareOrderAddrsMgr) GetFromHospitalName(hospitalName string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_name = ?", hospitalName).Find(&results).Error

	return
}

// GetBatchFromHospitalName 批量唯一主键查找 医院名称
func (obj *_CareOrderAddrsMgr) GetBatchFromHospitalName(hospitalNames []string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_name IN (?)", hospitalNames).Find(&results).Error

	return
}

// GetFromLocName 通过loc_name获取内容 病区名称
func (obj *_CareOrderAddrsMgr) GetFromLocName(locName string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loc_name = ?", locName).Find(&results).Error

	return
}

// GetBatchFromLocName 批量唯一主键查找 病区名称
func (obj *_CareOrderAddrsMgr) GetBatchFromLocName(locNames []string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loc_name IN (?)", locNames).Find(&results).Error

	return
}

// GetFromBedNum 通过bed_num获取内容 床号
func (obj *_CareOrderAddrsMgr) GetFromBedNum(bedNum string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bed_num = ?", bedNum).Find(&results).Error

	return
}

// GetBatchFromBedNum 批量唯一主键查找 床号
func (obj *_CareOrderAddrsMgr) GetBatchFromBedNum(bedNums []string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bed_num IN (?)", bedNums).Find(&results).Error

	return
}

// GetFromHospitalNo 通过hospital_no获取内容 住院号
func (obj *_CareOrderAddrsMgr) GetFromHospitalNo(hospitalNo string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_no = ?", hospitalNo).Find(&results).Error

	return
}

// GetBatchFromHospitalNo 批量唯一主键查找 住院号
func (obj *_CareOrderAddrsMgr) GetBatchFromHospitalNo(hospitalNos []string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_no IN (?)", hospitalNos).Find(&results).Error

	return
}

// GetFromDisease 通过disease获取内容 病种
func (obj *_CareOrderAddrsMgr) GetFromDisease(disease string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("disease = ?", disease).Find(&results).Error

	return
}

// GetBatchFromDisease 批量唯一主键查找 病种
func (obj *_CareOrderAddrsMgr) GetBatchFromDisease(diseases []string) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("disease IN (?)", diseases).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_CareOrderAddrsMgr) GetFromIsDeleted(isDeleted time.Time) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_CareOrderAddrsMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCareOrderID 通过care_order_id获取内容 care_order_id
func (obj *_CareOrderAddrsMgr) GetFromCareOrderID(careOrderID int) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id = ?", careOrderID).Find(&results).Error

	return
}

// GetBatchFromCareOrderID 批量唯一主键查找 care_order_id
func (obj *_CareOrderAddrsMgr) GetBatchFromCareOrderID(careOrderIDs []int) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id IN (?)", careOrderIDs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_CareOrderAddrsMgr) GetFromCreateAt(createAt time.Time) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_CareOrderAddrsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_CareOrderAddrsMgr) GetFromUpdateAt(updateAt time.Time) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_CareOrderAddrsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容 年龄
func (obj *_CareOrderAddrsMgr) GetFromAge(age int) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量唯一主键查找 年龄
func (obj *_CareOrderAddrsMgr) GetBatchFromAge(ages []int) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age IN (?)", ages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CareOrderAddrsMgr) FetchByPrimaryKey(id int) (result CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByCareOrderID  获取多个内容
func (obj *_CareOrderAddrsMgr) FetchIndexByCareOrderID(careOrderID int) (results []*CareOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("care_order_id = ?", careOrderID).Find(&results).Error

	return
}
