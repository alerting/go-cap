package cap

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Circle struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
	Radius      float64   `json:"radius"`
}
type Circles []*Circle

func (m *Circles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	err := d.DecodeElement(&str, &start)

	if err != nil {
		return err
	}

	// Ignore empty strings
	if str == "" {
		return nil
	}

	// Create the list if it doesn't already exist
	if *m == nil {
		*m = make([]*Circle, 0)
	}

	stra := strings.Split(str, " ")
	coordPair := strings.Split(stra[0], ",")

	lat, err := strconv.ParseFloat(coordPair[0], 64)
	if err != nil {
		return err
	}

	lon, err := strconv.ParseFloat(coordPair[1], 64)
	if err != nil {
		return err
	}

	rad, err := strconv.ParseFloat(stra[1], 64)
	if err != nil {
		return err
	}

	c := Circle{
		Type:        "circle",
		Coordinates: []float64{lon, lat},
		// Convert km to m
		Radius: float64(rad / 1000.0),
	}
	*m = append(*m, &c)
	return nil
}

func (m *Circles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, circle := range *m {
		if len(circle.Coordinates) != 2 {
			return errors.New("Invalid number of coordinates in circle")
		}

		str := fmt.Sprintf("%s,%s %s",
			strconv.FormatFloat(circle.Coordinates[1], 'f', -1, 64),
			strconv.FormatFloat(circle.Coordinates[0], 'f', -1, 64),
			strconv.FormatFloat(circle.Coordinates[2], 'f', -1, 64))
		err := e.EncodeElement(str, start)
		if err != nil {
			return err
		}
	}

	return nil
}
