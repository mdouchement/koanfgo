package koanfgo_test

import (
	"os"
	"testing"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/mdouchement/koanfgo"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	cfg := []byte(`package config
	import (
		"maps"
		"os"
	)

	func Load() (map[string]interface{}, error) {
		config := map[string]interface{}{
			"brokers": []string{"localhost:42", "localhost:4242"},
			"listen":  "localhost",
			"redis": map[string]interface{}{
				"addr":     "localhost:6379",
				"password": "trololo",
				"db":       1,
			},
		}

		if os.Getenv("MAGIC_FEATURE") == "enabled" {
			maps.Set(config, "feature", "testouille")
		}

		return config, nil
	}`)

	//

	os.Setenv("MAGIC_FEATURE", "enabled")

	konf := koanf.New(".")
	err := konf.Load(rawbytes.Provider(cfg), koanfgo.Parser())
	assert.NoError(t, err)

	assert.Equal(t, []string{"localhost:42", "localhost:4242"}, konf.Strings("brokers"))
	assert.Equal(t, "localhost", konf.String("listen"))

	assert.Equal(t, "localhost:6379", konf.String("redis.addr"))
	assert.Equal(t, "trololo", konf.String("redis.password"))
	assert.Equal(t, 1, konf.Int("redis.db"))

	assert.Equal(t, "testouille", konf.String("feature"))
}
