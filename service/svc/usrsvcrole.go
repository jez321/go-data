package svc

import (
	"github.com/socialdog-inc/go-data/service/usrdata"
	"github.com/socialdog-inc/go-data/service/usrsvcroledata"
)

type UserServiceRoleService struct {
	u   usrdata.Repository
	usr usrsvcroledata.Repository
}

func (s UserServiceRoleService) Get(id int) {
	s.u.Get(id)
	s.usr.Get(id)
	// do something
}
