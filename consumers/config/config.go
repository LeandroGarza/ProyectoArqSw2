package config

var (
	RABBITUSER     = "user"
	RABBITPASSWORD = "password"
	RABBITHOST     = "localhost"
	EXCHANGE       = "users"
	RABBITPORT     = 5672

	ENDPOINTS = []string{"http://localhost:8090/items/user", "http://localhost:8081", "http://localhost:9001/messages/user"}
)
