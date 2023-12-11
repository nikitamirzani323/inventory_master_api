package models

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/db"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nleeper/goment"
)

const database_warehouse_local = configs.DB_tbl_mst_warehouse

func Fetch_warehouseHome(idbranch string) (helpers.Responsewarehouse, error) {
	var obj entities.Model_warehouse
	var arraobj []entities.Model_warehouse
	var objbranch entities.Model_branchshare
	var arraobjbranch []entities.Model_branchshare
	var res helpers.Responsewarehouse
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			A.idwarehouse , A.idbranch, B.nmbranch, 
			A.nmwarehouse , A.alamatwarehouse, A.phone1warehouse, A.phone2warehouse, A.statuswarehouse, 
			A.createwarehouse, to_char(COALESCE(A.createdatewarehouse,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			A.updatewarehouse, to_char(COALESCE(A.updatedatewarehouse,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_warehouse_local + ` AS A   
			JOIN ` + configs.DB_tbl_mst_branch + ` AS B ON B.idbranch = A.idbranch    
			WHERE A.idbranch=$1 
			ORDER BY A.createdatewarehouse DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idwarehouse_db, idbranch_db, nmbranch_db                                                       string
			nmwarehouse_db, alamatwarehouse_db, phone1warehouse_db, phone2warehouse_db, statuswarehouse_db string
			createwarehouse_db, createdatewarehouse_db, updatewarehouse_db, updatedatewarehouse_db         string
		)

		err = row.Scan(&idwarehouse_db, &idbranch_db, &nmbranch_db,
			&nmwarehouse_db, &alamatwarehouse_db, &phone1warehouse_db, &phone2warehouse_db, &statuswarehouse_db,
			&createwarehouse_db, &createdatewarehouse_db, &updatewarehouse_db, &updatedatewarehouse_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createwarehouse_db != "" {
			create = createwarehouse_db + ", " + createdatewarehouse_db
		}
		if updatewarehouse_db != "" {
			update = updatewarehouse_db + ", " + updatedatewarehouse_db
		}
		if statuswarehouse_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Warehouse_id = idwarehouse_db
		obj.Warehouse_idbranch = idbranch_db
		obj.Warehouse_nmbranch = nmbranch_db
		obj.Warehouse_name = nmwarehouse_db
		obj.Warehouse_alamat = alamatwarehouse_db
		obj.Warehouse_phone1 = phone1warehouse_db
		obj.Warehouse_phone2 = phone2warehouse_db
		obj.Warehouse_status = statuswarehouse_db
		obj.Warehouse_status_css = status_css
		obj.Warehouse_create = create
		obj.Warehouse_update = update
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

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listbranch = arraobjbranch
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_warehouse(admin, idrecord, idbranch, name, alamat, phone1, phone2, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_warehouse_local, "idwarehouse", strings.ToUpper(idrecord))
		if !flag {
			sql_insert := `
				insert into
				` + database_warehouse_local + ` (
					idwarehouse , idbranch, 
					nmwarehouse , alamatwarehouse, phone1warehouse, phone2warehouse, statuswarehouse,  
					createwarehouse, createdatewarehouse 
				) values (
					$1, $2,    
					$3, $4, $5, $6, $7,    
					$8, $9   
				)
			`
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_warehouse_local, "INSERT",
				strings.ToUpper(idrecord), idbranch,
				name, alamat, phone1, phone2, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				msg = "Succes"
			} else {
				fmt.Println(msg_insert)
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_warehouse_local + `  
				SET nmwarehouse=$1, alamatwarehouse=$2, phone1warehouse=$3, phone2warehouse=$4, statuswarehouse=$5, 
				updatewarehouse=$6, updatedatewarehouse=$7      
				WHERE idwarehouse=$8 
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_warehouse_local, "UPDATE",
			name, alamat, phone1, phone2, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

		if flag_update {
			flag = true
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
