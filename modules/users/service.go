package users

import (
	"github.com/emicklei/go-restful"
	"strconv"
	"net/http"
	"api-application/helpers"
	"encoding/json"
)

func (self *User) UserWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/users")
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(self.UserList).
		Doc("Список пользователей").
		Operation("GetUsers"))
	ws.Route(ws.POST("/").To(self.CreateUser).
		Doc("Создание пользователей").
		Operation("CreateUser"))
	ws.Route(ws.PUT("/").To(self.UpdateUser).
		Doc("Обновление пользователей").
		Operation("UpdateUser"))
	ws.Route(ws.DELETE("/").To(self.DeleteUser).
		Doc("Удаление пользователей").
		Operation("DeleteUser"))
	id := ws.PathParameter("id", "ID!").DataType("int")
	ws.Route(ws.GET("/{id}").To(self.UserById).
		Doc("Список пользователей").
		Operation("GetUsers").Param(id))
	return ws
}

func (self *User) UserList(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity(self.GetAll())
}

func (self *User) UserById(req *restful.Request, resp *restful.Response) {
	id,err := strconv.Atoi(req.PathParameter("id"))
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest, err.Error()))
		return
	}
	if self.FindById(int64(id)) {
		resp.WriteEntity(self)
		return
	}
	resp.WriteHeaderAndEntity(404, helpers.AnwerMessage(404, "Ничего не найдено."))
}

func (self *User) CreateUser(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(self)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	self.Password = self.GetMD5Hash(self.Password)
	if self.Create() {
		resp.WriteEntity(self)
		return
	}
	resp.WriteHeaderAndEntity(102, helpers.AnwerMessage(102, "Ошибка создания пользователя"))
}

func (self *User) UpdateUser(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(self)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	self.Password = self.GetMD5Hash(self.Password)
	if self.Id != 0 {
		if self.Update() {
			resp.WriteEntity(self)
			return
		}
		resp.WriteHeaderAndEntity(102, helpers.AnwerMessage(102, "Ошибка обновления пользователя"))
		return
	}
	if self.Create() {
		resp.WriteEntity(self)
		return
	}
	resp.WriteHeaderAndEntity(102, helpers.AnwerMessage(102, "Ошибка создания пользователя"))
}
func (self *User) DeleteUser(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(&self)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	if self.Id != 0 {
		if self.Delete() {
			resp.WriteEntity(helpers.AnwerMessage(200, "Пользователь успешно удален"))
			return
		}
		resp.WriteHeaderAndEntity(102, helpers.AnwerMessage(102, "Ошибка удаления пользователя"))
		return
	}
	resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest, "отсутствует номер пользователя"))
}

