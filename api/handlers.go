package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RinksHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(Rinks)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(resp)
}

func RinkInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	i, err := strconv.Atoi(vars["rinkID"])
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	rinkInfo, err := fetchRinkInfo(i)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	resp, err := json.Marshal(rinkInfo)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(resp)
}
