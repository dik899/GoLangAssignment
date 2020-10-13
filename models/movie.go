package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Title        string        `bson:"title" json:"title"`
	ReleaseYear  int          `bson:"release_year" json:"release_year"`
	Rating       int        `bson:"rating" json:"rating"`
	Genres       []string   `bson:"genres" json:"genres"`
}
