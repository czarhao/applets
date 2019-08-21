package models

import "applets/system"

type Course struct {
	Id          int    `xorm:"not null pk autoincr comment('课程id') INT(11)"`
	Time        int    `xorm:"not null comment('上课时间') INT(11)"`
	Type        int    `xorm:"not null comment('课程类型') INT(11)"`
	WeekStart   int    `xorm:"not null comment('第几周开始') INT(11)"`
	WeekEnd     int    `xorm:"not null comment('第几周结束') INT(11)"`
	DayWeek     int    `xorm:"not null comment('周几上课') INT(11)"`
	DayStart    int    `xorm:"not null comment('第几节开始') INT(11)"`
	DayEnd      int    `xorm:"not null comment('第几节结束') INT(11)"`
	CourseName  string `xorm:"not null comment('课程名字') VARCHAR(96)"`
	TeacherName string `xorm:"not null comment('教师名字') VARCHAR(24)"`
	CourWhere   string `xorm:"not null comment('上课地点') VARCHAR(24)"`
	Jud         int    `xorm:"not null comment('单双周') INT(11)"`
}

// 课程是否存在在数据库中
func (db *DbController) CourseExit(cname string, courStart int, dayWeek int, tName string, start int, end int) (bool, int) {
	cnameCheck := &Course{CourseName: cname, DayStart: courStart, DayWeek: dayWeek, TeacherName: tName, WeekStart: start, WeekEnd: end}
	has, err := db.eng.Get(cnameCheck)
	if err != nil {
		system.Save.DbError("CourseExit()", err)
		return false, 0
	}
	if !has {
		return false, 0
	}
	return true, cnameCheck.Id
}

// 将课程存入数据库中
func (db *DbController) CourseSave(course Course) (bool, error, int) {
	if _, err := db.eng.Insert(course); err != nil {
		system.Save.DbError("CourseSave()", err)
		return false, err, 0
	}
	returnCourse := &Course{
		Time:        course.Time,
		Type:        course.Type,
		DayWeek:     course.DayWeek,
		WeekStart:   course.WeekStart,
		WeekEnd:     course.WeekEnd,
		DayStart:    course.DayStart,
		DayEnd:      course.DayEnd,
		TeacherName: course.TeacherName,
		CourWhere:   course.CourWhere,
		Jud:         course.Jud,
	}
	if _, err := db.eng.Get(returnCourse); err != nil {
		system.Save.DbError("CourseSave()", err)
		return false, err, 0
	}
	return true, nil, returnCourse.Id
}

// 删除已经有的绑定关系
func (db *DbController) DeleteBand(sid int) bool {
	tmpDelete := StudentCourse{Sid: sid}
	if _, err := db.eng.Where("sid=?", sid).Delete(tmpDelete); err != nil {
		system.Save.DbError("DeleteBand()", err)
		return false
	}
	return true
}
