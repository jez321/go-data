package usrsvcroledata

import (
	"database/sql"
)

type UserGetter interface {
	Get(id int)
}

type Repository struct {
	DB   sql.DB
	User UserGetter
}

func NewRepository(db sql.DB, u UserGetter) Repository {
	return Repository{
		DB:   db,
		User: u,
	}
}

func (r Repository) Get(id int) {
	r.DB.QueryRow("SELECT ... from user_service_roles")
	r.User.Get(id)
	// do something
}
