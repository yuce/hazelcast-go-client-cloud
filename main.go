package main

import (
	"context"
	"fmt"
	"github.com/hazelcast/hazelcast-go-client"
	"log"
)

func main() {
	config := hazelcast.NewConfig()
	cc := &config.ClusterConfig
	cc.Name = "XXXXX"
	cc.HazelcastCloudConfig.Enabled = true
	cc.HazelcastCloudConfig.Token = "XXXXXX"
	client, err := hazelcast.StartNewClientWithConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	m, err := client.GetMap(ctx, "my-ma[")

	for i := 0; i < 10_000; i++ {
		key := fmt.Sprintf("key-%d", i)
		value := fmt.Sprintf("value-%d", i)
		if err := m.Set(ctx, key, value); err != nil {
			log.Fatal(err)
		}
		readValue, err := m.Get(ctx, key)
		if err != nil {
			log.Fatal(err)
		}
		if value != readValue {
			log.Fatalf("set/got valeus differ: %s != %s", value, readValue)
		}
		log.Printf("OK key: %s, value: %s", key, value)
	}

	client.Shutdown()
}
