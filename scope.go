package cap

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

const (
	ScopeUnknown = iota
	ScopePublic
	ScopeRestricted
	ScopePrivate
)

type Scope int

// UnmarshalString unmarshals the string into a Scope value.
func (scope *Scope) UnmarshalString(str string) error {
	if str == "Public" {
		*scope = ScopePublic
	} else if str == "Restricted" {
		*scope = ScopeRestricted
	} else if str == "Private" {
		*scope = ScopePrivate
	} else {
		return errors.New("Unknown Scope value")
	}

	return nil
}

// UnmarshalXML unmarshals the XML into a Scope value.
func (scope *Scope) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}
	return scope.UnmarshalString(str)
}

// UnmarshalJSON unmarshals the JSON into a Scope value.
func (scope *Scope) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	return scope.UnmarshalString(str)
}

// MarshalJSON returns the string version of the scope.
func (scope *Scope) MarshalJSON() ([]byte, error) {
	return json.Marshal(scope.String())
}

// MarshalXML returns the string version of the scope.
func (scope *Scope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(scope.String(), start)
}

// String returns a Scope as a string
func (scope Scope) String() string {
	if scope == ScopePublic {
		return "Public"
	} else if scope == ScopeRestricted {
		return "Restricted"
	} else if scope == ScopePrivate {
		return "Private"
	} else if scope == ScopeUnknown {
		return "Unknown"
	}

	return ""
}
