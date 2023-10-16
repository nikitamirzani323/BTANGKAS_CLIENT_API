package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/models"
)

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

	//idcompany, username string, round_game_all,round_bet, bet, c_before, c_after, win, idpoin int
	result, err := models.Save_transaksi(client.Transaksi_company, client.Transaksi_username, client.Transaksi_status, client.Transaksi_resultcardwin,
		client.Transaksi_roundgameall, client.Transaksi_roundbet, client.Transaksi_bet, client.Transaksi_cbefore, client.Transaksi_cafter,
		client.Transaksi_win, client.Transaksi_idpoin)

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

	//idtransaksi, resulcard_win string, round_bet, bet, c_before, c_after, win, idpoin int
	result, err := models.Save_transaksidetail(client.Transaksidetail_company,
		client.Transaksidetail_idtransaksi, client.Transaksidetail_resultcardwin, client.Transaksidetail_status,
		client.Transaksidetail_roundbet, client.Transaksidetail_bet, client.Transaksidetail_cbefore, client.Transaksidetail_cafter,
		client.Transaksidetail_win, client.Transaksidetail_idpoin)

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
