package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	cfg := Init("./config.example.yaml")
	assert.True(t, cfg.Default.Enable, "they should be true")
}
