package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.ibm.com/diksha/movies-restapi/config"
	. "github.ibm.com/diksha/movies-restapi/dao"
	. "github.ibm.com/diksha/movies-restapi/models"
)

var config = Config{}
var dao = DAOMovies{}

// GET list of movies
func AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

// GET a movie by its ID
func FindMovieByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

// POST a new movie
func CreateMovieInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dao.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

// PUT update an existing movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing movie
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
func FindMovieByYear(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Print("params var ::",params["release_year"])
	i1, err := strconv.Atoi(params["release_year"])
	if err == nil {
		fmt.Println(i1)
	}
	movie, err := dao.FindByYear(i1)
	log.Print("in main function",movie)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie year")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

func FindMovieBelowRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Print("params var ::",params["rating"])
	i1, err := strconv.Atoi(params["rating"])
	if err == nil {
		fmt.Println(i1)
	}
	movie, err := dao.FindBelowRating(i1)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid  rating")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}
func FindMovieAboveRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Print("params var ::",params["rating"])
	i1, err := strconv.Atoi(params["rating"])
	if err == nil {
		fmt.Println(i1)
	}
	movie, err := dao.FindAboveRating(i1)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid  rating")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

func FindMovieUptoYear(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Print("params var ::",params["release_year"])
	i1, err := strconv.Atoi(params["release_year"])
	if err == nil {
		fmt.Println(i1)
	}
	movie, err := dao.FindUptoYear(i1)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie year")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	log.Print(config.Server)
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMovies).Methods("GET")
	r.HandleFunc("/movies", CreateMovieInfo).Methods("POST")
	r.HandleFunc("/movies", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieByID).Methods("GET")
	r.HandleFunc("/movies/release_year/{release_year}",FindMovieByYear).Methods("GET")

	r.HandleFunc("/movies/rating/below/{rating}",FindMovieBelowRating).Methods("GET")
	r.HandleFunc("/movies/rating/above/{rating}",FindMovieAboveRating).Methods("GET")
	r.HandleFunc("/movies/release_year/upto/{release_year}",FindMovieUptoYear).Methods("GET")


	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
