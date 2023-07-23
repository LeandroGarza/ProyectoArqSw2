package config

import "time"

var (
	CCMAXSIZE      int64  = 1000
	CCITEMSTOPRUNE uint32 = 100
	CCDEFAULTTTL          = 30 * time.Second

	DBHOST     = "localhost"
	DBPORT     = 27017
	COLLECTION = "items"

	RABBITHOST = "localhost"
	RABBITPORT = 5672
	EXCHANGE   = "users"
)
