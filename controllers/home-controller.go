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

const invoice_super_redis = "COMPANYINVOICE_BACKEND"
const invoice_home_redis = "LISTINVOICE"

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
		listbet, _ := models.Fetch_listbetHome("AJUNA")
		return c.JSON(fiber.Map{
			"status":           fiber.StatusOK,
			"client_idcompany": "ajuna",
			"client_name":      "developer",
			"client_username":  "developer212",
			"client_credit":    100000,
			"client_listbet":   listbet,
		})

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
	resultredis, flag := helpers.GetRedis(invoice_home_redis + "_" + strings.ToLower(client.Invoice_company) + "_" + strings.ToLower(client.Invoice_username))
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
		helpers.SetRedis(invoice_home_redis+"_"+strings.ToLower(client.Invoice_company)+"_"+strings.ToLower(client.Invoice_username), result, 60*time.Minute)
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
	val_invoice := helpers.DeleteRedis(invoice_home_redis + "_" + strings.ToLower(company) + "_" + strings.ToLower(username))
	fmt.Printf("Redis Delete INVOICE : %d - %s %s\n", val_invoice, company, username)

	val_invoice_super := helpers.DeleteRedis(invoice_super_redis + "_" + strings.ToLower(company))
	fmt.Printf("Redis Delete INVOICE SUPER : %d - %s %s\n", val_invoice_super, company, username)
}
