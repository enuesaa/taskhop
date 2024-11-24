package subscription

import (
	"context"
	"time"

	"github.com/enuesaa/taskhop/gql/model"
)

func (r *SubscriptionResolver) SubscribeRunCmdOutput(ctx context.Context) (<-chan *model.RunCmdOutput, error) {
	ch := make(chan *model.RunCmdOutput)

	go func() {
		defer close(ch)
		for {
			time.Sleep(1 * time.Second)

			res := &model.RunCmdOutput{
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
