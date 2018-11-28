package dao

import (
	"bufio"
	"github.com/UPrefer/StorageService/model"
	"io"
	"os"
)

func NewFileBlobDao(path string) *FileBlobDao {
	return &FileBlobDao{path: path}
}

type FileBlobDao struct {
	path string
}

func (dao *FileBlobDao) ReadData(artifactId string) (io.ReadCloser, error) {
	return os.Open(dao.path + artifactId)
}

func (dao *FileBlobDao) SaveData(dto *model.ArtifactDTO, contentType string, reader io.Reader) error {
	const chunkSize = 255
	var outputFile, err = os.OpenFile(dao.path+dto.Uuid, os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		return err
	}

	var fileWriter = bufio.NewWriterSize(outputFile, chunkSize)
	var requestReader = bufio.NewReaderSize(reader, chunkSize)

	_, err = requestReader.WriteTo(fileWriter)
	return err
}
