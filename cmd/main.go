package main

import (
	"biling"
	"biling/configs"
	"biling/database"
	"biling/internal/handler"
	"biling/internal/repository"
	"biling/internal/service"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := InitConfigs(); err != nil {
		log.Fatal("error while reading config file due to: ", err.Error())
	}

	var cfg configs.DatabaseConnConfig

	if err := viper.UnmarshalExact(&cfg); err != nil {
		log.Fatal("could not unmarshal the config into struct due to: ", err.Error())
	}

	cfg.Password = os.Getenv("DB_PASSWORD")

	conn, err := database.GetDBConnection(cfg)
	if err != nil {
		log.Fatalf("erro while opening DB. error is: %s", err.Error())
	}

	//----------Dependency Injection--------------
	newRepository := repository.NewRepository(conn)
	newService := service.NewService(newRepository)
	newHandler := handler.NewHandler(newService.Acc)
	//----------------------------------------------

	server := new(biling.Server)
	go func() {
		if err := server.Run(os.Getenv("PORT"), newHandler.InitRoutes()); err != nil {
			log.Fatalf("error while running http.server. error is: %s", err.Error())
		}
	}()
	fmt.Printf("Server is listening to port %s\n", os.Getenv("PORT"))

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	if err := conn.Close(); err != nil {
		log.Fatalf("error while closing DB. Error is: %s\n", err.Error())
	}

	fmt.Println("server is shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("error while shutting server down. Error is: %s\n", err.Error())
	}
}

func InitConfigs() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
