package category

import (
	"database/sql"
	"log"
	"api-application/helpers"
	"api-application/modules/session"
	"time"
)
func (self *Category) dbInit() {
	db := self.getDb()
	_, e := db.Exec("CREATE TABLE IF NOT EXISTS categories (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT UNIQUE, date NUMERIC)")
	if e != nil {
		log.Panic(e)
	}
	helpers.Info("База категорий определена")
}

func (self *Category) getDb() *sql.DB {
	if self.sessMng == nil {
		self.sessMng = session.NewSession()
	}
	return self.sessMng.GetStorageDb()
}

func (self *Category) GetAll() []Category {
	db := self.getDb()
	s := []Category{}
	u := CategorySql{}
	rows, err := db.Query("SELECT id, title, date FROM categories")
	if err != nil {
		helpers.Info("Ошибка запроса списка категорий %v", err)
		return s
	}
	for rows.Next() {
		err = rows.Scan(&u.Id,&u.Title,&u.Date)
		if err != nil {
			helpers.Info("Ошибка сканирования списка категорий %v", err)
			continue
		}
		s = append(s, Category{Id:u.Id.Int64,Title:u.Title.String,Date:u.Date.Int64})
	}
	return s
}


func (u *Category) FindByKey(key string) bool {
	db := u.getDb()
	str := "SELECT id, title, date FROM categories WHERE keys = ?"
	err := db.QueryRow(str, key).Scan(&u.Id, &u.Title, &u.Date)
	if err != nil {
		helpers.Info("Ошибка запроса списка категорий %v", err)
		return false
	}
	return true
}

func (u *Category) FindById(id int64) bool {
	db := u.getDb()
	str := "SELECT id, title, date FROM categories WHERE id = ?"
	err := db.QueryRow(str, id).Scan(&u.Id, &u.Title, &u.Date)
	if err != nil {
		helpers.Info("Ошибка запроса списка категорий %v", err)
		return false
	}
	return true
}

func (u *Category) IsEmpty() bool {
	return len(u.Title) == 0 && u.Date == 0
}

func (u *Category) Create() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	str := "INSERT INTO categories( title, date ) VALUES( ?, ? )"
	res, err := db.Exec(str, u.Title, time.Now().Unix())
	if err != nil {
		helpers.Info("Ошибка создания категорий %v", err)
		return false
	}
	u.Id,_ = res.LastInsertId()
	return true
}

func (u *Category) Update() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	smtp := "UPDATE categories SET title=?, date=? WHERE id = ?"
	res, err := db.Exec(smtp, u.Title, u.Date, u.Id)
	if err != nil {
		helpers.Info("Ошибка обновления категорий %v", err)
		return false
	}
	u.Id,_ = res.LastInsertId()
	return true
}

func (u *Category) Delete() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	str := "DELETE FROM categories WHERE id = ?"
	_, err := db.Exec(str, u.Id)
	if err != nil {
		helpers.Info("Ошибка удаления категорий %v", err)
		return false
	}
	return true
}
