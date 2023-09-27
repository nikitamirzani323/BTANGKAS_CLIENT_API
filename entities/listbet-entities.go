package entities

type Model_lisbet struct {
	Lisbet_id     int     `json:"lisbet_id"`
	Lisbet_minbet float64 `json:"lisbet_minbet"`
	Lisbet_create string  `json:"lisbet_create"`
	Lisbet_update string  `json:"lisbet_update"`
}
type Controller_listbetsave struct {
	Page          string  `json:"page" validate:"required"`
	Sdata         string  `json:"sdata" validate:"required"`
	Lisbet_id     int     `json:"lisbet_id"`
	Lisbet_minbet float64 `json:"lisbet_minbet" validate:"required"`
}
