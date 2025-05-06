package service

import (
	"website-monitoring/configs/dbconfig"
	"website-monitoring/internal/model"
)

func PostSite(site model.Site) (model.Site, error) {
	conn, err := dbconfig.OpenConn()
	if err != nil {
		return model.Site{}, err
	}
	defer conn.Close()

	sql := `INSERT INTO site (name, url, freq) VALUES ($1, $2, $3)`
	_, err = conn.Exec(sql, site.Name, site.Url, site.Freq)
	if err != nil {
		return model.Site{}, err
	}

	return site, nil
}

func GetSiteById(id string) (model.Site, error) {

	conn, err := dbconfig.OpenConn()
	if err != nil {
		return model.Site{}, err
	}
	defer conn.Close()

	sql := `SELECT name, url, freq FROM site WHERE id = $1`
	row := conn.QueryRow(sql, id)

	var site model.Site
	if err = row.Scan(
		&site.Name,
		&site.Url,
		&site.Freq,
	); err != nil {
		return model.Site{}, err
	}

	return site, nil
}

func GetAllSites() ([]model.Site, error) {
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

	var siteList []model.Site
	for rows.Next() {
		var site model.Site
		if err := rows.Scan(
			&site.Name,
			&site.Url,
			&site.Freq,
			&site.Id,
		); err != nil {
			return nil, err
		}

		siteList = append(siteList, site)
	}

	return siteList, nil
}
