package routes

import (
	"coinbase-go/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//InitRoutes is...
func InitRoutes() {
	router := mux.NewRouter()
	err := godotenv.Load()
	api := router.PathPrefix("/api/v1").Subrouter()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("succesfully load env file")

	api.HandleFunc("/metas", handlerMetas)
	api.HandleFunc("/metadata/ids", handlerMetasIds)
	api.HandleFunc("/ticker", handlerTicker)
	api.HandleFunc("/ticker/ids", handlerTickerIds)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func handlerMetas(w http.ResponseWriter, r *http.Request) {
	// switch r.Method {
	// case "GET":
	// 	controllers.GetAllMetadata(w, r)
	// default:
	// 	http.Error(w, "", http.StatusBadRequest)
	// }
	controllers.GetAllMetadata(w, r)
}

func handlerMetasIds(w http.ResponseWriter, r *http.Request) {
	// switch r.Method {
	// case "POST":
	// 	controllers.GetMetadataIds(w, r)
	// default:
	// 	http.Error(w, "", http.StatusBadRequest)
	// }
	controllers.GetMetadataIds(w, r)
}

func handlerTicker(w http.ResponseWriter, r *http.Request) {
	// switch r.Method {
	// case "GET":
	// 	controllers.GetTickerAll(w, r)
	// default:
	// 	http.Error(w, "", http.StatusBadRequest)
	// }

	controllers.GetTickerAll(w, r)
}

func handlerTickerIds(w http.ResponseWriter, r *http.Request) {
	// switch r.Method {
	// case "POST":
	// 	controllers.GetTickerIds(w, r)
	// default:
	// 	http.Error(w, "", http.StatusBadRequest)
	// }

	controllers.GetTickerIds(w, r)
}
