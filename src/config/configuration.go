package config

type Configuration struct {
	MongoIp       string `env:"STORAGESERVICE_MONGO_IP" required:"true"`
	MongoPort     string `env:"STORAGESERVICE_MONGO_PORT" default:"27017"`
	MongoUser     string `env:"STORAGESERVICE_MONGO_USER" required:"true"`
	MongoPassword string `env:"STORAGESERVICE_MONGO_PASSWORD" required:"true"`
	MongoDB       string `env:"STORAGESERVICE_MONGO_DB" required:"true"`
	Port          string `env:"STORAGESERVICE_PORT" default:"80"`
}
