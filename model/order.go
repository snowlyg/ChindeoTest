package model

import (
	"database/sql"
	"time"
)

// APIAddrs [...]
type APIAddrs struct {
	ID           int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name         string       `gorm:"column:name;type:varchar(50);not null" json:"name"`   // 联系人姓名
	Phone        string       `gorm:"column:phone;type:varchar(30);not null" json:"phone"` // 手机
	Addr         string       `gorm:"column:addr;type:varchar(200);not null" json:"addr"`  // 地址
	Sex          int8         `gorm:"column:sex;type:tinyint;not null" json:"sex"`
	IsDefault    bool         `gorm:"column:is_default;type:tinyint(1);not null" json:"is_default"`        // 默认地址
	UserID       int          `gorm:"index:user_id;column:user_id;type:int;not null" json:"user_id"`       // user_id
	HospitalName string       `gorm:"column:hospital_name;type:varchar(50);not null" json:"hospital_name"` // 医院名称
	LocName      string       `gorm:"column:loc_name;type:varchar(50);not null" json:"loc_name"`           // 病区名称
	BedNum       string       `gorm:"column:bed_num;type:varchar(10);not null" json:"bed_num"`             // 床号
	HospitalNo   string       `gorm:"column:hospital_no;type:varchar(20);not null" json:"hospital_no"`     // 住院号
	Disease      string       `gorm:"column:disease;type:varchar(150);not null" json:"disease"`            // 病种
	Age          int          `gorm:"column:age;type:int;not null" json:"age"`                             // 年龄
	Province     string       `gorm:"column:province;type:varchar(20);not null" json:"province"`           // 省份
	City         string       `gorm:"column:city;type:varchar(20);not null" json:"city"`                   // 城市
	County       string       `gorm:"column:county;type:varchar(20);not null" json:"county"`               // 县城
	Postcode     string       `gorm:"column:postcode;type:char(10);not null" json:"postcode"`              // 邮编
	CreateAt     time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt     time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted    sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// APIApplications [...]
type APIApplications struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name          string       `gorm:"column:name;type:varchar(50);not null" json:"name"`
	AppID         string       `gorm:"column:app_id;type:varchar(30);not null" json:"app_id"`          // app_id
	AppSecret     string       `gorm:"column:app_secret;type:varchar(100);not null" json:"app_secret"` // app_secret
	Tel           string       `gorm:"column:tel;type:varchar(100);not null" json:"tel"`
	Addr          string       `gorm:"column:addr;type:varchar(100);not null" json:"addr"`
	Describle     string       `gorm:"column:describle;type:varchar(100);not null" json:"describle"` // 应用描述
	AppType       string       `gorm:"column:app_type;type:varchar(50);not null" json:"app_type"`    // app_type
	CreateAt      time.Time    `gorm:"column:create_at;type:timestamp;not null" json:"create_at"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	UserID        int          `gorm:"column:user_id;type:int;not null" json:"user_id"`
	CnareaID      int          `gorm:"column:cnarea_id;type:int;not null" json:"cnarea_id"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:timestamp;not null" json:"update_at"`
	BusinessHours string       `gorm:"column:business_hours;type:varchar(50);not null" json:"business_hours"` // 营业时间
	CtHospitalID  int          `gorm:"column:ct_hospital_id;type:int;not null" json:"ct_hospital_id"`         // ct_hospital_id
}

// APIAuths [...]
type APIAuths struct {
	ID         int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name       string       `gorm:"column:name;type:varchar(30);not null" json:"name"`              // 权限名称
	Controller string       `gorm:"column:controller;type:varchar(100);not null" json:"controller"` // 控制器
	Action     string       `gorm:"column:action;type:varchar(100);not null" json:"action"`         // 操作方法
	Method     string       `gorm:"column:method;type:varchar(100);not null" json:"method"`         // 请求方法
	Desc       string       `gorm:"column:desc;type:varchar(30);not null" json:"desc"`              // 描述
	Remark     string       `gorm:"column:remark;type:varchar(100);not null" json:"remark"`         // 备注
	Status     bool         `gorm:"column:status;type:tinyint(1);not null" json:"status"`           // 账号状态
	CreateAt   time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt   time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted  sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// APIMenuTags [...]
type APIMenuTags struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name      string       `gorm:"column:name;type:varchar(50);not null" json:"name"` // 标签名称
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
}

// APIMenuTypes [...]
type APIMenuTypes struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name          string       `gorm:"column:name;type:varchar(50);not null" json:"name"` // 菜品类型名称
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	ApplicationID int          `gorm:"column:application_id;type:int;not null" json:"application_id"`
}

// APIMenus [...]
type APIMenus struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name          string       `gorm:"column:name;type:varchar(50);not null" json:"name"`     // 菜品名称
	TimeType      MenuTimeType `gorm:"column:time_type;type:int;not null" json:"time_type"`   // 菜品时段类型
	Desc          string       `gorm:"column:desc;type:varchar(100);not null" json:"desc"`    // 菜品介绍
	Status        bool         `gorm:"column:status;type:tinyint(1);not null" json:"status"`  // 菜品状态：上下架
	Amount        int          `gorm:"column:amount;type:int;not null" json:"amount"`         // 销量
	Price         float64      `gorm:"column:price;type:decimal(10,2);not null" json:"price"` // 价格
	Cover         string       `gorm:"column:cover;type:varchar(1200);not null" json:"cover"` // 封面
	Sort          int          `gorm:"column:sort;type:int;not null" json:"sort"`             // 排序
	Pics          string       `gorm:"column:pics;type:varchar(1200);not null" json:"pics"`   // 餐品图片
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	ApplicationID int          `gorm:"column:application_id;type:int;not null" json:"application_id"`
	MenuTypeID    int          `gorm:"index:menu_type_id;column:menu_type_id;type:int;not null" json:"menu_type_id"` // menu_type_id
	MenuTags      []*APIMenuTags
	MenuType      *APIMenuTypes
}

// APIOOrderAddrs [...]
type APIOOrderAddrs struct {
	ID           int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name         string       `gorm:"column:name;type:varchar(50);not null" json:"name"`   // 联系人姓名
	Phone        string       `gorm:"column:phone;type:varchar(30);not null" json:"phone"` // 手机
	Addr         string       `gorm:"column:addr;type:varchar(200);not null" json:"addr"`  // 地址
	Sex          int          `gorm:"column:sex;type:int;not null" json:"sex"`             // 性别:1:男,0:女
	CreateAt     time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt     time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted    sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	OOrderID     int          `gorm:"index:o_order_id;column:o_order_id;type:int;not null" json:"o_order_id"` // o_order_id
	HospitalName string       `gorm:"column:hospital_name;type:varchar(50);not null" json:"hospital_name"`    // 医院名称
	LocName      string       `gorm:"column:loc_name;type:varchar(50);not null" json:"loc_name"`              // 病区名称
	BedNum       string       `gorm:"column:bed_num;type:varchar(10);not null" json:"bed_num"`                // 床号
	HospitalNo   string       `gorm:"column:hospital_no;type:varchar(20);not null" json:"hospital_no"`        // 住院号
	Disease      string       `gorm:"column:disease;type:varchar(150);not null" json:"disease"`               // 病种
	Age          int          `gorm:"column:age;type:int;not null" json:"age"`                                // 年龄
}

// APIOOrderMenus [...]
type APIOOrderMenus struct {
	ID           int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	MenuName     string       `gorm:"column:menu_name;type:varchar(50);not null" json:"menu_name"`           // 菜品名称
	MenuType     string       `gorm:"column:menu_type;type:varchar(30);not null" json:"menu_type"`           // 菜品类型
	MenuTimeType string       `gorm:"column:menu_time_type;type:varchar(30);not null" json:"menu_time_type"` // 菜品时段类型
	MenuDesc     string       `gorm:"column:menu_desc;type:varchar(100);not null" json:"menu_desc"`          // 菜品介绍
	MenuID       int          `gorm:"index:o_order_id;column:menu_id;type:int;not null" json:"menu_id"`      // menu_id
	Price        float64      `gorm:"column:price;type:decimal(10,0);not null" json:"price"`                 // 价格
	Amount       int          `gorm:"column:amount;type:int;not null" json:"amount"`                         // 数量
	Cover        string       `gorm:"column:cover;type:varchar(255);not null" json:"cover"`                  // 菜品图片
	CreateAt     time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt     time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted    sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	OOrderID     int          `gorm:"index:o_order_id;column:o_order_id;type:int;not null" json:"o_order_id"` // o_order_id
}

// APIOOrders [...]
type APIOOrders struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	OrderNo       string       `gorm:"index:user_id;column:order_no;type:varchar(50);not null" json:"order_no"` // 订单号
	Status        int          `gorm:"column:status;type:int;not null" json:"status"`                           // 订单状态：1：待付款，2：已付款，3，配送中，4：已取消，5：已完成
	Amount        int          `gorm:"column:amount;type:int;not null" json:"amount"`                           // 总数量
	Total         float64      `gorm:"column:total;type:decimal(10,2);not null" json:"total"`                   // 总金额
	PayType       int          `gorm:"column:pay_type;type:tinyint(1);not null" json:"pay_type"`
	Rmk           string       `gorm:"column:rmk;type:varchar(500)" json:"rmk"`                       // 备注
	AppType       int          `gorm:"column:app_type;type:int;not null" json:"app_type"`             // 订单应用类型，1：小程序，2：床旁
	OpenID        string       `gorm:"column:open_id;type:varchar(100);not null" json:"open_id"`      // open_id
	TradeType     string       `gorm:"column:trade_type;type:varchar(10);not null" json:"trade_type"` // JSAPI、NATIVE、APP
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	IsReturn      bool         `gorm:"column:is_return;type:tinyint(1);not null" json:"is_return"`                  // 是否退款
	ApplicationID int          `gorm:"index:user_id;column:application_id;type:int;not null" json:"application_id"` // application_id
	UserID        int          `gorm:"index:user_id;column:user_id;type:int;not null" json:"user_id"`               // user_id
	OrderAddr     *APIOOrderAddrs
	OrderMenus    []*APIOOrderMenus
	OrderComments []*CommonComments
}

