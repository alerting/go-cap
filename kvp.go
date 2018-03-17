package cap

import (
	"encoding/xml"
)

type KeyValue map[string][]string

type keyValuePair struct {
	Name  string `xml:"valueName"`
	Value string `xml:"value"`
}

func (m *KeyValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// If the map is not already created, create it.
	if *m == nil {
		*m = make(map[string][]string)
	}

	var kvp keyValuePair

	err := d.DecodeElement(&kvp, &start)

	if err != nil {
		return err
	}

	if _, ok := (*m)[kvp.Name]; !ok {
		(*m)[kvp.Name] = make([]string, 0)
	}

	if len(kvp.Value) > 0 {
		(*m)[kvp.Name] = append((*m)[kvp.Name], kvp.Value)
	}

	return nil
}

func (m *KeyValue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// If empty, don't need to proceed
	if *m == nil || len(*m) == 0 {
		return nil
	}

	for k, vs := range *m {
		for _, v := range vs {
			err := e.EncodeElement(keyValuePair{Name: k, Value: v}, start)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
