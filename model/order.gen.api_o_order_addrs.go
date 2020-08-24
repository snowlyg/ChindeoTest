package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIOOrderAddrsMgr struct {
	*_BaseMgr
}

// APIOOrderAddrsMgr open func
func APIOOrderAddrsMgr(db *gorm.DB) *_APIOOrderAddrsMgr {
	if db == nil {
		panic(fmt.Errorf("APIOOrderAddrsMgr need init by db"))
	}
	return &_APIOOrderAddrsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_o_order_addrs"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIOOrderAddrsMgr) GetTableName() string {
	return "api_o_order_addrs"
}

// Get 获取
func (obj *_APIOOrderAddrsMgr) Get() (result APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIOOrderAddrsMgr) Gets() (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIOOrderAddrsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 联系人姓名
func (obj *_APIOOrderAddrsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPhone phone获取 手机
func (obj *_APIOOrderAddrsMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithAddr addr获取 地址
func (obj *_APIOOrderAddrsMgr) WithAddr(addr string) Option {
	return optionFunc(func(o *options) { o.query["addr"] = addr })
}

// WithSex sex获取 性别:1:男,0:女
func (obj *_APIOOrderAddrsMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithCreateAt create_at获取
func (obj *_APIOOrderAddrsMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIOOrderAddrsMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithIsDeleted is_deleted获取
func (obj *_APIOOrderAddrsMgr) WithIsDeleted(IsDeleted sql.NullTime) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = IsDeleted })
}

// WithOOrderID o_order_id获取 o_order_id
func (obj *_APIOOrderAddrsMgr) WithOOrderID(oOrderID int) Option {
	return optionFunc(func(o *options) { o.query["o_order_id"] = oOrderID })
}

// WithHospitalName hospital_name获取 医院名称
func (obj *_APIOOrderAddrsMgr) WithHospitalName(hospitalName string) Option {
	return optionFunc(func(o *options) { o.query["hospital_name"] = hospitalName })
}

// WithLocName loc_name获取 病区名称
func (obj *_APIOOrderAddrsMgr) WithLocName(locName string) Option {
	return optionFunc(func(o *options) { o.query["loc_name"] = locName })
}

// WithBedNum bed_num获取 床号
func (obj *_APIOOrderAddrsMgr) WithBedNum(bedNum string) Option {
	return optionFunc(func(o *options) { o.query["bed_num"] = bedNum })
}

// WithHospitalNo hospital_no获取 住院号
func (obj *_APIOOrderAddrsMgr) WithHospitalNo(hospitalNo string) Option {
	return optionFunc(func(o *options) { o.query["hospital_no"] = hospitalNo })
}

// WithDisease disease获取 病种
func (obj *_APIOOrderAddrsMgr) WithDisease(disease string) Option {
	return optionFunc(func(o *options) { o.query["disease"] = disease })
}

// WithAge age获取 年龄
func (obj *_APIOOrderAddrsMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// GetByOption 功能选项模式获取
func (obj *_APIOOrderAddrsMgr) GetByOption(opts ...Option) (result APIOOrderAddrs, err error) {
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
func (obj *_APIOOrderAddrsMgr) GetByOptions(opts ...Option) (results []*APIOOrderAddrs, err error) {
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
func (obj *_APIOOrderAddrsMgr) GetFromID(id int) (result APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIOOrderAddrsMgr) GetBatchFromID(ids []int) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 联系人姓名
func (obj *_APIOOrderAddrsMgr) GetFromName(name string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 联系人姓名
func (obj *_APIOOrderAddrsMgr) GetBatchFromName(names []string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机
func (obj *_APIOOrderAddrsMgr) GetFromPhone(phone string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 手机
func (obj *_APIOOrderAddrsMgr) GetBatchFromPhone(phones []string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromAddr 通过addr获取内容 地址
func (obj *_APIOOrderAddrsMgr) GetFromAddr(addr string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addr = ?", addr).Find(&results).Error

	return
}

// GetBatchFromAddr 批量唯一主键查找 地址
func (obj *_APIOOrderAddrsMgr) GetBatchFromAddr(addrs []string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("addr IN (?)", addrs).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 性别:1:男,0:女
func (obj *_APIOOrderAddrsMgr) GetFromSex(sex int) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找 性别:1:男,0:女
func (obj *_APIOOrderAddrsMgr) GetBatchFromSex(sexs []int) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", sexs).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIOOrderAddrsMgr) GetFromCreateAt(createAt time.Time) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIOOrderAddrsMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIOOrderAddrsMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIOOrderAddrsMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIOOrderAddrsMgr) GetFromIsDeleted(IsDeleted sql.NullTime) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", IsDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIOOrderAddrsMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromOOrderID 通过o_order_id获取内容 o_order_id
func (obj *_APIOOrderAddrsMgr) GetFromOOrderID(oOrderID int) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("o_order_id = ?", oOrderID).Find(&results).Error

	return
}

// GetBatchFromOOrderID 批量唯一主键查找 o_order_id
func (obj *_APIOOrderAddrsMgr) GetBatchFromOOrderID(oOrderIDs []int) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("o_order_id IN (?)", oOrderIDs).Find(&results).Error

	return
}

// GetFromHospitalName 通过hospital_name获取内容 医院名称
func (obj *_APIOOrderAddrsMgr) GetFromHospitalName(hospitalName string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_name = ?", hospitalName).Find(&results).Error

	return
}

// GetBatchFromHospitalName 批量唯一主键查找 医院名称
func (obj *_APIOOrderAddrsMgr) GetBatchFromHospitalName(hospitalNames []string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_name IN (?)", hospitalNames).Find(&results).Error

	return
}

// GetFromLocName 通过loc_name获取内容 病区名称
func (obj *_APIOOrderAddrsMgr) GetFromLocName(locName string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loc_name = ?", locName).Find(&results).Error

	return
}

// GetBatchFromLocName 批量唯一主键查找 病区名称
func (obj *_APIOOrderAddrsMgr) GetBatchFromLocName(locNames []string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loc_name IN (?)", locNames).Find(&results).Error

	return
}

// GetFromBedNum 通过bed_num获取内容 床号
func (obj *_APIOOrderAddrsMgr) GetFromBedNum(bedNum string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bed_num = ?", bedNum).Find(&results).Error

	return
}

// GetBatchFromBedNum 批量唯一主键查找 床号
func (obj *_APIOOrderAddrsMgr) GetBatchFromBedNum(bedNums []string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bed_num IN (?)", bedNums).Find(&results).Error

	return
}

// GetFromHospitalNo 通过hospital_no获取内容 住院号
func (obj *_APIOOrderAddrsMgr) GetFromHospitalNo(hospitalNo string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_no = ?", hospitalNo).Find(&results).Error

	return
}

// GetBatchFromHospitalNo 批量唯一主键查找 住院号
func (obj *_APIOOrderAddrsMgr) GetBatchFromHospitalNo(hospitalNos []string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_no IN (?)", hospitalNos).Find(&results).Error

	return
}

// GetFromDisease 通过disease获取内容 病种
func (obj *_APIOOrderAddrsMgr) GetFromDisease(disease string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("disease = ?", disease).Find(&results).Error

	return
}

// GetBatchFromDisease 批量唯一主键查找 病种
func (obj *_APIOOrderAddrsMgr) GetBatchFromDisease(diseases []string) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("disease IN (?)", diseases).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容 年龄
func (obj *_APIOOrderAddrsMgr) GetFromAge(age int) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量唯一主键查找 年龄
func (obj *_APIOOrderAddrsMgr) GetBatchFromAge(ages []int) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age IN (?)", ages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIOOrderAddrsMgr) FetchByPrimaryKey(id int) (result APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByOOrderID  获取多个内容
func (obj *_APIOOrderAddrsMgr) FetchIndexByOOrderID(oOrderID int) (results []*APIOOrderAddrs, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("o_order_id = ?", oOrderID).Find(&results).Error

	return
}
