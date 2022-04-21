package interactor

import (
	entities "github.com/t-east/pr-sounds-backend/src/domains/entities"
	port "github.com/t-east/pr-sounds-backend/src/usecases/port"
)

type UserHandler struct {
	OutputPort port.UserOutputPort
	Repository port.UserRepository
}

// NewUserInputPort はUserInputPortを取得します．
func NewUserInputPort(outputPort port.UserOutputPort, repository port.UserRepository) port.UserInputPort {
	return &UserHandler{
		OutputPort: outputPort,
		Repository: repository,
	}
}

//* ユーザ登録
func (uc *UserHandler) Create(user *entities.User) (*entities.User, error) {
	// ドメインのvalidation
	if err := user.ValidateName(); err != nil {
		uc.OutputPort.RenderError(err, 400)
	}
	if err := user.ValidateEmail(); err != nil {
		uc.OutputPort.RenderError(err, 400)
	}
	// メールアドレスが重複しているユーザがいないか
	_, err := uc.Repository.FindByEmail(user.Email)
	if err != nil {
		uc.OutputPort.RenderError(err, 400)
	}
	// データベースにユーザを登録
	createdUser, err := uc.Repository.Create(user)
	if err != nil {
		uc.OutputPort.RenderError(err, 400)
	}
	uc.OutputPort.Render(user, 201)
	return createdUser, nil
}

//* ユーザ情報を取得
func (uc *UserHandler) FindByID(id uint) (*entities.User, error) {
	// userが存在することを確認する
	user, err := uc.Repository.FindByID(id)
	if err != nil {
		uc.OutputPort.RenderError(err, 400)
		return nil, err
	}
	// userを返す
	uc.OutputPort.Render(user, 200)
	return user, nil
}

//* ユーザ情報を更新
func (uc *UserHandler) Update(user *entities.User) (*entities.User, error) {
	updatedUser, err := uc.Repository.Update(user)
	if err != nil {
		uc.OutputPort.RenderError(err, 400)
	}
	uc.OutputPort.Render(user, 201)
	return updatedUser, nil
}
