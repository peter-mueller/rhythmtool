package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/peter-mueller/rhythmtool"
)

type RhythmShelfServer struct {
	r *mux.Router
}

func (s *RhythmShelfServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}

func GenerateBjorklund(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pulses, err := strconv.Atoi(r.Form.Get("pulses"))
	if err != nil || pulses < 0 {
		http.Error(w, "Query parameter 'pulses' must be at least 0!", http.StatusBadRequest)
		return
	}
	length, err := strconv.Atoi(r.Form.Get("length"))
	if err != nil || length <= 1 {
		http.Error(w, "Query parameter 'length' must be bigger than 1!", http.StatusBadRequest)
		return
	}

	if pulses > length {
		http.Error(w, "'pulses' must not be bigger than 'length'!", http.StatusBadRequest)
		return
	}

	result := rhythmtool.GenerateBjorklund(pulses, length)
	json.NewEncoder(w).Encode(result)
}

func UseString(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s := r.Form.Get("s")
	result := rhythmtool.UseString(s)
	json.NewEncoder(w).Encode(result)
}
