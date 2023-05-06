package main

import (
	"net/http"

	"github.com/controllers/DataPerikananController"
)

func main() {

	// Data Perikanan
	http.HandleFunc("/", DataPerikananController.Index)
	http.HandleFunc("/DataPerikanan", DataPerikananController.Index)
	http.HandleFunc("/DataPerikanan/index", DataPerikananController.Index)
	http.HandleFunc("/DataPerikanan/add", DataPerikananController.Add)
	http.HandleFunc("/DataPerikanan/edit", DataPerikananController.Edit)
	http.HandleFunc("/DataPerikanan/delete", DataPerikananController.Delete)

	http.ListenAndServe(":8000", nil)
}
