package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Event struct {
		Id             bson.ObjectId `json:"id" bson:"_id"`
		Name           string        `json:"name" bson:"name"`
		Details        string        `json:"details" bson:"details"`
		Location       string        `json:"location" bson:"location"`
		StartDate      time.Time     `json:"start_time" bson:"start_time"`
		AttendingCount int             `json:"attending_count" bson:"attending_count"`
		MaxCount       int             `json:"max_count" bson:"max_count"`
		CreatedAt      time.Time     `json:"created_at" bson:"created_at"`
	}
)


//curl http://go.dev/events/568ad56b945dbe4c3bad75dc
//curl -XPOST -H 'Content-Type: application/json' -d '{"name": "Party", "details": "Super party", "location": "Belarus, Minsk", "max_count": 10 }' http://go.dev/events