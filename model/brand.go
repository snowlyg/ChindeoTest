package model

import (
	"database/sql"
	"fmt"
	"time"
)

var BrandCount int
var BrandWithCateCount int

func CreateBrand(isCate bool) *ShopBrands {
	brand := &ShopBrands{
		Name:      Fake.Paragraph(1, true),
		Image:     Fake.URL(),
		Letter:    Fake.Words(1, true)[0],
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		IsDeleted: sql.NullTime{},
	}
	if err := DB.Create(brand).Error; err != nil {
		fmt.Println(fmt.Sprintf("brand create error :%v", err))
	}

	if isCate {
		var cates []*ShopCates
		cate := CreateCate(0, 1)
		brandCate := ShopBrandCate{BrandID: brand.ID, CateID: cate.ID, UpdateAt: time.Now(), CreateAt: time.Now()}
		if err := DB.Table("shop_brand_cate").Create(&brandCate).Error; err != nil {
			fmt.Println(fmt.Sprintf("brandCate create error :%v", err))
		}
		cates = append(cates, cate)
		brand.Cates = cates
		BrandWithCateCount++
	}

	BrandCount++

	return brand
}
