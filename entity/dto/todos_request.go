package dto

type TodosRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoCheckFlag struct {
	CheckFlag bool `json:"check_flag"`
}
