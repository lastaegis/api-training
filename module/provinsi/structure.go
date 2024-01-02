package provinsi

type Provinsi struct {
	ID       int32  `json:"id" db:"ID"`
	PROVINSI string `json:"provinsi" db:"PROVINSI"`
}
