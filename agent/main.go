package main

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/vflame6/astra-vdi-activity-analyzer/agent/http"
	"github.com/vflame6/astra-vdi-activity-analyzer/agent/utils"
	"github.com/vflame6/astra-vdi-activity-analyzer/agent/worker"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	register = kingpin.Flag("register", "Register the agent").Short('r').Bool()
	noSender = kingpin.Flag("offline", "Offline mode. This will just make screenshots and save them to data/ directory.").Bool()
)

func main() {
	kingpin.Version("1.0.0")

	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	if *register && *noSender {
		log.Fatal("You can't set --register and --offline at the same time.")
	}

	conf, err := utils.LoadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	serverURL := http.GetURL(conf.Address, conf.UseTLS)

	if *register {
		log.Println("Registering agent")

		log.Println("Enter password for storage server: ")
		pwd, err := utils.ReadPassword()
		if err != nil {
			log.Fatal(err)
		}

		key, err := worker.Register(serverURL, conf.ClientName, pwd)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Saving the configuration file to ./config.json")
		newConf := &utils.Config{
			ClientName: conf.ClientName,
			Address:    conf.Address,
			UseTLS:     conf.UseTLS,
			Key:        key,
		}
		err = utils.SaveConfig("./config.json", newConf)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Successfully registered agent")
		os.Exit(0)
	}

	if conf.Key == "" && !*noSender {
		log.Fatal("The agent is not registered. Try again with --register.")
	}

	agent := worker.NewAgent(conf, serverURL, *noSender)

	if !*noSender {
		log.Println("Starting agent in online mode")
		log.Println("Checking storage server availability")
		err = worker.Ping(serverURL)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Checking agent authenticity")
		err = agent.HealthCheck()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Starting agent in offline mode")
	}
	go agent.Start()

	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Println("CTRL+C received... Gracefully shutting down the agent")
	agent.Stop()
	os.Exit(0)
}
