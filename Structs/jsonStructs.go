package Structs

type Graphhopper struct {
	Hits []struct {
		Point struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"point"`
		Extent      []float64 `json:"extent,omitempty"`
		Name        string    `json:"name"`
		Country     string    `json:"country"`
		Countrycode string    `json:"countrycode"`
		OsmId       int       `json:"osm_id"`
		OsmType     string    `json:"osm_type"`
		OsmKey      string    `json:"osm_key"`
		OsmValue    string    `json:"osm_value"`
		City        string    `json:"city,omitempty"`
		Street      string    `json:"street,omitempty"`
		Postcode    string    `json:"postcode,omitempty"`
		State       string    `json:"state,omitempty"`
		Housenumber string    `json:"housenumber,omitempty"`
		HouseNumber string    `json:"house_number,omitempty"`
	} `json:"hits"`
	Locale string `json:"locale"`
}

type Place struct {
	Country string `json:"country"`
	State   string `json:"state,omitempty"`
	City    string `json:"city,omitempty"`
	Street  string `json:"street,omitempty"`
	Name    string `json:"name"`
}
