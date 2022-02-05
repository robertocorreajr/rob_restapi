package server

import "github.com/gorilla/mux"

func Router() *mux.Router {
	return mux.NewRouter()
}
