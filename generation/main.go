package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/naoina/toml"
)

type Config struct {
	Database struct {
		User     string
		Password string
		Name     string
		Host     string
	}
}

var devices []int

func parseConfig() (Config, error) {
	var config Config
	f, err := os.Open("config.toml")
	if err != nil {
		return config, err
	}
	defer f.Close()
	if err := toml.NewDecoder(f).Decode(&config); err != nil {
		return config, err
	}
	return config, nil
}

func connectToDB(config Config) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		config.Database.User, config.Database.Password, config.Database.Name, config.Database.Host)
	db, err := sql.Open("postgres", dbinfo)
	fmt.Println(dbinfo)
	if err != nil {
		return db, err
	}
	return db, nil
}

func generateMetrics(config Config, db *sql.DB) {
	queryDevices := "SELECT id FROM metrics.public.devices"
	queryMetrics := `INSERT INTO
						device_metrics(device_id,metric_1,metric_2,metric_3,metric_4, metric_5, local_time) 
						VALUES($1,$2,$3,$4,$5,$6,$7);`

	rows, err := db.Query(queryDevices)
	if err != nil {
		fmt.Fprintln(os.Stderr, "(POSTGRESQL) ERROR:", err)
	}
	var v int
	for rows.Next() {
		err = rows.Scan(&v)
		devices = append(devices, v)
	}

	for i := 0; i <= devices[rand.Intn(len(devices))]; i++ {
		_, err := db.Exec(queryMetrics,
			devices[rand.Intn(len(devices))],
			rand.Int31(),
			rand.Int31(),
			rand.Int31(),
			rand.Int31(),
			rand.Int31(),
			time.Now())
		if err != nil {
			fmt.Fprintln(os.Stderr, "(POSTGRESQL) ERROR:", err)
		}
	}
}
func main() {
	config, err := parseConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "(CONFIG) ERROR:", err)
		os.Exit(1)
	}
	db, err := connectToDB(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "(POSTGRESQL) ERROR:", err)
		os.Exit(1)
	}
	for {
		fmt.Println("generating metrics")
		generateMetrics(config, db)
		time.Sleep(5 * time.Second)
	}
}
