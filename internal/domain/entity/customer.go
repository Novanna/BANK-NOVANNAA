package entity

import (
	"time"
)

type Customer struct {
	ID           int       `json:"id"`
	NamaLengkap  string    `json:"nama_lengkap"`
	Alamat       string    `json:"alamat"`
	TanggalLahir string    `json:"tanggal_lahir"`
	TempatLahir  string    `json:"tempat_lahir"`
	JenisKelamin string    `json:"jenis_kelamin"`
	NoKTP        int64     `json:"no_ktp"`
	NoHP         int64     `json:"no_hp"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type CustomerViewModel struct {
	ID           int    `json:"id"`
	NamaLengkap  string `json:"nama_lengkap"`
	Alamat       string `json:"alamat"`
	TanggalLahir string `json:"tanggal_lahir"`
	TempatLahir  string `json:"tempat_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoKTP        int64  `json:"no_ktp"`
	NoHP         int64  `json:"no_hp"`
}

func (c *CustomerViewModel) Validate() map[string]string {
	var errorMessages = make(map[string]string)

	if c.NamaLengkap == "" || c.NamaLengkap == "null" {
		errorMessages["nama_lengkap_null"] = "Nama lengkap belum diisi"
	}
	if c.Alamat == "" || c.Alamat == "null" {
		errorMessages["alamat_null"] = "Alamat lengkap belum diisi"
	}
	if c.JenisKelamin == "" || c.JenisKelamin == "null" {
		errorMessages["jenis_kelamin_null"] = "Gender belum diisi"
	}
	if c.NoKTP-c.NoKTP != 0 {
		errorMessages["no_ktp_mustint"] = "No. KTP harus angka"
	}
	stringKTP := string(c.NoKTP)
	if len([]rune(stringKTP)) != 16 {
		errorMessages["no_ktp_must16"] = "No. KTP harus 16 angka"
	}

	return errorMessages
}
