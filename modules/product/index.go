package product

import (
	"database/sql"
	"api-application/modules/session"
	"github.com/emicklei/go-restful"
	"time"
)

type Product struct {
	Id int64 `json:"id"`
	ParentId int64 `json:"parentId"`
	Title string `json:"title"`
	Article string `json:"article"`
	PrimePrice float64 `json:"primePrice"`
	Price float64 `json:"price"`
	Date int64 `json:"date"`
	Status bool `json:"status"`
	sessMng *session.SessionManager
}

type ProductSql struct {
	Id sql.NullInt64
	ParentId sql.NullInt64
	Title sql.NullString
	Date sql.NullInt64
	Article sql.NullString
	PrimePrice sql.NullFloat64
	Price sql.NullFloat64
	Status sql.NullBool
}

func NewProductService(sessMng *session.SessionManager) *restful.WebService {
	s := &Product{sessMng: sessMng, Status:false, Date: time.Now().Unix()}
	s.dbInit()
	return s.ProductWebService()
}

func NewProduct() Product {
	s := Product{sessMng: session.NewSession(),Status:false, Date: time.Now().Unix()}
	return s
}
func (u *Product) SetSession(sessMng *session.SessionManager) {
	u.sessMng = sessMng
}