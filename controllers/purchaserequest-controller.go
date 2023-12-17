package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/models"
)

const Fieldpurchaserequest_home_redis = "PURCAHSEREQUEST_BACKEND"

func Purchaserequesthome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_purchaserequest)
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

	if client.Purchaserequest_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldpurchaserequest_home_redis + "_" + strconv.Itoa(client.Purchaserequest_page) + "_" + client.Purchaserequest_search)
		fmt.Printf("Redis Delete BACKEND DEPARTEMENT : %d", val_pattern)
	}

	var obj entities.Model_purchaserequest
	var arraobj []entities.Model_purchaserequest
	var objdepartement entities.Model_departementshare
	var arraobjdepartement []entities.Model_departementshare
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldpurchaserequest_home_redis + "_" + strconv.Itoa(client.Purchaserequest_page) + "_" + client.Purchaserequest_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	Listdepartement_RD, _, _, _ := jsonparser.Get(jsonredis, "listdepartement")
	Listcurr_RD, _, _, _ := jsonparser.Get(jsonredis, "listcurr")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		purchaserequest_id, _ := jsonparser.GetString(value, "purchaserequest_id")
		purchaserequest_iddepartement, _ := jsonparser.GetString(value, "purchaserequest_iddepartement")
		purchaserequest_idemployee, _ := jsonparser.GetString(value, "purchaserequest_idemployee")
		purchaserequest_idcurr, _ := jsonparser.GetString(value, "purchaserequest_idcurr")
		purchaserequest_tipedoc, _ := jsonparser.GetString(value, "purchaserequest_tipedoc")
		purchaserequest_periodedoc, _ := jsonparser.GetString(value, "purchaserequest_periodedoc")
		purchaserequest_nmdepartement, _ := jsonparser.GetString(value, "purchaserequest_nmdepartement")
		purchaserequest_nmemployee, _ := jsonparser.GetString(value, "purchaserequest_nmemployee")
		purchaserequest_totalitem, _ := jsonparser.GetFloat(value, "purchaserequest_totalitem")
		purchaserequest_totalpr, _ := jsonparser.GetFloat(value, "purchaserequest_totalpr")
		purchaserequest_totalpo, _ := jsonparser.GetFloat(value, "purchaserequest_totalpo")
		purchaserequest_status, _ := jsonparser.GetString(value, "purchaserequest_status")
		purchaserequest_status_css, _ := jsonparser.GetString(value, "purchaserequest_status_css")
		purchaserequest_create, _ := jsonparser.GetString(value, "purchaserequest_create")
		purchaserequest_update, _ := jsonparser.GetString(value, "purchaserequest_update")

		obj.Purchaserequest_id = purchaserequest_id
		obj.Purchaserequest_iddepartement = purchaserequest_iddepartement
		obj.Purchaserequest_idemployee = purchaserequest_idemployee
		obj.Purchaserequest_idcurr = purchaserequest_idcurr
		obj.Purchaserequest_tipedoc = purchaserequest_tipedoc
		obj.Purchaserequest_periodedoc = purchaserequest_periodedoc
		obj.Purchaserequest_nmdepartement = purchaserequest_nmdepartement
		obj.Purchaserequest_nmemployee = purchaserequest_nmemployee
		obj.Purchaserequest_totalitem = float64(purchaserequest_totalitem)
		obj.Purchaserequest_totalpr = float64(purchaserequest_totalpr)
		obj.Purchaserequest_totalpo = float64(purchaserequest_totalpo)
		obj.Purchaserequest_status = purchaserequest_status
		obj.Purchaserequest_status_css = purchaserequest_status_css
		obj.Purchaserequest_create = purchaserequest_create
		obj.Purchaserequest_update = purchaserequest_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(Listdepartement_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		departement_id, _ := jsonparser.GetString(value, "departement_id")
		departement_name, _ := jsonparser.GetString(value, "departement_name")

		objdepartement.Departement_id = departement_id
		objdepartement.Departement_name = departement_name
		arraobjdepartement = append(arraobjdepartement, objdepartement)
	})
	jsonparser.ArrayEach(Listcurr_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		curr_id, _ := jsonparser.GetString(value, "curr_id")

		objcurr.Curr_id = curr_id
		arraobjcurr = append(arraobjcurr, objcurr)
	})
	if !flag {
		result, err := models.Fetch_purchaserequestHome(client.Purchaserequest_search, client.Purchaserequest_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldpurchaserequest_home_redis+"_"+strconv.Itoa(client.Purchaserequest_page)+"_"+client.Purchaserequest_search, result, 60*time.Minute)
		fmt.Println("PURCHASE REQUEST MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("PURCHASE REQUEST CACHE")
		return c.JSON(fiber.Map{
			"status":          fiber.StatusOK,
			"message":         "Success",
			"record":          arraobj,
			"listdepartement": arraobjdepartement,
			"listcurr":        arraobjcurr,
			"perpage":         perpage_RD,
			"totalrecord":     totalrecord_RD,
			"time":            time.Since(render_page).String(),
		})
	}
}
func PurchaserequestSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_purchaserequestsave)
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

	// admin, idrecord, iddepartement, idemployee, idcurr, tipedoc, status, sData string
	result, err := models.Save_purchaserequest(
		client_admin,
		client.Purchaserequest_id, client.Purchaserequest_iddepartement, client.Purchaserequest_idemployee, client.Purchaserequest_idcurr,
		client.Purchaserequest_tipedoc, client.Purchaserequest_status, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_purchaserequest(client.Purchaserequest_search, client.Purchaserequest_page)
	return c.JSON(result)
}
func _deleteredis_purchaserequest(search string, page int) {
	val_master := helpers.DeleteRedis(Fieldpurchaserequest_home_redis + "_" + strconv.Itoa(page) + "_" + search)
	fmt.Printf("Redis Delete BACKEND PURCHASE REQUEST : %d\n", val_master)

}
