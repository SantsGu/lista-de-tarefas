package app

import (
	"github.com/SantsGu/lista-de-tarefa/controller"
	"github.com/SantsGu/lista-de-tarefa/repository"
	"github.com/SantsGu/lista-de-tarefa/service"
)

func Build() {

	r := repository.NewMapa()
	s := service.NewService(r)
	c := controller.NewController(s)

	Routes(c)
}