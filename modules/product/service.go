package product

import (
	"github.com/emicklei/go-restful"
	"strconv"
	"net/http"
	"api-application/helpers"
	"encoding/json"
)

func (self *Product) ProductWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/products")
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(self.ProductList).
		Doc("Список товаров").
		Operation("ProductList"))
	ws.Route(ws.POST("/").To(self.CreateProduct).
		Doc("Создание товаров").
		Operation("CreateProduct"))
	ws.Route(ws.PUT("/").To(self.UpdateProduct).
		Doc("Обновление товаров").
		Operation("UpdateProduct"))
	ws.Route(ws.DELETE("/").To(self.DeleteProduct).
		Doc("Удаление товаров").
		Operation("DeleteProduct"))
	id := ws.PathParameter("id", "ID!").DataType("int")
	ws.Route(ws.GET("/{id}").To(self.ProductById).
		Doc("Поиск товаров").
		Operation("ProductById").Param(id))
	ws.Route(ws.GET("/parents/{id}").To(self.ProductByParent).
		Doc("Поиск товаров").
		Operation("ProductByParent").Param(id))
	return ws
}

func (self *Product) ProductList(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity(self.GetAll())
}

func (self *Product) ProductById(req *restful.Request, resp *restful.Response) {
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
func (self *Product) ProductByParent(req *restful.Request, resp *restful.Response) {
	id,err := strconv.Atoi(req.PathParameter("id"))
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest, err.Error()))
		return
	}
	resp.WriteEntity(self.FindByParentId(int64(id)))
}

func (self *Product) CreateProduct(req *restful.Request, resp *restful.Response) {
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
	resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка создания товара"))
}

func (self *Product) UpdateProduct(req *restful.Request, resp *restful.Response) {
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
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка обновления товара"))
		return
	}
	if self.Create() {
		resp.WriteEntity(self)
		return
	}
	resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка создания товара"))
}
func (self *Product) DeleteProduct(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(&self)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	if self.Id != 0 {
		if self.Delete() {
			resp.WriteEntity(helpers.AnwerMessage(http.StatusOK, "Товар успешно удален"))
			return
		}
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка удаления товара"))
		return
	}
	resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest, "отсутствует номер товара"))
}


