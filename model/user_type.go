package model

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

var UserTypeCount int

func CreateUserType() *UserTypes {
	letters := []byte("NDO")
	userType := &UserTypes{
		UtpID:       rand.Int(),
		UtpCode:     Fake.CellPhoneNumber(),
		UtpDesc:     Fake.Paragraph(1, true),
		UtpType:     string(letters[rand.Intn(2)]),
		UtpActive:   true,
		UtpContrast: Fake.Paragraph(1, true),
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
		IsDeleted:   sql.NullTime{},
	}
	if err := DB.Create(userType).Error; err != nil {
		fmt.Println(fmt.Sprintf("userType create error :%v", err))
	}
	UserTypeCount++
	return userType
}
