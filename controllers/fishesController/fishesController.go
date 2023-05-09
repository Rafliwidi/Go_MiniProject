package fishesController

import (
	"html/template"

	"net/http"

	"mini-project/entities"

	models "mini-project/models/fishesmodel"
)

var fishesmodel = models.Newfishesmodel()

func Index(response http.ResponseWriter, request *http.Request) {

	fishes, _ := fishesmodel.FindAll()

	data := map[string]interface{}{
		"fishes": fishes,
	}

	temp, err := template.ParseFiles("views/fishes/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)

}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/fishes/add.html")

		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var fishes entities.Fishes
		fishes.NamaLengkap = request.Form.Get("Nama_Lengkap")
		fishes.NIK = request.Form.Get("NIK")
		fishes.UkuranKapal = request.Form.Get("Ukuran_Kapal")
		fishes.AlatTangkap = request.Form.Get("Alat_Tangkap")
		fishes.JumlahTangkapan = request.Form.Get("Jumlah_Tangkapan")
		fishes.Alamat = request.Form.Get("Alamat")
		fishes.Nohp = request.Form.Get("no_hp")

		fishesmodel.Create(fishes)

		data := map[string]interface{}{
			"pesan": "Data Anda Berhasil Disimpan",
		}

		temp, _ := template.ParseFiles("views/fishes/add.html")

		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

}

func Delete(response http.ResponseWriter, request *http.Request) {

}
