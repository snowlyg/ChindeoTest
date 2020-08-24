package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _APIProfilesMgr struct {
	*_BaseMgr
}

// APIProfilesMgr open func
func APIProfilesMgr(db *gorm.DB) *_APIProfilesMgr {
	if db == nil {
		panic(fmt.Errorf("APIProfilesMgr need init by db"))
	}
	return &_APIProfilesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_profiles"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APIProfilesMgr) GetTableName() string {
	return "api_profiles"
}

// Get 获取
func (obj *_APIProfilesMgr) Get() (result APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APIProfilesMgr) Gets() (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APIProfilesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 联系人姓名
func (obj *_APIProfilesMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPhone phone获取 手机
func (obj *_APIProfilesMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithSex sex获取 性别:1:男,0:女
func (obj *_APIProfilesMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithHospitalName hospital_name获取 医院名称
func (obj *_APIProfilesMgr) WithHospitalName(hospitalName string) Option {
	return optionFunc(func(o *options) { o.query["hospital_name"] = hospitalName })
}

// WithLocName loc_name获取 病区名称
func (obj *_APIProfilesMgr) WithLocName(locName string) Option {
	return optionFunc(func(o *options) { o.query["loc_name"] = locName })
}

// WithBedNum bed_num获取 床号
func (obj *_APIProfilesMgr) WithBedNum(bedNum string) Option {
	return optionFunc(func(o *options) { o.query["bed_num"] = bedNum })
}

// WithHospitalNo hospital_no获取 住院号
func (obj *_APIProfilesMgr) WithHospitalNo(hospitalNo string) Option {
	return optionFunc(func(o *options) { o.query["hospital_no"] = hospitalNo })
}

// WithDisease disease获取 病种
func (obj *_APIProfilesMgr) WithDisease(disease string) Option {
	return optionFunc(func(o *options) { o.query["disease"] = disease })
}

// WithAge age获取 年龄
func (obj *_APIProfilesMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// WithCreateAt create_at获取
func (obj *_APIProfilesMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取
func (obj *_APIProfilesMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithUserID user_id获取
func (obj *_APIProfilesMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithIsDeleted is_deleted获取
func (obj *_APIProfilesMgr) WithIsDeleted(isDeleted time.Time) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// GetByOption 功能选项模式获取
func (obj *_APIProfilesMgr) GetByOption(opts ...Option) (result APIProfiles, err error) {
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
func (obj *_APIProfilesMgr) GetByOptions(opts ...Option) (results []*APIProfiles, err error) {
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
func (obj *_APIProfilesMgr) GetFromID(id int) (result APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_APIProfilesMgr) GetBatchFromID(ids []int) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 联系人姓名
func (obj *_APIProfilesMgr) GetFromName(name string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 联系人姓名
func (obj *_APIProfilesMgr) GetBatchFromName(names []string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机
func (obj *_APIProfilesMgr) GetFromPhone(phone string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量唯一主键查找 手机
func (obj *_APIProfilesMgr) GetBatchFromPhone(phones []string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("phone IN (?)", phones).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 性别:1:男,0:女
func (obj *_APIProfilesMgr) GetFromSex(sex int) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量唯一主键查找 性别:1:男,0:女
func (obj *_APIProfilesMgr) GetBatchFromSex(sexs []int) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sex IN (?)", sexs).Find(&results).Error

	return
}

// GetFromHospitalName 通过hospital_name获取内容 医院名称
func (obj *_APIProfilesMgr) GetFromHospitalName(hospitalName string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_name = ?", hospitalName).Find(&results).Error

	return
}

// GetBatchFromHospitalName 批量唯一主键查找 医院名称
func (obj *_APIProfilesMgr) GetBatchFromHospitalName(hospitalNames []string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_name IN (?)", hospitalNames).Find(&results).Error

	return
}

// GetFromLocName 通过loc_name获取内容 病区名称
func (obj *_APIProfilesMgr) GetFromLocName(locName string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loc_name = ?", locName).Find(&results).Error

	return
}

// GetBatchFromLocName 批量唯一主键查找 病区名称
func (obj *_APIProfilesMgr) GetBatchFromLocName(locNames []string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loc_name IN (?)", locNames).Find(&results).Error

	return
}

// GetFromBedNum 通过bed_num获取内容 床号
func (obj *_APIProfilesMgr) GetFromBedNum(bedNum string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bed_num = ?", bedNum).Find(&results).Error

	return
}

// GetBatchFromBedNum 批量唯一主键查找 床号
func (obj *_APIProfilesMgr) GetBatchFromBedNum(bedNums []string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("bed_num IN (?)", bedNums).Find(&results).Error

	return
}

// GetFromHospitalNo 通过hospital_no获取内容 住院号
func (obj *_APIProfilesMgr) GetFromHospitalNo(hospitalNo string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_no = ?", hospitalNo).Find(&results).Error

	return
}

// GetBatchFromHospitalNo 批量唯一主键查找 住院号
func (obj *_APIProfilesMgr) GetBatchFromHospitalNo(hospitalNos []string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("hospital_no IN (?)", hospitalNos).Find(&results).Error

	return
}

// GetFromDisease 通过disease获取内容 病种
func (obj *_APIProfilesMgr) GetFromDisease(disease string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("disease = ?", disease).Find(&results).Error

	return
}

// GetBatchFromDisease 批量唯一主键查找 病种
func (obj *_APIProfilesMgr) GetBatchFromDisease(diseases []string) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("disease IN (?)", diseases).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容 年龄
func (obj *_APIProfilesMgr) GetFromAge(age int) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量唯一主键查找 年龄
func (obj *_APIProfilesMgr) GetBatchFromAge(ages []int) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("age IN (?)", ages).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容
func (obj *_APIProfilesMgr) GetFromCreateAt(createAt time.Time) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找
func (obj *_APIProfilesMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_APIProfilesMgr) GetFromUpdateAt(updateAt time.Time) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量唯一主键查找
func (obj *_APIProfilesMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("update_at IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_APIProfilesMgr) GetFromUserID(userID int) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_APIProfilesMgr) GetBatchFromUserID(userIDs []int) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_APIProfilesMgr) GetFromIsDeleted(isDeleted time.Time) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量唯一主键查找
func (obj *_APIProfilesMgr) GetBatchFromIsDeleted(isDeleteds []time.Time) (results []*APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_APIProfilesMgr) FetchByPrimaryKey(id int) (result APIProfiles, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
