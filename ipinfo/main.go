package main

import (
	"fmt"
	"io/ioutil"
	"flag"
	"net/http"
	"encoding/json"
	"strconv"
	"net"
)

type Data struct {
	IPaddress string `json:"query"`
	Organization string `json:"org"`
	City string `json:"city"`
	Region string `json:"regionName"`
	Country string `json:"country"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Postal string `json:"zip"`
}

func ValidateIP(ip string) {
	if net.ParseIP(ip) == nil {
     	err := "Wrong ip address!"
        panic(err)
     }
}

func main() {
	var data Data
	// set normal flag
	 ip := flag.String("ip", "81.190.40.214", "ip address")
	 //set geo flg
	 geo := flag.Bool("geo", false, "geo")
     flag.Parse()
     ipstr := *ip // change type - string
     // validate ip address
     ValidateIP(ipstr)
     // get respond
	resp, err := http.Get("http://ip-api.com/json/"+ipstr+"?fields=query,org,city,regionName,country,zip,lat,lon")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// read all data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        panic(err)
    }
	// decoding data
	errDecoding := json.Unmarshal(body, &data)
    if errDecoding != nil {
        panic(errDecoding)
    }
    // output 
    if *geo {
    	fmt.Println("Location: "+strconv.FormatFloat(data.Lat, 'f', 4, 64)+", "+strconv.FormatFloat(data.Lon, 'f', 4, 64))
    } else {
	    fmt.Println("IP address: "+data.IPaddress)
	    fmt.Println("Organization: "+data.Organization)
	    fmt.Println("City: "+data.City)
	    fmt.Println("Region: "+data.Region)
	    fmt.Println("Country: "+data.Country)
	    fmt.Println("Location: "+strconv.FormatFloat(data.Lat, 'f', 4, 64)+", "+strconv.FormatFloat(data.Lon, 'f', 4, 64))
	    fmt.Println("Postal: "+data.Postal)
	}
}
