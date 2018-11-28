package model

import "time"

type ArtifactDTO struct {
	Uuid        string     `json:"uuid" bson:"_id,omitempty"`
	Name        string     `json:"name" binding:"required"bson:"filename"`
	Length      int64      `json:"length,omitempty" bson:"length,omitempty"`
	ContentType string     `json:"contentType,omitempty" bson:"contentType,omitempty"`
	Md5         string     `json:"md5sum,omitempty" bson:"md5,omitempty"`
	UploadDate  *time.Time `json:"uploadDate,omitempty" bson:"uploadDate,omitempty"`
}
