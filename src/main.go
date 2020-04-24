package main

import (
	db "src/database"
	"src/redis"
)

func main()  {
	defer db.MyDb.Close()
	defer redis.RedisClient.Close()
	app:=initRouter()
	app.Run("127.0.0.1:8089")
}
