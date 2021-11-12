package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
	"lenslocked.com/models"
)

const (
	host     = "localhost"
	port     = 5435
	user     = "juanluisrubalcava"
	password = "my-password"
	dbname   = "lenslocked_dev"
	sslmode  = "disable"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s"+
		" sslmode=disable",
		host, port, user, dbname)
	us, err := models.NewUserService(psqlInfo)
	must(err)
	defer us.Close()
	us.AutoMigrate()

	usersC := controllers.NewUsers(us)
	staticC := controllers.NewStatic()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
