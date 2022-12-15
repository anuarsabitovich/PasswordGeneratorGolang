package main

import (
	"fmt"
	"log"
	"net/http"
	"password/handlers"
)

func main() {
	fmt.Println("hello")
	log.Print("http://localhost:8080")
	http.HandleFunc("/", handlers.MainPage)
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))
	log.Fatalln(http.ListenAndServe(":8080", nil))

}
