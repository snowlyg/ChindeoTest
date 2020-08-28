package model

import (
	"database/sql"
	"fmt"
	"time"
)

var ArticleCount int

func CreateArticle() *Articles {
	artice := &Articles{
		Title:            Fake.JobTitle(),
		Digest:           Fake.Paragraph(1, true),
		Content:          Fake.Paragraph(1, true),
		Author:           Fake.Name(),
		LocalURL:         Fake.URL(),
		ContentSourceURL: Fake.URL(),
		CreateAt:         time.Now(),
		UpdateAt:         time.Now(),
		IsDeleted:        sql.NullTime{},
	}
	if err := DB.Create(&artice).Error; err != nil {
		fmt.Println(fmt.Sprintf("artice create error :%v", err))
	}
	ArticleCount++
	return artice
}
