package gopom

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/charmap"
	"io"
)

// ParentPOM represents the parent element in a POM xml
type ParentPOM struct {
	XMLName      xml.Name `xml:"parent"`
	GroupID      string   `xml:"groupId"`
	ArtifactID   string   `xml:"artifactId"`
	Version      string   `xml:"version"`
	RelativePath string   `xml:"relativePath"`
}

// PropertySet represents the properties element in a POM xml
type PropertySet struct {
	XMLName xml.Name        `xml:"properties"`
	Entries []PropertyEntry `xml:",any"`
}

// PropertyEntry represents a property entry element in POM xml
type PropertyEntry struct {
	Name  string
	Value string `xml:",chardata"`
}

// UnmarshalXML unmarshals a property entry in POM xml
func (e *PropertyEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	e.Name = start.Name.Local
	for {
		token, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.CharData:
			e.Value = string(token.(xml.CharData).Copy())
		case xml.EndElement:
			return nil
		}
	}
}

// Exclusion represents an exclusion element in POM xml
type Exclusion struct {
	XMLName    xml.Name `xml:"exclusion"`
	GroupID    string   `xml:"groupId"`
	ArtifactID string   `xml:"artifactId"`
}

// Dependency represents a dependency element in POM xml
type Dependency struct {
	XMLName    xml.Name    `xml:"dependency"`
	GroupID    string      `xml:"groupId"`
	ArtifactID string      `xml:"artifactId"`
	Version    string      `xml:"version"`
	Type       string      `xml:"type"`
	Classifier string      `xml:"classifier"`
	SystemPath string      `xml:"systemPath"`
	Scope      string      `xml:"scope"`
	Optional   bool        `xml:"optional"`
	Exclusions []Exclusion `xml:"exclusions>exclusion"`
}

// DependencyManagement represents the dependencyManagement element in POM xml
type DependencyManagement struct {
	XMLName      xml.Name     `xml:"dependencyManagement"`
	Dependencies []Dependency `xml:"dependencies>dependency"`
}

// POM represents the `maven` project object model
type POM struct {
	XMLName              xml.Name             `xml:"project"`
	ModelVersion         string               `xml:"modelVersion"`
	Parent               ParentPOM            `xml:"parent"`
	GroupID              string               `xml:"groupId"`
	ArtifactID           string               `xml:"artifactId"`
	Version              string               `xml:"version"`
	Packaging            string               `xml:"packaging"`
	Name                 string               `xml:"name"`
	Description          string               `xml:"description"`
	Dependencies         []Dependency         `xml:"dependencies>dependency"`
	DependencyManagement DependencyManagement `xml:"dependencyManagement"`
	Properties           PropertySet          `xml:"properties"`
}

// UnmarshalPOM parses the data to a POM
func UnmarshalPOM(data []byte) (*POM, error) {
	var pom POM
	decoder := xml.NewDecoder(bytes.NewBuffer(data))
	decoder.CharsetReader = func(charset string, reader io.Reader) (io.Reader, error) {
		if charset == "ISO-8859-1" {
			return charmap.ISO8859_1.NewDecoder().Reader(reader), nil
		}
		if charset == "Windows-1252" {
			return charmap.Windows1252.NewDecoder().Reader(reader), nil
		}
		return nil, fmt.Errorf("unmarshal: unsupported charset: %s", charset)
	}

	if err := decoder.Decode(&pom); err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}
	return &pom, nil
}
