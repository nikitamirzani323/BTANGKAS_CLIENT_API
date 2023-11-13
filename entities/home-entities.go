package entities

type CheckToken struct {
	Token string `json:"token" validate:"required"`
}
type Controller_listbet struct {
	Client_token   string `json:"client_token" validate:"required"`
	Client_company string `json:"client_company" validate:"required"`
}
type Model_lisbet struct {
	Lisbet_id     int         `json:"lisbet_id"`
	Lisbet_minbet float64     `json:"lisbet_minbet"`
	Lisbet_conf   interface{} `json:"lisbet_config"`
}

type Model_lispoin struct {
	Lispoin_id     string `json:"lispoin_id"`
	Lispoin_nmpoin string `json:"lispoin_nmpoin"`
	Lispoin_poin   int    `json:"lispoin_poin"`
}
