package service

import (
	"website-monitoring/configs/dbconfig"
	"website-monitoring/internal/model"
)

func GetAllChecksHistory() ([]model.Check, error) {
	conn, err := dbconfig.OpenConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	sql := "SELECT * FROM history_checks"
	rows, err := conn.Query(sql)
	if err != nil {
		return nil, err
	}

	var checkList []model.Check
	for rows.Next() {
		var check model.Check
		if err := rows.Scan(
			&check.Id,
			&check.SiteId,
			&check.Status,
			&check.TimeResponse,
			&check.HttpCode,
			&check.TimeCreated,
		); err != nil {
			return nil, err
		}

		checkList = append(checkList, check)
	}

	return checkList, nil
}
