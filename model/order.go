package model

import (
	"database/sql"
	"github.com/snowlyg/ChindeoTest/common"
	"time"
)

// APIAddrs [...]
type APIAddrs struct {
	ID           int          `gorm:"primary_key;column:id;type:int;not null"`
	Name         string       `gorm:"column:name;type:varchar(50);not null"`  // 联系人姓名
	Phone        string       `gorm:"column:phone;type:varchar(30);not null"` // 手机
	Addr         string       `gorm:"column:addr;type:varchar(200);not null"` // 地址
	Sex          int8         `gorm:"column:sex;type:tinyint;not null"`
	IsDefault    bool         `gorm:"column:is_default;type:tinyint(1);not null"` // 默认地址
	CreateAt     time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt     time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted    sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	UserID       int          `gorm:"index:user_id;column:user_id;type:int;not null"` // user_id
	HospitalName string       `gorm:"column:hospital_name;type:varchar(50);not null"` // 医院名称
	LocName      string       `gorm:"column:loc_name;type:varchar(50);not null"`      // 病区名称
	BedNum       string       `gorm:"column:bed_num;type:varchar(10);not null"`       // 床号
	HospitalNo   string       `gorm:"column:hospital_no;type:varchar(20);not null"`   // 住院号
	Disease      string       `gorm:"column:disease;type:varchar(150);not null"`      // 病种
	Age          int          `gorm:"column:age;type:int;not null"`                   // 年龄
}

// APIApplications [...]
type APIApplications struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null"`
	Name          string       `gorm:"column:name;type:varchar(50);not null"`
	AppID         string       `gorm:"column:app_id;type:varchar(30);not null"`      // app_id
	AppSecret     string       `gorm:"column:app_secret;type:varchar(100);not null"` // app_secret
	Tel           string       `gorm:"column:tel;type:varchar(100);not null"`
	Addr          string       `gorm:"column:addr;type:varchar(100);not null"`
	Describle     string       `gorm:"column:describle;type:varchar(100);not null"` // 应用描述
	AppType       string       `gorm:"column:app_type;type:varchar(50);not null"`   // app_type
	CreateAt      time.Time    `gorm:"column:create_at;type:timestamp;not null"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	UserID        int          `gorm:"column:user_id;type:int;not null"`
	CnareaID      int          `gorm:"column:cnarea_id;type:int;not null"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:timestamp;not null"`
	BusinessHours string       `gorm:"column:business_hours;type:varchar(50);not null"` // 营业时间
}

// APIAuths [...]
type APIAuths struct {
	ID         int          `gorm:"primary_key;column:id;type:int;not null"`
	Name       string       `gorm:"column:name;type:varchar(30);not null"`        // 权限名称
	Controller string       `gorm:"column:controller;type:varchar(100);not null"` // 控制器
	Action     string       `gorm:"column:action;type:varchar(100);not null"`     // 操作方法
	Method     string       `gorm:"column:method;type:varchar(100);not null"`     // 请求方法
	Desc       string       `gorm:"column:desc;type:varchar(30);not null"`        // 描述
	Remark     string       `gorm:"column:remark;type:varchar(100);not null"`     // 备注
	Status     bool         `gorm:"column:status;type:tinyint(1);not null"`       // 账号状态
	CreateAt   time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt   time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted  sql.NullTime `gorm:"column:is_deleted;type:datetime"`
}

// APIMenuTags [...]
type APIMenuTags struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null"`
	Name      string       `gorm:"column:name;type:varchar(50);not null"` // 标签名称
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null"`
}

// APIMenuTypes [...]
type APIMenuTypes struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null"`
	Name          string       `gorm:"column:name;type:varchar(50);not null"` // 菜品类型名称
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	ApplicationID int          `gorm:"column:application_id;type:int;not null"`
}

// APIMenus [...]
type APIMenus struct {
	ID            int                 `gorm:"primary_key;column:id;type:int;not null"`
	Name          string              `gorm:"column:name;type:varchar(50);not null"`    // 菜品名称
	TimeType      common.MenuTimeType `gorm:"column:time_type;type:int;not null"`       // 菜品时段类型
	Desc          string              `gorm:"column:desc;type:varchar(100);not null"`   // 菜品介绍
	Status        bool                `gorm:"column:status;type:tinyint(1);not null"`   // 菜品状态：上下架
	Amount        int                 `gorm:"column:amount;type:int;not null"`          // 销量
	Price         float64             `gorm:"column:price;type:decimal(10,2);not null"` // 价格
	Cover         string              `gorm:"column:cover;type:varchar(1200);not null"` // 封面
	Sort          int                 `gorm:"column:sort;type:int;not null"`            // 排序
	Pics          string              `gorm:"column:pics;type:varchar(1200);not null"`  // 餐品图片
	CreateAt      time.Time           `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt      time.Time           `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted     sql.NullTime        `gorm:"column:is_deleted;type:datetime"`
	ApplicationID int                 `gorm:"column:application_id;type:int;not null"`
	MenuTypeID    int                 `gorm:"index:menu_type_id;column:menu_type_id;type:int;not null"` // menu_type_id
}

// APIOOrderAddrs [...]
type APIOOrderAddrs struct {
	ID           int          `gorm:"primary_key;column:id;type:int;not null"`
	Name         string       `gorm:"column:name;type:varchar(50);not null"`  // 联系人姓名
	Phone        string       `gorm:"column:phone;type:varchar(30);not null"` // 手机
	Addr         string       `gorm:"column:addr;type:varchar(200);not null"` // 地址
	Sex          int          `gorm:"column:sex;type:int;not null"`           // 性别:1:男,0:女
	CreateAt     time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt     time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted    sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	OOrderID     int          `gorm:"index:o_order_id;column:o_order_id;type:int;not null"` // o_order_id
	HospitalName string       `gorm:"column:hospital_name;type:varchar(50);not null"`       // 医院名称
	LocName      string       `gorm:"column:loc_name;type:varchar(50);not null"`            // 病区名称
	BedNum       string       `gorm:"column:bed_num;type:varchar(10);not null"`             // 床号
	HospitalNo   string       `gorm:"column:hospital_no;type:varchar(20);not null"`         // 住院号
	Disease      string       `gorm:"column:disease;type:varchar(150);not null"`            // 病种
	Age          int          `gorm:"column:age;type:int;not null"`                         // 年龄
}

