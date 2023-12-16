package entities

type Model_vendor struct {
	Vendor_id         string `json:"vendor_id"`
	Vendor_name       string `json:"vendor_name"`
	Vendor_pic        string `json:"vendor_pic"`
	Vendor_alamat     string `json:"vendor_alamat"`
	Vendor_email      string `json:"vendor_email"`
	Vendor_phone1     string `json:"vendor_phone1"`
	Vendor_phone2     string `json:"vendor_phone2"`
	Vendor_status     string `json:"vendor_status"`
	Vendor_status_css string `json:"vendor_status_css"`
	Vendor_create     string `json:"vendor_create"`
	Vendor_update     string `json:"vendor_update"`
}
type Controller_vendorsave struct {
	Page          string `json:"page" validate:"required"`
	Sdata         string `json:"sdata" validate:"required"`
	Vendor_search string `json:"vendor_search"`
	Vendor_page   int    `json:"vendor_page"`
	Vendor_id     string `json:"vendor_id"`
	Vendor_name   string `json:"vendor_name" validate:"required"`
	Vendor_pic    string `json:"vendor_pic" validate:"required"`
	Vendor_alamat string `json:"vendor_alamat"`
	Vendor_email  string `json:"vendor_email"`
	Vendor_phone1 string `json:"vendor_phone1" validate:"required"`
	Vendor_phone2 string `json:"vendor_phone2"`
	Vendor_status string `json:"vendor_status" validate:"required"`
}
type Controller_vendor struct {
	Vendor_search string `json:"vendor_search"`
	Vendor_page   int    `json:"vendor_page"`
}
