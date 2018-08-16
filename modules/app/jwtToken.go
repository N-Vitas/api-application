package app

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"api-application/helpers"
	"github.com/emicklei/go-restful"
	"api-application/modules/users"
	"api-application/modules/api"
	"encoding/json"
	"net/http"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func (self *App) TokenWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/token")
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(self.Welcome).
		Doc("Welcome").
		Operation("Welcome"))
	ws.Route(ws.POST("/").To(self.GetToken).
		Doc("GetToken").
		Operation("GetToken"))
	return ws
}


func (self *App) Welcome(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity(helpers.AnwerMessage(http.StatusOK,"Добро пожаловать"))
}

func (s *App) ParseClaims(claims jwt.MapClaims) api.Auth {
	id := claims["id"].(int64)
	return api.Auth{
		id,
		fmt.Sprint("s%",claims["login"]),
		fmt.Sprint("s%",claims["role"]),
		fmt.Sprint("s%",claims["fullName"]),
	}
}

func (s *App) CreateToken(auth api.Auth) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": auth.Id,
		"login": auth.Login,
		"role": auth.Role,
		"fullName": auth.FullName,
	})
	helpers.Info("%s авторизировался",auth.FullName)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(s.Api.Token)
	return  tokenString
}

func (d *App) GetToken(req *restful.Request, resp *restful.Response)  {
	// Get aliases from session manager
	document := Login{}
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(&document)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	user := users.NewUser()
	user.SetSession(d.sessMng)
	if user.FindByName(document.Username) {
		if user.Password == user.GetMD5Hash(document.Password) {
			token := d.Api.CreateToken(api.Auth{user.Id,user.Login,user.Role,user.FullName})
			resp.WriteEntity(helpers.AnwerMessage(http.StatusOK,token))
			return
		}
	}
	resp.WriteHeaderAndEntity(http.StatusForbidden, helpers.AnwerMessage(http.StatusForbidden,"Неверный логин или пароль"))
}