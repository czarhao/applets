package api

import (
	"applets/models"
	"applets/system"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func getParameters(ctx *gin.Context) (string, string, string) {
	return ctx.Param("sno"), ctx.Param("spw"), ctx.Param("types")
}

// 从内网服务器获取json
func getJson(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		system.Save.OtherError("从内网服务器获取消息出错！", err)
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			system.Save.OtherError("close getJson has made some errors ", err)
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

// 返回排序的课程
func returnSortJson(sid int) SchedulesJson {
	relation := models.DbContr.SidGetGid(sid)
	var returnCourses []CourseJson
	for _, value := range relation {
		tmpCourse := models.DbContr.SidGetCourse(value.Cid)
		returnCourses = append(returnCourses, CourseModelToJson(tmpCourse))
	}
	return buildSchedulesJson(true, "", sortCourse(returnCourses))
}

// 对课程按照周一到周日排序
func sortCourse(courses []CourseJson) [7][6]ScheduleJson {
	returnCourses := [7][6]ScheduleJson{}
	for _, value := range courses {
		if value.Start <= system.SemesterConf.TeachWeek && value.End >= system.SemesterConf.TeachWeek {
			returnCourses[value.DayWeek][(value.CourStart-1)/2] = CourseToSchedule(value)
		}
	}
	return returnCourses
}

func CourseModelToJson(course models.Course) CourseJson {
	return CourseJson{
		Id:          course.Id,
		Type:        course.Type,
		Start:       course.WeekStart,
		End:         course.WeekEnd,
		DayWeek:     course.DayWeek,
		CourStart:   course.DayStart,
		CourLength:  course.DayEnd - course.DayStart + 1,
		CourName:    course.CourseName,
		TeacherName: course.TeacherName,
		CourWhere:   course.CourWhere,
		Jud:         course.Jud,
	}
}
