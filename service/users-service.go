package service

import (
	"context"

	"github.com/kudabab/mr-test/db"
	"github.com/kudabab/mr-test/entity"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
}

type userService struct {
	users []entity.User
}

func New() UserService {
	return &userService{}
}

func (s *userService) Save(user entity.User) entity.User {
	s.users = append(s.users, user)
	db.GetDB().Exec(context.Background(), "INSERT INTO users (id, login, password) VALUES ($1,$2,$3)", user.ID, user.Login, user.Password)

	return user
}

func (s *userService) FindAll() []entity.User {
	return s.users
}

/*func CreateUser(user entity.User) error {
	_, err := db.GetDB().Exec(context.Background(), "INSERT INTO users (id, login, password) VALUES ($1)", user.Login)
	return err
}*/
