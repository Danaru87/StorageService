package service

import (
	"errors"
)

var (
	ErrArtifactNotFound        = errors.New("Artifact not found")
	ErrArtifactAlreadyUploaded = errors.New("Artifact already uploaded")
)
