package cmdsfile

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func (i *Impl) Validate(cmdsfile CmdsFile) error {
	err := validator.New().Struct(cmdsfile)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%s: %s", err.Field(), err.Tag())
		}
	}

	return nil
}
