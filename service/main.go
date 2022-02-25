package main

import (
	"database/sql"

	"github.com/socialdog-inc/go-data/service/svc"
)

func main() {
	s := svc.NewServices(sql.DB{})
	s.User.Get(1)
	s.UserSvcRole.Get(2)
}
