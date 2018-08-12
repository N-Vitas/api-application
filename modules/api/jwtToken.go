package api

import (
	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	FullName string
	UserId   int64
}

func (s *Api) ParseClaims(claims jwt.MapClaims) Auth {
	a := Auth{}
	a.FullName, _ = claims["userId"].(string)
	a.UserId, _ = claims["userAgentId"].(int64)
	//"userName": FullName.String,
	//	"userId": IDSysUser.Int64,
	//	"userLogin": Name.String,
	//	"userAgentId": IDAgents.Int64,
	//	"userPhone" : CellPhone.String,
	//	"created": time.Now().Unix(),
	return a
}
