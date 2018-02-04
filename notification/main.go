package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := parseConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "(CONFIG) ERROR:", err)
		os.Exit(1)
	}
	db, err := connectToPostgre(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "(POSTGRESQL) ERROR:", err)
		os.Exit(1)
	}
	redisClient := connectToRedis(config)
	metricsNotification(redisClient, config, db)
}
