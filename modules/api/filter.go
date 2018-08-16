package api

import (
	"strings"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/emicklei/go-restful"
	"api-application/helpers"
	"net/http"
)

type Auth struct {
	Id int64 `json:"id"`
	Login string `json:"login"`
	Role string `json:"role"`
	FullName string `json:"full_name"`
}

func (s *Api) JWTFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {

	if "OPTIONS" != req.Request.Method && strings.Index(req.Request.URL.String(), "token") != -1 {
		chain.ProcessFilter(req, resp)
		return
	}
	if "OPTIONS" != req.Request.Method && strings.Index(req.Request.URL.String(), "telegrams") != -1 {
		chain.ProcessFilter(req, resp)
		return
	}
	if "OPTIONS" != req.Request.Method && strings.Index(req.Request.URL.String(), "settings") != -1 {
		chain.ProcessFilter(req, resp)
		return
	}
	if "OPTIONS" != req.Request.Method && strings.Index(req.Request.URL.String(), "tickets") != -1 {
		chain.ProcessFilter(req, resp)
		return
	}
	if "OPTIONS" != req.Request.Method && strings.Index(req.Request.URL.String(), "images") != -1 {
		chain.ProcessFilter(req, resp)
		return
	}
	tokenHeader := req.Request.Header.Get("authorization")
	//Bearer
	if tokenHeader != "" {
		token, _ := jwt.Parse(strings.Replace(tokenHeader, "Bearer ", "", 1), func(token *jwt.Token) (interface{}, error) {

			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return s.Token, nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			s.ParseClaims(claims)
			chain.ProcessFilter(req, resp)
			return
		}
	}
	resp.WriteHeaderAndEntity(http.StatusForbidden, map[string]string{"code":"403","error":"Неверный токен авторизации"})
}

func (s *Api) Course(cors bool) {
	// Настройка доступа для ajax запросов
	if cors {
		corsRule := restful.CrossOriginResourceSharing{
			//ExposeHeaders: []string{"Content-Type"},
			AllowedDomains: []string{"http://invo.dev","http://herobrine.my","http://cabinet.kassa24.dev","http://192.168.150.52","http://pult.kassa24.kz","https://pult.kassa24.kz"},
			AllowedHeaders: []string{"content-type", "authorization"},
			CookiesAllowed: false,
			Container:      s.container,
		}
		s.container.Filter(corsRule.Filter)
	}
}
func (s *Api) Filter(filter bool) {
	// Настройка доступа для ajax запросов
	if filter {
		s.container.Filter(s.JWTFilter)
	}
}


func (s *Api) ParseClaims(claims jwt.MapClaims) Auth {
	id := int64(claims["id"].(float64))
	return Auth{
		id,
		fmt.Sprint("s%",claims["login"]),
		fmt.Sprint("s%",claims["role"]),
		fmt.Sprint("s%",claims["fullName"]),
	}
}

func (s *Api) CreateToken(auth Auth) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": auth.Id,
		"login": auth.Login,
		"role": auth.Role,
		"fullName": auth.FullName,
	})
	helpers.Info("%s авторизировался",auth.FullName, s.Token, string(s.Token))
	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(s.Token)
	return  tokenString
}