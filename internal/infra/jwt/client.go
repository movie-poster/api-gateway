package jwt

import (
	objectValues "api-gateway/internal/domain/object_values"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

func NewJwtClient() JwtClient {
	return JwtClient{}
}

type JwtClient struct{}

func (j *JwtClient) GenerateToken(entityJWT objectValues.JwtEntity) (string, error) {
	claims := objectValues.JwtCustomClaims{
		Name:     entityJWT.Name,
		Surname:  entityJWT.Surname,
		Email:    entityJWT.Email,
		UserId:   entityJWT.UserId,
		NickName: entityJWT.NickName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}

func (j *JwtClient) VerifyToken(token string) (bool, *objectValues.JwtCustomClaims) {
	obj, err := jwt.ParseWithClaims(token, &objectValues.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("algoritmo de firma inválido")
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return false, &objectValues.JwtCustomClaims{}
	}

	return obj.Valid, obj.Claims.(*objectValues.JwtCustomClaims)
}

func (j *JwtClient) GenerateTokenResetPassword(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, objectValues.JwtCustomClaimsResetPassword{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	})

	return token.SignedString([]byte("token_for_reset_password_Nestor&Bugitech_asdfghjkl"))
}

func (j *JwtClient) VerifyTokenResetPassword(token string) bool {
	obj, err := jwt.ParseWithClaims(token, &objectValues.JwtCustomClaimsResetPassword{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("algoritmo de firma inválido")
		}
		return []byte("token_for_reset_password_Nestor&Bugitech_asdfghjkl"), nil
	})

	if err != nil {
		return false
	}

	return obj.Valid
}

func (j *JwtClient) GenerateTokenRegistrationTemporalLink(email, name string, businessId, typePeopleId uint64, mustExpire bool, expiresIn uint64) (string, error) {

	claims := objectValues.JwtCustomClaimsRegistrationLink{
		Email:        email,
		Name:         name,
		BusinessID:   businessId,
		TypePeopleID: typePeopleId,
		MustExpire:   mustExpire,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiresIn)).Unix(),
		},
	}

	if !mustExpire {
		claims.StandardClaims = jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1000).Unix(),
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("token_for_reset_password_Nestor&Bugitech_asdfghjkl"))
}

func (j *JwtClient) VerifyTokenRegistrationTemporalLink(token string) (bool, *objectValues.JwtCustomClaimsRegistrationLink) {
	obj, err := jwt.ParseWithClaims(token, &objectValues.JwtCustomClaimsRegistrationLink{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("algoritmo de firma inválido")
		}
		return []byte("token_for_reset_password_Nestor&Bugitech_asdfghjkl"), nil
	})

	if err != nil {
		return false, &objectValues.JwtCustomClaimsRegistrationLink{}
	}

	return obj.Valid, obj.Claims.(*objectValues.JwtCustomClaimsRegistrationLink)
}

func (j *JwtClient) GetConfig() middleware.JWTConfig {
	config := middleware.JWTConfig{
		Claims:     &objectValues.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	return config
}
