package main

import (
	"Net3/Structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const GRAPHHOPPER_API_KEY = "c7e2a639-9ca1-4745-ae35-2cb8b3f9667b"

func main() {
	q := getQueryName()
	r := getGraphhopperJson(q)
	gh := parseGraphhopperJson(r)
	p := getAnyPlaceOptions(gh)
	printPlacesToUser(p)
	cp := choosePlaceFromConsole(p)
	fmt.Printf("%v", cp)

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
