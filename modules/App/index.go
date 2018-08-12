package App

import (
	"github.com/magiconair/properties"
	"api-application/modules/session"
	"api-application/modules/api"
	"flag"
	"log"
	"path/filepath"
	"fmt"
)

type App struct {
	sessMng        *session.SessionManager
	props          *properties.Properties
	propertiesFile *string
	jwtSecret      string
	Api            *api.Api
}

func NewApp(name string, value string, usage string) *App {
	flag.Parse()
	app := &App{}
	app.propertiesFile = flag.String(name, value, usage)
	var err error
	if app.props, err = properties.LoadFile(*app.propertiesFile, properties.UTF8); err != nil {
		log.Fatal("Unable to read properties:%v\n", err)
	}
	app.jwtSecret = app.GetSettingsString("jwt.secret", "")
	app.Api.Token = []byte(app.jwtSecret)
	return app
}
func (s *App) NewApi(icon string){
	s.Api = api.NewApi()
	s.Api.ApiIcon = filepath.Join(s.Api.SwaggerPath, icon)
	s.Api.SwaggerPath = s.GetSettingsString("swagger.path", "")
}

func (s *App) GetSettingsString (settings string, def string) string {
	return s.props.GetString(settings, def)
}
func (s *App) GetSettingsBool (settings string, def bool) bool {
	return s.props.GetBool(settings, def)
}

func (s *App) FilterPrefix(prefix string) *properties.Properties {
	return s.props.FilterPrefix(prefix)
}

func (s *App) AddSession(sessMng *session.SessionManager) {
	s.sessMng = sessMng
}
func (s* App) GetAddress() string {
	return fmt.Sprintf("%s:%s",s.props.MustGet("http.server.host"),s.props.MustGet("http.server.port"))
}
func (s *App) CloseSession() {
	s.sessMng.CloseAll()
}