package cap

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
)

type Alert struct {
	XMLName xml.Name `xml:"alert" json:"-"`

	Identifier  string      `xml:"identifier" json:"identifier"`
	Sender      string      `xml:"sender" json:"sender"`
	Sent        Time        `xml:"sent" json:"sent"`
	Status      Status      `xml:"status" json:"status"`
	MessageType MessageType `xml:"msgType" json:"message_type"`
	Source      *string     `xml:"source" json:"source"`
	Scope       Scope       `xml:"scope" json:"scope"`
	Restriction *string     `xml:"restriction" json:"restriction"`
	Addresses   List        `xml:"addresses" json:"addresses"`
	Codes       []string    `xml:"code" json:"codes"`
	Note        *string     `xml:"note" json:"note"`
	References  References  `xml:"references" json:"references"`
	Incidents   List        `xml:"incidents" json:"incidents"`
	Infos       []Info      `xml:"info" json:"infos"`
}

func (alert *Alert) Id() string {
	hash := sha1.New()
	hash.Write([]byte(fmt.Sprintf("%s,%s,%s", alert.Sender, alert.Sent.FormatCAP(), alert.Identifier)))
	return hex.EncodeToString(hash.Sum(nil))
}
