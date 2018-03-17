package cap

import (
	"encoding/xml"
	"testing"
)

func TestSinglePolygonFromXML(t *testing.T) {
	sourceXML := []byte("<polygon>43.5481,-65.8528 43.613,-66.0816 43.8302,-66.2464 43.9343,-66.2609 44.0034,-66.154 44.0035,-66.1537 44.1327,-65.8891 44.2147,-65.458 44.227,-65.3927 44.227,-65.3927 44.227,-65.3927 43.7465,-65.6514 43.6446,-65.6121 43.5749,-65.7862 43.5734,-65.7898 43.5481,-65.8528</polygon>")
	expectedCoords := [][]float64{
		[]float64{-65.8528, 43.5481},
		[]float64{-66.0816, 43.613},
		[]float64{-66.2464, 43.8302},
		[]float64{-66.2609, 43.9343},
		[]float64{-66.154, 44.0034},
		[]float64{-66.1537, 44.0035},
		[]float64{-65.8891, 44.1327},
		[]float64{-65.458, 44.2147},
		[]float64{-65.3927, 44.227},
		[]float64{-65.6514, 43.7465},
		[]float64{-65.6121, 43.6446},
		[]float64{-65.7862, 43.5749},
		[]float64{-65.7898, 43.5734},
		[]float64{-65.8528, 43.5481},
	}

	var polygon Polygon
	if err := xml.Unmarshal(sourceXML, &polygon); err != nil {
		t.Fatal(err)
	}

	if polygon.Type != "polygon" {
		t.Errorf("Unexpected polygon type, got: %s, want: %s.", polygon.Type, "polygon")
	}

	if len(polygon.Coordinates) != 1 {
		t.Errorf("Unexpected number of polygons, got: %d, want: %d.", len(polygon.Coordinates), 1)
	}

	if len(polygon.Coordinates[0]) != len(expectedCoords) {
		t.Fatalf("Unexpected number of coordinates, got: %d, want: %d.", len(polygon.Coordinates[0]), len(expectedCoords))
	}

	for i, coord := range polygon.Coordinates[0] {
		if coord[0] != expectedCoords[i][0] || coord[1] != expectedCoords[i][1] {
			t.Errorf("Unexpected coordinate at index %d, got: %f,%f, want: %f,%f",
				i,
				coord[0], coord[1],
				expectedCoords[i][0], expectedCoords[i][1])
		}
	}
}

func TestSinglePolygonMissingFinalPointFromXML(t *testing.T) {
	sourceXML := []byte("<polygon>43.5481,-65.8528 43.613,-66.0816 43.8302,-66.2464 43.9343,-66.2609 44.0034,-66.154 44.0035,-66.1537 44.1327,-65.8891 44.2147,-65.458 44.227,-65.3927 44.227,-65.3927 44.227,-65.3927 43.7465,-65.6514 43.6446,-65.6121 43.5749,-65.7862 43.5734,-65.7898</polygon>")
	expectedCoords := [][]float64{
		[]float64{-65.8528, 43.5481},
		[]float64{-66.0816, 43.613},
		[]float64{-66.2464, 43.8302},
		[]float64{-66.2609, 43.9343},
		[]float64{-66.154, 44.0034},
		[]float64{-66.1537, 44.0035},
		[]float64{-65.8891, 44.1327},
		[]float64{-65.458, 44.2147},
		[]float64{-65.3927, 44.227},
		[]float64{-65.6514, 43.7465},
		[]float64{-65.6121, 43.6446},
		[]float64{-65.7862, 43.5749},
		[]float64{-65.7898, 43.5734},
		[]float64{-65.8528, 43.5481},
	}

	var polygon Polygon
	if err := xml.Unmarshal(sourceXML, &polygon); err != nil {
		t.Fatal(err)
	}

	if polygon.Type != "polygon" {
		t.Errorf("Unexpected polygon type, got: %s, want: %s.", polygon.Type, "polygon")
	}

	if len(polygon.Coordinates) != 1 {
		t.Errorf("Unexpected number of polygons, got: %d, want: %d.", len(polygon.Coordinates), 1)
	}

	if len(polygon.Coordinates[0]) != len(expectedCoords) {
		t.Fatalf("Unexpected number of coordinates, got: %d, want: %d.", len(polygon.Coordinates[0]), len(expectedCoords))
	}

	for i, coord := range polygon.Coordinates[0] {
		if coord[0] != expectedCoords[i][0] || coord[1] != expectedCoords[i][1] {
			t.Errorf("Unexpected coordinate at index %d, got: %f,%f, want: %f,%f",
				i,
				coord[0], coord[1],
				expectedCoords[i][0], expectedCoords[i][1])
		}
	}
}

func TestMultiplePoylgonsFromXML(t *testing.T) {
	sourceXML := []byte(`
    <area>
      <polygon>1,1 2,2, 3,3 1,1</polygon>
      <polygon>-1,-1 -2,-2, -3,-3 -4,-4 -1,-1</polygon>
    </area>
   `)
   expectedCoords := [][][]float64{
    [][]float64{
      []float64{ 1, 1 },
      []float64{ 2, 2 },
      []float64{ 3, 3 },
      []float64{ 1, 1 },
    },
    [][]float64{
      []float64{ -1, -1 },
      []float64{ -2, -2 },
      []float64{ -3, -3 },
      []float64{ -4, -4 },
      []float64{ -1, -1 },
    },
   }

	var area struct {
		Polygons Polygons `xml:"polygon"`
	}

	if err := xml.Unmarshal(sourceXML, &area); err != nil {
		t.Fatal(err)
	}

	if len(area.Polygons) != 2 {
		t.Errorf("Unexpected number of polyongs, got: %d, expected: %d", len(area.Polygons), 2)
	}

	for i, p := range area.Polygons {
		if p.Type != "polygon" {
			t.Errorf("Unexpected polygon type, got: %s, want: %s.", p.Type, "polygon")
		}

    if len(expectedCoords[i]) != len(p.Coordinates[0]) {
      t.Errorf("Unexpected number of coordinates, got: %d, want: %d.", len(p.Coordinates[0]), len(expectedCoords[i]))
    }
	}
}
