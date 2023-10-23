package main

import (
	"fmt"
	"log"

	"github.com/KrizzMU/delivery-service/internal/db"
	"github.com/KrizzMU/delivery-service/internal/repository"
	"github.com/KrizzMU/delivery-service/internal/service"
	"github.com/KrizzMU/delivery-service/internal/transport/ns/sender"
	"github.com/KrizzMU/delivery-service/internal/transport/ns/subscriber"
	"github.com/KrizzMU/delivery-service/internal/transport/rest"
	"github.com/KrizzMU/delivery-service/internal/transport/rest/handlers"
	"github.com/KrizzMU/delivery-service/pkg/cache"
	"github.com/joho/godotenv"
)

func main() {
	natsURL := "stan://localhost:4222"
	clusterID := "alancid"
	channel := "Krizzis"
	subID := "subs"
	sendID := "send"

	sender := sender.NewSender(natsURL, clusterID, sendID)
	fmt.Println("Started sending to channel")
	go sender.SendFake(channel)

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	c := cache.NewCache()

	repo := repository.NewRepository(db.GetConnection())
	services := service.NewService(repo, c)
	go services.Order.RecoveryCache(repo.Order.GetAll())
	handler := handlers.NewHandler(services)

	sub := subscriber.NewSub(natsURL, clusterID, subID, services)

	fmt.Println("Started listening to the channel")
	go sub.SubToChannel(channel)

	s := new(rest.Server)

	if err := s.Run("8088", handler.InitRoutes()); err != nil {
		log.Fatal("ERROR start Server!")
	}
}
