package model

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100,min=8"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type UpdateUserRequest struct {
	Name     string `json:"name,omitempty" validate:"max=100"`
	Username string `validate:"max=100"`
	Password string `json:"password,omitempty" validate:"max=100"`
}

type GetUserRequest struct {
	Username string `json:"username,omitempty"`
}

type VerifyUserRequest struct {
	AccessToken string `json:"access_token,omitempty"`
}

type UserResponse struct {
	ID        uint    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Username  string `json:"username,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
}
