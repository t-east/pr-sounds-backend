package interactor

import (
	"github.com/t-east/pr-sounds-backend/src/domains/entities"
	port "github.com/t-east/pr-sounds-backend/src/usecases/port"
)

type AuthHandler struct {
	OutputPort port.AuthOutputPort
	Auth       port.Auth
}

// NewUserInputPort はUserInputPortを取得します．
func NewAuthInputPort(outputPort port.AuthOutputPort, auth port.Auth) port.AuthInputPort {
	return &AuthHandler{
		OutputPort: outputPort,
		Auth:       auth,
	}
}

func (ah *AuthHandler) Login(user *entities.LoginUser) (string, error) {
	token, err := ah.Auth.Login(user)
	if err != nil {
		ah.OutputPort.RenderError(err, 400)
		return "", err
	}
	ah.OutputPort.Render(token, 200)
	return token, nil
}
