package models

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/mesment/microservice/user-service/pkg/util"
	"time"
)

// NilUser is the nil value for User
var NilUser User

type User struct {
	ID  string   `gorm:"primary_key"`
	UserName string
	Password string
	Email 	string `gorm:"not null;unique_index"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"create_at"`
	UpdatedAt time.Time  `gorm:"default:current_timestamp on update current_timestamp" json:"update_at"`
	DeletedAt *time.Time
}


// GetToken returns the User's JWT
func (u User) GetToken() string {
	byteSlc, _ := json.Marshal(u)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": string(byteSlc),
	})
	tokenString, _ := token.SignedString(util.JWTSecret)
	return tokenString
}

// UserFromToken returns the User which is authenticated with this Token
func UserFromToken(token string) (*User, error ){
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return util.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	var result = &User{}
	err = json.Unmarshal([]byte(claims["data"].(string)), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}