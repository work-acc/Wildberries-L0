package app

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/caarlos0/env/v6"

	"work-acc/wildberries-L0/internal/config"
	"work-acc/wildberries-L0/internal/server"
	"work-acc/wildberries-L0/internal/service"
	"work-acc/wildberries-L0/internal/storage/pdb"
	api "work-acc/wildberries-L0/internal/transport/api"
	nats "work-acc/wildberries-L0/pkg/client/nats-streaming"
	postgresql "work-acc/wildberries-L0/pkg/client/postgreSQL"
)

func Run() {
	data, err := os.ReadFile("./configs/config.json")
	if err != nil {
		log.Fatalf("error for read config.json: %v", err)
	}

	cfg := new(config.Config)
	if err := json.Unmarshal(data, cfg); err != nil {
		log.Fatalf("error for read config.json: %v", err)
	}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("error for parse env: %v", err)
	}

	router := http.NewServeMux()
	server := new(server.Server)

	postgreClient, err := postgresql.NewPostgresDB(&cfg.PostgreSQL)
	if err != nil {
		log.Fatalf("error for connection to DB: %v\n", err)

		return
	}
	defer postgreClient.Close()

	sc, err := nats.NewNatsConnect(cfg)
	if err != nil {
		log.Fatalf("error for connection to nats-streaming: %v\n", err)

		return
	}
	defer sc.Close()

	storageOrder := pdb.NewStorageOrder(postgreClient)
	serviceOrder := service.NewServiceOrder(storageOrder)
	serviceNats := service.NewServiceNatsStreaming(storageOrder)

	if err := serviceOrder.RecoveryInMemory(); err != nil {
		log.Fatalf("error for recovery cash: %v\n", err)
	}

	serviceNats.Subscribe(sc)

	handlerOrder := api.NewHandlerOrder(serviceOrder)
	handlerOrder.Init(router)

	// Publishing data to a channel
	//Publish()

	if err := server.Run(cfg.Api.Port, router); err != nil {
		log.Fatalf("error for run server: %v\n", err)

		return
	}
}
