package common

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/snowlyg/ChindeoTest/model"
	"time"
)

var MenuTypeCount int

func CreateMenuType(db *gorm.DB) *model.APIMenuTypes {
	menuType := model.APIMenuTypes{Name: "菜单类型", ApplicationID: AppId, CreateAt: time.Now(), UpdateAt: time.Now(), IsDeleted: sql.NullTime{}}
	if err := db.Create(&menuType).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuType create error :%v", err))
	}
	MenuTypeCount++
	return &menuType
}
