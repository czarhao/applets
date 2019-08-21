package api

import (
	"applets/models"
	"applets/system"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Schedule(ctx *gin.Context) {
	sno, spw, _ := getParameters(ctx)
	inDb, studentInfo := models.DbContr.StudentExit(sno)
	if inDb && studentInfo.Password == spw {
		// 相同直接返回数据库中的信息
		ctx.JSON(http.StatusOK, returnSortJson(studentInfo.Id))
	} else {
		if inDb == false {
			ctx.JSON(http.StatusOK, buildInfoJson(false, "同学尚未注册我们的服务。。。", models.Student{}))
		} else {
			ctx.JSON(http.StatusOK, buildInfoJson(false, "同学输入的密码错误！", models.Student{}))
		}
	}
}

// 刷新课程表
func updateSchedule(sno string, spw string) SchedulesJson {
	jud, student := models.DbContr.StudentExit(sno)
	if !jud || !models.DbContr.DeleteBand(student.Id) {
		return buildSchedulesJson(false, "可能是删除关系时出错了", [7][6]ScheduleJson{})
	}
	body, _ := getJson(system.CrawlServer + "schedule/" + sno + "/" + spw)
	coursesJson := CoursesJson{}
	_ = json.Unmarshal(body, &coursesJson)
	if !coursesJson.Success {
		return buildSchedulesJson(false, "课程获取出错！"+coursesJson.Info, [7][6]ScheduleJson{})
	}
	for _, course := range coursesJson.Schedule {
		jud, cid := models.DbContr.CourseExit(course.CourName, course.CourStart, course.DayWeek, course.TeacherName, course.Start, course.End)
		if !jud {
			_, _, cid := models.DbContr.CourseSave(CourseJsonToModel(course))
			_, _ = models.DbContr.CourseBand(cid, student.Id)
		} else {
			_, _ = models.DbContr.CourseBand(cid, student.Id)
		}
	}
	return returnSortJson(student.Id)
}
