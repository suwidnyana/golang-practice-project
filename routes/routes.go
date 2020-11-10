package routes

import (
	"coinbase-go/controllers"
	"net/http"
)

//InitRoutes is...
func InitRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write([]byte("get"))
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	http.HandleFunc("/metadata", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetAllMetadata(w, r)
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	http.HandleFunc("/metadata/ids", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			controllers.GetMetadataIds(w, r)
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	http.HandleFunc("/ticker", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetTickerAll(w, r)
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	http.HandleFunc("/ticker/ids", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			controllers.GetTickerIds(w, r)
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})
}
