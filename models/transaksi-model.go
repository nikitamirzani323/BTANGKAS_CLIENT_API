package models

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/helpers"
	"github.com/nleeper/goment"
)

func Save_transaksi(idcompany, username, status string, round_bet, bet, c_before, c_after, win, idpoin int) (helpers.Responsetransaksi, error) {
	var res helpers.Responsetransaksi
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	_, _, tbl_trx_transaksi, tbl_trx_transaksidetail := Get_mappingdatabase(idcompany)
	sql_insert := `
			insert into
			` + tbl_trx_transaksi + ` (
				idtransaksi , idcompany, datetransaksi, 
				username_client, roundbet,  resultcard, 
				create_transaksi, createdate_transaksi 
			) values (
				$1, $2, $3, 
				$4, $5, $6,     
				$7, $8   
			)
			`

	field_column := tbl_trx_transaksi + tglnow.Format("YYYY") + tglnow.Format("MM")
	idrecord_counter := Get_counter(field_column)
	idrecrodparent_value := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
	date_transaksi := tglnow.Format("YYYY-MM-DD HH:mm:ss")
	resultcard := _GenerateCard()
	log.Println(resultcard)
	flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_trx_transaksi, "INSERT",
		idrecrodparent_value, idcompany, date_transaksi,
		username, 0, resultcard,
		"SYSTEM", date_transaksi)

	if flag_insert {
		msg = "Succes"
		if round_bet == 1 {
			sql_insertdetail := `
				insert into
				` + tbl_trx_transaksidetail + ` (
					idtransaksidetail, idtransaksi , roundbet_detail, 
					bet, credit_before,  credit_after, 
					win, idpoin, resultcard_win, status_transaksidetail, 
					create_transaksidetail, createdate_transaksidetail  
				) values (
					$1, $2, $3, 
					$4, $5, $6,     
					$7, $8, $9, $10, 
					$11, $12   
				)
			`

			fielddetail_column := tbl_trx_transaksidetail + tglnow.Format("YYYY") + tglnow.Format("MM")
			idrecorddetail_counter := Get_counter(fielddetail_column)
			idrecroddetail_value := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecorddetail_counter)
			flag_insert, msg_insert := Exec_SQL(sql_insertdetail, tbl_trx_transaksidetail, "INSERT",
				idrecroddetail_value, idrecrodparent_value, round_bet,
				bet, c_before, c_after,
				win, idpoin, "", status,
				"SYSTEM", tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				msg = "Succes"
			} else {
				fmt.Println(msg_insert)
			}
		}

	} else {
		fmt.Println(msg_insert)
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Idtransaksi = idrecrodparent_value
	res.Card_game = resultcard
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_transaksidetail(idtransaksi, resulcard_win, status string, round_bet, bet, c_before, c_after, win, idpoin int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	sql_insert := `
			insert into
			` + configs.DB_tbl_trx_transaksidetail + ` (
				idtransaksidetail, idtransaksi , roundbet_detail, 
				bet, credit_before,  credit_after, 
				win, idpoin, resultcard_win, status_transaksidetail, 
				create_transaksidetail, createdate_transaksidetail  
			) values (
				$1, $2, $3, 
				$4, $5, $6,     
				$7, $8, $9, $10, 
				$11, $12    
			)
			`

	field_column := configs.DB_tbl_trx_transaksidetail + tglnow.Format("YYYY") + tglnow.Format("MM")
	idrecord_counter := Get_counter(field_column)
	idrecrod_value := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
	flag_insert, msg_insert := Exec_SQL(sql_insert, configs.DB_tbl_trx_transaksidetail, "INSERT",
		idrecrod_value, idtransaksi, round_bet,
		bet, c_before, c_after,
		win, idpoin, resulcard_win, status,
		"SYSTEM", tglnow.Format("YYYY-MM-DD HH:mm:ss"))

	if flag_insert {
		msg = "Succes"
	} else {
		fmt.Println(msg_insert)
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func _GenerateCard() string {
	var a [7]int
	min := 0
	max := 54
	result := ""
	for i := 0; i < 7; i++ {
		var n = rand.Intn(max-min) + min
		a[i] = n
		if i == 6 {
			result += strconv.Itoa(n)
		} else {
			result += strconv.Itoa(n) + ","
		}

	}
	return result
}
