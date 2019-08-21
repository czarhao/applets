package api

import "applets/models"

// 将json转化为models

func infoToModel(tmp InfoJson) models.Student {
	return models.Student{
		Name:     tmp.Name,
		No:       tmp.Sno,
		Password: tmp.Spw,
		College:  tmp.College,
		Class:    tmp.Class,
		Prof:     tmp.Prof,
		Sex:      tmp.Sex,
		Birth:    tmp.Birth,
		Grade:    tmp.Grade,
	}
}

func GradeToModel(gradeJson GradeJson) models.Grade {
	return models.Grade{
		Grade:      gradeJson.Grade,
		Year:       gradeJson.Year,
		Semester:   gradeJson.Semester,
		CourseName: gradeJson.Cname,
		Credit:     gradeJson.Credit,
		Point:      gradeJson.Point,
		Usually:    gradeJson.Usually,
		Mid:        gradeJson.Mid,
		Final:      gradeJson.Final,
		Experiment: gradeJson.Experiment,
	}
}

func CourseJsonToModel(CourseJson CourseJson) models.Course {
	return models.Course{
		Id:          CourseJson.Id,
		Type:        CourseJson.Type,
		WeekStart:   CourseJson.Start,
		WeekEnd:     CourseJson.End,
		DayWeek:     CourseJson.DayWeek,
		DayStart:    CourseJson.CourStart,
		DayEnd:      CourseJson.CourStart + CourseJson.CourLength - 1,
		CourseName:  CourseJson.CourName,
		TeacherName: CourseJson.TeacherName,
		CourWhere:   CourseJson.CourWhere,
		Jud:         CourseJson.Jud,
	}
}