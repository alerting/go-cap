package cap

import (
	"encoding/xml"
)

type Area struct {
	XMLName xml.Name `xml:"area" json:"-"`

	Description string   `xml:"areaDesc" json:"description"`
	Polygons    Polygons `xml:"polygon" json:"polygons"`
	Circles     Circles  `xml:"circle" json:"circles"`
	GeoCodes    KeyValue `xml:"geocode" json:"geocodes"`
	Altitude    *int     `xml:"altitude" json:"altitude"`
	Ceiling     *int     `xml:"ceiling" json:"ceiling"`
}
