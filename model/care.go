package model

import (
	"database/sql"
	"fmt"
	"time"
)

var CareAmount int
var CareCount int

func CreateCare(careTagId, careTypeId int, isTag bool) *Cares {
	care := &Cares{
		Name:          "护理项目名称",
		Desc:          "护理项目名称",
		TimeType:      "时",
		Status:        true,
		Amount:        CareAmount,
		MinPrice:      CareMiniPrice,
		MaxPrice:      CareMaxPrice,
		Cover:         "sdfsdfds",
		CareDetail:    "dsfsdfdsf",
		CareTypeID:    careTypeId,
		ApplicationID: AppId,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}

	if err := DB.Create(care).Error; err != nil {
		fmt.Println(fmt.Sprintf("care create error :%v", err))
	}

	if isTag {
		tagCare := CareCareTag{CareID: care.ID, CareTagID: careTagId, UpdateAt: time.Now(), CreateAt: time.Now()}
		if err := DB.Table("care_care_tag").Create(&tagCare).Error; err != nil {
			fmt.Println(fmt.Sprintf("tagCare create error :%v", err))
		}
	}
	CareCount++
	return care
}
