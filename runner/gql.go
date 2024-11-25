package runner

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/enuesaa/taskhop/gql/client"
)

func fetch() {
	c := client.NewClient(http.DefaultClient, "http://localhost:3000/graphql", nil)

	data, err := c.GetHealth(context.Background())
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	fmt.Printf("%+v", data.GetHealth)
}
