package user

import (
	"errors"
	"refactoring/model"
	"refactoring/storage"
	"time"
)

var (
	ErrInternalError = errors.New("internal error")
	ErrUserNotFound  = storage.ErrUserNotFound
)

func GetList() (*model.UserList, error) {
	u, err := storage.GetUserList()
	if err != nil {
		return nil, ErrInternalError
	}
	return u, nil
}
func Create(request *model.CreateUserRequest) (string, error) {
	u := model.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.DisplayName,
	}
	id, err := storage.AddUser(u)
	if err != nil {
		return "", ErrInternalError
	}
	return id, nil
}
func GetByID(id string) (*model.User, error) {
	u, err := storage.GetUser(id)
	if err == storage.ErrUserNotFound {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, ErrInternalError
	}
	return u, nil
}
func Update(id string, request model.UpdateUserRequest) error {
	u, err := storage.GetUser(id)
	if err == storage.ErrUserNotFound {
		return ErrUserNotFound
	}
	u.DisplayName = request.DisplayName

	if err = storage.UpdateUser(id, *u); err != nil {
		return ErrInternalError
	}
	return nil
}
func Delete(id string) error {
	if err := storage.DeleteUser(id); err != nil {
		return ErrInternalError
	}
	return nil
}
