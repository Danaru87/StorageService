package controller

import (
	"github.com/UPrefer/StorageService/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlobController struct {
	blobService service.IBlobService
}

func NewBlobController(blobService service.IBlobService) *BlobController {
	return &BlobController{blobService: blobService}
}

func (blobController *BlobController) Get(ctx *gin.Context) {
	var contentType, contentLength, readCloser, err = blobController.blobService.ReadBlob(ctx.Param("artifact_id"))

	if readCloser != nil {
		defer readCloser.Close()
	}

	if err == service.ErrArtifactNotFound {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.DataFromReader(http.StatusOK, contentLength, contentType, readCloser, map[string]string{"Content-Disposition": "attachment; filename=\"toto.zip\""})
}

func (blobController *BlobController) Put(ctx *gin.Context) {
	var err = blobController.blobService.SaveBlob(ctx.Param("artifact_id"), ctx.GetHeader("Content-Type"), ctx.Request.Body)

	if err == service.ErrArtifactNotFound {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	} else if err == service.ErrArtifactAlreadyUploaded {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	} else if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
