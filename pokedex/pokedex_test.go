package pokedex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPokedex(t *testing.T) {
	pokedex, err := GetPokedex()
	assert.Nil(t, err)
	assert.Equal(t, "フシギダネ", pokedex[0].Name)
}
