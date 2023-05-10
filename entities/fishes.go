package entities

type Fishes struct {
	Id              int64
	NamaLengkap    	string 	`validate:"required"` 
	NIK             string	`validate:"required"`
	UkuranKapal     string	`validate:"required"` 
	AlatTangkap     string	`validate:"required"`  
	JumlahTangkapan string	`validate:"required"`  
	Alamat          string	`validate:"required"`
	Nohp            string	`validate:"required"` 
}
