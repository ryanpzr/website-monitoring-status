package service

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"website-monitoring/configs/dbconfig"
	"website-monitoring/internal/model"
)

const (
	StatusOnline  = "Online"
	StatusOffline = "Offline"
)

func VerifyWebStatus() {
	rowsCh := make(chan *sql.Rows)
	go getRowsBd(rowsCh)

	rows := <-rowsCh
	if rows == nil {
		log.Println("Erro ao buscar dados")
		return
	}

	siteList := getSiteList(rows)
	for _, site := range siteList {
		go monitor(site)
	}
}

func getRowsBd(ch chan<- *sql.Rows) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Verificando se a novos sites cadastrados...")
			conn, err := dbconfig.OpenConn()
			if err != nil {
				fmt.Println("Erro ao abrir conexão com o banco", err)
				ch <- nil
				return
			}
			defer conn.Close()

			sql := "SELECT name, url, freq, id FROM site"
			resp, err := conn.Query(sql)
			if err != nil {
				log.Fatal("Erro ao buscar dados no banco.", err)
				ch <- nil
				return
			}

			ch <- resp
		}
	}
}

func getSiteList(rows *sql.Rows) []model.Site {
	var siteList []model.Site
	for rows.Next() {
		var id int
		var name string
		var url string
		var freq int

		err := rows.Scan(&name, &url, &freq, &id)
		if err != nil {
			fmt.Println("Erro ao ler linha referente aos resultados do banco.", err)
		}

		site := model.Site{
			Name: name,
			Url:  url,
			Freq: freq,
			Id:   id,
		}

		siteList = append(siteList, site)
	}
	return siteList
}

func monitor(site model.Site) {
	ticker := time.NewTicker(time.Duration(site.Freq) * time.Second)
	defer ticker.Stop()
	errTimes := 0

	conn, err := dbconfig.OpenConn()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	for {
		select {
		case <-ticker.C:
			sql := "INSERT INTO history_checks (site_id, status, time_response, http_code) VALUES ($1, $2, $3, $4)"

			start := time.Now()
			resp, err := http.Get(site.Url)
			duration := time.Since(start)

			if err != nil {
				_, err := conn.Exec(sql, site.Id, "Offline", duration.Milliseconds(), 0)
				if err != nil {
					fmt.Println(err)
					return
				}

				errTimes++
				if errTimes >= 3 {
					fmt.Println("====================")
					fmt.Println("O site ", site.Name, " falhou 3 vezes consecutivas. Verifique a disponibilidade do serviço.")
					fmt.Println("====================")
					errTimes = 0
				}
				log.Printf("[%s] Site %s (%s) está %s - Código: %d - Tempo: %dms",
					time.Now().Format(time.RFC3339),
					site.Name,
					site.Url,
					StatusOffline,
					404,
					duration.Milliseconds())

				continue
			}

			_, err = conn.Exec(sql, site.Id, "Online", duration.Milliseconds(), resp.StatusCode)
			if err != nil {
				fmt.Println(err)
				return
			}

			log.Printf("[%s] Site %s (%s) está %s - Código: %d - Tempo: %dms",
				time.Now().Format(time.RFC3339),
				site.Name,
				site.Url,
				StatusOnline,
				resp.StatusCode,
				duration.Milliseconds())

		}
	}
}
