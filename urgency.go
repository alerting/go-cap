package cap

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

const (
	UrgencyUnknown = iota
	UrgencyImmediate
	UrgencyExpected
	UrgencyFuture
	UrgencyPast
)

type Urgency int

// UnmarshalString unmarshals the string into a Urgency value.
func (urgency *Urgency) UnmarshalString(str string) error {
	if str == "Immediate" {
		*urgency = UrgencyImmediate
	} else if str == "Expected" {
		*urgency = UrgencyExpected
	} else if str == "Future" {
		*urgency = UrgencyFuture
	} else if str == "Past" {
		*urgency = UrgencyPast
	} else if str == "Unknown" {
		*urgency = UrgencyUnknown
	} else {
		return errors.New("Unknown Urgency value: " + str)
	}

	return nil
}

// UnmarshalXML unmarshals the XML into a Urgency value.
func (urgency *Urgency) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}
	return urgency.UnmarshalString(str)
}

// UnmarshalJSON unmarshals the JSON into a Urgency value.
func (urgency *Urgency) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	return urgency.UnmarshalString(str)
}

// MarshalJSON returns the string version of the urgency.
func (urgency *Urgency) MarshalJSON() ([]byte, error) {
	return json.Marshal(urgency.String())
}

// MarshalXML returns the string version of the urgency.
func (urgency *Urgency) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(urgency.String(), start)
}

// String returns a Urgency as a string
func (urgency Urgency) String() string {
	if urgency == UrgencyImmediate {
		return "Immediate"
	} else if urgency == UrgencyExpected {
		return "Expected"
	} else if urgency == UrgencyFuture {
		return "Future"
	} else if urgency == UrgencyPast {
		return "Past"
	} else if urgency == UrgencyUnknown {
		return "Unknown"
	}

	return ""
}
