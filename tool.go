package krptool

import (
	"encoding/xml"
	"os"
)

type Tool struct {
}

func FromPath(path string) (*Krpano, error) {
	// read from file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return FromByte(data)
}

func FromXML(xmlStr string) (*Krpano, error) {
	return FromByte([]byte(xmlStr))
}

func FromByte(data []byte) (*Krpano, error) {
	var krpano Krpano
	err := xml.Unmarshal(data, &krpano)
	if err != nil {
		return nil, err
	}
	return &krpano, err
}
