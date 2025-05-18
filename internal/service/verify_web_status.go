package service

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"website-monitoring/internal/model"
	"website-monitoring/internal/repository"
)

const (
	StatusOnline  = "Online"
	StatusOffline = "Offline"
)

func VerifyWebStatus() {
	siteCh := make(chan []model.Site)
	go getRowsBd(siteCh)

	go func() {
		for {
			siteList := <-siteCh
			if siteList == nil {
				log.Println("Erro ao buscar dados do site")
				return
			}

			for _, site := range siteList {
				go monitor(site)
			}
		}
	}()
}

func getRowsBd(ch chan<- []model.Site) {
	for {
		fmt.Println("Verificando se a novos sites cadastrados...")
		time.Sleep(10 * time.Second)

		siteList, err := fetchSite()
		if err != nil {

		}
		ch <- siteList
	}
}

func fetchSite() ([]model.Site, error) {
	rows, err := repository.GetBdAllSites()
	if err != nil {
		log.Fatal("Erro ao buscar dados no banco.", err)
		return nil, err
	}
	defer rows.Close()

	var siteList []model.Site
	for rows.Next() {
		var id int
		var name string
		var url string
		var freq int

		err := rows.Scan(&name, &url, &freq, &id)
		if err != nil {
			fmt.Println("Erro ao ler linha referente aos resultados do banco.", err)
			continue
		}

		site := model.Site{
			Name: name,
			Url:  url,
			Freq: freq,
			Id:   id,
		}

		siteList = append(siteList, site)
	}

	return siteList, nil
}

func monitor(site model.Site) {
	ticker := time.NewTicker(time.Duration(site.Freq) * time.Second)
	defer ticker.Stop()
	errTimes := 0

	for {
		select {
		case <-ticker.C:

			start := time.Now()
			resp, err := http.Get(site.Url)
			duration := time.Since(start)

			if err != nil {
				err = repository.PostBdSiteStatus(site, duration, "Offline")
				if err != nil {
					fmt.Println("Erro ao guardar status do site. ", err)
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

			err = repository.PostBdSiteStatus(site, duration, "Online")
			if err != nil {
				fmt.Println("Erro ao guardar status do site. ", err)
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
