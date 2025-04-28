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

	sql := `INSERT INTO site (name, url, freq, id) VALUES ($1, $2, $3, $4)`
	_, err = conn.Exec(sql, site.Name, site.Url, site.Freq, site.Id)
	if err != nil {
		return model.Site{}, err
	}

	return site, nil
}
