package models

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type CreateUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

type UpdateUserRequest struct {
	User User `json:"user"`
}

type UpdateUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type DeleteUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

type GetUsersResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Users   []User `json:"users"`
}

type Void struct{}

type UserFilter struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	User         User   `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
