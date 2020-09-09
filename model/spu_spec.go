package model

import (
	"database/sql"
	"fmt"
	"time"
)

func CreateSpuSpec(specParamID, supId int, isGeneric bool, value string) *ShopSpuSpecs {
	supSpec := &ShopSpuSpecs{
		IsGeneric:   isGeneric,
		Value:       value,
		SpuID:       supId,
		SpecParamID: specParamID,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
		IsDeleted:   sql.NullTime{},
	}
	if err := DB.Create(supSpec).Error; err != nil {
		fmt.Println(fmt.Sprintf("supSpec create error :%v", err))
	}

	return supSpec
}
