package main

import (
	"fmt"
	"log"

	i2l "github.com/theplant/ip2location-nginx/go-client"
	"github.com/theplant/testingutils"
)

var _ip2locationClient *i2l.Client

func MustGetIp2locationClient() *i2l.Client {
	if _ip2locationClient != nil {
		return _ip2locationClient
	}

	_ip2locationClient = i2l.NewClient("https://ip2location.theplant-dev.com") // 放在config

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

func main() {

	lookupData := &LookupData{}

	var ips = []string{
		"133.18.5.188",   // Japan
		"216.24.176.149", // The US
		"216.24.176.149", // China
	}

	for _, ip := range ips {
		err := fillIpAndLocation(lookupData, ip)

		if err != nil {
			log.Fatal("failed: ")
		}

		testingutils.PrintlnJson(lookupData)
	}
}
