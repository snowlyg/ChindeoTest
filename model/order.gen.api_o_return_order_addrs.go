package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIOReturnOrderAddrsMgr struct {
	*_BaseMgr
}

// APIOReturnOrderAddrsMgr open func
func APIOReturnOrderAddrsMgr(db *gorm.DB) *_APIOReturnOrderAddrsMgr {
	if db == nil {
		panic(fmt.Errorf("APIOReturnOrderAddrsMgr need init by db"))
	}
	return &_APIOReturnOrderAddrsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_o_return_order_addrs"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIOReturnOrderAddrsMgr) GetTableName() string {
	return "api_o_return_order_addrs"
}

// Get 获取
func (obj *_APIOReturnOrderAddrsMgr) Get() (result APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIOReturnOrderAddrsMgr) Gets() (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIOReturnOrderAddrsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 联系人姓名
func (obj *_APIOReturnOrderAddrsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPhone phone获取 手机
func (obj *_APIOReturnOrderAddrsMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithAddr addr获取 地址
func (obj *_APIOReturnOrderAddrsMgr) WithAddr(addr string) Option {
	return optionFunc(func(o *options) { o.query["addr"] = addr })
}

// WithSex sex获取 性别:1:男,0:女
func (obj *_APIOReturnOrderAddrsMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithCreateAt create_at获取
func (obj *_APIOReturnOrderAddrsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIOReturnOrderAddrsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIOReturnOrderAddrsMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithApplicationID application_id获取 application_id
func (obj *_APIOReturnOrderAddrsMgr) WithApplicationID(applicationID int) Option {
	return optionFunc(func(o *options) { o.query["application_id"] = applicationID })
}

// WithOReturnOrderID o_return_order_id获取 o_return_order_id
func (obj *_APIOReturnOrderAddrsMgr) WithOReturnOrderID(oReturnOrderID int) Option {
	return optionFunc(func(o *options) { o.query["o_return_order_id"] = oReturnOrderID })
}

// WithHospitalName hospital_name获取 医院名称
func (obj *_APIOReturnOrderAddrsMgr) WithHospitalName(hospitalName string) Option {
	return optionFunc(func(o *options) { o.query["hospital_name"] = hospitalName })
}

// WithLocName loc_name获取 病区名称
func (obj *_APIOReturnOrderAddrsMgr) WithLocName(locName string) Option {
	return optionFunc(func(o *options) { o.query["loc_name"] = locName })
}

// WithBedNum bed_num获取 床号
func (obj *_APIOReturnOrderAddrsMgr) WithBedNum(bedNum string) Option {
	return optionFunc(func(o *options) { o.query["bed_num"] = bedNum })
}

// WithHospitalNo hospital_no获取 住院号
func (obj *_APIOReturnOrderAddrsMgr) WithHospitalNo(hospitalNo string) Option {
	return optionFunc(func(o *options) { o.query["hospital_no"] = hospitalNo })
}

// WithDisease disease获取 病种
func (obj *_APIOReturnOrderAddrsMgr) WithDisease(disease string) Option {
	return optionFunc(func(o *options) { o.query["disease"] = disease })
}

// WithAge age获取 年龄
func (obj *_APIOReturnOrderAddrsMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// GetByOption 功能选项模式获取
func (obj *_APIOReturnOrderAddrsMgr) GetByOption(opts ...Option) (result APIOReturnOrderAddrs, err error) {
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
func (obj *_APIOReturnOrderAddrsMgr) GetByOptions(opts ...Option) (results []*APIOReturnOrderAddrs, err error) {
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
func (obj *_APIOReturnOrderAddrsMgr) GetFromID(id int) (result APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromID(ids []int) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 联系人姓名
func (obj *_APIOReturnOrderAddrsMgr) GetFromName(name string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 联系人姓名
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromName(names []string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机
func (obj *_APIOReturnOrderAddrsMgr) GetFromPhone(phone string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 手机
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromPhone(phones []string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromAddr 通过addr获取内容 地址
func (obj *_APIOReturnOrderAddrsMgr) GetFromAddr(addr string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addr = ?", addr).Find(&results).Error

	return
}

// GetBatchFromAddr 批量唯一主键查找 地址
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromAddr(addrs []string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addr IN (?)", addrs).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 性别:1:男,0:女
func (obj *_APIOReturnOrderAddrsMgr) GetFromSex(sex int) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找 性别:1:男,0:女
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromSex(sexs []int) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", sexs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIOReturnOrderAddrsMgr) GetFromCreateAt(createAt time.Time) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIOReturnOrderAddrsMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIOReturnOrderAddrsMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromApplicationID 通过application_id获取内容 application_id
func (obj *_APIOReturnOrderAddrsMgr) GetFromApplicationID(applicationID int) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id = ?", applicationID).Find(&results).Error

	return
}

// GetBatchFromApplicationID 批量唯一主键查找 application_id
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromApplicationID(applicationIDs []int) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("application_id IN (?)", applicationIDs).Find(&results).Error

	return
}

// GetFromOReturnOrderID 通过o_return_order_id获取内容 o_return_order_id
func (obj *_APIOReturnOrderAddrsMgr) GetFromOReturnOrderID(oReturnOrderID int) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("o_return_order_id = ?", oReturnOrderID).Find(&results).Error

	return
}

// GetBatchFromOReturnOrderID 批量唯一主键查找 o_return_order_id
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromOReturnOrderID(oReturnOrderIDs []int) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("o_return_order_id IN (?)", oReturnOrderIDs).Find(&results).Error

	return
}

// GetFromHospitalName 通过hospital_name获取内容 医院名称
func (obj *_APIOReturnOrderAddrsMgr) GetFromHospitalName(hospitalName string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_name = ?", hospitalName).Find(&results).Error

	return
}

// GetBatchFromHospitalName 批量唯一主键查找 医院名称
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromHospitalName(hospitalNames []string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_name IN (?)", hospitalNames).Find(&results).Error

	return
}

// GetFromLocName 通过loc_name获取内容 病区名称
func (obj *_APIOReturnOrderAddrsMgr) GetFromLocName(locName string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loc_name = ?", locName).Find(&results).Error

	return
}

// GetBatchFromLocName 批量唯一主键查找 病区名称
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromLocName(locNames []string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loc_name IN (?)", locNames).Find(&results).Error

	return
}

// GetFromBedNum 通过bed_num获取内容 床号
func (obj *_APIOReturnOrderAddrsMgr) GetFromBedNum(bedNum string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bed_num = ?", bedNum).Find(&results).Error

	return
}

// GetBatchFromBedNum 批量唯一主键查找 床号
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromBedNum(bedNums []string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bed_num IN (?)", bedNums).Find(&results).Error

	return
}

// GetFromHospitalNo 通过hospital_no获取内容 住院号
func (obj *_APIOReturnOrderAddrsMgr) GetFromHospitalNo(hospitalNo string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_no = ?", hospitalNo).Find(&results).Error

	return
}

// GetBatchFromHospitalNo 批量唯一主键查找 住院号
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromHospitalNo(hospitalNos []string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_no IN (?)", hospitalNos).Find(&results).Error

	return
}

// GetFromDisease 通过disease获取内容 病种
func (obj *_APIOReturnOrderAddrsMgr) GetFromDisease(disease string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("disease = ?", disease).Find(&results).Error

	return
}

// GetBatchFromDisease 批量唯一主键查找 病种
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromDisease(diseases []string) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("disease IN (?)", diseases).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容 年龄
func (obj *_APIOReturnOrderAddrsMgr) GetFromAge(age int) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量唯一主键查找 年龄
func (obj *_APIOReturnOrderAddrsMgr) GetBatchFromAge(ages []int) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age IN (?)", ages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIOReturnOrderAddrsMgr) FetchByPrimaryKey(id int) (result APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByOReturnOrderID  获取多个内容
func (obj *_APIOReturnOrderAddrsMgr) FetchIndexByOReturnOrderID(oReturnOrderID int) (results []*APIOReturnOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("o_return_order_id = ?", oReturnOrderID).Find(&results).Error

	return
}
