package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"./handler/userhandler"
	"./handler/noteshandler"
	"./databaseConnect"
	"./middleware"
)

func main(){

	databaseConnect.InitDatabase()
	router := mux.NewRouter()

	router.HandleFunc("/users",userhandler.GetUser).Methods("GET");
	router.HandleFunc("/users",userhandler.CreateUser).Methods("POST");
	
	router.HandleFunc("/login",userhandler.AuthenticateUser).Methods("POST")
	router.HandleFunc("/home",userhandler.UserDetails).Methods("GET")
	router.HandleFunc("/refreshtoken",userhandler.RefreshToken).Methods("POST")

	router.HandleFunc("/users/notes", middleware.CookieMiddleware(noteshandler.GetAllNotes)).Methods("GET")
	router.HandleFunc("/users/notes", middleware.CookieMiddleware(noteshandler.CreateNote)).Methods("POST")
	router.HandleFunc("/users/notes/{id}", middleware.CookieMiddleware(noteshandler.DeleteNote)).Methods("DELETE")




	e := http.ListenAndServe(":7005",router)
	if e != nil {
		log.Fatal("ListenAndServe: ", e)
	}	
}
