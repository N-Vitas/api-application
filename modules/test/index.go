package test

import "github.com/emicklei/go-restful"

type Test struct {
	Status int `json:"status"`
	Message string `json:"message"`
}
func NewTestService() *restful.WebService{
	t := Test{}
	return t.TicketWebService()
}

func (self *Test) TicketWebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/tests")
	ws.Consumes("*/*")
	ws.Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(self.NotFound).
		Doc("NotFound").
		Operation("NotFound"))
	return ws
}

func (self *Test) NotFound(req *restful.Request, resp *restful.Response) {
	self.Status = 200
	self.Message = "Hello World"
	resp.WriteEntity(self)
}
