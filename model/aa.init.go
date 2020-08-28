package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/snowlyg/ChindeoTest/config"
)

func Init() *gorm.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", config.Config.User, config.Config.Password, config.Config.Host, config.Config.Port, config.Config.Name, "Asia%2FShanghai")
	Db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(fmt.Sprintf("conn %s Error:%s", conn, err))
	}

	Db.DB().SetMaxOpenConns(1)
	Db.LogMode(false)

	return Db

}

// 关闭连接
func Close(db *gorm.DB) {
	if db != nil {
		_ = db.Close()
		db = nil
	}
}

// 清除测试数据
func ClearTables(db *gorm.DB) {
	db.Unscoped().Delete(APIWarnTimes{}, "id > ?", 0)

	db.Unscoped().Delete(APIAddrs{}, "id > ?", 0)
	db.Unscoped().Delete(APIMenuTypes{}, "id > ?", 0)
	db.Unscoped().Delete(APIMenus{}, "id > ?", 0)

	db.Unscoped().Table("menu_menu_tag").Delete(MenuMenuTag{}, "id > ?", 0)

	db.Unscoped().Delete(APIMenuTags{}, "id > ?", 0)
	db.Unscoped().Delete(APIOOrders{}, "id > ?", 0)
	db.Unscoped().Delete(APIOOrderAddrs{}, "id > ?", 0)
	db.Unscoped().Delete(APIOOrderMenus{}, "id > ?", 0)
	db.Unscoped().Delete(APIOReturnOrderAddrs{}, "id > ?", 0)
	db.Unscoped().Delete(APIOReturnOrders{}, "id > ?", 0)
	db.Unscoped().Delete(APIOReturnOrderMenus{}, "id > ?", 0)

	db.Unscoped().Delete(CareTypes{}, "id > ?", 0)
	db.Unscoped().Delete(CareTags{}, "id > ?", 0)
	db.Unscoped().Delete(Cares{}, "id > ?", 0)

	db.Unscoped().Delete(CarerTags{}, "id > ?", 0)
	db.Unscoped().Delete(Carers{}, "id > ?", 0)
	db.Unscoped().Delete(CarerServerTimes{}, "id > ?", 0)

	db.Unscoped().Table("care_care_tag").Delete(CareCareTag{}, "id > ?", 0)
	db.Unscoped().Table("carer_carer_tag").Delete(CarerCarerTag{}, "id > ?", 0)

	db.Unscoped().Delete(CareOrders{}, "id > ?", 0)
	db.Unscoped().Delete(CareOrderInfos{}, "id > ?", 0)
	db.Unscoped().Delete(CareOrderCarerInfos{}, "id > ?", 0)
	db.Unscoped().Delete(CareOrderAddrs{}, "id > ?", 0)

	db.Unscoped().Delete(CareReturnOrderAddrs{}, "id > ?", 0)
	db.Unscoped().Delete(CareReturnOrders{}, "id > ?", 0)
	db.Unscoped().Delete(CareReturnOrderInfos{}, "id > ?", 0)
	db.Unscoped().Delete(CareReturnOrderCarerInfos{}, "id > ?", 0)

	db.Unscoped().Delete(Articles{}, "id > ?", 0)
	db.Unscoped().Delete(Quals{}, "id > ?", 0)
	db.Unscoped().Delete(LocTypes{}, "id > ?", 0)
	db.Unscoped().Delete(Locs{}, "id > ?", 0)
	db.Unscoped().Delete(UserTypes{}, "id > ?", 0)
	db.Unscoped().Delete(Patients{}, "id > ?", 0)
	db.Unscoped().Delete(PatientProfiles{}, "id > ?", 0)
	db.Unscoped().Table("user_patient").Delete(UserPatient{}, "id > ?", 0)
	db.Unscoped().Delete(Comments{}, "id > ?", 0)

}
