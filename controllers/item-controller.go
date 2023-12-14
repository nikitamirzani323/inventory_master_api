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

const Fieldcateitem_home_redis = "LISTCATEITEM_BACKEND"
const Fielditem_home_redis = "LISTITEM_BACKEND"

func Cateitemhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_cateitem)
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
	fmt.Println(client.Cateitem_page)
	if client.Cateitem_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldcateitem_home_redis + "_" + strconv.Itoa(client.Cateitem_page) + "_" + client.Cateitem_search)
		fmt.Printf("Redis Delete BACKEND CATEITEM : %d", val_pattern)
	}

	var obj entities.Model_cateitem
	var arraobj []entities.Model_cateitem
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcateitem_home_redis + "_" + strconv.Itoa(client.Cateitem_page) + "_" + client.Cateitem_search)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		cateitem_id, _ := jsonparser.GetInt(value, "cateitem_id")
		cateitem_name, _ := jsonparser.GetString(value, "cateitem_name")
		cateitem_status, _ := jsonparser.GetString(value, "cateitem_status")
		cateitem_status_css, _ := jsonparser.GetString(value, "cateitem_status_css")
		cateitem_create, _ := jsonparser.GetString(value, "cateitem_create")
		cateitem_update, _ := jsonparser.GetString(value, "cateitem_update")

		obj.Cateitem_id = int(cateitem_id)
		obj.Cateitem_name = cateitem_name
		obj.Cateitem_status = cateitem_status
		obj.Cateitem_status_css = cateitem_status_css
		obj.Cateitem_create = cateitem_create
		obj.Cateitem_update = cateitem_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_catetemHome(client.Cateitem_search, client.Cateitem_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcateitem_home_redis+"_"+strconv.Itoa(client.Cateitem_page)+"_"+client.Cateitem_search, result, 60*time.Minute)
		fmt.Println("CATE ITEM MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("CATE ITEM CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"perpage":     perpage_RD,
			"totalrecord": totalrecord_RD,
			"time":        time.Since(render_page).String(),
		})
	}
}
func Itemhome(c *fiber.Ctx) error {
	var obj entities.Model_item
	var arraobj []entities.Model_item
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fielditem_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		item_id, _ := jsonparser.GetString(value, "item_id")
		item_idcateitem, _ := jsonparser.GetInt(value, "item_idcateitem")
		item_nmcateitem, _ := jsonparser.GetString(value, "item_nmcateitem")
		item_name, _ := jsonparser.GetString(value, "item_name")
		item_descp, _ := jsonparser.GetString(value, "item_descp")
		item_inventory, _ := jsonparser.GetString(value, "item_inventory")
		item_sales, _ := jsonparser.GetString(value, "item_sales")
		item_purchase, _ := jsonparser.GetString(value, "item_purchase")
		item_status, _ := jsonparser.GetString(value, "item_status")
		item_status_css, _ := jsonparser.GetString(value, "item_status_css")
		item_create, _ := jsonparser.GetString(value, "item_create")
		item_update, _ := jsonparser.GetString(value, "item_update")

		obj.Item_id = item_id
		obj.Item_idcateitem = int(item_idcateitem)
		obj.Item_nmcateitem = item_nmcateitem
		obj.Item_name = item_name
		obj.Item_descp = item_descp
		obj.Item_inventory = item_inventory
		obj.Item_sales = item_sales
		obj.Item_purchase = item_purchase
		obj.Item_status = item_status
		obj.Item_status_css = item_status_css
		obj.Item_create = item_create
		obj.Item_update = item_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_itemHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fielditem_home_redis, result, 60*time.Minute)
		fmt.Println("ITEM MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("ITEM CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CateitemSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_cateitemsave)
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

	// admin, idrecord, name, status, sData string
	result, err := models.Save_cateitem(
		client_admin,
		client.Cateitem_name, client.Cateitem_status, client.Sdata, client.Cateitem_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_item(client.Cateitem_search, client.Cateitem_page)
	return c.JSON(result)
}
func ItemSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_itemsave)
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

	// admin, idrecord, name, descp, inventory, sales, purchase, status, sData string, idcateitem int
	result, err := models.Save_item(
		client_admin,
		client.Item_id, client.Item_name, client.Item_descp, client.Item_inventory, client.Item_sales, client.Item_purchase, client.Item_status,
		client.Sdata, client.Item_idcateitem)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_item("", 0)
	return c.JSON(result)
}
func _deleteredis_item(search string, page int) {
	val_master := helpers.DeleteRedis(Fieldcateitem_home_redis + "_" + strconv.Itoa(page) + "_" + search)
	fmt.Printf("Redis Delete BACKEND CATEITEM : %d", val_master)

}
