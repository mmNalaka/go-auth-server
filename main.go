package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

var products = []Product{
	{Id: 1, Name: "World of Authcraft", Slug: "world-of-authcraft", Description: "Battle bugs and protect yourself from invaders while you explore a scary world with no security"},
	{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.Handle("/status", http.HandlerFunc(StatusHandler)).Methods("GET")
	r.Handle("/products", http.HandlerFunc(ProductsHandler)).Methods("GET")
	r.Handle("/products/{slug}/feedback", http.HandlerFunc(NotImplemented)).Methods("POST")

	addr := ":4000"
	log.Println("Server listing on port", addr)
	http.ListenAndServe(addr, r)
}

var NotImplemented = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Not implemented!!"))
})

var StatusHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("X-Origin", "nm-auth-server")
	rw.Write([]byte("API is up and running"))
})

var ProductsHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(products)

	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(payload))
})
