package dao

import (
	"crypto/rand"
	"github.com/UPrefer/StorageService/config"
	"github.com/UPrefer/StorageService/model"
	"github.com/google/uuid"
	"io"
	"testing"
)

func Benchmark_10MoBlobPersistenceSpeed(b *testing.B) {
	//GIVEN
	var (
		database   = config.NewDatabase("mongodb://root:root@localhost:27017", "StorageService_bench")
		mongoDao   = NewMongoBlobDao(database)
		dataReader = io.LimitReader(rand.Reader, 10*1024*1024)
	)

	//WHEN
	mongoDao.SaveData(&model.ArtifactDTO{Uuid: uuid.New().String()}, "random", dataReader)
}

func Benchmark_100MoBlobPersistenceSpeed(b *testing.B) {
	//GIVEN
	var (
		database   = config.NewDatabase("mongodb://root:root@localhost:27017", "StorageService_bench")
		mongoDao   = NewMongoBlobDao(database)
		dataReader = io.LimitReader(rand.Reader, 100*1024*1024)
	)

	//WHEN
	mongoDao.SaveData(&model.ArtifactDTO{Uuid: uuid.New().String()}, "random", dataReader)
}

func Benchmark_1000MoBlobPersistenceSpeed(b *testing.B) {
	//GIVEN
	var (
		database   = config.NewDatabase("mongodb://root:root@localhost:27017", "StorageService_bench")
		mongoDao   = NewMongoBlobDao(database)
		dataReader = io.LimitReader(rand.Reader, 1000*1024*1024)
	)

	//WHEN
	mongoDao.SaveData(&model.ArtifactDTO{Uuid: uuid.New().String()}, "random", dataReader)
}
