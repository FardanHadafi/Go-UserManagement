package web

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Token     string `json:"token,omitempty"`
	CreatedAt string `json:"created_at"`
}