package catalog

import (
	"api-application/modules/category"
)

type Category struct {
	category.Category
	Product map[int64]Product `json:"product"`
}

func NewCategory() Category {
	p := NewProduct()
	c := Category{
		Category: category.NewCategory(),
		Product: p.NewProducts(),
	}
	return c
}

func (c *Category) setCategory(cat category.Category) bool {
	if cat.IsEmpty() || len (c.Product) == 0 { return false }
	c.Category = cat
	return true
}

func (c *Category) AddProduct(p Product) bool {
	if p.IsEmpty() { return false }
	c.Product[p.Id] = p
	return true
}

func (c *Category) GetById(id int64) (Product, bool) {
	cat, ok := c.Product[id]
	return cat, ok
}