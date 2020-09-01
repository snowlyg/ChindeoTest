package model

import (
	"database/sql"
	"fmt"
	"time"
)

var ParentCateCount int
var CateCount int

func CreateCate(pid, level int) *ShopCates {
	cate := &ShopCates{
		Title:     Fake.Paragraph(1, true),
		Pid:       pid,
		Icon:      Fake.Name(),
		URL:       Fake.URL(),
		Sort:      0,
		Level:     level,
		Status:    true,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		IsDeleted: sql.NullTime{},
	}
	if err := DB.Create(cate).Error; err != nil {
		fmt.Println(fmt.Sprintf("cate create error :%v", err))
	}

	if pid == 0 {
		ParentCateCount++
	}
	CateCount++
	return cate
}
