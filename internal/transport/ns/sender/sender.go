package sender

import (
	"log"

	//ns "github.com/KrizzMU/delivery-service/internal/transport/nats"
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
	//fakedate := ns.Order{}

	// jsonData, err := os.ReadFile("./model.json")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// for i := 0; i < 3; i++ {
	// 	err = s.sc.Publish(subject, jsonData)
	// 	fmt.Println("Json Sended", i)
	// 	//time.Sleep(time.Second * 3)
	// }

}
