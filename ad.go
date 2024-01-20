package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running")

	http.HandleFunc("/Signup", SignUp)
	http.HandleFunc("/Login", Login)
	http.HandleFunc("/New", New)

	http.ListenAndServe(":2222",nil)
}

func SignUp(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Signup")	
}

func Login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Login")
}

func New(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("New")
}