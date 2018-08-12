package api

import (
	"strings"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/emicklei/go-restful"
)

func (s *Api) JWTFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {

	if "OPTIONS" != req.Request.Method && strings.Index(req.Request.URL.String(), "login") != -1 {
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
	resp.WriteHeaderAndEntity(403, map[string]string{"code":"403","error":"Неверный токен авторизации"})
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
