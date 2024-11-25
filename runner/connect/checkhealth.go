package connect

import (
	"context"
	"fmt"
	"net/http"

	"github.com/enuesaa/taskhop/runner/client"
)

func checkHealth(address string) error {
	url := fmt.Sprintf("http://%s/graphql", address)
	cli := client.NewClient(http.DefaultClient, url, nil)
	
	data, err := cli.GetHealth(context.Background())
	if err != nil {
		return err
	}

	if !data.Health.Ok {
		return fmt.Errorf("commander not healthy")
	}
	return nil
}
