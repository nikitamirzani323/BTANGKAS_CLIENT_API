package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/models"
)

const listbet_client_redis = "CLIENT_LISTBET"
const invoice_super_redis = "COMPANYINVOICE_BACKEND"
const invoice_agen_redis = "AGEN_TRANSAKSI"
const invoice_client_redis = "CLIENT_LISTINVOICE"

type c_tai struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
	Time    string      `json:"time"`
}

func CheckToken(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.CheckToken)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	// result, ruleadmin, err := models.Login_Model(client.Username, client.Password, client.Ipaddress)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	result := false
	if client.Token == "qC5YmBvXzabGp34jJlKvnC6wCrr3pLCwBzsLoSzl4k=" {
		result = true
	}

	if !result {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": "Data Not Found",
			})

	} else {
		// dataclient := client.Username + "==" + ruleadmin
		// dataclient_encr, keymap := helpers.Encryption(dataclient)
		// dataclient_encr_final := dataclient_encr + "|" + strconv.Itoa(keymap)
		// t, err := helpers.GenerateNewAccessToken(dataclient_encr_final)
		// if err != nil {
		// 	return c.SendStatus(fiber.StatusInternalServerError)
		// }
		// var obj_clistbet c_listbet
		// var arraobj_clistbet []c_listbet
		var obj_record c_tai

		var obj entities.Model_lisbet
		var arraobj []entities.Model_lisbet
		resultredis, flag := helpers.GetRedis(listbet_client_redis + "_" + strings.ToLower("ajuna"))
		jsonredis := []byte(resultredis)
		status_RD, _ := jsonparser.GetInt(jsonredis, "status")
		message_RD, _ := jsonparser.GetString(jsonredis, "message")
		record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
		jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			lisbet_id, _ := jsonparser.GetInt(value, "lisbet_id")
			lisbet_minbet, _ := jsonparser.GetFloat(value, "lisbet_minbet")

			var obj_listpoin entities.Model_lispoin
			var arraobj_listpoin []entities.Model_lispoin
			record_lisbet_config_RD, _, _, _ := jsonparser.Get(value, "lisbet_config")
			jsonparser.ArrayEach(record_lisbet_config_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				lispoin_id, _ := jsonparser.GetString(value, "lispoin_id")
				lispoin_nmpoin, _ := jsonparser.GetString(value, "lispoin_nmpoin")
				lispoin_poin, _ := jsonparser.GetInt(value, "lispoin_poin")

				obj_listpoin.Lispoin_id = lispoin_id
				obj_listpoin.Lispoin_nmpoin = lispoin_nmpoin
				obj_listpoin.Lispoin_poin = int(lispoin_poin)
				arraobj_listpoin = append(arraobj_listpoin, obj_listpoin)
			})

			obj.Lisbet_id = int(lisbet_id)
			obj.Lisbet_minbet = float64(lisbet_minbet)
			obj.Lisbet_conf = arraobj_listpoin
			arraobj = append(arraobj, obj)
		})
		obj_record.Status = int(status_RD)
		obj_record.Message = message_RD
		obj_record.Record = arraobj
		if !flag {
			fmt.Println("LISTBET MYSQL")
			listbet, _ := models.Fetch_listbetHome("AJUNA")
			helpers.SetRedis(listbet_client_redis+"_"+strings.ToLower("ajuna"), listbet, 1440*time.Minute)
			return c.JSON(fiber.Map{
				"status":          fiber.StatusOK,
				"client_company":  "ajuna",
				"client_name":     "developer",
				"client_username": "developer212",
				"client_listbet":  listbet,
				"client_credit":   100000,
			})
		} else {
			fmt.Println("LISTBET CACHE")
			return c.JSON(fiber.Map{
				"status":          fiber.StatusOK,
				"client_company":  "ajuna",
				"client_name":     "developer",
				"client_username": "developer212",
				"client_listbet":  obj_record,
				"client_credit":   100000,
			})
		}

	}
}

