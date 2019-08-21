package api

import (
	"applets/models"
	"applets/system"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 查询成绩
func Grade(ctx *gin.Context) {
	sno, spw, _ := getParameters(ctx)
	inDb, info := models.DbContr.StudentExit(sno)
	// 判断该学生信息是否存在数据库中，密码是否与数据库中的相同
	if inDb && info.Password == spw {
		// 相同直接返回数据库中的信息
		grades := models.DbContr.GradeReturn(info.Id)
		var tmpGrades []GradeJson
		for _, grade := range grades {
			tmpGrades = append(tmpGrades, GradeModelToJson(grade))
		}
		ctx.JSON(http.StatusOK, buildGradesJson(true, "", tmpGrades))
	} else {
		if inDb == false {
			ctx.JSON(http.StatusOK, buildGradesJson(false, "同学尚未注册我们的服务。。。", nil))
		} else {
			ctx.JSON(http.StatusOK, buildGradesJson(false, "同学输入的密码错误！", nil))
		}
	}
}

// 更新成绩
func updateGrade(sno string, spw string) GradesJson {
	jud, student := models.DbContr.StudentExit(sno)
	if !jud {
		return buildGradesJson(false, "同学还没有注册我们的服务呀", nil)
	} else {
		body, _ := getJson(system.CrawlServer + "grade/" + sno + "/" + spw)
		tmpJson := GradesJson{}
		err := json.Unmarshal(body, &tmpJson)
		if err != nil {
			return buildGradesJson(false, err.Error(), nil)
		}
		if tmpJson.Success == false {
			return buildGradesJson(false, tmpJson.Info, nil)
		}
		// 删除原有的成绩
		deleteJud := models.DbContr.GradeDelete(student.Id)
		if deleteJud == false {
			return buildGradesJson(false, "删除原来成绩时出错", nil)
		}
		// 重新插入新的成绩
		for _, grade := range tmpJson.Grades {
			_, _ = models.DbContr.GradeCreate(student.Id, GradeToModel(grade))
		}
		return tmpJson
	}
}
