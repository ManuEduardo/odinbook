package main

import (
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Email    string
	Password string
}

func main() {

	user_login := User{}

	loginHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("login.html"))
		templ.Execute(w, nil)
	}

	requestLoginHandler := func(w http.ResponseWriter, r *http.Request) {
		email := r.PostFormValue("email-user")
		password := r.PostFormValue("password-user")
		log.Printf("Email: %v | Password: %v", email, password)

		user_login.Email = email
		user_login.Password = password

		templ := template.Must(template.ParseFiles("home.html"))
		templ.Execute(w, user_login)
	}

	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/home", requestLoginHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
