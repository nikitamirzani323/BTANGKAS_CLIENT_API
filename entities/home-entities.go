package entities

type CheckToken struct {
	Token string `json:"token" validate:"required"`
}
type Controller_listbet struct {
	Client_token   string `json:"client_token" validate:"required"`
	Client_company string `json:"client_company" validate:"required"`
}
