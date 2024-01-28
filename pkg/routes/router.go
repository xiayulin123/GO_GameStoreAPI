package routes

import (
	"github.com/gorilla/mux"
	"github.com/xiayulin123/GO_GameStoreAPI/pkg/controllers"
)

var routes = func(router *mux.Router) {
	router.HandleFunc("/toy/", controllers.ShowToys).Methods("GET")
	router.HandleFunc("/toy/", controllers.CreateToys).Methods("post")
	router.HandleFunc("/toy/{toyId}", controllers.ShowToyById).Methods("GET")
	router.HandleFunc("/toy/{toyId}", controllers.DeleteToys).Methods("DELETE")
	router.HandleFunc("/toy/{toyId}", controllers.UpdateToys).Methods("PATCH")
}
