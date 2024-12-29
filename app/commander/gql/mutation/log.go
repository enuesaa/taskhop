package mutation

import (
	"context"
	"fmt"

	"github.com/enuesaa/taskhop/app/commander/gql/model"
)

func (r *MutationResolver) Log(ctx context.Context, input model.LogInput) (bool, error) {
	if input.Type == model.LogTypeCommand {
		r.Lib.Log.Faint(context.Background(), "")
		text := fmt.Sprintf("> %s", input.Output)
		r.Lib.Log.Faint(context.Background(), text)
		return true, nil
	}

	r.Lib.Log.Info(context.Background(), input.Output)

	return true, nil
}
