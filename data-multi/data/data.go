package data

import (
	"database/sql"

	"github.com/socialdog-inc/go-data/data-multi/usrdata"
	"github.com/socialdog-inc/go-data/data-multi/usrsvcroledata"
)

type Repositories struct {
	User        usrdata.Repository
	UserSvcRole usrsvcroledata.Repository
}

func NewRepositories(db sql.DB) Repositories {
	u := usrdata.Repository{DB: db}
	usr := usrsvcroledata.Repository{DB: db}

	// !!??
	u.UserServiceRole = usr
	usr.User = u

	return Repositories{
		User:        u,
		UserSvcRole: usr,
	}
}
