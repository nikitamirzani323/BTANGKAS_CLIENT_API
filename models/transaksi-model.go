package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/db"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/helpers"
	"github.com/nleeper/goment"
)

type Card_Strc struct {
	TypePattern string `json:"typepattern"`
	Pattern     string `json:"pattern"`
	Total       int    `json:"total"`
}

func Fetch_invoice(idcompany, username string) (helpers.Response, error) {
	var obj entities.Model_invoice
	var arraobj []entities.Model_invoice
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	_, _, tbl_trx_transaksi, _ := Get_mappingdatabase(idcompany)

	tglnow, _ := goment.New()
	tglskrg := tglnow.Format("YYYY-MM-DD HH:mm:ss")
	tglbefore := tglnow.Add(-31, "days").Format("YYYY-MM-DD HH:mm:ss")
	fmt.Println("tgl skrg :" + tglskrg)
	fmt.Println("tgl before :" + tglbefore)

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "idtransaksi , to_char(COALESCE(createdate_transaksi,now()), 'YYYY-MM-DD HH24:MI:SS') as datetransaksi,  "
	sql_select += "roundbet, total_bet, total_win,   "
	sql_select += "card_codepoin, card_result, card_win    "
	sql_select += "FROM " + tbl_trx_transaksi + " "
	sql_select += "WHERE createdate_transaksi >='" + tglbefore + "' "
	sql_select += "AND createdate_transaksi <='" + tglskrg + "' "
	sql_select += "ORDER BY createdate_transaksi DESC  LIMIT 31 "

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idtransaksi_db, datetransaksi_db              string
			roundbet_db, total_bet_db, total_win_db       int
			card_codepoin_db, card_result_db, card_win_db string
		)

		err = row.Scan(&idtransaksi_db, &datetransaksi_db,
			&roundbet_db, &total_bet_db, &total_win_db,
			&card_codepoin_db, &card_result_db, &card_win_db)

		helpers.ErrorCheck(err)
		status := "LOSE"
		status_css := configs.STATUS_CANCEL
		if card_win_db != "" {
			status_css = configs.STATUS_COMPLETE
			status = "WIN"
		}

		obj.Invoice_id = idtransaksi_db
		obj.Invoice_date = datetransaksi_db
		obj.Invoice_round = roundbet_db
		obj.Invoice_totalbet = total_bet_db
		obj.Invoice_totalwin = total_win_db
		obj.Invoice_card_result = card_result_db
		obj.Invoice_card_win = card_win_db
		obj.Invoice_nmpoin = _GetInfoPoint(card_codepoin_db)
		obj.Invoice_status = status
		obj.Invoice_status_css = status_css
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_transaksi(idcompany, username, status, resultcardwin, codepoin string, round_game_all, round_bet, bet, c_before, c_after, win int) (helpers.Responsetransaksi, error) {
	var res helpers.Responsetransaksi
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	_, _, tbl_trx_transaksi, tbl_trx_transaksidetail := Get_mappingdatabase(idcompany)
	sql_insert := `
			insert into
			` + tbl_trx_transaksi + ` (
				idtransaksi , idcompany, datetransaksi, 
				username_client, roundbet,  total_bet, total_win,  
				card_codepoin,  card_pattern, card_result,  card_win, 
				create_transaksi, createdate_transaksi 
			) values (
				$1, $2, $3, 
				$4, $5, $6, $7,      
				$8, $9, $10, $11,    
				$12, $13   
			)
		`

	field_column := tbl_trx_transaksi + tglnow.Format("YYYY") + tglnow.Format("MM")
	idrecord_counter := Get_counter(field_column)
	idrecrodparent_value := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
	date_transaksi := tglnow.Format("YYYY-MM-DD HH:mm:ss")

	idlistpattern := ""
	pattern := ""
	total_card := 0
	field_redis := "CLIENT_PATTERN_" + strings.ToLower(idcompany) + "_" + strings.ToLower(username)
	var card_res Card_Strc

	if round_game_all < 1 {
		fmt.Println("Database Card Pattern")
		pattern, total_card, idlistpattern = _GenerateCardRandomDB()
		card_res.TypePattern = idlistpattern
		card_res.Pattern = pattern
		card_res.Total = total_card
		helpers.SetRedis(field_redis, card_res, 30*time.Minute)
	} else {
		resultredis, flag := helpers.GetRedis(field_redis)
		jsonredis := []byte(resultredis)
		pattern_redis, _ := jsonparser.GetString(jsonredis, "pattern")
		typepattern_redis, _ := jsonparser.GetString(jsonredis, "typepattern")
		total_redis, _ := jsonparser.GetInt(jsonredis, "total")
		fmt.Println("Total Card REDIS : " + strconv.Itoa(int(total_redis)))
		fmt.Println("Total Game : " + strconv.Itoa(round_game_all))
		if int(total_redis-1) == round_game_all {
			val_pattern := helpers.DeleteRedis(field_redis)
			fmt.Printf("Redis Delete Card Pattern : %d\n", val_pattern)
			flag = false
			round_game_all = 0
		}
		if !flag {
			fmt.Println("Database Card Pattern")
			pattern, total_card, idlistpattern = _GenerateCardRandomDB()
			card_res.TypePattern = idlistpattern
			card_res.Pattern = pattern
			card_res.Total = total_card
			helpers.SetRedis(field_redis, card_res, 30*time.Minute)
		} else {
			fmt.Println("Cache Card Pattern")
			pattern = pattern_redis
			idlistpattern = typepattern_redis
			total_card = int(total_redis)
		}

	}
	fmt.Println("Total Card : " + strconv.Itoa(total_card))
	fmt.Println("Total Game : " + strconv.Itoa(round_game_all))
	fmt.Println("Type Pattern : " + idlistpattern)
	resultcard := strings.Split(pattern, "|")

	flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_trx_transaksi, "INSERT",
		idrecrodparent_value, idcompany, date_transaksi,
		username, 0, 0, 0,
		"", idlistpattern, resultcard[round_game_all], "",
		"SYSTEM", date_transaksi)

	if flag_insert {
		msg = "Succes"
		if round_bet == 1 || round_bet == 4 {
			sql_insertdetail := `
				insert into
				` + tbl_trx_transaksidetail + ` (
					idtransaksidetail, idtransaksi , roundbet_detail, 
					bet, credit_before,  credit_after, 
					win, codepoin, resultcard_win, status_transaksidetail, 
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
			flag_insertdetail, msg_insertdetail := Exec_SQL(sql_insertdetail, tbl_trx_transaksidetail, "INSERT",
				idrecroddetail_value, idrecrodparent_value, round_bet,
				bet, c_before, c_after,
				win, codepoin, resultcardwin, status,
				"SYSTEM", tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insertdetail {
				msg_insertdetail = "Succes"

				sql_update := `
					UPDATE 
					` + tbl_trx_transaksi + `  
					SET roundbet=$1, total_bet=$2,
					update_transaksi=$3, updatedate_transaksi=$4          
					WHERE idtransaksi=$5         
				`

				flag_update, msg_update := Exec_SQL(sql_update, database_listpoint_local, "UPDATE",
					round_bet, _GetTotalBet_Transaksi(tbl_trx_transaksidetail, idrecrodparent_value),
					"SYSTEM", tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecrodparent_value)

				if flag_update {
					msg = "Succes"
				} else {
					fmt.Println(msg_update)
				}
				log.Println(msg_insertdetail)
			} else {
				fmt.Println(msg_insertdetail)
			}
		}

	} else {
		fmt.Println(msg_insert)
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Idtransaksi = idrecrodparent_value
	res.Card_game = resultcard[round_game_all]
	res.Card_length = total_card
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_transaksidetail(idcompany, idtransaksi, resulcard_win, status, codepoin string, round_bet, bet, c_before, c_after, win int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	_, _, tbl_trx_transaksi, tbl_trx_transaksidetail := Get_mappingdatabase(idcompany)

	sql_insert := `
			insert into
			` + tbl_trx_transaksidetail + ` (
				idtransaksidetail, idtransaksi , roundbet_detail, 
				bet, credit_before,  credit_after, 
				win, codepoin, resultcard_win, status_transaksidetail, 
				create_transaksidetail, createdate_transaksidetail  
			) values (
				$1, $2, $3, 
				$4, $5, $6,     
				$7, $8, $9, $10, 
				$11, $12    
			)
			`

	field_column := tbl_trx_transaksidetail + tglnow.Format("YYYY") + tglnow.Format("MM")
	idrecord_counter := Get_counter(field_column)
	idrecrod_value := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
	flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_trx_transaksidetail, "INSERT",
		idrecrod_value, idtransaksi, round_bet,
		bet, c_before, c_after,
		win, codepoin, resulcard_win, status,
		"SYSTEM", tglnow.Format("YYYY-MM-DD HH:mm:ss"))

	if flag_insert {
		msg = "Succes"
		sql_update := `
				UPDATE 
				` + tbl_trx_transaksi + `  
				SET roundbet=$1, total_bet=$2, total_win=$3, card_codepoin=$4, card_win=$5, 
				update_transaksi=$6, updatedate_transaksi=$7         
				WHERE idtransaksi=$8        
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_listpoint_local, "UPDATE",
			round_bet, _GetTotalBet_Transaksi(tbl_trx_transaksidetail, idtransaksi), win, codepoin, resulcard_win,
			"SYSTEM", tglnow.Format("YYYY-MM-DD HH:mm:ss"), idtransaksi)

		if flag_update {
			msg = "Succes"
		} else {
			fmt.Println(msg_update)
		}
	} else {
		fmt.Println(msg_insert)
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}

