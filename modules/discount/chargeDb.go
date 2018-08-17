package discount

import (
	"database/sql"
	"log"
	"api-application/helpers"
	"api-application/modules/session"
	"time"
)
func (self *Discount) dbInit() {
	db := self.getDb()
	pr := `CREATE TABLE IF NOT EXISTS charges (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		productId INTEGER,
		title TEXT UNIQUE,
		price NUMERIC,
		percent NUMERIC,
		types TEXT,
		list TEXT,
		dateStart NUMERIC,
		dateStop NUMERIC
	)`
	_, e := db.Exec(pr)
	if e != nil {
		log.Panic(e)
	}
	helpers.Info("База скидки определена")
}

func (self *Discount) getDb() *sql.DB {
	if self.sessMng == nil {
		self.sessMng = session.NewSession()
	}
	return self.sessMng.GetStorageDb()
}

func (self *Discount) GetAll() []Discount {
	db := self.getDb()
	s := []Discount{}
	u := DiscountSql{}
	rows, err := db.Query("SELECT id,productId,title,price,percent,types,list,dateStart,dateStop FROM charges")
	if err != nil {
		helpers.Info("Ошибка запроса списка скидки %v", err)
		return s
	}
	for rows.Next() {
		err = rows.Scan(&u.Id,&u.ProductId,&u.Title,&u.Price,&u.Percent,&u.Types,&u.List,&u.DateStart,&u.DateStop)
		if err != nil {
			helpers.Info("Ошибка сканирования списка скидки %v", err)
			continue
		}
		s = append(s, self.Merge(u))
	}
	return s
}

func (self *Discount) FindByProducts(productId int64) []Discount {
	db := self.getDb()
	s := []Discount{}
	u := DiscountSql{}
	rows, err := db.Query("SELECT id,productId,title,price,percent,types,list,dateStart,dateStop FROM charges WHERE productId = ?", productId)
	if err != nil {
		helpers.Info("Ошибка запроса списка скидки %v", err)
		return s
	}
	for rows.Next() {
		err = rows.Scan(&u.Id,&u.ProductId,&u.Title,&u.Price,&u.Percent,&u.Types,&u.List,&u.DateStart,&u.DateStop)
		if err != nil {
			helpers.Info("Ошибка сканирования списка скидки %v", err)
			continue
		}
		s = append(s, self.Merge(u))
	}
	return s
}

func (u *Discount) FindByTitle(title string) bool {
	db := u.getDb()
	str := "SELECT id,productId,title,price,percent,types,list,dateStart,dateStop FROM charges WHERE title = ?"
	err := db.QueryRow(str, title).Scan(&u.Id,&u.ProductId,&u.Title,&u.Price,&u.Percent,&u.Types,&u.List,&u.DateStart,&u.DateStop)
	if err != nil {
		helpers.Info("Ошибка запроса скидки %v", err)
		return false
	}
	return true
}

func (u *Discount) FindById(id int64) bool {
	db := u.getDb()
	str := "SELECT id,productId,title,price,percent,types,list,dateStart,dateStop FROM charges WHERE id = ?"
	err := db.QueryRow(str, id).Scan(&u.Id,&u.ProductId,&u.Title,&u.Price,&u.Percent,&u.Types,&u.List,&u.DateStart,&u.DateStop)
	if err != nil {
		helpers.Info("Ошибка запроса скидки %v", err)
		return false
	}
	return true
}

func (u *Discount) IsEmpty() bool {
	return len(u.Title) == 0 || u.Price == 0 || u.Percent == 0 || len(u.Types) == 0 || len(u.List) == 0 || u.DateStart == 0
}

func (u *Discount) Create() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	str := "INSERT INTO charges( productId,title,price,percent,types,list,dateStart,dateStop ) VALUES( ?, ?, ?, ?, ?, ?, ?, ? )"
	res, err := db.Exec(str, u.Title, time.Now().Unix())
	if err != nil {
		helpers.Info("Ошибка создания скидки %v", err)
		return false
	}
	u.Id,_ = res.LastInsertId()
	return true
}

func (u *Discount) Update() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }

	smtp := "UPDATE charges SET productId = ?, title = ?, price = ?, percent = ?, types = ?, list = ?, dateStart = ?, dateStop = ? WHERE id = ?"

	res, err := db.Exec(smtp,u.ProductId,u.Title,u.Price,u.Percent,u.Types,u.List,u.DateStart,u.DateStop,u.Id)
	if err != nil {
		helpers.Info("Ошибка обновления скидки %v", err)
		return false
	}
	u.Id,_ = res.LastInsertId()
	return true
}

func (u *Discount) Delete() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	str := "DELETE FROM charges WHERE id = ?"
	_, err := db.Exec(str, u.Id)
	if err != nil {
		helpers.Info("Ошибка удаления скидки %v", err)
		return false
	}
	return true
}
