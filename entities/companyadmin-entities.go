package entities

type Model_companyadmin struct {
	Companyadmin_id         string `json:"companyadmin_id"`
	Companyadmin_idrule     int    `json:"companyadmin_idrule"`
	Companyadmin_idcompany  string `json:"companyadmin_idcompany"`
	Companyadmin_tipe       string `json:"companyadmin_tipe"`
	Companyadmin_nmrule     string `json:"companyadmin_nmrule"`
	Companyadmin_username   string `json:"companyadmin_username"`
	Companyadmin_ipaddress  string `json:"companyadmin_ipaddress"`
	Companyadmin_lastlogin  string `json:"companyadmin_lastlogin"`
	Companyadmin_name       string `json:"companyadmin_name"`
	Companyadmin_phone1     string `json:"companyadmin_phone1"`
	Companyadmin_phone2     string `json:"companyadmin_phone2"`
	Companyadmin_status     string `json:"companyadmin_status"`
	Companyadmin_status_css string `json:"companyadmin_status_css"`
	Companyadmin_create     string `json:"companyadmin_create"`
	Companyadmin_update     string `json:"companyadmin_update"`
}

type Controller_companyadminsave struct {
	Sdata                  string `json:"sdata" validate:"required"`
	Page                   string `json:"page" validate:"required"`
	Companyadmin_id        string `json:"companyadmin_id"`
	Companyadmin_idrule    int    `json:"companyadmin_idrule" validate:"required"`
	Companyadmin_idcompany string `json:"companyadmin_idcompany" validate:"required"`
	Companyadmin_username  string `json:"companyadmin_username" validate:"required"`
	Companyadmin_password  string `json:"companyadmin_password" `
	Companyadmin_name      string `json:"companyadmin_name" validate:"required"`
	Companyadmin_phone1    string `json:"companyadmin_phone1" validate:"required"`
	Companyadmin_phone2    string `json:"companyadmin_phone2"`
	Companyadmin_status    string `json:"companyadmin_status" validate:"required"`
}
type Controller_companygroupcompany struct {
	Company_id string `json:"company_id" validate:"required"`
}
