package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

var SpuCount int

func CreateSpu(brandId, cateId, skuNum int, name, subTitle string, spec *SpecGroup) *ShopSpus {
	if len(name) == 0 {
		name = Fake.Name()
	}
	if len(subTitle) == 0 {
		subTitle = Fake.Paragraph(1, true)
	}
	sup := &ShopSpus{
		Name:      name,
		SubTitle:  subTitle,
		BrandID:   brandId,
		Status:    true,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		IsDeleted: sql.NullTime{},
	}
	if err := DB.Create(sup).Error; err != nil {
		fmt.Println(fmt.Sprintf("sup create error :%v", err))
	}

	sup.Detail = CreateSpuDetail(sup.ID)

	var specs []*ShopSpuSpecs
	sup.Specs = GetSpecs(sup.ID, specs, spec)

	var skus []*ShopSkus
	for skuNum > 0 {
		skus = append(skus, CreateSku(sup.ID, name))
		skuNum--
	}
	sup.Skus = skus
	cateSpu := ShopCateSpu{SpuID: sup.ID, CateID: cateId, UpdateAt: time.Now(), CreateAt: time.Now()}
	if err := DB.Table("shop_cate_spu").Create(&cateSpu).Error; err != nil {
		fmt.Println(fmt.Sprintf("cateSpu create error :%v", err))
	}

	SpuCount++
	return sup
}

func GetSpecs(supId int, specs []*ShopSpuSpecs, specGroup *SpecGroup) []*ShopSpuSpecs {
	gsps := GetSpecParams()
	for _, gsp := range gsps {
		if specGroup != nil {
			for _, sp := range specGroup.Params {
				if sp.Name == gsp.Name {
					gs, _ := json.Marshal(sp.Value)
					spec := CreateSpuSpec(gsp.ID, supId, gsp.Generic, string(gs))
					specs = append(specs, spec)
				}
			}
		}
	}

	return specs
}
