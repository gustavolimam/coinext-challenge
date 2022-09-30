package services

import (
	"github.com/gustavolimam/coinext-challenge/internal/model"
	"github.com/gustavolimam/coinext-challenge/internal/repository"
)

type UserI interface {
	CreateUser() error
	AddOrRemoveItem() error
	Trade() error
}

type user struct {
	repo repository.Queries
}

func New() UserI {
	return &user{repository.New()}
}

func (u *user) CreateUser(user *model.User) error {

	return nil
}

func (u *user) AddOrRemoveItem() error {

	return nil
}

func (u *user) Trade() error {

	return nil
}
