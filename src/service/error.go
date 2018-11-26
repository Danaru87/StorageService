package service

import (
	"errors"
)

var (
	ErrArtifactNofFound        = errors.New("Artifact not found")
	ErrArtifactAlreadyUploaded = errors.New("Artifact already uploaded")
)