// APIOOrderMenus [...]
type APIOOrderMenus struct {
	ID           int          `gorm:"primary_key;column:id;type:int;not null"`
	MenuName     string       `gorm:"column:menu_name;type:varchar(50);not null"`        // 菜品名称
	MenuType     string       `gorm:"column:menu_type;type:varchar(30);not null"`        // 菜品类型
	MenuTimeType string       `gorm:"column:menu_time_type;type:varchar(30);not null"`   // 菜品时段类型
	MenuDesc     string       `gorm:"column:menu_desc;type:varchar(100);not null"`       // 菜品介绍
	MenuID       int          `gorm:"index:o_order_id;column:menu_id;type:int;not null"` // menu_id
	Price        float64      `gorm:"column:price;type:decimal(10,0);not null"`          // 价格
	Amount       int          `gorm:"column:amount;type:int;not null"`                   // 数量
	Cover        string       `gorm:"column:cover;type:varchar(255);not null"`           // 菜品图片
	CreateAt     time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt     time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted    sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	OOrderID     int          `gorm:"index:o_order_id;column:o_order_id;type:int;not null"` // o_order_id
}

// APIOOrders [...]
type APIOOrders struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null"`
	OrderNo       string       `gorm:"index:user_id;column:order_no;type:varchar(50);not null"` // 订单号
	Status        int          `gorm:"column:status;type:int;not null"`                         // 订单状态：1：待付款，2：已付款，3，配送中，4：已取消，5：已完成
	Amount        int          `gorm:"column:amount;type:int;not null"`                         // 总数量
	Total         float64      `gorm:"column:total;type:decimal(10,2);not null"`                // 总金额
	PayType       int          `gorm:"column:pay_type;type:tinyint(1);not null"`
	Rmk           string       `gorm:"column:rmk;type:varchar(500)"`                // 备注
	AppType       int          `gorm:"column:app_type;type:int;not null"`           // 订单应用类型，1：小程序，2：床旁
	OpenID        string       `gorm:"column:open_id;type:varchar(100);not null"`   // open_id
	TradeType     string       `gorm:"column:trade_type;type:varchar(10);not null"` // JSAPI、NATIVE、APP
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	IsReturn      bool         `gorm:"column:is_return;type:tinyint(1);not null"`             // 是否退款
	ApplicationID int          `gorm:"index:user_id;column:application_id;type:int;not null"` // application_id
	UserID        int          `gorm:"index:user_id;column:user_id;type:int;not null"`        // user_id
}

// APIOReturnOrderAddrs [...]
type APIOReturnOrderAddrs struct {
	ID             int       `gorm:"primary_key;column:id;type:int;not null"`
	Name           string    `gorm:"column:name;type:varchar(50);not null"`  // 联系人姓名
	Phone          string    `gorm:"column:phone;type:varchar(30);not null"` // 手机
	Addr           string    `gorm:"column:addr;type:varchar(200);not null"` // 地址
	Sex            int       `gorm:"column:sex;type:int;not null"`           // 性别:1:男,0:女
	CreateAt       time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt       time.Time `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted      time.Time `gorm:"column:is_deleted;type:datetime"`
	ApplicationID  int       `gorm:"column:application_id;type:int;not null"`                            // application_id
	OReturnOrderID int       `gorm:"index:o_return_order_id;column:o_return_order_id;type:int;not null"` // o_return_order_id
	HospitalName   string    `gorm:"column:hospital_name;type:varchar(50);not null"`                     // 医院名称
	LocName        string    `gorm:"column:loc_name;type:varchar(50);not null"`                          // 病区名称
	BedNum         string    `gorm:"column:bed_num;type:varchar(10);not null"`                           // 床号
	HospitalNo     string    `gorm:"column:hospital_no;type:varchar(20);not null"`                       // 住院号
	Disease        string    `gorm:"column:disease;type:varchar(150);not null"`                          // 病种
	Age            int       `gorm:"column:age;type:int;not null"`                                       // 年龄
}

// APIOReturnOrderMenus [...]
type APIOReturnOrderMenus struct {
	ID             int       `gorm:"primary_key;column:id;type:int;not null"`
	MenuName       string    `gorm:"column:menu_name;type:varchar(50);not null"`               // 菜品名称
	MenuType       string    `gorm:"column:menu_type;type:varchar(30);not null"`               // 菜品类型
	MenuTimeType   string    `gorm:"column:menu_time_type;type:varchar(30);not null"`          // 菜品时段类型
	MenuDesc       string    `gorm:"column:menu_desc;type:varchar(100);not null"`              // 菜品介绍
	MenuID         int       `gorm:"index:o_return_order_id;column:menu_id;type:int;not null"` // menu_id
	Price          float64   `gorm:"column:price;type:decimal(10,0);not null"`                 // 价格
	Amount         int       `gorm:"column:amount;type:int;not null"`                          // 数量
	Cover          string    `gorm:"column:cover;type:varchar(255);not null"`                  // 菜品图片
	CreateAt       time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt       time.Time `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted      time.Time `gorm:"column:is_deleted;type:datetime"`
	ApplicationID  int       `gorm:"column:application_id;type:int;not null"`                            // application_id
	OReturnOrderID int       `gorm:"index:o_return_order_id;column:o_return_order_id;type:int;not null"` // o_return_order_id
}

// APIOReturnOrders [...]
type APIOReturnOrders struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null"`
	OrderNo       string       `gorm:"index:o_order_id;column:order_no;type:varchar(50);not null"` // 订单号
	Status        int          `gorm:"column:status;type:int;not null"`                            // 订单状态：1，待审核，2：审核通过，3：审核驳回，4：退款成功，5：退款失败，
	Amount        int          `gorm:"column:amount;type:int;not null"`                            // 总数量
	Total         float64      `gorm:"column:total;type:decimal(10,0);not null"`                   // 总金额
	OpenID        string       `gorm:"column:open_id;type:varchar(100);not null"`                  // open_id
	AppType       int          `gorm:"column:app_type;type:int;not null"`                          // 订单应用类型，1：小程序，2：床旁
	TradeType     string       `gorm:"column:trade_type;type:varchar(10);not null"`                // JSAPI、NATIVE、APP
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	OOrderID      int          `gorm:"index:o_order_id;column:o_order_id;type:int;not null"`     // o_order_id
	ApplicationID int          `gorm:"index:o_order_id;column:application_id;type:int;not null"` // application_id
	PayType       bool         `gorm:"column:pay_type;type:tinyint(1);not null"`
	UserID        int          `gorm:"column:user_id;type:int;not null"` // user_id
}

// APIProfiles [...]
type APIProfiles struct {
	ID           int          `gorm:"primary_key;column:id;type:int;not null"`
	Name         string       `gorm:"column:name;type:varchar(50);not null"`          // 联系人姓名
	Phone        string       `gorm:"column:phone;type:varchar(30);not null"`         // 手机
	Sex          int          `gorm:"column:sex;type:int;not null"`                   // 性别:1:男,0:女
	HospitalName string       `gorm:"column:hospital_name;type:varchar(50);not null"` // 医院名称
	LocName      string       `gorm:"column:loc_name;type:varchar(50);not null"`      // 病区名称
	BedNum       string       `gorm:"column:bed_num;type:varchar(10);not null"`       // 床号
	HospitalNo   string       `gorm:"column:hospital_no;type:varchar(20);not null"`   // 住院号
	Disease      string       `gorm:"column:disease;type:varchar(150);not null"`      // 病种
	Age          int          `gorm:"column:age;type:int;not null"`                   // 年龄
	CreateAt     time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt     time.Time    `gorm:"column:update_at;type:datetime;not null"`
	UserID       int          `gorm:"column:user_id;type:int;not null"`
	IsDeleted    sql.NullTime `gorm:"column:is_deleted;type:datetime"`
}

