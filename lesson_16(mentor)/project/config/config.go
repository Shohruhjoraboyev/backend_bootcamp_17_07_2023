package config

type Config struct {
	Limit   int
	Page    int
	Methods []string
	Objects []string
}

func Load() *Config {
	return &Config{
		Limit: 10,
		Page:  1,
		Methods: []string{
			"create", "update", "get", "getAll", "update", "delete",
		},
		Objects: []string{
			"branch", "staff", "sales", "transaction", "tariff",
		},
	}
}
