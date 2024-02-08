package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/SantsGu/lista-de-tarefa/service"
)



func GetAll(w http.ResponseWriter, r *http.Request) {
	lista := service.FindAll()
	jsonResp, err := json.Marshal(lista)
	if err != nil{
		http.Error(w, "Erro ao conveter para json", http.StatusInternalServerError)
		return
	}

	response(w, http.StatusOK, jsonResp)
}

func GetId(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Create(w http.ResponseWriter, r *http.Request) {
	 path := r.URL.Path
	 pieces := strings.Split(path, "/")
	 if len(pieces) <3{
		http.NotFound(w, r)
		return 
	 }
	 service.SaveTask(pieces[2])
	 response(w, http.StatusOK, nil)

}

func Delete(w http.ResponseWriter, r *http.Request) {

}

func response(w http.ResponseWriter, statusCode int, json []byte){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}