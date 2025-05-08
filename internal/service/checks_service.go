package service

import (
	"fmt"
	"website-monitoring/internal/model"
	"website-monitoring/internal/repository"
)

func GetAllChecksHistory() ([]model.Check, error) {
	rows, err := repository.GetBdAllChecks()
	if err != nil {
		fmt.Println("Erro ao buscar checks no banco de dados. ", err)
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
