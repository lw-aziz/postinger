package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// ServerConfig - Server configuration
type ServerConfig struct {
	AppName         string
	Port            int
	ServerUri       string
	AllowOrigins    []string
	LogLevel        string
	DefaultLanguage string
	Languages       []string
	APIGateway      string
	APIGatewayAuth  string
}

// GRPCConfig - GRPC configuration
type GRPCConfig struct {
	Secure   bool
	CertFile string
	KeyFile  string
	Port     string
}

// MongoDBConfig - MongoDB configuration
type MongoDBConfig struct {
	URL      string
	Database string
}

// MqttConfig - Mqtt configuration
type MqttConfig struct {
	URL      string
	User     string
	Password string
}

// Config structure
type Config struct {
	Server  ServerConfig
	GRPC    GRPCConfig
	MongoDB MongoDBConfig
	Mqtt    MqttConfig
}

// AppConfig - Appconfig object
var AppConfig = &Config{
	Server: ServerConfig{
		AppName:         "",
		Port:            50053,
		ServerUri:       "",
		AllowOrigins:    []string{"*"},
		LogLevel:        "info",
		DefaultLanguage: "en",
		Languages:       []string{"en"},
		APIGateway:      "",
		APIGatewayAuth:  "",
	},
	GRPC: GRPCConfig{
		Secure:   false,
		CertFile: "",
		KeyFile:  "",
		Port:     "50052",
	},
	MongoDB: MongoDBConfig{
		URL:      "",
		Database: "",
	},
	Mqtt: MqttConfig{
		URL:      "",
		User:     "",
		Password: "",
	},
}

// LoadEnv - function load Enviroment variable from .env file
func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := &Config{
		Server: ServerConfig{
			AppName:         getEnv("APP_NAME", ""),
			Port:            getEnvAsInt("API_PORT", 50053),
			ServerUri:       getEnv("SERVER_URI", ""),
			AllowOrigins:    strings.Split(getEnv("ALLOW_ORIGIN", "*"), ","),
			LogLevel:        getEnv("LOG_LEVEL", "info"),
			DefaultLanguage: getEnv("DEFAULT_LANGUAGE", "en"),
			Languages:       strings.Split(getEnv("LANGUAGES", "en"), ","),
			APIGateway:      getEnv("API_GATEWAY", ""),
			APIGatewayAuth:  getEnv("API_GATEWAY_AUTH", ""),
		},
		GRPC: GRPCConfig{
			Secure:   getEnvAsBool("GRPC_SECURE", false),
			CertFile: getEnv("GRPC_CERT_NAME", ""),
			KeyFile:  getEnv("GRPC_KEY_NAME", ""),
			Port:     getEnv("GRPC_PORT", "50052"),
		},
		MongoDB: MongoDBConfig{
			URL:      getEnv("MONGODB_URL", ""),
			Database: getEnv("MONGODB_DATABASE", ""),
		},
		Mqtt: MqttConfig{
			URL:      getEnv("MQTT_URL", ""),
			User:     getEnv("MQTT_USERNAME", ""),
			Password: getEnv("MQTT_PASSWORD", ""),
		},
	}
	AppConfig = config
	return config
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// // Helper to read an environment variable into a string slice or return default value
// func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
// 	valStr := getEnv(name, "")

// 	if valStr == "" {
// 		return defaultVal
// 	}

// 	val := strings.Split(valStr, sep)

// 	return val
// }
