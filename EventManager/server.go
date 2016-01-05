package main

import (
	"net"
	"./controllers"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"

	"net/http/fcgi"
)

type FastCGIServer struct{}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:9001")

	r := mux.NewRouter().StrictSlash(true)

	ec := controllers.NewEventsController(getSession())

	r.HandleFunc("/events", ec.GetEvents).Methods("GET")
	r.HandleFunc("/events/{event_id}", ec.GetEventById).Methods("GET")

	r.HandleFunc("/events", ec.AddEvent).Methods("POST")
	r.HandleFunc("/events/{event_id}/attend", ec.RegisterMember).Methods("POST")

	fcgi.Serve(listener, r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")

	s.SetMode(mgo.Strong, true)

	if err != nil {
		panic(err)
	}
	return s
}
