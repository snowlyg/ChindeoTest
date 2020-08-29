package model

import (
	"database/sql"
	"fmt"
	"time"
)

var MenuNoTagCount int
var MenuCount int
var MenuAmount int

func CreateMenu(isTag bool) *APIMenus {
	menuType := CreateMenuType()
	menu := APIMenus{
		Name:          "菜单名称",
		TimeType:      MenuTimeTypeB,
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

	if err := DB.Create(&menu).Error; err != nil {
		fmt.Println(fmt.Sprintf("menu create error :%v", err))
	}

	var menuTags []*APIMenuTags
	if isTag {
		menuTag := CreateMenuTag()
		menuTags = append(menuTags, menuTag)
		tagMenu := MenuMenuTag{MenuID: menu.ID, MenuTagID: menuTag.ID, UpdateAt: time.Now(), CreateAt: time.Now()}
		if err := DB.Table("menu_menu_tag").Create(&tagMenu).Error; err != nil {
			fmt.Println(fmt.Sprintf("tagMenu create error :%v", err))
		}
		MenuCount++
	}

	menu.MenuTags = menuTags
	menu.MenuType = menuType
	MenuNoTagCount++
	return &menu
}
