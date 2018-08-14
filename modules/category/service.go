package category

import (
"github.com/emicklei/go-restful"
"strconv"
"net/http"
"api-application/helpers"
"encoding/json"
)

func (self *Category) CategoryWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/categories")
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(self.CategoryList).
		Doc("Список категорий").
		Operation("CategoryList"))
	ws.Route(ws.POST("/").To(self.CreateCategory).
		Doc("Создание категорий").
		Operation("CreateCategory"))
	ws.Route(ws.PUT("/").To(self.UpdateCategory).
		Doc("Обновление категорий").
		Operation("UpdateCategory"))
	ws.Route(ws.DELETE("/").To(self.DeleteCategory).
		Doc("Удаление категорий").
		Operation("DeleteCategory"))
	id := ws.PathParameter("id", "ID!").DataType("int")
	ws.Route(ws.GET("/{id}").To(self.CategoryById).
		Doc("Поиск категории").
		Operation("CategoryById").Param(id))
	return ws
}

func (self *Category) CategoryList(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity(self.GetAll())
}

func (self *Category) CategoryById(req *restful.Request, resp *restful.Response) {
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

func (self *Category) CreateCategory(req *restful.Request, resp *restful.Response) {
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
	resp.WriteHeaderAndEntity(102, helpers.AnwerMessage(102, "Ошибка создания категории"))
}

func (self *Category) UpdateCategory(req *restful.Request, resp *restful.Response) {
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
		resp.WriteHeaderAndEntity(102, helpers.AnwerMessage(102, "Ошибка обновления категории"))
		return
	}
	if self.Create() {
		resp.WriteEntity(self)
		return
	}
	resp.WriteHeaderAndEntity(102, helpers.AnwerMessage(102, "Ошибка создания категории"))
}
func (self *Category) DeleteCategory(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(&self)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	if self.Id != 0 {
		if self.Delete() {
			resp.WriteEntity(helpers.AnwerMessage(200, "Категория успешно удален"))
			return
		}
		resp.WriteHeaderAndEntity(102, helpers.AnwerMessage(102, "Ошибка удаления категории"))
		return
	}
	resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest, "отсутствует номер категории"))
}


