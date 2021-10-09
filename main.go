package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	router := httprouter.New()
	uc := control.newUserController(getSession())
	router.GET("/user/:id", uc.getUser)
	router.POST("/user", uc.createUser)

	http.ListenAndServe("localhost:8080", router)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return session
}
