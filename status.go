package cap

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

const (
	StatusUnknown = iota
	StatusActual
	StatusExcercise
	StatusSystem
	StatusTest
	StatusDraft
)

type Status int

// UnmarshalString unmarshals the string into a Status value.
func (status *Status) UnmarshalString(str string) error {
	if str == "Actual" {
		*status = StatusActual
	} else if str == "Excercise" {
		*status = StatusExcercise
	} else if str == "System" {
		*status = StatusSystem
	} else if str == "Test" {
		*status = StatusTest
	} else if str == "Draft" {
		*status = StatusDraft
	} else {
		return errors.New("Unknown Status value: " + str)
	}

	return nil
}

// UnmarshalXML unmarshals the XML into a Status value.
func (status *Status) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}
	return status.UnmarshalString(str)
}

// UnmarshalJSON unmarshals the JSON into a Status value.
func (status *Status) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	return status.UnmarshalString(str)
}

// MarshalJSON returns the string version of the status.
func (status *Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(status.String())
}

// MarshalXML returns the string version of the status.
func (status *Status) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(status.String(), start)
}

// String returns a Status as a string
func (status Status) String() string {
	if status == StatusActual {
		return "Actual"
	} else if status == StatusExcercise {
		return "Excercise"
	} else if status == StatusSystem {
		return "System"
	} else if status == StatusTest {
		return "Test"
	} else if status == StatusDraft {
		return "Draft"
	} else if status == StatusUnknown {
		return "Unknown"
	}

	return ""
}
