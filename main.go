package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/davidlick/digitalocean-ddns/digitalocean"
)

var (
	client = &http.Client{}
)

func main() {
	token := os.Getenv("DIGITALOCEAN_TOKEN")
	doBaseUrl := os.Getenv("DIGITALOCEAN_URL")
	doDomain := os.Getenv("DIGITALOCEAN_DOMAIN")

	client := digitalocean.NewClient(doBaseUrl, token, doDomain)

	ip, err := getPublicIP()
	if err != nil {
		log.Fatalf("failed to get public ip: %v", err)
	}

	id, err := client.GetARecordID()
	if err != nil {
		log.Fatalf("failed to get A record ID: %v", err)
	}

	err = client.SetARecord(id, ip)
	if err != nil {
		log.Fatalf("failed to set A record ip: %v", err)
	}

	log.Printf("successfully set DNS A record id %s to %s", id, ip)
}

func getPublicIP() (string, error) {
	resp, err := http.Get("https://icanhazip.com")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
