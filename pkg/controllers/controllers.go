package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xiayulin123/GO_GameStoreAPI/pkg/model"
)

var NewToy model.Toys

func ShowToys(w http.ResponseWriter, r *http.Request) {
	NewToy := model.GetToys()
	res, _ := json.Marshal(NewToy)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ShowToyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	toyId := vars["toyId"]
	ID, err := strconv.ParseInt(toyId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	toyDetail, _ := model.GetToyById(ID)

	res, _ := json.Marshal(toyDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
