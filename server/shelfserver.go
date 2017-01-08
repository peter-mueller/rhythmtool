// rhythm-shelf-server project rhythm-shelf-server.go
package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer(port string) {
	r := mux.NewRouter()
	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	r.PathPrefix("/app/").Handler(http.StripPrefix("/app/", http.FileServer(http.Dir("."))))

	apiV1.HandleFunc("/useString", UseString).Methods("GET")
	apiV1.HandleFunc("/generateBjorklund", GenerateBjorklund).Methods("GET")
	apiV1.HandleFunc("/random", Random).Methods("GET")

	http.Handle("/", &RhythmShelfServer{r})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