// APIOReturnOrderAddrs [...]
type APIOReturnOrderAddrs struct {
	ID             int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name           string       `gorm:"column:name;type:varchar(50);not null" json:"name"`   // 联系人姓名
	Phone          string       `gorm:"column:phone;type:varchar(30);not null" json:"phone"` // 手机
	Addr           string       `gorm:"column:addr;type:varchar(200);not null" json:"addr"`  // 地址
	Sex            int          `gorm:"column:sex;type:int;not null" json:"sex"`             // 性别:1:男,0:女
	CreateAt       time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt       time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted      sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	ApplicationID  int          `gorm:"column:application_id;type:int;not null" json:"application_id"`                               // application_id
	OReturnOrderID int          `gorm:"index:o_return_order_id;column:o_return_order_id;type:int;not null" json:"o_return_order_id"` // o_return_order_id
	HospitalName   string       `gorm:"column:hospital_name;type:varchar(50);not null" json:"hospital_name"`                         // 医院名称
	LocName        string       `gorm:"column:loc_name;type:varchar(50);not null" json:"loc_name"`                                   // 病区名称
	BedNum         string       `gorm:"column:bed_num;type:varchar(10);not null" json:"bed_num"`                                     // 床号
	HospitalNo     string       `gorm:"column:hospital_no;type:varchar(20);not null" json:"hospital_no"`                             // 住院号
	Disease        string       `gorm:"column:disease;type:varchar(150);not null" json:"disease"`                                    // 病种
	Age            int          `gorm:"column:age;type:int;not null" json:"age"`                                                     // 年龄
}

// APIOReturnOrderMenus [...]
type APIOReturnOrderMenus struct {
	ID             int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	MenuName       string       `gorm:"column:menu_name;type:varchar(50);not null" json:"menu_name"`             // 菜品名称
	MenuType       string       `gorm:"column:menu_type;type:varchar(30);not null" json:"menu_type"`             // 菜品类型
	MenuTimeType   string       `gorm:"column:menu_time_type;type:varchar(30);not null" json:"menu_time_type"`   // 菜品时段类型
	MenuDesc       string       `gorm:"column:menu_desc;type:varchar(100);not null" json:"menu_desc"`            // 菜品介绍
	MenuID         int          `gorm:"index:o_return_order_id;column:menu_id;type:int;not null" json:"menu_id"` // menu_id
	Price          float64      `gorm:"column:price;type:decimal(10,0);not null" json:"price"`                   // 价格
	Amount         int          `gorm:"column:amount;type:int;not null" json:"amount"`                           // 数量
	Cover          string       `gorm:"column:cover;type:varchar(255);not null" json:"cover"`                    // 菜品图片
	CreateAt       time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt       time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted      sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	ApplicationID  int          `gorm:"column:application_id;type:int;not null" json:"application_id"`                               // application_id
	OReturnOrderID int          `gorm:"index:o_return_order_id;column:o_return_order_id;type:int;not null" json:"o_return_order_id"` // o_return_order_id
}

// APIOReturnOrders [...]
type APIOReturnOrders struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	OrderNo       string       `gorm:"index:o_order_id;column:order_no;type:varchar(50);not null" json:"order_no"` // 订单号
	Status        int          `gorm:"column:status;type:int;not null" json:"status"`                              // 订单状态：1，待审核，2：审核通过，3：审核驳回，4：退款成功，5：退款失败，
	Amount        int          `gorm:"column:amount;type:int;not null" json:"amount"`                              // 总数量
	Total         float64      `gorm:"column:total;type:decimal(10,0);not null" json:"total"`                      // 总金额
	OpenID        string       `gorm:"column:open_id;type:varchar(100);not null" json:"open_id"`                   // open_id
	AppType       int          `gorm:"column:app_type;type:int;not null" json:"app_type"`                          // 订单应用类型，1：小程序，2：床旁
	TradeType     string       `gorm:"column:trade_type;type:varchar(10);not null" json:"trade_type"`              // JSAPI、NATIVE、APP
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	OOrderID      int          `gorm:"index:o_order_id;column:o_order_id;type:int;not null" json:"o_order_id"`         // o_order_id
	ApplicationID int          `gorm:"index:o_order_id;column:application_id;type:int;not null" json:"application_id"` // application_id
	PayType       bool         `gorm:"column:pay_type;type:tinyint(1);not null" json:"pay_type"`
	UserID        int          `gorm:"column:user_id;type:int;not null" json:"user_id"` // user_id
}

// APIRoles [...]
type APIRoles struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name      string       `gorm:"column:name;type:varchar(30);not null" json:"name"`      // 角色名称
	Desc      string       `gorm:"column:desc;type:varchar(30);not null" json:"desc"`      // 描述
	Remark    string       `gorm:"column:remark;type:varchar(100);not null" json:"remark"` // 备注
	Status    bool         `gorm:"column:status;type:tinyint(1);not null" json:"status"`   // 账号状态
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// APIUsers [...]
type APIUsers struct {
	ID        int       `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Username  string    `gorm:"unique_index:open_id;column:username;type:varchar(30);not null" json:"username"`  // 用户名
	Nickname  string    `gorm:"column:nickname;type:varchar(30);not null" json:"nickname"`                       // 名称
	Phone     string    `gorm:"unique_index:open_id;column:phone;type:varchar(30);not null" json:"phone"`        // 手机
	Email     string    `gorm:"unique_index:open_id;column:email;type:varchar(50);not null" json:"email"`        // 邮箱
	Sex       int       `gorm:"column:sex;type:int;not null" json:"sex"`                                         // 性别
	Password  string    `gorm:"column:password;type:varchar(32);not null" json:"password"`                       // 用户密码
	Status    bool      `gorm:"column:status;type:tinyint(1);not null" json:"status"`                            // 账号状态
	AvatarURL string    `gorm:"column:avatar_url;type:varchar(500);not null" json:"avatar_url"`                  // 用户头像
	OpenID    string    `gorm:"unique_index:open_id;column:open_id;type:varchar(100);not null" json:"open_id"`   // open_id
	UnionID   string    `gorm:"unique_index:open_id;column:union_id;type:varchar(100);not null" json:"union_id"` // union_id
	Country   string    `gorm:"column:country;type:varchar(100);not null" json:"country"`                        // 国家
	Province  string    `gorm:"column:province;type:varchar(100);not null" json:"province"`                      // 省份
	City      string    `gorm:"column:city;type:varchar(100);not null" json:"city"`                              // 城镇
	Mac       string    `gorm:"column:mac;type:varchar(100);not null" json:"mac"`                                // mac 地址
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IDCardNo  string    `gorm:"column:id_card_no;type:varchar(20);not null" json:"id_card_no"` // 身份证
	IsAuth    bool      `gorm:"column:is_auth;type:tinyint(1);not null" json:"is_auth"`        // 是否实名认证
	Realname  string    `gorm:"column:realname;type:varchar(30);not null" json:"realname"`     // 真实姓名
}

// APIWarnTimes [...]
type APIWarnTimes struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Weeks     string       `gorm:"column:weeks;type:varchar(30);not null" json:"weeks"`  // 星期
	Time      string       `gorm:"column:time;type:varchar(20);not null" json:"time"`    // 提醒时间
	Status    bool         `gorm:"column:status;type:tinyint(1);not null" json:"status"` // 是否启用
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	UserID    int          `gorm:"index:user_id;column:user_id;type:int;not null" json:"user_id"` // user_id
}

// Articles [...]
type OnlineArticles struct {
	ID               int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Title            string       `gorm:"column:title;type:varchar(50);not null" json:"title"`                            // 文章标题
	Digest           string       `gorm:"column:digest;type:varchar(200);not null" json:"digest"`                         // 文章摘要
	Content          string       `gorm:"column:content;type:text;not null" json:"content"`                               // 文章内容
	Author           string       `gorm:"column:author;type:varchar(50);not null" json:"author"`                          // 文章作者
	LocalURL         string       `gorm:"column:local_url;type:varchar(300);not null" json:"local_url"`                   // 封面
	ContentSourceURL string       `gorm:"column:content_source_url;type:varchar(200);not null" json:"content_source_url"` // 原文地址
	CreateAt         time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt         time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted        sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// CareCareTag 护理标签和护理关系中间表
type CareCareTag struct {
	ID        int       `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	CareID    int       `gorm:"index:care_id;column:care_id;type:int;not null" json:"care_id"`         // care_id
	CareTagID int       `gorm:"index:care_id;column:care_tag_id;type:int;not null" json:"care_tag_id"` // care_tag_id
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
}

// CareOrderAddrs 订单地址表
type CareOrderAddrs struct {
	ID           int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name         string       `gorm:"column:name;type:varchar(50);not null" json:"name"`                   // 联系人姓名
	Phone        string       `gorm:"column:phone;type:varchar(30);not null" json:"phone"`                 // 手机
	Addr         string       `gorm:"column:addr;type:varchar(200);not null" json:"addr"`                  // 地址
	Sex          int          `gorm:"column:sex;type:int;not null" json:"sex"`                             // 性别:1:男,0:女
	HospitalName string       `gorm:"column:hospital_name;type:varchar(50);not null" json:"hospital_name"` // 医院名称
	LocName      string       `gorm:"column:loc_name;type:varchar(50);not null" json:"loc_name"`           // 病区名称
	BedNum       string       `gorm:"column:bed_num;type:varchar(10);not null" json:"bed_num"`             // 床号
	HospitalNo   string       `gorm:"column:hospital_no;type:varchar(20);not null" json:"hospital_no"`     // 住院号
	Disease      string       `gorm:"column:disease;type:varchar(150);not null" json:"disease"`            // 病种
	IsDeleted    sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	CareOrderID  int          `gorm:"index:care_order_id;column:care_order_id;type:int;not null" json:"care_order_id"` // care_order_id
	CreateAt     time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt     time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	Age          int          `gorm:"column:age;type:int;not null" json:"age"` // 年龄
}

