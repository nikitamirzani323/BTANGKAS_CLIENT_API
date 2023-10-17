package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/db"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/helpers"
	"github.com/nleeper/goment"
)

type Card_Strc struct {
	TypePattern string `json:"typepattern"`
	Pattern     string `json:"pattern"`
	Total       int    `json:"total"`
}

func Save_transaksi(idcompany, username, status, resultcardwin string, round_game_all, round_bet, bet, c_before, c_after, win, idpoin int) (helpers.Responsetransaksi, error) {
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

	// pattern_new := _GenerateCardRandomDB()
	// log.Println("Pattern Baru : " + pattern_new)

	idlistpattern := ""
	pattern := ""
	total_card := 0
	field_redis := "PATTERN_" + idcompany + "_" + username
	var card_res Card_Strc

	if round_game_all < 1 {
		fmt.Println("Database Card Pattern")
		pattern, total_card, idlistpattern = _GenerateCardRandomDB()
		card_res.TypePattern = idlistpattern
		card_res.Pattern = pattern
		card_res.Total = total_card
		helpers.SetRedis(field_redis, card_res, 10*time.Minute)
	} else {
		resultredis, flag := helpers.GetRedis(field_redis)
		jsonredis := []byte(resultredis)
		pattern_redis, _ := jsonparser.GetString(jsonredis, "pattern")
		typepattern_redis, _ := jsonparser.GetString(jsonredis, "typepattern")
		total_redis, _ := jsonparser.GetInt(jsonredis, "total")
		if int(total_redis-1) == round_game_all {
			val_pattern := helpers.DeleteRedis(field_redis)
			fmt.Printf("Redis Delete Card Pattern : %d\n", val_pattern)
			flag = false
		}
		if !flag {
			fmt.Println("Database Card Pattern")
			pattern, total_card, idlistpattern = _GenerateCardRandomDB()
			card_res.TypePattern = idlistpattern
			card_res.Pattern = pattern
			card_res.Total = total_card
			helpers.SetRedis(field_redis, card_res, 5*time.Minute)
		} else {
			fmt.Println("Cache Card Pattern")
			pattern = pattern_redis
			idlistpattern = typepattern_redis
			total_card = int(total_redis)
		}

	}
	log.Println("Total Card : " + strconv.Itoa(total_card))
	log.Println("Total Game : " + strconv.Itoa(round_game_all))
	log.Println("Type Pattern : " + idlistpattern)
	resultcard := strings.Split(pattern, "|")

	flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_trx_transaksi, "INSERT",
		idrecrodparent_value, idcompany, date_transaksi,
		username, 0, resultcard[round_game_all],
		"SYSTEM", date_transaksi)

	if flag_insert {
		msg = "Succes"
		if round_bet == 1 || round_bet == 4 {
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
			flag_insertdetail, msg_insertdetail := Exec_SQL(sql_insertdetail, tbl_trx_transaksidetail, "INSERT",
				idrecroddetail_value, idrecrodparent_value, round_bet,
				bet, c_before, c_after,
				win, idpoin, resultcardwin, status,
				"SYSTEM", tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insertdetail {
				msg_insertdetail = "Succes"
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
func Save_transaksidetail(idcompany, idtransaksi, resulcard_win, status string, round_bet, bet, c_before, c_after, win, idpoin int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	_, _, _, tbl_trx_transaksidetail := Get_mappingdatabase(idcompany)

	sql_insert := `
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

	field_column := tbl_trx_transaksidetail + tglnow.Format("YYYY") + tglnow.Format("MM")
	idrecord_counter := Get_counter(field_column)
	idrecrod_value := tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
	flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_trx_transaksidetail, "INSERT",
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

func _GenerateCardRandom() (string, int) {
	pattern := ""
	total_card := 0
	listpattern := [...]string{
		"37-18-6-0-3-21-10|40-47-52-5-33-21-0|47-13-19-12-24-10-28|2-0-14-27-50-22-19|41-7-49-47-32-30-46|41-15-29-18-48-12-40|41-15-29-18-48-12-40|36-18-27-25-50-48-26|34-42-24-51-53-4-16|47-11-20-43-32-28-6",
		"16-33-31-52-22-43-18|40-47-52-5-33-21-0|47-13-19-12-24-10-28|2-0-14-27-50-22-19|41-7-49-47-32-30-46|41-15-29-18-48-12-40|41-15-29-18-48-12-40|36-18-27-25-50-48-26|34-42-24-51-53-4-16|47-11-20-43-32-28-6",
		"14-39-25-45-32-2-35|37-1-48-42-4-49-29|47-39-45-30-36-7-38|2-0-14-27-50-22-19|41-7-49-47-32-30-46|41-15-29-18-48-12-40|41-15-29-18-48-12-40|36-18-27-25-50-48-26|34-42-24-51-53-4-16|47-11-20-43-32-28-6"}

	min := 0
	max := len(listpattern)
	var n = rand.Intn(max-min) + min

	pattern = listpattern[n]
	total_card = len(strings.Split(pattern, "|"))
	fmt.Println()

	return pattern, total_card
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
