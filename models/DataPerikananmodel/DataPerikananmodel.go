package DataPerikananmodel

import (
	"database/sql"
	"fmt"

	"github.com/config"

	"github.com/entities"
)

type DataPerikananmodel struct {
	con *sql.DB
}

func newDataPerikananmodel() *DataPerikananmodel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &DataPerikananmodel{
		conn: conn,
	}
}

func (p *DataPerikananmodel) Create(DataPerikanan entities.DataPerikanan) bool {

	result, err := p.conn.Exec("insert into DataPerikanan (Nama_Lengkap, NIK, Ukuran_Kapal, Alat_Tangkap, Jumlah_Tangkapan, Alamat, no_hp) values(?,?,?,?,?,?,?)",
		DataPerikanan.NamaLengkap, DataPerikanan.NIK, DataPerikanan.UkuranKapal, DataPerikanan.AlatTangkap, DataPerikanan.JumlahTangkapan, DataPerikanan.Alamat, DataPerikanan.Nohp)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.lastInsertId()

	return lastInsertId > 0
}
