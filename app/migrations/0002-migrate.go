package mixtures

import (
	"github.com/ezn-go/mixture"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
)

func init() {
	mx := &gormigrate.Migration{
		ID:       "0009",
		Migrate:  mixture.CreateTableM(&service.Payment{}),
		Rollback: mixture.DropTableR(&service.Payment{}),
	}

	mixture.Add(mixture.ForAnyEnv, mx)
}
