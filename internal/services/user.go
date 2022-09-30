package services

import (
	"errors"

	"github.com/gustavolimam/coinext-challenge/internal/model"
	"github.com/gustavolimam/coinext-challenge/internal/repository"
	"github.com/rs/zerolog/log"
)

type UserI interface {
	CreateUser(user *model.User) error
	AddOrRemoveItem(item *model.Inventory) error
	Trade(userID string) error
	perfectMatch(itemToTrade model.Inventory, users []model.User) *model.User
}

type user struct {
	repo repository.Queries
}

func New() UserI {
	return &user{repository.New()}
}

func (u *user) CreateUser(user *model.User) error {

	return u.repo.CreateUser(user)
}

func (u *user) AddOrRemoveItem(item *model.Inventory) error {

	return u.repo.AddOrRemoveItem(item)
}

func (u *user) Trade(userID string) error {
	user, err := u.repo.GetUser(userID)
	if err != nil {
		log.Error().Err(err).Msg("Error trying to get user data")
		return err
	}

	users, err := u.repo.ListUsers()
	if err != nil {
		log.Error().Err(err).Msg("Error trying to list users")
		return err
	}

	match := u.perfectMatch(*user.Inventory, users)
	if match == nil {
		log.Warn().Msg("No perfect match found")
		return errors.New("no perfect match found")
	}

	user.Inventory.UserID = *match.ID
	if err := u.repo.AddOrRemoveItem(user.Inventory); err != nil {
		log.Error().Err(err).Msg("Error trying to add or remove item to user")
		return err
	}

	match.ID = &userID
	if err := u.repo.AddOrRemoveItem(match.Inventory); err != nil {
		log.Error().Err(err).Msg("Error trying to add or remove item to match user")
		return err
	}

	return nil
}

func (u *user) perfectMatch(itemToTrade model.Inventory, users []model.User) *model.User {
	var perfectMatch *model.User

	for _, user := range users {
		if user.Inventory != nil {
			if sumItemsPoints(itemToTrade) == sumItemsPoints(*user.Inventory) {
				perfectMatch = &user
			}
		}
	}

	return perfectMatch
}
