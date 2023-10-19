package subscriber

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/KrizzMU/delivery-service/internal/core"

	"github.com/KrizzMU/delivery-service/internal/service"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type Subscriber struct {
	sc       stan.Conn
	services *service.Service
}

func NewSub(natsURL, clusterID, clientID string, s *service.Service) *Subscriber {
	nc, err := nats.Connect(natsURL)

	if err != nil {
		log.Fatal(err.Error())
	}

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))

	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, natsURL)
	}

	return &Subscriber{
		sc:       sc,
		services: s,
	}
}

func (s *Subscriber) SubToChannel(subject string) error {
	_, err := s.sc.Subscribe(subject, func(msg *stan.Msg) {
		s.handleOrderMessage(msg.Data)
		msg.Ack()
	}, stan.DeliverAllAvailable(), stan.DurableName("dur"))

	if err != nil {
		return err
	}

	return nil
}

func (s *Subscriber) handleOrderMessage(msg []byte) error {

	var orderData core.Order
	err := json.Unmarshal(msg, &orderData)
	if err != nil {
		fmt.Printf("JSON Error: %v", err)
		return err
	}

	fmt.Println()
	fmt.Println("Прочитал сообщение!")
	fmt.Println()

	if err := s.services.Order.Create(orderData); err != nil {
		fmt.Printf("Created Error: %v", err)
	}

	return nil
}
