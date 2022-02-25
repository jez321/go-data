package main

import (
	"database/sql"

	"github.com/socialdog-inc/go-data/data-multi/data"
)

func main() {
	r := data.NewRepositories(sql.DB{})
	r.User.Get(1)
	r.UserSvcRole.Get(2)
}
