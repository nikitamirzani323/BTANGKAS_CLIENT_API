package entities

type Model_companyadminrule struct {
	Companyadminrule_id        int    `json:"companyadminrule_id"`
	Companyadminrule_idcompany string `json:"companyadminrule_idcompany"`
	Companyadminrule_nmrule    string `json:"companyadminrule_nmrule"`
	Companyadminrule_rule      string `json:"companyadminrule_rule"`
	Companyadminrule_create    string `json:"companyadminrule_create"`
	Companyadminrule_update    string `json:"companyadminrule_update"`
}
type Model_companyadminrule_share struct {
	Companyadminrule_id        int    `json:"companyadminrule_id"`
	Companyadminrule_idcompany string `json:"companyadminrule_idcompany"`
	Companyadminrule_nmrule    string `json:"companyadminrule_nmrule"`
}

type Controller_companyadminrulesave struct {
	Sdata                      string `json:"sdata" validate:"required"`
	Page                       string `json:"page" validate:"required"`
	Companyadminrule_id        int    `json:"companyadminrule_id" `
	Companyadminrule_idcompany string `json:"companyadminrule_idcompany" validate:"required"`
	Companyadminrule_nmrule    string `json:"companyadminrule_nmrule" validate:"required"`
	Companyadminrule_rule      string `json:"companyadminrule_rule"`
}
