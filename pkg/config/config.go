// package config contains configuration for waffy
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	// DefaultWAFListen is the default
	DefaultAPIListen = "0.0.0.0:8500"

	// DefaultCertPath is the default path to certificates
	DefaultCertPath = "./etc"

	// DefaultRPCName is the hostname of the RPC
	DefaultRPCName = "waffy.local"
)

// Config is the main configuration for waffy
type Config struct {
	// API Listener is the address the API should listen on
	APIListen string

	// CertPath is the path to certificates for the system
	CertPath string

	// RPCName is the hostname of the RPC client
	RPCName string
}

var cfg *Config

func init() {
	c, err := godotenv.Read()
	if err != nil {
		panic("cannot read configuration environment")
	}

	cfg = &Config{
		APIListen: getEnv("WAFFY_API_LISTEN", c, DefaultAPIListen),
		CertPath:  getEnv("WAFFY_CERT_PATH", c, DefaultCertPath),
		RPCName:   getEnv("WAFFY_RPC_NAME", c, DefaultRPCName),
	}
}

func getEnv(cfg string, c map[string]string, defValue string) string {
	val := os.Getenv(cfg)
	if val == "" {
		var ok bool
		if val, ok = c[cfg]; !ok {
			val = defValue
		}
	}

	return val
}

// Load returns the loaded configuration
func Load() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	return nil, fmt.Errorf("Error reading configuration")
}
