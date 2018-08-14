package users

import (
	"database/sql"
	"log"
	"api-application/helpers"
	"time"
	"api-application/modules/session"
)
func (self *User) dbInit() {
	db := self.getDb()
	_, e := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, login UNIQUE, password TEXT, role TEXT, full_name TEXT, status NUMERIC, date NUMERIC)")
	if e != nil {
		log.Panic(e)
	}
	helpers.Info("База пользователей определена")

	u := User{Id: 1, Login: "root", Role: "admin", FullName: "Администратор", Status: 1, Date: time.Now().Unix()}
	u.Password = u.GetMD5Hash("123456789")

	if u.FindByName("root") {
		helpers.Info("Администратор %s уже существует", u.Login)
		return
	}
	_, err := db.Exec("INSERT INTO users( id, login, password, role, full_name, status, date ) VALUES( ?, ?, ?, ?, ?, ?, ?)", u.Id, u.Login, u.Password, u.Role, u.FullName, u.Status, u.Date)
	if err != nil {
		helpers.Info("База пользователей определена")
		return
	}
	helpers.Info("Администратор %s создан", u.Login)
}

func (self *User) getDb() *sql.DB {
	if self.sessMng == nil {
		self.sessMng = session.NewSession()
	}
	return self.sessMng.GetStorageDb()
}

func (self *User) GetAll() []User {
	db := self.getDb()
	users := []User{}
	rows, err := db.Query("SELECT id, login, password, role, full_name, status, date FROM users LIMIT 25")
	if err != nil {
		helpers.Info("Ошибка запроса списка пользователей %v", err)
		return users
	}
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.Id,&u.Login,&u.Password,&u.Role,&u.FullName,&u.Status,&u.Date)
		if err != nil {
			helpers.Info("Ошибка сканирования списка пользователей %v", err)
			continue
		}
		users = append(users, u)
	}
	return users
}


func (u *User) FindByName(login string) bool {
	db := u.getDb()
	str := "SELECT id, login, password, role, full_name, status, date FROM users WHERE login = ?"
	err := db.QueryRow(str, login).Scan(&u.Id, &u.Login, &u.Password, &u.Role, &u.FullName, &u.Status, &u.Date)
	if err != nil {
		helpers.Info("Ошибка запроса списка пользователей %v", err)
		return false
	}
	return true
}

func (u *User) FindById(id int64) bool {
	db := u.getDb()
	str := "SELECT id, login, password, role, full_name, status, date FROM users WHERE id = ?"
	err := db.QueryRow(str, id).Scan(&u.Id, &u.Login, &u.Password, &u.Role, &u.FullName, &u.Status, &u.Date)
	if err != nil {
		helpers.Info("Ошибка запроса списка пользователей %v", err)
		return false
	}
	return true
}
