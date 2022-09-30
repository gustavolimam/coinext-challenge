package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gustavolimam/coinext-challenge/internal/model"
	"github.com/gustavolimam/coinext-challenge/internal/repository"
	mockRepo "github.com/gustavolimam/coinext-challenge/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_user_CreateUser(t *testing.T) {
	type fields struct {
		repo repository.Queries
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should create a new user",
			fields: fields{
				repo: mockRepo.NewQueries(t),
			},
			args: args{
				user: &model.User{
					Name:   "Gustavo",
					Age:    25,
					Gender: "M",
					LastLocation: model.Location{
						Latitude:  "1",
						Longitude: "1",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Should return error",
			fields: fields{
				repo: mockRepo.NewQueries(t),
			},
			args: args{
				user: &model.User{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				repo: tt.fields.repo,
			}

			switch tt.name {
			case "Should create a new user":
				tt.fields.repo.(*mockRepo.Queries).On("CreateUser", tt.args.user).Return(nil)
			case "Should return error":
				tt.fields.repo.(*mockRepo.Queries).On("CreateUser", tt.args.user).Return(errors.New("error"))
			}

			if err := u.CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("user.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_user_AddOrRemoveItem(t *testing.T) {
	type fields struct {
		repo repository.Queries
	}
	type args struct {
		item *model.Inventory
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should add or remove item",
			fields: fields{
				repo: mockRepo.NewQueries(t),
			},
			args: args{
				item: &model.Inventory{
					UserID: "1",
					Water:  1,
					Food:   1,
					Drug:   1,
					Ammo:   1,
				},
			},
			wantErr: false,
		},
		{
			name: "Should return error",
			fields: fields{
				repo: mockRepo.NewQueries(t),
			},
			args: args{
				item: &model.Inventory{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				repo: tt.fields.repo,
			}

			switch tt.name {
			case "Should add or remove item":
				tt.fields.repo.(*mockRepo.Queries).On("AddOrRemoveItem", tt.args.item).Return(nil)
			case "Should return error":
				tt.fields.repo.(*mockRepo.Queries).On("AddOrRemoveItem", tt.args.item).Return(errors.New("error"))
			}

			if err := u.AddOrRemoveItem(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("user.AddOrRemoveItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_user_Trade(t *testing.T) {
	userID := "1"
	userAlternativeID := "2"
	userT := &model.User{
		ID:   &userID,
		Name: "Alice",
		Age:  22,
		Inventory: &model.Inventory{
			UserID: userID,
			Water:  1,
			Food:   1,
			Drug:   1,
			Ammo:   1,
		},
	}

	users := []model.User{
		{
			ID:   &userAlternativeID,
			Name: "Gustavo",
			Age:  25,
			Inventory: &model.Inventory{
				UserID: userAlternativeID,
				Water:  1,
				Food:   1,
				Drug:   1,
				Ammo:   1,
			},
		},
	}

	notMatchUsers := []model.User{
		{
			ID:   &userAlternativeID,
			Name: "Gustavo",
			Age:  25,
			Inventory: &model.Inventory{
				UserID: userAlternativeID,
				Water:  1,
				Food:   1,
				Drug:   1,
				Ammo:   2,
			},
		},
	}

	type fields struct {
		repo repository.Queries
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should trade items",
			fields: fields{
				repo: mockRepo.NewQueries(t),
			},
			args: args{
				userID: userID,
			},
			wantErr: false,
		},
		{
			name: "Should not trade items because user does not match",
			fields: fields{
				repo: mockRepo.NewQueries(t),
			},
			args: args{
				userID: userID,
			},
			wantErr: true,
		},
		{
			name: "Should not trade items because user does not exist",
			fields: fields{
				repo: mockRepo.NewQueries(t),
			},
			args: args{
				userID: "3",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				repo: tt.fields.repo,
			}

			switch tt.name {
			case "Should trade items":
				tt.fields.repo.(*mockRepo.Queries).On("GetUser", tt.args.userID).Return(userT, nil)
				tt.fields.repo.(*mockRepo.Queries).On("ListUsers").Return(users, nil)
				tt.fields.repo.(*mockRepo.Queries).On("AddOrRemoveItem", mock.Anything).Return(nil)
				tt.fields.repo.(*mockRepo.Queries).On("AddOrRemoveItem", mock.Anything).Return(nil)
			case "Should not trade items because user does not match":
				tt.fields.repo.(*mockRepo.Queries).On("GetUser", tt.args.userID).Return(userT, nil)
				tt.fields.repo.(*mockRepo.Queries).On("ListUsers").Return(notMatchUsers, nil)
			case "Should not trade items because user does not exist":
				tt.fields.repo.(*mockRepo.Queries).On("GetUser", tt.args.userID).Return(nil, errors.New("error"))
			}

			if err := u.Trade(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("user.Trade() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_user_perfectMatch(t *testing.T) {
	userID := "1"
	userAlternativeID := "2"

	type fields struct {
		repo repository.Queries
	}
	type args struct {
		itemToTrade model.Inventory
		users       []model.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.User
	}{
		{
			name: "Should return a perfect match",
			args: args{
				itemToTrade: model.Inventory{
					UserID: userID,
					Water:  1,
					Food:   1,
					Drug:   1,
					Ammo:   1,
				},
				users: []model.User{
					{
						ID: &userAlternativeID,
						Inventory: &model.Inventory{
							UserID: userAlternativeID,
							Water:  1,
							Food:   1,
							Drug:   1,
							Ammo:   1,
						},
					},
				},
			},
			want: &model.User{
				ID: &userAlternativeID,
				Inventory: &model.Inventory{
					UserID: userAlternativeID,
					Water:  1,
					Food:   1,
					Drug:   1,
					Ammo:   1,
				},
			},
		},
		{
			name: "Should return error",
			args: args{
				itemToTrade: model.Inventory{
					UserID: userID,
					Water:  1,
					Food:   2,
					Drug:   1,
					Ammo:   1,
				},
				users: []model.User{
					{
						ID: &userAlternativeID,
						Inventory: &model.Inventory{
							UserID: userAlternativeID,
							Water:  1,
							Food:   1,
							Drug:   1,
							Ammo:   1,
						},
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				repo: tt.fields.repo,
			}
			if got := u.perfectMatch(tt.args.itemToTrade, tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("user.perfectMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
