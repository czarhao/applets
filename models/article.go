package models

import (
	"applets/system"
)

type Article struct {
	Id      int    `xorm:"not null pk autoincr comment('文章id') INT(11)"`
	Date    string `xorm:"not null comment('创建时间') VARCHAR(24)"`
	Title   string `xorm:"not null comment('标题') VARCHAR(100)"`
	Head    string `xorm:"not null comment('作者头像') VARCHAR(100)"`
	Img     string `xorm:"not null comment('标题图片') VARCHAR(100)"`
	Content string `xorm:"not null comment('内容简介') TEXT"`
	Author  string `xorm:"not null comment('作者') VARCHAR(100)"`
	Passage string `xorm:"not null comment('文章内容') TEXT"`
}

func (db *DbController) GetArticle (start int) []Article {
	returnArticle := make([]Article, 0)
	if err := db.eng.Where("id > ? ", start).Where("id < ?", start + system.ArticleNum + 1).Find(&returnArticle); err != nil {
		system.Save.DbError("GetArticle()", err)
	}
	return returnArticle
}
