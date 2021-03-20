package config

import "os"

var (
	Environment = getEnviron()
	DbHost      = getDatabaseHost()
	DbPort      = getDatabasePort()
	DbUser      = getDatabaseUser()
	DbName      = getDatabaseUser()
	DbPassword  = getDatabasePassword()
	IdentityKey = getIdentityKey()
	SecretKey   = getSecretKey()
)

func getEnviron() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "DEV"
	}
	return env
}

func getDatabaseHost() string {
	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		host = "localhost"
	}
	return host
}

func getDatabasePort() string {
	port := os.Getenv("DATABASE_HOST")
	if port == "" {
		port = "5432"
	}
	return port
}

func getDatabaseName() string {

	name := os.Getenv("DATABASE_NAME")
	if name == "" {
		name = ""
	}
	return name
}

func getDatabaseUser() string {

	user := os.Getenv("DATABASE_USER")
	if user == "" {
		user = "postgres"
	}
	return user
}

func getDatabasePassword() string {

	password := os.Getenv("DATABASE_PASSWORD")
	if password == "" {
		password = "postgres"
	}
	return password
}

func getIdentityKey() string {

	identityKey := os.Getenv("IDENTITY_KEY")
	if identityKey == "" {
		identityKey = "id"
	}
	return identityKey
}

func getSecretKey() string {

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "my_secret_key_3219kd"
	}
	return secretKey
}