// CareOrderCarerInfos 订单护工信息表
type CareOrderCarerInfos struct {
	ID              int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name            string       `gorm:"column:name;type:varchar(50);not null" json:"name"`                         // 姓名
	Age             int          `gorm:"column:age;type:int;not null" json:"age"`                                   // 年龄
	WorkExp         int          `gorm:"column:work_exp;type:int;not null" json:"work_exp"`                         // 工作年限
	Price           float64      `gorm:"column:price;type:decimal(10,2);not null" json:"price"`                     // 价格
	Sex             int          `gorm:"column:sex;type:int;not null" json:"sex"`                                   // 性别:1:男,0:女
	Phone           string       `gorm:"column:phone;type:varchar(30);not null" json:"phone"`                       // 手机
	Desc            string       `gorm:"column:desc;type:varchar(500);not null" json:"desc"`                        // 简介
	Avatar          string       `gorm:"column:avatar;type:varchar(200);not null" json:"avatar"`                    // 头像
	CarerTags       string       `gorm:"column:carer_tags;type:varchar(200);not null" json:"carer_tags"`            // 标签
	ApplicationName string       `gorm:"column:application_name;type:varchar(50);not null" json:"application_name"` // 医院名称
	CreateAt        time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt        time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted       sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	CareOrderID     int          `gorm:"index:care_order_id;column:care_order_id;type:int;not null" json:"care_order_id"` // care_order_id
	TimeType        string       `gorm:"column:time_type;type:varchar(5);not null" json:"time_type"`                      // 时间类型：时，天，月
	CarerDetail     string       `gorm:"column:carer_detail;type:varchar(2000);not null" json:"carer_detail"`             // 服务内容:采用富文本
}

// CareOrderInfos 护理项目表
type CareOrderInfos struct {
	ID              int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name            string       `gorm:"column:name;type:varchar(50);not null" json:"name"`                         // 名称
	Desc            string       `gorm:"column:desc;type:varchar(100);not null" json:"desc"`                        // 简介
	TimeType        string       `gorm:"column:time_type;type:varchar(5);not null" json:"time_type"`                // 时间类型：时，天，月
	CareType        string       `gorm:"column:care_type;type:varchar(30);not null" json:"care_type"`               // 类型
	CareTags        string       `gorm:"column:care_tags;type:varchar(80);not null" json:"care_tags"`               // 标签：用 | 分隔
	MinPrice        float64      `gorm:"column:min_price;type:decimal(10,2);not null" json:"min_price"`             // 最小价格
	MaxPrice        float64      `gorm:"column:max_price;type:decimal(10,2);not null" json:"max_price"`             // 最大价格
	Cover           string       `gorm:"column:cover;type:varchar(1200);not null" json:"cover"`                     // 封面
	CareDetail      string       `gorm:"column:care_detail;type:varchar(2000);not null" json:"care_detail"`         // 服务内容:采用富文本
	ApplicationName string       `gorm:"column:application_name;type:varchar(50);not null" json:"application_name"` // 医院名称
	CreateAt        time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt        time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted       sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	CareOrderID     int          `gorm:"index:care_order_id;column:care_order_id;type:int;not null" json:"care_order_id"` // care_order_id
}

// CareOrders 护理订单表
type CareOrders struct {
	ID                 int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	OrderNo            string       `gorm:"column:order_no;type:varchar(50);not null" json:"order_no"`                   // 订单号
	Status             int          `gorm:"column:status;type:int;not null" json:"status"`                               // 订单状态：1：待付款，2：已付款，3，进行中，4：已取消，5：已完成
	PayType            int          `gorm:"column:pay_type;type:int;not null" json:"pay_type"`                           // 订单支付类型：1：微信，2：支付宝
	Total              float64      `gorm:"column:total;type:decimal(10,2);not null" json:"total"`                       // 总金额
	Rmk                string       `gorm:"column:rmk;type:varchar(500);not null" json:"rmk"`                            // 备注
	StartAt            time.Time    `gorm:"column:start_at;type:datetime;not null" json:"start_at"`                      // 服务开始时间
	EndAt              time.Time    `gorm:"column:end_at;type:datetime;not null" json:"end_at"`                          // 服务结束时间
	OpenID             string       `gorm:"column:open_id;type:varchar(100);not null" json:"open_id"`                    // open_id
	TradeType          string       `gorm:"column:trade_type;type:varchar(10);not null" json:"trade_type"`               // JSAPI、NATIVE、APP
	IsReturn           bool         `gorm:"column:is_return;type:tinyint(1);not null" json:"is_return"`                  // 是否退款
	ApplicationID      int          `gorm:"index:user_id;column:application_id;type:int;not null" json:"application_id"` // application_id
	UserID             int          `gorm:"index:user_id;column:user_id;type:int;not null" json:"user_id"`               // user_id
	IsDeleted          sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	CreateAt           time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt           time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	AppType            int          `gorm:"column:app_type;type:int;not null" json:"app_type"`                // 订单应用类型，1：小程序，2：床旁
	CarerID            int          `gorm:"index:carer_id;column:carer_id;type:int;not null" json:"carer_id"` // 护工id
	CareOrderAddr      *CareOrderAddrs
	CareOrderInfo      *CareOrderInfos
	CareOrderCarerInfo *CareOrderCarerInfos
	CareOrderComments  []*CommonComments
}

// CareReturnOrderAddrs 退款订单地址表
type CareReturnOrderAddrs struct {
	ID                int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name              string       `gorm:"column:name;type:varchar(50);not null" json:"name"`                   // 联系人姓名
	Phone             string       `gorm:"column:phone;type:varchar(30);not null" json:"phone"`                 // 手机
	Addr              string       `gorm:"column:addr;type:varchar(200);not null" json:"addr"`                  // 地址
	Sex               int          `gorm:"column:sex;type:int;not null" json:"sex"`                             // 性别:1:男,0:女
	HospitalName      string       `gorm:"column:hospital_name;type:varchar(50);not null" json:"hospital_name"` // 医院名称
	LocName           string       `gorm:"column:loc_name;type:varchar(50);not null" json:"loc_name"`           // 病区名称
	BedNum            string       `gorm:"column:bed_num;type:varchar(10);not null" json:"bed_num"`             // 床号
	HospitalNo        string       `gorm:"column:hospital_no;type:varchar(20);not null" json:"hospital_no"`     // 住院号
	Disease           string       `gorm:"column:disease;type:varchar(150);not null" json:"disease"`            // 病种
	CreateAt          time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt          time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted         sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	ApplicationID     int          `gorm:"column:application_id;type:int;not null" json:"application_id"`                                        // application_id
	CareReturnOrderID int          `gorm:"index:care_return_order_id;column:care_return_order_id;type:int;not null" json:"care_return_order_id"` // care_return_order_id
	Age               int          `gorm:"column:age;type:int;not null" json:"age"`                                                              // 年龄
}

// CareReturnOrderCarerInfos 订单护工信息表
type CareReturnOrderCarerInfos struct {
	ID                int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name              string       `gorm:"column:name;type:varchar(50);not null" json:"name"`                         // 姓名
	Age               int          `gorm:"column:age;type:int;not null" json:"age"`                                   // 年龄
	WorkExp           int          `gorm:"column:work_exp;type:int;not null" json:"work_exp"`                         // 工作年限
	Price             float64      `gorm:"column:price;type:decimal(10,2);not null" json:"price"`                     // 价格
	Sex               int          `gorm:"column:sex;type:int;not null" json:"sex"`                                   // 性别:1:男,0:女
	Phone             string       `gorm:"column:phone;type:varchar(30);not null" json:"phone"`                       // 手机
	Desc              string       `gorm:"column:desc;type:varchar(500);not null" json:"desc"`                        // 简介
	Avatar            string       `gorm:"column:avatar;type:varchar(200);not null" json:"avatar"`                    // 头像
	CarerTags         string       `gorm:"column:carer_tags;type:varchar(200);not null" json:"carer_tags"`            // 标签
	ApplicationName   string       `gorm:"column:application_name;type:varchar(50);not null" json:"application_name"` // 医院名称
	CreateAt          time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt          time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted         sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	CareReturnOrderID int          `gorm:"index:care_return_order_id;column:care_return_order_id;type:int;not null" json:"care_return_order_id"` // care_return_order_id
	TimeType          string       `gorm:"column:time_type;type:varchar(5);not null" json:"time_type"`                                           // 时间类型：时，天，月
	CarerDetail       string       `gorm:"column:carer_detail;type:varchar(2000);not null" json:"carer_detail"`                                  // 服务内容:采用富文本
}

// CareReturnOrderInfos 护理项目表
type CareReturnOrderInfos struct {
	ID                int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name              string       `gorm:"column:name;type:varchar(50);not null" json:"name"`                         // 名称
	Desc              string       `gorm:"column:desc;type:varchar(100);not null" json:"desc"`                        // 简介
	TimeType          string       `gorm:"column:time_type;type:varchar(5);not null" json:"time_type"`                // 时间类型：时，天，月
	CareType          string       `gorm:"column:care_type;type:varchar(30);not null" json:"care_type"`               // 类型
	CareTags          string       `gorm:"column:care_tags;type:varchar(80);not null" json:"care_tags"`               // 标签：用 | 分隔
	MinPrice          float64      `gorm:"column:min_price;type:decimal(10,2);not null" json:"min_price"`             // 最小价格
	MaxPrice          float64      `gorm:"column:max_price;type:decimal(10,2);not null" json:"max_price"`             // 最大价格
	Cover             string       `gorm:"column:cover;type:varchar(1200);not null" json:"cover"`                     // 封面
	CareDetail        string       `gorm:"column:care_detail;type:varchar(2000);not null" json:"care_detail"`         // 服务内容:采用富文本
	ApplicationName   string       `gorm:"column:application_name;type:varchar(50);not null" json:"application_name"` // 医院名称
	CreateAt          time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt          time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted         sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	CareReturnOrderID int          `gorm:"index:care_return_order_id;column:care_return_order_id;type:int;not null" json:"care_return_order_id"` // care_return_order_id
}

