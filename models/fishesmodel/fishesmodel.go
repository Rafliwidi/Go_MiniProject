package models

import (
	"database/sql"
	"fmt"

	"mini-project/config"
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

func (p *fishesmodel) FindAll() ([]entities.fishes, error) {

	rows, err := p.conn.Query("select * from fishes")
	if err != nil {
		return []entities.fishes{}, err
	}
	defer rows.Close()

	var datafishes []entities.fishes
	for rows.Next() {
		var fishes entities.fishes
		rows.Scan(&fishes.id,
			&fishes.NamaLengkap,
			&fishes.NIK,
			&fishes.UkuranKapal,
			&fishes.AlatTangkap,
			&fishes.JumlahTangkapan,
			&fishes.Alamat,
			&fishes.Nohp)

		fishes = append(datafishes, fishes)
	}
}

func (p *fishesmodel) Create(fishes entities.fishes) bool {

	result, err := p.conn.Exec("insert into fishes (Nama_Lengkap, NIK, Ukuran_Kapal, Alat_Tangkap, Jumlah_Tangkapan, Alamat, no_hp) values(?,?,?,?,?,?,?)",
		fishes.NamaLengkap, fishes.NIK, fishes.UkuranKapal, fishes.AlatTangkap, fishes.JumlahTangkapan, fishes.Alamat, fishes.Nohp)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
