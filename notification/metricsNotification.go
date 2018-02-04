package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type Metrics struct {
	Metric_1 int
	Metric_2 int
	Metric_3 int
	Metric_4 int
	Metric_5 int
}

func metricsNotification(redisClient *redis.Client, config Config, db *sql.DB) {
	baseQuery := ` SELECT 
					coalesce(metric_1,0),
					coalesce(metric_2, 0),
					coalesce(metric_3,0),
					coalesce(metric_4,0),
					coalesce(metric_5, 0),
					device_id,
					local_time 
				FROM device_metrics`

	var deviceID int
	allQueries := []Metrics{}

	ticker := time.Tick(1 * time.Second)
	var lastTime time.Time
	for now := range ticker {
		query := baseQuery
		if !lastTime.IsZero() {
			query = query + " WHERE local_time > '" + lastTime.Format(time.RFC3339Nano) + "'"
		}
		rows, err := db.Query(query)
		if err != nil {
			fmt.Fprintln(os.Stderr, "(POSTGRESQL) ERROR:", err)
			return
		}
		for rows.Next() {
			var m Metrics
			var metricTime time.Time

			err = rows.Scan(&m.Metric_1, &m.Metric_2, &m.Metric_3, &m.Metric_4, &m.Metric_5, &deviceID, &metricTime)
			if err != nil {
				fmt.Fprintln(os.Stderr, "(POSTGRESQL) ERROR:", err)
				continue
			}

			criticalMetrics := checkMetrics(m, config)
			if criticalMetrics != "" {
				_, err := db.Exec("INSERT INTO device_alerts(device_id, message) VALUES($1,$2);", deviceID, criticalMetrics)
				if err != nil {
					fmt.Fprintln(os.Stderr, "(POSTGRESQL) ERROR:", err)
					continue
				}
				err = redisClient.Set(strconv.Itoa(deviceID), criticalMetrics, 0).Err()
				if err != nil {
					fmt.Fprintln(os.Stderr, "(REDIS) ERROR:", err)
				}

				var email string
				emailQuery := `SELECT u.email from users u 
								join devices d on d.user_id = u.id 
								join device_metrics dm on dm.device_id = d.id 
								WHERE dm.device_id = $1`

				err = db.QueryRow(emailQuery, deviceID).Scan(&email)

				if err != nil {
					fmt.Fprintln(os.Stderr, "(POSTGRESQL) ERROR:", err)
				}
				message := fmt.Sprintf("DeviceID: %d\nMetrics: %s\n", deviceID, criticalMetrics)

				sendMail(config, message, email)
			}

			allQueries = append(allQueries, m)

			if metricTime.After(lastTime) {
				lastTime = metricTime
			}
		}
		_ = now
	}
}
