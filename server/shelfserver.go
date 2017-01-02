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

	apiV1.HandleFunc("/useString", UseString).Methods("GET")
	apiV1.HandleFunc("/generateBjorklund", GenerateBjorklund).Methods("GET")

	http.Handle("/", &RhythmShelfServer{r})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
