package cap

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Polygon struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}
type Polygons []*Polygon

func (p *Polygon) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}

	coords := strings.Split(str, " ")
	coordinates := make([][]float64, 0)

	var lastCoord []float64
	for _, coordStr := range coords {
		coordPair := strings.Split(coordStr, ",")

		lat, err := strconv.ParseFloat(coordPair[0], 64)
		if err != nil {
			return err
		}

		lon, err := strconv.ParseFloat(coordPair[1], 64)
		if err != nil {
			return err
		}

		// Only add coord if it's not the same as the previous point
		coord := []float64{lon, lat}
		if lastCoord == nil || (coord[0] != lastCoord[0] && coord[1] != lastCoord[1]) {
			coordinates = append(coordinates, coord)
			lastCoord = coord
		}
	}

	// Close the polygon, if it isn't already
	if coordinates[0][0] != coordinates[len(coordinates)-1][0] || coordinates[0][1] != coordinates[len(coordinates)-1][1] {
		coordinates = append(coordinates, coordinates[0])
	}

	p.Type = "polygon"
	p.Coordinates = [][][]float64{coordinates}
	return nil
}

func (m *Polygons) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, polygon := range *m {
		if len(polygon.Coordinates) != 1 {
			return errors.New("Invalid number of coordinates in polygon")
		}

		p := polygon.Coordinates[0]
		coords := make([]string, len(p))
		for i, c := range p {
			if len(c) != 2 {
				return errors.New("Invalid number of coordinates in polygon point")
			}
			coords[i] = fmt.Sprintf("%s,%s",
				strconv.FormatFloat(c[1], 'f', -1, 64),
				strconv.FormatFloat(c[0], 'f', -1, 64))
		}

		err := e.EncodeElement(strings.Join(coords, " "), start)
		if err != nil {
			return err
		}
	}

	return nil
}
