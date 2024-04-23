package entity

import "time"

type TrxLaundry struct {
	Id              int
	NoTrx           string
	Customer_id     int
	Tanggal_masuk   time.Time
	Tanggal_selesai time.Time
	Penerima        string
	Total_biaya     int
}

type TrxLaundryDetail struct {
	Id             int
	Trx_laundry_id int
	Layanan_id     int
	Jumlah         int
}
