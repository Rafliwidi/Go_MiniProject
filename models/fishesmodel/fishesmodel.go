package models

import (
	"database/sql"
	
	"fmt"

	"mini-project/config"

	"mini-project/entities"
)

type fishesmodel struct {
	conn *sql.DB
}

func Newfishesmodel() *fishesmodel {
	conn, err := config.DBconnection()
	if err != nil {
		panic(err)
	}

	return &fishesmodel{
		conn: conn,
	}
}

func (p *fishesmodel) FindAll() ([]entities.Fishes, error) {

	rows, err := p.conn.Query("select * from fishes")
	if err != nil {
		return []entities.Fishes{}, err
	}
	defer rows.Close()

	var datafishes []entities.Fishes
	for rows.Next() {
		var fishes entities.Fishes
		rows.Scan(&fishes.Id,
			&fishes.NamaLengkap,
			&fishes.NIK,
			&fishes.UkuranKapal,
			&fishes.AlatTangkap,
			&fishes.JumlahTangkapan,
			&fishes.Alamat,
			&fishes.Nohp)

		datafishes = append(datafishes, fishes)
	}

	return datafishes, nil
}

func (p *fishesmodel) Create(fishes entities.Fishes) bool {

	result, err := p.conn.Exec("insert into fishes (nama_lengkap, nik, ukuran_kapal, alat_tangkap, jumlah_tangkapan, alamat, no_hp) values(?,?,?,?,?,?,?)",
		fishes.NamaLengkap, fishes.NIK, fishes.UkuranKapal, fishes.AlatTangkap, fishes.JumlahTangkapan, fishes.Alamat, fishes.Nohp)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *fishesmodel) Find(id int64, fishes *entities.Fishes) error {
	
	return p.conn.QueryRow("select *From fishes where id = ?", id).Scan(&fishes.Id, 
		&fishes.NamaLengkap, 
		&fishes.NIK, 
		&fishes.UkuranKapal, 
		&fishes.AlatTangkap, 
		&fishes.JumlahTangkapan, 
		&fishes.Alamat, 
		&fishes.Nohp)
}
