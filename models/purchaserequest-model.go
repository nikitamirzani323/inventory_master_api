package models

import (
	"context"
	"database/sql"
	"fmt"
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

const database_purchaserequest_local = configs.DB_tbl_trx_purchaserequest
const database_purchaserequestdetail_local = configs.DB_tbl_trx_purchaserequest_detail

func Fetch_purchaserequestHome(search string, page int) (helpers.Responsepurchaserequest, error) {
	var obj entities.Model_purchaserequest
	var arraobj []entities.Model_purchaserequest
	var objbranch entities.Model_branchshare
	var arraobjbranch []entities.Model_branchshare
	var objdepartement entities.Model_departementshare
	var arraobjdepartement []entities.Model_departementshare
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	var res helpers.Responsepurchaserequest
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
	sql_selectcount += "COUNT(idpurchaserequest) as totalpurchase  "
	sql_selectcount += "FROM " + database_purchaserequest_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(idpurchaserequest) LIKE '%" + strings.ToLower(search) + "%' "
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
	sql_select += "A.idpurchaserequest, A.idbranch,B.nmbranch, A.iddepartement, C.nmdepartement, "
	sql_select += "A.idemployee, D.nmemployee , A.idcurr, A.tipe_document, A.periode_document, A.statupurchaserequest,  "
	sql_select += "A.createpurchaserequest, to_char(COALESCE(A.createdatepurchaserequest,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "A.updatepurchaserequest, to_char(COALESCE(A.updatedatepurchaserequest,now()), 'YYYY-MM-DD HH24:MI:SS')  "
	sql_select += "FROM " + database_purchaserequest_local + " as A   "
	sql_select += "JOIN " + configs.DB_tbl_mst_branch + " as B ON B.idbranch = A.idbranch   "
	sql_select += "JOIN " + configs.DB_tbl_mst_departement + " as C ON C.iddepartement = A.iddepartement   "
	sql_select += "JOIN " + configs.DB_tbl_mst_employee + " as D ON D.idemployee = A.idemployee   "
	if search == "" {
		sql_select += "ORDER BY A.createdatepurchaserequest DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)

	} else {
		sql_selectcount += "WHERE LOWER(A.idpurchaserequest) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY A.createdatepurchaserequest DESC   LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idpurchaserequest_db, idbranch_db, nmbranch_db, iddepartement_db, nmdepartement_db                             string
			idemployee_db, nmemployee_db, idcurr_db, tipe_document_db, periode_document_db, statupurchaserequest_db        string
			createpurchaserequest_db, createdatepurchaserequest_db, updatepurchaserequest_db, updatedatepurchaserequest_db string
		)

		err = row.Scan(&idpurchaserequest_db, &idbranch_db, &nmbranch_db, &iddepartement_db, &nmdepartement_db,
			&idemployee_db, &nmemployee_db, &idcurr_db, &tipe_document_db, &periode_document_db, &statupurchaserequest_db,
			&createpurchaserequest_db, &createdatepurchaserequest_db, &updatepurchaserequest_db, &updatedatepurchaserequest_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createpurchaserequest_db != "" {
			create = createpurchaserequest_db + ", " + createdatepurchaserequest_db
		}
		if updatepurchaserequest_db != "" {
			update = updatepurchaserequest_db + ", " + updatedatepurchaserequest_db
		}
		switch statupurchaserequest_db {
		case "OPEN":
			status_css = configs.STATUS_NEW2
		case "PROCESS":
			status_css = configs.STATUS_RUNNING
		case "COMPLETE":
			status_css = configs.STATUS_COMPLETE
		case "CANCEL":
			status_css = configs.STATUS_CANCEL
		}

		obj.Purchaserequest_id = idpurchaserequest_db
		obj.Purchaserequest_idcurr = idcurr_db
		obj.Purchaserequest_idbranch = idbranch_db
		obj.Purchaserequest_nmbranch = nmbranch_db
		obj.Purchaserequest_iddepartement = iddepartement_db
		obj.Purchaserequest_nmdepartement = nmdepartement_db
		obj.Purchaserequest_idemployee = idemployee_db
		obj.Purchaserequest_nmemployee = nmemployee_db
		obj.Purchaserequest_tipedoc = tipe_document_db
		obj.Purchaserequest_periodedoc = periode_document_db
		obj.Purchaserequest_totalitem = 0
		obj.Purchaserequest_totalpr = 0
		obj.Purchaserequest_totalpo = 0
		obj.Purchaserequest_status = statupurchaserequest_db
		obj.Purchaserequest_status_css = status_css
		obj.Purchaserequest_create = create
		obj.Purchaserequest_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	sql_selectbranch := `SELECT 
			idbranch, nmbranch  
			FROM ` + configs.DB_tbl_mst_branch + ` 
			WHERE statusbranch = 'Y' 
			ORDER BY nmbranch ASC    
	`
	rowbranch, errbranch := con.QueryContext(ctx, sql_selectbranch)
	helpers.ErrorCheck(errbranch)
	for rowbranch.Next() {
		var (
			idbranch_db, nmbranch_db string
		)

		errbranch = rowbranch.Scan(&idbranch_db, &nmbranch_db)

		helpers.ErrorCheck(errbranch)

		objbranch.Branch_id = idbranch_db
		objbranch.Branch_name = nmbranch_db
		arraobjbranch = append(arraobjbranch, objbranch)
		msg = "Success"
	}
	defer rowbranch.Close()

	sql_selectdepartement := `SELECT 
			iddepartement, nmdepartement  
			FROM ` + configs.DB_tbl_mst_departement + ` 
			WHERE statusdepartement = 'Y' 
			ORDER BY nmdepartement ASC    
	`
	rowdepartement, errdepartement := con.QueryContext(ctx, sql_selectdepartement)
	helpers.ErrorCheck(errdepartement)
	for rowdepartement.Next() {
		var (
			iddepartement_db, nmdepartement_db string
		)

		errdepartement = rowdepartement.Scan(&iddepartement_db, &nmdepartement_db)

		helpers.ErrorCheck(errdepartement)

		objdepartement.Departement_id = iddepartement_db
		objdepartement.Departement_name = nmdepartement_db
		arraobjdepartement = append(arraobjdepartement, objdepartement)
		msg = "Success"
	}
	defer rowdepartement.Close()

	sql_selectcurr := `SELECT 
			idcurr 
			FROM ` + configs.DB_tbl_mst_curr + ` 
			ORDER BY idcurr ASC    
	`
	rowcurr, errcurr := con.QueryContext(ctx, sql_selectcurr)
	helpers.ErrorCheck(errcurr)
	for rowcurr.Next() {
		var (
			idcurr_db string
		)

		errcurr = rowcurr.Scan(&idcurr_db)

		helpers.ErrorCheck(errcurr)

		objcurr.Curr_id = idcurr_db
		arraobjcurr = append(arraobjcurr, objcurr)
		msg = "Success"
	}
	defer rowcurr.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listbranch = arraobjbranch
	res.Listdepartement = arraobjdepartement
	res.Listcurr = arraobjcurr
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_purchaserequest(admin, idrecord, idbranch, iddepartement, idemployee, idcurr, tipedoc, status, listdetail, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
				insert into
				` + database_purchaserequest_local + ` (
					idpurchaserequest , idbranch, iddepartement, idemployee, idcurr,  
					tipe_document , periode_document, statupurchaserequest,  
					createpurchaserequest, createdatepurchaserequest 
				) values (
					$1, $2, $3, $4, $5,      
					$6, $7, $8, 
					$9, $10   
				)
			`
		field_column := database_purchaserequest_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		idrecord := "PR_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
		periode_doc := tglnow.Format("MM") + tglnow.Format("DDDD")
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_purchaserequest_local, "INSERT",
			idrecord, idbranch, iddepartement, idemployee, idcurr,
			tipedoc, periode_doc, "OPEN",
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"

			json := []byte(listdetail)
			jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				detail_iditem, _ := jsonparser.GetString(value, "detail_iditem")
				detail_nmitem, _ := jsonparser.GetString(value, "detail_nmitem")
				detail_descpitem, _ := jsonparser.GetString(value, "detail_descpitem")
				detail_purpose, _ := jsonparser.GetString(value, "detail_purpose")
				detail_iduom, _ := jsonparser.GetString(value, "detail_iduom")
				detail_qty, _ := jsonparser.GetFloat(value, "detail_qty")
				detail_estimateprice, _ := jsonparser.GetFloat(value, "detail_estimateprice")

				Save_purchaserequestdetail(admin, "", idrecord,
					detail_iditem, detail_nmitem, detail_descpitem, detail_purpose, detail_iduom,
					"PROCESS", "New", detail_qty, detail_estimateprice)
			})
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_purchaserequest_local + `  
				SET statupurchaserequest=$1, 
				updatepurchaserequest=$2, updatedatepurchaserequest=$3         
				WHERE idpurchaserequest=$4       
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_purchaserequest_local, "UPDATE",
			status,
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
func Save_purchaserequestdetail(admin, idrecord, idpurchaserequest, iditem, nmitem, descpitem, purpose, iduom, status, sData string, qty, estimateprice float64) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
				insert into
				` + database_purchaserequestdetail_local + ` (
					idpurchaserequestdetail, idpurchaserequest ,  
					iditem , nmitem, descitem,  purpose,
					qty , iduom, estimateprice,  statupurchaserequestdetail,
					createpurchaserequestdetail, createdatepurchaserequestdetail 
				) values (
					$1, $2, 
					$3, $4, $5, $6, 
					$7, $8, $9, $10,    
					$11, $12 
				)
			`
		field_column := database_purchaserequestdetail_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		idrecord := "PRDETAIL_" + tglnow.Format("YY") + tglnow.Format("MM") + tglnow.Format("DD") + tglnow.Format("HH") + strconv.Itoa(idrecord_counter)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_purchaserequestdetail_local, "INSERT",
			idrecord, idpurchaserequest,
			iditem, nmitem, descpitem, purpose,
			qty, iduom, estimateprice, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_purchaserequestdetail_local + `  
				SET iditem=$1, nmitem=$2, descitem=$3,  purpose=$4,
				qty=$5 , iduom=$6, estimateprice=$7,  statupurchaserequestdetail=$8,
				updatepurchaserequest=$9, updatedatepurchaserequest=$10          
				WHERE idpurchaserequest=$11       
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_purchaserequestdetail_local, "UPDATE",
			iditem, nmitem, descpitem, purpose,
			qty, iduom, estimateprice, status,
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
