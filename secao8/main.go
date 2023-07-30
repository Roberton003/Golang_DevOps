package main

import (
	"crypto/tls"
    "encoding/csv"
    "log"
    "math"
    "net"
    "os"
    "sync"
    "time"
	"errors"
	"net/http"
	"fmt"
	"io"
			
)
	
type TopURL struct {
		GlobalRanking 	string // position 0
		DomainRanking 	string // position 1
		Endereco 		string // position 2
		Pais 			string // position 3
}

func CheckCert(server string, port string, ranking string) {
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 180 * time.Second},"tcp", server+":"+port, nil)
	defer conn.Close()
	if err != nil {
		log.Println("Server %s ranking %s nao suporta certificacao SSL", server, ranking)
		return
	}
		expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	currentTime := time.Now()
	diff := expiry.Sub(currentTime)
	log.Printf("Tempo restante para o server %s expirar: %1.f", server, math.Round(diff.Hours()/24), ranking)
}

func CriarListaURLS(urlLine  *os.File) []TopURL {
		csvReader := csv.NewReader(urlLine)
		data, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		var urls []TopURL
		for _, line := range data {
			if line[3] == "br" {
				url := TopURL{
					GlobalRanking: 	line[0],
					DomainRanking: 	line[1],
					Endereco: 		line[2],
					Pais: 			line[3],
				}
				urls = append(urls, url)
			}
			return urls
		}
		return nil		
	}

func DownloadMillionDomains() error {
	if _, err := os.Stat("majestic_million.csv"); errors.Is(err, os.ErrNotExist) {
		log.Println("Downloading file...")
		url := "https://downloads.majestic.com/majestic_million.csv"
		out, err := os.Create("majestic_million.csv")
		if err != nil {
			return err 
		} 
		defer out.Close()
		resp, err := http.Get(url)
		if err != nil {
		return err
		}
		defer resp.Body.Close()
		if (resp.StatusCode != http.StatusOK) {
			return fmt.Errorf("bad status: %s", resp.Status)
		}

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
		log.Println("Downloaded file Finish...")
		} else {
			log.Println("File CSV majestic already exists...")
	}
	return nil
}
func main() {
	file, err := os.OpenFile("logss.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	
	err = DownloadMillionDomains()
	if err != nil {
		log.Fatal(err)
	}
	urlList, err := os.Open("majestic_million.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer urlList.Close()
	urls := CriarListaURLS(urlList)
	var wg sync.WaitGroup
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go func (url TopURL) {

			CheckCert(urls[1].Endereco, "443", urls[1].DomainRanking)
			defer wg.Done()
		}(urls[i])

	}
	wg.Wait()
}
