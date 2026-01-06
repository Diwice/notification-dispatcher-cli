package user

import (
	"log"
	"context"
	"database/sql"
	"not-dis/internal/db"
)

func CreateUser(database *db.Queries, pref, mail, ph_num string) (db.User, error) {
	createParams := db.CreateUserParams{
		Preferences: sql.NullString{String: pref, Valid: true},
		Mail:        sql.NullString{String: mail, Valid: true},
		PhoneNumber: sql.NullString{String: ph_num, Valid: true},
	}

	newUser, err := database.CreateUser(context.Background(), createParams)
	if err != nil {
		return db.User{}, error
	}

	return newUser, nil
}

func DeleteUser(database *db.Queries, id int64) error {
	if err := database.DeleteUserByID(context.Background(), id); err != nil {
		return err
	}

	return nil
}

func ModifyUser(database *db.Queries, id int64, pref, mail, ph_num string) (db.User, error) {
	modifyParams = db.UpdateUserByIDParams{
		ID:          id,
		Preferences: sql.NullString{String: pref, Valid: true},
		Mail:        sql.NullString{String: mail, Valid: true},
		PhoneNumber: sql.NullString{String: ph_num, Valid: true},
	}

	modifiedUser, err := database.UpdateUserByID(context.Background(), modifyParams)
	if err != nil {
		return db.User{}, err
	}

	return modifiedUser, nil
}

func DeleteUsers(database *db.Queries) error {
	if err := database.DeleteAllUsers(context.Background()); err != nil {
		return err
	}

	return nil
}

func GetUserInfo(database *db.Queries, id int64) (db.User, error) {
	user, err := database.GetUserByID(context.Background(), id)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}
