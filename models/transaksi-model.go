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

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/db"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/helpers"
	"github.com/nleeper/goment"
)

var a []int

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

	pattern := ""
	field_redis := "PATTERN_" + idcompany + "_" + username
	if round_game_all == 0 {
		pattern = _GenerateCardRandom()
		helpers.SetRedis(field_redis, pattern, 5*time.Minute)
	} else {
		resultredis, _ := helpers.GetRedis(field_redis)
		// jsonredis := []byte(resultredis)
		// record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
		log.Println("Data redis " + resultredis)
		pattern = _GenerateCardRandom()
	}
	resultcard := strings.Split(pattern, "|")

	log.Println("Generate :" + pattern)
	flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_trx_transaksi, "INSERT",
		idrecrodparent_value, idcompany, date_transaksi,
		username, 0, resultcard[round_game_all],
		"SYSTEM", date_transaksi)

	if flag_insert {
		msg = "Succes"
		log.Printf("round %d", round_bet)
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

func _GenerateCard() string {
	// var a [7]int
	min := 0
	max := 54
	result := ""
	// for i := 0; i < 7; i++ {
	// 	var n = rand.Intn(max-min) + min
	// 	a[i] = n
	// 	if i == 6 {
	// 		result += strconv.Itoa(n)
	// 	} else {
	// 		result += strconv.Itoa(n) + ","
	// 	}

	// }
	var i = 0
	for {
		var n = rand.Intn(max-min) + min

		if !search_array(n) {
			a = append(a, n)

			if i == 6 {
				result += strconv.Itoa(n)
			} else {
				result += strconv.Itoa(n) + ","
			}
			i++
		}
		if i == 7 {
			break
		}

	}
	return result
}
func search_array(key int) bool {
	for _, element := range a {
		if element == key { // check the condition if its true return index
			return true
		}
	}
	return false
}
func _GenerateCardDB() string {
	con := db.CreateCon()
	ctx := context.Background()
	var idpattern string
	sql_select := `
			SELECT
			idpattern     
			FROM ` + configs.DB_tbl_trx_pattern + ` 
			ORDER BY random()
			LIMIT 1
		`

	fmt.Println(sql_select)
	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&idpattern); e {
	case sql.ErrNoRows:

	case nil:

	default:

	}

	return idpattern
}
func _GenerateCardRandom() string {

	pattern := "37-18-6-0-3-21-10|40-47-52-5-33-21-0|47-13-19-12-24-10-28|2-0-14-27-50-22-19|41-7-49-47-32-30-46|"
	pattern += "41-15-29-18-48-12-40|41-15-29-18-48-12-40|36-18-27-25-50-48-26|34-42-24-51-53-4-16|47-11-20-43-32-28-6"

	return pattern
}
