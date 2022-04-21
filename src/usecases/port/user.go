package port

import (
	entities "github.com/t-east/pr-sounds-backend/src/domains/entities"
)

type UserInputPort interface {
	Create(*entities.User) (*entities.User, error)
	FindByID(uint) (*entities.User, error)
	Update(*entities.User) (*entities.User, error)
}

type UserOutputPort interface {
	Render(*entities.User, int)
	RenderError(error, int)
}

type UserRepository interface {
	Create(*entities.User) (*entities.User, error)
	Update(*entities.User) (*entities.User, error)
	FindByID(uint) (*entities.User, error)
	FindByEmail(string) (*entities.User, error)
}
