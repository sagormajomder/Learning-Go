package main

import (
	"fmt"
	"net/http"
)

 func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello World")
 }
 func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "I'm sagor. I'm software engineer")
 }

func main(){
	mux := http.NewServeMux() //router

	mux.HandleFunc("/hello", helloHandler ) //route

	mux.HandleFunc("/about", aboutHandler ) //router

	fmt.Println("Server running on :3000")

	err:= http.ListenAndServe(":3000", mux) // start server

	if err !=nil {
		fmt.Println("Error starting the server", err)
	}
}