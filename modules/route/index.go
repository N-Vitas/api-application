package route

import "net/http"

var (
	SwaggerPath string
	ApiIcon     string
)

// If swagger is not on `/` redirect to it
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, SwaggerPath, http.StatusMovedPermanently)
}

func Icon(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, ApiIcon, http.StatusMovedPermanently)
}
