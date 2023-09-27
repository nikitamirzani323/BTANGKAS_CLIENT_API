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

const Fieldlistpoint_home_redis = "LISTPOINT_BACKEND"
const Fieldlistpoint_home_client_redis = "LISTPOINT_FRONTEND"

func Listpointhome(c *fiber.Ctx) error {
	var obj entities.Model_listpoint
	var arraobj []entities.Model_listpoint
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldlistpoint_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		lispoint_id, _ := jsonparser.GetInt(value, "lispoint_id")
		lispoint_code, _ := jsonparser.GetString(value, "lispoint_code")
		lispoint_name, _ := jsonparser.GetString(value, "lispoint_name")
		lispoint_point, _ := jsonparser.GetInt(value, "lispoint_point")
		lispoint_display, _ := jsonparser.GetInt(value, "lispoint_display")
		lispoint_create, _ := jsonparser.GetString(value, "lispoint_create")
		lispoint_update, _ := jsonparser.GetString(value, "lispoint_update")

		obj.Lispoint_id = int(lispoint_id)
		obj.Lispoint_code = lispoint_code
		obj.Lispoint_name = lispoint_name
		obj.Lispoint_point = int(lispoint_point)
		obj.Lispoint_display = int(lispoint_display)
		obj.Lispoint_create = lispoint_create
		obj.Lispoint_update = lispoint_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_listpointHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldlistpoint_home_redis, result, 60*time.Minute)
		fmt.Println("LISTPOINT MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("LISTPOINT CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func ListpointSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_listpointsave)
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

	// admin, code, name, sData string, idrecord int
	result, err := models.Save_listpoint(
		client_admin,
		client.Lispoint_code, client.Lispoint_name, client.Sdata, client.Lispoint_id, client.Lispoint_point, client.Lispoint_display)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_listpoint()
	return c.JSON(result)
}
func _deleteredis_listpoint() {
	val_master := helpers.DeleteRedis(Fieldlistpoint_home_redis)
	fmt.Printf("Redis Delete BACKEND LISTPOINT : %d", val_master)

}
