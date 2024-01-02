package general_structure

type ResponseGet struct {
	Status    int32       `json:"status"`
	Message   string      `json:"message"`
	TotalData int32       `json:"total_data"`
	Data      interface{} `json:"data"`
}

type ResponsePost struct {
	Status  int32  `json:"status"`
	Message string `json:"message"`
}

type ResponsePut struct {
	Status  int32  `json:"status"`
	Message string `json:"message"`
}

type ResponseDelete struct {
	Status  int32  `json:"status"`
	Message string `json:"message"`
}
