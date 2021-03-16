package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

type stationRecord struct {
	Name        string `json:"name"`
	Codec       string `json:"codec"`
	Bitrate     json.Number `json:"bitrate"`
	Countrycode string `json:"countrycode"`
	Tags        string `json:"tags"`
	Url         string `json:"url"`
	Lastcheck   int `json:"lastcheck"`
}


func RandomIP(iplist []net.IP) net.IP {
	rand.Seed(time.Now().Unix())
	randomIndex := rand.Intn(len(iplist))
	return iplist[randomIndex]
}

func nslookup(hostname string) net.IP {
	iprecords, _ := net.LookupIP(hostname)
	randomIp := RandomIP(iprecords)
	return randomIp
}

func reverseLookup(ipAddr net.IP) string {
	ptr, _ := net.LookupAddr(ipAddr.String())
	return ptr[0]
}

func GetApiHost() string {
	appHostIp := nslookup(api())
	apiHost := reverseLookup(appHostIp)
	return apiHost
}

func GetStations(qstring string) ([]stationRecord, error){
	urlstr := fmt.Sprintf("https://%s/json/stations/search?%s&limit=100000",GetApiHost(),qstring)
	resp, err := http.Get(urlstr)
	if err != nil {
		log.Printf(err.Error())
	}
	defer resp.Body.Close()

	var data []stationRecord
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, err
}


