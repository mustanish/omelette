package schemas

import user "github.com/mustanish/omelette/app/schemas/user"

//Schema is exported to be in validation middleware
var Schema = map[string]interface{}{
	"/auth:POST":  user.AuthenticateOpts,
	"/login:POST": user.LoginOpts,
	"/user:PATCH": user.UpdateUserOpts,
}
