package cap

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

const (
	ResponseTypeUnknown = iota
	ResponseTypeShelter
	ResponseTypeEvacuate
	ResponseTypePrepare
	ResponseTypeExecute
	ResponseTypeAvoid
	ResponseTypeMonitor
	ResponseTypeAssess
	ResponseTypeAllClear
	ResponseTypeNone
)

type ResponseType int

// UnmarshalString unmarshals the string into a ResponseType value.
func (responseType *ResponseType) UnmarshalString(str string) error {
	if str == "Shelter" {
		*responseType = ResponseTypeShelter
	} else if str == "Evacuate" {
		*responseType = ResponseTypeEvacuate
	} else if str == "Prepare" {
		*responseType = ResponseTypePrepare
	} else if str == "Execute" {
		*responseType = ResponseTypeExecute
	} else if str == "Avoid" {
		*responseType = ResponseTypeAvoid
	} else if str == "Monitor" {
		*responseType = ResponseTypeMonitor
	} else if str == "Assess" {
		*responseType = ResponseTypeAssess
	} else if str == "AllClear" || str == "All Clear" {
		*responseType = ResponseTypeAllClear
	} else if str == "None" {
		*responseType = ResponseTypeNone
	} else {
		return errors.New("Unknown ResponseType value: " + str)
	}

	return nil
}

// UnmarshalXML unmarshals the XML into a ResponseType value.
func (responseType *ResponseType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}
	return responseType.UnmarshalString(str)
}

// UnmarshalJSON unmarshals the JSON into a ResponseType value.
func (responseType *ResponseType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	return responseType.UnmarshalString(str)
}

// MarshalJSON returns the string version of the response type.
func (responseType *ResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(responseType.String())
}

// MarshalXML returns the string version of the response type.
func (responseType *ResponseType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	str := responseType.String()

	// Some values differ in XML
	if str == "All Clear" {
		str = "AllClear"
	}

	return e.EncodeElement(str, start)
}

// String returns a ResponseType as a string
func (responseType ResponseType) String() string {
	if responseType == ResponseTypeShelter {
		return "Shelter"
	} else if responseType == ResponseTypeEvacuate {
		return "Evacuate"
	} else if responseType == ResponseTypePrepare {
		return "Prepare"
	} else if responseType == ResponseTypeExecute {
		return "Execute"
	} else if responseType == ResponseTypeAvoid {
		return "Avoid"
	} else if responseType == ResponseTypeMonitor {
		return "Monitor"
	} else if responseType == ResponseTypeAssess {
		return "Assess"
	} else if responseType == ResponseTypeAllClear {
		return "All Clear"
	} else if responseType == ResponseTypeNone {
		return "None"
	} else if responseType == ResponseTypeUnknown {
		return "Unknown"
	}

	return ""
}
