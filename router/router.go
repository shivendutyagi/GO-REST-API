package router

import (
	"github.com/gorilla/mux"
	"github.com/shivendutyagi/newapi/controllers"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("api/getmovie", controllers.Getallmovies).Methods("GET")
	r.HandleFunc("api/create", controllers.Createmovie).Methods("POST")
	r.HandleFunc("api/getmovie/{id}", controllers.Marksaswatched).Methods("PUT")
	r.HandleFunc("api/getmovie/{id}", controllers.Deleteamovie).Methods("DELETE")

	return r
}
