package model

import "fmt"

func GetApp() *APIApplications {
	app := &APIApplications{}
	if err := DB.Where("id = ?", AppId).First(app).Error; err != nil {
		fmt.Println(fmt.Sprintf("orderInfo create error :%v", err))
	}
	return app
}
