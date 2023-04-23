package model

import "time"

type Mahasiswa struct {
	Nama    string `json:"nama"`
	Usia    int    `json:"usia"`
	Gender  int    `json:"gender"`
	Jurusan string `json:"jurusan"`
	Hobi    string `json:"hobi"`
}

type GetMahasiswa struct {
	ID         int       `json:"id"`
	Nama       string    `json:"nama"`
	Usia       int       `json:"usia"`
	Gender     int       `json:"gender"`
	TanggalReg time.Time `json:"tanggal_reg"`
	Jurusan    []string  `json:"jurusan"`
	Hobi       []string  `json:"hobi"`
}
