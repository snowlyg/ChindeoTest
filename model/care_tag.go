package model

import (
	"database/sql"
	"fmt"
	"time"
)

var CareTagCount int

func CreateCareTag() *CareTags {
	careTag := &CareTags{
		Name:      "护理项目标签",
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		IsDeleted: sql.NullTime{},
		Icon:      "icon",
	}
	if err := DB.Create(careTag).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuTag create error :%v", err))
	}

	CareTagCount++
	return careTag
}
