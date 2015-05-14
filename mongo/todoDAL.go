package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty" `
	Name      string        `json:"name" bson:"name"`
	Completed bool          `json:"completed" bson:"completed"`
	Due       time.Time     `json:"due" bson:"due"`
}

type Todos []Todo

func InsertToDO(collection *mgo.Collection, todo Todo) {

	err = collection.Insert(&todo)
	if err != nil {
		log.Printf("%s", err)

	}
}

func FindToDos(collection *mgo.Collection) Todos {

	var result []Todo
	err = collection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Printf("%s", err)

	}

	log.Printf("%s", result)
	if len(result) > 0 {
		return result
	} else {
		return nil
	}

}

func DeleteToDo(collection *mgo.Collection, id string) int {
	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		log.Printf("%s", err)
		return http.StatusServiceUnavailable
	} else {
		return http.StatusOK
	}

}

func UpdateToDo(collection *mgo.Collection, id string, completed bool) int {
	err = collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"completed": completed}})
	if err != nil {
		log.Printf("%s", err)
		return http.StatusServiceUnavailable
	} else {
		return http.StatusOK
	}
}
