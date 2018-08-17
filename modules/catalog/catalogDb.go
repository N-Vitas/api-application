package catalog

import (
	"database/sql"
	"api-application/modules/session"
	"api-application/modules/category"
	"api-application/modules/discount"
	"api-application/modules/charge"
)

func (self *Сatalog) getDb() *sql.DB {
	if self.sessMng == nil {
		self.sessMng = session.NewSession()
	}
	return self.sessMng.GetStorageDb()
}

func (self *Сatalog) GetAll() []Category {
	r := []Category{}
	s := category.NewCategory()
	d := discount.NewDiscount()
	ch := charge.NewCharge()
	for _,cat := range s.GetAll(){
		c := NewCategory()
		pr := NewProduct()
		for _, prod := range pr.FindByParentId(cat.Id) {
			pr.setProduct(prod)
			disks := d.FindByProducts(prod.Id)
			for _, disk := range disks {
				pr.AddDiscount(disk)
			}
			chars := ch.FindByProducts(prod.Id)
			for _, char := range chars {
				pr.AddCharge(char)
			}
			c.AddProduct(pr)
		}
		c.setCategory(cat)
		if !c.IsEmpty() {
			r = append(r, c)
			self.Category[cat.Id] = c
		}
	}
	if self.IsEmpty() {
		return []Category{}
	}
	return r
}

func (self *Сatalog) FindByProducts(productId int64) []Product {
	s := []Product{}
	//p,_ := self.G
	//for _,cat := range self.Product{
	//	s = append(s,cat)
	//}
	return s
}

func (u *Сatalog) FindByTitle(title string) bool {
	//db := u.getDb()
	//str := "SELECT id,productId,title,price,percent,types,list,dateStart,dateStop FROM charges WHERE title = ?"
	//err := db.QueryRow(str, title).Scan(&u.Id,&u.ProductId,&u.Title,&u.Price,&u.Percent,&u.Types,&u.List,&u.DateStart,&u.DateStop)
	//if err != nil {
	//	helpers.Info("Ошибка запроса наценки %v", err)
	//	return false
	//}
	return true
}

func (u *Сatalog) FindById(id int64) bool {
	//db := u.getDb()
	//str := "SELECT id,productId,title,price,percent,types,list,dateStart,dateStop FROM charges WHERE id = ?"
	//err := db.QueryRow(str, id).Scan(&u.Id,&u.ProductId,&u.Title,&u.Price,&u.Percent,&u.Types,&u.List,&u.DateStart,&u.DateStop)
	//if err != nil {
	//	helpers.Info("Ошибка запроса наценки %v", err)
	//	return false
	//}
	return true
}

func (u *Сatalog) Create() bool {
	//db := u.getDb()
	//if u.IsEmpty(){	return false }
	//str := "INSERT INTO charges( productId,title,price,percent,types,list,dateStart,dateStop ) VALUES( ?, ?, ?, ?, ?, ?, ?, ? )"
	//res, err := db.Exec(str, u.Title, time.Now().Unix())
	//if err != nil {
	//	helpers.Info("Ошибка создания наценки %v", err)
	//	return false
	//}
	//u.Id,_ = res.LastInsertId()
	return true
}

func (u *Сatalog) Update() bool {
	//db := u.getDb()
	//if u.IsEmpty(){	return false }
	//
	//smtp := "UPDATE charges SET productId = ?, title = ?, price = ?, percent = ?, types = ?, list = ?, dateStart = ?, dateStop = ? WHERE id = ?"
	//
	//res, err := db.Exec(smtp,u.ProductId,u.Title,u.Price,u.Percent,u.Types,u.List,u.DateStart,u.DateStop,u.Id)
	//if err != nil {
	//	helpers.Info("Ошибка обновления наценки %v", err)
	//	return false
	//}
	//u.Id,_ = res.LastInsertId()
	return true
}

func (u *Сatalog) Delete() bool {
	//db := u.getDb()
	//if u.IsEmpty(){	return false }
	//str := "DELETE FROM charges WHERE id = ?"
	//_, err := db.Exec(str, u.Id)
	//if err != nil {
	//	helpers.Info("Ошибка удаления наценки %v", err)
	//	return false
	//}
	return true
}
