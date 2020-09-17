package model

import (
	"database/sql"
	"fmt"
	"time"
)

func CreateIndexVitalTime(startAt, endAt time.Time) *IndexVitalTimes {
	ir := &IndexVitalTimes{
		Status:    true,
		StartAt:   startAt,
		EndAt:     endAt,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		IsDeleted: sql.NullTime{},
	}
	if err := DB.Create(ir).Error; err != nil {
		fmt.Println(fmt.Sprintf("ir create error :%v", err))
	}
	return ir
}

func CreateIndexVitalTimes() {
	t, _ := time.LoadLocation("Asia/Shanghai")
	n, _ := time.ParseInLocation("15:04", "09:00", t)
	nn, _ := time.ParseInLocation("15:04", "16:00", t)
	z, _ := time.ParseInLocation("15:04", "23:59", t)

	irs := []map[string]time.Time{
		{
			"start_at": z,
			"end_at":   n,
		},
		{
			"start_at": n,
			"end_at":   nn,
		},
		{
			"start_at": nn,
			"end_at":   z,
		},
	}

	for _, ir := range irs {
		CreateIndexVitalTime(ir["start_at"], ir["end_at"])
	}
}
