package model

import (
	"database/sql"
	"fmt"
	"time"
)

var MenuTypeCount int

func CreateMenuType() *APIMenuTypes {
	menuType := APIMenuTypes{Name: "菜单类型", ApplicationID: AppId, CreateAt: time.Now(), UpdateAt: time.Now(), IsDeleted: sql.NullTime{}}
	if err := DB.Create(&menuType).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuType create error :%v", err))
	}
	MenuTypeCount++
	return &menuType
}
