package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"

	"github.com/gofrs/uuid"
)

type root struct {
	XMLName xml.Name `xml:"root"`
	Site    site     `xml:"site"`
}

type site struct {
	Name         string `xml:"name,attr"`
	SiteNodeName string `xml:"siteNodeName,attr"`

	Nodes []node `xml:"nodes>node"`
}

func (s site) visitNodes(f func(n node)) {
	for _, n := range s.Nodes {
		n.visitNodes(f)
	}
}

type node struct {
	Identifier string `xml:"identifier,attr"`
	NodeName   string `xml:"nodeName,attr"`

	Nodes []node `xml:"node"`
}

func (n node) visitNodes(f func(n node)) {
	f(n)
	for _, n := range n.Nodes {
		n.visitNodes(f)
	}
}

func main() {
	// Read XML from stdin
	xmlData, err := ioutil.ReadAll(os.Stdin)
	checkErr(err)

	// Unmarshal XML
	var r root
	err = xml.Unmarshal(xmlData, &r)
	checkErr(err)

	// Remember all node identifiers
	nodeIdentifiers := make(map[string]string)
	r.Site.visitNodes(func(n node) {
		nodeIdentifiers[n.Identifier] = ""
	})

	// Generate new uuid for each node identifier
	for k, _ := range nodeIdentifiers {
		nodeIdentifiers[k] = uuid.Must(uuid.NewV4()).String()
	}

	// Replace node identifiers in data and output to stdout
	for oldIdentifier, newIdentifier := range nodeIdentifiers {
		xmlData = bytes.ReplaceAll(xmlData, []byte(oldIdentifier), []byte(newIdentifier))
	}
	_, _ = os.Stdout.Write(xmlData)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
