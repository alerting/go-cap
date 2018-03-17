package cap

import (
	"encoding/xml"
)

type Info struct {
	XMLName xml.Name `xml:"info" json:"-"`

	Language      string         `xml:"language" json:"language"`
	Categories    []Category     `xml:"category" json:"categories"`
	Event         string         `xml:"event" json:"event"`
	ResponseTypes []ResponseType `xml:"responseType" json:"response_types"`
	Urgency       Urgency        `xml:"urgency" json:"urgency"`
	Severity      Severity       `xml:"severity" json:"severity"`
	Certainty     Certainty      `xml:"certainty" json:"certainty"`
	Audience      string         `xml:"audience" json:"audience"`
	EventCodes    KeyValue       `xml:"eventCode" json:"event_codes"`
	Effective     *Time          `xml:"effective" json:"effective"`
	Onset         *Time          `xml:"onset" json:"onset"`
	Expires       *Time          `xml:"expires" json:"expires"`
	SenderName    string         `xml:"senderName" json:"sender_name"`
	Headline      string         `xml:"headline" json:"headline"`
	Description   string         `xml:"description" json:"description"`
	Instruction   string         `xml:"instruction" json:"instruction"`
	Web           string         `xml:"web" json:"web"`
	Contact       string         `xml:"contact" json:"contact"`
	Parameters    KeyValue       `xml:"parameter" json:"parameters"`
	Resources     []Resource     `xml:"resource" json:"resources"`
	Areas         []Area         `xml:"area" json:"areas"`
}
