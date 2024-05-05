package main

import (
	"go-web-native/config"
	"go-web-native/controllers/homecontroller"
	"go-web-native/controllers/categorycontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// Panggil Home Page

	http.HandleFunc("/", homecontroller.Welcome)

	// Categories Page

	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)


	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}