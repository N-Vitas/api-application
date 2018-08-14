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
	_, e := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, login TEXT UNIQUE, password TEXT, role TEXT, full_name TEXT, status NUMERIC, date NUMERIC)")
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
	if u.Create() {
		helpers.Info("Администратор %s создан", u.Login)
		return
	}
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

func (u *User) IsEmpty() bool {
	return len(u.FullName) == 0 && len(u.Login) == 0 && len(u.Password) == 0 && len(u.Role) == 0
}

func (u *User) Create() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	str := "INSERT INTO users( login, password, role, full_name, status, date ) VALUES( ?, ?, ?, ?, ?, ?)"
	u.Status = 1
	u.Date = time.Now().Unix()
	res, err := db.Exec(str, u.Login, u.Password, u.Role, u.FullName, u.Status, u.Date)
	if err != nil {
		helpers.Info("Ошибка создания пользователя %v", err)
		return false
	}
	u.Id,_ = res.LastInsertId()
	return true
}

func (u *User) Update() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	smtp := "UPDATE users SET login=?, password=?, role=?, full_name=?, status=?, date=? WHERE id = ?"
	res, err := db.Exec(smtp, u.Login, u.Password, u.Role, u.FullName, u.Status, u.Date, u.Id)
	if err != nil {
		helpers.Info("Ошибка обновления пользователя %v", err)
		return false
	}
	u.Id,_ = res.LastInsertId()
	return true
}

func (u *User) Delete() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	str := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(str, u.Id)
	if err != nil {
		helpers.Info("Ошибка удаления пользователя %v", err)
		return false
	}
	return true
}