package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/helpers"
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
				"message": "Username or Password Not Found",
			})

	} else {
		// dataclient := client.Username + "==" + ruleadmin
		// dataclient_encr, keymap := helpers.Encryption(dataclient)
		// dataclient_encr_final := dataclient_encr + "|" + strconv.Itoa(keymap)
		// t, err := helpers.GenerateNewAccessToken(dataclient_encr_final)
		// if err != nil {
		// 	return c.SendStatus(fiber.StatusInternalServerError)
		// }

		return c.JSON(fiber.Map{
			"status":           fiber.StatusOK,
			"client_idcompany": "ajuna",
			"client_name":      "developer",
			"client_credit":    10000,
		})

	}
}
