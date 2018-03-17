package cap

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

const (
	MessageTypeUnknown = iota
	MessageTypeAlert
	MessageTypeUpdate
	MessageTypeCancel
	MessageTypeAck
	MessageTypeError
)

type MessageType int

// UnmarshalString unmarshals the string into a MessageType value.
func (messageType *MessageType) UnmarshalString(str string) error {
	if str == "Alert" {
		*messageType = MessageTypeAlert
	} else if str == "Update" {
		*messageType = MessageTypeUpdate
	} else if str == "Cancel" {
		*messageType = MessageTypeCancel
	} else if str == "Ack" {
		*messageType = MessageTypeAck
	} else if str == "Error" {
		*messageType = MessageTypeError
	} else {
		return errors.New("Unknown MessageType value")
	}

	return nil
}

// UnmarshalXML unmarshals the XML into a MessageType value.
func (messageType *MessageType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}
	return messageType.UnmarshalString(str)
}

// UnmarshalJSON unmarshals the XML into a MessageType value.
func (messageType *MessageType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	return messageType.UnmarshalString(str)
}

// MarshalJSON returns the string version of the message type.
func (messageType *MessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(messageType.String())
}

// MarshalXML returns the string version of the message type.
func (messageType *MessageType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(messageType.String(), start)
}

// String returns a MessageType as a string
func (messageType MessageType) String() string {
	if messageType == MessageTypeAlert {
		return "Alert"
	} else if messageType == MessageTypeUpdate {
		return "Update"
	} else if messageType == MessageTypeCancel {
		return "Cancel"
	} else if messageType == MessageTypeAck {
		return "Ack"
	} else if messageType == MessageTypeError {
		return "Error"
	} else if messageType == MessageTypeUnknown {
		return "Unknown"
	}

	return ""
}
