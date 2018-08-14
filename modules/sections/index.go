package sections

import (
	"github.com/emicklei/go-restful"
	"api-application/modules/session"
	"database/sql"
)

type Section struct {
	Id int64 `json:"id"`
	Key string `json:"key"`
	Title string `json:"title"`
	Value float64 `json:"value"`
	sessMng *session.SessionManager
}

type SectionSql struct {
	Id sql.NullInt64
	Key sql.NullString
	Title sql.NullString
	Value sql.NullFloat64
}

func NewSectionService(sessMng *session.SessionManager) *restful.WebService {
	s := &Section{sessMng: sessMng}
	s.dbInit()
	return s.SectionWebService()
}

func NewSection() Section {
	s := Section{1,"nds","Без НДС",0,session.NewSession()}
	return s
}
func (u *Section) SetSession(sessMng *session.SessionManager) {
	u.sessMng = sessMng
}