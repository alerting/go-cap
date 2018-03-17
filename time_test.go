package cap

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestTimeParse(t *testing.T) {
	b := []byte("2018-03-16T22:02:34-04:00")
	var ti Time

	if err := ti.UnmarshalText(b); err != nil {
		t.Fatal(err)
	}

	str := ti.FormatCAP()
	if string(b) != str {
		t.Errorf("Unexpected date response, got: %s, want: %s.", str, string(b))
	}
}

func TestTimeFormatUTC(t *testing.T) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		t.Fatal(err)
	}

	ti := Time{time.Date(2018, time.March, 16, 22, 13, 31, 5, loc)}

	str := ti.FormatCAP()
	exp := "2018-03-16T22:13:31-00:00"
	if exp != str {
		t.Errorf("Unexpected date response, got: %s, want: %s.", str, exp)
	}
}

func TestTimeFormatTimezone(t *testing.T) {
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		t.Fatal(err)
	}

	ti := Time{time.Date(2018, time.March, 16, 22, 13, 31, 5, loc)}

	str := ti.FormatCAP()
	exp := "2018-03-16T22:13:31-04:00"
	if exp != str {
		t.Errorf("Unexpected date response, got: %s, want: %s.", str, exp)
	}
}

func TestTimeMarshalUTCXML(t *testing.T) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		t.Fatal(err)
	}

	ti := Time{time.Date(2018, time.March, 16, 22, 13, 31, 5, loc)}

	b, err := xml.Marshal(ti)
	if err != nil {
		t.Fatal(err)
	}

	str := string(b)
	exp := "<Time>2018-03-16T22:13:31-00:00</Time>"
	if exp != str {
		t.Errorf("Unexpected date response, got: %s, want: %s.", str, exp)
	}
}

func TestTimeMarshalTimezoneXML(t *testing.T) {
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		t.Fatal(err)
	}

	ti := Time{time.Date(2018, time.March, 16, 22, 13, 31, 5, loc)}

	b, err := xml.Marshal(ti)
	if err != nil {
		t.Fatal(err)
	}

	str := string(b)
	exp := "<Time>2018-03-16T22:13:31-04:00</Time>"
	if exp != str {
		t.Errorf("Unexpected date response, got: %s, want: %s.", str, exp)
	}
}
