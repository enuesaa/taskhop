package taskfx

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func (i *TaskSrv) Validate(task Task) error {
	err := validator.New().Struct(task)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%s: %s", err.Field(), err.Tag())
		}
	}
	return nil
}
