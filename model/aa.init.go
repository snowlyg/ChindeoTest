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
	db.Unscoped().Delete(APIAddrs{}, "id > ?", 0)
	db.Unscoped().Delete(APIMenuTypes{}, "id > ?", 0)
	db.Unscoped().Delete(APIMenus{}, "id > ?", 0)
	db.Unscoped().Delete(MenuMenuTag{}, "id > ?", 0)
	db.Unscoped().Delete(APIOOrders{}, "id > ?", 0)
	db.Unscoped().Delete(APIOOrderAddrs{}, "id > ?", 0)
	db.Unscoped().Delete(APIOOrderMenus{}, "id > ?", 0)
}
