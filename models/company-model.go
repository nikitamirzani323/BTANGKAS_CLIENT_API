package models

import (
	"context"
	"database/sql"
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

const database_company_local = configs.DB_tbl_mst_company
const database_companyadminrule_local = configs.DB_tbl_mst_company_adminrule
const database_companyadmin_local = configs.DB_tbl_mst_company_admin

func Fetch_companyHome() (helpers.Responsecompany, error) {
	var obj entities.Model_company
	var arraobj []entities.Model_company
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	var res helpers.Responsecompany
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idcompany ,    
			to_char(COALESCE(startjoincompany,now()), 'YYYY-MM-DD HH24:MI:SS'),
			to_char(COALESCE(endjoincompany,now()), 'YYYY-MM-DD HH24:MI:SS'),
			idcurr, nmcompany, nmowner, phoneowner, emailowner, companyurl,statuscompany,
			createcompany, to_char(COALESCE(createdatecompany,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updatecompany, to_char(COALESCE(updatedatecompany,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_company_local + `  
			ORDER BY createdatecompany DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcompany_db, startjoincompany_db, endjoincompany_db                                               string
			idcurr_db, nmcompany_db, nmowner_db, phoneowner_db, emailowner_db, companyurl_db, statuscompany_db string
			createcompany_db, createdatecompany_db, updatecompany_db, updatedatecompany_db                     string
		)

		err = row.Scan(&idcompany_db, &startjoincompany_db, &endjoincompany_db,
			&idcurr_db, &nmcompany_db, &nmowner_db, &phoneowner_db, &emailowner_db, &companyurl_db, &statuscompany_db,
			&createcompany_db, &createdatecompany_db, &updatecompany_db, &updatedatecompany_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createcompany_db != "" {
			create = createcompany_db + ", " + createdatecompany_db
		}
		if updatecompany_db != "" {
			update = updatecompany_db + ", " + updatedatecompany_db
		}
		if statuscompany_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}
		if startjoincompany_db == endjoincompany_db {
			endjoincompany_db = ""
		}
		obj.Company_id = idcompany_db
		obj.Company_startjoin = startjoincompany_db
		obj.Company_endjoin = endjoincompany_db
		obj.Company_idcurr = idcurr_db
		obj.Company_name = nmcompany_db
		obj.Company_nmowner = nmowner_db
		obj.Company_phoneowner = phoneowner_db
		obj.Company_emailowner = emailowner_db
		obj.Company_url = companyurl_db
		obj.Company_status = statuscompany_db
		obj.Company_status_css = status_css
		obj.Company_create = create
		obj.Company_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	sql_selectcurr := `SELECT 
			idcurr  
			FROM ` + configs.DB_tbl_mst_curr + ` 
			ORDER BY idcurr ASC    
	`
	rowcurr, errcurr := con.QueryContext(ctx, sql_selectcurr)
	helpers.ErrorCheck(errcurr)
	for rowcurr.Next() {
		var (
			idcurr_db string
		)

		errcurr = rowcurr.Scan(&idcurr_db)

		helpers.ErrorCheck(errcurr)

		objcurr.Curr_id = idcurr_db
		arraobjcurr = append(arraobjcurr, objcurr)
		msg = "Success"
	}
	defer rowcurr.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listcurr = arraobjcurr
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_companyadminruleHome() (helpers.Responsecompanyadminrule, error) {
	var obj entities.Model_companyadminrule
	var arraobj []entities.Model_companyadminrule
	var objcompany entities.Model_companyshare
	var arraobjcompany []entities.Model_companyshare
	var res helpers.Responsecompanyadminrule
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			companyrule_adminrule, idcompany, companyrule_name, companyrule_rule, 
			create_companyrule, to_char(COALESCE(createdate_companyrule,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			update_companyrule, to_char(COALESCE(updatedate_companyrule,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_companyadminrule_local + `  
			ORDER BY createdate_companyrule DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			companyrule_adminrule_db                                                                           int
			idcompany_db, companyrule_name_db, companyrule_rule_db                                             string
			create_companyrule_db, createdate_companyrule_db, update_companyrule_db, updatedate_companyrule_db string
		)

		err = row.Scan(&companyrule_adminrule_db, &idcompany_db, &companyrule_name_db, &companyrule_rule_db,
			&create_companyrule_db, &createdate_companyrule_db, &update_companyrule_db, &updatedate_companyrule_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if create_companyrule_db != "" {
			create = create_companyrule_db + ", " + createdate_companyrule_db
		}
		if update_companyrule_db != "" {
			update = update_companyrule_db + ", " + updatedate_companyrule_db
		}
		obj.Companyadminrule_id = companyrule_adminrule_db
		obj.Companyadminrule_idcompany = idcompany_db
		obj.Companyadminrule_nmrule = companyrule_name_db
		obj.Companyadminrule_rule = companyrule_rule_db
		obj.Companyadminrule_create = create
		obj.Companyadminrule_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	sql_selectcompany := `SELECT 
			idcompany, nmcompany  
			FROM ` + database_company_local + ` 
			WHERE statuscompany = 'Y' 
			ORDER BY idcompany ASC    
	`
	rowcompany, errcompany := con.QueryContext(ctx, sql_selectcompany)
	helpers.ErrorCheck(errcompany)
	for rowcompany.Next() {
		var (
			idcompany_db, nmcompany_db string
		)

		errcompany = rowcompany.Scan(&idcompany_db, &nmcompany_db)

		helpers.ErrorCheck(errcompany)

		objcompany.Company_id = idcompany_db
		objcompany.Company_name = nmcompany_db
		arraobjcompany = append(arraobjcompany, objcompany)
		msg = "Success"
	}
	defer rowcompany.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listcompany = arraobjcompany
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_companyadminHome() (helpers.Responsecompanyadmin, error) {
	var obj entities.Model_companyadmin
	var arraobj []entities.Model_companyadmin
	var objrule entities.Model_companyadminrule_share
	var arraobjrule []entities.Model_companyadminrule_share
	var objcompany entities.Model_companyshare
	var arraobjcompany []entities.Model_companyshare
	var res helpers.Responsecompanyadmin
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			company_idadmin, companyrule_adminrule, idcompany, tipeadmincompany, 
			company_username, company_ipaddress, 
			to_char(COALESCE(company_lastloginadmin,now()), 'YYYY-MM-DD HH24:MI:SS'),
			company_name, company_phone1, company_phone2, company_status,
			createadmin_company, to_char(COALESCE(createadmindate_company,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updateadmin_company, to_char(COALESCE(updateadmindate_company,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_companyadmin_local + `  
			ORDER BY createadmindate_company DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			companyrule_adminrule_db                                                                                                                       int
			company_idadmin_db, idcompany_db, tipeadmincompany_db                                                                                          string
			company_username_db, company_ipaddress_db, company_lastloginadmin_db, company_name_db, company_phone1_db, company_phone2_db, company_status_db string
			createadmin_company_db, createadmindate_company_db, updateadmin_company_db, updateadmindate_company_db                                         string
		)

		err = row.Scan(&company_idadmin_db, &companyrule_adminrule_db, &idcompany_db, &tipeadmincompany_db,
			&company_username_db, &company_ipaddress_db, &company_lastloginadmin_db, &company_name_db,
			&company_phone1_db, &company_phone2_db, &company_status_db,
			&createadmin_company_db, &createadmindate_company_db, &updateadmin_company_db, &updateadmindate_company_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createadmin_company_db != "" {
			create = createadmin_company_db + ", " + createadmindate_company_db
		}
		if updateadmin_company_db != "" {
			update = updateadmin_company_db + ", " + updateadmindate_company_db
		}
		if company_status_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}
		if company_lastloginadmin_db == createadmindate_company_db {
			company_lastloginadmin_db = ""
		}
		obj.Companyadmin_id = company_idadmin_db
		obj.Companyadmin_idrule = companyrule_adminrule_db
		obj.Companyadmin_idcompany = idcompany_db
		obj.Companyadmin_tipe = tipeadmincompany_db
		obj.Companyadmin_nmrule = _Get_infoadminrule(idcompany_db, companyrule_adminrule_db)
		obj.Companyadmin_username = company_username_db
		obj.Companyadmin_ipaddress = company_ipaddress_db
		obj.Companyadmin_lastlogin = company_lastloginadmin_db
		obj.Companyadmin_name = company_name_db
		obj.Companyadmin_phone1 = company_phone1_db
		obj.Companyadmin_phone2 = company_phone2_db
		obj.Companyadmin_status = company_status_db
		obj.Companyadmin_status_css = status_css
		obj.Companyadmin_create = create
		obj.Companyadmin_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	sql_selectcompany := `SELECT 
			idcompany, nmcompany  
			FROM ` + database_company_local + ` 
			WHERE statuscompany = 'Y' 
			ORDER BY idcompany ASC    
	`
	rowcompany, errcompany := con.QueryContext(ctx, sql_selectcompany)
	helpers.ErrorCheck(errcompany)
	for rowcompany.Next() {
		var (
			idcompany_db, nmcompany_db string
		)

		errcompany = rowcompany.Scan(&idcompany_db, &nmcompany_db)

		helpers.ErrorCheck(errcompany)

		objcompany.Company_id = idcompany_db
		objcompany.Company_name = nmcompany_db
		arraobjcompany = append(arraobjcompany, objcompany)
		msg = "Success"
	}

	sql_selectcompanyadminrule := `SELECT 
			A.companyrule_adminrule, A.idcompany, A.companyrule_name  
			FROM ` + database_companyadminrule_local + ` as A 
			JOIN ` + database_company_local + ` as B ON B.idcompany = A.idcompany  
			WHERE B.statuscompany = 'Y' 
			ORDER BY A.companyrule_adminrule ASC    
	`
	rowcompanyadminrule, errcompanyadminrule := con.QueryContext(ctx, sql_selectcompanyadminrule)
	helpers.ErrorCheck(errcompanyadminrule)
	for rowcompanyadminrule.Next() {
		var (
			companyrule_adminrule_db          int
			idcompany_db, companyrule_name_db string
		)

		errcompanyadminrule = rowcompanyadminrule.Scan(&companyrule_adminrule_db, &idcompany_db, &companyrule_name_db)

		helpers.ErrorCheck(errcompanyadminrule)

		objrule.Companyadminrule_id = companyrule_adminrule_db
		objrule.Companyadminrule_idcompany = idcompany_db
		objrule.Companyadminrule_nmrule = companyrule_name_db
		arraobjrule = append(arraobjrule, objrule)
		msg = "Success"
	}
	defer row.Close()
	defer rowcompany.Close()
	defer rowcompanyadminrule.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listcompany = arraobjcompany
	res.Listrule = arraobjrule
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_companyadminByCompany(idcompany string) (helpers.Response, error) {
	var obj entities.Model_companyadmin
	var arraobj []entities.Model_companyadmin
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			company_idadmin, companyrule_adminrule, idcompany, tipeadmincompany, 
			company_username, company_ipaddress, 
			to_char(COALESCE(company_lastloginadmin,now()), 'YYYY-MM-DD HH24:MI:SS'),
			company_name, company_phone1, company_phone2, company_status,
			createadmin_company, to_char(COALESCE(createadmindate_company,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updateadmin_company, to_char(COALESCE(updateadmindate_company,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_companyadmin_local + `  
			WHERE idcompany=$1 
			ORDER BY createadmindate_company DESC   `

	row, err := con.QueryContext(ctx, sql_select, idcompany)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			companyrule_adminrule_db                                                                                                                       int
			company_idadmin_db, idcompany_db, tipeadmincompany_db                                                                                          string
			company_username_db, company_ipaddress_db, company_lastloginadmin_db, company_name_db, company_phone1_db, company_phone2_db, company_status_db string
			createadmin_company_db, createadmindate_company_db, updateadmin_company_db, updateadmindate_company_db                                         string
		)

		err = row.Scan(&company_idadmin_db, &companyrule_adminrule_db, &idcompany_db, &tipeadmincompany_db,
			&company_username_db, &company_ipaddress_db, &company_lastloginadmin_db, &company_name_db,
			&company_phone1_db, &company_phone2_db, &company_status_db,
			&createadmin_company_db, &createadmindate_company_db, &updateadmin_company_db, &updateadmindate_company_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createadmin_company_db != "" {
			create = createadmin_company_db + ", " + createadmindate_company_db
		}
		if updateadmin_company_db != "" {
			update = updateadmin_company_db + ", " + updateadmindate_company_db
		}
		if company_status_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}
		if company_lastloginadmin_db == createadmindate_company_db {
			company_lastloginadmin_db = ""
		}
		obj.Companyadmin_id = company_idadmin_db
		obj.Companyadmin_idrule = companyrule_adminrule_db
		obj.Companyadmin_idcompany = idcompany_db
		obj.Companyadmin_tipe = tipeadmincompany_db
		obj.Companyadmin_nmrule = _Get_infoadminrule(idcompany_db, companyrule_adminrule_db)
		obj.Companyadmin_username = company_username_db
		obj.Companyadmin_ipaddress = company_ipaddress_db
		obj.Companyadmin_lastlogin = company_lastloginadmin_db
		obj.Companyadmin_name = company_name_db
		obj.Companyadmin_phone1 = company_phone1_db
		obj.Companyadmin_phone2 = company_phone2_db
		obj.Companyadmin_status = company_status_db
		obj.Companyadmin_status_css = status_css
		obj.Companyadmin_create = create
		obj.Companyadmin_update = update
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
func Fetch_companyListBet(idcompany string) (helpers.Response, error) {
	var obj entities.Model_company_listbet
	var arraobj []entities.Model_company_listbet
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	tbl_mst_listbet, _, _, _ := Get_mappingdatabase(idcompany)

	sql_select := `SELECT 
			idbet_listbet, minbet_listbet, 
			create_listbet, to_char(COALESCE(createdate_listbet,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			update_listbet, to_char(COALESCE(updatedate_listbet,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + tbl_mst_listbet + `  
			ORDER BY createdate_listbet DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idbet_listbet_db                                                                   int
			minbet_listbet_db                                                                  float64
			create_listbet_db, createdate_listbet_db, update_listbet_db, updatedate_listbet_db string
		)

		err = row.Scan(&idbet_listbet_db, &minbet_listbet_db,
			&create_listbet_db, &createdate_listbet_db, &update_listbet_db, &updatedate_listbet_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if create_listbet_db != "" {
			create = create_listbet_db + ", " + createdate_listbet_db
		}
		if update_listbet_db != "" {
			update = update_listbet_db + ", " + update_listbet_db
		}
		obj.Companylistbet_id = idbet_listbet_db
		obj.Companylistbet_minbet = minbet_listbet_db
		obj.Companylistbet_create = create
		obj.Companylistbet_update = update
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
func Fetch_companyConfPoint(idbet int, idcompany string) (helpers.Response, error) {
	var obj entities.Model_company_conf
	var arraobj []entities.Model_company_conf
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	_, tbl_mst_config, _, _ := Get_mappingdatabase(idcompany)

	sql_select := `SELECT 
			A.idconf_conf, A.idbet_listbet, A.idpoin, A.poin_conf,  
			A.create_conf, to_char(COALESCE(A.createdate_conf,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			A.update_conf, to_char(COALESCE(A.updatedate_conf,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + tbl_mst_config + ` as A   
			JOIN ` + configs.DB_tbl_mst_listpoint + ` as B ON B.idpoin = A.idpoin    
			WHERE A.idbet_listbet=$1 
			ORDER BY B.display_listpoint ASC   `

	row, err := con.QueryContext(ctx, sql_select, idbet)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idconf_conf_db, idbet_listbet_db, idpoin_db, poin_conf_db              int
			create_conf_db, createdate_conf_db, update_conf_db, updatedate_conf_db string
		)

		err = row.Scan(&idconf_conf_db, &idbet_listbet_db, &idpoin_db, &poin_conf_db,
			&create_conf_db, &createdate_conf_db, &update_conf_db, &updatedate_conf_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if create_conf_db != "" {
			create = create_conf_db + ", " + createdate_conf_db
		}
		if update_conf_db != "" {
			update = update_conf_db + ", " + updatedate_conf_db
		}
		obj.Companyconf_id = idconf_conf_db
		obj.Companyconf_idbet = idbet_listbet_db
		obj.Companyconf_nmpoin = _Get_infomasterpoint(idpoin_db)
		obj.Companyconf_poin = poin_conf_db
		obj.Companyconf_create = create
		obj.Companyconf_update = update
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
func Save_company(admin, idrecord, idcurr, nmcompany, nmowner, phoneowner, emailowner, url, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_company_local, "idcompany", idrecord)
		if !flag {
			sql_insert := `
				insert into
				` + database_company_local + ` (
					idcompany , startjoincompany, endjoincompany,  
					idcurr , nmcompany, nmowner, phoneowner, emailowner, companyurl, statuscompany,  
					createcompany, createdatecompany 
				) values (
					$1, $2, $3,    
					$4, $5, $6, $7, $8, $9, $10, 
					$11, $12   
				)
			`
			startjoin := tglnow.Format("YYYY-MM-DD HH:mm:ss")
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_company_local, "INSERT",
				idrecord, startjoin, startjoin,
				idcurr, nmcompany, nmowner, phoneowner, emailowner, url, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				msg = "Succes"
			} else {
				fmt.Println(msg_insert)
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		if status == "Y" {
			sql_update := `
				UPDATE 
				` + database_company_local + `  
				SET nmcompany=$1, nmowner=$2, phoneowner=$3, emailowner=$4, companyurl=$5, statuscompany=$6,   
				updatecompany=$7, updatedatecompany=$8      
				WHERE idcompany=$9     
			`

			flag_update, msg_update := Exec_SQL(sql_update, database_company_local, "UPDATE",
				nmcompany, nmowner, phoneowner, emailowner, url, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"
			} else {
				fmt.Println(msg_update)
			}
		} else {
			endjoin := tglnow.Format("YYYY-MM-DD HH:mm:ss")
			sql_update := `
				UPDATE 
				` + database_company_local + `  
				SET endjoincompany=$1, nmcompany=$2, nmowner=$3, phoneowner=$4, emailowner=$5, companyurl=$6, statuscompany=$7,   
				updatecompany=$8, updatedatecompany=$9       
				WHERE idcompany=$10     
			`

			flag_update, msg_update := Exec_SQL(sql_update, database_company_local, "UPDATE",
				endjoin, nmcompany, nmowner, phoneowner, emailowner, url, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"
			} else {
				fmt.Println(msg_update)
			}
		}

	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_companyadminrule(admin, idcompany, name, rule, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
			insert into
			` + database_companyadminrule_local + ` (
				companyrule_adminrule , idcompany, companyrule_name, companyrule_rule,   
				create_companyrule, createdate_companyrule 
			) values (
				$1, $2, $3, $4,    
				$5, $6   
			)
		`
		field_column := database_companyadminrule_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_companyadminrule_local, "INSERT",
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter), idcompany, name, rule,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_companyadminrule_local + `  
				SET companyrule_name=$1, companyrule_rule=$2, 
				update_companyrule=$3, updatedate_companyrule=$4 
				WHERE idcompany=$5 AND companyrule_adminrule=$6  
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_company_local, "UPDATE",
			name, rule,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idcompany, idrecord)

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
func Save_companyadmin(admin, idrecord, idcompany, username, password, name, phone1, phone2, status, sData string, idrule int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDBTwoField(database_companyadmin_local, "company_idadmin", idrecord, "idcompany", idcompany)
		if !flag {
			sql_insert := `
				insert into
				` + database_companyadmin_local + ` (
					company_idadmin , companyrule_adminrule, idcompany, tipeadmincompany,  
					company_username , company_password, company_lastloginadmin,  
					company_name , company_phone1, company_phone2,  company_status,
					createadmin_company, createadmindate_company 
				) values (
					$1, $2, $3, $4,    
					$5, $6, $7,     
					$8, $9, $10, $11,    
					$12, $13   
				)
			`
			startjoin := tglnow.Format("YYYY-MM-DD HH:mm:ss")
			hashpass := helpers.HashPasswordMD5(password)
			field_column := database_companyadmin_local + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_companyadmin_local, "INSERT",
				idcompany+tglnow.Format("YY")+strconv.Itoa(idrecord_counter), idrule, idcompany, "MASTER",
				username, hashpass, startjoin,
				name, phone1, phone2, status,
				admin, startjoin)

			if flag_insert {
				msg = "Succes"
			} else {
				fmt.Println(msg_insert)
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		if password != "" {
			hashpass := helpers.HashPasswordMD5(password)
			sql_update := `
				UPDATE 
				` + database_companyadmin_local + `  
				SET company_password=$1, companyrule_adminrule=$2, company_name=$3, 
				company_phone1=$4, company_phone2=$5, company_status=$6, 
				updateadmin_company=$7, updateadmindate_company=$8  
				WHERE idcompany=$9 AND company_idadmin=$10   
			`

			flag_update, msg_update := Exec_SQL(sql_update, database_companyadmin_local, "UPDATE",
				hashpass, idrule, name, phone1, phone2, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idcompany, idrecord)

			if flag_update {
				msg = "Succes"
			} else {
				fmt.Println(msg_update)
			}
		} else {
			sql_update := `
				UPDATE 
				` + database_companyadmin_local + `  
				SET companyrule_adminrule=$1, company_name=$2, 
				company_phone1=$3, company_phone2=$4, company_status=$5, 
				updateadmin_company=$6, updateadmindate_company=$7  
				WHERE idcompany=$8 AND company_idadmin=$9  
			`

			flag_update, msg_update := Exec_SQL(sql_update, database_companyadmin_local, "UPDATE",
				idrule, name, phone1, phone2, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idcompany, idrecord)

			if flag_update {
				msg = "Succes"
			} else {
				fmt.Println(msg_update)
			}
		}

	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_companyListBet(admin, idcompany, sData string, idrecord, minbet int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false
	tbl_mst_listbet, _, _, _ := Get_mappingdatabase(idcompany)

	if sData == "New" {
		flag = CheckDBTwoField(tbl_mst_listbet, "idcompany", idcompany, "minbet_listbet", strconv.Itoa(minbet))
		if !flag {
			sql_insert := `
			insert into
			` + tbl_mst_listbet + ` (
				idbet_listbet , idcompany, minbet_listbet,    
				create_listbet, createdate_listbet 
			) values (
				$1, $2, $3,    
				$4, $5    
			)
		`
			field_column := idcompany + "_" + tbl_mst_listbet + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_mst_listbet, "INSERT",
				tglnow.Format("YY")+strconv.Itoa(idrecord_counter), idcompany, minbet,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				msg = "Succes"
			} else {
				fmt.Println(msg_insert)
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		sql_update := `
			UPDATE 
			` + tbl_mst_listbet + `  
			SET minbet_listbet=$1,  
			updatecompany=$2, updatedatecompany=$3       
			WHERE idbet_listbet=$4 AND idcompany=$5      
		`

		flag_update, msg_update := Exec_SQL(sql_update, tbl_mst_listbet, "UPDATE",
			minbet,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord, idcompany)

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
func Save_companyConfPoint(admin, idcompany, sData string, idrecord, idbet, point int) (helpers.Response, error) {
	var res helpers.Response
	con := db.CreateCon()
	ctx := context.Background()
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	_, tbl_mst_config, _, _ := Get_mappingdatabase(idcompany)

	if sData == "New" {
		sql_select := `SELECT 
				idpoin , poin
				FROM ` + configs.DB_tbl_mst_listpoint + `  
				ORDER BY display_listpoint ASC   `

		row, err := con.QueryContext(ctx, sql_select)
		helpers.ErrorCheck(err)
		for row.Next() {
			var (
				idpoin_db, poin_db int
			)

			err = row.Scan(&idpoin_db, &poin_db)
			helpers.ErrorCheck(err)
			flag = CheckDBThreeField(tbl_mst_config, "idbet_listbet", strconv.Itoa(idbet), "idcompany", idcompany, "idpoin", strconv.Itoa(idpoin_db))
			if !flag {
				sql_insert := `
					insert into
					` + tbl_mst_config + ` (
						idconf_conf , idbet_listbet, idcompany, idpoin, poin_conf,     
						create_conf, createdate_conf 
					) values (
						$1, $2, $3, $4, $5,     
						$6, $7     
					)
				`
				field_column := idcompany + "_" + tbl_mst_config + tglnow.Format("YYYY")
				idrecord_counter := Get_counter(field_column)
				flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_mst_config, "INSERT",
					tglnow.Format("YY")+strconv.Itoa(idrecord_counter), idbet, idcompany, idpoin_db, poin_db,
					admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

				if flag_insert {
					msg = "Succes"
				} else {
					fmt.Println(msg_insert)
				}
			}
		}

	} else {
		sql_update := `
			UPDATE 
			` + tbl_mst_config + `  
			SET poin_conf=$1,  
			update_conf=$2, updatedate_conf=$3       
			WHERE idconf_conf=$4 AND idcompany=$5      
		`

		flag_update, msg_update := Exec_SQL(sql_update, tbl_mst_config, "UPDATE",
			point,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord, idcompany)

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
func _Get_infoadminrule(idcompany string, idrecord int) string {
	con := db.CreateCon()
	ctx := context.Background()
	companyrule_name := ""
	sql_select := `SELECT
			companyrule_name    
			FROM ` + database_companyadminrule_local + `  
			WHERE companyrule_adminrule=` + strconv.Itoa(idrecord) + ` AND idcompany='` + idcompany + `'    
		`

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&companyrule_name); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return companyrule_name
}
func _Get_infomasterpoint(idpoin int) string {
	con := db.CreateCon()
	ctx := context.Background()
	nmpoin := ""
	sql_select := `SELECT
			nmpoin    
			FROM ` + configs.DB_tbl_mst_listpoint + `  
			WHERE idpoin='` + strconv.Itoa(idpoin) + `'     
		`

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&nmpoin); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return nmpoin
}
