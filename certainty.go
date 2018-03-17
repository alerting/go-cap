package cap

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

const (
	CertaintyUnknown = iota
	CertaintyObserved
	CertaintyLikely
	CertaintyPossible
	CertaintyUnlikely
)

type Certainty int

// UnmarshalString unmarshals the string into a Certainty value.
func (certainty *Certainty) UnmarshalString(str string) error {
	if str == "Observed" {
		*certainty = CertaintyObserved
	} else if str == "Likely" || str == "Very Likely" {
		*certainty = CertaintyLikely
	} else if str == "Possible" {
		*certainty = CertaintyObserved
	} else if str == "Unlikely" {
		*certainty = CertaintyUnlikely
	} else if str == "Unknown" {
		*certainty = CertaintyUnknown
	} else {
		return errors.New("Unknown Certainty value: " + str)
	}

	return nil
}

// UnmarshalXML unmarshals the XML into a Certainty value.
func (certainty *Certainty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}
	return certainty.UnmarshalString(str)
}

// UnmarshalJSON unmarshals the JSON into a Certainty value.
func (certainty *Certainty) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	return certainty.UnmarshalString(str)
}

// MarshalJSON returns the string version of the certainty.
func (certainty *Certainty) MarshalJSON() ([]byte, error) {
	return json.Marshal(certainty.String())
}

// MarshalXML returns the string version of the certainty.
func (certainty *Certainty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(certainty.String(), start)
}

// String returns a Certainty as a string
func (certainty Certainty) String() string {
	if certainty == CertaintyObserved {
		return "Observed"
	} else if certainty == CertaintyLikely {
		return "Likely"
	} else if certainty == CertaintyPossible {
		return "Possible"
	} else if certainty == CertaintyUnlikely {
		return "Unlikley"
	} else if certainty == CertaintyUnknown {
		return "Unknown"
	}

	return ""
}
