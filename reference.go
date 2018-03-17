package cap

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"strings"
)

type Reference struct {
	Sender     string `json:"sender"`
	Identifier string `json:"identifier"`
	Sent       Time   `json:"sent"`
}
type References []*Reference

func (reference *Reference) Id() string {
	hash := sha1.New()
	hash.Write([]byte(fmt.Sprintf("%s,%s,%s", reference.Sender, reference.Sent.FormatCAP(), reference.Identifier)))
	return hex.EncodeToString(hash.Sum(nil))
}

func (m *References) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}

	// Ignore empty strings
	if str == "" {
		return nil
	}

	// Create the list if it doesn't already exist
	references := strings.Split(str, " ")

	for _, reference := range references {
		components := strings.Split(reference, ",")

		var sent Time
		if err := sent.UnmarshalText([]byte(components[2])); err != nil {
			return err
		}

		*m = append(*m, &Reference{
			Sender:     components[0],
			Identifier: components[1],
			Sent:       sent,
		})
	}

	return nil
}

func (m *References) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	values := make([]string, len(*m))

	for i, ref := range *m {
		values[i] = fmt.Sprintf("%s,%s,%s", ref.Sender, ref.Identifier, ref.Sent.FormatCAP())
	}

	return e.EncodeElement(strings.Join(values, " "), start)
}
