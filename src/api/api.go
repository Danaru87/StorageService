package api

import (
	"github.com/UPrefer/StorageService/config"
	"github.com/UPrefer/StorageService/controller"
	"github.com/UPrefer/StorageService/dao"
	"github.com/UPrefer/StorageService/service"
	"github.com/gin-gonic/gin"
)

func Handlers() *gin.Engine {
	engine := gin.Default()

	var (
		database = config.NewDatabase("mongodb://root:root@localhost:27017", "StorageService")

		artifactDao = dao.NewArtifactDao(database)

		utilsService    = service.NewUtilsService()
		artifactService = service.NewArtifactService(utilsService, artifactDao)

		artifactController = controller.NewArtifactController(artifactService)
	)

	v1 := engine.Group("/v1")
	{
		artifactResource := v1.Group("/artifact")
		{
			artifactResource.POST("", artifactController.Post)
		}
	}
	return engine
}
