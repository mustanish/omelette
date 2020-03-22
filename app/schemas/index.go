package schemas

import user "github.com/mustanish/omelette/app/schemas/user"

//Schema is exported to be in validation middleware
var Schema = map[string]interface{}{
	"/user/auth:POST": user.Authenticate,
}
