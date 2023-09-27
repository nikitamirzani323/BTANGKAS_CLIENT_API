package models

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/db"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/helpers"
	"github.com/nleeper/goment"
)

const database_listpoint_local = configs.DB_tbl_mst_listpoint

func Fetch_listpointHome() (helpers.Response, error) {
	var obj entities.Model_listpoint
	var arraobj []entities.Model_listpoint
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idpoin , codepoin, nmpoin, poin,  display_listpoint,  
			create_listpoint, to_char(COALESCE(createdate_listpoint,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			update_listpoint, to_char(COALESCE(updatedate_listpoint,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_listpoint_local + `  
			ORDER BY display_listpoint ASC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idpoin_db, poin_db, display_listpoint_db                                                   int
			codepoin_db, nmpoin_db                                                                     string
			create_listpoint_db, createdate_listpoint_db, update_listpoint_db, updatedate_listpoint_db string
		)

		err = row.Scan(&idpoin_db, &codepoin_db, &nmpoin_db, &poin_db, &display_listpoint_db,
			&create_listpoint_db, &createdate_listpoint_db, &update_listpoint_db, &updatedate_listpoint_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if create_listpoint_db != "" {
			create = create_listpoint_db + ", " + createdate_listpoint_db
		}
		if update_listpoint_db != "" {
			update = update_listpoint_db + ", " + updatedate_listpoint_db
		}

		obj.Lispoint_id = idpoin_db
		obj.Lispoint_code = codepoin_db
		obj.Lispoint_name = nmpoin_db
		obj.Lispoint_point = poin_db
		obj.Lispoint_display = display_listpoint_db
		obj.Lispoint_create = create
		obj.Lispoint_update = update
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
func Save_listpoint(admin, code, name, sData string, idrecord, poin, display int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
			insert into
			` + database_listpoint_local + ` (
				idpoin , codepoin, nmpoin, poin, display_listpoint,  
				create_listpoint, createdate_listpoint 
			) values (
				$1, $2, $3, $4, $5,    
				$6, $7  
			)
			`

		field_column := database_listpoint_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_listpoint_local, "INSERT",
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter), code, name, poin, display,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_listpoint_local + `  
				SET codepoin=$1, nmpoin=$2, poin=$3, display_listpoint=$4,   
				update_listpoint=$5, updatedate_listpoint=$6      
				WHERE idpoin=$7     
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_listpoint_local, "UPDATE",
			code, name, poin, display,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

		if flag_update {
			msg = "Succes"
		} else {
			fmt.Println(msg_update)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
