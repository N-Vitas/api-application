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
	"api-application/modules/route"
	"api-application/modules/session"
	. "api-application/helpers"
)

var (
	props          *properties.Properties
	propertiesFile = flag.String("config", "config/app.properties", "the configuration file")
	JwtSecret   string
)

func main() {
	flag.Parse()
	route.ApiIcon = filepath.Join(route.SwaggerPath, "files/favicon.ico")

	// Загрузка конфигурации из файла
	Info("loading configuration from [%s]", *propertiesFile)
	var err error
	if props, err = properties.LoadFile(*propertiesFile, properties.UTF8); err != nil {
		log.Fatal("Unable to read properties:%v\n", err)
	}

	route.SwaggerPath = props.GetString("swagger.path", "")
	JwtSecret = props.GetString("jwt.secret", "")

	// New, shared session manager
	sessMng := session.NewSessionManager(props.FilterPrefix("database."))
	defer sessMng.CloseAll()

	// accept and respond in JSON unless told otherwise
	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	// gzip if accepted
	restful.DefaultContainer.EnableContentEncoding(true)
	// faster router
	restful.DefaultContainer.Router(restful.CurlyRouter{})
	// no need to access body more than once
	//restful.SetCacheReadEntity(false)

	apiCors := props.GetBool("http.server.cors", true)
	JwtToken := props.GetString("jwt.secret", "")
	app.Token = []byte(JwtToken)

	resurce := app.Register(sessMng, restful.DefaultContainer, apiCors)
	resurce.Conn = techquest.NewServer("/socket")

	addr := props.MustGet("http.server.host") + ":" + props.MustGet("http.server.port")
	//if os.Args[1]!="" {
	//	addr = props.MustGet("http.server.host") + ":" + os.Args[1]
	//}
	basePath := "http://" + addr

	// Serve favicon.ico
	http.HandleFunc(route.ApiIcon, route.Icon)

	tikets.NewTicketService(resurce,restful.DefaultContainer)
	pultForum.NewForumService(resurce,restful.DefaultContainer)
	images.NewImageService(resurce,restful.DefaultContainer)

	//http.Handle("/socket", websocket.Handler(ticketService.NewClient))

	go resurce.Conn.Listen()
	go resurce.Worker()

	fs := http.FileServer(http.Dir("./modules/techquest/webroot"))
	http.Handle("/webroot/", http.StripPrefix("/webroot/", fs))

	Info("ready to serve on %s", basePath)
	log.Fatal(http.ListenAndServe(addr, nil))

}
