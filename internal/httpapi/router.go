
package httpapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)


func health(response http.ResponseWriter, _ *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	_, _ = response.Write([]byte(`{"status":"ok"}`))
}


func Router() http.Handler {

	router := chi.NewRouter()
	router.Get("/health", health)
	return router

}
