package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/models"
)

const Fieldwarehouse_home_redis = "LISTWAREHOUSE_BACKEND"

func Warehousehome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_warehouse)
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

	var obj entities.Model_warehouse
	var arraobj []entities.Model_warehouse
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldwarehouse_home_redis + "_" + strings.ToUpper(client.Branch_id))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		warehouse_id, _ := jsonparser.GetString(value, "warehouse_id")
		warehouse_idbranch, _ := jsonparser.GetString(value, "warehouse_idbranch")
		warehouse_nmbranch, _ := jsonparser.GetString(value, "warehouse_nmbranch")
		warehouse_name, _ := jsonparser.GetString(value, "warehouse_name")
		warehouse_alamat, _ := jsonparser.GetString(value, "warehouse_alamat")
		warehouse_phone1, _ := jsonparser.GetString(value, "warehouse_phone1")
		warehouse_phone2, _ := jsonparser.GetString(value, "warehouse_phone2")
		warehouse_status, _ := jsonparser.GetString(value, "warehouse_status")
		warehouse_status_css, _ := jsonparser.GetString(value, "warehouse_status_css")
		warehouse_create, _ := jsonparser.GetString(value, "warehouse_create")
		warehouse_update, _ := jsonparser.GetString(value, "warehouse_update")

		obj.Warehouse_id = warehouse_id
		obj.Warehouse_idbranch = warehouse_idbranch
		obj.Warehouse_nmbranch = warehouse_nmbranch
		obj.Warehouse_name = warehouse_name
		obj.Warehouse_alamat = warehouse_alamat
		obj.Warehouse_phone1 = warehouse_phone1
		obj.Warehouse_phone2 = warehouse_phone2
		obj.Warehouse_status = warehouse_status
		obj.Warehouse_status_css = warehouse_status_css
		obj.Warehouse_create = warehouse_create
		obj.Warehouse_update = warehouse_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_warehouseHome(client.Branch_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldwarehouse_home_redis+"_"+strings.ToUpper(client.Branch_id), result, 60*time.Minute)
		fmt.Println("WAREHOUSE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("WAREHOUSE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func WarehouseSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_warehousesave)
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

	// admin, idrecord, idbranch, name, alamat, phone1, phone2, status, sData string
	result, err := models.Save_warehouse(
		client_admin,
		client.Warehouse_id, client.Warehouse_idbranch, client.Warehouse_name, client.Warehouse_alamat,
		client.Warehouse_phone1, client.Warehouse_phone2, client.Warehouse_status,
		client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_warehouse(client.Warehouse_idbranch)
	return c.JSON(result)
}
func _deleteredis_warehouse(idbranch string) {
	val_master := helpers.DeleteRedis(Fieldwarehouse_home_redis + "_" + strings.ToUpper(idbranch))
	fmt.Printf("Redis Delete BACKEND WAREHOUSE : %d", val_master)

}
