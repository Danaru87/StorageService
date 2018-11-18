package model

type ArtifactDTO struct {
	BaseModel
	Name string `json:"name" binding:"required"`
}
