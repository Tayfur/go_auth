package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	return os.Getenv(key)
}

func PortConfig(key string) uint64 {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	p := os.Getenv(key)
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println(err)
	}
	return port
}
