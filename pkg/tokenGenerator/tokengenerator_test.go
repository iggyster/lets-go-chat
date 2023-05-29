package tokengenerator_test

import (
	"testing"

	"github.com/iggyster/lets-go-chat/pkg/tokengenerator"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	token := tokengenerator.Generate(16)

	assert.Equal(t, 32, len(token))
}
