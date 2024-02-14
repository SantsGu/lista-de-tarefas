package entity


type Lista struct {
	Tasks []Task `json:"tarefas"`
}

type Task struct {
	Id int `json:"id"`
	Status string `json:"status"`
	Tarefa string `json:"tarefa"`
}

