package model

import (
	"database/sql"
	"fmt"
	"time"
)

var CarerCount int
var CarerNoTagCount int
var CarerAmount int

func CreateCarer(carerTagId int, isTag bool) *Carers {
	carer := &Carers{
		Name:          "护理项目名称",
		Desc:          "护理项目名称",
		TimeType:      "时",
		Status:        true,
		Amount:        CarerAmount,
		Price:         CarerPrice,
		Avatar:        "dsfsdfs",
		CarerDetail:   "dfsdfsf",
		ApplicationID: AppId,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}
	if err := DB.Create(carer).Error; err != nil {
		fmt.Println(fmt.Sprintf("carer create error :%v", err))
	}
	if isTag {
		tagCarer := CarerCarerTag{CarerID: carer.ID, CarerTagID: carerTagId, UpdateAt: time.Now(), CreateAt: time.Now()}
		if err := DB.Table("carer_carer_tag").Create(&tagCarer).Error; err != nil {
			fmt.Println(fmt.Sprintf("tagCarer create error :%v", err))
		}
		CarerCount++
	}
	CarerNoTagCount++
	return carer
}
