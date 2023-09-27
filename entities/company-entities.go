package entities

type Model_company struct {
	Company_id         string `json:"company_id"`
	Company_startjoin  string `json:"company_startjoin"`
	Company_endjoin    string `json:"company_endjoin"`
	Company_name       string `json:"company_name"`
	Company_idcurr     string `json:"company_idcurr"`
	Company_nmowner    string `json:"company_nmowner"`
	Company_phoneowner string `json:"company_phoneowner"`
	Company_emailowner string `json:"company_emailowner"`
	Company_url        string `json:"company_url"`
	Company_status     string `json:"company_status"`
	Company_status_css string `json:"company_status_css"`
	Company_create     string `json:"company_create"`
	Company_update     string `json:"company_update"`
}
type Model_company_listbet struct {
	Companylistbet_id     int     `json:"companylistbet_id"`
	Companylistbet_minbet float64 `json:"companylistbet_minbet"`
	Companylistbet_create string  `json:"companylistbet_create"`
	Companylistbet_update string  `json:"companylistbet_update"`
}
type Model_company_conf struct {
	Companyconf_id     int    `json:"companyconf_id"`
	Companyconf_idbet  int    `json:"companyconf_idbet"`
	Companyconf_nmpoin string `json:"companyconf_nmpoin"`
	Companyconf_poin   int    `json:"companyconf_poin"`
	Companyconf_create string `json:"companyconf_create"`
	Companyconf_update string `json:"companyconf_update"`
}

type Model_companyshare struct {
	Company_id   string `json:"company_id"`
	Company_name string `json:"company_name"`
}

type Controller_companysave struct {
	Page               string `json:"page" validate:"required"`
	Sdata              string `json:"sdata" validate:"required"`
	Company_id         string `json:"company_id"`
	Company_name       string `json:"company_name" validate:"required"`
	Company_idcurr     string `json:"company_idcurr" validate:"required"`
	Company_nmowner    string `json:"company_nmowner" validate:"required"`
	Company_phoneowner string `json:"company_phoneowner" validate:"required"`
	Company_emailowner string `json:"company_emailowner" `
	Company_url        string `json:"company_url" validate:"required"`
	Company_status     string `json:"company_status" validate:"required"`
}

type Controller_companylistbetsave struct {
	Page                     string `json:"page" validate:"required"`
	Sdata                    string `json:"sdata" validate:"required"`
	Companylistbet_id        int    `json:"companylistbet_id"`
	Companylistbet_idcompany string `json:"companylistbet_idcompany" validate:"required"`
	Companylistbet_minbet    int    `json:"companylistbet_minbet" validate:"required"`
}
type Controller_companylistbet struct {
	Company_id string `json:"company_id" validate:"required"`
}
type Controller_companyconfpoint struct {
	Company_idbet int    `json:"company_idbet" validate:"required"`
	Company_id    string `json:"company_id" validate:"required"`
}
type Controller_companyconfpointsave struct {
	Page                       string `json:"page" validate:"required"`
	Sdata                      string `json:"sdata" validate:"required"`
	Companyconfpoint_id        int    `json:"companyconfpoint_id"`
	Companyconfpoint_idbet     int    `json:"companyconfpoint_idbet" validate:"required"`
	Companyconfpoint_idcompany string `json:"companyconfpoint_idcompany" validate:"required"`
	Companyconfpoint_point     int    `json:"companyconfpoint_point" `
}
