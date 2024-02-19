package routes

import (
	"net/http"
)

func SetupAuthenticationRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/signin", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST /api/signin"))
	})

	mux.HandleFunc("POST /api/signup", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST /api/signup"))
	})
}
