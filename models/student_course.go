package models

import "applets/system"

type StudentCourse struct {
	Id  int `xorm:"not null pk autoincr comment('关联id') INT(11)"`
	Sid int `xorm:"not null comment('学生id') INT(11)"`
	Cid int `xorm:"not null comment('课程id') INT(11)"`
}

// 绑定课程用户
func (db *DbController) CourseBand(cid int, sid int) (bool, error) {
	band := &StudentCourse{
		Sid: sid,
		Cid: cid,
	}
	if _, err := db.eng.Insert(band); err != nil {
		system.Save.DbError("CourseBand()", err)
		return false, err
	}
	return true, nil
}

// 根据sid返回课程id
func (db DbController) SidGetGid(sid int) []StudentCourse {
	returnStudentInfo := make([]StudentCourse, 0)
	if err := db.eng.Where("sid = ?", sid).Find(&returnStudentInfo); err != nil {
		system.Save.DbError("SidGetGid()", err)
	}
	return returnStudentInfo
}

// 根据课程id返回课程
func (db *DbController) SidGetCourse(cid int) Course {
	course := Course{
		Id: cid,
	}
	if _, err := db.eng.Get(&course); err != nil {
		system.Save.DbError("SidGetCourse()", err)
	}
	return course
}
