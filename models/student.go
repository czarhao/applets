package models

import (
	"applets/system"
)

type Student struct {
	Id       int    `xorm:"not null pk autoincr comment('学生id') INT(11)"`
	Name     string `xorm:"not null comment('学生姓名') VARCHAR(20)"`
	No       string `xorm:"not null comment('学生学号') CHAR(8)"`
	Password string `xorm:"not null comment('教务网密码') VARCHAR(64)"`
	College  string `xorm:"not null comment('学院') VARCHAR(20)"`
	Class    string `xorm:"not null comment('班级') VARCHAR(20)"`
	Prof     string `xorm:"not null comment('专业') VARCHAR(20)"`
	Sex      string `xorm:"not null comment('性别') VARCHAR(4)"`
	Birth    string `xorm:"not null comment('生日') VARCHAR(20)"`
	Grade    string `xorm:"not null comment('年级') VARCHAR(10)"`
}

// 判断学号是否存在
func (db *DbController) StudentExit(no string) (bool, *Student) {
	usernameCheck := &Student{No: no}
	has, err := db.eng.Get(usernameCheck)
	if err != nil {
		system.Save.DbError("StudentExit()", err)
		return false, nil
	}
	if !has {
		return false, nil
	}
	return true, usernameCheck
}

// 更新学生信息
func (db *DbController) StudentUpdate(student Student) (bool, error, int) {
	affected, err := db.eng.In("no", student.No).Update(student)
	if affected != 1 || err != nil {
		system.Save.DbError("StudentUpdate()", err)
		return false, err, 0
	}
	return true, nil, student.Id
}

// 创建学生信息
func (db *DbController) StudentCreate(student Student) (bool, int, error) {
	if _, err := db.eng.Insert(student); err != nil {
		system.Save.DbError("StudentCreate()", err)
		return false, 0, err
	}
	returnId := &Student{
		No: student.No,
	}
	if _, err := db.eng.Get(returnId); err != nil {
		system.Save.DbError("StudentCreate()", err)
	}
	return true, returnId.Id, nil
}
