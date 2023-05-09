package main

import (
	"net/http"

	"mini-project/controllers/fishesController"
)

func main() {

	// fishes
	http.HandleFunc("/", fishesController.Index)
	http.HandleFunc("/fishes", fishesController.Index)
	http.HandleFunc("/fishes/index", fishesController.Index)
	http.HandleFunc("/fishes/add", fishesController.Add)
	http.HandleFunc("/fishes/edit", fishesController.Edit)
	http.HandleFunc("/fishes/delete", fishesController.Delete)

	http.ListenAndServe(":8000", nil)
}
