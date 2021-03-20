package config

import "os"

func GetEnviron() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "DEV"
	}
	return env
}

func GetDatabaseHost() string {
	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		host = "localhost"
	}
	return host
}

func GetDatabasePort() string {
	port := os.Getenv("DATABASE_HOST")
	if port == "" {
		port = "5432"
	}
	return port
}

func GetDatabaseName() string {

	name := os.Getenv("DATABASE_NAME")
	if name == "" {
		name = ""
	}
	return name
}

func GetDatabaseUser() string {

	user := os.Getenv("DATABASE_USER")
	if user == "" {
		user = "postgres"
	}
	return user
}

func GetDatabasePassword() string {

	password := os.Getenv("DATABASE_PASSWORD")
	if password == "" {
		password = "postgres"
	}
	return password
}
