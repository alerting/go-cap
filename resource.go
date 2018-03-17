package cap

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
)

type Resource struct {
	XMLName xml.Name `xml:"resource" json:"-"`

	Description string  `xml:"resourceDesc" json:"description"`
	MimeType    string  `xml:"mimeType" json:"mime_type"`
	Size        int     `xml:"size" json:"size"`
	Uri         string  `xml:"uri" json:"uri"`
	Digest      string  `xml:"digest" json:"digest"`
	DerefUri    *string `xml:"derefUri" json:"deref_uri"`
}

func (res *Resource) Checksum() string {
	if res.Digest != "" {
		return res.Digest
	}

	// If we have the contents, sha1sum that
	if res.DerefUri != nil {
		hash := sha1.New()
		hash.Write([]byte(*res.DerefUri))
		return hex.EncodeToString(hash.Sum(nil))
	}

	return ""
}
