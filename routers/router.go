package routers

import (
	"factly/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/user", middleware.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/delete/{Id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")

	return router
}
