package users

import (
	"github.com/emicklei/go-restful"
	"strconv"
	"net/http"
	"api-application/helpers"
)

func (self *User) UserWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/users")
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(self.UserList).
		Doc("Список пользователей").
		Operation("GetUsers"))
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

