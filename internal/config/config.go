package config

type Config struct {
	Api        Api        `json:"api"`
	Nats       Nats       `json:"nats"`
	PostgreSQL PostgreSQL `json:"postgres"`
}

type Api struct {
	Port int `json:"port"`
}

type Nats struct {
	URL       string `json:"url"`
	ClusterID string `json:"cluster_id"`
	ClientID  string `json:"client_id"`
}

type PostgreSQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"db_name"`
	SslMode  string `json:"sslmode"`
	User     string `env:"POSTGRES_USER,notEmpty"`
	Password string `env:"POSTGRES_PASSWORD,notEmpty"`
}
