package main

import (
	db "src/database"
)

func main()  {
	defer db.MyDb.Close()
	app:=initRouter()
	app.Run("127.0.0.1:8089")
}
