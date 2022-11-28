package models

// JwtToken is a JWT token.
type JwtToken struct {
	Token      string `json:"token"`
	Expiration int64  `json:"expiration"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
