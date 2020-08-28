package model

import (
	"database/sql"
	"fmt"
	"time"
)

var CareTypeCount int

func CreateCareType() *CareTypes {
	careType := &CareTypes{
		Name:      "护理项目类型",
		EnName:    "en_name",
		Status:    true,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		IsDeleted: sql.NullTime{},
		Icon:      "icon",
	}
	if err := DB.Create(careType).Error; err != nil {
		fmt.Println(fmt.Sprintf("careType create error :%v", err))
	}
	CareTypeCount++
	return careType
}
