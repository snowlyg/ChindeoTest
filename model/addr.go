package model

import (
	"database/sql"
	"fmt"
	"time"
)

var AddrCount int

func CreateAddr(userId int) *APIAddrs {
	addr := APIAddrs{
		Name:         "name",
		Phone:        "13412569874",
		Addr:         "",
		Sex:          1,
		UserID:       userId,
		HospitalName: "HospitalName",
		LocName:      "LocName",
		BedNum:       "BedNum",
		HospitalNo:   "9556854545",
		Disease:      "Disease",
		Age:          10,
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
		IsDeleted:    sql.NullTime{},
	}
	if err := DB.Create(&addr).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuType create error :%v", err))
	}
	AddrCount++
	return &addr
}
