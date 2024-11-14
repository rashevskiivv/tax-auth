package entity

// User cross layers user entity.
type User struct {
	ID       *int64  `json:"id,omitempty"`
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}