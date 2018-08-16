package sections

import (
	"database/sql"
	"log"
	"api-application/helpers"
	"api-application/modules/session"
)
func (self *Section) dbInit() {
	db := self.getDb()
	_, e := db.Exec("CREATE TABLE IF NOT EXISTS sections (id INTEGER PRIMARY KEY AUTOINCREMENT, keys TEXT UNIQUE, title TEXT, percent NUMERIC)")
	if e != nil {
		log.Panic(e)
	}
	helpers.Info("База секций определена")

	u := NewSection()

	if u.FindByKey("nds") {
		helpers.Info("Секция %s уже существует", u.Title)
		return
	}
	if u.Create() {
		helpers.Info("Секция %s создана", u.Title)
		return
	}
}

func (self *Section) getDb() *sql.DB {
	if self.sessMng == nil {
		self.sessMng = session.NewSession()
	}
	return self.sessMng.GetStorageDb()
}

func (self *Section) GetAll() []Section {
	db := self.getDb()
	s := []Section{}
	u := SectionSql{}
	rows, err := db.Query("SELECT id, keys, title, percent FROM sections")
	if err != nil {
		helpers.Info("Ошибка запроса списка секций %v", err)
		return s
	}
	for rows.Next() {
		err = rows.Scan(&u.Id,&u.Key,&u.Title,&u.Value)
		if err != nil {
			helpers.Info("Ошибка сканирования списка секций %v", err)
			continue
		}
		s = append(s, Section{Id:u.Id.Int64,Key:u.Key.String,Title:u.Title.String,Value:u.Value.Float64})
	}
	return s
}


func (u *Section) FindByKey(key string) bool {
	db := u.getDb()
	str := "SELECT id, keys, title, percent FROM sections WHERE keys = ?"
	err := db.QueryRow(str, key).Scan(&u.Id, &u.Key, &u.Title, &u.Value)
	if err != nil {
		helpers.Info("Ошибка запроса списка секций %v", err)
		return false
	}
	return true
}

func (u *Section) FindById(id int64) bool {
	db := u.getDb()
	str := "SELECT id, keys, title, percent FROM sections WHERE id = ?"
	err := db.QueryRow(str, id).Scan(&u.Id, &u.Key, &u.Title, &u.Value)
	if err != nil {
		helpers.Info("Ошибка запроса списка секций %v", err)
		return false
	}
	return true
}

func (u *Section) IsEmpty() bool {
	return len(u.Key) == 0 || len(u.Title) == 0
}

func (u *Section) Create() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	str := "INSERT INTO sections( keys, title, percent ) VALUES( ?, ?, ? )"
	res, err := db.Exec(str, u.Key, u.Title, u.Value)
	if err != nil {
		helpers.Info("Ошибка создания секций %v", err)
		return false
	}
	u.Id,_ = res.LastInsertId()
	return true
}

func (u *Section) Update() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	smtp := "UPDATE sections SET keys=?, title=?, percent=? WHERE id = ?"
	res, err := db.Exec(smtp, u.Key, u.Title, u.Value, u.Id)
	if err != nil {
		helpers.Info("Ошибка обновления секций %v", err)
		return false
	}
	u.Id,_ = res.LastInsertId()
	return true
}

func (u *Section) Delete() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	str := "DELETE FROM sections WHERE id = ?"
	_, err := db.Exec(str, u.Id)
	if err != nil {
		helpers.Info("Ошибка удаления секций %v", err)
		return false
	}
	return true
}
