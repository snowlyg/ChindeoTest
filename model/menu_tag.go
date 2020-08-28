package model

import (
	"database/sql"
	"fmt"
	"time"
)

var MenuTagCount int

func CreateMenuTag() *APIMenuTags {
	menuTag := APIMenuTags{Name: "菜单标签", CreateAt: time.Now(), UpdateAt: time.Now(), IsDeleted: sql.NullTime{}}
	if err := DB.Create(&menuTag).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuTag create error :%v", err))
	}
	MenuTagCount++
	return &menuTag
}
