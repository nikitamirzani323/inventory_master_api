package entities

type Model_branch struct {
	Branch_id     string `json:"branch_id"`
	Branch_name   string `json:"branch_name"`
	Branch_create string `json:"branch_create"`
	Branch_update string `json:"branch_update"`
}
type Controller_branchsave struct {
	Page        string `json:"page" validate:"required"`
	Sdata       string `json:"sdata" validate:"required"`
	Branch_id   string `json:"branch_id" validate:"required"`
	Branch_name string `json:"branch_name" validate:"required"`
}
