package catalog

import (
"github.com/emicklei/go-restful"
"strconv"
"net/http"
"api-application/helpers"
"encoding/json"
)

func (self *Сatalog) СatalogWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/catalog")
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(self.СatalogList).
		Doc("Список каталога").
		Operation("СatalogList"))
	ws.Route(ws.POST("/").To(self.CreateСatalog).
		Doc("Создание каталога").
		Operation("CreateСatalog"))
	ws.Route(ws.PUT("/").To(self.UpdateСatalog).
		Doc("Обновление каталога").
		Operation("UpdateСatalog"))
	ws.Route(ws.DELETE("/").To(self.DeleteСatalog).
		Doc("Удаление каталога").
		Operation("DeleteСatalog"))
	id := ws.PathParameter("id", "ID!").DataType("int")
	ws.Route(ws.GET("/{id}").To(self.СatalogById).
		Doc("Поиск каталога").
		Operation("СatalogById").Param(id))
	return ws
}

func (self *Сatalog) СatalogList(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity(self.GetAll())
}

func (self *Сatalog) СatalogById(req *restful.Request, resp *restful.Response) {
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

func (self *Сatalog) CreateСatalog(req *restful.Request, resp *restful.Response) {
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
	resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка создания каталога"))
}

func (self *Сatalog) UpdateСatalog(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(self)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	if self.IsEmpty() {
		if self.Create() {
			resp.WriteEntity(self)
			return
		}
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка обновления каталога"))
		return
	}
	if self.Update() {
		resp.WriteEntity(self)
		return
	}
	resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка создания каталога"))
}
func (self *Сatalog) DeleteСatalog(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(&self)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	if !self.IsEmpty() {
		if self.Delete() {
			resp.WriteEntity(helpers.AnwerMessage(http.StatusOK, "Удаление каталога выполненно успешно"))
			return
		}
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка удаления каталога"))
		return
	}
	resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest, "отсутствует номер каталога"))
}



