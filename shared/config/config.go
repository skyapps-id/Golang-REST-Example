package config

import cfg "golang-rest-example/utils/config"

type (
	Configuration struct {
		// - MySQL database
		DbUser            string `envconfig:"DB_USER" required:"true" default:"user"`
		DbPass            string `envconfig:"DB_PASS" required:"true" default:"test"`
		DbName            string `envconfig:"DB_NAME" required:"true" default:"example"`
		DbHost            string `envconfig:"DB_HOST" required:"true" default:"localhost"`
		DbPort            int    `envconfig:"DB_PORT" required:"true" default:"3306"`
		EnableLogSqlQuery bool   `envconfig:"LOG_SQL_QUERY" default:"false"`
	}
)

func New() (*Configuration, error) {
	var (
		c = Configuration{}
	)

	if err := cfg.New(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
