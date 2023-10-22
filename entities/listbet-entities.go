package entities

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
