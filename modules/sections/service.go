package sections

import (
"github.com/emicklei/go-restful"
"strconv"
"net/http"
"api-application/helpers"
"encoding/json"
)

func (self *Section) SectionWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/sections")
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(self.SectionList).
		Doc("Список секций").
		Operation("GetSections"))
	ws.Route(ws.POST("/").To(self.CreateSection).
		Doc("Создание секций").
		Operation("CreateSection"))
	ws.Route(ws.PUT("/").To(self.UpdateSection).
		Doc("Обновление секций").
		Operation("UpdateSection"))
	ws.Route(ws.DELETE("/").To(self.DeleteSection).
		Doc("Удаление секций").
		Operation("DeleteSection"))
	id := ws.PathParameter("id", "ID!").DataType("int")
	ws.Route(ws.GET("/{id}").To(self.SectionById).
		Doc("Список секций").
		Operation("GetSections").Param(id))
	return ws
}

func (self *Section) SectionList(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity(self.GetAll())
}

func (self *Section) SectionById(req *restful.Request, resp *restful.Response) {
	id,err := strconv.Atoi(req.PathParameter("id"))
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest, err.Error()))
		return
	}
	if self.FindById(int64(id)) {
		resp.WriteEntity(self)
		return
	}
	resp.WriteHeaderAndEntity(http.StatusNotFound, helpers.AnwerMessage(http.StatusNotFound, "Ничего не найдено."))
}

func (self *Section) CreateSection(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(self)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	if self.Create() {
		resp.WriteEntity(self)
		return
	}
	resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка создания пользователя"))
}

func (self *Section) UpdateSection(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(self)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	if self.Id != 0 {
		if self.Update() {
			resp.WriteEntity(self)
			return
		}
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка обновления пользователя"))
		return
	}
	if self.Create() {
		resp.WriteEntity(self)
		return
	}
	resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка создания пользователя"))
}
func (self *Section) DeleteSection(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(&self)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	if self.Id != 0 {
		if self.Delete() {
			resp.WriteEntity(helpers.AnwerMessage(http.StatusOK, "Пользователь успешно удален"))
			return
		}
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка удаления пользователя"))
		return
	}
	resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest, "отсутствует номер пользователя"))
}


