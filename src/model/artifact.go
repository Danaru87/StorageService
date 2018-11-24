package model

type ArtifactDTO struct {
	Uuid string `json:"uuid" bson:"_id,omitempty"`
}
