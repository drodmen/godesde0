package modelos

import (
	"time"
)

// solo para definiciones
type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// funcionalidad
func (usuario *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreatedAt = createdAt
	usuario.Status = status
}
