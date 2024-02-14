package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/SantsGu/lista-de-tarefa/entity"
	"github.com/SantsGu/lista-de-tarefa/service"
)

type Controller struct {
	Service *service.Service
}

func NewController(Service *service.Service) *Controller {
	return &Controller{Service: Service}
}

func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	lista := c.Service.FindAll()
	jsonResp, err := json.Marshal(lista)
	if err != nil {
		http.Error(w, "Erro ao conveter para json", http.StatusInternalServerError)
		return
	}

	response(w, http.StatusOK, jsonResp)
}

func (c *Controller) GetId(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var newTask entity.Task

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o body da solicitação", http.StatusInternalServerError)
	}

	err = json.Unmarshal(body, newTask)
	if err != nil {
		http.Error(w, "Não foi possivel realizar conversão do body", http.StatusInternalServerError)
	}
	c.Service.SaveTask(newTask.Tarefa)
	response(w, http.StatusOK, nil)
	return

}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {

}

func response(w http.ResponseWriter, statusCode int, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}
