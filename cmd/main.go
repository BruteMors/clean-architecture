package main

import (
	"awesomeProject1/internal/composites"
	"awesomeProject1/internal/config"
	"awesomeProject1/pkg/logging"
	"context"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("config init")
	cfg := config.GetConfig()

	logger.Info("router init")
	router := httprouter.New()

	logger.Info("create mongo db composite")
	mongoDBC, err := composites.NewMongoDBComposite(context.Background(), cfg.MongoDB.Host, "", "", "", "", "")
	if err != nil {
		logger.Fatal("mongo db composite failed")
	}

	logger.Info("create author composite")
	authorComposite, err := composites.NewAuthorComposite(mongoDBC)
	if err != nil {
		logger.Fatal("author composite failed")
	}
	authorComposite.Handler.Register(router)

	logger.Info("create book composite")
	bookComposite, err := composites.NewAuthorComposite(mongoDBC)
	if err != nil {
		logger.Fatal("book composite failed")
	}

	bookComposite.Handler.Register(router)
}
