package charge

import (
"github.com/emicklei/go-restful"
"strconv"
"net/http"
"api-application/helpers"
"encoding/json"
)

func (self *Charge) ChargeWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/charges")
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(self.ChargeList).
		Doc("Список наценки").
		Operation("ChargeList"))
	ws.Route(ws.POST("/").To(self.CreateCharge).
		Doc("Создание наценки").
		Operation("CreateCharge"))
	ws.Route(ws.PUT("/").To(self.UpdateCharge).
		Doc("Обновление наценки").
		Operation("UpdateCharge"))
	ws.Route(ws.DELETE("/").To(self.DeleteCharge).
		Doc("Удаление наценки").
		Operation("DeleteCharge"))
	id := ws.PathParameter("id", "ID!").DataType("int")
	ws.Route(ws.GET("/{id}").To(self.ChargeById).
		Doc("Поиск наценки").
		Operation("ChargeById").Param(id))
	return ws
}

func (self *Charge) ChargeList(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity(self.GetAll())
}

func (self *Charge) ChargeById(req *restful.Request, resp *restful.Response) {
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

func (self *Charge) CreateCharge(req *restful.Request, resp *restful.Response) {
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
	resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка создания наценки"))
}

func (self *Charge) UpdateCharge(req *restful.Request, resp *restful.Response) {
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
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка обновления наценки"))
		return
	}
	if self.Create() {
		resp.WriteEntity(self)
		return
	}
	resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка создания наценки"))
}
func (self *Charge) DeleteCharge(req *restful.Request, resp *restful.Response) {
	decoder := json.NewDecoder(req.Request.Body)
	err := decoder.Decode(&self)
	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest,err.Error()))
		return
	}
	if self.Id != 0 {
		if self.Delete() {
			resp.WriteEntity(helpers.AnwerMessage(http.StatusOK, "Удаление наценки выполненно успешно"))
			return
		}
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, helpers.AnwerMessage(http.StatusInternalServerError, "Ошибка удаления наценки"))
		return
	}
	resp.WriteHeaderAndEntity(http.StatusBadRequest, helpers.AnwerMessage(http.StatusBadRequest, "отсутствует номер наценки"))
}


