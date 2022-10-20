package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Awesome Golang Application is ready to go !!!")
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("name")
	if len(username) != 0 {
		fmt.Fprintf(w, "Hello from Awesome Golang Application, Have a nice day "+username+"!!!")
	} else {
		fmt.Fprintf(w, "Hello from Awesome Golang Application")
	}
}
func setupRoutes() {
	// Default Router ,you can add endpoints
	http.HandleFunc("/", homePage)
	http.HandleFunc("/sayHello/", sayHello)
}
func main() {
	fmt.Println("Your Golang Application is running in port 3000")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}
