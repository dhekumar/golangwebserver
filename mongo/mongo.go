package mongo

import (
	"golangwebserver/constants"
	"gopkg.in/mgo.v2"
)

var session, err = mgo.Dial(constants.MongoDBConnectionString)

func init() {

	if err != nil {
		panic(err)
	}

	// defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

}

func GetToDoModel() *mgo.Collection {
	m := session.DB(constants.DatabaseName).C("todo")
	return m
}