// CareReturnOrders 退款订单表
type CareReturnOrders struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	OrderNo       string       `gorm:"column:order_no;type:varchar(50);not null" json:"order_no"`     // 订单号
	Status        int          `gorm:"column:status;type:int;not null" json:"status"`                 // 订单状态：1，待审核，2：审核通过，3：审核驳回，4：退款成功，5：退款失败，
	PayType       int          `gorm:"column:pay_type;type:int;not null" json:"pay_type"`             // 订单支付类型：1：微信，2：支付宝
	Total         float64      `gorm:"column:total;type:decimal(10,2);not null" json:"total"`         // 总金额
	OpenID        string       `gorm:"column:open_id;type:varchar(100);not null" json:"open_id"`      // open_id
	TradeType     string       `gorm:"column:trade_type;type:varchar(10);not null" json:"trade_type"` // JSAPI、NATIVE、APP
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	CareOrderID   int          `gorm:"index:care_order_id;column:care_order_id;type:int;not null" json:"care_order_id"`   // care_order_id
	ApplicationID int          `gorm:"index:care_order_id;column:application_id;type:int;not null" json:"application_id"` // application_id
	UserID        int          `gorm:"column:user_id;type:int;not null" json:"user_id"`                                   // user_id
	AppType       int          `gorm:"column:app_type;type:int;not null" json:"app_type"`                                 // 订单应用类型，1：小程序，2：床旁
	CarerID       int          `gorm:"index:carer_id;column:carer_id;type:int;not null" json:"carer_id"`                  // 护工id
}

// CareTags 护理标签表
type CareTags struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name      string       `gorm:"column:name;type:varchar(50);not null" json:"name"` // 名称
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	Icon      string       `gorm:"column:icon;type:varchar(200);not null" json:"icon"` // 图标
}

// CareTypes 护理类型表
type CareTypes struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name      string       `gorm:"column:name;type:varchar(50);not null" json:"name"`       // 名称
	EnName    string       `gorm:"column:en_name;type:varchar(50);not null" json:"en_name"` // 英文名称
	Status    bool         `gorm:"column:status;type:tinyint(1);not null" json:"status"`    // 状态：启用，禁用
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	Icon      string       `gorm:"column:icon;type:varchar(200);not null" json:"icon"` // 图标
}

// CarerCarerTag 护工标签和护工关系中间表
type CarerCarerTag struct {
	ID         int       `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	CarerID    int       `gorm:"index:carer_id;column:carer_id;type:int;not null" json:"carer_id"`         // carer_id
	CarerTagID int       `gorm:"index:carer_id;column:carer_tag_id;type:int;not null" json:"carer_tag_id"` // carer_tag_id
	CreateAt   time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt   time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
}

// CarerServerTimes 护工服务时间表
type CarerServerTimes struct {
	ID          int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	StartAt     time.Time    `gorm:"column:start_at;type:datetime;not null" json:"start_at"`                          // 服务开始时间
	EndAt       time.Time    `gorm:"column:end_at;type:datetime;not null" json:"end_at"`                              // 服务结束时间
	CarerID     int          `gorm:"index:care_order_id;column:carer_id;type:int;not null" json:"carer_id"`           // carer_id
	TimeType    string       `gorm:"column:time_type;type:varchar(5);not null" json:"time_type"`                      // 时间类型：时，天
	CareOrderID int          `gorm:"index:care_order_id;column:care_order_id;type:int;not null" json:"care_order_id"` // care_order_id
	CreateAt    time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt    time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted   sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// CarerTags 护工标签表
type CarerTags struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name      string       `gorm:"column:name;type:varchar(50);not null" json:"name"` // 名称
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	Icon      string       `gorm:"column:icon;type:varchar(200);not null" json:"icon"` // 图标
}

// Carers 护工表
type Carers struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name          string       `gorm:"column:name;type:varchar(50);not null" json:"name"`      // 姓名
	Age           int          `gorm:"column:age;type:int;not null" json:"age"`                // 年龄
	WorkExp       int          `gorm:"column:work_exp;type:int;not null" json:"work_exp"`      // 工作年限
	Price         float64      `gorm:"column:price;type:decimal(10,2);not null" json:"price"`  // 价格
	Sex           int          `gorm:"column:sex;type:int;not null" json:"sex"`                // 性别:1:男,0:女
	Phone         string       `gorm:"column:phone;type:varchar(30);not null" json:"phone"`    // 手机
	Desc          string       `gorm:"column:desc;type:varchar(500);not null" json:"desc"`     // 简介
	Avatar        string       `gorm:"column:avatar;type:varchar(200);not null" json:"avatar"` // 头像
	Amount        int          `gorm:"column:amount;type:int;not null" json:"amount"`          // 服务数量
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	ApplicationID int          `gorm:"index:application_id;column:application_id;type:int;not null" json:"application_id"` // application_id
	Status        bool         `gorm:"column:status;type:tinyint(1);not null" json:"status"`                               // 状态：上下架
	TimeType      string       `gorm:"column:time_type;type:varchar(5);not null" json:"time_type"`                         // 时间类型：时，天，月
	CarerDetail   string       `gorm:"column:carer_detail;type:varchar(2000);not null" json:"carer_detail"`                // 服务内容:采用富文本
	CardNo        string       `gorm:"column:card_no;type:varchar(12);not null" json:"card_no"`                            // 工号
}

// Cares 护理项目表
type Cares struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name          string       `gorm:"column:name;type:varchar(50);not null" json:"name"`                 // 名称
	Desc          string       `gorm:"column:desc;type:varchar(100);not null" json:"desc"`                // 简介
	TimeType      string       `gorm:"column:time_type;type:varchar(5);not null" json:"time_type"`        // 时间类型：时，天，月
	Status        bool         `gorm:"column:status;type:tinyint(1);not null" json:"status"`              // 状态：上下架
	Amount        int          `gorm:"column:amount;type:int;not null" json:"amount"`                     // 销量
	MinPrice      float64      `gorm:"column:min_price;type:decimal(10,2);not null" json:"min_price"`     // 最小价格
	MaxPrice      float64      `gorm:"column:max_price;type:decimal(10,2);not null" json:"max_price"`     // 最大价格
	Cover         string       `gorm:"column:cover;type:varchar(1200);not null" json:"cover"`             // 封面
	Sort          int          `gorm:"column:sort;type:int;not null" json:"sort"`                         // 排序
	CareDetail    string       `gorm:"column:care_detail;type:varchar(2000);not null" json:"care_detail"` // 服务内容:采用富文本
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	CareTypeID    int          `gorm:"index:care_type_id;column:care_type_id;type:int;not null" json:"care_type_id"`     // care_type_id
	ApplicationID int          `gorm:"index:care_type_id;column:application_id;type:int;not null" json:"application_id"` // application_id
}

// Cnareas 中国行政地区表
type Cnareas struct {
	ID         string  `gorm:"primary_key;column:id;type:mediumint unsigned;not null" json:"-"`
	Level      uint8   `gorm:"column:level;type:tinyint unsigned;not null" json:"level"`                                  // 层级
	ParentCode uint64  `gorm:"index:idx_parent_code;column:parent_code;type:bigint unsigned;not null" json:"parent_code"` // 父级行政代码
	AreaCode   uint64  `gorm:"unique;column:area_code;type:bigint unsigned;not null" json:"area_code"`                    // 行政代码
	ZipCode    string  `gorm:"column:zip_code;type:mediumint(6) unsigned zerofill;not null" json:"zip_code"`              // 邮政编码
	CityCode   string  `gorm:"column:city_code;type:char(6);not null" json:"city_code"`                                   // 区号
	Name       string  `gorm:"column:name;type:varchar(50);not null" json:"name"`                                         // 名称
	ShortName  string  `gorm:"column:short_name;type:varchar(50);not null" json:"short_name"`                             // 简称
	MergerName string  `gorm:"column:merger_name;type:varchar(50);not null" json:"merger_name"`                           // 组合名
	Pinyin     string  `gorm:"column:pinyin;type:varchar(30);not null" json:"pinyin"`                                     // 拼音
	Lng        float64 `gorm:"column:lng;type:decimal(10,6);not null" json:"lng"`                                         // 经度
	Lat        float64 `gorm:"column:lat;type:decimal(10,6);not null" json:"lat"`                                         // 纬度
}

// LocLocType [...]
type OnlineLocLocType struct {
	ID        int       `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	LocID     int       `gorm:"index:loc_id;column:loc_id;type:int;not null" json:"loc_id"`           // loc_id
	LocTypeID int       `gorm:"index:loc_id;column:loc_type_id;type:int;not null" json:"loc_type_id"` // loc_type_id
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
}

