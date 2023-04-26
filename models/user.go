package models

import "database/sql"

type NullShortUser struct {
	ID        sql.NullString `json:"id"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Image     sql.NullString `json:"image"`
}
