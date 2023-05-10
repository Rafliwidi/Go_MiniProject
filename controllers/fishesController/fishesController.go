package fishesController

import (
	"html/template"
	"strconv"

	"net/http"

	"mini-project/entities"
	"mini-project/libraries"

	"mini-project/models/fishesmodel"
)

var validation = libraries.NewValidation()

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

		var data = make(map[string]interface{})
		
		vErrors := validation.Struct(fishes)

		if vErrors != nil {
			data["fishes"] = fishes
			data["validation"] =vErrors
		} else {
			data["pesan"] = "Data Fishes Berhasil Disimpan"
			fishesmodel.Create(fishes)
		
		}
		
		temp, _ := template.ParseFiles("views/fishes/add.html")

		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var fishes entities.Fishes
		fishesmodel.Find(id, &fishes)

		data := map[string]interface{}{
			"fishes": fishes,
		
		}

		temp, err := template.ParseFiles("views/fishes/edit.html")

		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {
		
		request.ParseForm()

		var fishes entities.Fishes
		fishes.Id, _ = strconv.ParseInt(request.Form.Get("id"),10 , 64)
		fishes.NamaLengkap = request.Form.Get("Nama_Lengkap")
		fishes.NIK = request.Form.Get("NIK")
		fishes.UkuranKapal = request.Form.Get("Ukuran_Kapal")
		fishes.AlatTangkap = request.Form.Get("Alat_Tangkap")
		fishes.JumlahTangkapan = request.Form.Get("Jumlah_Tangkapan")
		fishes.Alamat = request.Form.Get("Alamat")
		fishes.Nohp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})
		
		vErrors := validation.Struct(fishes)

		if vErrors != nil {
			data["validation"] =vErrors
		} else {
			data["pesan"] = "Data Fishes Berhasil Diperbarui"
			fishesmodel.Update(fishes)
		
		}
		
		temp, _ := template.ParseFiles("views/fishes/add.html")

		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id,_ := strconv.ParseInt(queryString.Get("id"),10, 64)

	fishesmodel.Delete(id)

	http.Redirect(response, request, "/fishes", http.StatusSeeOther)




}
