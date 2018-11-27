package model

type ArtifactDTO struct {
	Uuid        string `json:"uuid" bson:"_id,omitempty"`
	Length      int64  `json:"length" bson:"length"`
	ContentType string `json:"contentType" bson:"contentType"`
}
