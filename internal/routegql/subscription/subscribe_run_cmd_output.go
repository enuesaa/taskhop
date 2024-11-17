package subscription

import (
	"context"
	"time"

	"github.com/enuesaa/taskhop/internal/routegql/schema"
)

func (r *SubscriptionResolver) SubscribeRunCmdOutput(ctx context.Context) (<-chan *schema.RunCmdOutput, error) {
	ch := make(chan *schema.RunCmdOutput)

	go func() {
		defer close(ch)
		for {
			time.Sleep(1 * time.Second)

			res := &schema.RunCmdOutput{
				Output: "text",
			}
			select {
			case <-ctx.Done():
				return
			case ch <- res:
			}
		}
	}()

	return ch, nil
}
