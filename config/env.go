package config

import "os"

var (
	Environment = getEnvironment()
	DbHost      = getDatabaseHost()
	DbPort      = getDatabasePort()
	DbUser      = getDatabaseUser()
	DbName      = getDatabaseName()
	DbPassword  = getDatabasePassword()
	IdentityKey = getIdentityKey()
	SecretKey   = getSecretKey()
)

func getEnvVar(variable, defaultValue string) string {
	if envVar, ok := os.LookupEnv(variable); ok {
		return envVar
	}
	return defaultValue
}

func getEnvironment() string {
	return getEnvVar("ENV", "DEV")
}

func getDatabaseHost() string {
	return getEnvVar("DATABASE_HOST", "localhost")
}

func getDatabasePort() string {
	return getEnvVar("DATABASE_PORT", "5432")
}

func getDatabaseName() string {
	return getEnvVar("DATABASE_NAME", "tasktracker")
}

func getDatabaseUser() string {
	return getEnvVar("DATABASE_USER", "postgres")
}

func getDatabasePassword() string {
	return getEnvVar("DATABASE_PASSWORD", "postgres")
}

func getIdentityKey() string {
	return getEnvVar("IDENTITY_KEY", "id")
}

func getSecretKey() string {
	return getEnvVar("SECRET_KEY", "my_secret_key_3219kd")
}
