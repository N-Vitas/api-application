package discount

import (
	"github.com/emicklei/go-restful"
	"api-application/modules/session"
	"database/sql"
	"time"
)

type Discount struct {
	Id        int64    `json:"id"`
	Title     string   `json:"title"`
	ProductId int64    `json:"productId"`
	Price     int64    `json:"price"`
	Percent   int64    `json:"percent"`
	Types     string   `json:"types"`
	List      []string `json:"list"`
	DateStart int64    `json:"dateStart"`
	DateStop  int64    `json:"dateStop"`
	sessMng   *session.SessionManager
}

type DiscountSql struct {
	Id        sql.NullInt64
	ProductId sql.NullInt64
	Price     sql.NullInt64
	Percent   sql.NullInt64
	DateStart sql.NullInt64
	DateStop  sql.NullInt64
	Title     sql.NullString
	Types     sql.NullString
	List      []string
}

func (u *Discount) Merge(s DiscountSql) Discount {
	n := NewDiscount()
	n.Id = s.Id.Int64
	n.ProductId = s.ProductId.Int64
	n.Price = s.Price.Int64
	n.Percent = s.Percent.Int64
	n.DateStart = s.DateStart.Int64
	n.DateStop = s.DateStop.Int64
	n.Title = s.Title.String
	n.Types = s.Types.String
	n.List = s.List
	return n
}

func NewDiscountService(sessMng *session.SessionManager) *restful.WebService {
	s := &Discount{sessMng: sessMng}
	s.dbInit()
	return s.DiscountWebService()
}

func NewDiscount() Discount {
	s := Discount{0,"",0,0,0,"percent",[]string{"percent","price"},time.Now().Unix(),0,session.NewSession()}
	return s
}
func (u *Discount) SetSession(sessMng *session.SessionManager) {
	u.sessMng = sessMng
}