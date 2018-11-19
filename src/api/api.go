package api

import (
	"github.com/UPrefer/StorageService/controller"
	"github.com/UPrefer/StorageService/dao"
	"github.com/UPrefer/StorageService/service"
	"github.com/gin-gonic/gin"
)

func Handlers() *gin.Engine {
	engine := gin.Default()

	v1 := engine.Group("/v1")
	{
		artifactResource := v1.Group("/artifact")
		{
			artifactResource.POST("", (controller.NewArtifactController(service.NewArtifactService(dao.NewArtifactDao()))).Post)
		}
	}
	return engine
}
