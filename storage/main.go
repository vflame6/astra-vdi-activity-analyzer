package main

import (
	"github.com/vflame6/astra-vdi-activity-analyzer/storage/database"
	"github.com/vflame6/astra-vdi-activity-analyzer/storage/router"
	"github.com/vflame6/astra-vdi-activity-analyzer/storage/utils"
	"log"
)

func main() {
	conf, err := utils.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	if conf.Processing {
		log.Println("Starting processing ...")
		w := utils.NewWorker(conf.ProcessingAddress)
		w.Start()
		defer w.Shutdown()
	}

	_ = database.Init()

	r := router.InitRouter(conf.Password)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
