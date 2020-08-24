package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _SystemQueueMgr struct {
	*_BaseMgr
}

// SystemQueueMgr open func
func SystemQueueMgr(db *gorm.DB) *_SystemQueueMgr {
	if db == nil {
		panic(fmt.Errorf("SystemQueueMgr need init by db"))
	}
	return &_SystemQueueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_queue"), isRelated: globalIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemQueueMgr) GetTableName() string {
	return "system_queue"
}

// Get 获取
func (obj *_SystemQueueMgr) Get() (result SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemQueueMgr) Gets() (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SystemQueueMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 任务编号
func (obj *_SystemQueueMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithTitle title获取 任务名称
func (obj *_SystemQueueMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithCommand command获取 执行指令
func (obj *_SystemQueueMgr) WithCommand(command string) Option {
	return optionFunc(func(o *options) { o.query["command"] = command })
}

// WithExecPid exec_pid获取 执行进程
func (obj *_SystemQueueMgr) WithExecPid(execPid int64) Option {
	return optionFunc(func(o *options) { o.query["exec_pid"] = execPid })
}

// WithExecData exec_data获取 执行参数
func (obj *_SystemQueueMgr) WithExecData(execData string) Option {
	return optionFunc(func(o *options) { o.query["exec_data"] = execData })
}

// WithExecTime exec_time获取 执行时间
func (obj *_SystemQueueMgr) WithExecTime(execTime int64) Option {
	return optionFunc(func(o *options) { o.query["exec_time"] = execTime })
}

// WithExecDesc exec_desc获取 执行描述
func (obj *_SystemQueueMgr) WithExecDesc(execDesc string) Option {
	return optionFunc(func(o *options) { o.query["exec_desc"] = execDesc })
}

// WithEnterTime enter_time获取 开始时间
func (obj *_SystemQueueMgr) WithEnterTime(enterTime float64) Option {
	return optionFunc(func(o *options) { o.query["enter_time"] = enterTime })
}

// WithOuterTime outer_time获取 结束时间
func (obj *_SystemQueueMgr) WithOuterTime(outerTime float64) Option {
	return optionFunc(func(o *options) { o.query["outer_time"] = outerTime })
}

// WithLoopsTime loops_time获取 循环时间
func (obj *_SystemQueueMgr) WithLoopsTime(loopsTime int64) Option {
	return optionFunc(func(o *options) { o.query["loops_time"] = loopsTime })
}

// WithAttempts attempts获取 执行次数
func (obj *_SystemQueueMgr) WithAttempts(attempts int64) Option {
	return optionFunc(func(o *options) { o.query["attempts"] = attempts })
}

// WithRscript rscript获取 任务类型(0单例,1多例)
func (obj *_SystemQueueMgr) WithRscript(rscript bool) Option {
	return optionFunc(func(o *options) { o.query["rscript"] = rscript })
}

// WithStatus status获取 任务状态(1新任务,2处理中,3成功,4失败)
func (obj *_SystemQueueMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取 创建时间
func (obj *_SystemQueueMgr) WithCreateAt(createAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// GetByOption 功能选项模式获取
func (obj *_SystemQueueMgr) GetByOption(opts ...Option) (result SystemQueue, err error) {
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
func (obj *_SystemQueueMgr) GetByOptions(opts ...Option) (results []*SystemQueue, err error) {
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
func (obj *_SystemQueueMgr) GetFromID(id int64) (result SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_SystemQueueMgr) GetBatchFromID(ids []int64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 任务编号
func (obj *_SystemQueueMgr) GetFromCode(code string) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("code = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量唯一主键查找 任务编号
func (obj *_SystemQueueMgr) GetBatchFromCode(codes []string) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("code IN (?)", codes).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 任务名称
func (obj *_SystemQueueMgr) GetFromTitle(title string) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找 任务名称
func (obj *_SystemQueueMgr) GetBatchFromTitle(titles []string) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromCommand 通过command获取内容 执行指令
func (obj *_SystemQueueMgr) GetFromCommand(command string) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("command = ?", command).Find(&results).Error

	return
}

// GetBatchFromCommand 批量唯一主键查找 执行指令
func (obj *_SystemQueueMgr) GetBatchFromCommand(commands []string) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("command IN (?)", commands).Find(&results).Error

	return
}

// GetFromExecPid 通过exec_pid获取内容 执行进程
func (obj *_SystemQueueMgr) GetFromExecPid(execPid int64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exec_pid = ?", execPid).Find(&results).Error

	return
}

// GetBatchFromExecPid 批量唯一主键查找 执行进程
func (obj *_SystemQueueMgr) GetBatchFromExecPid(execPids []int64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exec_pid IN (?)", execPids).Find(&results).Error

	return
}

// GetFromExecData 通过exec_data获取内容 执行参数
func (obj *_SystemQueueMgr) GetFromExecData(execData string) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exec_data = ?", execData).Find(&results).Error

	return
}

// GetBatchFromExecData 批量唯一主键查找 执行参数
func (obj *_SystemQueueMgr) GetBatchFromExecData(execDatas []string) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exec_data IN (?)", execDatas).Find(&results).Error

	return
}

// GetFromExecTime 通过exec_time获取内容 执行时间
func (obj *_SystemQueueMgr) GetFromExecTime(execTime int64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exec_time = ?", execTime).Find(&results).Error

	return
}

// GetBatchFromExecTime 批量唯一主键查找 执行时间
func (obj *_SystemQueueMgr) GetBatchFromExecTime(execTimes []int64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exec_time IN (?)", execTimes).Find(&results).Error

	return
}

// GetFromExecDesc 通过exec_desc获取内容 执行描述
func (obj *_SystemQueueMgr) GetFromExecDesc(execDesc string) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exec_desc = ?", execDesc).Find(&results).Error

	return
}

// GetBatchFromExecDesc 批量唯一主键查找 执行描述
func (obj *_SystemQueueMgr) GetBatchFromExecDesc(execDescs []string) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exec_desc IN (?)", execDescs).Find(&results).Error

	return
}

// GetFromEnterTime 通过enter_time获取内容 开始时间
func (obj *_SystemQueueMgr) GetFromEnterTime(enterTime float64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("enter_time = ?", enterTime).Find(&results).Error

	return
}

// GetBatchFromEnterTime 批量唯一主键查找 开始时间
func (obj *_SystemQueueMgr) GetBatchFromEnterTime(enterTimes []float64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("enter_time IN (?)", enterTimes).Find(&results).Error

	return
}

// GetFromOuterTime 通过outer_time获取内容 结束时间
func (obj *_SystemQueueMgr) GetFromOuterTime(outerTime float64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("outer_time = ?", outerTime).Find(&results).Error

	return
}

// GetBatchFromOuterTime 批量唯一主键查找 结束时间
func (obj *_SystemQueueMgr) GetBatchFromOuterTime(outerTimes []float64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("outer_time IN (?)", outerTimes).Find(&results).Error

	return
}

// GetFromLoopsTime 通过loops_time获取内容 循环时间
func (obj *_SystemQueueMgr) GetFromLoopsTime(loopsTime int64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loops_time = ?", loopsTime).Find(&results).Error

	return
}

// GetBatchFromLoopsTime 批量唯一主键查找 循环时间
func (obj *_SystemQueueMgr) GetBatchFromLoopsTime(loopsTimes []int64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("loops_time IN (?)", loopsTimes).Find(&results).Error

	return
}

// GetFromAttempts 通过attempts获取内容 执行次数
func (obj *_SystemQueueMgr) GetFromAttempts(attempts int64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attempts = ?", attempts).Find(&results).Error

	return
}

// GetBatchFromAttempts 批量唯一主键查找 执行次数
func (obj *_SystemQueueMgr) GetBatchFromAttempts(attemptss []int64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("attempts IN (?)", attemptss).Find(&results).Error

	return
}

// GetFromRscript 通过rscript获取内容 任务类型(0单例,1多例)
func (obj *_SystemQueueMgr) GetFromRscript(rscript bool) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("rscript = ?", rscript).Find(&results).Error

	return
}

// GetBatchFromRscript 批量唯一主键查找 任务类型(0单例,1多例)
func (obj *_SystemQueueMgr) GetBatchFromRscript(rscripts []bool) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("rscript IN (?)", rscripts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 任务状态(1新任务,2处理中,3成功,4失败)
func (obj *_SystemQueueMgr) GetFromStatus(status bool) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找 任务状态(1新任务,2处理中,3成功,4失败)
func (obj *_SystemQueueMgr) GetBatchFromStatus(statuss []bool) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_SystemQueueMgr) GetFromCreateAt(createAt time.Time) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量唯一主键查找 创建时间
func (obj *_SystemQueueMgr) GetBatchFromCreateAt(createAts []time.Time) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at IN (?)", createAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SystemQueueMgr) FetchByPrimaryKey(id int64) (result SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxSystemQueueCode  获取多个内容
func (obj *_SystemQueueMgr) FetchIndexByIDxSystemQueueCode(code string) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("code = ?", code).Find(&results).Error

	return
}

// FetchIndexByIDxSystemQueueTitle  获取多个内容
func (obj *_SystemQueueMgr) FetchIndexByIDxSystemQueueTitle(title string) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// FetchIndexByIDxSystemQueueExecTime  获取多个内容
func (obj *_SystemQueueMgr) FetchIndexByIDxSystemQueueExecTime(execTime int64) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("exec_time = ?", execTime).Find(&results).Error

	return
}

// FetchIndexByIDxSystemQueueRscript  获取多个内容
func (obj *_SystemQueueMgr) FetchIndexByIDxSystemQueueRscript(rscript bool) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("rscript = ?", rscript).Find(&results).Error

	return
}

// FetchIndexByIDxSystemQueueStatus  获取多个内容
func (obj *_SystemQueueMgr) FetchIndexByIDxSystemQueueStatus(status bool) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// FetchIndexByIDxSystemQueueCreateAt  获取多个内容
func (obj *_SystemQueueMgr) FetchIndexByIDxSystemQueueCreateAt(createAt time.Time) (results []*SystemQueue, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("create_at = ?", createAt).Find(&results).Error

	return
}
