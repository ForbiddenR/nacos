package options

import (
	"github.com/caarlos0/env"
	"github.com/go-playground/validator/v10"
)

type Options struct {
	Url         string `env:"NACOS_HOST" validate:"required"`
	Port        uint64 `validate:"required"`
	NamespaceId string `env:"NACOS_NAMESPACE" validate:"required"`
	GroupName   string `env:"NACOS_GROUP" validate:"required"`
	Username    string `env:"NACOS_USERNAME" validate:"required"`
	Password    string `env:"NACOS_PASSWORD" validate:"required"`
	DataId      string `env:"NACOS_CONFIG_NAME" validate:"required"`
}

// InitConfig reads nacos configuration from environmental variables.
func InitConfig(validate *validator.Validate, port *uint64) (*Options, error) {
	op := &Options{}
	if err := env.Parse(op); err != nil {
		return nil, err
	}
	if port != nil {
		op.Port = *port
	} else {
		// default port of naocs.
		op.Port = uint64(8848)
	}

	if err := validate.Struct(op); err != nil {
		return nil, err
	}

	return op, nil
}
