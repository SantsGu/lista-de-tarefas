package service

import (
	"errors"
	"fmt"

	"github.com/SantsGu/lista-de-tarefa/entity"
	"github.com/SantsGu/lista-de-tarefa/repository"
)

const statusDefault = ("Pendente")

type Service struct {
	Repository *repository.Mapa
}

func NewService(Repository *repository.Mapa) *Service {
	return &Service{Repository: Repository}
}

func (s *Service) FindAll() entity.Lista {

	var lista entity.Lista

	for id := 1; id <= len(s.Repository.Tarefa); id++ {

		lista.Tasks = append(lista.Tasks, entity.Task{Id: id, Tarefa: s.Repository.Tarefa[id], Status: s.Repository.Status[id]})
	}
	return lista
}

func (s *Service) FindTask(Task string) (string, error) {

	for chave, valor := range s.Repository.Tarefa {

		if valor == Task {
			return fmt.Sprintf("%v - %s, status: %s\n", chave, s.Repository.Tarefa[chave], s.Repository.Status[chave]), nil
		}
	}

	return "", errors.New("Tarefa não encontrada")
}

func (s *Service) UpdateStatus(task, status string) ([]string, error) {

	var lista []string

	for chave, valor := range s.Repository.Tarefa {

		if valor == task {

			s.Repository.Status[chave] = status
			lista = append(lista, fmt.Sprintf("%v - %s, status: %s", chave, s.Repository.Tarefa[chave], s.Repository.Status[chave]))
			return lista, nil
		}

	}
	return nil, errors.New("Não foi possivel atualizar a tarefa porque ela nao existe")

}

func (s Service) SaveTask(Novatarefa string) {
	*s.Repository.Id++

	s.Repository.Tarefa[*s.Repository.Id] = Novatarefa

	s.Repository.Status[*s.Repository.Id] = statusDefault
}

func (s *Service) DeleteTask(task string) error {

	for chave, valor := range s.Repository.Tarefa {
		if valor == task {
			delete(s.Repository.Tarefa, chave)
			break
		} else {
			return errors.New("A tarefa que você está buscando não foi encontrada na lista de tarefas")
		}
	}

	return nil
}
