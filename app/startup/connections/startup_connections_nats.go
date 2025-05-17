package connections

import (
	"log"
	"wrench/app/manifest/connection_settings"

	"github.com/nats-io/nats.go"
)

var connections map[string]*nats.Conn

func GetNatsConnectionById(natsConnectionId string) *nats.Conn {
	if len(natsConnectionId) == 0 || connections == nil {
		return nil
	}

	return connections[natsConnectionId]
}

func loadConnectionNats(connNatsSetting []*connection_settings.ConnectionNatsSettings) error {
	if len(connNatsSetting) > 0 {
		if connections == nil {
			connections = make(map[string]*nats.Conn)
		}

		for _, conn := range connNatsSetting {
			nc, err := nats.Connect(conn.ServerAddress)

			if err != nil {
				log.Printf("Error nats connection: %v", err)
				return err
			}

			connections[conn.Id] = nc
		}
	}

	return nil
}
