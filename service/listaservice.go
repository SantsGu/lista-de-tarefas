package service

import (
	"errors"
	"fmt"

	"github.com/SantsGu/lista-de-tarefa/repository"
)

const statusDefault = "Pendente"

func FindAll() []string {
	 
	var m repository.Mapa
	var lista []string

	for id:=1;id<=len(m.Tarefa);id++ {
		
		lista = append(lista, fmt.Sprintf("%v - %s, status: %s\n", id,m.Tarefa[id],m.Status[id])) 
	}
	return lista
}

func FindTask() (string, error) {
	return "", nil
}

func UpdateStatus() ([]string, error) {
	return nil, nil
	
}

func SaveTask(Novatarefa string) {
	var m repository.Mapa
	*m.Id++

	m.Tarefa[*m.Id] = Novatarefa

	m.Status[*m.Id] = statusDefault
}

func DeleteTask(task string) error {
	var m repository.Mapa

	for chave, valor := range m.Tarefa{
		if valor == task{
			delete(m.Tarefa, chave)
			break
		} else{
			return errors.New("A tarefa que você está buscando não foi encontrada na lista de tarefas")
		}
	}

	return nil
}