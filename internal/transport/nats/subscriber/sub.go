package subscriber

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type Subscriber struct {
	sc stan.Conn
}

func NewSub(natsURL, clusterID, clientID string) *Subscriber {
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

	return &Subscriber{sc: sc}
}

func (s *Subscriber) SubToChannel(subject string) error {
	_, err := s.sc.Subscribe(subject, func(msg *stan.Msg) {
		handleOrderMessage(msg.Data)
		msg.Ack()
	}, stan.DeliverAllAvailable(), stan.DurableName("dur"))

	if err != nil {
		return err
	}

	return nil
}

func handleOrderMessage(msg []byte) error {

	var orderData map[string]interface{}
	err := json.Unmarshal(msg, &orderData)
	if err != nil {
		fmt.Printf("JSON Error: %v", err)
		return err
	}

	jsonData, err := json.MarshalIndent(orderData, "", "  ") // Создаем отформатированную JSON-строку
	if err != nil {
		fmt.Printf("Ошибка при маршалинге JSON: %v", err)
		return err
	}

	fmt.Println(string(jsonData))
	fmt.Println()
	fmt.Println("Прочитал сообщение!")
	fmt.Println()

	return nil
}
