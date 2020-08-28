package model

import (
	"database/sql"
	"fmt"
	"time"
)

var OrderCount int

func CreateOrderAddr(orderId int) *APIOOrderAddrs {
	orderAddr := APIOOrderAddrs{
		Name:         "name",
		Phone:        "13412569874",
		Addr:         "",
		Sex:          1,
		OOrderID:     orderId,
		HospitalName: "HospitalName",
		LocName:      "LocName",
		BedNum:       "BedNum",
		HospitalNo:   "9556854545",
		Disease:      "Disease",
		Age:          10,
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
		IsDeleted:    sql.NullTime{},
	}
	if err := DB.Create(&orderAddr).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderAddr create error :%v", err))
	}

	return &orderAddr
}

var OrderMenuCount int

func CreateOrderMenu(menu *APIMenus, menuType string, orderId int) *APIOOrderMenus {
	orderMenu := APIOOrderMenus{
		MenuName:     menu.Name,
		MenuType:     menuType,
		MenuTimeType: menu.TimeType.String(),
		MenuDesc:     menu.Desc,
		MenuID:       menu.ID,
		Price:        menu.Price,
		Amount:       menu.Amount,
		Cover:        menu.Cover,
		OOrderID:     orderId,
		CreateAt:     time.Now(),
		UpdateAt:     time.Now(),
		IsDeleted:    sql.NullTime{},
	}
	if err := DB.Create(&orderMenu).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderAddr create error :%v", err))
	}
	OrderMenuCount++
	return &orderMenu
}

func CreateOrderMenus(orderMenu *APIOOrderMenus) []*APIOOrderMenus {
	return []*APIOOrderMenus{orderMenu}
}

func CreateOrderComment(userId, orderId int, commentableType string) *Comments {
	orderComment := &Comments{
		Content:         Fake.Paragraph(1, true),
		UserID:          userId,
		CommentableID:   orderId,
		CommentableType: commentableType,
		ApplicationID:   AppId,
		Pics:            Pics,
		Star:            5,
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
		IsDeleted:       sql.NullTime{},
	}
	if err := DB.Create(orderComment).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderComment create error :%v", err))
	}
	return orderComment
}

func CreateOrder(orderNo string, userId, appType, payType int) *APIOOrders {
	order := APIOOrders{
		OrderNo:       orderNo,
		Status:        I_ORDER_STATUS_FOR_DELIVERY,
		Amount:        0,
		Total:         10.00,
		ApplicationID: AppId,
		PayType:       payType,
		Rmk:           "备注",
		AppType:       appType,
		OpenID:        "",
		TradeType:     "",
		IsReturn:      false,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		IsDeleted:     sql.NullTime{},
		UserID:        userId,
	}
	if err := DB.Create(&order).Error; err != nil {
		fmt.Println(fmt.Sprintf("order create error :%v", err))
	}

	OrderCount++
	return &order
}
