package model

import (
	"database/sql"
	"fmt"
	"time"
)

func CreateIndexRelationShip(code, name string) *IndexRelationships {
	ir := &IndexRelationships{
		Code:      code,
		Name:      name,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		IsDeleted: sql.NullTime{},
	}
	if err := DB.Create(ir).Error; err != nil {
		fmt.Println(fmt.Sprintf("ir create error :%v", err))
	}
	return ir
}

func CreateIndexRelationShips() {
	irs := []map[string]string{
		{
			"code": "parent",
			"name": "父母",
		}, {
			"code": "children",
			"name": "子女",
		}, {
			"code": "brother",
			"name": "兄弟",
		}, {
			"code": "sister",
			"name": "姐妹",
		}, {
			"code": "couple",
			"name": "夫妻",
		},
	}

	for _, ir := range irs {
		CreateIndexRelationShip(ir["code"], ir["name"])
	}
}
