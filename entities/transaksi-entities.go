package entities

type Model_invoice struct {
	Invoice_id          string `json:"invoice_id"`
	Invoice_date        string `json:"invoice_date"`
	Invoice_round       int    `json:"invoice_round"`
	Invoice_totalbet    int    `json:"invoice_totalbet"`
	Invoice_totalwin    int    `json:"invoice_totalwin"`
	Invoice_nmpoin      string `json:"invoice_nmpoin"`
	Invoice_status      string `json:"invoice_status"`
	Invoice_status_css  string `json:"invoice_status_css"`
	Invoice_card_result string `json:"invoice_card_result"`
	Invoice_card_win    string `json:"invoice_card_win"`
}
type Model_invoicedetail struct {
	Invoicedetail_id     string `json:"invoicedetail_id"`
	Invoicedetail_date   string `json:"invoicedetail_date"`
	Invoicedetail_round  int    `json:"invoicedetail_round"`
	Invoicedetail_bet    int    `json:"invoicedetail_bet"`
	Invoicedetail_win    int    `json:"invoicedetail_win"`
	Invoicedetail_status string `json:"invoicedetail_status"`
}

type Controller_invoice struct {
	Invoice_company  string `json:"invoice_company" validate:"required"`
	Invoice_username string `json:"invoice_username" validate:"required"`
}
type Controller_invoicedetail struct {
	Invoice_id      string `json:"invoice_id" validate:"required"`
	Invoice_company string `json:"invoice_company" validate:"required"`
}
type Controller_transaksisave struct {
	Transaksi_company       string `json:"transaksi_company" validate:"required"`
	Transaksi_username      string `json:"transaksi_username" validate:"required"`
	Transaksi_roundgameall  int    `json:"transaksi_roundgameall"`
	Transaksi_roundbet      int    `json:"transaksi_roundbet"`
	Transaksi_bet           int    `json:"transaksi_bet"`
	Transaksi_cbefore       int    `json:"transaksi_cbefore"`
	Transaksi_cafter        int    `json:"transaksi_cafter"`
	Transaksi_win           int    `json:"transaksi_win"`
	Transaksi_codepoin      string `json:"transaksi_codepoin"`
	Transaksi_resultcardwin string `json:"transaksi_resultcardwin" `
	Transaksi_status        string `json:"transaksi_status" validate:"required"`
}

// idtransaksi, resulcard_win string, round_bet, bet, c_before, c_after, win, idpoin int
type Controller_transaksidetailsave struct {
	Transaksidetail_company       string `json:"transaksidetail_company" validate:"required"`
	Transaksidetail_idtransaksi   string `json:"transaksidetail_idtransaksi" validate:"required"`
	Transaksidetail_roundbet      int    `json:"transaksidetail_roundbet"`
	Transaksidetail_bet           int    `json:"transaksidetail_bet"`
	Transaksidetail_cbefore       int    `json:"transaksidetail_cbefore"`
	Transaksidetail_cafter        int    `json:"transaksidetail_cafter"`
	Transaksidetail_win           int    `json:"transaksidetail_win"`
	Transaksidetail_codepoin      string `json:"transaksidetail_codepoin"`
	Transaksidetail_resultcardwin string `json:"transaksidetail_resultcardwin"`
	Transaksidetail_status        string `json:"transaksidetail_status" validate:"required"`
}
