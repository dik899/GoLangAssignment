package dao

import (
	"log"	


	."github.com/diksha/movies-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

type DAOMovies struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

func (m *DAOMovies) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of movies
func (m *DAOMovies) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

// Find a movie by its id
func (m *DAOMovies) FindById(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// Insert a movie into database
func (m *DAOMovies) Insert(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

// Delete an existing movie
func (m *DAOMovies) Delete(movie Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

// Update an existing movie
func (m *DAOMovies) Update(movie Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}

// Find a movie by its release year
func (m *DAOMovies) FindByYear(release_year int) ([]Movie, error) {
	query := bson.M{"release_year": release_year}
	movie := make([]Movie, 0, 10) 
	err := db.C(COLLECTION).Find(query).Sort("title").All(&movie)
	if err != nil {
	    log.Fatal(err)
	}
	return movie, err
}

// Find a set of movies below a perticular rating
func (m *DAOMovies) FindBelowRating(rating int) ([]Movie, error) {	
	query :=bson.M{"rating": bson.M{"$gt": 0, "$lt": rating}}
	movie := make([]Movie, 0, 10) 
	err := db.C(COLLECTION).Find(query).Sort("rating").All(&movie)
	if err != nil {
	    log.Fatal(err)
	}
	return movie, err
}

// Find a set of movies above a perticular rating
func (m *DAOMovies) FindAboveRating(rating int) ([]Movie, error) {	
	query :=bson.M{"rating": bson.M{"$gt": rating}}
	movie := make([]Movie, 0, 10) 
	err := db.C(COLLECTION).Find(query).Sort("rating").All(&movie)
	if err != nil {
	    log.Fatal(err)
	}
	return movie, err
}

// Find a set of movies upto perticular year
func (m *DAOMovies) FindUptoYear(year int) ([]Movie, error) {
	query := bson.M{"release_year": bson.M{"$gt": 0, "$lt": year+1}}
	movie := make([]Movie, 0, 10) 
	err := db.C(COLLECTION).Find(query).Sort("release_year").All(&movie)
	if err != nil {
	    log.Fatal(err)
	}
	return movie, err
}
