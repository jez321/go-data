package main

import (
	"database/sql"

	"github.com/socialdog-inc/go-data/data-one/data"
)

func main() {
	r := data.NewRepositories(sql.DB{})
	r.GetUser(1)
	r.GetUserServiceRoles(2)
}
