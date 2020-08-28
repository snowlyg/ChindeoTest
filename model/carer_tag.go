package model

import (
	"database/sql"
	"fmt"
	"time"
)

var CarerTagCount int

func CreateCarerTag() *CarerTags {
	carerTag := &CarerTags{
		Name:      "护工标签",
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		IsDeleted: sql.NullTime{},
		Icon:      "icon",
	}
	if err := DB.Create(carerTag).Error; err != nil {
		fmt.Println(fmt.Sprintf("carerTag create error :%v", err))
	}
	CarerTagCount++
	return carerTag
}
