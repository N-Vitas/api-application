package api

import "github.com/emicklei/go-restful"

func (s *Api) AddService(services *restful.WebService)  {
	s.services = append(s.services,)
}

func (s *Api) Register()  {
	for _, service := range s.services {
		s.container.Add(service)
	}
}
