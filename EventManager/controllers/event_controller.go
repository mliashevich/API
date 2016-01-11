package controllers

import (
	"fmt"
	"../models"
	"time"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"

	"net/http"
	"encoding/json"
)

type (
	EventsController struct {
		session *mgo.Session
	}
)

func NewEventsController(s *mgo.Session) *EventsController {
	return &EventsController{s}
}

/* Get all events */
func (uc EventsController) GetEvents(resp http.ResponseWriter, req *http.Request) {


	var allEvents []models.Event

	if err := uc.session.DB("bsu_api").C("events").Find(nil).All(&allEvents); err != nil {
		resp.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(allEvents)



	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(200)

	fmt.Fprintf(resp, "%s", uj)
}

/* Get event by Id */
func (uc EventsController) GetEventById(resp http.ResponseWriter, req *http.Request) {

	event := models.Event{}

	if err := uc.session.DB("bsu_api").C("events").FindId(bson.ObjectIdHex(mux.Vars(req)["event_id"])).One(&event); err != nil {
		resp.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(event)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(200)

	fmt.Fprintf(resp, "%s", uj)
}

/* Add new event */
func (uc EventsController) AddEvent(resp http.ResponseWriter, req *http.Request) {

	u := models.Event{}

	json.NewDecoder(req.Body).Decode(&u)

	u.Id = bson.NewObjectId()
	u.CreatedAt = time.Now()
	u.PlacesLeft = u.MaxCount

	uc.session.DB("bsu_api").C("events").Insert(u)

	uj, _ := json.Marshal(u)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(201)

	fmt.Fprintf(resp, "%s", uj)
}

func (uc EventsController) RegisterMember(resp http.ResponseWriter, req *http.Request) {

	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"places_left": -1}},
	}

	event := models.Event{}

	//	info, err := uc.session.DB("bsu_api").C("events").Find(
	//		bson.M{
	//			"_id" : bson.ObjectIdHex(mux.Vars(req)["event_id"])}).Apply(change, &event)

	info, err := uc.session.DB("bsu_api").C("test").Find(
		bson.M{
			"_id" : bson.ObjectIdHex(mux.Vars(req)["event_id"]), "places_left": bson.M{"$gt": 0 }}).Apply(change, &event)

	//	user := models.User{}
	//	json.NewDecoder(req.Body).Decode(&user)
	//	user.Id = bson.NewObjectId()
	//	uc.session.DB("bsu_api").C("test").Insert(user)

	if err != nil {
		resp.WriteHeader(404)
		return
	}

	fmt.Fprintf(resp, "%s", info)
}