// APIRoles [...]
type APIRoles struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null"`
	Name      string       `gorm:"column:name;type:varchar(30);not null"`    // 角色名称
	Desc      string       `gorm:"column:desc;type:varchar(30);not null"`    // 描述
	Remark    string       `gorm:"column:remark;type:varchar(100);not null"` // 备注
	Status    bool         `gorm:"column:status;type:tinyint(1);not null"`   // 账号状态
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime"`
}

// APIUsers [...]
type APIUsers struct {
	ID        int       `gorm:"primary_key;column:id;type:int;not null"`
	Username  string    `gorm:"unique_index:open_id;column:username;type:varchar(30);not null"`  // 用户名
	Nickname  string    `gorm:"column:nickname;type:varchar(30);not null"`                       // 名称
	Phone     string    `gorm:"unique_index:open_id;column:phone;type:varchar(30);not null"`     // 手机
	Email     string    `gorm:"unique_index:open_id;column:email;type:varchar(50);not null"`     // 邮箱
	Sex       int       `gorm:"column:sex;type:int;not null"`                                    // 性别
	Password  string    `gorm:"column:password;type:varchar(32);not null"`                       // 用户密码
	Status    bool      `gorm:"column:status;type:tinyint(1);not null"`                          // 账号状态
	AvatarURL string    `gorm:"column:avatar_url;type:varchar(500);not null"`                    // 用户头像
	OpenID    string    `gorm:"unique_index:open_id;column:open_id;type:varchar(100);not null"`  // open_id
	UnionID   string    `gorm:"unique_index:open_id;column:union_id;type:varchar(100);not null"` // union_id
	Country   string    `gorm:"column:country;type:varchar(100);not null"`                       // 国家
	Province  string    `gorm:"column:province;type:varchar(100);not null"`                      // 省份
	City      string    `gorm:"column:city;type:varchar(100);not null"`                          // 城镇
	Mac       string    `gorm:"column:mac;type:varchar(100);not null"`                           // mac 地址
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;not null"`
	IDCardNo  string    `gorm:"column:id_card_no;type:varchar(20);not null"` // 身份证
	IsAuth    bool      `gorm:"column:is_auth;type:tinyint(1);not null"`     // 是否实名认证
	Realname  string    `gorm:"column:realname;type:varchar(30);not null"`   // 真实姓名
}

// APIWarnTimes [...]
type APIWarnTimes struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null"`
	Weeks     string       `gorm:"column:weeks;type:varchar(30);not null"` // 星期
	Time      string       `gorm:"column:time;type:varchar(20);not null"`  // 提醒时间
	Status    bool         `gorm:"column:status;type:tinyint(1);not null"` // 是否启用
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	UserID    int          `gorm:"index:user_id;column:user_id;type:int;not null"` // user_id
}

// CareCareTag 护理标签和护理关系中间表
type CareCareTag struct {
	ID        int       `gorm:"primary_key;column:id;type:int;not null"`
	CareID    int       `gorm:"index:care_id;column:care_id;type:int;not null"`     // care_id
	CareTagID int       `gorm:"index:care_id;column:care_tag_id;type:int;not null"` // care_tag_id
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;not null"`
}

// CareOrderAddrs 订单地址表
type CareOrderAddrs struct {
	ID           int          `gorm:"primary_key;column:id;type:int;not null"`
	Name         string       `gorm:"column:name;type:varchar(50);not null"`          // 联系人姓名
	Phone        string       `gorm:"column:phone;type:varchar(30);not null"`         // 手机
	Addr         string       `gorm:"column:addr;type:varchar(200);not null"`         // 地址
	Sex          int          `gorm:"column:sex;type:int;not null"`                   // 性别:1:男,0:女
	HospitalName string       `gorm:"column:hospital_name;type:varchar(50);not null"` // 医院名称
	LocName      string       `gorm:"column:loc_name;type:varchar(50);not null"`      // 病区名称
	BedNum       string       `gorm:"column:bed_num;type:varchar(10);not null"`       // 床号
	HospitalNo   string       `gorm:"column:hospital_no;type:varchar(20);not null"`   // 住院号
	Disease      string       `gorm:"column:disease;type:varchar(150);not null"`      // 病种
	IsDeleted    sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	CareOrderID  int          `gorm:"index:care_order_id;column:care_order_id;type:int;not null"` // care_order_id
	CreateAt     time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt     time.Time    `gorm:"column:update_at;type:datetime;not null"`
	Age          int          `gorm:"column:age;type:int;not null"` // 年龄
}

// CareOrderCarerInfos 订单护工信息表
type CareOrderCarerInfos struct {
	ID              int       `gorm:"primary_key;column:id;type:int;not null"`
	Name            string    `gorm:"column:name;type:varchar(50);not null"`             // 姓名
	Age             int       `gorm:"column:age;type:int;not null"`                      // 年龄
	WorkExp         int       `gorm:"column:work_exp;type:int;not null"`                 // 工作年限
	Price           float64   `gorm:"column:price;type:decimal(10,2);not null"`          // 价格
	Sex             int       `gorm:"column:sex;type:int;not null"`                      // 性别:1:男,0:女
	Phone           string    `gorm:"column:phone;type:varchar(30);not null"`            // 手机
	Desc            string    `gorm:"column:desc;type:varchar(500);not null"`            // 简介
	Avatar          string    `gorm:"column:avatar;type:varchar(200);not null"`          // 头像
	CarerTags       string    `gorm:"column:carer_tags;type:varchar(200);not null"`      // 标签
	ApplicationName string    `gorm:"column:application_name;type:varchar(50);not null"` // 医院名称
	CreateAt        time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt        time.Time `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted       time.Time `gorm:"column:is_deleted;type:datetime"`
	CareOrderID     int       `gorm:"index:care_order_id;column:care_order_id;type:int;not null"` // care_order_id
	TimeType        string    `gorm:"column:time_type;type:varchar(5);not null"`                  // 时间类型：时，天，月
}

// CareOrderInfos 护理项目表
type CareOrderInfos struct {
	ID              int       `gorm:"primary_key;column:id;type:int;not null"`
	Name            string    `gorm:"column:name;type:varchar(50);not null"`             // 名称
	Desc            string    `gorm:"column:desc;type:varchar(100);not null"`            // 简介
	TimeType        string    `gorm:"column:time_type;type:varchar(5);not null"`         // 时间类型：时，天，月
	CareType        string    `gorm:"column:care_type;type:varchar(30);not null"`        // 类型
	CareTags        string    `gorm:"column:care_tags;type:varchar(80);not null"`        // 标签：用 | 分隔
	MinPrice        float64   `gorm:"column:min_price;type:decimal(10,2);not null"`      // 最小价格
	MaxPrice        float64   `gorm:"column:max_price;type:decimal(10,2);not null"`      // 最大价格
	Cover           string    `gorm:"column:cover;type:varchar(1200);not null"`          // 封面
	CareDetail      string    `gorm:"column:care_detail;type:varchar(2000);not null"`    // 服务内容:采用富文本
	ApplicationName string    `gorm:"column:application_name;type:varchar(50);not null"` // 医院名称
	CreateAt        time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt        time.Time `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted       time.Time `gorm:"column:is_deleted;type:datetime"`
	CareOrderID     int       `gorm:"index:care_order_id;column:care_order_id;type:int;not null"` // care_order_id
}

