package model

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

var LocCount int

func CreateLoc() *OnlineLocs {
	loc := &OnlineLocs{
		LocID:         rand.Int(),
		CtHospitalID:  rand.Int(),
		LocWardFlag:   rand.Int(),
		LocActiveFlag: rand.Int(),
		LocDesc:       Fake.Paragraph(1, true),
		ApplicationID: AppId,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}
	if err := DB.Create(loc).Error; err != nil {
		fmt.Println(fmt.Sprintf("loc create error :%v", err))
	}
	LocCount++
	return loc
}
