package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"time"
)

type stationRecord struct {
	Name        string `json:"name"`
	Codec       string `json:"codec"`
	Bitrate     json.Number `json:"bitrate"`
	Countrycode string `json:"countrycode"`
	Tags        string `json:"tags"`
	Url         string `json:"url"`
	Lastcheck   int `json:"lastcheckok"`
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
	urlstr := fmt.Sprintf("https://%s/json/stations/search?%s&limit=%d",GetApiHost(),qstring,maxitems())
	resp, err := http.Get(urlstr)
	if err != nil {
		log.Print(err.Error())
	}
	defer resp.Body.Close()

	var data []stationRecord
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, err
}

func pruneStations(stations []stationRecord) []stationRecord {
	filteredStations := stations[:0]
	for _, station := range stations {
		if station.Lastcheck == 1 {
			filteredStations = append(filteredStations, station)
		}
	}
	return filteredStations
}

func StationSearch(name string, country string, state string, tags string, notok bool) ([]stationRecord, error) {
	params := url.Values{}
	if name != ""{
		params.Add("name", name)
	}
	if country != "" {
		params.Add("country", country)
	}
	if state != ""{
		params.Add("state", state)
	}
	if tags != ""{
		params.Add("tag",tags)
	}

	stations, err := GetStations(params.Encode())
	if err != nil{
		return nil, err
	}
	if notok{
		return stations, err
	} // otherwise prune the list
	prunedStations := pruneStations(stations) // eliminate stations that are reporting down.
	return prunedStations, err
}


