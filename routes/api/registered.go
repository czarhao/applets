package api

import (
	"applets/models"
	"applets/system"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Registered(ctx *gin.Context) {
	sno, spw, _ := getParameters(ctx)
	inDb, _ := models.DbContr.StudentExit(sno)
	if inDb == true {
		ctx.JSON(http.StatusOK, buildSchedulesJson(false, "同学已经注册了", [7][6]ScheduleJson{}))
	} else {
		body, err := getJson(system.CrawlServer + "all/" + sno + "/" + spw)
		if err != nil {
			ctx.JSON(http.StatusOK, buildSchedulesJson(false, "内部出现错误...", [7][6]ScheduleJson{}))
			return
		}
		registeredJson := RegisteredJson{}
		if err := json.Unmarshal(body, &registeredJson); err != nil {
			system.Save.OtherError("Unmarshal json", err)
			ctx.JSON(http.StatusOK, buildSchedulesJson(false, "内部出现错误...", [7][6]ScheduleJson{}))
			return
		}
		if registeredJson.Success == false {
			ctx.JSON(http.StatusOK, buildSchedulesJson(false, registeredJson.Info, [7][6]ScheduleJson{}))
		} else {
			// 确保消息获取完全正确
			if !registeredJson.InfoJson.Success {
				ctx.JSON(http.StatusOK, buildSchedulesJson(false, "信息获取出错！", [7][6]ScheduleJson{}))
				system.Save.Warning("信息获取出错！", errors.New(registeredJson.InfoJson.Info))
				return
			}
			if !registeredJson.GradeJson.Success {
				ctx.JSON(http.StatusOK, buildSchedulesJson(false, "成绩获取出错！", [7][6]ScheduleJson{}))
				system.Save.Warning("成绩获取出错！", errors.New(registeredJson.InfoJson.Info))
				return
			}
			if !registeredJson.CourseJson.Success {
				ctx.JSON(http.StatusOK, buildSchedulesJson(false, "课程获取出错！", [7][6]ScheduleJson{}))
				system.Save.Warning("课程获取出错！", errors.New(registeredJson.InfoJson.Info))
				return
			}
			// 防止数据库重复添加，再加一道验证
			inDb, _ = models.DbContr.StudentExit(sno)
			if inDb == true {
				ctx.JSON(http.StatusOK, buildSchedulesJson(false, "同学已经注册了", [7][6]ScheduleJson{}))
				return
			}
			_, sid, _ := models.DbContr.StudentCreate(infoToModel(registeredJson.InfoJson))
			for _, grade := range registeredJson.GradeJson.Grades {
				_, _ = models.DbContr.GradeCreate(sid, GradeToModel(grade))
			}
			for _, course := range registeredJson.CourseJson.Schedule {
				jud, cid := models.DbContr.CourseExit(course.CourName, course.CourStart, course.DayWeek, course.TeacherName, course.Start, course.End)
				if !jud {
					_, _, cid := models.DbContr.CourseSave(CourseJsonToModel(course))
					_, _ = models.DbContr.CourseBand(cid, sid)
				} else {
					_, _ = models.DbContr.CourseBand(cid, sid)
				}
			}
			// 返回课程表
			ctx.JSON(http.StatusOK, returnSortJson(sid))
		}
	}
}
