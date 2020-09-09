package model

import (
	"database/sql"
	"fmt"
	"time"
)

func CreateSpuDetail(supId int) *ShopSpuDetails {
	supDetail := &ShopSpuDetails{
		Description:  Fake.Paragraph(1, true),
		PackingList:  Fake.Paragraph(1, true),
		AfterService: Fake.Paragraph(1, true),
		SpuID:        supId,
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
		IsDeleted:    sql.NullTime{},
	}
	if err := DB.Create(supDetail).Error; err != nil {
		fmt.Println(fmt.Sprintf("supDetail create error :%v", err))
	}

	return supDetail
}
