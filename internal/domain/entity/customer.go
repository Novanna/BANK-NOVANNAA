package entity

import (
	"strconv"
	"time"
)

type Customer struct {
	ID           int       `gorm:"column:id;primary_key" json:"id"`
	NamaLengkap  string    `gorm:"column:nama_lengkap" json:"nama_lengkap"`
	Alamat       string    `gorm:"column:alamat" json:"alamat"`
	TanggalLahir string    `gorm:"column:tanggal_lahir" json:"tanggal_lahir"`
	TempatLahir  string    `gorm:"column:tempat_lahir" json:"tempat_lahir"`
	JenisKelamin string    `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	NoKTP        int64     `gorm:"column:no_ktp;unique" json:"no_ktp"`
	NoHP         int64     `gorm:"column:no_hp" json:"no_hp"`
	CreatedAt    time.Time `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updatedAt;autoUpdateTime" json:"updatedAt"`
}

type CustomerViewModel struct {
	ID           int    `gorm:"column:id;primary_key" json:"id"`
	NamaLengkap  string `gorm:"column:nama_lengkap" json:"nama_lengkap"`
	Alamat       string `gorm:"column:alamat" json:"alamat"`
	TanggalLahir string `gorm:"column:tanggal_lahir" json:"tanggal_lahir"`
	TempatLahir  string `gorm:"column:tempat_lahir" json:"tempat_lahir"`
	JenisKelamin string `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	NoKTP        int64  `gorm:"column:no_ktp;unique" json:"no_ktp"`
	NoHP         int64  `gorm:"column:no_hp" json:"no_hp"`
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
	intKTP := int(c.NoKTP)
	stringKTP := strconv.Itoa(intKTP)
	count := len([]rune(stringKTP))
	if count != 16 {
		errorMessages["no_ktp_must16"] = "No. KTP harus 16 angka"
	}

	return errorMessages
}
