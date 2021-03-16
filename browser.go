package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

//func BytesToString(b []byte) string {
//	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
//	sh := reflect.StringHeader{bh.Data, bh.Len}
//	return *(*string)(unsafe.Pointer(&sh))
//}

func ResponseToJson(resp io.ReadCloser) interface{} {
	var jbody interface{}
	err := json.NewDecoder(resp).Decode(&jbody)
	if err != nil {
		panic(err)
	}
	return jbody
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

func GetStations(qstring string) interface{} {
	urlstr := fmt.Sprintf("https://%s/json/stations/search?%s&limit=100000",GetApiHost(),qstring)
	resp, err := http.Get(urlstr)
	if err != nil {
		log.Printf(err.Error())
	}
	defer resp.Body.Close()
	bodyJson := ResponseToJson(resp.Body)
	//body, err := io.ReadAll(resp.Body)
	//bodyString := bytes.NewBuffer(body).String()
	return bodyJson
}