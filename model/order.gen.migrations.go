package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _MigrationsMgr struct {
	*_BaseMgr
}

// MigrationsMgr open func
func MigrationsMgr(db *gorm.DB) *_MigrationsMgr {
	if db == nil {
		panic(fmt.Errorf("MigrationsMgr need init by db"))
	}
	return &_MigrationsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("migrations"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MigrationsMgr) GetTableName() string {
	return "migrations"
}

// Get 获取
func (obj *_MigrationsMgr) Get() (result Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MigrationsMgr) Gets() (results []*Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithVersion version获取
func (obj *_MigrationsMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithMigrationName migration_name获取
func (obj *_MigrationsMgr) WithMigrationName(migrationName string) Option {
	return optionFunc(func(o *options) { o.query["migration_name"] = migrationName })
}

// WithStartTime start_time获取
func (obj *_MigrationsMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取
func (obj *_MigrationsMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithBreakpoint breakpoint获取
func (obj *_MigrationsMgr) WithBreakpoint(breakpoint bool) Option {
	return optionFunc(func(o *options) { o.query["breakpoint"] = breakpoint })
}

// GetByOption 功能选项模式获取
func (obj *_MigrationsMgr) GetByOption(opts ...Option) (result Migrations, err error) {
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
func (obj *_MigrationsMgr) GetByOptions(opts ...Option) (results []*Migrations, err error) {
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

// GetFromVersion 通过version获取内容
func (obj *_MigrationsMgr) GetFromVersion(version int64) (result Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("version = ?", version).Find(&result).Error

	return
}

// GetBatchFromVersion 批量唯一主键查找
func (obj *_MigrationsMgr) GetBatchFromVersion(versions []int64) (results []*Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("version IN (?)", versions).Find(&results).Error

	return
}

// GetFromMigrationName 通过migration_name获取内容
func (obj *_MigrationsMgr) GetFromMigrationName(migrationName string) (results []*Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("migration_name = ?", migrationName).Find(&results).Error

	return
}

// GetBatchFromMigrationName 批量唯一主键查找
func (obj *_MigrationsMgr) GetBatchFromMigrationName(migrationNames []string) (results []*Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("migration_name IN (?)", migrationNames).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容
func (obj *_MigrationsMgr) GetFromStartTime(startTime time.Time) (results []*Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("start_time = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量唯一主键查找
func (obj *_MigrationsMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("start_time IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容
func (obj *_MigrationsMgr) GetFromEndTime(endTime time.Time) (results []*Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("end_time = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量唯一主键查找
func (obj *_MigrationsMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("end_time IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromBreakpoint 通过breakpoint获取内容
func (obj *_MigrationsMgr) GetFromBreakpoint(breakpoint bool) (results []*Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("breakpoint = ?", breakpoint).Find(&results).Error

	return
}

// GetBatchFromBreakpoint 批量唯一主键查找
func (obj *_MigrationsMgr) GetBatchFromBreakpoint(breakpoints []bool) (results []*Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("breakpoint IN (?)", breakpoints).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_MigrationsMgr) FetchByPrimaryKey(version int64) (result Migrations, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("version = ?", version).Find(&result).Error

	return
}
