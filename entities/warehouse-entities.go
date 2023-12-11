package entities

type Model_warehouse struct {
	Warehouse_id         string `json:"warehouse_id"`
	Warehouse_idbranch   string `json:"warehouse_idbranch"`
	Warehouse_nmbranch   string `json:"warehouse_nmbranch"`
	Warehouse_name       string `json:"warehouse_name"`
	Warehouse_alamat     string `json:"warehouse_alamat"`
	Warehouse_phone1     string `json:"warehouse_phone1"`
	Warehouse_phone2     string `json:"warehouse_phone2"`
	Warehouse_status     string `json:"warehouse_status"`
	Warehouse_status_css string `json:"warehouse_status_css"`
	Warehouse_create     string `json:"warehouse_create"`
	Warehouse_update     string `json:"warehouse_update"`
}
type Controller_warehousesave struct {
	Page               string `json:"page" validate:"required"`
	Sdata              string `json:"sdata" validate:"required"`
	Warehouse_id       string `json:"warehouse_id" validate:"required"`
	Warehouse_idbranch string `json:"warehouse_idbranch" validate:"required"`
	Warehouse_name     string `json:"warehouse_name" validate:"required"`
	Warehouse_alamat   string `json:"warehouse_alamat" validate:"required"`
	Warehouse_phone1   string `json:"warehouse_phone1" `
	Warehouse_phone2   string `json:"warehouse_phone2" `
	Warehouse_status   string `json:"warehouse_status" validate:"required"`
}
type Controller_warehouse struct {
	Branch_id string `json:"branch_id" `
}
