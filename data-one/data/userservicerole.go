package data

func (r *Repositories) GetUserServiceRoles(id int) {
	r.db.Query("SELECT ... from user_service_roles")
	r.GetUser(id)
	// do something
}

func userServiceRoleHelper() {}
