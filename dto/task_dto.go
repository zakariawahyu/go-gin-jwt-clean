package dto

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int64  `json:"user_id"`
}
