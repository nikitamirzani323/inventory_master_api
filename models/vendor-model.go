package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/db"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nleeper/goment"
)

const database_vendor_local = configs.DB_tbl_mst_vendor

func Fetch_vendorHome(search string, page int) (helpers.Responsepaging, error) {
	var obj entities.Model_vendor
	var arraobj []entities.Model_vendor
	var res helpers.Responsepaging
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	perpage := configs.PAGING_PAGE
	totalrecord := 0
	offset := page
	sql_selectcount := ""
	sql_selectcount += ""
	sql_selectcount += "SELECT "
	sql_selectcount += "COUNT(idvendor) as totalvendor  "
	sql_selectcount += "FROM " + database_vendor_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(idvendor) LIKE '%" + strings.ToLower(search) + "%' "
		sql_selectcount += "OR LOWER(nmvendor) LIKE '%" + strings.ToLower(search) + "%' "
	}
	row_selectcount := con.QueryRowContext(ctx, sql_selectcount)
	switch e_selectcount := row_selectcount.Scan(&totalrecord); e_selectcount {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e_selectcount)
	}

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "idvendor , nmvendor, picvendor, "
	sql_select += "alamatvendor, emailvendor, phone1vendor, phone2vendor, statusvendor,  "
	sql_select += "createvendor, to_char(COALESCE(createdatevendor,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "updatevendor, to_char(COALESCE(updatedatevendor,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_vendor_local + "   "
	if search == "" {
		sql_selectcount += "WHERE LOWER(idvendor) LIKE '%" + strings.ToLower(search) + "%' "
		sql_selectcount += "OR LOWER(nmvendor) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY createdatevendor DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_selectcount += "WHERE LOWER(idvendor) LIKE '%" + strings.ToLower(search) + "%' "
		sql_selectcount += "OR LOWER(nmvendor) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY createdatevendor DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idvendor_db, nmvendor_db, picvendor_db                                             string
			alamatvendor_db, emailvendor_db, phone1vendor_db, phone2vendor_db, statusvendor_db string
			createvendor_db, createdatevendor_db, updatevendor_db, updatedatevendor_db         string
		)

		err = row.Scan(&idvendor_db, &nmvendor_db, &picvendor_db, &alamatvendor_db,
			&emailvendor_db, &phone1vendor_db, &phone2vendor_db, &statusvendor_db,
			&createvendor_db, &createdatevendor_db, &updatevendor_db, &updatedatevendor_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createvendor_db != "" {
			create = createvendor_db + ", " + createdatevendor_db
		}
		if updatevendor_db != "" {
			update = updatevendor_db + ", " + updatedatevendor_db
		}
		if statusvendor_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Vendor_id = idvendor_db
		obj.Vendor_name = nmvendor_db
		obj.Vendor_pic = picvendor_db
		obj.Vendor_alamat = alamatvendor_db
		obj.Vendor_email = emailvendor_db
		obj.Vendor_phone1 = phone1vendor_db
		obj.Vendor_phone2 = phone2vendor_db
		obj.Vendor_status = statusvendor_db
		obj.Vendor_status_css = status_css
		obj.Vendor_create = create
		obj.Vendor_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_vendor(admin, idrecord, name, pic, alamat, email, phone1, phone2, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
				insert into
				` + database_vendor_local + ` (
					idvendor , nmvendor, picvendor, 
					alamatvendor,emailvendor, phone1vendor, phone2vendor, statusvendor,
					createvendor, createdatevendor 
				) values (
					$1, $2, $3,    
					$4, $5, $6, $7, $8,  
					$9, $10 
				)
			`
		field_column := database_vendor_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		idrecord := "VENDOR_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_vendor_local, "INSERT",
			idrecord, name, pic,
			alamat, email, phone1, phone2, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_vendor_local + `  
				SET nmvendor=$1, picvendor=$2, alamatvendor=$3,  
				emailvendor=$4, phone1vendor=$5, phone2vendor=$6 ,statusvendor=$7, 
				updatevendor=$8, updatedatevendor=$9       
				WHERE idvendor=$10     
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_vendor_local, "UPDATE",
			name, pic, alamat, email, phone1, phone2, status,
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
