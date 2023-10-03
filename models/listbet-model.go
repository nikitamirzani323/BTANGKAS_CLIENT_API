package models

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
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

	tbl_mst_listbet, _, _, _ := Get_mappingdatabase(idcompany)

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

		obj.Lisbet_id = idbet_db
		obj.Lisbet_minbet = minbet_listbet_db
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
