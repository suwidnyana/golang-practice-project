package controllers

import (
	"coinbase-go/request"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetAllMetadata(w http.ResponseWriter, r *http.Request) {

	secretkey := os.Getenv("API_KEY")

	url := fmt.Sprintf("https://api.nomics.com/v1/currencies?key=%s", secretkey)

	request, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

func GetMetadataIds(w http.ResponseWriter, r *http.Request) {
	var body request.Ids
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return

	}
	secretkey := os.Getenv("API_KEY")

	url := fmt.Sprintf("https://api.nomics.com/v1/currencies/?key=%s&ids=%s&attributes=%s", secretkey, strings.Join(body.IDS, ","), strings.Join(body.Attributes, ","))
	request, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func GetTickerIds(w http.ResponseWriter, r *http.Request) {

	var body request.Ids
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return

	}
	secretkey := os.Getenv("API_KEY")

	url := fmt.Sprintf("https://api.nomics.com/v1/currencies/ticker?key=%s&ids=%s", secretkey, strings.Join(body.IDS, ","))
	request, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func GetTickerAll(w http.ResponseWriter, r *http.Request) {
	secretkey := os.Getenv("API_KEY")

	url := fmt.Sprintf("https://api.nomics.com/v1/currencies/ticker?key=%s", secretkey)
	request, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
