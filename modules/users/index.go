package users

import (
	"github.com/emicklei/go-restful"
	"api-application/modules/session"
	"crypto/md5"
	"encoding/hex"
	"database/sql"
)

type User struct {
	Id int64 `json:"id"`
	Login string `json:"login"`
	Password string `json:"password"`
	Role string `json:"role"`
	FullName string `json:"full_name"`
	Status int64 `json:"status"`
	Date int64 `json:"date"`
	sessMng *session.SessionManager
}

type UserSql struct {
	Id sql.NullInt64
	Status sql.NullInt64
	Date sql.NullInt64
	Login sql.NullString
	Password sql.NullString
	Role sql.NullString
	FullName sql.NullString
}

func NewUserService(sessMng *session.SessionManager) *restful.WebService {
	u := &User{sessMng: sessMng}
	u.dbInit()
	return u.UserWebService()
}

func NewUser() *User {
	u := &User{}
	return u
}
func (u *User) SetSession(sessMng *session.SessionManager) {
	u.sessMng = sessMng
}

// Хеширование строки
func (s *User) GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}