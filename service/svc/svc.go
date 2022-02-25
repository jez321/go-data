package svc

import (
	"database/sql"

	"github.com/socialdog-inc/go-data/service/usrdata"
	"github.com/socialdog-inc/go-data/service/usrsvcroledata"
)

type Services struct {
	User        UserService
	UserSvcRole UserServiceRoleService
}

func NewServices(db sql.DB) Services {
	u := usrdata.NewRepository(db)
	usr := usrsvcroledata.NewRepository(db)

	return Services{
		User:        UserService{u: u, usr: usr},
		UserSvcRole: UserServiceRoleService{u: u, usr: usr},
	}
}
