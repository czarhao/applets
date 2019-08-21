package api

import "applets/models"

// 发送的课程表json
type SchedulesJson struct {
	Success    bool               `json:"success"`
	Info       string             `json:"info"`
	SchoolYear string             `json:"school_year"`
	Semester   string             `json:"semester"`
	TeachWeek  int                `json:"teach_week"`
	Schedule   [7][6]ScheduleJson `json:"schedule"`
}

type ScheduleJson struct {
	Id          int    `json:"id"`
	Type        int    `json:"type"`
	Start       int    `json:"start"`
	End         int    `json:"end"`
	CourName    string `json:"cour_name"`
	TeacherName string `json:"teacher_name"`
	CourWhere   string `json:"cour_where"`
}

// 获取的课程表json
type CoursesJson struct {
	Success  bool         `json:"success"`
	Info     string       `json:"info"`
	Schedule []CourseJson `json:"courses"`
}

type CourseJson struct {
	Id          int    `json:"id"`
	Type        int    `json:"type"`
	Start       int    `json:"start"`
	End         int    `json:"end"`
	DayWeek     int    `json:"day_week"`
	CourStart   int    `json:"cour_start"`
	CourLength  int    `json:"cour_length"`
	CourName    string `json:"cour_name"`
	TeacherName string `json:"teacher_name"`
	CourWhere   string `json:"cour_where"`
	Jud         int    `json:"jud"`
}

type RegisteredJson struct {
	Success    bool        `json:"success"`
	Info       string      `json:"info"`
	InfoJson   InfoJson    `json:"info_json"`
	CourseJson CoursesJson `json:"schedule_json"`
	GradeJson  GradesJson  `json:"grade_json"`
}

type InfoJson struct {
	Success bool   `json:"success"`
	Info    string `json:"info"`
	Name    string `json:"name"`
	Sno     string `json:"sno"`
	Spw     string `json:"spw"`
	College string `json:"college"`
	Class   string `json:"class"`
	Prof    string `json:"prof"`
	Sex     string `json:"sex"`
	Birth   string `json:"birth"`
	Grade   string `json:"grade"`
}

type GradesJson struct {
	Success bool        `json:"success"`
	Info    string      `json:"info"`
	Grades  []GradeJson `json:"grade"`
}

type GradeJson struct {
	Year       string  `json:"year"`
	Semester   int     `json:"semester"`
	Cname      string  `json:"cname"`
	Credit     float64 `json:"credit"`
	Point      float64 `json:"point"`
	Usually    string  `json:"usually"`
	Mid        string  `json:"mid"`
	Final      string  `json:"final"`
	Experiment string  `json:"experiment"`
	Grade      string  `json:"grade"`
}

func CourseToSchedule(CourseJson CourseJson) ScheduleJson {
	return ScheduleJson{
		Id:          CourseJson.Id,
		Type:        CourseJson.Type,
		Start:       CourseJson.Start,
		End:         CourseJson.End,
		CourName:    CourseJson.CourName,
		TeacherName: CourseJson.TeacherName,
		CourWhere:   CourseJson.CourWhere,
	}
}

func buildGradesJson(success bool, info string, grades []GradeJson) GradesJson {
	return GradesJson{
		Success: success,
		Info:    info,
		Grades:  grades,
	}
}

func GradeModelToJson(grade models.Grade) GradeJson {
	return GradeJson{
		Year:       grade.Year,
		Semester:   grade.Semester,
		Cname:      grade.CourseName,
		Credit:     grade.Credit,
		Point:      grade.Point,
		Usually:    grade.Usually,
		Mid:        grade.Mid,
		Final:      grade.Final,
		Experiment: grade.Experiment,
		Grade:      grade.Grade,
	}
}

type ArticlesJson struct {
	Success  bool          `json:"success"`
	Info     string        `json:"info"`
	Articles []ArticleJson `json:"articles"`
}

type ArticleJson struct {
	Date       string `json:"date"`
	Title      string `json:"title"`
	ImgSrc     string `json:"img_src"`
	AuthorImg  string `json:"author_img"`
	Content    string `json:"content"`
	AuthorName string `json:"author_name"`
	Passage    string `json:"passage"`
	PostId     int    `json:"post_id"`
}
