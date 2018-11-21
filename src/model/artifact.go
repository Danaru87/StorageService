package model

type ArtifactDTO struct {
	Uuid string `json:"uuid" bson:"_id,omitempty"`
	Name string `json:"name" binding:"required" bson:"name"`
}
