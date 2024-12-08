package archivefx

import (
	"testing"

	"github.com/enuesaa/taskhop/lib/archivefx/repository"
	"go.uber.org/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestArchive(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRepo := repository.NewMockI(ctl)
	mockRepo.EXPECT().IsDir("/").Return(true, nil)

	is, err := mockRepo.IsDir("/")
	assert.NoError(t, err)
	assert.Equal(t, true, is)
}
