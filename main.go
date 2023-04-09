package main

import (
	"go-jwt/database"
	"go-jwt/router"
	"os"
)

func main() {	
	database.StartDB()
	var PORT = os.Getenv("PORT")
	r:= router.StartApp()
	r.Run(":" + PORT)
}