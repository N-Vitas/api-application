package category

import (
	"github.com/emicklei/go-restful"
	"api-application/modules/session"
	"database/sql"
)

type Category struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Date int64 `json:"date"`
	sessMng *session.SessionManager
}

type CategorySql struct {
	Id sql.NullInt64
	Title sql.NullString
	Date sql.NullInt64
}

func NewCategoryService(sessMng *session.SessionManager) *restful.WebService {
	s := &Category{sessMng: sessMng}
	s.dbInit()
	return s.CategoryWebService()
}

func NewCategory() Category {
	s := Category{0,"",0,session.NewSession()}
	return s
}
func (u *Category) SetSession(sessMng *session.SessionManager) {
	u.sessMng = sessMng
}