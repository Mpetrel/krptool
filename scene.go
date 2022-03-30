package krptool

import "encoding/xml"

func NewScene(name, thumburl, title string) *Scene {
	return &Scene{
		Name:     name,
		Thumburl: thumburl,
		Title:    title,
	}
}

func (s *Scene) ToXML() (string, error) {
	data, err := xml.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
