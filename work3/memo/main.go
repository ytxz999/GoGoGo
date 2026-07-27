package main

import (
	"memo/database"
	_ "memo/docs"
	"memo/router"
)

func main() {
	database.Init()

	r := router.SetupRouter()

	r.Run(":8080")
}
