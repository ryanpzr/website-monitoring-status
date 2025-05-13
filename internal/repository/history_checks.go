package repository

import (
	"database/sql"
	"time"
	"website-monitoring/configs/db"
	"website-monitoring/internal/model"
)

func PostBdSiteStatus(site model.Site, duration time.Duration, status string) error {
	conn, err := db.OpenConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := "INSERT INTO history_checks (site_id, status, time_response, http_code) VALUES ($1, $2, $3, $4)"
	_, err = conn.Exec(sql, site.Id, status, duration.Milliseconds(), 0)
	if err != nil {
		return err
	}

	return nil
}

func GetBdAllChecks() (*sql.Rows, error) {
	conn, err := db.OpenConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	sql := "SELECT * FROM history_checks"
	rows, err := conn.Query(sql)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
