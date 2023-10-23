package sender

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/KrizzMU/delivery-service/internal/core"
	"github.com/go-faker/faker/v4"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type Sender struct {
	sc stan.Conn
}

func NewSender(natsURL, clusterID, clientID string) *Sender {
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

	return &Sender{sc: sc}
}

func (s *Sender) SendFake(subject string) {
	for {
		var order core.Order
		fakedate := faker.FakeData(&order)
		if fakedate != nil {
			fmt.Println(fakedate)
			continue
		}

		order.Payment.Transaction = order.OrderUID
		if len(order.Items) > 3 {
			order.Items = order.Items[:3]
		}
		for i := 0; i < len(order.Items); i++ {
			order.Items[i].TrackNumber = order.TrackNumber
		}

		fmt.Println("Sended Order: OrderUID = ", order.OrderUID)

		jsondata, err := json.Marshal(order)
		if err != nil {
			fmt.Println(err)
			continue
		}
		s.sc.Publish(subject, jsondata)

		time.Sleep(time.Second * 120)
	}
}
