package main

import (
	"memo/database"
	"memo/router"
)

func main() {
	database.Init()

	r := router.SetupRouter()

	r.Run(":8080")
}
