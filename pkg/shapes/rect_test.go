package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	r := Rectangle{2, 4}

	assert.Equal(t, 8.0, r.Area())
}

func TestPerimeter(t *testing.T) {
	r := Rectangle{2, 4}

	assert.Equal(t, 12.0, r.Perimeter())
}
