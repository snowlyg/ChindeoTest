package common

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/snowlyg/ChindeoTest/model"
	"time"
)

var AddrCount int

func CreateAddr(db *gorm.DB) *model.APIAddrs {
	addr := model.APIAddrs{
		Name:         "name",
		Phone:        "13412569874",
		Addr:         "",
		Sex:          1,
		UserID:       User.ID,
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
	if err := db.Create(&addr).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuType create error :%v", err))
	}
	AddrCount++
	return &addr
}
