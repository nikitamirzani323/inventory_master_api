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
	Item_iduom         string `json:"item_iduom"`
	Item_name          string `json:"item_name"`
	Item_descp         string `json:"item_descp"`
	Item_urlimg        string `json:"item_urlimg"`
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
type Model_itemuom struct {
	Itemuom_id          int     `json:"itemuom_id"`
	Itemuom_iduom       string  `json:"itemuom_iduom"`
	Itemuom_nmuom       string  `json:"itemuom_nmuom"`
	Itemuom_default     string  `json:"itemuom_default"`
	Itemuom_default_css string  `json:"itemuom_default_css"`
	Itemuom_conversion  float32 `json:"itemuom_conversion"`
	Itemuom_create      string  `json:"itemuom_create"`
	Itemuom_update      string  `json:"itemuom_update"`
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
	Item_iduom      string `json:"item_iduom"  `
	Item_name       string `json:"item_name" validate:"required"`
	Item_descp      string `json:"item_descp"`
	Item_urlimg     string `json:"item_urlimg"`
	Item_inventory  string `json:"item_inventory" validate:"required"`
	Item_sales      string `json:"item_sales" validate:"required"`
	Item_purchase   string `json:"item_purchase" validate:"required"`
	Item_status     string `json:"item_status" validate:"required"`
}
type Controller_itemuomsave struct {
	Page               string  `json:"page" validate:"required"`
	Sdata              string  `json:"sdata" validate:"required"`
	Itemuom_search     string  `json:"itemuom_search"`
	Itemuom_page       int     `json:"itemuom_page"`
	Itemuom_id         int     `json:"itemuom_id"`
	Itemuom_iditem     string  `json:"itemuom_iditem"  validate:"required"`
	Itemuom_iduom      string  `json:"itemuom_iduom"  validate:"required"`
	Itemuom_default    string  `json:"itemuom_default" validate:"required"`
	Itemuom_conversion float32 `json:"itemuom_conversion" validate:"required"`
}
type Controller_itemuomdelete struct {
	Page           string `json:"page" validate:"required"`
	Itemuom_search string `json:"itemuom_search"`
	Itemuom_page   int    `json:"itemuom_page"`
	Itemuom_id     int    `json:"itemuom_id" validate:"required"`
	Itemuom_iditem string `json:"itemuom_iditem"  validate:"required"`
}
type Controller_item struct {
	Item_search string `json:"item_search"`
	Item_page   int    `json:"item_page"`
}
type Controller_cateitem struct {
	Cateitem_search string `json:"cateitem_search"`
	Cateitem_page   int    `json:"cateitem_page"`
}
type Controller_itemuom struct {
	Item_id string `json:"item_id"`
}
