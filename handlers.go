package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golangwebserver/mongo"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type greetings struct {
	Intro    string
	Messages []string
}

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Welcome!")
	passedObj := greetings{
		Intro:    "Hello from Go!",
		Messages: []string{"Hello!", "Hi!", "¡Hola!", "Bonjour!", "Ciao!", "<script>evilScript()</script>"},
	}
	templates.ExecuteTemplate(w, "homePage", passedObj)

}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := mongo.Todos{
		mongo.Todo{Name: "Write presentation"},
		mongo.Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}

	todo := mongo.Todo{
		Name:      "Write presentation",
		Completed: false,
		Due:       time.Now(),
	}

	log.Printf("%s", todo)
	mongo.InsertToDO(mongo.GetToDoModel(), todo)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(mongo.FindToDos(mongo.GetToDoModel())); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
	// mongo.FindToDos(mongo.GetToDoModel())

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(mongo.FindToDos(mongo.GetToDoModel())); err != nil {
		panic(err)
	}

}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo mongo.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	log.Printf("body is %s", body)
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// t := RepoCreateTodo(todo)
	mongo.InsertToDO(mongo.GetToDoModel(), todo)
	log.Printf("%s", todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	todoId := vars["todoId"]

	log.Printf(" Deleting ToDO with Id %s", todoId)
	status := mongo.DeleteToDo(mongo.GetToDoModel(), todoId)
	// status := mongo.DeleteToDO(mongo.GetToDoModel(), todo.Name)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

}

func TodoUpdate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	todoId := vars["todoId"]

	var todo mongo.Todo

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	log.Printf("body is %s", body)

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	log.Printf(" updating ToDO with Id %s", todoId)
	status := mongo.UpdateToDo(mongo.GetToDoModel(), todoId, todo.Completed)
	// status := mongo.DeleteToDO(mongo.GetToDoModel(), todo.Name)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

}