func _GenerateCardRandomDB() (string, int, string) {
	con := db.CreateCon()
	ctx := context.Background()
	listpattern := ""
	total_detail := 0
	idlistpattern := ""
	sql_select := `SELECT
			idlistpattern
			FROM ` + configs.DB_tbl_trx_listpattern + `  
			WHERE status_listpattern='Y'     
			ORDER BY random()  LIMIT 1  
		`

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&idlistpattern); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	if idlistpattern != "" {
		sql_selecttotal := `SELECT
				COUNT(idlistpatterndetail) as total
				FROM ` + configs.DB_tbl_trx_listpatterndetail + `  
				WHERE idlistpattern=$1       
			`

		row_selecttotal := con.QueryRowContext(ctx, sql_selecttotal, idlistpattern)
		switch e := row_selecttotal.Scan(&total_detail); e {
		case sql.ErrNoRows:
		case nil:
		default:
			helpers.ErrorCheck(e)
		}
		// log.Println("TOTAL LISTPATTERNDETAIL : " + strconv.Itoa(total_detail))
		// log.Println("IDLISTPATTERN : " + idlistpattern)
		sql_selectdua := `SELECT 
			status_card , idpoin 
			FROM ` + configs.DB_tbl_trx_listpatterndetail + `  
			WHERE idlistpattern='` + idlistpattern + `' 
			ORDER BY idlistpatterndetail ASC   `

		row_dua, err_dua := con.QueryContext(ctx, sql_selectdua)
		helpers.ErrorCheck(err_dua)
		no := 0
		for row_dua.Next() {
			no = no + 1
			var (
				idpoin_db      int
				status_card_db string
			)

			err_dua = row_dua.Scan(&status_card_db, &idpoin_db)
			helpers.ErrorCheck(err_dua)
			// log.Println("MASUK COY")
			if no != total_detail {
				listpattern += _GetPattern(status_card_db, idpoin_db) + "|"
			} else {
				listpattern += _GetPattern(status_card_db, idpoin_db)
			}
		}
		defer row_dua.Close()
	}

	return listpattern, total_detail, idlistpattern
}

