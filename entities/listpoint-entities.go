package entities

type Model_listpoint struct {
	Lispoint_id      int    `json:"lispoint_id"`
	Lispoint_code    string `json:"lispoint_code"`
	Lispoint_name    string `json:"lispoint_name"`
	Lispoint_point   int    `json:"lispoint_point"`
	Lispoint_display int    `json:"lispoint_display"`
	Lispoint_create  string `json:"lispoint_create"`
	Lispoint_update  string `json:"lispoint_update"`
}
type Controller_listpointsave struct {
	Page             string `json:"page" validate:"required"`
	Sdata            string `json:"sdata" validate:"required"`
	Lispoint_id      int    `json:"lispoint_id"`
	Lispoint_code    string `json:"lispoint_code" validate:"required"`
	Lispoint_name    string `json:"lispoint_name" validate:"required"`
	Lispoint_point   int    `json:"lispoint_point" validate:"required"`
	Lispoint_display int    `json:"lispoint_display" validate:"required"`
}
