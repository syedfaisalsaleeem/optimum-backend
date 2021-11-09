package main

import (
	"database/sql"
)

type todo struct {
	ID       int    `json:"id"`
	Todolist string `json:"todolist"`
}

func (p *todo) createtodolistitem(db *sql.DB) error {
	data := removeExtraSpaces(p.Todolist)
	err := db.QueryRow(
		"INSERT INTO todo(Todolist) VALUES($1) RETURNING id",
		data).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func gettodolist(db *sql.DB) ([]todo, error) {
	rows, err := db.Query(
		"SELECT id, todolist FROM todo ")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todolistitems := []todo{}

	for rows.Next() {
		var p todo
		if err := rows.Scan(&p.ID, &p.Todolist); err != nil {
			return nil, err
		}
		todolistitems = append(todolistitems, p)
	}

	return todolistitems, nil
}
