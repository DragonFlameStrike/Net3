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
	Point   struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"point"`
}

type Weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}
