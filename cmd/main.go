package main

import (
	"fmt"

	"github.com/KrizzMU/delivery-service/internal/db"
	"github.com/KrizzMU/delivery-service/internal/repository"
	"github.com/KrizzMU/delivery-service/internal/service"
	"github.com/KrizzMU/delivery-service/internal/transport/ns/sender"
	"github.com/KrizzMU/delivery-service/internal/transport/ns/subscriber"
	"github.com/KrizzMU/delivery-service/pkg/cache"
	"github.com/joho/godotenv"
)

func main() {
	natsURL := "stan://localhost:4222"
	clusterID := "alancid"
	channel := "Krizzis"
	subID := "subs"
	sendID := "send"

	//s := new(rest.Server)

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	c := cache.NewCache()

	repo := repository.NewRepository(db.GetConnection())
	services := service.NewService(repo, c)

	go services.Order.RecoveryCache(repo.Order.GetAll())

	sender := sender.NewSender(natsURL, clusterID, sendID)
	sub := subscriber.NewSub(natsURL, clusterID, subID, services)

	go sender.SendFake(channel)
	//time.Sleep(time.Second * 5)
	fmt.Println("Started listening to the channel")
	go sub.SubToChannel(channel)

	select {}
}