func _GetPattern(status string, idpoin int) string {
	con := db.CreateCon()
	ctx := context.Background()
	idpattern := ""
	sql_select := ""
	sql_select += "SELECT "
	sql_select += "idpattern "
	sql_select += "FROM " + configs.DB_tbl_trx_pattern + " "
	if status == "N" {
		sql_select += "WHERE status_pattern='" + status + "' "
	} else {
		sql_select += "WHERE status_pattern='" + status + "' "
		sql_select += "AND idpoin='" + strconv.Itoa(idpoin) + "' "
	}
	sql_select += "ORDER BY random()  LIMIT 1 "

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&idpattern); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return idpattern
}
func _GetTotalBet_Transaksi(table, idtransaksi string) int {
	con := db.CreateCon()
	ctx := context.Background()
	total_bet := 0
	sql_select := ""
	sql_select += "SELECT "
	sql_select += "(bet*roundbet_detail) AS total_bet "
	sql_select += "FROM " + table + " "
	sql_select += "WHERE idtransaksi='" + idtransaksi + "'  AND status_transaksidetail='LOSE' "
	sql_select += "ORDER BY idtransaksidetail DESC LIMIT 1 "

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&total_bet); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return total_bet
}
func _GetInfoPoint(codepoin string) string {
	con := db.CreateCon()
	ctx := context.Background()
	nmpoin := ""
	sql_select := ""
	sql_select += "SELECT "
	sql_select += "nmpoin "
	sql_select += "FROM " + configs.DB_tbl_mst_listpoint + " "
	sql_select += "WHERE codepoin='" + codepoin + "' "

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&nmpoin); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return nmpoin
}
