package settings

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Setting env...
type Setting struct {
	Name    string `env:"NAME"`
	Version string `env:"VERSION"`
}

// Build ...
func Build(file ...string) *Setting {
	settings := new(Setting)

	godotenv.Load(file...)
	err := env.Parse(settings)

	if err != nil {
		fmt.Println(err)
	}

	return settings
}
