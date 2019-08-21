package models

import "applets/system"

type Grade struct {
	Id         int     `xorm:"not null pk autoincr comment('成绩id') INT(11)"`
	Year       string  `xorm:"not null comment('学年') VARCHAR(24)"`
	Semester   int     `xorm:"not null comment('学期') INT(11)"`
	CourseName string  `xorm:"not null comment('课程名字') VARCHAR(96)"`
	Credit     float64 `xorm:"not null comment('学分') DOUBLE"`
	Point      float64 `xorm:"not null comment('绩点') DOUBLE"`
	Usually    string  `xorm:"not null comment('平时成绩') VARCHAR(8)"`
	Mid        string  `xorm:"not null comment('期中成绩') VARCHAR(8)"`
	Final      string  `xorm:"not null comment('期末成绩') VARCHAR(8)"`
	Experiment string  `xorm:"not null comment('实验成绩') VARCHAR(8)"`
	Grade      string  `xorm:"not null comment('总成绩') VARCHAR(8)"`
	Sid        int     `xorm:"not null comment('学生学号') INT(11)"`
}

// 判断课程是否存在在数据库中
func (db *DbController) GradeExit(cname string, sid int) (bool, int) {
	gradeCheck := &Grade{
		CourseName: cname,
		Sid:        sid,
	}
	has, err := db.eng.Get(gradeCheck)
	if err != nil {
		system.Save.DbError("GradeExit()", err)
		return false, 0
	}
	if !has {
		return false, 0
	}
	return true, gradeCheck.Id
}

// 创建一个课程
func (db *DbController) GradeCreate(sid int, grade Grade) (bool, error) {
	grade.Sid = sid
	if _, err := db.eng.Insert(grade); err != nil {
		system.Save.DbError("GradeCreate()", err)
		return false, err
	}
	return true, nil
}

// 根据学号返回成绩
func (db *DbController) GradeReturn(sid int) []Grade {
	returnGrades := make([]Grade, 0)
	if err := db.eng.Where("sid = ?", sid).Find(&returnGrades); err != nil {
		system.Save.DbError("GradeReturn()", err)
	}
	return returnGrades
}

// 删除成绩
func (db *DbController) GradeDelete(sid int) bool {
	tmpDelete := Grade{Sid: sid}
	if _, err := db.eng.Where("sid=?", sid).Delete(tmpDelete); err != nil {
		system.Save.DbError("GradeDelete()", err)
		return false
	}
	return true
}
