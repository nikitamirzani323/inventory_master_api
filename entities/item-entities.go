package entities

type Model_cateitem struct {
	Cateitem_id         int    `json:"cateitem_id"`
	Cateitem_name       string `json:"cateitem_name"`
	Cateitem_status     string `json:"cateitem_status"`
	Cateitem_status_css string `json:"cateitem_status_css"`
	Cateitem_create     string `json:"cateitem_create"`
	Cateitem_update     string `json:"cateitem_update"`
}
type Model_cateitemshare struct {
	Cateitem_id   int    `json:"cateitem_id"`
	Cateitem_name string `json:"cateitem_name"`
}
type Model_item struct {
	Item_id            string `json:"item_id"`
	Item_idcateitem    int    `json:"item_idcateitem"`
	Item_nmcateitem    string `json:"item_nmcateitem"`
	Item_name          string `json:"item_name"`
	Item_descp         string `json:"item_descp"`
	Item_inventory     string `json:"item_inventory"`
	Item_sales         string `json:"item_sales"`
	Item_purchase      string `json:"item_purchase"`
	Item_inventory_css string `json:"item_inventory_css"`
	Item_sales_css     string `json:"item_sales_css"`
	Item_purchase_css  string `json:"item_purchase_css"`
	Item_status        string `json:"item_status"`
	Item_status_css    string `json:"item_status_css"`
	Item_create        string `json:"item_create"`
	Item_update        string `json:"item_update"`
}
type Controller_cateitemsave struct {
	Page            string `json:"page" validate:"required"`
	Sdata           string `json:"sdata" validate:"required"`
	Cateitem_search string `json:"cateitem_search"`
	Cateitem_page   int    `json:"cateitem_page"`
	Cateitem_id     int    `json:"cateitem_id" `
	Cateitem_name   string `json:"cateitem_name" validate:"required"`
	Cateitem_status string `json:"cateitem_status" validate:"required"`
}
type Controller_itemsave struct {
	Page            string `json:"page" validate:"required"`
	Sdata           string `json:"sdata" validate:"required"`
	Item_search     string `json:"item_search"`
	Item_page       int    `json:"item_page"`
	Item_id         string `json:"item_id"`
	Item_idcateitem int    `json:"item_idcateitem"  validate:"required"`
	Item_iduom      string `json:"item_iduom"  validate:"required"`
	Item_name       string `json:"item_name" validate:"required"`
	Item_descp      string `json:"item_descp"`
	Item_inventory  string `json:"item_inventory" validate:"required"`
	Item_sales      string `json:"item_sales" validate:"required"`
	Item_purchase   string `json:"item_purchase" validate:"required"`
	Item_status     string `json:"item_status" validate:"required"`
}
type Controller_item struct {
	Item_search string `json:"item_search"`
	Item_page   int    `json:"item_page"`
}
type Controller_cateitem struct {
	Cateitem_search string `json:"cateitem_search"`
	Cateitem_page   int    `json:"cateitem_page"`
}
