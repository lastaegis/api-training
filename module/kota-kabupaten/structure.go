package kota_kabupaten

type StructKotaKabupaten struct {
	ID             int    `json:"id" db:"ID"`
	PROVINSI       string `json:"provinsi" db:"PROVINSI"`
	KOTA_KABUPATEN string `json:"kota_kabupaten" db:"KOTA_KABUPATEN"`
}
