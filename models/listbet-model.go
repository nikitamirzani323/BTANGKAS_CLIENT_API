package models

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/db"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/helpers"
)

func Fetch_listbetHome(idcompany string) (helpers.Response, error) {
	var obj entities.Model_lisbet
	var arraobj []entities.Model_lisbet
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()
	tbl_mst_listbet, tbl_mst_config, _, _ := Get_mappingdatabase(idcompany)
	sql_select := `SELECT 
			idbet_listbet , minbet_listbet
			FROM ` + tbl_mst_listbet + `  
			WHERE idcompany=$1 
			ORDER BY minbet_listbet ASC   `
	row, err := con.QueryContext(ctx, sql_select, idcompany)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idbet_db          int
			minbet_listbet_db float64
		)

		err = row.Scan(&idbet_db, &minbet_listbet_db)
		helpers.ErrorCheck(err)

		var obj_listpoin entities.Model_lispoin
		var arraobj_listpoin []entities.Model_lispoin
		sql_select_listpoin := `SELECT 
			B.codepoin , B.nmpoin, A.poin_conf
			FROM ` + tbl_mst_config + ` as A  
			JOIN ` + configs.DB_tbl_mst_listpoint + ` as B ON B.idpoin = A.idpoin  
			WHERE A.idbet_listbet=$1 
			ORDER BY B.display_listpoint ASC   `

		row_listpoin, err_listpoin := con.QueryContext(ctx, sql_select_listpoin, idbet_db)
		helpers.ErrorCheck(err_listpoin)
		for row_listpoin.Next() {
			var (
				codepoin_db, nmpoin_db string
				poin_conf_db           int
			)

			err_listpoin = row_listpoin.Scan(&codepoin_db, &nmpoin_db, &poin_conf_db)
			helpers.ErrorCheck(err_listpoin)

			obj_listpoin.Lispoin_id = codepoin_db
			obj_listpoin.Lispoin_nmpoin = nmpoin_db
			obj_listpoin.Lispoin_poin = poin_conf_db
			arraobj_listpoin = append(arraobj_listpoin, obj_listpoin)
		}

		obj.Lisbet_id = idbet_db
		obj.Lisbet_minbet = minbet_listbet_db
		obj.Lisbet_conf = arraobj_listpoin
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
