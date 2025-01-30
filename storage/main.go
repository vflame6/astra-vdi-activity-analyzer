package main

import (
	"github.com/vflame6/astra-vdi-activity-analyzer/storage/database"
	"github.com/vflame6/astra-vdi-activity-analyzer/storage/router"
)

func main() {
	_ = database.Init()

	r := router.InitRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
