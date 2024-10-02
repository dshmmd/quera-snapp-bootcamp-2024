package config

type Config struct {
	// BindAddress is address on which application will try
	// to serve HTTP Server
	//
	// Example: 0.0.0.0:8080
	//
	BindAddress string `yaml:"bind_address"`
}
