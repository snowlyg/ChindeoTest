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

	sup.Detail = CreateSpuDetail(sup.ID, GetGenericSpec(spec), GetSpecialSpec(spec))

	var skus []*ShopSkus
	for skuNum > 0 {
		skus = append(skus, CreateSku(sup.ID))
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

func GetGenericSpec(spec *SpecGroup) string {
	genericSpec := make(GenericSpec)
	if spec != nil {
		gsps := GetGenericSpecParams()
		for _, gsp := range gsps {
			for _, sp := range spec.Params {
				if sp.Name == gsp.Name {
					genericSpec[gsp.ID] = sp.Value
				}
			}
		}
	}
	gs, _ := json.Marshal(genericSpec)
	return string(gs)
}

func GetSpecialSpec(spec *SpecGroup) string {
	specialSpec := make(SpecialSpec)
	if spec != nil {
		gsps := GetNotGenericSpecParams()
		for _, gsp := range gsps {
			for _, sp := range spec.Params {
				if sp.Name == gsp.Name {
					specialSpec[gsp.ID] = sp.Value
				}
			}
		}
	}
	ss, _ := json.Marshal(specialSpec)
	return string(ss)
}
