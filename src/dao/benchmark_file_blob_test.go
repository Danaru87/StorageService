package dao

import (
	"crypto/rand"
	"github.com/UPrefer/StorageService/model"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"testing"
)

func Benchmark_10MoFileBlobPersistenceSpeed(b *testing.B) {
	//GIVEN
	var (
		tmpFile, _ = ioutil.TempFile("", "bench10Mb*")
		fileDao    = NewFileBlobDao(tmpFile.Name())
		dataReader = io.LimitReader(rand.Reader, 10*1024*1024)
	)

	//WHEN
	fileDao.SaveData(&model.ArtifactDTO{Uuid: uuid.New().String()}, "random", dataReader)
}

func Benchmark_100MoFileBlobPersistenceSpeed(b *testing.B) {
	//GIVEN
	var (
		tmpFile, _ = ioutil.TempFile("", "bench100Mb*")
		fileDao    = NewFileBlobDao(tmpFile.Name())
		dataReader = io.LimitReader(rand.Reader, 100*1024*1024)
	)

	//WHEN
	fileDao.SaveData(&model.ArtifactDTO{Uuid: uuid.New().String()}, "random", dataReader)
}

func Benchmark_1000MoFileBlobPersistenceSpeed(b *testing.B) {
	//GIVEN
	var (
		tmpFile, _ = ioutil.TempFile("", "bench100Mb*")
		fileDao    = NewFileBlobDao(tmpFile.Name())
		dataReader = io.LimitReader(rand.Reader, 1000*1024*1024)
	)

	//WHEN
	fileDao.SaveData(&model.ArtifactDTO{Uuid: uuid.New().String()}, "random", dataReader)
}
