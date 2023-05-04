package main

import (
	"net/http"
	"github.com/controllers/DataPerikananController"
)

func main() {
	// 1. Homepage
	http.HandleFunc("/", DataPerikananController.Add)
	
	// 2. Data Perikanan
	http.HandleFunc("/", DataPerikananController.Index)
	http.HandleFunc("/DataPerikanan", DataPerikananController.Index)
	http.HandleFunc("/DataPerikanan/index", DataPerikananController.Index)
	http.HandleFunc("/DataPerikanan/add", DataPerikananController.Add)
	http.HandleFunc("/DataPerikanan/edit", DataPerikananController.Edit)
	http.HandleFunc("/DataPerikanan/delete", DataPerikananController.Delete)

	http.ListenAndServe(":4000", nil)
}
