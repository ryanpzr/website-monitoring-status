package service

import (
	"website-monitoring/internal/model"
	"website-monitoring/internal/repository"
)

func PostSite(site model.Site) (model.Site, error) {
	err := repository.PostBdSite(site)
	if err != nil {
		return model.Site{}, err
	}
	return site, nil
}

func GetSiteById(id string) (model.Site, error) {
	row := repository.GetBdSiteById(id)
	var site model.Site
	row.Scan(
		&site.Name,
		&site.Url,
		&site.Freq,
	)

	return site, nil
}

func GetAllSites() ([]model.Site, error) {
	rows, err := repository.GetBdAllSites()
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
