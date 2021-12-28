//Package storage implements functions to manipulate stored data
package storage

import (
	"errors"
	"refactoring/model"
	"refactoring/storage/file"
	"strconv"
)

var (
	ErrCouldNotRead  = errors.New("could not read from storage")
	ErrCouldNotWrite = errors.New("could not write new data to storage")
	ErrUserNotFound  = errors.New("user not found")
)

type (
	//UserStorage specifies additional data of storing objects, and the way how data stored
	UserStorage struct {
		Increment int            `json:"increment"`
		List      model.UserList `json:"list"`
	}
)

func GetUserList() (*model.UserList, error) {
	s := UserStorage{}
	if err := file.Read(&s); err != nil {
		return nil, ErrCouldNotRead
	}
	return &s.List, nil
}

func AddUser(u model.User) (string, error) {
	s := UserStorage{}
	if err := file.Read(&s); err != nil {
		return "", ErrCouldNotRead
	}
	id := strconv.Itoa(s.Increment)
	s.List[id] = u
	if err := file.Write(&s); err != nil {
		return "", ErrCouldNotWrite
	}
	return id, nil
}

func GetUser(id string) (*model.User, error) {
	s := UserStorage{}
	if err := file.Read(&s); err != nil {
		return nil, ErrCouldNotRead
	}
	u, ok := s.List[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return &u, nil
}
func UpdateUser(id string, u model.User) error {
	s := UserStorage{}
	if err := file.Read(&s); err != nil {
		return ErrCouldNotRead
	}
	if _, ok := s.List[id]; !ok {
		return ErrUserNotFound
	}
	s.List[id] = u
	if err := file.Write(&s); err != nil {
		return ErrCouldNotWrite
	}
	return nil
}

func DeleteUser(id string) error {
	s := UserStorage{}
	if err := file.Read(&s); err != nil {
		return ErrCouldNotRead
	}
	if _, ok := s.List[id]; !ok {
		return ErrUserNotFound
	}
	delete(s.List, id)
	if err := file.Write(&s); err != nil {
		return ErrCouldNotWrite
	}
	return nil
}
