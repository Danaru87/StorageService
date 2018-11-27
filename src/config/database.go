package config

import (
	"errors"
	"github.com/globalsign/mgo"
	"log"
	"time"
)

type GridFileReadCloser struct {
	gridFile *mgo.GridFile
	session  *mgo.Session
}

func (readCloser *GridFileReadCloser) Close() error {
	var err = readCloser.gridFile.Close()
	readCloser.session.Close()
	return err
}

func (readCloser *GridFileReadCloser) Read(p []byte) (n int, err error) {
	return readCloser.gridFile.Read(p)
}

type Request func(*mgo.Database)

type Database struct {
	session      *mgo.Session
	databaseName string
}

func NewDatabase(url string, databaseName string) *Database {
	var (
		session *mgo.Session = nil
		err                  = errors.New("DB no connection attempt yet")
	)

	for err != nil {
		session, err = mgo.Dial(url)
		if err != nil {
			log.Println("DB : Connection failed at " + url + ", waiting 5 secondes before retrying")
			time.Sleep(5 * time.Second)
		}
	}
	log.Println("DB : Connexion succeed !")

	return &Database{session: session, databaseName: databaseName}
}

func (database *Database) OpenGridFsReader(prefix string, objectId string) (*GridFileReadCloser, error) {
	var newSession = database.session.Copy()
	var fileReader, err = newSession.DB(database.databaseName).GridFS(prefix).OpenId(objectId)
	return &GridFileReadCloser{session: newSession, gridFile: fileReader}, err
}

func (database *Database) HandleRequest(request Request) {
	var newSession = database.session.Copy()
	defer newSession.Close()

	request(newSession.DB(database.databaseName))
}
