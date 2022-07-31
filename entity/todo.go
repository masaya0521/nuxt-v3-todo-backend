package entity

type Todo struct {
	ID      string `json:"todo_id"`
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`	
	Status  string `json:"status"`
}