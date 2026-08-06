package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mohammedimrankasab/go-design-patterns/behavioral/strategy"
)

func main() {
	ctx := context.Background()

	connectors := []strategy.Connector{
		strategy.PowerBIConnector{},
		strategy.TableauConnector{},
		strategy.MLflowConnector{},
	}

	for _, connector := range connectors {
		processor, err := strategy.NewProcessor(connector)
		if err != nil {
			log.Fatal(err)
		}

		metadata, err := processor.Process(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n%s Connector\n", connector.Name())
		fmt.Println("-----------------------------")

		for _, item := range metadata {
			fmt.Printf(
				"ID: %-15s Name: %-25s Type: %s\n",
				item.ID,
				item.Name,
				item.Type,
			)
		}
	}
}
