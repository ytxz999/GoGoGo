package main

import (
	"fmt"
	"memo/database"
	_ "memo/docs"
	"memo/router"
)

func main() {
	database.Init()

	r := router.SetupRouter()

	err := r.Run(":8080")
	if err != nil {
		fmt.Print(err)
		return
	}
}
