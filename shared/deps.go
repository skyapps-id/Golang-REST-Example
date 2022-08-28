package shared

import (
	"golang-rest-example/shared/config"
)

type (
	Deps struct {
		Config *config.Configuration
	}
)

func NewDeps(config *config.Configuration) Deps {
	return Deps{Config: config}
}
