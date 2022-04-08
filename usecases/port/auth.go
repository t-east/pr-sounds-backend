package port

import (
	entities "github.com/t-east/pr-sounds-backend/domains/entities"
)

type Auth interface {
	Login(*entities.LoginUser) (string, error)
}

type AuthInputPort interface {
	Login(*entities.LoginUser) (string, error)
}

type AuthOutputPort interface {
	Render(string, int)
	RenderError(error, int)
}
