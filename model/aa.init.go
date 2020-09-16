package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/snowlyg/ChindeoTest/config"
)

var DB *gorm.DB
var err error

func init() {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", config.Config.User, config.Config.Password, config.Config.Host, config.Config.Port, config.Config.Name, "Asia%2FShanghai")
	DB, err = gorm.Open("mysql", conn)
	if err != nil {
		panic(fmt.Sprintf("conn %s Error:%s", conn, err))
	}

	if DB == nil {
		panic(fmt.Sprintf("conn %s db is nil", conn))
	}

	DB.DB().SetMaxOpenConns(1)
	DB.LogMode(false)

}

// 关闭连接
func Close() {
	if DB != nil {
		_ = DB.Close()
		DB = nil
	}
}

// 清除测试数据
func ClearTables(cleanTables bool) {
	if cleanTables {
		DB.Unscoped().Delete(APIWarnTimes{}, "id > ?", 0)

		DB.Unscoped().Delete(APIAddrs{}, "id > ?", 0)
		DB.Unscoped().Delete(APIMenuTypes{}, "id > ?", 0)
		DB.Unscoped().Delete(APIMenus{}, "id > ?", 0)

		DB.Unscoped().Table("menu_menu_tag").Delete(MenuMenuTag{}, "id > ?", 0)

		DB.Unscoped().Delete(APIMenuTags{}, "id > ?", 0)
		DB.Unscoped().Delete(APIOOrders{}, "id > ?", 0)
		DB.Unscoped().Delete(APIOOrderAddrs{}, "id > ?", 0)
		DB.Unscoped().Delete(APIOOrderMenus{}, "id > ?", 0)
		DB.Unscoped().Delete(APIOReturnOrderAddrs{}, "id > ?", 0)
		DB.Unscoped().Delete(APIOReturnOrders{}, "id > ?", 0)
		DB.Unscoped().Delete(APIOReturnOrderMenus{}, "id > ?", 0)

		DB.Unscoped().Delete(CareTypes{}, "id > ?", 0)
		DB.Unscoped().Delete(CareTags{}, "id > ?", 0)
		DB.Unscoped().Delete(Cares{}, "id > ?", 0)

		DB.Unscoped().Delete(CarerTags{}, "id > ?", 0)
		DB.Unscoped().Delete(Carers{}, "id > ?", 0)
		DB.Unscoped().Delete(CarerServerTimes{}, "id > ?", 0)

		DB.Unscoped().Table("care_care_tag").Delete(CareCareTag{}, "id > ?", 0)
		DB.Unscoped().Table("carer_carer_tag").Delete(CarerCarerTag{}, "id > ?", 0)

		DB.Unscoped().Delete(CareOrders{}, "id > ?", 0)
		DB.Unscoped().Delete(CareOrderInfos{}, "id > ?", 0)
		DB.Unscoped().Delete(CareOrderCarerInfos{}, "id > ?", 0)
		DB.Unscoped().Delete(CareOrderAddrs{}, "id > ?", 0)

		DB.Unscoped().Delete(CareReturnOrderAddrs{}, "id > ?", 0)
		DB.Unscoped().Delete(CareReturnOrders{}, "id > ?", 0)
		DB.Unscoped().Delete(CareReturnOrderInfos{}, "id > ?", 0)
		DB.Unscoped().Delete(CareReturnOrderCarerInfos{}, "id > ?", 0)

		DB.Unscoped().Delete(OnlineArticles{}, "id > ?", 0)
		DB.Unscoped().Delete(OnlineQuals{}, "id > ?", 0)
		DB.Unscoped().Delete(OnlineLocTypes{}, "id > ?", 0)
		DB.Unscoped().Delete(OnlineLocs{}, "id > ?", 0)
		DB.Unscoped().Delete(OnlineUserTypes{}, "id > ?", 0)
		DB.Unscoped().Delete(OnlinePatients{}, "id > ?", 0)
		DB.Unscoped().Delete(OnlinePatientProfiles{}, "id > ?", 0)
		DB.Unscoped().Table("online_user_patient").Delete(OnlineUserPatient{}, "id > ?", 0)
		DB.Unscoped().Delete(CommonComments{}, "id > ?", 0)

		DB.Unscoped().Delete(ShopCates{}, "id > ?", 0)
		DB.Unscoped().Table("shop_brand_cate").Delete(ShopBrandCate{}, "id > ?", 0)
		DB.Unscoped().Delete(ShopBrands{}, "id > ?", 0)
		DB.Unscoped().Delete(ShopSpecGroups{}, "id > ?", 0)
		DB.Unscoped().Delete(ShopSpecParams{}, "id > ?", 0)

		DB.Unscoped().Delete(ShopSpus{}, "id > ?", 0)
		DB.Unscoped().Delete(ShopSpuDetails{}, "id > ?", 0)
		DB.Unscoped().Delete(ShopSkus{}, "id > ?", 0)
		DB.Unscoped().Delete(ShopSpuSpecs{}, "id > ?", 0)
		DB.Unscoped().Table("shop_cate_spu").Delete(ShopCateSpu{}, "id > ?", 0)

		DB.Unscoped().Delete(ShopOrders{}, "id > ?", 0)
		DB.Unscoped().Delete(ShopOrderAddrs{}, "id > ?", 0)
		DB.Unscoped().Delete(ShopOrderSkus{}, "id > ?", 0)
		DB.Unscoped().Delete(ShopOrderComments{}, "id > ?", 0)
		DB.Unscoped().Delete(UserSpu{}, "id > ?", 0)
	}
}
