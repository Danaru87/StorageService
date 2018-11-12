package api

import (
	"github.com/gin-gonic/gin"
)

func Handlers() *gin.Engine {
	engine := gin.Default()

	return engine
}
