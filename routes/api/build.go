package api

import (
	"applets/models"
	"applets/system"
)
// 将models转化成json

func buildSchedulesJson(success bool, errInfo string, schedule [7][6]ScheduleJson) SchedulesJson {
	return SchedulesJson{
		Success:    success,
		Info:       errInfo,
		SchoolYear: system.SemesterConf.SchoolYear,
		TeachWeek:  system.SemesterConf.TeachWeek,
		Semester:   system.SemesterConf.Semester,
		Schedule:   schedule,
	}
}

func buildInfoJson(success bool, info string, studentInfo models.Student) InfoJson {
	return InfoJson{
		Success: success,
		Info:    info,
		Name:    studentInfo.Name,
		Spw:     studentInfo.Password,
		Grade:   studentInfo.Grade,
		Sno:     studentInfo.No,
		College: studentInfo.College,
		Class:   studentInfo.Class,
		Prof:    studentInfo.Prof,
		Sex:     studentInfo.Sex,
		Birth:   studentInfo.Birth,
	}
}

func buildArticlesJson(success bool, info string, article []ArticleJson) ArticlesJson {
	return ArticlesJson{
		Success:  success,
		Info:     info,
		Articles: article,
	}
}

func buildArticleJson(articles []models.Article) []ArticleJson {
	var returnArticles []ArticleJson
	for _, value := range articles {
		returnArticles = append(returnArticles, ArticleJson{
			Date:       value.Date,
			Title:      value.Title,
			ImgSrc:     value.Img,
			AuthorImg:  value.Head,
			Content:    value.Content,
			AuthorName: value.Author,
			Passage:    value.Passage,
			PostId:     value.Id,
		})
	}
	return returnArticles
}
