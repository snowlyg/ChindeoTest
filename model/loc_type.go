package model

import (
	"database/sql"
	"fmt"
	"time"
)

var LocTypeCount int

func CreateLocType() *LocTypes {
	locType := &LocTypes{
		Name:          Fake.JobTitle(),
		ApplicationID: AppId,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}
	if err := DB.Create(locType).Error; err != nil {
		fmt.Println(fmt.Sprintf("locType create error :%v", err))
	}
	LocTypeCount++
	return locType
}
