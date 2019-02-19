package blueprint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBootstrappingServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewBootstrappingService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

// TODO
