package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/jinzhu/configor"
	jsoniter "github.com/json-iterator/go"
	i2l "github.com/theplant/ip2location-nginx/go-client"
)

type Ip2locationConfig struct {
	ServiceUrl string `required:"true"`
}

var _ip2locationConfig *Ip2locationConfig

func MustGetIp2locationConfig() *Ip2locationConfig {
	if _ip2locationConfig != nil {
		return _ip2locationConfig
	}

	cfg := &Ip2locationConfig{}
	err := configor.New(&configor.Config{ENVPrefix: "IP2LOCATION"}).Load(cfg)
	if err != nil {
		panic(err)
	}

	_ip2locationConfig = cfg
	return _ip2locationConfig
}

var _ip2locationClient *i2l.Client

func MustGetIp2locationClient() *i2l.Client {
	if _ip2locationClient != nil {
		return _ip2locationClient
	}

	// cfg := MustGetIp2locationConfig()
	// export IP2LOCATION_SERVICEURL=https://ip2location.theplant-dev.com

	_ip2locationClient = i2l.NewClient("https://ip2location.theplant-dev.com")

	return _ip2locationClient
}

type IpLocation struct {
	City        string  `json:"city"`
	CityCode    string  `json:"cityCode"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Area        string  `json:"area"`
	AreaCode    string  `json:"areaCode"`
	Region      string  `json:"region"`
	RegionCode  string  `json:"regionCode"`
}

type LookupData struct {
	IpAddress  string     `json:"ipAddress"`
	IpLocation IpLocation `json:"ipLocation"`
}

func fillIpAndLocation(d *LookupData, ip string) error {
	var i2lClient = MustGetIp2locationClient()

	location, err := i2lClient.Ip2location(ip)
	if err != nil {
		if err == i2l.ErrInvalidIpAddress {
			fmt.Println(
				"msg", "invalid ip address",
				"ip", ip,
			)
			return nil
		}
		return err
	}

	d.IpAddress = ip
	d.IpLocation = IpLocation{
		City:        location.City,
		CityCode:    "",
		Country:     location.CountryName,
		CountryCode: location.CountryCode,
		Latitude:    location.Latitude,
		Longitude:   location.Longitude,
		Area:        "",
		AreaCode:    "",
		Region:      location.Region,
		RegionCode:  "",
	}

	return nil
}

func proxy(r *http.Request) []string {
	if ips := r.Header.Get("X-Forwarded-For"); ips != "" {
		return strings.Split(ips, ",")
	}

	return nil
}

func IP(r *http.Request) string {
	if r == nil {
		return ""
	}

	ips := proxy(r)

	if len(ips) > 0 && ips[0] != "" {
		rip, _, err := net.SplitHostPort(ips[0])
		if err != nil {
			rip = ips[0]
		}
		return rip
	}

	if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return ip
	}

	return r.RemoteAddr
}

func getLocationData(w http.ResponseWriter, r *http.Request) {
	lookupData := &LookupData{}

	err := fillIpAndLocation(lookupData, IP(r))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "get ip address failed")
		return
	}

	ret, err := jsoniter.Marshal(lookupData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
}

func main() {
	http.HandleFunc("/", getLocationData)    // 设置访问的路由
	err := http.ListenAndServe(":4000", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	log.Fatal("ListenAndServe: 4000 succeed")

}
