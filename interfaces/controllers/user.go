package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/t-east/pr-sounds-backend/core"
	"github.com/t-east/pr-sounds-backend/domains/entities"
	"github.com/t-east/pr-sounds-backend/interfaces/gateways"
	"github.com/t-east/pr-sounds-backend/interfaces/presenters"
	"github.com/t-east/pr-sounds-backend/usecases/interactor"
	"github.com/t-east/pr-sounds-backend/usecases/port"

	"gorm.io/gorm"
)

type UserController struct {
	// -> gateway.NewUserRepository
	RepoFactory func(c *gorm.DB) port.UserRepository
	// -> presenter.NewUserOutputPort
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	// -> interactor.NewUserInputPort
	InputFactory func(
		o port.UserOutputPort,
		u port.UserRepository,
	) port.UserInputPort
	Conn *gorm.DB
}

func LoadUserController(db *gorm.DB) *UserController {
	return &UserController{Conn: db}
}

func (uc *UserController) Post(w http.ResponseWriter, r *http.Request) {
	user := &entities.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	outputPort := presenters.NewUserOutputPort(w)
	repository := gateways.NewUserRepository(uc.Conn)
	inputPort := interactor.NewUserInputPort(outputPort, repository)
	_, _ = inputPort.Create(user)
}

func (uc *UserController) Get(w http.ResponseWriter, r *http.Request) {
	_, tail := core.ShiftPath(r.URL.Path)
	_, tail = core.ShiftPath(tail)
	id, _ := core.ShiftPath(tail)
	outputPort := presenters.NewUserOutputPort(w)
	repository := gateways.NewUserRepository(uc.Conn)
	inputPort := interactor.NewUserInputPort(outputPort, repository)
	_, _ = inputPort.FindByID(id)
}

func (uc *UserController) Dispatch(w http.ResponseWriter, r *http.Request) {
	a, _ := core.ShiftPath(r.URL.Path)
	switch r.Method {
	case "POST":
		uc.Post(w, r)
	case "GET":
		if a == "" {
			return
		}
		uc.Get(w, r)
	default:
		http.NotFound(w, r)
	}
}
