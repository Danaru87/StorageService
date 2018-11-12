package controller

import (
	"github.com/gin-gonic/gin"
)

type ArtifactController struct{}

func NewArtifactController() *ArtifactController {
	return new(ArtifactController)
}

func (*ArtifactController) Post(ctx *gin.Context) {
}
