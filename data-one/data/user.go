package data

func (r *Repositories) GetUser(id int) {
	r.db.QueryRow("SELECT ... from users")
	r.GetUserServiceRoles(id)
	// do something
}

func userHelper() {}
