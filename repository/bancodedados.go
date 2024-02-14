package repository

type Mapa struct {
	Tarefa map[int]string
	Status map[int]string
	Id     *int
}

func NewMapa() *Mapa {
	return &Mapa {
		Tarefa: make(map[int]string),
		Status: make(map[int]string),
		Id:     new(int),
	}
}
