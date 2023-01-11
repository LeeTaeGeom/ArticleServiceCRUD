package main

import (
	"articleCRUD/api"
	"articleCRUD/database"
)

func init() {
	database.NewPostgreSQLClient()
}
func main() {

	r := api.SetupRouter()
	r.Run(":8000")
}
