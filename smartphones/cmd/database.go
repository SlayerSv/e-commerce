package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	*sql.DB
}

const PAGE_LIMIT int = 20

func (db *PostgresDB) GetAll() ([]Smartphone, error) {
	rows, err := db.Query("SELECT * FROM smartphones ORDER BY price LIMIT $1", PAGE_LIMIT)
	if err != nil {
		return nil, err
	}
	smartphones, err := db.ExtractMany(rows)
	if err != nil {
		return nil, err
	}
	return smartphones, nil
}

func (db *PostgresDB) GetOne(id int) (Smartphone, error) {
	row := db.QueryRow("SELECT * FROM smartphones WHERE id = $1", id)
	sm, err := db.ExtractOne(row)
	if err != nil {
		return sm, err
	}
	return sm, nil
}

func (db *PostgresDB) Delete(id int) (Smartphone, error) {
	row := db.QueryRow("DELETE FROM smartphones where id = $1 returning *", id)
	sm, err := db.ExtractOne(row)
	if err != nil {
		return sm, err
	}
	return sm, nil
}

func (db *PostgresDB) Update(sm Smartphone) (Smartphone, error) {
	query := `
	UPDATE smartphones
	SET model = $1, producer = $2, color = $3, screen_size = $4,
	description = $5, image = $6, price = $7
	WHERE id = $8
	RETURNING *
	`
	row := db.QueryRow(query, sm.Model, sm.Producer, sm.Color, sm.ScreenSize,
		sm.Description, sm.Image, sm.Price, sm.ID)
	sm, err := db.ExtractOne(row)
	if err != nil {
		return sm, err
	}
	return sm, nil
}

func (db *PostgresDB) Create(sm Smartphone) (Smartphone, error) {
	query := `
	INSERT INTO smartphones (model, producer, color, screen_size, description, image, price)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING *
	`
	row := db.QueryRow(query, sm.Model, sm.Producer, sm.Color, sm.ScreenSize,
		sm.Description, sm.Image, sm.Price)
	sm, err := db.ExtractOne(row)
	if err != nil {
		return sm, err
	}
	return sm, nil
}

func (db *PostgresDB) ExtractOne(row *sql.Row) (Smartphone, error) {
	sm := Smartphone{}
	err := row.Scan(&sm.ID, &sm.Model, &sm.Producer, &sm.Color,
		&sm.ScreenSize, &sm.Description, &sm.Image, &sm.Price)
	if err != nil {
		return sm, err
	}
	return sm, nil
}

func (db *PostgresDB) ExtractMany(rows *sql.Rows) ([]Smartphone, error) {
	defer rows.Close()
	smartphones := []Smartphone{}
	for rows.Next() {
		sm := Smartphone{}
		err := rows.Scan(&sm.ID, &sm.Model, &sm.Producer, &sm.Color,
			&sm.ScreenSize, &sm.Description, &sm.Image, &sm.Price)
		if err != nil {
			return nil, err
		}
		smartphones = append(smartphones, sm)
	}
	return smartphones, nil
}