// LocTypes [...]
type OnlineLocTypes struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	ApplicationID int          `gorm:"index:name;column:application_id;type:int;not null" json:"application_id"` // application_id
	Name          string       `gorm:"index:name;column:name;type:varchar(255);not null" json:"name"`            // 标签名称
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// Locs [...]
type OnlineLocs struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	LocID         int          `gorm:"index:loc_id;column:loc_id;type:int;not null" json:"loc_id"`                   // utp_id
	CtHospitalID  int          `gorm:"column:ct_hospital_id;type:int;not null" json:"ct_hospital_id"`                // ct_hospital_id
	ApplicationID int          `gorm:"index:loc_id;column:application_id;type:int;not null" json:"application_id"`   // application_id
	LocWardFlag   int          `gorm:"index:loc_id;column:loc_ward_flag;type:int;not null" json:"loc_ward_flag"`     // loc_ward_flag
	LocActiveFlag int          `gorm:"index:loc_id;column:loc_active_flag;type:int;not null" json:"loc_active_flag"` // loc_active_flag
	LocDesc       string       `gorm:"column:loc_desc;type:varchar(255);not null" json:"loc_desc"`                   // 科室名称
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// MenuMenuTag [...]
type MenuMenuTag struct {
	ID        int       `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	MenuID    int       `gorm:"index:menu_id;column:menu_id;type:int;not null" json:"menu_id"`         // menu_id
	MenuTagID int       `gorm:"index:menu_id;column:menu_tag_id;type:int;not null" json:"menu_tag_id"` // menu_tag_id
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
}

// Migrations [...]
type Migrations struct {
	Version       int64     `gorm:"primary_key;column:version;type:bigint;not null" json:"-"`
	MigrationName string    `gorm:"column:migration_name;type:varchar(100)" json:"migration_name"`
	StartTime     time.Time `gorm:"column:start_time;type:timestamp;not null" json:"start_time"`
	EndTime       time.Time `gorm:"column:end_time;type:timestamp;not null" json:"end_time"`
	Breakpoint    bool      `gorm:"column:breakpoint;type:tinyint(1);not null" json:"breakpoint"`
}

// MiniApps [...]
type SystemMiniApps struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	AppID     string       `gorm:"unique_index:uuid;column:app_id;type:varchar(100);not null" json:"app_id"` // app_id
	AppSecret string       `gorm:"column:app_secret;type:varchar(100);not null" json:"app_secret"`           // app_secret
	UUID      string       `gorm:"unique_index:uuid;column:uuid;type:varchar(100);not null" json:"uuid"`     // uuid
	Name      string       `gorm:"index:name;column:name;type:varchar(50);not null" json:"name"`             // 应用名称
	Describle string       `gorm:"column:describle;type:varchar(100);not null" json:"describle"`             // 应用描述
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// PatientProfiles [...]
type OnlinePatientProfiles struct {
	ID          int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Realname    string       `gorm:"column:realname;type:varchar(30);not null" json:"realname"`          // 姓名
	Sex         int          `gorm:"column:sex;type:int;not null" json:"sex"`                            // 性别 1：男，0：女
	Age         int          `gorm:"column:age;type:int;not null" json:"age"`                            // 年龄
	Weight      int          `gorm:"column:weight;type:int;not null" json:"weight"`                      // 体重
	Hepatic     string       `gorm:"column:hepatic;type:varchar(150);not null" json:"hepatic"`           // 肝功能
	Kidney      string       `gorm:"column:kidney;type:varchar(150);not null" json:"kidney"`             // 肾功能
	BredInfo    string       `gorm:"column:bred_info;type:varchar(300);not null" json:"bred_info"`       // 孕育情况
	Allergic    string       `gorm:"column:allergic;type:varchar(300);not null" json:"allergic"`         // 过敏史
	MedicalCare string       `gorm:"column:medical_care;type:varchar(300);not null" json:"medical_care"` // 过往病史
	Briefly     string       `gorm:"column:briefly;type:varchar(300);not null" json:"briefly"`           // 主述
	MedicalDec  string       `gorm:"column:medical_dec;type:varchar(800);not null" json:"medical_dec"`   // 病史
	Suggestions string       `gorm:"column:suggestions;type:varchar(800);not null" json:"suggestions"`   // 建议
	CreateAt    time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt    time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	PatientID   int          `gorm:"index:patient_id;column:patient_id;type:int;not null" json:"patient_id"`
	IsDeleted   sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// Patients [...]
type OnlinePatients struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Username  string       `gorm:"unique_index:open_id;column:username;type:varchar(30);not null" json:"username"`  // 用户名
	Nickname  string       `gorm:"column:nickname;type:varchar(30);not null" json:"nickname"`                       // 名称
	Phone     string       `gorm:"unique_index:open_id;column:phone;type:varchar(30);not null" json:"phone"`        // 手机
	Email     string       `gorm:"unique_index:open_id;column:email;type:varchar(50);not null" json:"email"`        // 邮箱
	Sex       int          `gorm:"column:sex;type:int;not null" json:"sex"`                                         // 性别 1：男，0：女
	Password  string       `gorm:"column:password;type:varchar(32);not null" json:"password"`                       // 用户密码
	Status    bool         `gorm:"column:status;type:tinyint(1);not null" json:"status"`                            // 账号状态
	AvatarURL string       `gorm:"column:avatar_url;type:varchar(100);not null" json:"avatar_url"`                  // 用户头像
	OpenID    string       `gorm:"unique_index:open_id;column:open_id;type:varchar(100);not null" json:"open_id"`   // open_id
	UnionID   string       `gorm:"unique_index:open_id;column:union_id;type:varchar(100);not null" json:"union_id"` // union_id
	Country   string       `gorm:"column:country;type:varchar(100);not null" json:"country"`                        // 国家
	Province  string       `gorm:"column:province;type:varchar(100);not null" json:"province"`                      // 省份
	City      string       `gorm:"column:city;type:varchar(100);not null" json:"city"`                              // 城镇
	Mac       string       `gorm:"column:mac;type:varchar(100);not null" json:"mac"`                                // mac 地址
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// Profiles [...]
type Profiles struct {
	ID            int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Realname      string       `gorm:"column:realname;type:varchar(30);not null" json:"realname"`               // 真实姓名
	Hospital      string       `gorm:"column:hospital;type:varchar(50);not null" json:"hospital"`               // 医院名称
	Department    string       `gorm:"column:department;type:varchar(50);not null" json:"department"`           // 科室名称
	Postion       string       `gorm:"column:postion;type:varchar(50);not null" json:"postion"`                 // 职位名称
	School        string       `gorm:"column:school;type:varchar(50);not null" json:"school"`                   // 毕业学校
	Expertise     string       `gorm:"column:expertise;type:text;not null" json:"expertise"`                    // 专长
	Profile       string       `gorm:"column:profile;type:text;not null" json:"profile"`                        // 个人简介
	ServerDate    string       `gorm:"column:server_date;type:varchar(50);not null" json:"server_date"`         // 服务日期
	ServerStartAt string       `gorm:"column:server_start_at;type:varchar(50);not null" json:"server_start_at"` // 健康服务开始时间
	ServerEndAt   string       `gorm:"column:server_end_at;type:varchar(50);not null" json:"server_end_at"`     // 健康服务结束时间
	IsServerOpen  bool         `gorm:"column:is_server_open;type:tinyint(1);not null" json:"is_server_open"`    // 健康服务开关
	IsAuth        bool         `gorm:"column:is_auth;type:tinyint(1);not null" json:"is_auth"`                  // 是否实名认证
	Name          string       `gorm:"column:name;type:varchar(50);not null" json:"name"`                       // 联系人姓名
	Phone         string       `gorm:"column:phone;type:varchar(30);not null" json:"phone"`                     // 手机
	Sex           int          `gorm:"column:sex;type:int;not null" json:"sex"`                                 // 性别:1:男,0:女
	HospitalName  string       `gorm:"column:hospital_name;type:varchar(50);not null" json:"hospital_name"`     // 医院名称
	LocName       string       `gorm:"column:loc_name;type:varchar(50);not null" json:"loc_name"`               // 病区名称
	BedNum        string       `gorm:"column:bed_num;type:varchar(10);not null" json:"bed_num"`                 // 床号
	HospitalNo    string       `gorm:"column:hospital_no;type:varchar(20);not null" json:"hospital_no"`         // 住院号
	Disease       string       `gorm:"column:disease;type:varchar(150);not null" json:"disease"`                // 病种
	Age           int          `gorm:"column:age;type:int;not null" json:"age"`                                 // 年龄
	CreateAt      time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt      time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	UserID        int          `gorm:"column:user_id;type:int;not null" json:"user_id"`
	IsDeleted     sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// Quals [...]
type OnlineQuals struct {
	ID           int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Loc          string       `gorm:"column:loc;type:varchar(30);not null" json:"loc"`              // 所在科室
	Hospital     string       `gorm:"column:hospital;type:varchar(50);not null" json:"hospital"`    // 执业医院
	Postion      string       `gorm:"column:postion;type:varchar(50);not null" json:"postion"`      // 职位名称
	QualImg      string       `gorm:"column:qual_img;type:text;not null" json:"qual_img"`           // 执业证书图片
	ExpertiseImg string       `gorm:"column:expertise_img;type:text;not null" json:"expertise_img"` // 专业技术资格证图片
	LocImg       string       `gorm:"column:loc_img;type:text;not null" json:"loc_img"`             // 职称证图片
	Status       int          `gorm:"column:status;type:int;not null" json:"status"`                // 认证状态 1:未认证，2：审核中，3：已认证
	CreateAt     time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt     time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	UserID       int          `gorm:"column:user_id;type:int;not null" json:"user_id"`
	IsDeleted    sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// Rules [...]
type Rules struct {
	ID    int    `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Ptype string `gorm:"column:ptype;type:varchar(255)" json:"ptype"`
	V0    string `gorm:"column:v0;type:varchar(255)" json:"v0"`
	V1    string `gorm:"column:v1;type:varchar(255)" json:"v1"`
	V2    string `gorm:"column:v2;type:varchar(255)" json:"v2"`
	V3    string `gorm:"column:v3;type:varchar(255)" json:"v3"`
	V4    string `gorm:"column:v4;type:varchar(255)" json:"v4"`
	V5    string `gorm:"column:v5;type:varchar(255)" json:"v5"`
}

