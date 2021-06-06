package config_test

import (
	"os"
	"testing"

	"github.com/nurali-techie/rc/config"
	"github.com/stretchr/testify/assert"
)

func TestParamFromConfig(t *testing.T) {
	config.Init()
	assert.Equal(t, "", config.GetDBPath())
}

func TestParamFromEnv(t *testing.T) {
	env := os.Getenv("DB_PATH")
	os.Setenv("DB_PATH", "./out")
	defer os.Setenv("DB_PATH", env)

	config.Init()
	assert.Equal(t, "./out", config.GetDBPath())
}
