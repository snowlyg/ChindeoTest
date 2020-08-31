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
	return &orderMenu
}

func CreateOrderMenus(menu *APIMenus, menuType string, orderId, num int) []*APIOOrderMenus {
	var menus []*APIOOrderMenus
	for num > 0 {
		menu := CreateOrderMenu(menu, menuType, orderId)
		menus = append(menus, menu)
		num--
	}
	return menus
}

func CreateOrderComment(userId, orderId int, commentableType string) *Comments {
	orderComment := &Comments{
		Content:         Fake.Paragraph(1, true),
		UserID:          userId,
		CommentableID:   orderId,
		CommentableType: commentableType,
		ApplicationID:   AppId,
		Pics:            GetPics(),
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

func CreateOrderComments(num, userId, orderId int, commentableType string) []*Comments {
	var comments []*Comments
	for num > 0 {
		comment := CreateOrderComment(userId, orderId, commentableType)
		comments = append(comments, comment)
		num--
	}
	return comments
}

func CreateOrder(orderNo string, userId, appType, payType int) *APIOOrders {
	order := APIOOrders{
		OrderNo:       orderNo,
		Status:        IOrderStatusForDelivery,
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
