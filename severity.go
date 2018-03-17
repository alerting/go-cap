package cap

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

const (
	SeverityUnknown = iota
	SeverityExtreme
	SeveritySevere
	SeverityModerate
	SeverityMinor
)

type Severity int

// UnmarshalString unmarshals the string into a Severity value.
func (severity *Severity) UnmarshalString(str string) error {
	if str == "Extreme" {
		*severity = SeverityExtreme
	} else if str == "Severe" {
		*severity = SeveritySevere
	} else if str == "Moderate" {
		*severity = SeverityModerate
	} else if str == "Minor" {
		*severity = SeverityMinor
	} else if str == "Unknown" {
		*severity = SeverityUnknown
	} else {
		return errors.New("Unknown Severity value: " + str)
	}

	return nil
}

// UnmarshalXML unmarshals the XML into a Severity value.
func (severity *Severity) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}
	return severity.UnmarshalString(str)
}

// UnmarshalJSON unmarshals the JSON into a Severity value.
func (severity *Severity) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	return severity.UnmarshalString(str)
}

// MarshalJSON returns the string version of the severity.
func (severity *Severity) MarshalJSON() ([]byte, error) {
	return json.Marshal(severity.String())
}

// MarshalXML returns the string version of the severity.
func (severity *Severity) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(severity.String(), start)
}

// String returns a Severity as a string
func (severity Severity) String() string {
	if severity == SeverityExtreme {
		return "Extreme"
	} else if severity == SeveritySevere {
		return "Severe"
	} else if severity == SeverityModerate {
		return "Moderate"
	} else if severity == SeverityMinor {
		return "Minor"
	} else if severity == SeverityUnknown {
		return "Unknown"
	}

	return ""
}