// CareOrders 护理订单表
type CareOrders struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null"`
	OrderNo       string       `gorm:"column:order_no;type:varchar(50);not null"`             // 订单号
	Status        int          `gorm:"column:status;type:int;not null"`                       // 订单状态：1：待付款，2：已付款，3，进行中，4：已取消，5：已完成
	PayType       int          `gorm:"column:pay_type;type:int;not null"`                     // 订单支付类型：1：微信，2：支付宝
	Total         float64      `gorm:"column:total;type:decimal(10,2);not null"`              // 总金额
	Rmk           string       `gorm:"column:rmk;type:varchar(500);not null"`                 // 备注
	StartAt       time.Time    `gorm:"column:start_at;type:datetime;not null"`                // 服务开始时间
	EndAt         time.Time    `gorm:"column:end_at;type:datetime;not null"`                  // 服务结束时间
	OpenID        string       `gorm:"column:open_id;type:varchar(100);not null"`             // open_id
	TradeType     string       `gorm:"column:trade_type;type:varchar(10);not null"`           // JSAPI、NATIVE、APP
	IsReturn      bool         `gorm:"column:is_return;type:tinyint(1);not null"`             // 是否退款
	ApplicationID int          `gorm:"index:user_id;column:application_id;type:int;not null"` // application_id
	UserID        int          `gorm:"index:user_id;column:user_id;type:int;not null"`        // user_id
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null"`
	AppType       int          `gorm:"column:app_type;type:int;not null"`                // 订单应用类型，1：小程序，2：床旁
	CarerID       int          `gorm:"index:carer_id;column:carer_id;type:int;not null"` // 护工id
}

// CareReturnOrderAddrs 退款订单地址表
type CareReturnOrderAddrs struct {
	ID                int       `gorm:"primary_key;column:id;type:int;not null"`
	Name              string    `gorm:"column:name;type:varchar(50);not null"`          // 联系人姓名
	Phone             string    `gorm:"column:phone;type:varchar(30);not null"`         // 手机
	Addr              string    `gorm:"column:addr;type:varchar(200);not null"`         // 地址
	Sex               int       `gorm:"column:sex;type:int;not null"`                   // 性别:1:男,0:女
	HospitalName      string    `gorm:"column:hospital_name;type:varchar(50);not null"` // 医院名称
	LocName           string    `gorm:"column:loc_name;type:varchar(50);not null"`      // 病区名称
	BedNum            string    `gorm:"column:bed_num;type:varchar(10);not null"`       // 床号
	HospitalNo        string    `gorm:"column:hospital_no;type:varchar(20);not null"`   // 住院号
	Disease           string    `gorm:"column:disease;type:varchar(150);not null"`      // 病种
	CreateAt          time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt          time.Time `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted         time.Time `gorm:"column:is_deleted;type:datetime"`
	ApplicationID     int       `gorm:"column:application_id;type:int;not null"`                                  // application_id
	CareReturnOrderID int       `gorm:"index:care_return_order_id;column:care_return_order_id;type:int;not null"` // care_return_order_id
	Age               int       `gorm:"column:age;type:int;not null"`                                             // 年龄
}

// CareReturnOrderCarerInfos 订单护工信息表
type CareReturnOrderCarerInfos struct {
	ID                int       `gorm:"primary_key;column:id;type:int;not null"`
	Name              string    `gorm:"column:name;type:varchar(50);not null"`             // 姓名
	Age               int       `gorm:"column:age;type:int;not null"`                      // 年龄
	WorkExp           int       `gorm:"column:work_exp;type:int;not null"`                 // 工作年限
	Price             float64   `gorm:"column:price;type:decimal(10,2);not null"`          // 价格
	Sex               int       `gorm:"column:sex;type:int;not null"`                      // 性别:1:男,0:女
	Phone             string    `gorm:"column:phone;type:varchar(30);not null"`            // 手机
	Desc              string    `gorm:"column:desc;type:varchar(500);not null"`            // 简介
	Avatar            string    `gorm:"column:avatar;type:varchar(200);not null"`          // 头像
	CarerTags         string    `gorm:"column:carer_tags;type:varchar(200);not null"`      // 标签
	ApplicationName   string    `gorm:"column:application_name;type:varchar(50);not null"` // 医院名称
	CreateAt          time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt          time.Time `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted         time.Time `gorm:"column:is_deleted;type:datetime"`
	CareReturnOrderID int       `gorm:"index:care_return_order_id;column:care_return_order_id;type:int;not null"` // care_return_order_id
	TimeType          string    `gorm:"column:time_type;type:varchar(5);not null"`                                // 时间类型：时，天，月
}

// CareReturnOrderInfos 护理项目表
type CareReturnOrderInfos struct {
	ID                int       `gorm:"primary_key;column:id;type:int;not null"`
	Name              string    `gorm:"column:name;type:varchar(50);not null"`             // 名称
	Desc              string    `gorm:"column:desc;type:varchar(100);not null"`            // 简介
	TimeType          string    `gorm:"column:time_type;type:varchar(5);not null"`         // 时间类型：时，天，月
	CareType          string    `gorm:"column:care_type;type:varchar(30);not null"`        // 类型
	CareTags          string    `gorm:"column:care_tags;type:varchar(80);not null"`        // 标签：用 | 分隔
	MinPrice          float64   `gorm:"column:min_price;type:decimal(10,2);not null"`      // 最小价格
	MaxPrice          float64   `gorm:"column:max_price;type:decimal(10,2);not null"`      // 最大价格
	Cover             string    `gorm:"column:cover;type:varchar(1200);not null"`          // 封面
	CareDetail        string    `gorm:"column:care_detail;type:varchar(2000);not null"`    // 服务内容:采用富文本
	ApplicationName   string    `gorm:"column:application_name;type:varchar(50);not null"` // 医院名称
	CreateAt          time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt          time.Time `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted         time.Time `gorm:"column:is_deleted;type:datetime"`
	CareReturnOrderID int       `gorm:"index:care_return_order_id;column:care_return_order_id;type:int;not null"` // care_return_order_id
}

// CareReturnOrders 退款订单表
type CareReturnOrders struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null"`
	OrderNo       string       `gorm:"column:order_no;type:varchar(50);not null"`   // 订单号
	Status        int          `gorm:"column:status;type:int;not null"`             // 订单状态：1，待审核，2：审核通过，3：审核驳回，4：退款成功，5：退款失败，
	PayType       int          `gorm:"column:pay_type;type:int;not null"`           // 订单支付类型：1：微信，2：支付宝
	Total         float64      `gorm:"column:total;type:decimal(10,2);not null"`    // 总金额
	OpenID        string       `gorm:"column:open_id;type:varchar(100);not null"`   // open_id
	TradeType     string       `gorm:"column:trade_type;type:varchar(10);not null"` // JSAPI、NATIVE、APP
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	CareOrderID   int          `gorm:"index:care_order_id;column:care_order_id;type:int;not null"`  // care_order_id
	ApplicationID int          `gorm:"index:care_order_id;column:application_id;type:int;not null"` // application_id
	UserID        int          `gorm:"column:user_id;type:int;not null"`                            // user_id
	AppType       int          `gorm:"column:app_type;type:int;not null"`                           // 订单应用类型，1：小程序，2：床旁
	CarerID       int          `gorm:"index:carer_id;column:carer_id;type:int;not null"`            // 护工id
}

