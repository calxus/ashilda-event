package main

import (
	"log"
	"database/sql"
)

type Event struct {
	Event_id int `json:"id"`
	Name string `json:"name"`
}

func (e *Event) populate(rows *sql.Rows) {
	err := rows.Scan(&e.Event_id, &e.Name)
	if err != nil {
		log.Print(err)
	}
}