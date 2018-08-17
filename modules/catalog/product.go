package catalog

import (
	"api-application/modules/product"
	"api-application/modules/charge"
	"api-application/modules/discount"
)

type Product struct {
	product.Product
	Charge map[int64]charge.Charge `json:"charge"`
	Discount map[int64]discount.Discount `json:"discount"`
}

func NewProduct() Product {
	p := Product{Product: product.NewProduct()}
	p.NewCharges()
	p.NewDiscount()
	return p
}

func (c *Product) setProduct(p product.Product) bool {
	if p.IsEmpty() { return false }
	c.Product = p
	return true
}

func (p *Product) NewProducts() map[int64]Product {
	return make(map[int64]Product)
}

func (p *Product) NewDiscount() {
	p.Discount = make(map[int64]discount.Discount)
}

func (p *Product) NewCharges() {
	p.Charge = make(map[int64]charge.Charge)
}

func (c *Product) IsEmptyDiscountById(id int64) bool {
	empty := true
	if len(c.Discount) == 0 {
		return empty
	}
	if disk, ok := c.GetDiscountById(id); ok {
		return disk.IsEmpty()
	}
	return empty
}

func (c *Product) IsEmptyChargeById(id int64) bool {
	empty := true
	if len(c.Discount) == 0 {
		return empty
	}
	if char, ok := c.GetChargeById(id); ok {
		return char.IsEmpty()
	}
	return empty
}

func (c *Product) GetChargeById(id int64) (charge.Charge, bool) {
	cat, err := c.Charge[id]
	return cat, err
}
func (c *Product) GetDiscountById(id int64) (discount.Discount, bool) {
	cat, err := c.Discount[id]
	return cat, err
}

func (c *Product) AddCharge(char charge.Charge) bool {
	if char.Id == 0 || char.IsEmpty() { return false }
	c.Charge[char.Id] = char
	return true
}
func (c *Product) AddDiscount(d discount.Discount) bool {
	if d.Id == 0 || d.IsEmpty() { return false }
	c.Discount[d.Id] = d
	return true
}