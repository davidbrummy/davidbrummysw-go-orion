package model

type AuthCredentials struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

func NewAuthCredentials() *AuthCredentials {
	return &AuthCredentials{}
}

func NewAuthCredentialsFill(email *string, password *string) *AuthCredentials {
	return &AuthCredentials{Email: email, Password: password}
}
