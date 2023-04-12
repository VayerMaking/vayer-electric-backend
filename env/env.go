package env

// Load local env variables
import (
	_ "github.com/joho/godotenv/autoload"
)

var STATSD_FLUSH = getOptionalEnvAsInt("STATSD_FLUSH", 300)
var SHUTDOWN_TIMEOUT = getOptionalEnvAsInt("SHUTDOWN_TIMEOUT", 30)
var REQUEST_TIMEOUT = getOptionalEnvAsInt("REQUEST_TIMEOUT", 10)
var PORT = getOptionalEnvAsInt("PORT", 8080)
var DB_HOST = getOptionalEnv("DB_HOST", "localhost")
var DB_PORT = getOptionalEnvAsInt("DB_PORT", 5432)
var DB_USER = getOptionalEnv("DB_USER", "vayer-electric")
var DB_PASSWORD = getOptionalEnv("DB_PASSWORD", "vayer-electric")
var DB_NAME = getOptionalEnv("DB_NAME", "vayer-electric")
