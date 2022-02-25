package svc

import (
	"github.com/socialdog-inc/go-data/service/usrdata"
	"github.com/socialdog-inc/go-data/service/usrsvcroledata"
)

type UserService struct {
	u   usrdata.Repository
	usr usrsvcroledata.Repository
}

func (s UserService) Get(id int) {
	s.u.Get(id)
	s.usr.Get(id)
	// do something
}
