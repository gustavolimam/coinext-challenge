package environment

import (
	"fmt"
	"os"
)

const (
	Port = "PORT"
	// Database envs
	DBUser      = "DB_USER"
	DBPassword  = "DB_PASSWORD"
	DBHost      = "DB_HOST"
	DBName      = "DB_NAME"
	DBSuperUser = "DB_SUPER_USER"
)

func CheckEnvVars() {
	envVars := []string{
		Port,
		DBUser,
		DBPassword,
		DBHost,
		DBName,
		DBSuperUser,
	}

	for _, v := range envVars {
		if os.Getenv(v) == "" {
			panic(fmt.Sprintf("env variable %s must be defined", v))
		}
	}
}
