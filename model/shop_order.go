package model

import (
	"database/sql"
	"fmt"
	"time"
)

var ShopOrderCount int

func CreateShopOrder(orderNo string, userId, payType, appType int, skus []*ShopSkus) *ShopOrders {
	shopOrder := &ShopOrders{
		OrderNo:       orderNo,
		Status:        IOrderStatusForDelivery,
		PayType:       payType,
		Rmk:           "备注",
		AppType:       appType,
		OpenID:        "",
		TradeType:     "",
		IsReturn:      false,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
		ApplicationID: AppId,
		UserID:        userId,
	}
	if err := DB.Create(shopOrder).Error; err != nil {
		fmt.Println(fmt.Sprintf("shopOrder create error :%v", err))
	}

	addr := CreateAddr(userId)
	shopOrder.Addr = CreateShopOrderAddr(shopOrder.ID, addr)
	shopOrder.Skus = CreateShopOrderSkus(shopOrder.ID, userId, skus)

	ShopOrderCount++
	return shopOrder
}

func CreateShopOrderAddr(orderID int, addr *APIAddrs) *ShopOrderAddrs {
	shopOrderAddr := &ShopOrderAddrs{
		Name:          addr.Name,
		Phone:         addr.Phone,
		Addr:          addr.Addr,
		Sex:           int(addr.Sex),
		ShopOrderID:   orderID,
		ApplicationID: AppId,
		HospitalName:  addr.HospitalName,
		LocName:       addr.LocName,
		BedNum:        addr.BedNum,
		HospitalNo:    addr.HospitalNo,
		Disease:       addr.Disease,
		Age:           addr.Age,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}
	if err := DB.Create(shopOrderAddr).Error; err != nil {
		fmt.Println(fmt.Sprintf("shopOrderAddr create error :%v", err))
	}

	return shopOrderAddr
}

func CreateShopOrderSku(orderID, userId int, sku *ShopSkus) *ShopOrderSkus {
	shopOrderSku := &ShopOrderSkus{
		Title:         sku.Title,
		SkuNo:         sku.SkuNo,
		Images:        sku.Images,
		Price:         sku.Price,
		Indexes:       sku.Indexes,
		Num:           2,
		ShopOrderID:   orderID,
		SkuID:         sku.ID,
		ApplicationID: AppId,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}
	if err := DB.Create(shopOrderSku).Error; err != nil {
		fmt.Println(fmt.Sprintf("shopOrderSku create error :%v", err))
	}
	shopOrderSku.Comments = CreateShopOrderComments(userId, shopOrderSku.ID, 5, 4, sku)
	return shopOrderSku
}

func CreateShopOrderSkus(orderID, userId int, skus []*ShopSkus) []*ShopOrderSkus {
	var shopSkus []*ShopOrderSkus
	for _, sku := range skus {
		shopSku := CreateShopOrderSku(orderID, userId, sku)
		shopSkus = append(shopSkus, shopSku)
	}
	return shopSkus
}

func CreateShopOrderComment(userId, shopSkuId, star int, sku *ShopSkus) *ShopOrderComments {
	shopOrderComment := &ShopOrderComments{
		UserID:        userId,
		ShopSkuID:     shopSkuId,
		SkuID:         sku.ID,
		Content:       Fake.Paragraph(1, true),
		Pics:          GetPics(),
		Star:          star,
		ApplicationID: AppId,
		IsUnnamed:     false,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
	}
	if err := DB.Create(shopOrderComment).Error; err != nil {
		fmt.Println(fmt.Sprintf("shopOrderComment create error :%v", err))
	}

	return shopOrderComment
}

func CreateShopOrderComments(userId, shopSkuId, star, num int, sku *ShopSkus) []*ShopOrderComments {
	var comments []*ShopOrderComments
	for num > 0 {
		comment := CreateShopOrderComment(userId, shopSkuId, star, sku)
		comments = append(comments, comment)
		num--
	}
	return comments
}
