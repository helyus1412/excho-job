package routes

import (
	"excho-job/handler"

	"github.com/gin-gonic/gin"
)

func QuestionRoute(r *gin.Engine) {
	r.GET("/question", handler.QuestionFetch)
}
