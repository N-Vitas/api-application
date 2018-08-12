package api

import (
	"net/http"
	"github.com/emicklei/go-restful"
)

type Api struct {
	container   *restful.Container
	services    []*restful.WebService
	SwaggerPath string
	ApiIcon     string
	Token       []byte
}

func NewApi() *Api {
	api := &Api{}
	// accept and respond in JSON unless told otherwise
	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	// gzip if accepted
	restful.DefaultContainer.EnableContentEncoding(true)
	// faster router
	restful.DefaultContainer.Router(restful.CurlyRouter{})
	// no need to access body more than once
	//restful.SetCacheReadEntity(false)
	api.container = restful.DefaultContainer
	return api
}
// If swagger is not on `/` redirect to it
func (s *Api) Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, s.SwaggerPath, http.StatusMovedPermanently)
}

func (s *Api) Icon(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, s.ApiIcon, http.StatusMovedPermanently)
}
