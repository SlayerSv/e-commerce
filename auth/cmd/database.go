package main

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	*sql.DB
}

func (db *PostgresDB) GetUserByName(name string) (User, error) {
	row := db.QueryRow("SELECT * FROM Users WHERE name = $1", name)
	return db.ExtractOne(row)
}

func (db *PostgresDB) CreateUser(user User) error {
	query := `
	INSERT INTO users (name, password)
	VALUES ($1, $2)
	`
	_, err := db.GetUserByName(user.Name)
	if err == nil {
		return errUserExists
	}
	_, err = db.Exec(query, user.Name, user.Password)
	return err
}

func (db *PostgresDB) ExtractOne(row *sql.Row) (User, error) {
	user := User{}
	err := row.Scan(&user.Name, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, errUserNotFound
		}
		return User{}, err
	}
	return user, nil
}