// CareTags 护理标签表
type CareTags struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null"`
	Name      string       `gorm:"column:name;type:varchar(50);not null"` // 名称
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null"`
	Icon      string       `gorm:"column:icon;type:varchar(200);not null"` // 图标
}

// CareTypes 护理类型表
type CareTypes struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null"`
	Name      string       `gorm:"column:name;type:varchar(50);not null"`    // 名称
	EnName    string       `gorm:"column:en_name;type:varchar(50);not null"` // 英文名称
	Status    bool         `gorm:"column:status;type:tinyint(1);not null"`   // 状态：启用，禁用
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	Icon      string       `gorm:"column:icon;type:varchar(200);not null"` // 图标
}

// CarerCarerTag 护工标签和护工关系中间表
type CarerCarerTag struct {
	ID         int       `gorm:"primary_key;column:id;type:int;not null"`
	CarerID    int       `gorm:"index:carer_id;column:carer_id;type:int;not null"`     // carer_id
	CarerTagID int       `gorm:"index:carer_id;column:carer_tag_id;type:int;not null"` // carer_tag_id
	CreateAt   time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt   time.Time `gorm:"column:update_at;type:datetime;not null"`
}

// CarerServerTimes 护工服务时间表
type CarerServerTimes struct {
	ID          int       `gorm:"primary_key;column:id;type:int;not null"`
	StartAt     time.Time `gorm:"column:start_at;type:datetime;not null"`                     // 服务开始时间
	EndAt       time.Time `gorm:"column:end_at;type:datetime;not null"`                       // 服务结束时间
	CarerID     int       `gorm:"index:care_order_id;column:carer_id;type:int;not null"`      // carer_id
	TimeType    string    `gorm:"column:time_type;type:varchar(5);not null"`                  // 时间类型：时，天
	CareOrderID int       `gorm:"index:care_order_id;column:care_order_id;type:int;not null"` // care_order_id
	CreateAt    time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt    time.Time `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted   time.Time `gorm:"column:is_deleted;type:datetime"`
}

// CarerTags 护工标签表
type CarerTags struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null"`
	Name      string       `gorm:"column:name;type:varchar(50);not null"` // 名称
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	Icon      string       `gorm:"column:icon;type:varchar(200);not null"` // 图标
}

// Carers 护工表
type Carers struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null"`
	Name          string       `gorm:"column:name;type:varchar(50);not null"`    // 姓名
	Age           int          `gorm:"column:age;type:int;not null"`             // 年龄
	WorkExp       int          `gorm:"column:work_exp;type:int;not null"`        // 工作年限
	Price         float64      `gorm:"column:price;type:decimal(10,2);not null"` // 价格
	Sex           int          `gorm:"column:sex;type:int;not null"`             // 性别:1:男,0:女
	Phone         string       `gorm:"column:phone;type:varchar(30);not null"`   // 手机
	Desc          string       `gorm:"column:desc;type:varchar(500);not null"`   // 简介
	Avatar        string       `gorm:"column:avatar;type:varchar(200);not null"` // 头像
	Amount        int          `gorm:"column:amount;type:int;not null"`          // 服务数量
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	ApplicationID int          `gorm:"index:application_id;column:application_id;type:int;not null"` // application_id
	Status        bool         `gorm:"column:status;type:tinyint(1);not null"`                       // 状态：上下架
	TimeType      string       `gorm:"column:time_type;type:varchar(5);not null"`                    // 时间类型：时，天，月
	CarerDetail   string       `gorm:"column:carer_detail;type:varchar(2000);not null"`              // 服务内容:采用富文本
	CardNo        string       `gorm:"column:card_no;type:varchar(12);not null"`                     // 工号
}

// Cares 护理项目表
type Cares struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null"`
	Name          string       `gorm:"column:name;type:varchar(50);not null"`          // 名称
	Desc          string       `gorm:"column:desc;type:varchar(100);not null"`         // 简介
	TimeType      string       `gorm:"column:time_type;type:varchar(5);not null"`      // 时间类型：时，天，月
	Status        bool         `gorm:"column:status;type:tinyint(1);not null"`         // 状态：上下架
	Amount        int          `gorm:"column:amount;type:int;not null"`                // 销量
	MinPrice      float64      `gorm:"column:min_price;type:decimal(10,2);not null"`   // 最小价格
	MaxPrice      float64      `gorm:"column:max_price;type:decimal(10,2);not null"`   // 最大价格
	Cover         string       `gorm:"column:cover;type:varchar(1200);not null"`       // 封面
	Sort          int          `gorm:"column:sort;type:int;not null"`                  // 排序
	CareDetail    string       `gorm:"column:care_detail;type:varchar(2000);not null"` // 服务内容:采用富文本
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime"`
	CareTypeID    int          `gorm:"index:care_type_id;column:care_type_id;type:int;not null"`   // care_type_id
	ApplicationID int          `gorm:"index:care_type_id;column:application_id;type:int;not null"` // application_id
}

// Cnareas 中国行政地区表
type Cnareas struct {
	ID         string  `gorm:"primary_key;column:id;type:mediumint unsigned;not null"`
	Level      uint8   `gorm:"column:level;type:tinyint unsigned;not null"`                            // 层级
	ParentCode uint64  `gorm:"index:idx_parent_code;column:parent_code;type:bigint unsigned;not null"` // 父级行政代码
	AreaCode   uint64  `gorm:"unique;column:area_code;type:bigint unsigned;not null"`                  // 行政代码
	ZipCode    string  `gorm:"column:zip_code;type:mediumint(6) unsigned zerofill;not null"`           // 邮政编码
	CityCode   string  `gorm:"column:city_code;type:char(6);not null"`                                 // 区号
	Name       string  `gorm:"column:name;type:varchar(50);not null"`                                  // 名称
	ShortName  string  `gorm:"column:short_name;type:varchar(50);not null"`                            // 简称
	MergerName string  `gorm:"column:merger_name;type:varchar(50);not null"`                           // 组合名
	Pinyin     string  `gorm:"column:pinyin;type:varchar(30);not null"`                                // 拼音
	Lng        float64 `gorm:"column:lng;type:decimal(10,6);not null"`                                 // 经度
	Lat        float64 `gorm:"column:lat;type:decimal(10,6);not null"`                                 // 纬度
}

