# Koanf Go(lang)

Koanf Go(lang) allows to write dynamic configuration files for [koanf](https://github.com/knadh/koanf) configuration manager. Sometimes dynamic configuration is required.

Based on the idea of https://github.com/mdouchement/koanflua.

## Usage

```go
// config.go
package config

import (
	"maps" // https://github.com/containous/yaegi/issues/327
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
}
```

```go
package main

import (
	"fmt"
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/file"
	"github.com/mdouchement/koanfgo"
)

func main() {
	konf := koanf.New(".")
	err := konf.Load(file.Provider("config.go"), koanfgo.Parser())
	if err != nil {
		log.Fatal(err)
	}

	//

	redisAddr := konf.String("redis.addr")
	fmt.Println(redisAddr)
}
```

## Resources

- Yaegi is Another Elegant Go Interpreter: https://github.com/containous/yaegi

## License

**MIT**


## Contributing

All PRs are welcome.

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request
