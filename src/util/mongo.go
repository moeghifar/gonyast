package util

import "gopkg.in/mgo.v2"

// InitMongo ...
func InitMongo() (session *mgo.Session, err error) {
	mongoConnection := "localhost:27017"
	session, err = mgo.Dial(mongoConnection)
	return session, err
}