// MenuMenuTag [...]
type MenuMenuTag struct {
	ID        int       `gorm:"primary_key;column:id;type:int;not null"`
	MenuID    int       `gorm:"index:menu_id;column:menu_id;type:int;not null"`     // menu_id
	MenuTagID int       `gorm:"index:menu_id;column:menu_tag_id;type:int;not null"` // menu_tag_id
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;not null"`
}

// Migrations [...]
type Migrations struct {
	Version       int64     `gorm:"primary_key;column:version;type:bigint;not null"`
	MigrationName string    `gorm:"column:migration_name;type:varchar(100)"`
	StartTime     time.Time `gorm:"column:start_time;type:timestamp;not null"`
	EndTime       time.Time `gorm:"column:end_time;type:timestamp;not null"`
	Breakpoint    bool      `gorm:"column:breakpoint;type:tinyint(1);not null"`
}

// MiniApps [...]
type MiniApps struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null"`
	AppID     string       `gorm:"unique_index:uuid;column:app_id;type:varchar(100);not null"` // app_id
	AppSecret string       `gorm:"column:app_secret;type:varchar(100);not null"`               // app_secret
	UUID      string       `gorm:"unique_index:uuid;column:uuid;type:varchar(100);not null"`   // uuid
	Name      string       `gorm:"index:name;column:name;type:varchar(50);not null"`           // 应用名称
	Describle string       `gorm:"column:describle;type:varchar(100);not null"`                // 应用描述
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime"`
}

// Rules [...]
type Rules struct {
	ID    int    `gorm:"primary_key;column:id;type:int;not null"`
	Ptype string `gorm:"column:ptype;type:varchar(255)"`
	V0    string `gorm:"column:v0;type:varchar(255)"`
	V1    string `gorm:"column:v1;type:varchar(255)"`
	V2    string `gorm:"column:v2;type:varchar(255)"`
	V3    string `gorm:"column:v3;type:varchar(255)"`
	V4    string `gorm:"column:v4;type:varchar(255)"`
	V5    string `gorm:"column:v5;type:varchar(255)"`
}

// StateRecords 对账记录表
type StateRecords struct {
	ID           int       `gorm:"primary_key;column:id;type:int;not null"`
	StateType    int       `gorm:"column:state_type;type:int;not null"`           // 对账类型：1：微信，2：支付宝
	OrderType    int       `gorm:"column:order_type;type:int;not null"`           // 订单类型：1：付款，2：退款
	OrderNum     int       `gorm:"column:order_num;type:int;not null"`            // 订单数量
	Total        float64   `gorm:"column:total;type:decimal(10,2);not null"`      // 总金额
	RealOrderNum int       `gorm:"column:real_order_num;type:int;not null"`       // 支付平台订单数量
	RealTotal    float64   `gorm:"column:real_total;type:decimal(10,2);not null"` // 支付平台总金额
	IsException  bool      `gorm:"column:is_exception;type:tinyint(1);not null"`  // 是否异常
	OrderNos     string    `gorm:"column:order_nos;type:varchar(800);not null"`   // 异常订单号用 ， 隔开
	ExtraType    int       `gorm:"column:extra_type;type:int;not null"`           // 平账方式：1：盈余，2：亏损
	Extra        float64   `gorm:"column:extra;type:decimal(10,2);not null"`      // 平账金额
	StartAt      time.Time `gorm:"column:start_at;type:datetime;not null"`        // 对账订单开始时间范围
	EndAt        time.Time `gorm:"column:end_at;type:datetime;not null"`          // 对账订单结束时间范围
	Username     string    `gorm:"column:username;type:varchar(50);not null"`     // 平账人
	Rmk          string    `gorm:"column:rmk;type:varchar(500);not null"`         // 平账备注
	RecordAt     time.Time `gorm:"column:record_at;type:datetime;not null"`       // 对账时间
	CreateAt     time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt     time.Time `gorm:"column:update_at;type:datetime;not null"`
	IsDeleted    bool      `gorm:"column:is_deleted;type:tinyint(1)"`
}

// SystemAuth 系统-权限
type SystemAuth struct {
	ID       uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null"`
	Title    string    `gorm:"index:idx_system_auth_title;column:title;type:varchar(100)"`       // 权限名称
	Desc     string    `gorm:"column:desc;type:varchar(500)"`                                    // 备注说明
	Sort     uint64    `gorm:"column:sort;type:bigint unsigned"`                                 // 排序权重
	Status   uint8     `gorm:"index:idx_system_auth_status;column:status;type:tinyint unsigned"` // 权限状态(1使用,0禁用)
	CreateAt time.Time `gorm:"column:create_at;type:timestamp"`                                  // 创建时间
}

// SystemAuthNode 系统-授权
type SystemAuthNode struct {
	ID   uint64 `gorm:"primary_key;column:id;type:bigint unsigned;not null"`
	Auth uint64 `gorm:"index:idx_system_auth_auth;column:auth;type:bigint unsigned"` // 角色
	Node string `gorm:"index:idx_system_auth_node;column:node;type:varchar(200)"`    // 节点
}

// SystemConfig 系统-配置
type SystemConfig struct {
	Type  string `gorm:"index:idx_system_config_type;column:type;type:varchar(20)"`  // 分类
	Name  string `gorm:"index:idx_system_config_name;column:name;type:varchar(100)"` // 配置名
	Value string `gorm:"column:value;type:varchar(2048)"`                            // 配置值
}

// SystemData 系统-数据
type SystemData struct {
	ID    uint64 `gorm:"primary_key;column:id;type:bigint unsigned;not null"`
	Name  string `gorm:"index:idx_system_data_name;column:name;type:varchar(100)"` // 配置名
	Value string `gorm:"column:value;type:longtext"`                               // 配置值
}

