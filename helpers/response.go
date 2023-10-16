package helpers

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
	Time    string      `json:"time"`
}
type Responsepaging struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Perpage     int         `json:"perpage"`
	Totalrecord int         `json:"totalrecord"`
	Time        string      `json:"time"`
}
type Responsetransaksi struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Idtransaksi string      `json:"idtransaksi"`
	Card_game   interface{} `json:"card_game"`
	Card_length int         `json:"card_length"`
	Time        string      `json:"time"`
}
type Responsecompany struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Record   interface{} `json:"record"`
	Listcurr interface{} `json:"listcurr"`
	Time     string      `json:"time"`
}
type Responsecompanyadmin struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Listcompany interface{} `json:"listcompany"`
	Listrule    interface{} `json:"listrule"`
	Time        string      `json:"time"`
}
type Responsecompanyadminrule struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Listcompany interface{} `json:"listcompany"`
	Time        string      `json:"time"`
}
type Responseagenrule struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Record   interface{} `json:"record"`
	Listagen interface{} `json:"listagen"`
	Time     string      `json:"time"`
}
type ResponseAdmin struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Record   interface{} `json:"record"`
	Listrule interface{} `json:"listruleadmin"`
	Time     string      `json:"time"`
}
type ResponseEmployee struct {
	Status          int         `json:"status"`
	Message         string      `json:"message"`
	Record          interface{} `json:"record"`
	Listdepartement interface{} `json:"listdepartement"`
	Time            string      `json:"time"`
}
type ErrorResponse struct {
	Field string
	Tag   string
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
