package api

import (
	"applets/models"
	"applets/system"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 从数据库中读取学生的信息
func Info(ctx *gin.Context) {
	sno, spw, _ := getParameters(ctx)
	inDb, info := models.DbContr.StudentExit(sno)
	// 判断该学生信息是否存在数据库中，密码是否与数据库中的相同
	if inDb && info.Password == spw {
		// 相同直接返回数据库中的信息
		ctx.JSON(http.StatusOK, buildInfoJson(true, "", *info))
	} else {
		if inDb == false {
			ctx.JSON(http.StatusOK, buildInfoJson(false, "同学尚未注册我们的服务。。。", models.Student{}))
		} else {
			ctx.JSON(http.StatusOK, buildInfoJson(false, "同学输入的密码错误。。。", models.Student{}))
		}
	}
}

// 更新用户信息
func updateInfo(sno string, spw string) InfoJson {
	body, _ := getJson(system.CrawlServer + "info/" + sno + "/" + spw)
	tmpJson := InfoJson{}
	err := json.Unmarshal(body, &tmpJson)
	if tmpJson.Success == false || err != nil {
		return buildInfoJson(false, tmpJson.Info, models.Student{})
	} else {
		_, _, _ = models.DbContr.StudentUpdate(infoToModel(tmpJson))
		return tmpJson
	}
}
