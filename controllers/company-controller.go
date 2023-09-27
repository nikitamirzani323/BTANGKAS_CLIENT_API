package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_CLIENT_API/models"
)

const Fieldcompany_home_redis = "COMPANY_BACKEND"
const Fieldcompanylistbet_home_redis = "COMPANYLISTBET_BACKEND"
const Fieldcompanyconf_home_redis = "COMPANYCONF_BACKEND"
const Fieldcompanyadminrule_home_redis = "COMPANYADMINRULE_BACKEND"
const Fieldcompanyadmin_home_redis = "COMPANYADMIN_BACKEND"
const Fieldcompany_home_client_redis = "COMPANY_FRONTEND"

func Companyhome(c *fiber.Ctx) error {
	var obj entities.Model_company
	var arraobj []entities.Model_company
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompany_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listcurr_RD, _, _, _ := jsonparser.Get(jsonredis, "listcurr")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_id, _ := jsonparser.GetString(value, "company_id")
		company_startjoin, _ := jsonparser.GetString(value, "company_startjoin")
		company_endjoin, _ := jsonparser.GetString(value, "company_endjoin")
		company_name, _ := jsonparser.GetString(value, "company_name")
		company_idcurr, _ := jsonparser.GetString(value, "company_idcurr")
		company_nmowner, _ := jsonparser.GetString(value, "company_nmowner")
		company_phoneowner, _ := jsonparser.GetString(value, "company_phoneowner")
		company_emailowner, _ := jsonparser.GetString(value, "company_emailowner")
		company_url, _ := jsonparser.GetString(value, "company_url")
		company_status, _ := jsonparser.GetString(value, "company_status")
		company_status_css, _ := jsonparser.GetString(value, "company_status_css")
		company_create, _ := jsonparser.GetString(value, "company_create")
		company_update, _ := jsonparser.GetString(value, "company_update")

		obj.Company_id = company_id
		obj.Company_startjoin = company_startjoin
		obj.Company_endjoin = company_endjoin
		obj.Company_name = company_name
		obj.Company_idcurr = company_idcurr
		obj.Company_nmowner = company_nmowner
		obj.Company_phoneowner = company_phoneowner
		obj.Company_emailowner = company_emailowner
		obj.Company_url = company_url
		obj.Company_status = company_status
		obj.Company_status_css = company_status_css
		obj.Company_create = company_create
		obj.Company_update = company_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listcurr_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		curr_id, _ := jsonparser.GetString(value, "curr_id")

		objcurr.Curr_id = curr_id
		arraobjcurr = append(arraobjcurr, objcurr)
	})
	if !flag {
		result, err := models.Fetch_companyHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompany_home_redis, result, 60*time.Minute)
		fmt.Println("COMPANY MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY CACHE")
		return c.JSON(fiber.Map{
			"status":   fiber.StatusOK,
			"message":  "Success",
			"record":   arraobj,
			"listcurr": arraobjcurr,
			"time":     time.Since(render_page).String(),
		})
	}
}
func Companyadminrulehome(c *fiber.Ctx) error {
	var obj entities.Model_companyadminrule
	var arraobj []entities.Model_companyadminrule
	var objcompany entities.Model_companyshare
	var arraobjcompany []entities.Model_companyshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompanyadminrule_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listcompany_RD, _, _, _ := jsonparser.Get(jsonredis, "listcompany")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		companyadminrule_id, _ := jsonparser.GetInt(value, "companyadminrule_id")
		companyadminrule_idcompany, _ := jsonparser.GetString(value, "companyadminrule_idcompany")
		companyadminrule_nmrule, _ := jsonparser.GetString(value, "companyadminrule_nmrule")
		companyadminrule_rule, _ := jsonparser.GetString(value, "companyadminrule_rule")
		companyadminrule_create, _ := jsonparser.GetString(value, "companyadminrule_create")
		companyadminrule_update, _ := jsonparser.GetString(value, "companyadminrule_update")

		obj.Companyadminrule_id = int(companyadminrule_id)
		obj.Companyadminrule_idcompany = companyadminrule_idcompany
		obj.Companyadminrule_nmrule = companyadminrule_nmrule
		obj.Companyadminrule_rule = companyadminrule_rule
		obj.Companyadminrule_create = companyadminrule_create
		obj.Companyadminrule_update = companyadminrule_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listcompany_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_id, _ := jsonparser.GetString(value, "company_id")
		company_name, _ := jsonparser.GetString(value, "company_name")

		objcompany.Company_id = company_id
		objcompany.Company_name = company_name
		arraobjcompany = append(arraobjcompany, objcompany)
	})
	if !flag {
		result, err := models.Fetch_companyadminruleHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompanyadminrule_home_redis, result, 60*time.Minute)
		fmt.Println("COMPANY ADMIN GROUP MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY ADMIN GROUP CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"listcompany": arraobjcompany,
			"time":        time.Since(render_page).String(),
		})
	}
}
func Companyadminhome(c *fiber.Ctx) error {
	var obj entities.Model_companyadmin
	var arraobj []entities.Model_companyadmin
	var objcompany entities.Model_companyshare
	var arraobjcompany []entities.Model_companyshare
	var objrule entities.Model_companyadminrule_share
	var arraobjrule []entities.Model_companyadminrule_share
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompanyadmin_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listcompany_RD, _, _, _ := jsonparser.Get(jsonredis, "listcompany")
	listrule_RD, _, _, _ := jsonparser.Get(jsonredis, "listrule")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		companyadmin_id, _ := jsonparser.GetString(value, "companyadmin_id")
		companyadmin_idrule, _ := jsonparser.GetInt(value, "companyadmin_idrule")
		companyadmin_idcompany, _ := jsonparser.GetString(value, "companyadmin_idcompany")
		companyadmin_tipe, _ := jsonparser.GetString(value, "companyadmin_tipe")
		companyadmin_nmrule, _ := jsonparser.GetString(value, "companyadmin_nmrule")
		companyadmin_username, _ := jsonparser.GetString(value, "companyadmin_username")
		companyadmin_ipaddress, _ := jsonparser.GetString(value, "companyadmin_ipaddress")
		companyadmin_lastlogin, _ := jsonparser.GetString(value, "companyadmin_lastlogin")
		companyadmin_name, _ := jsonparser.GetString(value, "companyadmin_name")
		companyadmin_phone1, _ := jsonparser.GetString(value, "companyadmin_phone1")
		companyadmin_phone2, _ := jsonparser.GetString(value, "companyadmin_phone2")
		companyadmin_status, _ := jsonparser.GetString(value, "companyadmin_status")
		companyadmin_status_css, _ := jsonparser.GetString(value, "companyadmin_status_css")
		companyadmin_create, _ := jsonparser.GetString(value, "companyadmin_create")
		companyadmin_update, _ := jsonparser.GetString(value, "companyadmin_update")

		obj.Companyadmin_id = companyadmin_id
		obj.Companyadmin_idcompany = companyadmin_idcompany
		obj.Companyadmin_idrule = int(companyadmin_idrule)
		obj.Companyadmin_tipe = companyadmin_tipe
		obj.Companyadmin_nmrule = companyadmin_nmrule
		obj.Companyadmin_username = companyadmin_username
		obj.Companyadmin_ipaddress = companyadmin_ipaddress
		obj.Companyadmin_lastlogin = companyadmin_lastlogin
		obj.Companyadmin_name = companyadmin_name
		obj.Companyadmin_phone1 = companyadmin_phone1
		obj.Companyadmin_phone2 = companyadmin_phone2
		obj.Companyadmin_status = companyadmin_status
		obj.Companyadmin_status_css = companyadmin_status_css
		obj.Companyadmin_create = companyadmin_create
		obj.Companyadmin_update = companyadmin_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listcompany_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_id, _ := jsonparser.GetString(value, "company_id")
		company_name, _ := jsonparser.GetString(value, "company_name")

		objcompany.Company_id = company_id
		objcompany.Company_name = company_name
		arraobjcompany = append(arraobjcompany, objcompany)
	})
	jsonparser.ArrayEach(listrule_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		companyadminrule_id, _ := jsonparser.GetInt(value, "companyadminrule_id")
		companyadminrule_idcompany, _ := jsonparser.GetString(value, "companyadminrule_idcompany")
		companyadminrule_nmrule, _ := jsonparser.GetString(value, "companyadminrule_nmrule")

		objrule.Companyadminrule_id = int(companyadminrule_id)
		objrule.Companyadminrule_idcompany = companyadminrule_idcompany
		objrule.Companyadminrule_nmrule = companyadminrule_nmrule
		arraobjrule = append(arraobjrule, objrule)
	})
	if !flag {
		result, err := models.Fetch_companyadminHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompanyadmin_home_redis, result, 60*time.Minute)
		fmt.Println("COMPANY ADMIN  MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY ADMIN  CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"listcompany": arraobjcompany,
			"listrule":    arraobjrule,
			"time":        time.Since(render_page).String(),
		})
	}
}
func CompanyadminByCompany(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companygroupcompany)
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

	var obj entities.Model_companyadmin
	var arraobj []entities.Model_companyadmin
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompanyadmin_home_redis + "_" + client.Company_id)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		companyadmin_id, _ := jsonparser.GetString(value, "companyadmin_id")
		companyadmin_idrule, _ := jsonparser.GetInt(value, "companyadmin_idrule")
		companyadmin_idcompany, _ := jsonparser.GetString(value, "companyadmin_idcompany")
		companyadmin_tipe, _ := jsonparser.GetString(value, "companyadmin_tipe")
		companyadmin_nmrule, _ := jsonparser.GetString(value, "companyadmin_nmrule")
		companyadmin_username, _ := jsonparser.GetString(value, "companyadmin_username")
		companyadmin_ipaddress, _ := jsonparser.GetString(value, "companyadmin_ipaddress")
		companyadmin_lastlogin, _ := jsonparser.GetString(value, "companyadmin_lastlogin")
		companyadmin_name, _ := jsonparser.GetString(value, "companyadmin_name")
		companyadmin_phone1, _ := jsonparser.GetString(value, "companyadmin_phone1")
		companyadmin_phone2, _ := jsonparser.GetString(value, "companyadmin_phone2")
		companyadmin_status, _ := jsonparser.GetString(value, "companyadmin_status")
		companyadmin_status_css, _ := jsonparser.GetString(value, "companyadmin_status_css")
		companyadmin_create, _ := jsonparser.GetString(value, "companyadmin_create")
		companyadmin_update, _ := jsonparser.GetString(value, "companyadmin_update")

		obj.Companyadmin_id = companyadmin_id
		obj.Companyadmin_idcompany = companyadmin_idcompany
		obj.Companyadmin_idrule = int(companyadmin_idrule)
		obj.Companyadmin_tipe = companyadmin_tipe
		obj.Companyadmin_nmrule = companyadmin_nmrule
		obj.Companyadmin_username = companyadmin_username
		obj.Companyadmin_ipaddress = companyadmin_ipaddress
		obj.Companyadmin_lastlogin = companyadmin_lastlogin
		obj.Companyadmin_name = companyadmin_name
		obj.Companyadmin_phone1 = companyadmin_phone1
		obj.Companyadmin_phone2 = companyadmin_phone2
		obj.Companyadmin_status = companyadmin_status
		obj.Companyadmin_status_css = companyadmin_status_css
		obj.Companyadmin_create = companyadmin_create
		obj.Companyadmin_update = companyadmin_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_companyadminByCompany(client.Company_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompanyadmin_home_redis+"_"+client.Company_id, result, 60*time.Minute)
		fmt.Println("COMPANY ADMIN BY COMPANY  MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY ADMIN BY COMPANY CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companylistbethome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companylistbet)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	fmt.Println(client.Company_id)
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

	var obj entities.Model_company_listbet
	var arraobj []entities.Model_company_listbet
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompanylistbet_home_redis + "_" + client.Company_id)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		companylistbet_id, _ := jsonparser.GetInt(value, "companylistbet_id")
		companylistbet_minbet, _ := jsonparser.GetFloat(value, "companylistbet_minbet")
		companylistbet_create, _ := jsonparser.GetString(value, "companylistbet_create")
		companylistbet_update, _ := jsonparser.GetString(value, "companylistbet_update")

		obj.Companylistbet_id = int(companylistbet_id)
		obj.Companylistbet_minbet = float64(companylistbet_minbet)
		obj.Companylistbet_create = companylistbet_create
		obj.Companylistbet_update = companylistbet_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_companyListBet(client.Company_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompanylistbet_home_redis+"_"+client.Company_id, result, 60*time.Minute)
		fmt.Println("COMPANY LISTBET MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY LISTBET CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Companyconfpointhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyconfpoint)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	fmt.Println(client.Company_id)
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

	var obj entities.Model_company_conf
	var arraobj []entities.Model_company_conf
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompanyconf_home_redis + "_" + strconv.Itoa(client.Company_idbet) + "_" + client.Company_id)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		companyconf_id, _ := jsonparser.GetInt(value, "companyconf_id")
		companyconf_idbet, _ := jsonparser.GetInt(value, "companyconf_idbet")
		companyconf_nmpoin, _ := jsonparser.GetString(value, "companyconf_nmpoin")
		companyconf_poin, _ := jsonparser.GetInt(value, "companyconf_poin")
		companyconf_create, _ := jsonparser.GetString(value, "companyconf_create")
		companyconf_update, _ := jsonparser.GetString(value, "companyconf_update")

		obj.Companyconf_id = int(companyconf_id)
		obj.Companyconf_idbet = int(companyconf_idbet)
		obj.Companyconf_nmpoin = companyconf_nmpoin
		obj.Companyconf_poin = int(companyconf_poin)
		obj.Companyconf_create = companyconf_create
		obj.Companyconf_update = companyconf_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_companyConfPoint(client.Company_idbet, client.Company_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompanyconf_home_redis+"_"+strconv.Itoa(client.Company_idbet)+"_"+client.Company_id, result, 60*time.Minute)
		fmt.Println("COMPANY CONF MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY CONF CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CompanySave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companysave)
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

	// aadmin, idrecord, code, idcurr, nmcompany, nmowner, phoneowner, emailowner, url, status, sData string
	result, err := models.Save_company(
		client_admin,
		client.Company_id, client.Company_idcurr,
		client.Company_name, client.Company_nmowner, client.Company_phoneowner, client.Company_emailowner,
		client.Company_url, client.Company_status,
		client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_company(0, "")
	return c.JSON(result)
}
func CompanyadminruleSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyadminrulesave)
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

	// admin, idrecord, idcompany, name, rule, sData string
	result, err := models.Save_companyadminrule(
		client_admin,
		client.Companyadminrule_idcompany,
		client.Companyadminrule_nmrule, client.Companyadminrule_rule, client.Sdata, client.Companyadminrule_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_company(0, "")
	return c.JSON(result)
}
func CompanyadminSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyadminsave)
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

	// admin, idrecord, idcompany, idrule, username, password, name, phone1, phone2, status, sData string
	result, err := models.Save_companyadmin(
		client_admin,
		client.Companyadmin_id, client.Companyadmin_idcompany,
		client.Companyadmin_username, client.Companyadmin_password, client.Companyadmin_name,
		client.Companyadmin_phone1, client.Companyadmin_phone2, client.Companyadmin_status, client.Sdata, client.Companyadmin_idrule)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_company(0, "")
	return c.JSON(result)
}
func CompanylistbetSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companylistbetsave)
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

	// admin, idcompany, sData string, idrecord, minbet int
	result, err := models.Save_companyListBet(
		client_admin,
		client.Companylistbet_idcompany,
		client.Sdata, client.Companylistbet_id, client.Companylistbet_minbet)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_company(0, client.Companylistbet_idcompany)
	return c.JSON(result)
}
func CompanyconfpointSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyconfpointsave)
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

	// admin, idcompany, sData string, idrecord, idbet, point int
	result, err := models.Save_companyConfPoint(
		client_admin,
		client.Companyconfpoint_idcompany,
		client.Sdata, client.Companyconfpoint_id, client.Companyconfpoint_idbet, client.Companyconfpoint_point)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_company(client.Companyconfpoint_idbet, client.Companyconfpoint_idcompany)
	return c.JSON(result)
}
func _deleteredis_company(idbet int, idcompany string) {
	val_master := helpers.DeleteRedis(Fieldcompany_home_redis)
	fmt.Printf("Redis Delete BACKEND COMPANY : %d", val_master)

	val_master_admin_bycompany := helpers.DeleteRedis(Fieldcompanyadmin_home_redis + "_" + idcompany)
	fmt.Printf("Redis Delete BACKEND COMPANY ADMIN GROUP COMPANY  : %d", val_master_admin_bycompany)

	val_master_admin := helpers.DeleteRedis(Fieldcompanyadmin_home_redis)
	fmt.Printf("Redis Delete BACKEND COMPANY ADMIN  : %d", val_master_admin)

	val_master_adminrule := helpers.DeleteRedis(Fieldcompanyadminrule_home_redis)
	fmt.Printf("Redis Delete BACKEND COMPANY ADMIN RULE : %d", val_master_adminrule)

	val_master_listbet := helpers.DeleteRedis(Fieldcompanylistbet_home_redis + "_" + idcompany)
	fmt.Printf("Redis Delete BACKEND COMPANY LISTBET : %d", val_master_listbet)

	val_master_confbet := helpers.DeleteRedis(Fieldcompanyconf_home_redis + "_" + strconv.Itoa(idbet) + "_" + idcompany)
	fmt.Printf("Redis Delete BACKEND COMPANY CONF POINT : %d", val_master_confbet)
}
