package route

import (
	"net/http"

	"../biryani"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Routes() http.Handler {
	r := mux.NewRouter()
	ActualRoutes(r)
	return getCors(r)
}

func ActualRoutes(r *mux.Router) {
	r.HandleFunc("/take-biryani", biryani.TakeCheckenBiryani).Methods("GET")
	r.HandleFunc("/make-biryani", biryani.MakeCheckenBiryani).Methods("POST")
	r.HandleFunc("/remove-biryani", biryani.RemoveCheckenBiryani).Queries("id", "{key}").Methods("DELETE")
	r.HandleFunc("/take-one-biryani", biryani.TakeSingleBiryani).Queries("id", "{key}").Methods("GET")
	r.HandleFunc("/remake-biryani", biryani.RemakeSingleBiryani).Queries("id", "{key}").Methods("PUT")
}

func getCors(r http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8787"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{"Accept",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"Access-Control-Expose-Headers",
			"Access-Control-Request-Headers",
			"Access-Control-Request-Method",
			"Connection",
			"Host",
			"Origin",
			"User-Agent",
			"Referer",
			"Cache-Control",
			"X-header",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Credentials",
		},
	})
	return c.Handler(r)
}