// SystemMenu 系统-菜单
type SystemMenu struct {
	ID       uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null"`
	Pid      uint64    `gorm:"column:pid;type:bigint unsigned"`                                  // 上级ID
	Title    string    `gorm:"column:title;type:varchar(100)"`                                   // 菜单名称
	Icon     string    `gorm:"column:icon;type:varchar(100)"`                                    // 菜单图标
	Node     string    `gorm:"index:idx_system_menu_node;column:node;type:varchar(100)"`         // 节点代码
	URL      string    `gorm:"column:url;type:varchar(400)"`                                     // 链接节点
	Params   string    `gorm:"column:params;type:varchar(500)"`                                  // 链接参数
	Target   string    `gorm:"column:target;type:varchar(20)"`                                   // 打开方式
	Sort     uint      `gorm:"column:sort;type:int unsigned"`                                    // 排序权重
	Status   uint8     `gorm:"index:idx_system_menu_status;column:status;type:tinyint unsigned"` // 状态(0:禁用,1:启用)
	CreateAt time.Time `gorm:"column:create_at;type:timestamp"`                                  // 创建时间
}

// SystemOplog 系统-日志
type SystemOplog struct {
	ID       uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null"`
	Node     string    `gorm:"column:node;type:varchar(200);not null"`     // 当前操作节点
	Geoip    string    `gorm:"column:geoip;type:varchar(15);not null"`     // 操作者IP地址
	Action   string    `gorm:"column:action;type:varchar(200);not null"`   // 操作行为名称
	Content  string    `gorm:"column:content;type:varchar(1024);not null"` // 操作内容描述
	Username string    `gorm:"column:username;type:varchar(50);not null"`  // 操作人用户名
	CreateAt time.Time `gorm:"column:create_at;type:timestamp;not null"`   // 创建时间
}

// SystemQueue 系统-任务
type SystemQueue struct {
	ID        int64     `gorm:"primary_key;column:id;type:bigint;not null"`
	Code      string    `gorm:"index:idx_system_queue_code;column:code;type:varchar(20);not null"`         // 任务编号
	Title     string    `gorm:"index:idx_system_queue_title;column:title;type:varchar(50);not null"`       // 任务名称
	Command   string    `gorm:"column:command;type:varchar(500)"`                                          // 执行指令
	ExecPid   int64     `gorm:"column:exec_pid;type:bigint"`                                               // 执行进程
	ExecData  string    `gorm:"column:exec_data;type:longtext"`                                            // 执行参数
	ExecTime  int64     `gorm:"index:idx_system_queue_exec_time;column:exec_time;type:bigint"`             // 执行时间
	ExecDesc  string    `gorm:"column:exec_desc;type:varchar(500)"`                                        // 执行描述
	EnterTime float64   `gorm:"column:enter_time;type:decimal(20,4)"`                                      // 开始时间
	OuterTime float64   `gorm:"column:outer_time;type:decimal(20,4)"`                                      // 结束时间
	LoopsTime int64     `gorm:"column:loops_time;type:bigint"`                                             // 循环时间
	Attempts  int64     `gorm:"column:attempts;type:bigint"`                                               // 执行次数
	Rscript   bool      `gorm:"index:idx_system_queue_rscript;column:rscript;type:tinyint(1)"`             // 任务类型(0单例,1多例)
	Status    bool      `gorm:"index:idx_system_queue_status;column:status;type:tinyint(1)"`               // 任务状态(1新任务,2处理中,3成功,4失败)
	CreateAt  time.Time `gorm:"index:idx_system_queue_create_at;column:create_at;type:timestamp;not null"` // 创建时间
}

// SystemUser 系统-用户
type SystemUser struct {
	ID           int64     `gorm:"primary_key;column:id;type:bigint;not null"`
	Username     string    `gorm:"index:idx_system_user_username;column:username;type:varchar(50)"` // 用户账号
	Password     string    `gorm:"column:password;type:varchar(60)"`                                // 用户密码
	Nickname     string    `gorm:"column:nickname;type:varchar(50)"`                                // 用户昵称
	Headimg      string    `gorm:"column:headimg;type:varchar(255)"`                                // 头像地址
	Authorize    string    `gorm:"column:authorize;type:varchar(255)"`                              // 权限授权
	ContactQq    string    `gorm:"column:contact_qq;type:varchar(20)"`                              // 联系QQ
	ContactMail  string    `gorm:"column:contact_mail;type:varchar(20)"`                            // 联系邮箱
	ContactPhone string    `gorm:"column:contact_phone;type:varchar(20)"`                           // 联系手机
	LoginIP      string    `gorm:"column:login_ip;type:varchar(255)"`                               // 登录地址
	LoginAt      string    `gorm:"column:login_at;type:varchar(20)"`                                // 登录时间
	LoginNum     int64     `gorm:"column:login_num;type:bigint"`                                    // 登录次数
	Describe     string    `gorm:"column:describe;type:varchar(255)"`                               // 备注说明
	Status       bool      `gorm:"index:idx_system_user_status;column:status;type:tinyint(1)"`      // 状态(0禁用,1启用)
	Sort         int64     `gorm:"column:sort;type:bigint"`                                         // 排序权重
	IsDeleted    bool      `gorm:"index:idx_system_user_deleted;column:is_deleted;type:tinyint(1)"` // 删除(1删除,0未删)
	CreateAt     time.Time `gorm:"column:create_at;type:timestamp"`                                 // 创建时间
}

// UserCare 护理收藏中间表
type UserCare struct {
	ID       int       `gorm:"primary_key;column:id;type:int;not null"`
	CareID   int       `gorm:"index:care_id;column:care_id;type:int;not null"` // care_id
	UserID   int       `gorm:"index:care_id;column:user_id;type:int;not null"` // user_id
	Type     int       `gorm:"column:type;type:int;not null"`                  // 类型: 默认 ,1 收藏
	CreateAt time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt time.Time `gorm:"column:update_at;type:datetime;not null"`
}

// UserMenu [...]
type UserMenu struct {
	ID       int       `gorm:"primary_key;column:id;type:int;not null"`
	MenuID   int       `gorm:"index:menu_id;column:menu_id;type:int;not null"` // menu_id
	UserID   int       `gorm:"index:menu_id;column:user_id;type:int;not null"` // user_id
	Type     int       `gorm:"column:type;type:int;not null"`                  // 类型: 默认 1 收藏
	CreateAt time.Time `gorm:"column:create_at;type:datetime;not null"`
	UpdateAt time.Time `gorm:"column:update_at;type:datetime;not null"`
}

