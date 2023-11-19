package dto

type TodosResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CheckFlag   bool   `json:"check_flag"`
}
