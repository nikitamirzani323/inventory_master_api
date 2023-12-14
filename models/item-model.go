package models

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/db"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nleeper/goment"
)

const database_cateitem_local = configs.DB_tbl_mst_categoryitem
const database_item_local = configs.DB_tbl_mst_item

func Fetch_catetemHome() (helpers.Response, error) {
	var obj entities.Model_cateitem
	var arraobj []entities.Model_cateitem
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idcateitem , nmcateitem, statuscateitem,   
			createcateitem, to_char(COALESCE(createdatecateitem,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updatecateitem, to_char(COALESCE(updatedatecateitem,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_cateitem_local + `  
			ORDER BY createdatecateitem DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcateitem_db                                                                      int
			nmcateitem_db, statuscateitem_db                                                   string
			createcateitem_db, createdatecateitem_db, updatecateitem_db, updatedatecateitem_db string
		)

		err = row.Scan(&idcateitem_db, &nmcateitem_db, &statuscateitem_db,
			&createcateitem_db, &createdatecateitem_db, &updatecateitem_db, &updatedatecateitem_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createcateitem_db != "" {
			create = createcateitem_db + ", " + createdatecateitem_db
		}
		if updatecateitem_db != "" {
			update = updatecateitem_db + ", " + updatedatecateitem_db
		}
		if statuscateitem_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Cateitem_id = idcateitem_db
		obj.Cateitem_name = nmcateitem_db
		obj.Cateitem_status = statuscateitem_db
		obj.Cateitem_status_css = status_css
		obj.Cateitem_create = create
		obj.Cateitem_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_itemHome() (helpers.Response, error) {
	var obj entities.Model_item
	var arraobj []entities.Model_item
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			A.iditem , A.idcateitem, B.nmcateitem, 
			A.nmitem , A.descpitem, A.inventory_item, A.sales_item, A.purchase_item, A.statusitem,   
			A.createitem, to_char(COALESCE(A.createdateitem,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			A.updateitem, to_char(COALESCE(A.updatedateitem,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_item_local + ` as A  
			JOIN ` + database_cateitem_local + ` as B ON B.idcateitem = A.idcateitem   
			ORDER BY A.createdateitem DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcateitem_db                                                                                                        int
			iditem_db, nmcateitem_db, nmitem_db, descpitem_db, inventory_item_db, sales_item_db, purchase_item_db, statusitem_db string
			createitem_db, createdateitem_db, updateitem_db, updatedateitem_db                                                   string
		)

		err = row.Scan(&iditem_db, &idcateitem_db, &nmcateitem_db, &nmitem_db,
			&descpitem_db, &inventory_item_db, &sales_item_db, &purchase_item_db, &statusitem_db,
			&createitem_db, &createdateitem_db, &updateitem_db, &updatedateitem_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createitem_db != "" {
			create = createitem_db + ", " + createdateitem_db
		}
		if updateitem_db != "" {
			update = updateitem_db + ", " + updatedateitem_db
		}
		if statusitem_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Item_id = iditem_db
		obj.Item_idcateitem = idcateitem_db
		obj.Item_nmcateitem = nmcateitem_db
		obj.Item_name = nmitem_db
		obj.Item_descp = descpitem_db
		obj.Item_inventory = inventory_item_db
		obj.Item_sales = sales_item_db
		obj.Item_purchase = purchase_item_db
		obj.Item_status = statusitem_db
		obj.Item_status_css = status_css
		obj.Item_create = create
		obj.Item_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_cateitem(admin, idrecord, name, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {

		sql_insert := `
				insert into
				` + database_cateitem_local + ` (
					idcateitem , nmcateitem, statuscateitem,  
					createcateitem, createdatecateitem 
				) values (
					$1, $2, $3,   
					$4, $5 
				)
			`
		field_column := database_cateitem_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_cateitem_local, "INSERT",
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter), name, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_cateitem_local + `  
				SET nmcateitem=$1, statuscateitem=$2,  
				updatecateitem=$3, updatedatecateitem=$4    
				WHERE idcateitem=$5   
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_cateitem_local, "UPDATE",
			name, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

		if flag_update {
			msg = "Succes"
		} else {
			fmt.Println(msg_update)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_item(admin, idrecord, name, descp, inventory, sales, purchase, status, sData string, idcateitem int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		// A.iditem , A.idcateitem, B.nmcateitem,
		// A.nmitem , A.descpitem, A.inventory_item, A.sales_item, A.purchase_item, A.statusitem,
		// A.createitem, to_char(COALESCE(A.createdateitem,now()), 'YYYY-MM-DD HH24:MI:SS'),
		// A.updateitem, to_char(COALESCE(A.updatedateitem,now()), 'YYYY-MM-DD HH24:MI:SS')
		sql_insert := `
				insert into
				` + database_item_local + ` (
					iditem , idcateitem, nmitem,  
					descpitem , inventory_item, sales_item, purchase_item, statusitem,
					createitem, createdateitem 
				) values (
					$1, $2, $3,   
					$4, $5, $6, $7, $8,  
					$9, $10 
				)
			`
		field_column := database_item_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_item_local, "INSERT",
			"ITEM_"+tglnow.Format("YY")+tglnow.Format("MM")+tglnow.Format("DD")+tglnow.Format("HH")+strconv.Itoa(idrecord_counter), idcateitem, name,
			descp, inventory, sales, purchase, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_item_local + `  
				SET idcateitem=$1, nmitem=$2,  
				descpitem=$3, inventory_item=$4, sales_item=$5, purchase_item=$6, statusitem=$7,
				updateitem=$8, updatedateitem=$9     
				WHERE iditem=$10    
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_item_local, "UPDATE",
			idcateitem, name, descp, inventory, sales, purchase, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

		if flag_update {
			msg = "Succes"
		} else {
			fmt.Println(msg_update)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
