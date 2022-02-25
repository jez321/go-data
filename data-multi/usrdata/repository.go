package usrdata

import "database/sql"

type UserServiceRoleGetter interface {
	Get(id int)
}

type Repository struct {
	DB              sql.DB
	UserServiceRole UserServiceRoleGetter
}

func (r Repository) Get(id int) {
	r.DB.QueryRow("SELECT ... from users")
	r.UserServiceRole.Get(id)
	// do something
}
