package main

import (
	"fmt"
	"github.com/kuhotkin/jira-connector/db/postgres"
	"github.com/kuhotkin/jira-connector/pkg/handler"
	"github.com/kuhotkin/jira-connector/pkg/service"
	jira_connector "github.com/kuhotkin/jira-connector/rest"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.dbHost"),
		Port:     viper.GetString("db.dbPort"),
		User:     viper.GetString("db.dbUser"),
		DbName:   viper.GetString("db.dbName"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	jiraConnector := service.New(db)

	hand := handler.NewHandler(jiraConnector)

	hand.InitHandleFuncs()

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	server := new(jira_connector.Server)
	fmt.Println("Server started")
	if err := server.Run(viper.GetString("port")); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("error occured on db connection close: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
