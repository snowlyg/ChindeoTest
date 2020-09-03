package model

import (
	"database/sql"
	"fmt"
	"time"
)

var SpecParamCount int

func CreateSpecParam(groupCateId, groupId int, name, unit string, numeric, generic, searching bool) *ShopSpecParams {
	specParam := &ShopSpecParams{
		Name:        name,
		CateID:      groupCateId,
		SpecGroupID: groupId,
		Numeric:     numeric,
		Unit:        unit,
		Generic:     generic,
		Searching:   searching,
		Segments:    "0.1-0.5",
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
		IsDeleted:   sql.NullTime{},
	}
	if err := DB.Create(specParam).Error; err != nil {
		fmt.Println(fmt.Sprintf("specParam create error :%v", err))
	}

	SpecParamCount++
	return specParam
}

func GetGenericSpecParams() []*ShopSpecParams {
	var sp []*ShopSpecParams
	if err := DB.Where("generic = ?", true).Find(&sp).Error; err != nil {
		fmt.Println(fmt.Sprintf("GetUserByIdCardNo create error :%v", err))
	}
	return sp
}

func GetNotGenericSpecParams() []*ShopSpecParams {
	var sp []*ShopSpecParams
	if err := DB.Where("generic = ?", false).Find(&sp).Error; err != nil {
		fmt.Println(fmt.Sprintf("GetUserByIdCardNo create error :%v", err))
	}
	return sp
}
