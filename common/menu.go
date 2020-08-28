package common

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/snowlyg/ChindeoTest/model"
	"time"
)

var MenuCount int
var MenuAmount int

func CreateMenu(db *gorm.DB) *model.APIMenus {

	menuType := CreateMenuType(db)
	menuTag := CreateMenuTag(db)

	menu := model.APIMenus{
		Name:          "菜单名称",
		TimeType:      MENU_TIME_TYPE_B,
		Desc:          "菜品介绍",
		Status:        true,
		Amount:        MenuAmount,
		Price:         10.00,
		ApplicationID: AppId,
		MenuTypeID:    menuType.ID,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}

	if err := db.Create(&menu).Error; err != nil {
		fmt.Println(fmt.Sprintf("menu create error :%v", err))
	}

	tagMenu := model.MenuMenuTag{MenuID: menu.ID, MenuTagID: menuTag.ID, UpdateAt: time.Now(), CreateAt: time.Now()}
	if err := db.Table("menu_menu_tag").Create(&tagMenu).Error; err != nil {
		fmt.Println(fmt.Sprintf("tagMenu create error :%v", err))
	}
	MenuCount++
	return &menu
}
