package entities

type CheckToken struct {
	Token string `json:"token" validate:"required"`
}
