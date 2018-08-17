package catalog

import (
	"github.com/emicklei/go-restful"
	"api-application/modules/session"
	"api-application/modules/charge"
	"api-application/modules/discount"
)

type Сatalog struct {
	Category map[int64]Category `json:"category"`
	sessMng   *session.SessionManager
}


func (c *Сatalog) IsEmpty() bool {
	return len(c.Category) == 0
}

func (c *Сatalog) IsEmptyById(id int64) bool {
	empty := true
	if len(c.Category) == 0 {
		return empty
	}
	if prod, ok := c.GetById(id); ok {
		return prod.IsEmpty()
	}
	return empty
}

func (c *Сatalog) GetById(id int64) (Category, bool) {
	cat, err := c.Category[id]
	return cat, err
}
//func (c *Сatalog) Merge(ct category.Category, p product.Product, cr charge.Charge, d discount.Discount) bool {
//
//	p.AddProduct(p)
//	ct.AddProduct(p)
//	cat, ok := c.GetById(ct.Id)
//	if ok {
//	}
//	c.Category = ct
//	c.Product = p
//	c.Charge = cr
//	c.Discount = d
//	return c.IsEmpty()
//}

func NewСatalogService(sessMng *session.SessionManager) *restful.WebService {
	s := NewСatalog()
	s.SetSession(sessMng)
	return s.СatalogWebService()
}

func NewСatalog() *Сatalog {
	cr := make(map[int64]charge.Charge)
	d := make(map[int64]discount.Discount)
	p := make(map[int64]Product)
	c := make(map[int64]Category)
	cr[0] = charge.NewCharge()
	d[0] = discount.NewDiscount()
	p[0] = Product{Charge:cr, Discount:d}
	c[0] = Category{Product:p}
	s := &Сatalog{Category:c}
	return s
}
func (u *Сatalog) SetSession(sessMng *session.SessionManager) {
	u.sessMng = sessMng
}