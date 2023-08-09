package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURL string
	SecretKey  []byte
	ServerPort string
	UserDBName string
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		return
	}
}

func GetConfig() *Config {
	return &Config{
		MongoDBURL: getDBURL(),
		SecretKey:  getSecretKey(),
		ServerPort: os.Getenv("SERVER_PORT"),
		UserDBName: getDBName("USER_DB"),
	}
}

func getSecretKey() []byte {
	return []byte(os.Getenv("SECRET_KEY"))
}

func getDBName(dburl string) string {
	url := os.Getenv(dburl)
	rmv := strings.Replace(url, "//", "", 1)
	split := strings.Split(rmv, "/")

	if len(split) < 2 {
		return ""
	}
	dbname := split[1]
	return dbname
}

func getDBURL() string {
	dburl := os.Getenv("MONGODB_URI")
	if getDBName(dburl) == "" {
		return dburl
	}
	re := regexp.MustCompile(`\/[^/]+$`)
	result := re.ReplaceAllString(dburl, "")
	return result
}
