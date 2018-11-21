package config

import (
	"errors"
	"github.com/globalsign/mgo"
	"log"
	"time"
)

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

func (database *Database) HandleRequest(request Request) {
	var newSession = database.session.Copy()
	defer newSession.Close()

	request(newSession.DB(database.databaseName))
}