// StateRecords 对账记录表
type StateRecords struct {
	ID           int       `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	StateType    int       `gorm:"column:state_type;type:int;not null" json:"state_type"`            // 对账类型：1：微信，2：支付宝
	OrderType    int       `gorm:"column:order_type;type:int;not null" json:"order_type"`            // 订单类型：1：付款，2：退款
	OrderNum     int       `gorm:"column:order_num;type:int;not null" json:"order_num"`              // 订单数量
	Total        float64   `gorm:"column:total;type:decimal(10,2);not null" json:"total"`            // 总金额
	RealOrderNum int       `gorm:"column:real_order_num;type:int;not null" json:"real_order_num"`    // 支付平台订单数量
	RealTotal    float64   `gorm:"column:real_total;type:decimal(10,2);not null" json:"real_total"`  // 支付平台总金额
	IsException  bool      `gorm:"column:is_exception;type:tinyint(1);not null" json:"is_exception"` // 是否异常
	OrderNos     string    `gorm:"column:order_nos;type:varchar(800);not null" json:"order_nos"`     // 异常订单号用 ， 隔开
	ExtraType    int       `gorm:"column:extra_type;type:int;not null" json:"extra_type"`            // 平账方式：1：盈余，2：亏损
	Extra        float64   `gorm:"column:extra;type:decimal(10,2);not null" json:"extra"`            // 平账金额
	StartAt      time.Time `gorm:"column:start_at;type:datetime;not null" json:"start_at"`           // 对账订单开始时间范围
	EndAt        time.Time `gorm:"column:end_at;type:datetime;not null" json:"end_at"`               // 对账订单结束时间范围
	Username     string    `gorm:"column:username;type:varchar(50);not null" json:"username"`        // 平账人
	Rmk          string    `gorm:"column:rmk;type:varchar(500);not null" json:"rmk"`                 // 平账备注
	RecordAt     time.Time `gorm:"column:record_at;type:datetime;not null" json:"record_at"`         // 对账时间
	CreateAt     time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt     time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted    bool      `gorm:"column:is_deleted;type:tinyint(1)" json:"is_deleted"`
}

// SystemAuth 系统-权限
type SystemAuth struct {
	ID       uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null" json:"-"`
	Title    string    `gorm:"index:idx_system_auth_title;column:title;type:varchar(100)" json:"title"`        // 权限名称
	Desc     string    `gorm:"column:desc;type:varchar(500)" json:"desc"`                                      // 备注说明
	Sort     uint64    `gorm:"column:sort;type:bigint unsigned" json:"sort"`                                   // 排序权重
	Status   uint8     `gorm:"index:idx_system_auth_status;column:status;type:tinyint unsigned" json:"status"` // 权限状态(1使用,0禁用)
	CreateAt time.Time `gorm:"column:create_at;type:timestamp" json:"create_at"`                               // 创建时间
}

// SystemAuthNode 系统-授权
type SystemAuthNode struct {
	ID   uint64 `gorm:"primary_key;column:id;type:bigint unsigned;not null" json:"-"`
	Auth uint64 `gorm:"index:idx_system_auth_auth;column:auth;type:bigint unsigned" json:"auth"` // 角色
	Node string `gorm:"index:idx_system_auth_node;column:node;type:varchar(200)" json:"node"`    // 节点
}

// SystemConfig 系统-配置
type SystemConfig struct {
	Type  string `gorm:"index:idx_system_config_type;column:type;type:varchar(20)" json:"type"`  // 分类
	Name  string `gorm:"index:idx_system_config_name;column:name;type:varchar(100)" json:"name"` // 配置名
	Value string `gorm:"column:value;type:varchar(2048)" json:"value"`                           // 配置值
}

// SystemData 系统-数据
type SystemData struct {
	ID    uint64 `gorm:"primary_key;column:id;type:bigint unsigned;not null" json:"-"`
	Name  string `gorm:"index:idx_system_data_name;column:name;type:varchar(100)" json:"name"` // 配置名
	Value string `gorm:"column:value;type:longtext" json:"value"`                              // 配置值
}

// SystemMenu 系统-菜单
type SystemMenu struct {
	ID       uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null" json:"-"`
	Pid      uint64    `gorm:"column:pid;type:bigint unsigned" json:"pid"`                                     // 上级ID
	Title    string    `gorm:"column:title;type:varchar(100)" json:"title"`                                    // 菜单名称
	Icon     string    `gorm:"column:icon;type:varchar(100)" json:"icon"`                                      // 菜单图标
	Node     string    `gorm:"index:idx_system_menu_node;column:node;type:varchar(100)" json:"node"`           // 节点代码
	URL      string    `gorm:"column:url;type:varchar(400)" json:"url"`                                        // 链接节点
	Params   string    `gorm:"column:params;type:varchar(500)" json:"params"`                                  // 链接参数
	Target   string    `gorm:"column:target;type:varchar(20)" json:"target"`                                   // 打开方式
	Sort     uint      `gorm:"column:sort;type:int unsigned" json:"sort"`                                      // 排序权重
	Status   uint8     `gorm:"index:idx_system_menu_status;column:status;type:tinyint unsigned" json:"status"` // 状态(0:禁用,1:启用)
	CreateAt time.Time `gorm:"column:create_at;type:timestamp" json:"create_at"`                               // 创建时间
}

// SystemOplog 系统-日志
type SystemOplog struct {
	ID       uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null" json:"-"`
	Node     string    `gorm:"column:node;type:varchar(200);not null" json:"node"`        // 当前操作节点
	Geoip    string    `gorm:"column:geoip;type:varchar(15);not null" json:"geoip"`       // 操作者IP地址
	Action   string    `gorm:"column:action;type:varchar(200);not null" json:"action"`    // 操作行为名称
	Content  string    `gorm:"column:content;type:varchar(1024);not null" json:"content"` // 操作内容描述
	Username string    `gorm:"column:username;type:varchar(50);not null" json:"username"` // 操作人用户名
	CreateAt time.Time `gorm:"column:create_at;type:timestamp;not null" json:"create_at"` // 创建时间
}

// SystemQueue 系统-任务
type SystemQueue struct {
	ID        int64     `gorm:"primary_key;column:id;type:bigint;not null" json:"-"`
	Code      string    `gorm:"index:idx_system_queue_code;column:code;type:varchar(20);not null" json:"code"`              // 任务编号
	Title     string    `gorm:"index:idx_system_queue_title;column:title;type:varchar(50);not null" json:"title"`           // 任务名称
	Command   string    `gorm:"column:command;type:varchar(500)" json:"command"`                                            // 执行指令
	ExecPid   int64     `gorm:"column:exec_pid;type:bigint" json:"exec_pid"`                                                // 执行进程
	ExecData  string    `gorm:"column:exec_data;type:longtext" json:"exec_data"`                                            // 执行参数
	ExecTime  int64     `gorm:"index:idx_system_queue_exec_time;column:exec_time;type:bigint" json:"exec_time"`             // 执行时间
	ExecDesc  string    `gorm:"column:exec_desc;type:varchar(500)" json:"exec_desc"`                                        // 执行描述
	EnterTime float64   `gorm:"column:enter_time;type:decimal(20,4)" json:"enter_time"`                                     // 开始时间
	OuterTime float64   `gorm:"column:outer_time;type:decimal(20,4)" json:"outer_time"`                                     // 结束时间
	LoopsTime int64     `gorm:"column:loops_time;type:bigint" json:"loops_time"`                                            // 循环时间
	Attempts  int64     `gorm:"column:attempts;type:bigint" json:"attempts"`                                                // 执行次数
	Rscript   bool      `gorm:"index:idx_system_queue_rscript;column:rscript;type:tinyint(1)" json:"rscript"`               // 任务类型(0单例,1多例)
	Status    bool      `gorm:"index:idx_system_queue_status;column:status;type:tinyint(1)" json:"status"`                  // 任务状态(1新任务,2处理中,3成功,4失败)
	CreateAt  time.Time `gorm:"index:idx_system_queue_create_at;column:create_at;type:timestamp;not null" json:"create_at"` // 创建时间
}

// SystemUser 系统-用户
type SystemUser struct {
	ID           int64     `gorm:"primary_key;column:id;type:bigint;not null" json:"-"`
	Username     string    `gorm:"index:idx_system_user_username;column:username;type:varchar(50)" json:"username"`   // 用户账号
	Password     string    `gorm:"column:password;type:varchar(60)" json:"password"`                                  // 用户密码
	Nickname     string    `gorm:"column:nickname;type:varchar(50)" json:"nickname"`                                  // 用户昵称
	Headimg      string    `gorm:"column:headimg;type:varchar(255)" json:"headimg"`                                   // 头像地址
	Authorize    string    `gorm:"column:authorize;type:varchar(255)" json:"authorize"`                               // 权限授权
	ContactQq    string    `gorm:"column:contact_qq;type:varchar(20)" json:"contact_qq"`                              // 联系QQ
	ContactMail  string    `gorm:"column:contact_mail;type:varchar(20)" json:"contact_mail"`                          // 联系邮箱
	ContactPhone string    `gorm:"column:contact_phone;type:varchar(20)" json:"contact_phone"`                        // 联系手机
	LoginIP      string    `gorm:"column:login_ip;type:varchar(255)" json:"login_ip"`                                 // 登录地址
	LoginAt      string    `gorm:"column:login_at;type:varchar(20)" json:"login_at"`                                  // 登录时间
	LoginNum     int64     `gorm:"column:login_num;type:bigint" json:"login_num"`                                     // 登录次数
	Describe     string    `gorm:"column:describe;type:varchar(255)" json:"describe"`                                 // 备注说明
	Status       bool      `gorm:"index:idx_system_user_status;column:status;type:tinyint(1)" json:"status"`          // 状态(0禁用,1启用)
	Sort         int64     `gorm:"column:sort;type:bigint" json:"sort"`                                               // 排序权重
	IsDeleted    bool      `gorm:"index:idx_system_user_deleted;column:is_deleted;type:tinyint(1)" json:"is_deleted"` // 删除(1删除,0未删)
	CreateAt     time.Time `gorm:"column:create_at;type:timestamp" json:"create_at"`                                  // 创建时间
}

