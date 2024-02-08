package app

import (
	
	"net/http"
	"github.com/SantsGu/lista-de-tarefa/controller"
)

func Routes() {
	http.HandleFunc("/new-task/{tarefa}", controller.Create)
	http.HandleFunc("/complete-list", controller.GetAll)
	http.HandleFunc("/specific-task/{tarefa}", controller.GetId)
	http.HandleFunc("/change-status", controller.Update)
	
	http.HandleFunc("/delete/{tarefa}", controller.Delete)
	http.ListenAndServe(":8080",nil)

}
