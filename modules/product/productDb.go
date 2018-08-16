package product

import (
	"time"
	"database/sql"
	"log"
	"api-application/helpers"
	"api-application/modules/session"
)

func (self *Product) dbInit() {
	db := self.getDb()
	_, e := db.Exec("CREATE TABLE IF NOT EXISTS products (id INTEGER PRIMARY KEY AUTOINCREMENT, parentId NUMERIC, title TEXT UNIQUE, article Text, primePrice NUMERIC, price NUMERIC, status NUMERIC, date NUMERIC)")
	if e != nil {
		log.Panic(e)
	}
	helpers.Info("База товаров определена")
}

func (self *Product) getDb() *sql.DB {
	if self.sessMng == nil {
		self.sessMng = session.NewSession()
	}
	return self.sessMng.GetStorageDb()
}

func (self *Product) GetAll() []Product {
	db := self.getDb()
	s := []Product{}
	u := ProductSql{}
	rows, err := db.Query("SELECT id, parentId, title, article, primePrice, price, status, date FROM products")
	if err != nil {
		helpers.Info("Ошибка запроса списка товаров %v", err)
		return s
	}
	for rows.Next() {
		err = rows.Scan(&u.Id,&u.ParentId,&u.Title,&u.Article,&u.PrimePrice,&u.Price,&u.Status,&u.Date)
		if err != nil {
			helpers.Info("Ошибка сканирования списка товаров %v", err)
			continue
		}
		s = append(s, Product{Id:u.Id.Int64,ParentId:u.ParentId.Int64,Title:u.Title.String,Article:u.Article.String,PrimePrice:u.PrimePrice.Float64,Price:u.Price.Float64,Status:u.Status.Bool,Date:u.Date.Int64})
	}
	return s
}


func (self *Product) FindByParentId(key int64) []Product {
	db := self.getDb()
	r := []Product{}
	u := ProductSql{}
	str := "SELECT id, parentId, title, article, primePrice, price, status, date FROM products WHERE parentId = ?"
	rows, err := db.Query(str, key)
	if err != nil {
		helpers.Info("Ошибка запроса списка товаров %v", err)
		return r
	}
	for rows.Next() {
		err = rows.Scan(&u.Id, &u.ParentId, &u.Title, &u.Article, &u.PrimePrice, &u.Price, &u.Status, &u.Date)
		if err != nil {
			helpers.Info("Ошибка запроса списка товаров %v", err)
			continue
		}
		r = append(r, Product{Id:u.Id.Int64,ParentId:u.ParentId.Int64,Title:u.Title.String,Article:u.Article.String,PrimePrice:u.PrimePrice.Float64,Price:u.Price.Float64,Status:u.Status.Bool,Date:u.Date.Int64})
	}
	return r
}

func (u *Product) FindById(id int64) bool {
	db := u.getDb()
	str := "SELECT id, parentId, title, article, primePrice, price, status, date FROM products WHERE id = ?"
	err := db.QueryRow(str, id).Scan(&u.Id, &u.ParentId, &u.Title, &u.Article, &u.PrimePrice, &u.Price, &u.Status, &u.Date)
	if err != nil {
		helpers.Info("Ошибка запроса товара %v", err)
		return false
	}
	return true
}

func (u *Product) IsEmpty() bool {
	return len(u.Title) == 0 || u.Date == 0 || u.ParentId == 0 || len(u.Article) == 0 || u.PrimePrice == 0 || u.Price == 0
}

func (u *Product) Create() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	str := "INSERT INTO products( parentId, title, article, primePrice, price, status, date ) VALUES( ?, ?, ?, ?, ?, ?, ? )"
	res, err := db.Exec(str, u.ParentId, u.Title, u.Article, u.PrimePrice, u.Price, u.Status, time.Now().Unix())
	if err != nil {
		helpers.Info("Ошибка создания товара %v", err)
		return false
	}
	u.Id,_ = res.LastInsertId()
	return true
}

func (u *Product) Update() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	smtp := "UPDATE products SET parentId = ?, title = ?, article = ?, primePrice = ?, price = ?, status = ?, date  = ? WHERE id = ?"
	_, err := db.Exec(smtp, u.ParentId, u.Title, u.Article, u.PrimePrice, u.Price, u.Status, time.Now().Unix(),u.Id)
	if err != nil {
		helpers.Info("Ошибка обновления товара %v", err)
		return false
	}
	return true
}

func (u *Product) Delete() bool {
	db := u.getDb()
	if u.IsEmpty(){	return false }
	str := "DELETE FROM products WHERE id = ?"
	_, err := db.Exec(str, u.Id)
	if err != nil {
		helpers.Info("Ошибка удаления товара %v", err)
		return false
	}
	return true
}