// UserCare 护理收藏中间表
type UserCare struct {
	ID       int       `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	CareID   int       `gorm:"index:care_id;column:care_id;type:int;not null" json:"care_id"` // care_id
	UserID   int       `gorm:"index:care_id;column:user_id;type:int;not null" json:"user_id"` // user_id
	Type     int       `gorm:"column:type;type:int;not null" json:"type"`                     // 类型: 默认 ,1 收藏
	CreateAt time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
}

// UserMenu [...]
type UserMenu struct {
	ID       int       `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	MenuID   int       `gorm:"index:menu_id;column:menu_id;type:int;not null" json:"menu_id"` // menu_id
	UserID   int       `gorm:"index:menu_id;column:user_id;type:int;not null" json:"user_id"` // user_id
	Type     int       `gorm:"column:type;type:int;not null" json:"type"`                     // 类型: 默认 1 收藏
	CreateAt time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
}

// UserPatient [...]
type OnlineUserPatient struct {
	ID        int       `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	UserID    int       `gorm:"index:user_id;column:user_id;type:int;not null" json:"user_id"`       // user_id
	PatientID int       `gorm:"index:user_id;column:patient_id;type:int;not null" json:"patient_id"` // patient_id
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
}

// UserTypes [...]
type OnlineUserTypes struct {
	ID          int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	UtpID       int          `gorm:"index:utp_id;column:utp_id;type:int;not null" json:"utp_id"`              // utp_id
	UtpCode     string       `gorm:"column:utp_code;type:varchar(255);not null" json:"utp_code"`              // 代码
	UtpDesc     string       `gorm:"index:utp_id;column:utp_desc;type:varchar(255);not null" json:"utp_desc"` // 描述
	UtpType     string       `gorm:"column:utp_type;type:varchar(2);not null" json:"utp_type"`                // N:护士；D：医生；O:其他
	UtpActive   bool         `gorm:"column:utp_active;type:tinyint(1);not null" json:"utp_active"`            // 激活
	UtpContrast string       `gorm:"column:utp_contrast;type:varchar(255);not null" json:"utp_contrast"`      // 对照关系,HIS编码
	CreateAt    time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt    time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted   sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// WechatFans 微信-粉丝
type WechatFans struct {
	ID             uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null" json:"-"`
	Appid          string    `gorm:"column:appid;type:varchar(50)" json:"appid"`                                                // 公众号APPID
	Unionid        string    `gorm:"index:index_wechat_fans_unionid;column:unionid;type:varchar(100)" json:"unionid"`           // 粉丝unionid
	Openid         string    `gorm:"index:index_wechat_fans_openid;column:openid;type:varchar(100)" json:"openid"`              // 粉丝openid
	TagidList      string    `gorm:"column:tagid_list;type:varchar(100)" json:"tagid_list"`                                     // 粉丝标签id
	IsBlack        uint8     `gorm:"index:index_wechat_fans_is_back;column:is_black;type:tinyint unsigned" json:"is_black"`     // 是否为黑名单状态
	Subscribe      uint8     `gorm:"index:index_wechat_fans_subscribe;column:subscribe;type:tinyint unsigned" json:"subscribe"` // 关注状态(0未关注,1已关注)
	Nickname       string    `gorm:"column:nickname;type:varchar(200)" json:"nickname"`                                         // 用户昵称
	Sex            uint8     `gorm:"column:sex;type:tinyint unsigned" json:"sex"`                                               // 用户性别(1男性,2女性,0未知)
	Country        string    `gorm:"column:country;type:varchar(50)" json:"country"`                                            // 用户所在国家
	Province       string    `gorm:"column:province;type:varchar(50)" json:"province"`                                          // 用户所在省份
	City           string    `gorm:"column:city;type:varchar(50)" json:"city"`                                                  // 用户所在城市
	Language       string    `gorm:"column:language;type:varchar(50)" json:"language"`                                          // 用户的语言(zh_CN)
	Headimgurl     string    `gorm:"column:headimgurl;type:varchar(500)" json:"headimgurl"`                                     // 用户头像
	SubscribeTime  uint64    `gorm:"column:subscribe_time;type:bigint unsigned" json:"subscribe_time"`                          // 关注时间
	SubscribeAt    time.Time `gorm:"column:subscribe_at;type:datetime" json:"subscribe_at"`                                     // 关注时间
	Remark         string    `gorm:"column:remark;type:varchar(50)" json:"remark"`                                              // 备注
	SubscribeScene string    `gorm:"column:subscribe_scene;type:varchar(200)" json:"subscribe_scene"`                           // 扫码关注场景
	QrScene        string    `gorm:"column:qr_scene;type:varchar(100)" json:"qr_scene"`                                         // 二维码场景值
	QrSceneStr     string    `gorm:"column:qr_scene_str;type:varchar(200)" json:"qr_scene_str"`                                 // 二维码场景内容
	CreateAt       time.Time `gorm:"column:create_at;type:timestamp" json:"create_at"`                                          // 创建时间
}

// WechatFansTags 微信-标签
type WechatFansTags struct {
	ID       uint64    `gorm:"index:index_wechat_fans_tags_id;column:id;type:bigint unsigned;not null" json:"id"` // 标签ID
	Appid    string    `gorm:"index:index_wechat_fans_tags_appid;column:appid;type:varchar(50)" json:"appid"`     // 公众号APPID
	Name     string    `gorm:"column:name;type:varchar(35)" json:"name"`                                          // 标签名称
	Count    uint64    `gorm:"column:count;type:bigint unsigned" json:"count"`                                    // 总数
	CreateAt time.Time `gorm:"column:create_at;type:timestamp" json:"create_at"`                                  // 创建日期
}

// WechatKeys 微信-关键字
type WechatKeys struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint;not null" json:"-"`
	Appid      string    `gorm:"index:index_wechat_keys_appid;column:appid;type:char(100)" json:"appid"` // 公众号APPID
	Type       string    `gorm:"index:index_wechat_keys_type;column:type;type:varchar(20)" json:"type"`  // 类型(text,image,news)
	Keys       string    `gorm:"index:index_wechat_keys_keys;column:keys;type:varchar(100)" json:"keys"` // 关键字
	Content    string    `gorm:"column:content;type:text" json:"content"`                                // 文本内容
	ImageURL   string    `gorm:"column:image_url;type:varchar(255)" json:"image_url"`                    // 图片链接
	VoiceURL   string    `gorm:"column:voice_url;type:varchar(255)" json:"voice_url"`                    // 语音链接
	MusicTitle string    `gorm:"column:music_title;type:varchar(100)" json:"music_title"`                // 音乐标题
	MusicURL   string    `gorm:"column:music_url;type:varchar(255)" json:"music_url"`                    // 音乐链接
	MusicImage string    `gorm:"column:music_image;type:varchar(255)" json:"music_image"`                // 缩略图片
	MusicDesc  string    `gorm:"column:music_desc;type:varchar(255)" json:"music_desc"`                  // 音乐描述
	VideoTitle string    `gorm:"column:video_title;type:varchar(100)" json:"video_title"`                // 视频标题
	VideoURL   string    `gorm:"column:video_url;type:varchar(255)" json:"video_url"`                    // 视频URL
	VideoDesc  string    `gorm:"column:video_desc;type:varchar(255)" json:"video_desc"`                  // 视频描述
	NewsID     uint64    `gorm:"column:news_id;type:bigint unsigned" json:"news_id"`                     // 图文ID
	Sort       uint64    `gorm:"column:sort;type:bigint unsigned" json:"sort"`                           // 排序字段
	Status     uint8     `gorm:"column:status;type:tinyint unsigned" json:"status"`                      // 状态(0禁用,1启用)
	CreateBy   uint64    `gorm:"column:create_by;type:bigint unsigned" json:"create_by"`                 // 创建人
	CreateAt   time.Time `gorm:"column:create_at;type:timestamp" json:"create_at"`                       // 创建时间
}

// WechatMedia 微信-素材
type WechatMedia struct {
	ID       uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null" json:"-"`
	Appid    string    `gorm:"index:index_wechat_media_appid;column:appid;type:varchar(100)" json:"appid"`          // 公众号ID
	Md5      string    `gorm:"index:index_wechat_media_md5;column:md5;type:varchar(32)" json:"md5"`                 // 文件md5
	Type     string    `gorm:"index:index_wechat_media_type;column:type;type:varchar(20)" json:"type"`              // 媒体类型
	MediaID  string    `gorm:"index:index_wechat_media_media_id;column:media_id;type:varchar(100)" json:"media_id"` // 永久素材MediaID
	LocalURL string    `gorm:"column:local_url;type:varchar(300)" json:"local_url"`                                 // 本地文件链接
	MediaURL string    `gorm:"column:media_url;type:varchar(300)" json:"media_url"`                                 // 远程图片链接
	CreateAt time.Time `gorm:"column:create_at;type:timestamp" json:"create_at"`                                    // 创建时间
}

// WechatNews 微信-图文
type WechatNews struct {
	ID        uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null" json:"-"`
	MediaID   string    `gorm:"index:index_wechat_news_media_id;column:media_id;type:varchar(100)" json:"media_id"`     // 永久素材MediaID
	LocalURL  string    `gorm:"column:local_url;type:varchar(300)" json:"local_url"`                                    // 永久素材外网URL
	ArticleID string    `gorm:"index:index_wechat_news_artcle_id;column:article_id;type:varchar(60)" json:"article_id"` // 关联图文ID(用英文逗号做分割)
	IsDeleted uint8     `gorm:"column:is_deleted;type:tinyint unsigned" json:"is_deleted"`                              // 删除状态(0未删除,1已删除)
	CreateAt  time.Time `gorm:"column:create_at;type:timestamp" json:"create_at"`                                       // 创建时间
	CreateBy  int64     `gorm:"column:create_by;type:bigint" json:"create_by"`                                          // 创建人
}

