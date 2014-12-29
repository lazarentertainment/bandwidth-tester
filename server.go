package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/bandwidth", BandwidthHandler)
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", r)
}

func BandwidthHandler(response http.ResponseWriter, request *http.Request) {
	start := time.Now()
	if request.Method == "POST" {
		params := mux.Vars(request)
		log.Println(params["image"])
		response.Write([]byte("This is a server response"))
	} else {
		http.Error(response, "Can't touch this", http.StatusUnauthorized)
	}
	log.Println(request.URL)
	log.Println(request.Method)
	log.Println("Response took", time.Since(start), "seconds")
}