func ListInvoice(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_invoice)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	var obj entities.Model_invoice
	var arraobj []entities.Model_invoice
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(invoice_client_redis + "_" + strings.ToLower(client.Invoice_company) + "_" + strings.ToLower(client.Invoice_username))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		invoice_id, _ := jsonparser.GetString(value, "invoice_id")
		invoice_date, _ := jsonparser.GetString(value, "invoice_date")
		invoice_round, _ := jsonparser.GetInt(value, "invoice_round")
		invoice_totalbet, _ := jsonparser.GetInt(value, "invoice_totalbet")
		invoice_totalwin, _ := jsonparser.GetInt(value, "invoice_totalwin")
		invoice_nmpoin, _ := jsonparser.GetString(value, "invoice_nmpoin")
		invoice_status, _ := jsonparser.GetString(value, "invoice_status")
		invoice_status_css, _ := jsonparser.GetString(value, "invoice_status_css")
		invoice_card_result, _ := jsonparser.GetString(value, "invoice_card_result")
		invoice_card_win, _ := jsonparser.GetString(value, "invoice_card_win")

		obj.Invoice_id = invoice_id
		obj.Invoice_date = invoice_date
		obj.Invoice_round = int(invoice_round)
		obj.Invoice_totalbet = int(invoice_totalbet)
		obj.Invoice_totalwin = int(invoice_totalwin)
		obj.Invoice_nmpoin = invoice_nmpoin
		obj.Invoice_status = invoice_status
		obj.Invoice_status_css = invoice_status_css
		obj.Invoice_card_result = invoice_card_result
		obj.Invoice_card_win = invoice_card_win
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_invoice(client.Invoice_company, client.Invoice_username)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(invoice_client_redis+"_"+strings.ToLower(client.Invoice_company)+"_"+strings.ToLower(client.Invoice_username), result, 60*time.Minute)
		fmt.Printf("INVOICE MYSQL %s-%s\n", client.Invoice_company, client.Invoice_username)
		return c.JSON(result)
	} else {
		fmt.Printf("INVOICE CACHE %s-%s\n", client.Invoice_company, client.Invoice_username)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func TransaksiSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_transaksisave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	// user := c.Locals("jwt").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// name := claims["name"].(string)
	// temp_decp := helpers.Decryption(name)
	// client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	//idcompany, username, status, resultcardwin, codepoin string, round_game_all, round_bet, bet, c_before, c_after, win
	result, err := models.Save_transaksi(client.Transaksi_company, client.Transaksi_username, client.Transaksi_status, client.Transaksi_resultcardwin, client.Transaksi_codepoin,
		client.Transaksi_roundgameall, client.Transaksi_roundbet, client.Transaksi_bet, client.Transaksi_cbefore, client.Transaksi_cafter,
		client.Transaksi_win)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":      fiber.StatusBadRequest,
			"message":     err.Error(),
			"idtransaksi": "",
			"card_game":   "",
			"card_length": 0,
			"time":        "",
		})
	}
	_deleteredis_game(client.Transaksi_company, client.Transaksi_username)
	return c.JSON(result)
}
func TransaksidetailSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_transaksidetailsave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	// user := c.Locals("jwt").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// name := claims["name"].(string)
	// temp_decp := helpers.Decryption(name)
	// client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	//idcompany, idtransaksi, resulcard_win, status, codepoin string, round_bet, bet, c_before, c_after, win int
	result, err := models.Save_transaksidetail(client.Transaksidetail_company,
		client.Transaksidetail_idtransaksi, client.Transaksidetail_resultcardwin, client.Transaksidetail_status, client.Transaksidetail_codepoin,
		client.Transaksidetail_roundbet, client.Transaksidetail_bet, client.Transaksidetail_cbefore, client.Transaksidetail_cafter,
		client.Transaksidetail_win)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func _deleteredis_game(company, username string) {
	val_invoice := helpers.DeleteRedis(invoice_client_redis + "_" + strings.ToLower(company) + "_" + strings.ToLower(username))
	fmt.Printf("Redis Delete INVOICE : %d - %s %s\n", val_invoice, company, username)

	val_invoice_super := helpers.DeleteRedis(invoice_super_redis + "_" + strings.ToLower(company))
	fmt.Printf("Redis Delete INVOICE SUPER : %d - %s %s\n", val_invoice_super, company, username)

	val_invoice_agen := helpers.DeleteRedis(invoice_agen_redis + "_" + strings.ToLower(company))
	fmt.Printf("Redis Delete INVOICE AGEN : %d - %s \n", val_invoice_agen, company)
}
