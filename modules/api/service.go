package api

import (
	"github.com/emicklei/go-restful"
)

func (s *Api) AddService(service *restful.WebService)  {
	s.services = append(s.services, service)
}

func (s *Api) Register()  {
	for _, service := range s.services {
		s.container.Add(service)
	}
}
