package data

import "database/sql"

type Repositories struct {
	db sql.DB
}

func NewRepositories(db sql.DB) Repositories {
	return Repositories{db: db}
}
