package api

import (
	"github.com/UPrefer/StorageService/config"
	"github.com/UPrefer/StorageService/controller"
	"github.com/UPrefer/StorageService/dao"
	"github.com/UPrefer/StorageService/database"
	"github.com/UPrefer/StorageService/service"
	"github.com/gin-gonic/gin"
)

func Handlers(config *config.Configuration) *gin.Engine {
	engine := gin.Default()

	var (
		db = database.NewDatabase("mongodb://"+config.MongoUser+":"+config.MongoPassword+"@"+config.MongoIp+":"+config.MongoPort, config.MongoDB)

		artifactDao = dao.NewArtifactDao(db)
		blobDao     = dao.NewMongoBlobDao(db)

		utilsService    = service.NewUtilsService()
		artifactService = service.NewArtifactService(utilsService, artifactDao)
		blobService     = service.NewBlobService(blobDao, artifactDao)

		artifactController = controller.NewArtifactController(artifactService)
		blobController     = controller.NewBlobController(blobService)
	)

	v1 := engine.Group("/v1")
	{
		artifactGroup := v1.Group("/artifact")
		{
			artifactGroup.POST("", artifactController.Post)
		}

		artifactResource := artifactGroup.Group(":artifact_id")
		{
			artifactResource.GET("", artifactController.Get)
		}

		blobResource := artifactResource.Group("/blob")
		{
			blobResource.PUT("", blobController.Put)
			blobResource.GET("", blobController.Get)
		}
	}
	return engine
}
