package discount

import (
"github.com/emicklei/go-restful"
"strconv"
"net/http"
"api-application/helpers"
"encoding/json"
)

func (self *Discount) DiscountWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/discounts")
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(self.DiscountList).
		Doc("Список скидки").
		Operation("DiscountList"))
	ws.Route(ws.POST("/").To(self.CreateDiscount).
		Doc("Создание скидки").
		Operation("CreateDiscount"))
	ws.Route(ws.PUT("/").To(self.UpdateDiscount).
		Doc("Обновление скидки").
		Operation("UpdateDiscount"))
	ws.Route(ws.DELETE("/").To(self.DeleteDiscount).
		Doc("Удаление скидки").
		Operation("DeleteDiscount"))
	id := ws.PathParameter("id", "ID!").DataType("int")
	ws.Route(ws.GET("/{id}").To(self.DiscountById).
		Doc("Поиск скидки").
		Operation("DiscountById").Param(id))
	return ws
}

func (self *Discount) DiscountList(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity(self.GetAll())
}

func (self *Discount) DiscountById(req *restful.Request, resp *restful.Response) {
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

func (self *Discount) CreateDiscount(req *restful.Request, resp *restful.Response) {
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
	resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка создания скидки"))
}

func (self *Discount) UpdateDiscount(req *restful.Request, resp *restful.Response) {
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
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка обновления скидки"))
		return
	}
	if self.Create() {
		resp.WriteEntity(self)
		return
	}
	resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка создания скидки"))
}
func (self *Discount) DeleteDiscount(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(&self)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	if self.Id != 0 {
		if self.Delete() {
			resp.WriteEntity(helpers.AnwerMessage(http.StatusOK, "Удаление скидки выполненно успешно"))
			return
		}
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка удаления скидки"))
		return
	}
	resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest, "отсутствует номер скидки"))
}


