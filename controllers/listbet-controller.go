package controllers

import (
	"fmt"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/models"
)

const Fieldlistbet_home_redis = "LISTBET_BACKEND"
const Fieldlistbet_home_client_redis = "LISTBET_FRONTEND"

func Listbethome(c *fiber.Ctx) error {
	var obj entities.Model_lisbet
	var arraobj []entities.Model_lisbet
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldlistbet_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		lisbet_id, _ := jsonparser.GetInt(value, "lisbet_id")
		lisbet_minbet, _ := jsonparser.GetFloat(value, "lisbet_minbet")
		lisbet_create, _ := jsonparser.GetString(value, "lisbet_create")
		lisbet_update, _ := jsonparser.GetString(value, "lisbet_update")

		obj.Lisbet_id = int(lisbet_id)
		obj.Lisbet_minbet = float64(lisbet_minbet)
		obj.Lisbet_create = lisbet_create
		obj.Lisbet_update = lisbet_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_listbetHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldlistbet_home_redis, result, 60*time.Minute)
		fmt.Println("LISTBET MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("LISTBET CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func ListbetSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_listbetsave)
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
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	// admin, sData string, idrecord int, minbet float64
	result, err := models.Save_listbet(
		client_admin,
		client.Sdata, client.Lisbet_id, client.Lisbet_minbet)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_listbet()
	return c.JSON(result)
}
func _deleteredis_listbet() {
	val_master := helpers.DeleteRedis(Fieldlistbet_home_redis)
	fmt.Printf("Redis Delete BACKEND LISTBET : %d", val_master)

}
