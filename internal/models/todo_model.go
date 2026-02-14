type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UsersID   string `json:"users_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}