package app

import (
	
	"net/http"
	"github.com/SantsGu/lista-de-tarefa/controller"
)

func Routes(c *controller.Controller) {
	http.HandleFunc("/new-task/", c.Create)
	http.HandleFunc("/complete-list", c.GetAll)
	http.HandleFunc("/specific-task/", c.GetId)
	http.HandleFunc("/change-status", c.Update)
	http.HandleFunc("/delete/", c.Delete)
	http.ListenAndServe(":8080",nil)

}
