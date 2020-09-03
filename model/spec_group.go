package model

import (
	"database/sql"
	"fmt"
	"time"
)

var SpecGroupCount int

type SpecGroup struct {
	Name   string
	Params []*SpecParam
}

type SpecParam struct {
	Name      string
	Unit      string
	Numeric   bool
	Generic   bool
	Searching bool
	Value     []string
}

func CreateSpecGroup(cateId int, spec *SpecGroup) *ShopSpecGroups {
	specGroup := &ShopSpecGroups{
		Name:      spec.Name,
		CateID:    cateId,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		IsDeleted: sql.NullTime{},
	}

	if err := DB.Create(specGroup).Error; err != nil {
		fmt.Println(fmt.Sprintf("specGroup create error :%v", err))
	}

	var specParams []*ShopSpecParams
	for _, param := range spec.Params {
		specParams = append(specParams, CreateSpecParam(specGroup.CateID, specGroup.ID, param.Name, param.Unit, param.Numeric, param.Generic, param.Searching))
	}

	specGroup.SpecParams = specParams

	SpecGroupCount++
	return specGroup
}
