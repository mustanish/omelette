package connectors

var users = []map[string]interface{}{
	{"type": "persistent", "fields": []string{"phone"}, "name": "uniquePhone"},
	{"type": "persistent", "fields": []string{"email"}, "name": "uniqueEmail"},
}

var tokens = []map[string]interface{}{
	{"type": "ttl", "field": "expiresAt", "expireAfter": 10, "name": "tokenExpire"},
}

var index = map[string]interface{}{
	"tokens": tokens,
	"users":  users,
}
