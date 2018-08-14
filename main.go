package main

import (
	"flag"
	"log"
	"net/http"
	"api-application/modules/session"
	"api-application/modules/fileServer"
	"api-application/modules/test"
	. "api-application/modules/app"
	. "api-application/helpers"
	"api-application/modules/users"
)

func main() {
	flag.Parse()
	propertiesFile := flag.String("config", "./config/app.properties", "the configuration file")
	app := NewApp(propertiesFile)
	// Загрузка конфигурации из файла
	Info("Файл конфигурации загружен")
	// New, shared session manager
	app.AddSession(session.NewSessionManager(app.FilterPrefix("database.")))
	defer app.CloseSession()
	app.NewApi("public/favicon.ico")
	app.SetToken()
	app.Api.Course(app.GetSettingsBool("http.server.cors", true))
	app.Api.Filter(true)
	app.Api.AddService(test.NewTestService())
	app.Api.AddService(app.TokenWebService())
	app.Api.AddService(users.NewUserService(app.GetSession()))
	app.Api.Register()
	// Serve favicon.ico
	http.HandleFunc(app.Api.ApiIcon, app.Api.Icon)

	fileServer.NewFileServer("/","./public/www")
	Info("Сервер запущен на http://%s", app.GetAddress())
	log.Fatal(http.ListenAndServe(app.GetAddress(), nil))

}
