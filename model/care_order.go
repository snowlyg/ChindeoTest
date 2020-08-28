package model

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

var CareOrderCount int

func CreateCareOrder(timeType, orderNo string, userId, carerId, payType, appType int, price float64) *CareOrders {
	tt := time.Now()
	total := getTotalPrice(tt, timeType, price)

	order := &CareOrders{
		OrderNo:       orderNo,
		Status:        I_ORDER_STATUS_FOR_DELIVERY,
		Total:         total,
		ApplicationID: AppId,
		PayType:       payType,
		Rmk:           "备注",
		AppType:       appType,
		StartAt:       tt,
		EndAt:         tt.AddDate(0, 0, 1),
		OpenID:        "",
		TradeType:     "",
		IsReturn:      false,
		CreateAt:      tt,
		UpdateAt:      tt,
		IsDeleted:     sql.NullTime{},
		UserID:        userId,
		CarerID:       carerId,
	}
	if err := DB.Create(order).Error; err != nil {
		fmt.Println(fmt.Sprintf("order create error :%v", err))
	}

	CareOrderCount++
	return order
}

func getTotalPrice(tt time.Time, timeType string, price float64) float64 {

	var dd decimal.Decimal
	sub := int(tt.AddDate(0, 0, 1).Sub(tt).Hours())
	maxCarePrice := decimal.NewFromFloat(price)
	if timeType == "天" {
		dd = maxCarePrice.Mul(decimal.NewFromFloat(float64(sub / 24)))
	} else if timeType == "时" {
		dd = maxCarePrice.Mul(decimal.NewFromFloat(float64(sub)))
	}
	total, _ := dd.Float64()
	return total
}

var CareOrderCommentCount int

func CreateCareOrderComment(userId, orderId, star int, commentableType string) *Comments {
	orderComment := &Comments{
		Content:         Fake.Paragraph(1, true),
		UserID:          userId,
		CommentableID:   orderId,
		CommentableType: commentableType,
		ApplicationID:   AppId,
		Pics:            Pics,
		Star:            star,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := DB.Create(orderComment).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderComment create error :%v", err))
	}
	CareOrderCommentCount++
	return orderComment
}

func CreateCareOrderInfo(orderId int, care *Cares, appName, careTypeName, careTagName string) *CareOrderInfos {
	orderInfo := &CareOrderInfos{
		Name:            care.Name,
		Desc:            care.Desc,
		TimeType:        care.TimeType,
		CareType:        careTypeName,
		CareTags:        careTagName,
		MinPrice:        care.MinPrice,
		MaxPrice:        care.MaxPrice,
		CareDetail:      care.CareDetail,
		ApplicationName: appName,
		Cover:           care.Cover,
		CareOrderID:     orderId,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := DB.Create(orderInfo).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderInfo create error :%v", err))
	}
	return orderInfo
}

func CreateCareOrderAddr(orderID int, addr *APIAddrs) *CareOrderAddrs {
	orderAddr := &CareOrderAddrs{
		Name:         addr.Name,
		Phone:        addr.Phone,
		Addr:         addr.Addr,
		Sex:          int(addr.Sex),
		CareOrderID:  orderID,
		HospitalName: addr.HospitalName,
		LocName:      addr.LocName,
		BedNum:       addr.BedNum,
		HospitalNo:   addr.HospitalNo,
		Disease:      addr.Disease,
		Age:          addr.Age,
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
		IsDeleted:    sql.NullTime{},
	}
	if err := DB.Create(orderAddr).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderAddr create error :%v", err))
	}

	return orderAddr
}

func CreateCareOrderCarerInfo(orderId int, carer *Carers, AppName, carerTagName string) *CareOrderCarerInfos {
	orderCarerInfo := &CareOrderCarerInfos{
		Name:            carer.Name,
		Desc:            carer.Desc,
		TimeType:        carer.TimeType,
		CarerTags:       carerTagName,
		Price:           carer.Price,
		Age:             carer.Age,
		WorkExp:         carer.WorkExp,
		Sex:             carer.Sex,
		Phone:           carer.Phone,
		CarerDetail:     carer.CarerDetail,
		ApplicationName: AppName,
		Avatar:          carer.Avatar,
		CareOrderID:     orderId,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := DB.Create(orderCarerInfo).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderInfo create error :%v", err))
	}
	return orderCarerInfo
}
