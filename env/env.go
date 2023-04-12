package env

// Load local env variables
import (
	_ "github.com/joho/godotenv/autoload"
)

var STATSD_FLUSH = getOptionalEnvAsInt("STATSD_FLUSH", 300)
var SHUTDOWN_TIMEOUT = getOptionalEnvAsInt("SHUTDOWN_TIMEOUT", 30)
var REQUEST_TIMEOUT = getOptionalEnvAsInt("REQUEST_TIMEOUT", 10)
var PORT = getOptionalEnvAsInt("PORT", 8080)
