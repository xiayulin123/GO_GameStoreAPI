package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xiayulin123/GO_GameStoreAPI/pkg/model"
	"github.com/xiayulin123/GO_GameStoreAPI/pkg/utils"
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

func CreateToys(w http.ResponseWriter, r *http.Request) {
	CreateToy := &model.Toys{}
	utils.ParseBody(r, CreateToy)
	body := CreateToy.CreateToy()

	res, _ := json.Marshal(body)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteToys(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	toyId := vars["toyId"]
	Id, err := strconv.ParseInt(toyId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	toy := model.DeleteToy(Id)

	res, _ := json.Marshal(toy)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateToys(w http.ResponseWriter, r *http.Request) {
	var updateToy = &model.Toys{}
	utils.ParseBody(r, updateToy)
	vars := mux.Vars(r)
	toyId := vars["toyId"]
	ID, err := strconv.ParseInt(toyId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	toy, db := model.GetToyById(ID)
	if toy.Name != "" {
		toy.Name = updateToy.Name
	}
	if toy.Price != "" {
		toy.Price = updateToy.Price
	}
	if toy.Description != "" {
		toy.Description = updateToy.Description
	}
	db.Save(&toy)
	res, _ := json.Marshal(toy)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
