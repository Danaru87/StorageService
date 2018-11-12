package api

import (
	"github.com/UPrefer/StorageService/controller"
	"github.com/gin-gonic/gin"
)

func Handlers() *gin.Engine {
	engine := gin.Default()

	v1 := engine.Group("/v1")
	{
		artifactResource := v1.Group("/artifact")
		{
			artifactResource.POST("", (controller.NewArtifactController()).Post)
		}
	}
	return engine
}
