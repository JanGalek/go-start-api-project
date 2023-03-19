package router

import (
	"github.com/gorilla/mux"
	"jan-galek/api/src/endpoints"
)

func HandleRequest() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", endpoints.Homepage).Methods("GET")
	return myRouter
}
