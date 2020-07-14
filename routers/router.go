package routers

import (
	"factly/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/user", middleware.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/delete/{Id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/all", middleware.GetAllUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/usr/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/update/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")

	return router
}
