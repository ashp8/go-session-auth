package main

import (
	conn "com.ashp8/connection"
	"com.ashp8/router"
)

func main() {
	r := router.SetupRoutes()

	db, err := conn.SetupDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r.Run(":5000")
}
