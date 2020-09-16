package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

func CreateSku(supId int, name string, price float64) *ShopSkus {
	images := []string{"https://wx.qlogo.cn/mmopen/vi_32/08uGpO07ibo8JrANxTDWjPbZHOtZsSkFfTNRksMtibzSRjcQ76xewGZra1AgnKz2EstBFueFcJIJCqucmJAGPgNw/132", "https://wx.qlogo.cn/mmopen/vi_32/08uGpO07ibo8JrANxTDWjPbZHOtZsSkFfTNRksMtibzSRjcQ76xewGZra1AgnKz2EstBFueFcJIJCqucmJAGPgNw/132", "https://wx.qlogo.cn/mmopen/vi_32/08uGpO07ibo8JrANxTDWjPbZHOtZsSkFfTNRksMtibzSRjcQ76xewGZra1AgnKz2EstBFueFcJIJCqucmJAGPgNw/132"}
	i, _ := json.Marshal(images)
	sku := &ShopSkus{
		Title:     name,
		SkuNo:     Fake.PostCode(),
		Weight:    10,
		Images:    string(i),
		Stock:     100,
		Price:     price,
		Indexes:   "[\"白色\", \"3GB\",\"16GB\"]",
		OwnSpec:   "[\"白色\", \"3GB\",\"16GB\"]",
		Status:    true,
		SpuID:     supId,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		IsDeleted: sql.NullTime{},
	}
	if err := DB.Create(sku).Error; err != nil {
		fmt.Println(fmt.Sprintf("sku create error :%v", err))
	}

	return sku
}
