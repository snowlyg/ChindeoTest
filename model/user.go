package model

import (
	"fmt"
)

func GetUserByIdCardNo() *APIUsers {
	user := APIUsers{}
	if err := DB.Where("id_card_no = ?", IdCardNo).First(&user).Error; err != nil {
		fmt.Println(fmt.Sprintf("GetUserByIdCardNo create error :%v", err))
	}
	return &user
}
