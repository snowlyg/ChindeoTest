package model

import (
	"database/sql"
	"fmt"
	"time"
)

func CreateIndexVitalConfig(code, name, unit, alarming, color, bgImag, icon, appIcon string) *IndexVitalConfigs {
	ir := &IndexVitalConfigs{
		Code:              code,
		Name:              name,
		Unit:              unit,
		AlarmingCondition: alarming,
		Color:             color,
		BgImg:             bgImag,
		Icon:              icon,
		AppIcon:           appIcon,
		Status:            true,
		Sort:              1,
		StartAt:           time.Now(),
		EndAt:             time.Now(),
		CreateAt:          time.Now(),
		UpdateAt:          time.Now(),
		IsDeleted:         sql.NullTime{},
	}
	if err := DB.Create(ir).Error; err != nil {
		fmt.Println(fmt.Sprintf("ir create error :%v", err))
	}
	return ir
}

func CreateIndexVitalConfigs() {
	irs := []map[string]string{
		{
			"code":     "01",
			"name":     "体温",
			"unit":     "℃",
			"alarming": "[{\"cond\":\">\", \"value\":37.3, \"logic\":\"&\"}]",
			"color":    "#33CCFF",
			"bgImag":   "",
			"icon":     "/static/img/index/ICON_T.png",
			"appIcon":  "/static/img/index/APP_ICON_T.png",
		}, {
			"code":     "02",
			"name":     "心率",
			"unit":     "次/分",
			"alarming": "[{\n \"cond\": \"<\",\n \"value\": 60,\n \"logic\": \"&\"\n}, {\n \"cond\": \">\",\n \"value\": 100,\n \"logic\": \"&\"\n}]",
			"color":    "#9966FF",
			"bgImag":   "",
			"icon":     "/static/img/index/ICON_HR.png",
			"appIcon":  "/static/img/index/APP_ICON_HR.png",
		}, {
			"code":     "03",
			"name":     "血压",
			"unit":     "mmhg",
			"alarming": "[{\n \"cond\": \"<\",\n \"value\": 60,\n \"logic\": \"&\"\n}, {\n \"cond\": \">\",\n \"value\": 140,\n \"logic\": \"&\"\n}]",
			"color":    "#FFCC33",
			"bgImag":   "",
			"icon":     "/static/img/index/ICON_BP.png",
			"appIcon":  "/static/img/index/APP_ICON_BP.png",
		}, {
			"code":     "04",
			"name":     "血糖",
			"unit":     "mmol/L",
			"alarming": "[{\n \"cond\": \"<\",\n \"value\": 3.9,\n \"logic\": \"&\"\n}, {\n \"cond\": \">\",\n \"value\": 6.1\"logic\": \"&\"\n}]",
			"color":    "#FF6699 ",
			"bgImag":   "",
			"icon":     "/static/img/index/ICON_BG.png",
			"appIcon":  "/static/img/index/APP_ICON_BG.png",
		}, {
			"code":     "05",
			"name":     "血脂",
			"unit":     "mmol/L",
			"alarming": "[{\n \"cond\": \"<\",\n \"value\": 2.8,\n \"logic\": \"&\"\n}, {\n \"cond\": \">\",\n \"value\": 5.17,\n \"logic\": \"&\"\n}]",
			"color":    "#00CC00",
			"bgImag":   "",
			"icon":     "/static/img/index/ICON_BF.png",
			"appIcon":  "/static/img/index/APP_ICON_BF.png",
		},
	}

	for _, ir := range irs {
		CreateIndexVitalConfig(ir["code"], ir["name"], ir["unit"], ir["alarming"], ir["color"], ir["bgImag"], ir["icon"], ir["appIcon"])
	}
}
