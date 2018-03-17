package cap

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

type VisitFunc func(alert *Alert) error

func ParseAlert(path string) (*Alert, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var alert Alert
	err = xml.Unmarshal(contents, &alert)
	if err != nil {
		return nil, err
	}

	return &alert, nil
}

func WalkAlerts(dir string, visitFunc VisitFunc) error {
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		// Don't process folders
		if err != nil || f.IsDir() {
			return err
		}

		// Only handle xml files
		if !strings.HasSuffix(f.Name(), ".xml") {
			return nil
		}

		alert, err := ParseAlert(path)
		if err != nil {
			return err
		}

		return visitFunc(alert)
	})

	return err
}

func ParseAlerts(dir string) ([]*Alert, error) {
	alerts := make([]*Alert, 0)

	err := WalkAlerts(dir, func(alert *Alert) error {
		alerts = append(alerts, alert)
		return nil
	})

	return alerts, err
}
