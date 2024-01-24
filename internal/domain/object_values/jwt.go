package objectValues

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	UserId   uint64 `json:"user_id"`
	NickName string `json:"nick_name"`
	jwt.StandardClaims
}

type JwtCustomClaimsResetPassword struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

type JwtCustomClaimsRegistrationLink struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	BusinessID   uint64 `json:"business_id"`
	TypePeopleID uint64 `json:"type_people_id"`
	MustExpire   bool   `json:"must_expire"`
	jwt.StandardClaims
}

type JwtEntity struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	UserId   uint64 `json:"user_id"`
	NickName string `json:"nick_name"`
}
