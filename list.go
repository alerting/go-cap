package cap

import (
	"encoding/xml"
	"strings"
)

type List []string

func (m *List) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	err := d.DecodeElement(&str, &start)

	if err != nil {
		return err
	}

	// Ignore empty string
	if str == "" {
		*m = make([]string, 0)
		return nil
	}

	// TODO: Seperate by any whitespace, ignoring escaped test
	*m = strings.Split(str, " ")

	return nil
}
