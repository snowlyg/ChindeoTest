package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

func CreateSku(supId int) *ShopSkus {
	images := []string{"https://wx.qlogo.cn/mmopen/vi_32/08uGpO07ibo8JrANxTDWjPbZHOtZsSkFfTNRksMtibzSRjcQ76xewGZra1AgnKz2EstBFueFcJIJCqucmJAGPgNw/132", "https://wx.qlogo.cn/mmopen/vi_32/08uGpO07ibo8JrANxTDWjPbZHOtZsSkFfTNRksMtibzSRjcQ76xewGZra1AgnKz2EstBFueFcJIJCqucmJAGPgNw/132", "https://wx.qlogo.cn/mmopen/vi_32/08uGpO07ibo8JrANxTDWjPbZHOtZsSkFfTNRksMtibzSRjcQ76xewGZra1AgnKz2EstBFueFcJIJCqucmJAGPgNw/132"}
	i, _ := json.Marshal(images)
	sku := &ShopSkus{
		Title:     Fake.Paragraph(1, true),
		SkuNo:     Fake.PostCode(),
		Weight:    10,
		Images:    string(i),
		Stock:     100,
		Price:     100.00,
		Indexes:   "{\n\t\"4\": [\"白色\", \"金色\", \"玫瑰金\"],\n\t\"12\": [\"3GB\"],\n\t\"13\": [\"16GB\"]\n}",
		OwnSpec:   "{\n\t\"4\": \"白色\",\n\t\"12\": \"3GB\",\n\t\"13\": \"16GB\"\n}",
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
