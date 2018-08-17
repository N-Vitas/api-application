package charge

import (
	"database/sql"
	"log"
	"api-application/helpers"
	"api-application/modules/session"
	"time"
)
func (self *Charge) dbInit() {
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
	helpers.Info("База наценки определена")
}

func (self *Charge) getDb() *sql.DB {
	if self.sessMng == nil {
		self.sessMng = session.NewSession()
	}
	return self.sessMng.GetStorageDb()
}

func (self *Charge) GetAll() []Charge {
	db := self.getDb()
	s := []Charge{}
	u := ChargeSql{}
	rows, err := db.Query("SELECT id,productId,title,price,percent,types,list,dateStart,dateStop FROM charges")
	if err != nil {
		helpers.Info("Ошибка запроса списка наценки %v", err)
		return s
	}
	for rows.Next() {
		err = rows.Scan(&u.Id,&u.ProductId,&u.Title,&u.Price,&u.Percent,&u.Types,&u.List,&u.DateStart,&u.DateStop)
		if err != nil {
			helpers.Info("Ошибка сканирования списка наценки %v", err)
			continue
		}
		s = append(s, self.Merge(u))
	}
	return s
}

func (self *Charge) FindByProducts(productId int64) []Charge {
	db := self.getDb()
	s := []Charge{}
	u := ChargeSql{}
	rows, err := db.Query("SELECT id,productId,title,price,percent,types,list,dateStart,dateStop FROM charges WHERE productId = ?", productId)
	if err != nil {
		helpers.Info("Ошибка запроса списка наценки %v", err)
		return s
	}
	for rows.Next() {
		err = rows.Scan(&u.Id,&u.ProductId,&u.Title,&u.Price,&u.Percent,&u.Types,&u.List,&u.DateStart,&u.DateStop)
		if err != nil {
			helpers.Info("Ошибка сканирования списка наценки %v", err)
			continue
		}
		s = append(s, self.Merge(u))
	}
	return s
}

func (u *Charge) FindByTitle(title string) bool {
	db := u.getDb()
	str := "SELECT id,productId,title,price,percent,types,list,dateStart,dateStop FROM charges WHERE title = ?"
	err := db.QueryRow(str, title).Scan(&u.Id,&u.ProductId,&u.Title,&u.Price,&u.Percent,&u.Types,&u.List,&u.DateStart,&u.DateStop)
	if err != nil {
		helpers.Info("Ошибка запроса наценки %v", err)
		return false
	}
	return true
}

func (u *Charge) FindById(id int64) bool {
	db := u.getDb()
	str := "SELECT id,productId,title,price,percent,types,list,dateStart,dateStop FROM charges WHERE id = ?"
	err := db.QueryRow(str, id).Scan(&u.Id,&u.ProductId,&u.Title,&u.Price,&u.Percent,&u.Types,&u.List,&u.DateStart,&u.DateStop)
	if err != nil {
		helpers.Info("Ошибка запроса наценки %v", err)
		return false
	}
	return true
}

func (u *Charge) IsEmpty() bool {
	return len(u.Title) == 0 || u.Price == 0 || u.Percent == 0 || len(u.Types) == 0 || len(u.List) == 0 || u.DateStart == 0
}

func (u *Charge) Create() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	str := "INSERT INTO charges( productId,title,price,percent,types,list,dateStart,dateStop ) VALUES( ?, ?, ?, ?, ?, ?, ?, ? )"
	res, err := db.Exec(str, u.Title, time.Now().Unix())
	if err != nil {
		helpers.Info("Ошибка создания наценки %v", err)
		return false
	}
	u.Id,_ = res.LastInsertId()
	return true
}

func (u *Charge) Update() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }

	smtp := "UPDATE charges SET productId = ?, title = ?, price = ?, percent = ?, types = ?, list = ?, dateStart = ?, dateStop = ? WHERE id = ?"

	res, err := db.Exec(smtp,u.ProductId,u.Title,u.Price,u.Percent,u.Types,u.List,u.DateStart,u.DateStop,u.Id)
	if err != nil {
		helpers.Info("Ошибка обновления наценки %v", err)
		return false
	}
	u.Id,_ = res.LastInsertId()
	return true
}

func (u *Charge) Delete() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	str := "DELETE FROM charges WHERE id = ?"
	_, err := db.Exec(str, u.Id)
	if err != nil {
		helpers.Info("Ошибка удаления наценки %v", err)
		return false
	}
	return true
}
