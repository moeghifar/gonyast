package util

import "gopkg.in/mgo.v2"

// InitMongo ...
func InitMongo() (session *mgo.Session, err error) {
	mongoConnection := Config.MongoConfig["core"]
	session, err = mgo.Dial(mongoConnection)
	return session, err
}
