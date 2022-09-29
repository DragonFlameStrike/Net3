package main

import (
	"Net3/Structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const GRAPHHOPPER_API_KEY = "c7e2a639-9ca1-4745-ae35-2cb8b3f9667b"
const OPENWEATHER_API_KEY = "844d933d6725aee705f0914bce0bd024"
const OPENTRIPMAP_API_KEY = "5ae2e3f221c38a28845f05b69a9339aad3d316264e2201f22b172ac8"

func narrator(xid string) {
	r := getNarratorStory(xid)
	ns := parseNarratorStory(r)
	if ns.Name == "" {
		return
	}
	if ns.Info.Descr == "" {
		ns.Info.Descr = "No Description"
	}
	fmt.Printf("Interesting place: %s\n Drescription: %s\n\n", ns.Name, ns.Info.Descr)
}

func guide(p Structs.Place) {
	r := getNearPlaces(p)
	otp := parseOpentripmap(r)
	for i := 0; i < len(otp.Features); i++ {
		go narrator(otp.Features[i].Properties.Xid)
	}
}

func weatherForecaster(p Structs.Place) {
	r := getWeatherJson(p)
	w := parseWeatherJson(r)
	printWeather(w)
}

func main() {
	q := getQueryName()
	r := getGraphhopperJson(q)
	gh := parseGraphhopperJson(r)
	p := getAnyPlaceOptions(gh)
	printPlacesToUser(p)
	cp := choosePlaceFromConsole(p)
	go weatherForecaster(cp)
	go guide(cp)
	time.Sleep(100 * time.Second)
}

func getQueryName() string {
	if len(os.Args) >= 2 {
		var q string
		for i := 1; i < len(os.Args); i++ {
			q = q + os.Args[i]
			if i != len(os.Args)-1 {
				q = q + "+"
			}
		}
		return q
	} else {
		return "Цветной+Проезд"
	}
}

func getGraphhopperJson(q string) *http.Response {
	resp, err := http.Get("https://graphhopper.com/api/1/geocode?q=" + q + "&locale=en&key=" + GRAPHHOPPER_API_KEY)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

func parseGraphhopperJson(r *http.Response) Structs.Graphhopper {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var gh Structs.Graphhopper
	err = json.Unmarshal(b, &gh)
	if err != nil {
		log.Fatalln(err)
	}
	if len(gh.Hits) == 0 {
		log.Fatalln("Nothing found")
	}
	return gh
}

func getAnyPlaceOptions(gh Structs.Graphhopper) []Structs.Place {
	var p []Structs.Place
	for i := 0; i < len(gh.Hits); i++ {
		pn := Structs.Place{
			Country: gh.Hits[i].Country,
			State:   gh.Hits[i].State,
			City:    gh.Hits[i].City,
			Street:  gh.Hits[i].Street,
			Name:    gh.Hits[i].Name,
			Point:   gh.Hits[i].Point,
		}
		p = append(p, pn)
	}
	return p
}

func choosePlaceFromConsole(p []Structs.Place) Structs.Place {
	var id int
	for {
		fmt.Scan(&id)
		if id == -1 {

			return Structs.Place{}
		}
		if id > len(p) || id < 1 {
			fmt.Println("wrong input, try again")
			continue
		}
		return p[id-1]
	}

}

func printPlacesToUser(p []Structs.Place) {
	for i := 0; i < len(p); i++ {
		fmt.Printf("Place %d: Name - %s, Country - %s, City - %s, State - %s, Street - %s\n",
			i+1, p[i].Name, p[i].Country, p[i].City, p[i].State, p[i].Street)
	}
}

func getWeatherJson(p Structs.Place) *http.Response {
	lat := fmt.Sprintf("%f", p.Point.Lat)
	lon := fmt.Sprintf("%f", p.Point.Lng)
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&appid=" + OPENWEATHER_API_KEY)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

func parseWeatherJson(r *http.Response) Structs.Weather {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var w Structs.Weather
	err = json.Unmarshal(b, &w)
	if err != nil {
		log.Fatalln(err)
	}
	return w
}

func printWeather(w Structs.Weather) {
	fmt.Println("WEATHER START:")
	fmt.Printf("Temp: %.1f,Humidity: %d,Wind speed: %.1fm/s\n", w.Main.Temp-273.15, w.Main.Humidity, w.Wind.Speed)
	fmt.Println("WEATHER END.")
}

func getNearPlaces(p Structs.Place) *http.Response {
	lat := fmt.Sprintf("%f", p.Point.Lat)
	lng := fmt.Sprintf("%f", p.Point.Lng)
	resp, err := http.Get("https://api.opentripmap.com/0.1/en/places/radius?radius=10000&lon=" + lng + "&lat=" + lat + "&apikey=" + OPENTRIPMAP_API_KEY)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

func parseOpentripmap(r *http.Response) Structs.Opentripmap {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var otp Structs.Opentripmap
	err = json.Unmarshal(b, &otp)
	if err != nil {
		log.Fatalln(err)
	}
	return otp
}

func getNarratorStory(xid string) *http.Response {
	resp, err := http.Get("https://api.opentripmap.com/0.1/en/places/xid/" + xid + "?apikey=5ae2e3f221c38a28845f05b69a9339aad3d316264e2201f22b172ac8")
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

func parseNarratorStory(r *http.Response) Structs.NarratorStory {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var ns Structs.NarratorStory
	err = json.Unmarshal(b, &ns)
	if err != nil {
		log.Fatalln(err)
	}
	return ns
}
