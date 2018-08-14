package session

import (
	"database/sql"
	"sync"
	"gopkg.in/mgo.v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/magiconair/properties"
	"fmt"
	"strings"
	"errors"
	"log"
)

type SessionManager struct {
	configMap  map[string]*properties.Properties
	sessions   map[string]*sql.DB
	mongo MongoConnection
	accessLock *sync.RWMutex
}

type MongoConnection struct {
	Session *mgo.Session  // Соединение с сервером
	DB      *mgo.Database // Соединение с базой данных
}

func NewSessionManager(props *properties.Properties) *SessionManager {
	sess := &SessionManager{
		configMap:  make(map[string]*properties.Properties),
		sessions:   make(map[string]*sql.DB),
		accessLock: &sync.RWMutex{},
	}
	sess.SetConfig(props)
	return sess
}

func NewSession() *SessionManager {
	sess := &SessionManager{
		sessions:   make(map[string]*sql.DB),
	}
	return sess
}
// Переподключение или создание подключения к базе
func (self *SessionManager) GetStorageDb() *sql.DB {
	if self.sessions["storage"] != nil {
		return self.sessions["storage"]
	}
	db, err := sql.Open("sqlite3", "modules/session/application.db")
	if err != nil {
		info("Error Config %s", err)
		return nil
	}
	self.sessions["storage"] = db
	return self.sessions["storage"]
}

// Closes session based on `uri` or `host:port`
func (s *SessionManager) Close(sessionId string) {
	s.accessLock.Lock()
	if existing := s.sessions[sessionId]; existing != nil {
		existing.Close()
		delete(s.sessions, sessionId)
	}
	s.accessLock.Unlock()
}
// Closes all sessions.
func (s *SessionManager) CloseAll() {
	info("closing all sessions: ", len(s.sessions))
	s.accessLock.Lock()
	for _, each := range s.sessions {
		each.Close()
	}
	s.accessLock.Unlock()
}
// Returns slice containing all configured aliases
func (s *SessionManager) GetAliases() []string {
	aliases := []string{}
	for k := range s.configMap {
		aliases = append(aliases, k)
	}
	return aliases
}

func (s *SessionManager) Get(alias string) (*sql.DB, bool, error) {
	// Get alias configurations
	config, err := s.GetConfig(alias)
	if err != nil {
		info("Error Config %s", err)
		return nil, false, err
	}
	// var sessionId string
	sessionId := config.MustGet("host") + ":" + config.MustGet("port")
	// Check if session already exists
	s.accessLock.RLock()
	existing := s.sessions[sessionId]
	s.accessLock.RUnlock()

	// Clone and return if sessions exists
	if existing != nil {
		err = existing.Ping()
		if err != nil {
			existing = nil
		}else{
			//info("return connect to %s : %s",alias,sessionId)
			return existing, true, nil
		}
	}
	// Get timeout from configuration
	s.accessLock.Lock()

	var newSession *sql.DB
	var server string
	var user string
	var password string
	var port int
	var timeout int

	server = config.GetString("host","") // the database server
	user = config.GetString("username","") // the database user
	password = config.GetString("password","") // the database password
	port = config.GetInt("port",1433) // the database port
	timeout = config.GetInt("timeout",5) // the database port

	if alias == "remote" {
		connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;Connection Timeout=%d", server, user, password, port, timeout)
		newSession, err = sql.Open("mssql", connString)
	}else {
		connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, password, server, port, config.GetString("database","invo"))
		newSession, err = sql.Open("mysql", connString)//config.GetString("username","")+":"+config.GetString("password","")+"@tcp(127.0.0.1:"+config.GetInt("port",3306)+")/"+config.GetString("database","invo")+"?charset=utf8")
	}
	if err != nil {
		info("unable to connect to %s because : %v",alias, err.Error())
		newSession = nil
	}
	info("create new connect to %s : %s",alias,sessionId)
	s.sessions[sessionId] = newSession
	s.accessLock.Unlock()
	// s.sessions = newSession
	return newSession, true, err
}

func (s *SessionManager) SetConfig(props *properties.Properties) {
	for _, k := range props.Keys() {
		parts := strings.Split(k, ".")
		alias := parts[1]
		config := s.configMap[alias]
		if config == nil {
			config = properties.NewProperties()
			config.Set("alias", alias)
			s.configMap[alias] = config
		}
		config.Set(parts[2], props.MustGet(k))
	}
}

func (s *SessionManager) GetConfig(alias string) (*properties.Properties, error) {
	if config := s.configMap[alias]; config != nil {
		return config, nil
	}
	return nil, errors.New("Unknown alias: " + alias)
}
// Log wrapper
func info(template string, values ...interface{}) {
	log.Printf("[api] "+template+"\n", values...)
}

// Инициализация соединения с БД монго
func (s *SessionManager) GetMongoDB() (*mgo.Database,bool) {
	config, err := s.GetConfig("mongo")
	if err != nil {
		info("Error Config %s", err)
		return nil,false
	}
	var server string
	var database string

	server = config.GetString("host","")
	database = config.GetString("database","")
	existing := s.mongo.Session
	// Возвращаем существующую сессию mongo
	if existing != nil {
		err := existing.Ping()
		if err != nil {
			log.Println("Ping Mongo",err)
			existing = nil
		}else{
			//info("return connect to %s : %s",existing.DB(database))
			return s.mongo.DB.With(s.mongo.Session),true
		}
	}

	session, err := mgo.Dial(server)
	if err != nil {
		log.Println(server)
		log.Fatal("Ошибка подключения к базе",err)
	}
	s.mongo.Session = session
	db := session.DB(database)
	s.mongo.DB = db
	//s.accessLock.RUnlock()
	return s.mongo.DB,true
}