// WechatNewsArticle 微信-文章
type WechatNewsArticle struct {
	ID               uint64    `gorm:"primary_key;column:id;type:bigint unsigned;not null" json:"-"`
	Title            string    `gorm:"column:title;type:varchar(50)" json:"title"`                            // 素材标题
	LocalURL         string    `gorm:"column:local_url;type:varchar(300)" json:"local_url"`                   // 永久素材显示URL
	ShowCoverPic     uint8     `gorm:"column:show_cover_pic;type:tinyint unsigned" json:"show_cover_pic"`     // 显示封面(0不显示,1显示)
	Author           string    `gorm:"column:author;type:varchar(20)" json:"author"`                          // 文章作者
	Digest           string    `gorm:"column:digest;type:varchar(300)" json:"digest"`                         // 摘要内容
	Content          string    `gorm:"column:content;type:longtext" json:"content"`                           // 图文内容
	ContentSourceURL string    `gorm:"column:content_source_url;type:varchar(200)" json:"content_source_url"` // 原文地址
	ReadNum          uint64    `gorm:"column:read_num;type:bigint unsigned" json:"read_num"`                  // 阅读数量
	CreateAt         time.Time `gorm:"column:create_at;type:timestamp" json:"create_at"`                      // 创建时间
}

// Comments 评论表
type CommonComments struct {
	ID              int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	CommentableID   int          `gorm:"column:commentable_id;type:int;not null" json:"commentable_id"`              // commentable_id
	CommentableType string       `gorm:"column:commentable_type;type:varchar(255);not null" json:"commentable_type"` // commentable_type
	UserID          int          `gorm:"index:user_id;column:user_id;type:int;not null" json:"user_id"`              // 评价人id
	ApplicationID   int          `gorm:"column:application_id;type:int;not null" json:"application_id"`              // application_id
	Content         string       `gorm:"column:content;type:varchar(800);not null" json:"content"`                   // 评价内容
	Pics            string       `gorm:"column:pics;type:varchar(2000);not null" json:"pics"`                        // 评价图片数组
	Star            int          `gorm:"column:star;type:int;not null" json:"star"`                                  // 评价评分1-5
	CreateAt        time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt        time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted       sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// ShopBrands 商城品牌表
type ShopBrands struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name      string       `gorm:"column:name;type:varchar(64);not null" json:"name"`    // 品牌名称
	Image     string       `gorm:"column:image;type:varchar(500);not null" json:"image"` // 品牌图片地址
	Letter    string       `gorm:"column:letter;type:varchar(1);not null" json:"letter"` // 品牌的首字母
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	Cates     []*ShopCates
}

// ShopCateBrand 商品分类和品牌的中间表，两者是多对多关系
type ShopBrandCate struct {
	ID       int       `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	CateID   int       `gorm:"index:cate_id;column:cate_id;type:int;not null" json:"cate_id"`   // cate_id
	BrandID  int       `gorm:"index:cate_id;column:brand_id;type:int;not null" json:"brand_id"` // brand_id
	CreateAt time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
}

// ShopCates 商城分类表
type ShopCates struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Title     string       `gorm:"column:title;type:varchar(20);not null" json:"title"`  // 名称
	Pid       int          `gorm:"column:pid;type:int;not null" json:"pid"`              // pid
	Icon      string       `gorm:"column:icon;type:varchar(50);not null" json:"icon"`    // 图标
	Sort      int          `gorm:"column:sort;type:int;not null" json:"sort"`            // 排序
	Level     int          `gorm:"column:level;type:int;not null" json:"level"`          // level
	URL       string       `gorm:"column:url;type:varchar(400);not null" json:"url"`     // 链接
	Type      int          `gorm:"column:type;type:int;not null" json:"type"`            // 栏目类型
	Status    bool         `gorm:"column:status;type:tinyint(1);not null" json:"status"` // 是否启用
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// ShopSkus sku表,该表表示具体的商品实体,如黑色的 64g的iphone 8
type ShopSkus struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Title     string       `gorm:"column:title;type:varchar(256);not null" json:"title"`        // 商品标题
	SkuNo     string       `gorm:"column:sku_no;type:varchar(256);not null" json:"sku_no"`      // 商品编码
	Weight    float64      `gorm:"column:weight;type:decimal(10,4);not null" json:"weight"`     // 商品重量
	Images    string       `gorm:"column:images;type:varchar(1024);not null" json:"images"`     // 商品图片 (多个图片用,号分割)
	Stock     int          `gorm:"column:stock;type:int;not null" json:"stock"`                 // 库存
	Price     float64      `gorm:"column:price;type:decimal(10,2);not null" json:"price"`       // 价格
	Indexes   string       `gorm:"column:indexes;type:varchar(32);not null" json:"indexes"`     // 特有规格参数在SPU规格模板中对应的下标组合(如1_0_0)
	OwnSpec   string       `gorm:"column:own_spec;type:varchar(1024);not null" json:"own_spec"` // SKU的特有规格参数键值对 (json格式，反序列化时请使用linkedHashMap，保证有序)
	SpuID     int          `gorm:"column:spu_id;type:int;not null" json:"spu_id"`               // SPU Id
	Status    bool         `gorm:"column:status;type:tinyint(1);not null" json:"status"`        // 是否有效 (true或false)
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// ShopSpecGroups 规格参数的分组表，每个商品分类下有多个规格参数组
type ShopSpecGroups struct {
	ID         int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	CateID     int          `gorm:"index:cate_id;column:cate_id;type:int;not null" json:"cate_id"` // cate_id
	Name       string       `gorm:"column:name;type:varchar(32);not null" json:"name"`             // 名称
	CreateAt   time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt   time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted  sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	SpecParams []*ShopSpecParams
}

// ShopSpecParams 规格参数组下的参数名
type ShopSpecParams struct {
	ID          int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	CateID      int          `gorm:"index:cate_id;column:cate_id;type:int;not null" json:"cate_id"`             // cate_id
	SpecGroupID int          `gorm:"index:cate_id;column:spec_group_id;type:int;not null" json:"spec_group_id"` // group_id
	Name        string       `gorm:"column:name;type:varchar(128);not null" json:"name"`                        // 参数
	Numeric     bool         `gorm:"column:numeric;type:tinyint(1);not null" json:"numeric"`                    // 是否是数字类型参数 (true或false)
	Unit        string       `gorm:"column:unit;type:varchar(128);not null" json:"unit"`                        // 数字类型参数的单位 (非数字类型可以为空)
	Generic     bool         `gorm:"column:generic;type:tinyint(1);not null" json:"generic"`                    // 是否是SKU通用规格 (true或false)
	Searching   bool         `gorm:"column:searching;type:tinyint(1);not null" json:"searching"`                // 是否用于搜索过滤 (true或false)
	Segments    string       `gorm:"column:segments;type:varchar(1024);not null" json:"segments"`               // 区间 (数值类型参数的预设区间值，如果需要搜索，则添加分段间隔值，如CPU频率间隔：0.5-1.0)
	CreateAt    time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt    time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted   sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// ShopSpuDetails spu表描述的是一个抽象性的商品，比如 iphone8
type ShopSpuDetails struct {
	ID           int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Description  string       `gorm:"column:description;type:tinytext;not null" json:"description"`          // 商品描述信息
	PackingList  string       `gorm:"column:packing_list;type:varchar(1024);not null" json:"packing_list"`   // 包装清单
	AfterService string       `gorm:"column:after_service;type:varchar(1024);not null" json:"after_service"` // 售后服务
	SpuID        int          `gorm:"index:spu_id;column:spu_id;type:int;not null" json:"spu_id"`            // SPU Id
	CreateAt     time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt     time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted    sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// ShopSpuSpecs spu generic_spec 特殊规格参数
type ShopSpuSpecs struct {
	ID          int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	IsGeneric   bool         `gorm:"column:is_generic;type:tinyint(1);not null" json:"is_generic"`             // 是否为一般规格参数 (true或false)
	Value       string       `gorm:"column:value;type:varchar(1024);not null" json:"value"`                    // 特有规格可选值 (json格式)
	SpuID       int          `gorm:"index:spu_id;column:spu_id;type:int;not null" json:"spu_id"`               // SPU Id
	SpecParamID int          `gorm:"index:spu_id;column:spec_param_id;type:int;not null" json:"spec_param_id"` // spec_param_id
	CreateAt    time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt    time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted   sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
}

// ShopSpus spu表描述的是一个抽象性的商品，比如 iphone8
type ShopSpus struct {
	ID        int          `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Name      string       `gorm:"column:name;type:varchar(256);not null" json:"name"`               // 商品名称
	SubTitle  string       `gorm:"column:sub_title;type:varchar(256);not null" json:"sub_title"`     // 副标题 (一般是促销信息)
	BrandID   int          `gorm:"index:brand_id;column:brand_id;type:int;not null" json:"brand_id"` // 品牌Id (商品所属的品牌)
	Status    bool         `gorm:"column:status;type:tinyint(1);not null" json:"status"`             // 是否上架 (true或false)
	CreateAt  time.Time    `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt  time.Time    `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	IsDeleted sql.NullTime `gorm:"column:is_deleted;type:datetime" json:"is_deleted"`
	Detail    *ShopSpuDetails
	Skus      []*ShopSkus
	Specs     []*ShopSpuSpecs
}

type ShopCateSpu struct {
	ID       int       `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	CateID   int       `gorm:"index:cate_id;column:cate_id;type:int;not null" json:"cate_id"` // cate_id
	SpuID    int       `gorm:"index:cate_id;column:spu_id;type:int;not null" json:"spu_id"`   // spu_id
	CreateAt time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
}
