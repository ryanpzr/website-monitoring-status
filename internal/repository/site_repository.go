package repository

import (
	"database/sql"
	"website-monitoring/configs/dbconfig"
	"website-monitoring/internal/model"
)

func GetBdAllSites() (*sql.Rows, error) {
	conn, err := dbconfig.OpenConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	sql := "SELECT * FROM site"
	rows, err := conn.Query(sql)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func GetBdSiteById(id string) *sql.Row {
	conn, err := dbconfig.OpenConn()
	if err != nil {
		return nil
	}
	defer conn.Close()

	sql := `SELECT name, url, freq FROM site WHERE id = $1`
	row := conn.QueryRow(sql, id)

	return row
}

func PostBdSite(site model.Site) error {
	conn, err := dbconfig.OpenConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := `INSERT INTO site (name, url, freq) VALUES ($1, $2, $3)`
	_, err = conn.Exec(sql, site.Name, site.Url, site.Freq)
	if err != nil {
		return err
	}

	return nil
}
