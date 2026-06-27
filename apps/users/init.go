package users

import (
    "gohr/core"
	"gohr/apps/users/usersModels"
)
	

func init() {

    // Register routes
    core.Register(UserRoutes)

    // Register models for migrations
    core.RegisterModel(usersModels.User{})
    core.RegisterModel(usersModels.Role{})
    core.RegisterModel(usersModels.Permission{})
    core.RegisterModel(usersModels.Log{})
}
