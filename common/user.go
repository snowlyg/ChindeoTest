package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/snowlyg/ChindeoTest/model"
)

var User *model.APIUsers

func GetUserByIdCardNo(db *gorm.DB) {
	user := model.APIUsers{}
	if err := db.Where("id_card_no = ?", IdCardNo).First(&user).Error; err != nil {
		fmt.Println(fmt.Sprintf("menuType create error :%v", err))
	}
	User = &user
}
