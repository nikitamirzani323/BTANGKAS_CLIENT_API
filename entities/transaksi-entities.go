package entities

type Controller_transaksisave struct {
	Transaksi_company  string `json:"transaksi_company" validate:"required"`
	Transaksi_username string `json:"transaksi_username" validate:"required"`
	Transaksi_roundbet int    `json:"transaksi_roundbet"`
	Transaksi_bet      int    `json:"transaksi_bet"`
	Transaksi_cbefore  int    `json:"transaksi_cbefore"`
	Transaksi_cafter   int    `json:"transaksi_cafter"`
	Transaksi_win      int    `json:"transaksi_win"`
	Transaksi_idpoin   int    `json:"transaksi_idpoin"`
	Transaksi_status   string `json:"transaksi_status" validate:"required"`
}

// idtransaksi, resulcard_win string, round_bet, bet, c_before, c_after, win, idpoin int
type Controller_transaksidetailsave struct {
	Transaksidetail_idtransaksi   string `json:"transaksidetail_idtransaksi" validate:"required"`
	Transaksidetail_roundbet      int    `json:"transaksidetail_roundbet"`
	Transaksidetail_bet           int    `json:"transaksidetail_bet"`
	Transaksidetail_cbefore       int    `json:"transaksidetail_cbefore"`
	Transaksidetail_cafter        int    `json:"transaksidetail_cafter"`
	Transaksidetail_win           int    `json:"transaksidetail_win"`
	Transaksidetail_idpoin        int    `json:"transaksidetail_idpoin"`
	Transaksidetail_resultcardwin string `json:"transaksidetail_resultcardwin"`
	Transaksidetail_status        string `json:"transaksidetail_status" validate:"required"`
}
