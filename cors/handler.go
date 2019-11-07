package cors

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func Handler() func(http.Handler) http.Handler {

	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Accept", "token", "authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})

	return handlers.CORS(originsOk, headersOk, methodsOk)
}
