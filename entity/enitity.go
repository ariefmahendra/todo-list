package entity

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CheckFlag   bool   `json:"check_flag"`
}
