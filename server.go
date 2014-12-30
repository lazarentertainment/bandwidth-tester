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
	http.Handle("/", &LeServer{r})
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

type LeServer struct {
	r *mux.Router
}

func (serv *LeServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if origin := request.Header.Get("Origin"); origin != "" {
		response.Header().Set("Access-Control-Allow-Origin", origin)
		response.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		response.Header().Set("Access-Control-Allow-Headers",
			"Accept, contentType, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	if request.Method == "OPTIONS" {
		return
	}
	serv.r.ServeHTTP(response, request)
}

func BandwidthHandler(response http.ResponseWriter, request *http.Request) {
	start := time.Now()
	if request.Method == "POST" {
		params := request.ParseForm()
		log.Println(request)
		log.Println(params)
		//log.Println(params["image"])
		response.Write([]byte("This is a server response"))
	} else {
		http.Error(response, "Can't touch this", http.StatusUnauthorized)
	}
	log.Println(request.URL)
	log.Println(request.Method)
	log.Println("Response took", time.Since(start), "seconds")
}
