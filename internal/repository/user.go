package repository

import "github.com/gustavolimam/coinext-challenge/internal/model"

type UserQueries interface {
	CreateUser(user *model.User) error
	AddOrRemoveItem(inventory *model.Inventory) error
	GetUser(userID string) (*model.User, error)
	ListUsers() ([]model.User, error)
}

func (r *repository) CreateUser(user *model.User) error {
	query := `INSERTO INTO users (name, age, gender, longitude, latitude) VALUES ($1, $2, $3, $4, $5)`

	if _, err := r.db.Exec(query, user.Name, user.Age, user.Gender, user.LastLocation.Longitude, user.LastLocation.Latitude); err != nil {
		return err
	}

	return nil
}

func (r *repository) AddOrRemoveItem(inventory *model.Inventory) error {
	query := `INSERTO INTO users (user_id, water, food, drug, ammo) VALUES ($1, $2, $3, $4, $5) 
	ON CONFLICT (user_id)
	DO UPDATE SET water=$2, food=$3, drug=$4, ammo=$5 WHERE user_id=$1`

	if _, err := r.db.Exec(query, inventory.UserID, inventory.Water, inventory.Food, inventory.Drug, inventory.Ammo); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUser(userID string) (*model.User, error) {
	user := &model.User{}

	query := `SELECT * FROM users WHERE id=$1`
	row := r.db.QueryRow(query, userID)
	if err := row.Scan(&user.ID,
		&user.Name,
		&user.Age,
		&user.Gender,
		&user.LastLocation.Longitude,
		&user.LastLocation.Latitude,
		&user.CreatedAt,
		&user.DeadAt); err != nil {
		return nil, err
	}

	query2 := `SELECT * FROM inventory WHERE user_id=$1`
	row2 := r.db.QueryRow(query2, userID)
	if err := row2.Scan(&user.Inventory.UserID,
		&user.Inventory.Water,
		&user.Inventory.Food,
		&user.Inventory.Drug,
		&user.Inventory.Ammo); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) ListUsers() ([]model.User, error) {
	users := []model.User{}

	query := `SELECT * FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := model.User{}
		if err := rows.Scan(&user.ID,
			&user.Name,
			&user.Age,
			&user.Gender,
			&user.LastLocation.Longitude,
			&user.LastLocation.Latitude,
			&user.CreatedAt,
			&user.DeadAt); err != nil {
			return nil, err
		}

		query2 := `SELECT * FROM inventory WHERE user_id=$1`
		row2 := r.db.QueryRow(query2, user.ID)
		if err := row2.Scan(&user.Inventory.UserID,
			&user.Inventory.Water,
			&user.Inventory.Food,
			&user.Inventory.Drug,
			&user.Inventory.Ammo); err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}