// WechatFans 微信-粉丝
type WechatFans struct {
	ID             uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null"`
	Appid          string    `gorm:"column:appid;type:varchar(50)"`                                            // 公众号APPID
	Unionid        string    `gorm:"index:index_wechat_fans_unionid;column:unionid;type:varchar(100)"`         // 粉丝unionid
	Openid         string    `gorm:"index:index_wechat_fans_openid;column:openid;type:varchar(100)"`           // 粉丝openid
	TagidList      string    `gorm:"column:tagid_list;type:varchar(100)"`                                      // 粉丝标签id
	IsBlack        uint8     `gorm:"index:index_wechat_fans_is_back;column:is_black;type:tinyint unsigned"`    // 是否为黑名单状态
	Subscribe      uint8     `gorm:"index:index_wechat_fans_subscribe;column:subscribe;type:tinyint unsigned"` // 关注状态(0未关注,1已关注)
	Nickname       string    `gorm:"column:nickname;type:varchar(200)"`                                        // 用户昵称
	Sex            uint8     `gorm:"column:sex;type:tinyint unsigned"`                                         // 用户性别(1男性,2女性,0未知)
	Country        string    `gorm:"column:country;type:varchar(50)"`                                          // 用户所在国家
	Province       string    `gorm:"column:province;type:varchar(50)"`                                         // 用户所在省份
	City           string    `gorm:"column:city;type:varchar(50)"`                                             // 用户所在城市
	Language       string    `gorm:"column:language;type:varchar(50)"`                                         // 用户的语言(zh_CN)
	Headimgurl     string    `gorm:"column:headimgurl;type:varchar(500)"`                                      // 用户头像
	SubscribeTime  uint64    `gorm:"column:subscribe_time;type:bigint unsigned"`                               // 关注时间
	SubscribeAt    time.Time `gorm:"column:subscribe_at;type:datetime"`                                        // 关注时间
	Remark         string    `gorm:"column:remark;type:varchar(50)"`                                           // 备注
	SubscribeScene string    `gorm:"column:subscribe_scene;type:varchar(200)"`                                 // 扫码关注场景
	QrScene        string    `gorm:"column:qr_scene;type:varchar(100)"`                                        // 二维码场景值
	QrSceneStr     string    `gorm:"column:qr_scene_str;type:varchar(200)"`                                    // 二维码场景内容
	CreateAt       time.Time `gorm:"column:create_at;type:timestamp"`                                          // 创建时间
}

// WechatFansTags 微信-标签
type WechatFansTags struct {
	ID       uint64    `gorm:"index:index_wechat_fans_tags_id;column:id;type:bigint unsigned;not null"` // 标签ID
	Appid    string    `gorm:"index:index_wechat_fans_tags_appid;column:appid;type:varchar(50)"`        // 公众号APPID
	Name     string    `gorm:"column:name;type:varchar(35)"`                                            // 标签名称
	Count    uint64    `gorm:"column:count;type:bigint unsigned"`                                       // 总数
	CreateAt time.Time `gorm:"column:create_at;type:timestamp"`                                         // 创建日期
}

// WechatKeys 微信-关键字
type WechatKeys struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint;not null"`
	Appid      string    `gorm:"index:index_wechat_keys_appid;column:appid;type:char(100)"`  // 公众号APPID
	Type       string    `gorm:"index:index_wechat_keys_type;column:type;type:varchar(20)"`  // 类型(text,image,news)
	Keys       string    `gorm:"index:index_wechat_keys_keys;column:keys;type:varchar(100)"` // 关键字
	Content    string    `gorm:"column:content;type:text"`                                   // 文本内容
	ImageURL   string    `gorm:"column:image_url;type:varchar(255)"`                         // 图片链接
	VoiceURL   string    `gorm:"column:voice_url;type:varchar(255)"`                         // 语音链接
	MusicTitle string    `gorm:"column:music_title;type:varchar(100)"`                       // 音乐标题
	MusicURL   string    `gorm:"column:music_url;type:varchar(255)"`                         // 音乐链接
	MusicImage string    `gorm:"column:music_image;type:varchar(255)"`                       // 缩略图片
	MusicDesc  string    `gorm:"column:music_desc;type:varchar(255)"`                        // 音乐描述
	VideoTitle string    `gorm:"column:video_title;type:varchar(100)"`                       // 视频标题
	VideoURL   string    `gorm:"column:video_url;type:varchar(255)"`                         // 视频URL
	VideoDesc  string    `gorm:"column:video_desc;type:varchar(255)"`                        // 视频描述
	NewsID     uint64    `gorm:"column:news_id;type:bigint unsigned"`                        // 图文ID
	Sort       uint64    `gorm:"column:sort;type:bigint unsigned"`                           // 排序字段
	Status     uint8     `gorm:"column:status;type:tinyint unsigned"`                        // 状态(0禁用,1启用)
	CreateBy   uint64    `gorm:"column:create_by;type:bigint unsigned"`                      // 创建人
	CreateAt   time.Time `gorm:"column:create_at;type:timestamp"`                            // 创建时间
}

// WechatMedia 微信-素材
type WechatMedia struct {
	ID       uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null"`
	Appid    string    `gorm:"index:index_wechat_media_appid;column:appid;type:varchar(100)"`       // 公众号ID
	Md5      string    `gorm:"index:index_wechat_media_md5;column:md5;type:varchar(32)"`            // 文件md5
	Type     string    `gorm:"index:index_wechat_media_type;column:type;type:varchar(20)"`          // 媒体类型
	MediaID  string    `gorm:"index:index_wechat_media_media_id;column:media_id;type:varchar(100)"` // 永久素材MediaID
	LocalURL string    `gorm:"column:local_url;type:varchar(300)"`                                  // 本地文件链接
	MediaURL string    `gorm:"column:media_url;type:varchar(300)"`                                  // 远程图片链接
	CreateAt time.Time `gorm:"column:create_at;type:timestamp"`                                     // 创建时间
}

// WechatNews 微信-图文
type WechatNews struct {
	ID        uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null"`
	MediaID   string    `gorm:"index:index_wechat_news_media_id;column:media_id;type:varchar(100)"`   // 永久素材MediaID
	LocalURL  string    `gorm:"column:local_url;type:varchar(300)"`                                   // 永久素材外网URL
	ArticleID string    `gorm:"index:index_wechat_news_artcle_id;column:article_id;type:varchar(60)"` // 关联图文ID(用英文逗号做分割)
	IsDeleted uint8     `gorm:"column:is_deleted;type:tinyint unsigned"`                              // 删除状态(0未删除,1已删除)
	CreateAt  time.Time `gorm:"column:create_at;type:timestamp"`                                      // 创建时间
	CreateBy  int64     `gorm:"column:create_by;type:bigint"`                                         // 创建人
}

// WechatNewsArticle 微信-文章
type WechatNewsArticle struct {
	ID               uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null"`
	Title            string    `gorm:"column:title;type:varchar(50)"`               // 素材标题
	LocalURL         string    `gorm:"column:local_url;type:varchar(300)"`          // 永久素材显示URL
	ShowCoverPic     uint8     `gorm:"column:show_cover_pic;type:tinyint unsigned"` // 显示封面(0不显示,1显示)
	Author           string    `gorm:"column:author;type:varchar(20)"`              // 文章作者
	Digest           string    `gorm:"column:digest;type:varchar(300)"`             // 摘要内容
	Content          string    `gorm:"column:content;type:longtext"`                // 图文内容
	ContentSourceURL string    `gorm:"column:content_source_url;type:varchar(200)"` // 原文地址
	ReadNum          uint64    `gorm:"column:read_num;type:bigint unsigned"`        // 阅读数量
	CreateAt         time.Time `gorm:"column:create_at;type:timestamp"`             // 创建时间
}
