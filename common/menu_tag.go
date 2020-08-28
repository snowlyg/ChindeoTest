package common

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/snowlyg/ChindeoTest/model"
	"time"
)

var MenuTagCount int

func CreateMenuTag(db *gorm.DB) *model.APIMenuTags {
	menuTag := model.APIMenuTags{Name: "菜单标签", CreateAt: time.Now(), UpdateAt: time.Now(), IsDeleted: sql.NullTime{}}
	if err := db.Create(&menuTag).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuTag create error :%v", err))
	}
	MenuTagCount++
	return &menuTag
}
