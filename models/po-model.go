package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/db"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nleeper/goment"
)

const database_po_local = configs.DB_tbl_trx_po
const database_podetail_local = configs.DB_tbl_trx_po_detail

func Fetch_poHome(search string, page int) (helpers.Responsepaging, error) {
	var obj entities.Model_po
	var arraobj []entities.Model_po
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
	sql_selectcount += "COUNT(idrfq) as totalrfq  "
	sql_selectcount += "FROM " + database_po_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(idrfq) LIKE '%" + strings.ToLower(search) + "%' "
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
	sql_select += "A.idpo,  "
	sql_select += "A.idrfq, A.idbranch, B.nmbranch, A.idvendor, C.nmvendor, "
	sql_select += "A.idcurr, A.tipe_docpo, A.statuspo,   "
	sql_select += "A.po_discount, A.po_ppn, A.po_pph,   "
	sql_select += "A.po_totalitem, A.po_subtotal, A.po_grandtotal,   "
	sql_select += "A.createpo, to_char(COALESCE(A.createdatepo,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "A.updatepo, to_char(COALESCE(A.updatedatepo,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_po_local + " as A   "
	sql_select += "JOIN " + configs.DB_tbl_mst_branch + " as B ON B.idbranch = A.idbranch   "
	sql_select += "JOIN " + configs.DB_tbl_mst_vendor + " as C ON C.idvendor = A.idvendor   "
	if search == "" {
		sql_select += "ORDER BY A.createdatepo DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_select += "WHERE LOWER(A.idpo) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY A.createdatepo DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idpo_db, idrfq_db, idbranch_db, nmbranch_db, idvendor_db, nmvendor_db string
			idcurr_db, tipe_docpo_db, statuspo_db                                 string
			po_discount_db, po_ppn_db, po_pph_db                                  float64
			po_totalitem_db, po_subtotal_db, po_grandtotal_db                     float64
			createpo_db, createdatepo_db, updatepo_db, updatedatepo_db            string
		)

		err = row.Scan(&idpo_db, &idrfq_db, &idbranch_db, &nmbranch_db, &idvendor_db, &nmvendor_db,
			&idcurr_db, &tipe_docpo_db, &statuspo_db,
			&po_discount_db, &po_ppn_db, &po_pph_db,
			&po_totalitem_db, &po_subtotal_db,
			&createpo_db, &createdatepo_db, &updatepo_db, &updatedatepo_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createpo_db != "" {
			create = createpo_db + ", " + createdatepo_db
		}
		if updatepo_db != "" {
			update = updatepo_db + ", " + updatedatepo_db
		}
		switch statuspo_db {
		case "OPEN":
			status_css = configs.STATUS_NEW2
		case "PROCESS":
			status_css = configs.STATUS_RUNNING
		case "COMPLETE":
			status_css = configs.STATUS_COMPLETE
		case "CANCEL":
			status_css = configs.STATUS_CANCEL
		}

		obj.Po_id = idpo_db
		obj.Po_idrfq = idrfq_db
		obj.Po_idbranch = idbranch_db
		obj.Po_idvendor = idvendor_db
		obj.Po_idcurr = idcurr_db
		obj.Po_date = createdatepo_db
		obj.Po_nmbranch = nmbranch_db
		obj.Po_nmvendor = nmvendor_db
		obj.Po_tipedoc = tipe_docpo_db
		obj.Po_discount = po_discount_db
		obj.Po_ppn = po_ppn_db
		obj.Po_ppn = po_pph_db
		obj.Po_totalitem = po_totalitem_db
		obj.Po_subtotal = po_subtotal_db
		obj.Po_grandtotal = po_grandtotal_db
		obj.Po_status = statuspo_db
		obj.Po_status_css = status_css
		obj.Po_create = create
		obj.Po_update = update
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
func Fetch_poDetail(idrfq string) (helpers.Response, error) {
	var obj entities.Model_rfqdetail
	var arraobj []entities.Model_rfqdetail
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "A.idrfqdetail, A.idpurchaserequestdetail, A.idpurchaserequest, "
	sql_select += "C.nmdepartement, B.idemployee, D.nmemployee,  "
	sql_select += "A.iditem, A.nmitem, A.descitem, "
	sql_select += "A.qty, A.iduom, A.price, A.statusrfqdetail,  "
	sql_select += "A.createrfqdetail, to_char(COALESCE(A.createdaterfqdetaildetail,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "A.updaterfqdetaildetail, to_char(COALESCE(A.updatedaterfqdetail,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_rfqdetail_local + " as A   "
	sql_select += "JOIN " + configs.DB_tbl_trx_purchaserequest + " as B ON B.idpurchaserequest = A.idpurchaserequest   "
	sql_select += "JOIN " + configs.DB_tbl_mst_departement + " as C ON C.iddepartement = B.iddepartement   "
	sql_select += "JOIN " + configs.DB_tbl_mst_employee + " as D ON D.idemployee = B.idemployee   "
	sql_select += "WHERE A.idrfq='" + idrfq + "' "
	sql_select += "ORDER BY A.createdaterfqdetaildetail ASC   "

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idrfqdetail_db, idpurchaserequestdetail_db, idpurchaserequest_db                                   string
			nmdepartement_db, idemployee_db, nmemployee_db                                                     string
			iditem_db, nmitem_db, descitem_db, iduom_db, statusrfqdetail_db                                    string
			qty_db, price_db                                                                                   float64
			createrfqdetail_db, createdaterfqdetaildetail_db, updaterfqdetaildetail_db, updatedaterfqdetail_db string
		)

		err = row.Scan(&idrfqdetail_db, &idpurchaserequestdetail_db, &idpurchaserequest_db,
			&nmdepartement_db, &idemployee_db, &nmemployee_db,
			&iditem_db, &nmitem_db, &descitem_db,
			&qty_db, &iduom_db, &price_db, &statusrfqdetail_db,
			&createrfqdetail_db, &createdaterfqdetaildetail_db, &updaterfqdetaildetail_db, &updatedaterfqdetail_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createrfqdetail_db != "" {
			create = createrfqdetail_db + ", " + createdaterfqdetaildetail_db
		}
		if updaterfqdetaildetail_db != "" {
			update = updaterfqdetaildetail_db + ", " + updatedaterfqdetail_db
		}
		switch statusrfqdetail_db {
		case "OPEN":
			status_css = configs.STATUS_NEW2
		case "PROCESS":
			status_css = configs.STATUS_RUNNING
		case "COMPLETE":
			status_css = configs.STATUS_COMPLETE
		case "CANCEL":
			status_css = configs.STATUS_CANCEL
		}

		obj.Rfqdetail_id = idrfqdetail_db
		obj.Rfqdetail_idpurchaserequestdetail = idpurchaserequestdetail_db
		obj.Rfqdetail_idpurchaserequest = idpurchaserequest_db
		obj.Rfqdetail_nmdepartement = nmdepartement_db
		obj.Rfqdetail_nmemployee = idemployee_db + " - " + nmemployee_db
		obj.Rfqdetail_iditem = iditem_db
		obj.Rfqdetail_nmitem = nmitem_db
		obj.Rfqdetail_descitem = descitem_db
		obj.Rfqdetail_iduom = iduom_db
		obj.Rfqdetail_qty = float64(qty_db)
		obj.Rfqdetail_price = float64(price_db)
		obj.Rfqdetail_status = statusrfqdetail_db
		obj.Rfqdetail_status_css = status_css
		obj.Rfqdetail_create = create
		obj.Rfqdetail_update = update
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
func Save_po(admin, idrecord, idrfq, listdetail, sData string, discount, ppn, pph, total_item, subtotal, grandtotal float32) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	con := db.CreateCon()
	ctx := context.Background()
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		idbranch_rfq := ""
		idvendor_rfq := ""
		idcurr_rfq := ""
		tipedoc_rfq := ""

		sql_select := `SELECT
		 	idbranch, idvendor, idcurr, tipe_documentrfq   
			FROM ` + configs.DB_tbl_trx_rfq + `  
			WHERE idrfq='` + idrfq + `'     
		`
		row := con.QueryRowContext(ctx, sql_select)
		switch e := row.Scan(&idbranch_rfq, &idvendor_rfq, &idcurr_rfq, &tipedoc_rfq); e {
		case sql.ErrNoRows:
		case nil:
		default:
			helpers.ErrorCheck(e)
		}

		sql_insert := `
				insert into
				` + database_po_local + ` (
					idrfq , idbranch, idvendor, idcurr,  
					po_discount, po_ppn, po_pph, po_totalitem, po_subtotal, po_grandtotal,
					tipe_docpo , statuspo,
					createpo, createdatepo 
				) values (
					$1, $2, $3, $4,     
					$5, $6, $7, $8, $9, $10, 
					$11, $12, 
					$13, $14
				)
			`

		field_column := database_po_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		idrecord := "PO_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
		start_date := tglnow.Format("YYYY-MM-DD HH:mm:ss")
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_po_local, "INSERT",
			idrecord, idbranch_rfq, idvendor_rfq, idcurr_rfq,
			discount, ppn, pph, total_item, subtotal, grandtotal,
			tipedoc_rfq, "OPEN",
			admin, start_date)

		if flag_insert {
			msg = "Succes"

			json := []byte(listdetail)
			jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				detail_id, _ := jsonparser.GetString(value, "detail_id")
				detail_document, _ := jsonparser.GetString(value, "detail_document")
				detail_iditem, _ := jsonparser.GetString(value, "detail_iditem")
				detail_nmitem, _ := jsonparser.GetString(value, "detail_nmitem")
				detail_descpitem, _ := jsonparser.GetString(value, "detail_descpitem")
				detail_iduom, _ := jsonparser.GetString(value, "detail_iduom")
				detail_qty, _ := jsonparser.GetFloat(value, "detail_qty")
				detail_price, _ := jsonparser.GetFloat(value, "detail_price")

				//admin, idrecord, idrfq, idpurchaserequestdetail, idpurchaserequest, iditem, nmitem, descpitem, iduom, status, sData string, qty, price float64
				Save_podetail(admin, "", idrecord,
					detail_id, detail_document,
					detail_iditem, detail_nmitem, detail_descpitem, detail_iduom,
					"OPEN", "New", detail_qty, detail_price)
			})
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		_, totaldetail_db := _Get_info_po(idrecord)
		log.Println("total : ", totaldetail_db)
		log.Println("data : ", listdetail)
		if totaldetail_db > 0 {
			sql_delete := `
				DELETE FROM  
				` + database_podetail_local + `   
				WHERE idpo=$1  
			`

			flag_delete, msg_delete := Exec_SQL(sql_delete, database_podetail_local, "DELETE", idrecord)

			if flag_delete {
				msg = "Succes"
				//UPDATE
				sql_update := `
					UPDATE 
					` + database_po_local + `  
					SET po_totalitem=$1, po_subtotal=$2, po_grandtotal=$3, 
					updatepo=$4, updatedatepo=$5          
					WHERE idpo=$6        
				`

				flag_update, msg_update := Exec_SQL(sql_update, database_po_local, "UPDATE",
					total_item, subtotal, grandtotal,
					admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

				if flag_update {
					msg = "Succes"

					json := []byte(listdetail)
					jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
						detail_id, _ := jsonparser.GetString(value, "detail_id")
						detail_document, _ := jsonparser.GetString(value, "detail_document")
						detail_iditem, _ := jsonparser.GetString(value, "detail_iditem")
						detail_nmitem, _ := jsonparser.GetString(value, "detail_nmitem")
						detail_descpitem, _ := jsonparser.GetString(value, "detail_descpitem")
						detail_iduom, _ := jsonparser.GetString(value, "detail_iduom")
						detail_qty, _ := jsonparser.GetFloat(value, "detail_qty")
						detail_price, _ := jsonparser.GetFloat(value, "detail_price")

						//admin, idrecord, idrfq, idpurchaserequestdetail, idpurchaserequest, iditem, nmitem, descpitem, iduom, status, sData string, qty, price float64
						Save_podetail(admin, "", idrecord,
							detail_id, detail_document,
							detail_iditem, detail_nmitem, detail_descpitem, detail_iduom,
							"OPEN", "New", detail_qty, detail_price)
					})
				} else {
					fmt.Println(msg_update)
				}
			} else {
				fmt.Println(msg_delete)
			}
		} else {
			sql_update := `
					UPDATE 
					` + database_po_local + `  
					SET po_totalitem=$1, po_subtotal=$2, po_grandtotal=$3, 
					updatepo=$4, updatedatepo=$5          
					WHERE idpo=$6        
				`

			flag_update, msg_update := Exec_SQL(sql_update, database_po_local, "UPDATE",
				total_item, subtotal, grandtotal,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"

				json := []byte(listdetail)
				jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					detail_id, _ := jsonparser.GetString(value, "detail_id")
					detail_document, _ := jsonparser.GetString(value, "detail_document")
					detail_iditem, _ := jsonparser.GetString(value, "detail_iditem")
					detail_nmitem, _ := jsonparser.GetString(value, "detail_nmitem")
					detail_descpitem, _ := jsonparser.GetString(value, "detail_descpitem")
					detail_iduom, _ := jsonparser.GetString(value, "detail_iduom")
					detail_qty, _ := jsonparser.GetFloat(value, "detail_qty")
					detail_price, _ := jsonparser.GetFloat(value, "detail_price")

					//admin, idrecord, idrfq, idpurchaserequestdetail, idpurchaserequest, iditem, nmitem, descpitem, iduom, status, sData string, qty, price float64
					Save_podetail(admin, "", idrecord,
						detail_id, detail_document,
						detail_iditem, detail_nmitem, detail_descpitem, detail_iduom,
						"OPEN", "New", detail_qty, detail_price)
				})
			} else {
				fmt.Println(msg_update)
			}
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_poStatus(admin, idrecord, status string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	status_db := ""
	total_detail_db := 0

	status_db, total_detail_db = _Get_info_rfq(idrecord)
	if status_db == "OPEN" {
		if total_detail_db > 0 {
			sql_update := `
				UPDATE 
				` + database_rfq_local + `  
				SET statusrfq=$1, 
				updaterfq=$2, updatedaterfq=$3     
				WHERE idrfq=$4    
			`

			flag_update, msg_update := Exec_SQL(sql_update, database_rfq_local, "UPDATE",
				status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"
				//DETAIL
				sql_updatedetail := `
					UPDATE 
					` + database_rfqdetail_local + `  
					SET statusrfqdetail=$1, 
					updaterfqdetaildetail=$2, updatedaterfqdetail=$3     
					WHERE idrfq=$4    
				`

				flag_updatedetail, msg_updatedetail := Exec_SQL(sql_updatedetail, database_rfqdetail_local, "UPDATE",
					status,
					admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

				if flag_updatedetail {
					msg = "Succes"
				} else {
					fmt.Println(msg_updatedetail)
				}
			} else {
				fmt.Println(msg_update)
			}
		}
	} else if status_db == "PROCESS" {
		if status == "CANCEL" {
			sql_update := `
				UPDATE 
				` + database_rfq_local + `  
				SET statusrfq=$1, 
				updaterfq=$2, updatedaterfq=$3     
				WHERE idrfq=$4    
			`

			flag_update, msg_update := Exec_SQL(sql_update, database_rfq_local, "UPDATE",
				"CANCEL",
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"
				//DETAIL
				sql_updatedetail := `
					UPDATE 
					` + database_rfqdetail_local + `  
					SET statusrfqdetail=$1, 
					updaterfqdetaildetail=$2, updatedaterfqdetail=$3     
					WHERE idrfq=$4    
				`

				flag_updatedetail, msg_updatedetail := Exec_SQL(sql_updatedetail, database_rfqdetail_local, "UPDATE",
					"CANCEL",
					admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

				if flag_updatedetail {
					msg = "Succes"
				} else {
					fmt.Println(msg_updatedetail)
				}
			} else {
				fmt.Println(msg_update)
			}
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_podetail(admin, idrecord, idrfq, idpurchaserequestdetail, idpurchaserequest, iditem, nmitem, descpitem, iduom, status, sData string, qty, price float64) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDBTwoField(database_podetail_local, "idrfq", idrfq, "idpurchaserequestdetail", idpurchaserequestdetail)
		if !flag {
			sql_insert := `
				insert into
				` + database_podetail_local + ` (
					idrfqdetail, idrfq, 
					idpurchaserequestdetail, idpurchaserequest,  
					iditem , nmitem, descitem,  
					qty , iduom, price,  statusrfqdetail,
					createrfqdetail, createdaterfqdetaildetail 
				) values (
					$1, $2, 
					$3, $4, 
					$5, $6, $7, 
					$8, $9, $10, $11,    
					$12, $13  
				)
			`
			field_column := database_podetail_local + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			idrecord := "PODETAIL_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_podetail_local, "INSERT",
				idrecord, idrfq,
				idpurchaserequestdetail, idpurchaserequest,
				iditem, nmitem, descpitem,
				qty, iduom, price, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				msg = "Succes"
			} else {
				fmt.Println(msg_insert)
			}
		} else {
			msg = "Duplicate Entry"
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func _Get_info_po(idpo string) (string, int) {
	con := db.CreateCon()
	ctx := context.Background()
	status := ""
	total_detail := 0
	sql_select := `SELECT
			statuspo  
			FROM ` + database_po_local + `  
			WHERE idpo='` + idpo + `'     
		`
	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&status); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	sql_selectdetail := `SELECT
			COUNT(idpodetail) AS total 
			FROM ` + database_podetail_local + `  
			WHERE idpo='` + idpo + `'     
		`
	rowdetail := con.QueryRowContext(ctx, sql_selectdetail)
	switch e := rowdetail.Scan(&total_detail); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return status, total_detail
}
