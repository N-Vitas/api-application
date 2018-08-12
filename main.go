package api_application

import (
	"github.com/magiconair/properties"
	"flag"
	"path/filepath"
	"log"
	"github.com/emicklei/go-restful"
	"api-go-herobrine/app"
	"api-go-herobrine/modules/techquest"
	"net/http"
	"api-go-herobrine/modules/tikets"
	"api-go-herobrine/modules/pultForum"
	"api-go-herobrine/modules/images"
	"api-application/modules/session"
	"api-application/modules/fileServer"
	. "api-application/modules/app"
	. "api-application/helpers"
)

func main() {
	flag.Parse()
	app := NewApp("config", "config/app.properties", "the configuration file")
	// Загрузка конфигурации из файла
	Info("loading configuration")
	// New, shared session manager
	app.AddSession(session.NewSessionManager(app.FilterPrefix("database."))))
	defer app.CloseSession()
	app.NewApi("public/favicon.ico")
	app.Api.Course(app.GetSettingsBool("http.server.cors", true))
	app.Api.Filter(true)
	app.Api.Register()

	resurce := app.Register(sessMng, restful.DefaultContainer, apiCors)
	resurce.Conn = techquest.NewServer("/socket")

	basePath := "http://" + app.GetAddress()

	// Serve favicon.ico
	http.HandleFunc(app.Api.ApiIcon, app.Api.Icon)

	tikets.NewTicketService(resurce,restful.DefaultContainer)
	pultForum.NewForumService(resurce,restful.DefaultContainer)
	images.NewImageService(resurce,restful.DefaultContainer)

	//http.Handle("/socket", websocket.Handler(ticketService.NewClient))

	go resurce.Conn.Listen()
	go resurce.Worker()

	fileServer.NewFileServer("/","./public/www")
	Info("ready to serve on %s", basePath)
	log.Fatal(http.ListenAndServe(app.GetAddress(), nil))

}
