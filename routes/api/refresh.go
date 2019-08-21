package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 刷新课程
func Refresh(ctx *gin.Context) {
	sno, spw, types := getParameters(ctx)
	switch types {
	case "info":
		ctx.JSON(http.StatusOK, updateInfo(sno, spw))
	case "schedule":
		ctx.JSON(http.StatusOK, updateSchedule(sno, spw))
	case "grade":
		ctx.JSON(http.StatusOK, updateGrade(sno, spw))
	}
}
