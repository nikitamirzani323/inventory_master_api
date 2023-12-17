package entities

type Model_purchaserequest struct {
	Purchaserequest_id            string  `json:"purchaserequest_id"`
	Purchaserequest_iddepartement string  `json:"purchaserequest_iddepartement"`
	Purchaserequest_idemployee    string  `json:"purchaserequest_idemployee"`
	Purchaserequest_idcurr        string  `json:"purchaserequest_idcurr"`
	Purchaserequest_tipedoc       string  `json:"purchaserequest_tipedoc"`
	Purchaserequest_periodedoc    string  `json:"purchaserequest_periodedoc"`
	Purchaserequest_nmdepartement string  `json:"purchaserequest_nmdepartement"`
	Purchaserequest_nmemployee    string  `json:"purchaserequest_nmemployee"`
	Purchaserequest_totalitem     float64 `json:"purchaserequest_totalitem"`
	Purchaserequest_totalpr       float64 `json:"purchaserequest_totalpr"`
	Purchaserequest_totalpo       float64 `json:"purchaserequest_totalpo"`
	Purchaserequest_status        string  `json:"purchaserequest_status"`
	Purchaserequest_status_css    string  `json:"purchaserequest_status_css"`
	Purchaserequest_create        string  `json:"purchaserequest_create"`
	Purchaserequest_update        string  `json:"purchaserequest_update"`
}

type Controller_purchaserequestsave struct {
	Page                          string `json:"page" validate:"required"`
	Sdata                         string `json:"sdata" validate:"required"`
	Purchaserequest_search        string `json:"purchaserequest_search"`
	Purchaserequest_page          int    `json:"purchaserequest_page"`
	Purchaserequest_id            string `json:"purchaserequest_id"`
	Purchaserequest_iddepartement string `json:"purchaserequest_iddepartement" validate:"required"`
	Purchaserequest_idemployee    string `json:"purchaserequest_idemployee" validate:"required"`
	Purchaserequest_idcurr        string `json:"purchaserequest_idcurr" validate:"required"`
	Purchaserequest_tipedoc       string `json:"purchaserequest_tipedoc" validate:"required"`
	Purchaserequest_status        string `json:"purchaserequest_status" validate:"required"`
}
type Controller_purchaserequest struct {
	Purchaserequest_search string `json:"purchaserequest_search"`
	Purchaserequest_status string `json:"purchaserequest_status"`
	Purchaserequest_page   int    `json:"purchaserequest_page"`
}
