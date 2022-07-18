package main

import _(
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"encoding/json"
	"math/rand"
	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"


)

type Movie struct {
	id string `json:"id"`
	isbn string `json:"isbn"`
	title string `json:"title"`
	Director *Director `json:"director"`

}

type Director struct {
	firstname string `json:"firstname"`
	lastname string `json:"lastname"`
}

// movies slice
var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.id = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
	w.WriteHeader(http.StatusCreated, "Movie created")
	
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			w.WriteHeader(http.StatusOK, "Movie updated")
			return
		}
	}
	// json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{id: "1", isbn: "448743", title: "The Shawshank Redemption", Director: &Director{firstname: "Frank", lastname: "Darabont"}})
	movies = append(movies, Movie{id: "2", isbn: "448744", title: "The Godfather", Director: &Director{firstname: "Francis", lastname: "Ford Coppola"}})
	movies = append(movies, Movie{id: "3", isbn: "448745", title: "The Godfather: Part II", Director: &Director{firstname: "Francis", lastname: "Ford Coppola"}})
	movies = append(movies, Movie{id: "4", isbn: "448746", title: "Pulp Fiction", Director: &Director{firstname: "Quentin", lastname: "Tarantino"}})
	movies = append(movies, Movie{id: "5", isbn: "448747", title: "The Lord of the Rings: The Return of the King", Director: &Director{firstname: "Peter", lastname: "Jackson"}})
	movies = append(movies, Movie{id: "6", isbn: "448748", title: "The Good, the Bad and the Ugly", Director: &Director{firstname: "Sergio", lastname: "Leone"}})
	movies = append(movies, Movie{id: "7", isbn: "448749", title: "Fight Club", Director: &Director{firstname: "David", lastname: "Fincher"}})

	// routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, r))

}