package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// init env configs
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found on server")
	}
}

func APPName() string {
	return os.Getenv("APP_NAME")
}

func APPDomain() string {
	return os.Getenv("APP_DOMAIN")
}

func APPPort() string {
	return os.Getenv("APP_PORT")
}

func APPLogin() string {
	return os.Getenv("APP_LOGIN")
}

func APPPassw() string {
	return os.Getenv("APP_PASSW")
}

func DBHost() string {
	return os.Getenv("DB_HOST")
}

func DBName() string {
	return os.Getenv("DB_NAME")
}

func DBUser() string {
	return os.Getenv("DB_USER")
}

func DBPass() string {
	return os.Getenv("DB_PASS")
}

func REDISPass() string {
	return os.Getenv("REDIS_PASS")
}

func REDISHost() string {
	return os.Getenv("REDIS_HOST")
}

func BBBHost() string {
	return os.Getenv("BBB_HOST_URL")
}

func BBBKey() string {
	return os.Getenv("BBB_API_KEY")
}

func COOKIEHash() string {
	return os.Getenv("HASH_SALT")
}