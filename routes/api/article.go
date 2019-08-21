package api

import (
	"applets/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Article(ctx *gin.Context) {
	start := ctx.Param("start")
	startArticle, _ := strconv.Atoi(start)
	ctx.JSON(http.StatusOK, buildArticlesJson(true, "", buildArticleJson(models.DbContr.GetArticle(startArticle))))
}
