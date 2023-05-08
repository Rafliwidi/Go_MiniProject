package DataPerikananController

import (
	"fmt"
	"html/template"
	"net/http"
	
	"github.com/config"
	
	"github.com/entities"
)

var DataPerikananmodel = models.newDataPerikananmodel()


func Index(response http.ResponseWriter, request *http.Request) {

	temp, err := template.ParseFiles("views/DataPerikanan/index.html")
	if err != nil {
		panic(err)
	} 
	temp.Execute(response, nil)

}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/DataPerikanan/add.html")
		if err != nil {
		panic(err)
	}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var DataPerikanan entities.DataPerikanan 
		DataPerikanan.NamaLengkap = request.Form.Get("Nama_Lengkap")
		DataPerikanan.NIK = request.Form.Get("NIK")
		DataPerikanan.UkuranKapal = request.Form.Get("Ukuran_Kapal")
		DataPerikanan.AlatTangkap= request.Form.Get("Alat_Tangkap")
		DataPerikanan.JumlahTangkapan = request.Form.Get("Jumlah_Tangkapan")
		DataPerikanan.Alamat = request.Form.Get("Alamat")
		DataPerikanan.Nohp = request.Form.Get("no_hp")

		DataPerikananmodel.Create(DataPerikanan)
		data := map[string]interface{} {
		"pesan: "Data Anda Berhasil Disimpan",
		}
	
		temp, _ := template.ParseFiles("views/DataPerikanan/add.html")

		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

}

func Delete(response http.ResponseWriter, request *http.Request) {

}
