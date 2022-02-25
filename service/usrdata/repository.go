package usrdata

import "database/sql"

type Repository struct {
	db sql.DB
}

func NewRepository(db sql.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) Get(id int) {
	r.db.QueryRow("SELECT ... from users")
}
